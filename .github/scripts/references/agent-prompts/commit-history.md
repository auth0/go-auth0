Get the commit history for PR #{{PR_NUMBER}}.

First, get the commit list via GitHub API (works from any branch):
  gh pr view {{PR_NUMBER}} --json commits --jq '.commits[] | "\(.oid[:7]) \(.messageHeadline)"'

Then get full commit messages with bodies:
  gh pr view {{PR_NUMBER}} --json commits --jq '.commits[] | "\(.oid[:7]) \(.messageHeadline)\n\(.messageBody)\n---"'

Summarize:
1. Total number of commits
2. List of commit messages (short form)
3. Overall theme/pattern of the commits
4. Any conventional commit prefixes already used

Return a structured summary.
