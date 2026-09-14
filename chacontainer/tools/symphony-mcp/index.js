#!/usr/bin/env node
// MCP server that wraps Symphony's ask/reply HTTP API so a local Claude Code
// session can talk to it as a tool (registered via `claude mcp add`).
//
// A message sent through Symphony is treated as coming from the account
// holder with full authority to act (scheduling, contacting people, etc.).
// This server never invents that authorization — it only relays whatever
// message the caller explicitly provides.

import { McpServer } from "@modelcontextprotocol/sdk/server/mcp.js";
import { StdioServerTransport } from "@modelcontextprotocol/sdk/server/stdio.js";
import { z } from "zod";

const DEFAULT_BASE_URL = "https://symphony.wix.com/individuals-chat/poc/agent";
const DEFAULT_SESSION_ID = "chacontainer-mcp";
const DEFAULT_MAX_WAIT_MS = 110_000; // Symphony's own turns can take a minute or two.
const DEFAULT_POLL_INTERVAL_MS = 3_000;

const baseURL = (process.env.SYMPHONY_BASE_URL || DEFAULT_BASE_URL).replace(/\/$/, "");
const sessionID = process.env.SYMPHONY_SESSION_ID || DEFAULT_SESSION_ID;
const apiToken = process.env.SYMPHONY_API_TOKEN;

if (!apiToken) {
  console.error(
    "symphony-mcp: SYMPHONY_API_TOKEN is not set. Configure it when registering " +
      "this server (e.g. `claude mcp add symphony --env SYMPHONY_API_TOKEN=... -- node index.js`) " +
      "rather than hardcoding it here."
  );
  process.exit(1);
}

async function postJSON(path, body) {
  const res = await fetch(`${baseURL}${path}`, {
    method: "POST",
    headers: {
      Authorization: `Bearer ${apiToken}`,
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
  });
  if (!res.ok) {
    const text = await res.text().catch(() => "");
    throw new Error(`symphony ${res.status}: ${text}`);
  }
  return res.json();
}

async function ask(message) {
  return postJSON("/ask", { message, sessionId: sessionID });
}

async function reply(conversationId) {
  return postJSON("/reply", { conversationId });
}

function sleep(ms) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

/**
 * Sends message to Symphony and blocks until a reply is available, polling
 * as needed, mirroring the ask/reply contract Symphony documents.
 */
async function askAndWait(message, maxWaitMs = DEFAULT_MAX_WAIT_MS, pollIntervalMs = DEFAULT_POLL_INTERVAL_MS) {
  const first = await ask(message);
  if (first.reply) {
    return first.reply;
  }
  if (first.status !== "accepted" || !first.conversationId) {
    throw new Error(`symphony: unexpected response ${JSON.stringify(first)}`);
  }

  const deadline = Date.now() + maxWaitMs;
  while (true) {
    await sleep(pollIntervalMs);
    const r = await reply(first.conversationId);
    if (r.status === "answered" && r.reply) {
      return r.reply;
    }
    if (Date.now() > deadline) {
      throw new Error("symphony: timed out waiting for a reply");
    }
  }
}

const server = new McpServer({ name: "symphony", version: "1.0.0" });

server.registerTool(
  "symphony_ask",
  {
    description:
      "Send a message to Symphony, the user's team of AI business agents. " +
      "A message sent here carries the account holder's full authority to act " +
      "(it can schedule, reach contacts, follow up, and build reports on their " +
      "behalf) — only send what the user actually asked for, never an inference " +
      "or a guess at what they might want. Blocks until Symphony replies " +
      "(can take up to ~2 minutes).",
    inputSchema: {
      message: z.string().describe("The exact message to send to Symphony, as the user intends it."),
    },
  },
  async ({ message }) => {
    try {
      const text = await askAndWait(message);
      return { content: [{ type: "text", text }] };
    } catch (err) {
      return {
        content: [{ type: "text", text: `Error talking to Symphony: ${err.message}` }],
        isError: true,
      };
    }
  }
);

const transport = new StdioServerTransport();
await server.connect(transport);
console.error("symphony-mcp server running on stdio");
