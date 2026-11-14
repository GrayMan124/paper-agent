package semanticscholar

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const MAX_ATTEMPTS = 5

type SemanticScholarConfig struct {
	API_KEY string
	Client  *http.Client
	Muc     sync.Mutex
	LastReq time.Time
}

func (s SemanticScholarConfig) SemanticScholarAPI(query string) (SemanticResult, error) {
	baseUrl := "https://api.semanticscholar.org/graph/v1/paper/search"

	params := url.Values{}
	params.Add("query", query)
	params.Add("fields", "title,year,authors,url,abstract")
	params.Add("limit", "5")
	queryUrl := fmt.Sprintf("%s?%s", baseUrl, params.Encode())
	log.Printf("Seraching for %s using the URL:\n%s", query, queryUrl)

	var resp *http.Response

	// Spacing the requests appart, and also allowing for several attempts
	for attempt := 0; attempt <= MAX_ATTEMPTS; attempt++ {
		s.Muc.Lock()
		elapsed_time := time.Since(s.LastReq)
		if elapsed_time < 1100*time.Millisecond {
			time.Sleep(1100*time.Millisecond - elapsed_time)
		}
		s.LastReq = time.Now()
		s.Muc.Unlock()

		req, err := http.NewRequest("GET", queryUrl, nil)
		if err != nil {
			log.Fatalf("Failed to create request with error: %v", err)
		}

		req.Header.Set("x-api-key", s.API_KEY)

		resp, err = s.Client.Do(req)
		if err != nil {
			log.Fatalf("Failed to send the request due to error: %v", err)
			return SemanticResult{}, err
		}
		if resp.StatusCode != 200 {
			log.Printf("Retrieved code: {%v} from semanticScholar request", resp.StatusCode)
			continue
		}
		break

	}
	var response SemanticResult
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read the body due to error: %v", err)
		return SemanticResult{}, err
	}
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatalf("Failed to Unmarshall response with error: %v", err)
		return SemanticResult{}, err
	}
	for i, data := range response.Data {
		log.Printf("Data entry: {%v}\n ID: %v\nTitle:%s,URL:%v\nAuthors:%v\nAbstract:%s", i, data.PaperID, data.Title, data.URL, data.Authors, data.Abstract)
	}
	return response, nil

}
