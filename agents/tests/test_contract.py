from pathlib import Path
import pytest
from paper_agent.contract import Candidate, Result

SCHEMA = Path(__file__).resolve().parents[2] / "schema"


def test_not_found_decode():
    # print((SCHEMA / "result_example_corr.json").read_text())
    r = Result.model_validate_json((SCHEMA / "result_failed_nf.json").read_text())
    assert r.status == "not_found"
    assert r.canonical is None


def test_candidate_decode():
    c = Candidate.model_validate_json((SCHEMA / "candidate_example.json").read_text())
    assert c.year == 2017
    assert c.arxiv_id is None


def test_corr_responde_decode():
    r = Result.model_validate_json((SCHEMA / "result_example_corr.json").read_text())
    assert r.status == "verified"
    assert r.canonical is not None
