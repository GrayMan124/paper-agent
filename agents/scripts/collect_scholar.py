import asyncio
import json
from pathlib import Path
from pdb import run
from fastmcp import Client

CONFIG = Path(__file__).resolve().parents[1] / "mcp_servers.json"


async def main():
    config = json.loads(CONFIG.read_text())
    async with Client(config) as client:
        # tools = await client.list_tools()
        result = await client.call_tool(
            "search_google_scholar_key_words",
            {"query": "LLM analysis using graph", "num_results": 5},
        )
    for block in result.content:
        print(block.text if hasattr(block, "text") else block)
    print("Structured:", result.structured_content)


if __name__ == "__main__":
    asyncio.run(main())
