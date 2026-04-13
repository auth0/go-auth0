Do these tasks:

1. Check for linked issues using TWO methods:

   a. Get GitHub's native linked issues:
      gh pr view {{PR_NUMBER}} --json closingIssuesReferences --jq '.closingIssuesReferences[]'

   b. Also scan the current PR body for additional issue references (#123 patterns):
      ```
      {{PR_BODY}}
      ```

   For each unique issue number found (from either method), run:
      gh issue view <number> --json title,body,labels,state
   Summarize each linked issue.

2. Search for a PR template in the repo. Use `git rev-parse --show-toplevel` to get the repo root.

   Check for a **single template** at these paths (in order, return the FIRST one found):
   - .github/pull_request_template.md
   - .github/PULL_REQUEST_TEMPLATE.md
   - pull_request_template.md
   - PULL_REQUEST_TEMPLATE.md
   - docs/pull_request_template.md
   - docs/PULL_REQUEST_TEMPLATE.md

   If none found, check for a **multi-template directory** at these paths:
   - .github/PULL_REQUEST_TEMPLATE/
   - PULL_REQUEST_TEMPLATE/
   - docs/PULL_REQUEST_TEMPLATE/
   If a directory exists, list the `.md` files inside it and return the first one
   (or note that multiple templates exist so the user can choose).

   If a template is found, read its full content and analyze its structure:
   - List every section heading
   - Note any example/placeholder text (lines in HTML comments, italic text, or content
     after phrases like "e.g.", "Example:", "For example", "such as")
   - Note the formatting style used in examples (bullets vs prose, short vs detailed)
   - Note any special fields, dropdowns (`<details>`), or custom sections beyond the standard ones
   - Note if the template uses checkboxes, tables, or other structured formats

   Return the template path AND this structural analysis.
   If no template exists, return "NONE".

3. Check for CLAUDE.md at the repo root. If it exists, note any PR-related guidelines.

Return all findings in a structured summary.
