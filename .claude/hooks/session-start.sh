#!/bin/bash
set -euo pipefail

# Only needed in Claude Code on the web — local dev machines already have
# their toolchains set up.
if [ "${CLAUDE_CODE_REMOTE:-}" != "true" ]; then
  exit 0
fi

cd "$CLAUDE_PROJECT_DIR"

# Go module dependencies. The root go.mod covers both the gh CLI and
# chacontainer/ (same module), so this single download primes both.
go mod download

# chacontainer/ frontend (React + Vite).
if [ -d chacontainer/web ]; then
  (cd chacontainer/web && npm install)
fi
