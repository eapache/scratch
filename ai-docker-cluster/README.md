# Searxng MCP stack

Minimal SearxNG + MCP server setup for local LLM use. SearxNG runs in Docker and is only reachable from other containers on the default compose network. The MCP server runs via stdio so RooCode can launch it reliably.

Set `COMPOSE_PROJECT_NAME=ai-docker-cluster` in `.env` (see `.env.example`) so the network name matches the config below.

## MCP client config (JSON)

Use your MCP client config file and add a stdio server entry like this:

```json
{
  "mcpServers": {
    "searxng": {
      "command": "docker",
      "args": [
        "run", "-i", "--rm",
        "--network", "ai-docker-cluster_default",
        "-e", "SEARXNG_URL=http://searxng:8080",
        "isokoliuk/mcp-searxng:latest"
      ]
    }
  }
}
```
