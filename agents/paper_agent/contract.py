from typing import Literal
from pydantic import BaseModel, ConfigDict


class MetaData(BaseModel):
    title: str
    authors: list[str]
    year: int | None
    doi: str | None
    arxiv_id: str | None
    venue: str | None
    abstract: str | None


class Candidate(BaseModel):
    model_config = ConfigDict(from_attributes=True, extra="forbid")
    result_id: str
    title: str
    authors_raw: str
    year: int | None
    link: str | None
    doi: str | None
    arxiv_id: str | None


class Result(BaseModel):
    # model_config = ConfigDict(from_attributes=True, extra="forbid")
    result_id: str
    status: Literal["verified", "metadata_mismatch", "not_found", "api_error"]
    reason: str
    canonical: MetaData | None
