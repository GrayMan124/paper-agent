package semanticscholar

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func SemanticScholarAPI(query string) error {

	baseUrl := "https://api.semanticscholar.org/graph/v1/paper/search"
	queryUrl := fmt.Sprintf("%s?query=%s&fields=title,year,authors", baseUrl, query)
	log.Printf("Seraching for %s using the URL:\n%s", query, queryUrl)
	resp, err := http.Get(queryUrl)
	if err != nil {
		return err
	}
	var response any
	body, err := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Failed to Unmarshall response with error: %v", err)
		return err
	}
	log.Printf("Response:\n%v", response)
	return nil

}
