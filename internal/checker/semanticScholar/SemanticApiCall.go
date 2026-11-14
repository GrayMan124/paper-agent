package semanticscholar

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type SemanticScholarConfig struct {
	API_KEY string
}

func (s SemanticScholarConfig) SemanticScholarAPI(query string) error {
	baseUrl := "https://api.semanticscholar.org/graph/v1/paper/search"

	params := url.Values{}
	params.Add("query", query)
	params.Add("fields", "title,year,authors,url")
	params.Add("limit", "5")

	queryUrl := fmt.Sprintf("%s?%s", baseUrl, params.Encode())
	log.Printf("Seraching for %s using the URL:\n%s", query, queryUrl)
	req, err := http.NewRequest("GET", queryUrl, nil)
	if err != nil {
		log.Fatalf("Failed to make request with error: %v", err)
	}
	req.Header.Set("x-api-key", s.API_KEY)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	var response SemanticResult
	if resp.StatusCode != 400 {
		log.Printf("Retrieved code: {%v} from semanticScholar request", resp.StatusCode)
	}
	body, err := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Failed to Unmarshall response with error: %v", err)
		return err
	}

	for i, data := range response.Data {
		log.Printf("Data entry: {%v}\n ID: %v\nTitle:%s,URL:%v\nAuthors:%v", i, data.PaperID, data.Title, data.URL, data.Authors)
	}
	return nil

}
