# symphony-mcp

MCP server that wraps Symphony's `ask`/`reply` HTTP API so a **local** Claude
Code session can call it as a tool. This is separate from CHACONTAINER's own
Go integration (`chacontainer/internal/integrations/symphony`), which is what
the CHACONTAINER backend itself uses — this one is for talking to Symphony
directly from your terminal.

## Setup

```bash
cd chacontainer/tools/symphony-mcp
npm install
```

## Register it with Claude Code

Run this on **your own machine** (not in a remote session) — `claude mcp add`
needs a local interactive CLI:

```bash
claude mcp add symphony \
  --env SYMPHONY_API_TOKEN=your-token-here \
  -- node /absolute/path/to/chacontainer/tools/symphony-mcp/index.js
```

Never commit your token. Get/rotate it from Symphony → Settings → Connections.

Optional env vars:
- `SYMPHONY_BASE_URL` — override the API base (default: the production Symphony endpoint).
- `SYMPHONY_SESSION_ID` — keep replies in one thread (default: `chacontainer-mcp`).

## What it exposes

One tool, `symphony_ask(message)`: sends `message` to Symphony and blocks
(polling internally) until Symphony replies, up to ~2 minutes.

**A message sent through this tool carries your full authority to act** —
Symphony can schedule, contact people, and follow up as you. Only send
messages you actually intend to send.
