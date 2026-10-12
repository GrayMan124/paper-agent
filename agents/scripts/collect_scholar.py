import asyncio
import json
from pathlib import Path
from pdb import run
from fastmcp import Client
from datetime import datetime
from zoneinfo import ZoneInfo

CONFIG = Path(__file__).resolve().parents[1] / "mcp_servers.json"
QUERIES = [
    ("q01_ph_llm_hidden", "persistent homology LLM hidden states"),
    ("q02_graph_llm_hidden", "laplacian graph LLM hidden states"),
    ("q03_intr_dim_llm_adv", "intrinsic dimension LLM robustness"),
    ("q04_llm_geometry_methods", "LLM analysis geometry methods"),
]
OUT = Path(__file__).resolve().parents[2] / "fixtures" / "scholar"
OUT.mkdir(parents=True, exist_ok=True)


async def scholar_search(slag: str, query: str, client: Client) -> None:
    try:
        tool = "search_google_scholar_key_words"

        # async with Client(config) as client:
        result = await client.call_tool(
            tool,
            {"query": "LLM analysis using graph", "num_results": 5},
        )
        record = {
            "query": query,
            "tool": tool,
            "fetched_at": datetime.now(ZoneInfo("Asia/Tokyo")).isoformat(),
            "content": [b.model_dump(mode="json") for b in result.content],
            "structured_content": result.structured_content,
        }
        with open(OUT / f"{slag}.json", "w") as f:
            json.dump(record, f, indent=2, ensure_ascii=False)
    except BaseException as e:
        print(f"Failed to get results: {e}")


async def main():
    config = json.loads(CONFIG.read_text())
    async with Client(config) as client:
        for slaq, query in QUERIES:
            await scholar_search(slag=slaq, query=query, client=client)


if __name__ == "__main__":
    asyncio.run(main())
