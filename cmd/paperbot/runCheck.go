package main

import (
	"context"
	"log"

	"github.com/GrayMan124/paper-agent/internal/checker"
	semanticscholar "github.com/GrayMan124/paper-agent/internal/checker/semanticScholar"
	"github.com/GrayMan124/paper-agent/internal/sources/scholar"
)

func runCheck(ctx context.Context, args []string, semSearch *semanticscholar.SemanticScholarConfig, schol scholar.Scholar) error {
	candidates, err := ScholarResults(schol, ctx, args)
	log.Printf("Number of candidates retrieved: %v", len(candidates))
	if err != nil {
		return err
	}

	for i, cand := range candidates {
		log.Printf("Candidate: %v\nTitle:%s\nAuthors:%s\nYear:%v\nURL:%s\n\n", i, cand.Title, cand.AuthorsRaw, *cand.Year, *cand.Link)
	}
	//NOTE: I think it would be good if the LLM agent responsible for seraching could check the results here before passing them into contract checking

	results, err := checker.CheckScholarResults(candidates, semSearch)

	if err != nil {
		return err
	}

	for _, res := range results {
		log.Printf("ResultID: %v\nStatus:%s\nReason:%s\nTitle:%s\n\n", res.ResultID, res.Status, res.Reason, res.Canonical.Title)
	}
	return nil
}
