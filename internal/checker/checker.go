package checker

import (
	semanticscholar "github.com/GrayMan124/paper-agent/internal/checker/semanticScholar"
	"github.com/GrayMan124/paper-agent/internal/contract"
	"log"
	"strconv"
	"strings"
	"sync"
)

type compareResult struct {
	result string
	reason string
}

func CheckScholarResults(candidates []contract.Candidate,
	semanticScholarConfig *semanticscholar.SemanticScholarConfig) ([]contract.Result, error) {

	semSearchResults := make([]semanticscholar.SemanticResult, len(candidates))

	var wg sync.WaitGroup

	for i, sh := range candidates {
		wg.Go(func() {
			semSearchResult, err := semanticScholarConfig.SemanticScholarAPI(sh.Title)
			if err != nil {
				log.Printf("Failed to retrieve id: %v with title: %v\ndue to error: %v", i, sh.Title, err)
				return
			}
			semSearchResults[i] = semSearchResult
		})
	}
	wg.Wait()

	var results []contract.Result

	for i, c := range candidates {
		var res contract.Result
		id, compRes := compare(c, semSearchResults[i])
		if compRes.result == "accept" {
			res = contract.Result{
				ResultID:  strconv.Itoa(i),
				Status:    compRes.result,
				Reason:    compRes.reason,
				Relevance: -1,
				Canonical: contract.Canonical{
					Title:    c.Title,
					Authors:  nil,
					Abstract: &semSearchResults[i].Data[id].Abstract,
				},
			}
		} else {
			res = contract.Result{
				ResultID:  strconv.Itoa(i),
				Status:    compRes.result,
				Reason:    compRes.reason,
				Relevance: -1,
				Canonical: contract.Canonical{
					Title:   c.Title,
					Authors: nil,
				},
			}
		}
		results = append(results, res)
	}
	return results, nil
}

func compare(cand contract.Candidate, semHit semanticscholar.SemanticResult) (int, compareResult) {
	if len(semHit.Data) == 0 {
		return -1, compareResult{result: "reject", reason: "No matches found"}
	}

	for semId, semData := range semHit.Data {
		if strings.EqualFold(cand.Title, semData.Title) {
			return semId, compareResult{result: "accept", reason: "Title found"}
		}
	}
	return -1, compareResult{result: "reject", reason: "Mismatched titles"}
}
