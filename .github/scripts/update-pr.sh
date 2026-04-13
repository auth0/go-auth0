#!/usr/bin/env bash
# update-pr.sh — Generates and applies AI-powered PR title & description.
# Uses GITHUB_TOKEN to authenticate with GitHub Models API — no PAT needed.
set -euo pipefail

# ---------------------------------------------------------------------------
# Configuration
# ---------------------------------------------------------------------------
MAX_DIFF_CHARS=40000
API_URL="https://models.inference.ai.azure.com/chat/completions"
MODEL="${MODEL:-claude-opus-4.6}"
PR_NUMBER="${PR_NUMBER:?PR_NUMBER is required}"
SKIP_TITLE="${SKIP_TITLE:-false}"
SKIP_DESCRIPTION="${SKIP_DESCRIPTION:-false}"
DRY_RUN="${DRY_RUN:-false}"

if [ -z "${GITHUB_TOKEN:-}" ]; then
  echo "❌ GITHUB_TOKEN is not set."
  exit 1
fi

# ---------------------------------------------------------------------------
# Security: Redact secrets from diff before sending to AI
# ---------------------------------------------------------------------------

# Redact secrets from a string. Replaces matches with [REDACTED].
# Uses a temp file to handle large inputs safely.
sanitize_secrets() {
  local tmpfile
  tmpfile=$(mktemp)
  trap "rm -f '$tmpfile'" RETURN

  # Write input to temp file to avoid shell argument-length limits
  cat > "$tmpfile"

  # Redact explicit secret assignments (KEY=value, "key": "value", etc.)
  sed -i -E \
    -e 's/(password|passwd|secret|token|api_key|apikey|api-key|credential|private_key|access_key|client_secret)[[:space:]]*[:=][[:space:]]*["'"'"']?[^[:space:]"'"'"']{8,}/\1=[REDACTED]/gI' \
    -e 's/(Bearer |Basic |Authorization:[[:space:]]*)[A-Za-z0-9_.+\/=-]{20,}/\1[REDACTED]/gI' \
    "$tmpfile"

  # Redact well-known token formats
  sed -i -E \
    -e 's/AKIA[0-9A-Z]{16}/[REDACTED]/g' \
    -e 's/ghp_[A-Za-z0-9]{36}/[REDACTED]/g' \
    -e 's/gho_[A-Za-z0-9]{36}/[REDACTED]/g' \
    -e 's/ghs_[A-Za-z0-9]{36}/[REDACTED]/g' \
    -e 's/github_pat_[A-Za-z0-9_]{60,}/[REDACTED]/g' \
    -e 's/xox[baprs]-[A-Za-z0-9-]{10,}/[REDACTED]/g' \
    -e 's/sk-[A-Za-z0-9]{20,}/[REDACTED]/g' \
    -e 's/sk_live_[A-Za-z0-9]{20,}/[REDACTED]/g' \
    -e 's/sq0[a-z]{3}-[A-Za-z0-9_-]{22,}/[REDACTED]/g' \
    -e 's/eyJ[A-Za-z0-9_-]{20,}\.[A-Za-z0-9_-]{20,}/[REDACTED]/g' \
    "$tmpfile"

  cat "$tmpfile"
}

# Validate the API endpoint is the expected GitHub Models host.
# Prevents redirect-based exfiltration.
validate_api_url() {
  local url="$1"
  if [[ "$url" != "https://models.inference.ai.azure.com/"* ]]; then
    echo "❌ Security: API_URL must be https://models.inference.ai.azure.com/. Got: $url"
    exit 1
  fi
}
validate_api_url "$API_URL"

# ---------------------------------------------------------------------------
# 1. Gather PR details
# ---------------------------------------------------------------------------
echo "🔍 Gathering context for PR #${PR_NUMBER}..."

PR_JSON=$(gh pr view "$PR_NUMBER" --json number,title,body,baseRefName,headRefName,url,closingIssuesReferences)
PR_TITLE=$(echo "$PR_JSON" | jq -r '.title')
PR_BODY=$(echo "$PR_JSON" | jq -r '.body // ""')
BASE_BRANCH=$(echo "$PR_JSON" | jq -r '.baseRefName')
HEAD_BRANCH=$(echo "$PR_JSON" | jq -r '.headRefName')
PR_URL=$(echo "$PR_JSON" | jq -r '.url')

echo "  PR:     $PR_URL"
echo "  Title:  $PR_TITLE"
echo "  Branch: $HEAD_BRANCH → $BASE_BRANCH"

# ---------------------------------------------------------------------------
# 2. Comprehensive change detection
# ---------------------------------------------------------------------------

# 2a. Patch-format diff (includes rename/copy detection, binary file info)
RAW_DIFF=$(gh pr diff "$PR_NUMBER" --patch 2>/dev/null || gh pr diff "$PR_NUMBER" 2>/dev/null || echo "Could not fetch diff")

# Security: scrub secrets from the diff before sending to AI
DIFF=$(sanitize_secrets "$RAW_DIFF")
if [ ${#DIFF} -gt "$MAX_DIFF_CHARS" ]; then
  DIFF="${DIFF:0:$MAX_DIFF_CHARS}

... [diff truncated at ${MAX_DIFF_CHARS} characters] ..."
  echo "  ⚠️  Diff truncated to ${MAX_DIFF_CHARS} characters"
fi

# 2b. File list with change status (Added/Modified/Deleted/Renamed)
FILE_CHANGES=$(gh pr diff "$PR_NUMBER" --name-only 2>/dev/null || echo "Could not fetch file list")

# 2c. Diff stat summary (insertions/deletions per file + total)
PR_FILES_JSON=$(gh pr view "$PR_NUMBER" --json files 2>/dev/null || echo '{"files":[]}')
DIFF_STAT=$(echo "$PR_FILES_JSON" | jq -r '.files[] | "\(.path)\t+\(.additions)\t-\(.deletions)"' 2>/dev/null || echo "Could not fetch diff stats")

TOTAL_ADDITIONS=$(echo "$PR_FILES_JSON" | jq '[.files[]?.additions // 0] | add // 0' 2>/dev/null || echo "?")
TOTAL_DELETIONS=$(echo "$PR_FILES_JSON" | jq '[.files[]?.deletions // 0] | add // 0' 2>/dev/null || echo "?")
TOTAL_FILES=$(echo "$PR_FILES_JSON" | jq '.files | length' 2>/dev/null || echo "?")

echo "  📊 ${TOTAL_FILES} files changed, +${TOTAL_ADDITIONS}/-${TOTAL_DELETIONS} lines"

# 2d. Detect renames, binary changes, and large file modifications
RENAMES=$(echo "$DIFF" | grep -E "^(rename from|rename to|similarity index)" 2>/dev/null || true)
BINARY_FILES=$(echo "$DIFF" | grep -E "^(Binary files|GIT binary patch)" 2>/dev/null || true)

# 2e. PR metadata (labels, reviewers, milestone)
PR_LABELS=$(gh pr view "$PR_NUMBER" --json labels --jq '[.labels[].name] | join(", ")' 2>/dev/null || echo "")
PR_MILESTONE=$(gh pr view "$PR_NUMBER" --json milestone --jq '.milestone.title // ""' 2>/dev/null || echo "")

# ---------------------------------------------------------------------------
# 3. Commit history
# ---------------------------------------------------------------------------
COMMITS=$(gh pr view "$PR_NUMBER" --json commits \
  --jq '.commits[] | "\(.oid[:7]) \(.messageHeadline)"' 2>/dev/null || echo "No commits found")

COMMIT_BODIES=$(gh pr view "$PR_NUMBER" --json commits \
  --jq '.commits[] | "### \(.oid[:7]) \(.messageHeadline)\n\(.messageBody)\n"' 2>/dev/null || echo "")

# ---------------------------------------------------------------------------
# 4. Linked issues
# ---------------------------------------------------------------------------
ISSUE_NUMBERS=$(echo "$PR_JSON" | jq -r '.closingIssuesReferences[].number' 2>/dev/null || true)
ISSUE_DETAILS=""

if [ -n "$ISSUE_NUMBERS" ]; then
  while IFS= read -r issue_num; do
    [ -z "$issue_num" ] && continue
    detail=$(gh issue view "$issue_num" --json title,body,labels,state 2>/dev/null || echo "Could not fetch issue #$issue_num")
    ISSUE_DETAILS="${ISSUE_DETAILS}
--- Issue #${issue_num} ---
${detail}
"
  done <<< "$ISSUE_NUMBERS"
fi

# Scan PR body for #NNN references not already captured.
BODY_ISSUES=$(echo "$PR_BODY" | grep -oE '#[0-9]+' | sort -u || true)
if [ -n "$BODY_ISSUES" ]; then
  while IFS= read -r ref; do
    num="${ref#\#}"
    [ -z "$num" ] && continue
    if echo "$ISSUE_NUMBERS" | grep -qw "$num" 2>/dev/null; then continue; fi
    detail=$(gh issue view "$num" --json title,body,labels,state 2>/dev/null || true)
    if [ -n "$detail" ]; then
      ISSUE_DETAILS="${ISSUE_DETAILS}
--- Issue #${num} (referenced in body) ---
${detail}
"
    fi
  done <<< "$BODY_ISSUES"
fi

# ---------------------------------------------------------------------------
# 5. PR template
# ---------------------------------------------------------------------------
TEMPLATE=""
for tmpl_path in \
  ".github/PULL_REQUEST_TEMPLATE.md" \
  ".github/pull_request_template.md" \
  "PULL_REQUEST_TEMPLATE.md" \
  "pull_request_template.md" \
  "docs/pull_request_template.md"; do
  if [ -f "$tmpl_path" ]; then
    TEMPLATE=$(cat "$tmpl_path")
    echo "  📄 Found PR template: $tmpl_path"
    break
  fi
done

if [ -z "$TEMPLATE" ]; then
  TEMPLATE='## Summary

-

## Changes

## Related Issues

## Testing

## Breaking Changes

## Checklist

- [ ] Self-review completed
- [ ] Tests added/updated
- [ ] Documentation updated (if applicable)'
  echo "  📄 Using default PR template"
fi

echo ""
echo "🤖 Calling ${MODEL} via GitHub Models API..."

# ---------------------------------------------------------------------------
# 6. Build the AI prompt
# ---------------------------------------------------------------------------
# ---------------------------------------------------------------------------
# 6a. Load reference materials from bundled skill files
# ---------------------------------------------------------------------------
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REF_DIR="${SCRIPT_DIR}/references"

CONV_COMMITS=""
if [ -f "${REF_DIR}/conventional-commits.md" ]; then
  CONV_COMMITS=$(cat "${REF_DIR}/conventional-commits.md")
  echo "  📖 Loaded conventional-commits.md"
fi

DIFF_ANALYSIS_PROMPT=""
if [ -f "${REF_DIR}/agent-prompts/diff-analysis.md" ]; then
  DIFF_ANALYSIS_PROMPT=$(cat "${REF_DIR}/agent-prompts/diff-analysis.md")
fi

CODE_ANALYSIS_PROMPT=""
if [ -f "${REF_DIR}/agent-prompts/code-analysis.md" ]; then
  CODE_ANALYSIS_PROMPT=$(cat "${REF_DIR}/agent-prompts/code-analysis.md")
fi

read -r -d '' SYSTEM_PROMPT << SYSTEMPROMPT || true
You are an expert at analyzing pull request changes and generating clear, structured PR titles and descriptions.

## Conventional Commits Reference
${CONV_COMMITS}

## Diff Analysis Guidelines
${DIFF_ANALYSIS_PROMPT}

## Code Analysis Guidelines
${CODE_ANALYSIS_PROMPT}

## Output Format
You MUST respond with EXACTLY this format (no other text before or after):

<<<TITLE>>>
your conventional commits title here
<<<DESCRIPTION>>>
your full PR description here using the template
<<<END>>>

## Rules for the Description
- Use the provided PR template as the structure.
- Every section in the template MUST appear with real content.
- Replace all placeholder/example text (including HTML comments) with actual content derived from the diff.
- Use "N/A" for truly inapplicable sections — never delete a section.
- Reference actual filenames, function names, and code from the diff.
- Check checkboxes `[x]` that the diff confirms are done; leave others unchecked `[ ]`.
- Be specific and technical, not generic.
- Strip all HTML comments (<!-- -->) from your output.
SYSTEMPROMPT

USER_PROMPT="Analyze this pull request and generate a title and description.

## PR Details
- Number: #${PR_NUMBER}
- Current Title: ${PR_TITLE}
- Head Branch: ${HEAD_BRANCH}
- Base Branch: ${BASE_BRANCH}
- Labels: ${PR_LABELS:-None}
- Milestone: ${PR_MILESTONE:-None}

## Change Statistics
- Files changed: ${TOTAL_FILES}
- Lines added: +${TOTAL_ADDITIONS}
- Lines deleted: -${TOTAL_DELETIONS}

## Files Changed (with stats)
${DIFF_STAT}

## File List
${FILE_CHANGES}

## Renames Detected
${RENAMES:-None}

## Binary File Changes
${BINARY_FILES:-None}

## Commit History
${COMMITS}

## Commit Details
${COMMIT_BODIES}

## Linked Issues
${ISSUE_DETAILS:-None}

## PR Template to Fill
${TEMPLATE}

## Full Patch Diff
\`\`\`diff
${DIFF}
\`\`\`

Generate the title and description now. Use the exact <<<TITLE>>>, <<<DESCRIPTION>>>, <<<END>>> delimiters."

# ---------------------------------------------------------------------------
# 7. Call GitHub Models API using GITHUB_TOKEN
# ---------------------------------------------------------------------------
PAYLOAD=$(jq -n \
  --arg model "$MODEL" \
  --arg system "$SYSTEM_PROMPT" \
  --arg user "$USER_PROMPT" \
  '{
    model: $model,
    messages: [
      {role: "system", content: $system},
      {role: "user", content: $user}
    ],
    max_tokens: 8192,
    temperature: 0.3
  }')

MAX_RETRIES=3
RETRY=0
HTTP_CODE=""
RESPONSE_BODY=""

while [ "$RETRY" -lt "$MAX_RETRIES" ]; do
  RESPONSE=$(curl -s -w "\n%{http_code}" "$API_URL" \
    -H "Authorization: Bearer $GITHUB_TOKEN" \
    -H "Content-Type: application/json" \
    -d "$PAYLOAD" 2>&1) || true

  HTTP_CODE=$(echo "$RESPONSE" | tail -1)
  RESPONSE_BODY=$(echo "$RESPONSE" | sed '$d')

  if [ "$HTTP_CODE" = "200" ]; then
    break
  fi

  RETRY=$((RETRY + 1))
  if [ "$RETRY" -lt "$MAX_RETRIES" ]; then
    WAIT=$((RETRY * 5))
    echo "  ⚠️  API returned HTTP $HTTP_CODE, retrying in ${WAIT}s (attempt $((RETRY + 1))/$MAX_RETRIES)..."
    sleep "$WAIT"
  fi
done

if [ "$HTTP_CODE" != "200" ]; then
  echo "❌ GitHub Models API request failed after $MAX_RETRIES attempts (HTTP $HTTP_CODE)"
  # Only show error message, not full response (may contain tokens in headers)
  echo "$RESPONSE_BODY" | jq -r '.error.message // .error // "Unknown error"' 2>/dev/null || echo "Unknown error"
  exit 1
fi

# ---------------------------------------------------------------------------
# 8. Parse the AI response
# ---------------------------------------------------------------------------
AI_CONTENT=$(echo "$RESPONSE_BODY" | jq -r '.choices[0].message.content // empty')

if [ -z "$AI_CONTENT" ]; then
  echo "❌ Empty response from API"
  exit 1
fi

NEW_TITLE=$(echo "$AI_CONTENT" | awk '/^<<<TITLE>>>/{flag=1; next} /^<<<DESCRIPTION>>>/{flag=0} flag' | sed '/^$/d' | head -1)
NEW_DESCRIPTION=$(echo "$AI_CONTENT" | awk '/^<<<DESCRIPTION>>>/{flag=1; next} /^<<<END>>>/{flag=0} flag')

if [ -z "$NEW_TITLE" ]; then
  echo "❌ Could not parse title from AI response"
  exit 1
fi

if [ -z "$NEW_DESCRIPTION" ]; then
  echo "❌ Could not parse description from AI response"
  exit 1
fi

# Append footer.
NEW_DESCRIPTION="${NEW_DESCRIPTION}

---
*PR metadata updated by [GitHub Actions](${GITHUB_SERVER_URL:-https://github.com}/${GITHUB_REPOSITORY:-unknown}/actions/runs/${GITHUB_RUN_ID:-0}) using ${MODEL}*"

# ---------------------------------------------------------------------------
# Security: Validate AI output before applying
# ---------------------------------------------------------------------------
# 1. Sanitize the AI output (in case the model echoed back secrets from context)
NEW_TITLE=$(sanitize_secrets "$NEW_TITLE")
NEW_DESCRIPTION=$(sanitize_secrets "$NEW_DESCRIPTION")

# 2. Block if AI output contains suspicious patterns (e.g., prompt injection artifacts)
if echo "$NEW_TITLE" | grep -qiE '(<<<|>>>|SYSTEM|ASSISTANT|INSTRUCTION)'; then
  echo "⚠️  Title contains suspicious prompt artifacts, sanitizing..."
  NEW_TITLE=$(echo "$NEW_TITLE" | sed -E 's/(<<<|>>>|SYSTEM|ASSISTANT|INSTRUCTION)//gi' | xargs)
fi

# 3. Enforce title length limit
if [ ${#NEW_TITLE} -gt 100 ]; then
  echo "⚠️  Title too long (${#NEW_TITLE} chars), truncating to 100..."
  NEW_TITLE="${NEW_TITLE:0:100}"
fi

# 4. Enforce description size limit (prevent massive AI output)
MAX_DESC_CHARS=50000
if [ ${#NEW_DESCRIPTION} -gt "$MAX_DESC_CHARS" ]; then
  echo "⚠️  Description too long (${#NEW_DESCRIPTION} chars), truncating to ${MAX_DESC_CHARS}..."
  NEW_DESCRIPTION="${NEW_DESCRIPTION:0:$MAX_DESC_CHARS}

... [truncated]"
fi

echo ""
echo "✅ Generated PR metadata:"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
echo "  Title: $NEW_TITLE"
echo "  Description: ${#NEW_DESCRIPTION} characters"
echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"

# ---------------------------------------------------------------------------
# 9. Dry-run check
# ---------------------------------------------------------------------------
if [ "$DRY_RUN" = "true" ]; then
  echo ""
  echo "🏜️  DRY RUN — not updating the PR. Generated content:"
  echo ""
  echo "=== TITLE ==="
  echo "$NEW_TITLE"
  echo ""
  echo "=== DESCRIPTION ==="
  echo "$NEW_DESCRIPTION"
  exit 0
fi

# ---------------------------------------------------------------------------
# 10. Update the PR
# ---------------------------------------------------------------------------
if [ "$SKIP_TITLE" = "true" ] && [ "$SKIP_DESCRIPTION" = "true" ]; then
  echo "⚠️  Both title and description updates were skipped"
  exit 0
fi

if [ "$SKIP_TITLE" != "true" ] && [ "$SKIP_DESCRIPTION" != "true" ]; then
  echo "📝 Updating title and description..."
  echo "$NEW_DESCRIPTION" | gh pr edit "$PR_NUMBER" --title "$NEW_TITLE" --body-file -
elif [ "$SKIP_TITLE" != "true" ]; then
  echo "📝 Updating title only..."
  gh pr edit "$PR_NUMBER" --title "$NEW_TITLE"
else
  echo "📝 Updating description only..."
  echo "$NEW_DESCRIPTION" | gh pr edit "$PR_NUMBER" --body-file -
fi

echo ""
echo "🎉 PR #${PR_NUMBER} updated successfully!"
echo "  URL: $PR_URL"
echo "  New title: $NEW_TITLE"
