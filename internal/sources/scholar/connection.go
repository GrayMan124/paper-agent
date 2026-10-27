package scholar

import (
	"context"
	"encoding/json"
	"log"
	"os/exec"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Scholar struct {
	Config ScholarConfig
}

func (s Scholar) Search(ctx context.Context, query string, n int) ([]ScholarHit, error) {
	client := mcp.NewClient(&mcp.Implementation{
		Name:    "paperbot",
		Version: "0.0.1",
	}, nil)
	cmd := exec.Command(s.Config.McpServers.GoogleScholar.Command, s.Config.McpServers.GoogleScholar.Args...)
	session, err := client.Connect(ctx, &mcp.CommandTransport{Command: cmd}, nil)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer session.Close()
	params := &mcp.CallToolParams{
		Name:      "search_google_scholar_key_words",
		Arguments: map[string]any{"query": query, "num_results": n},
	}
	res, err := session.CallTool(ctx, params)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if res.IsError {
		log.Fatalf("CallTool Failed with error: %v", err)
		return nil, err
	}
	var scholarHits []ScholarHit
	for _, c := range res.Content {
		tc, ok := c.(*mcp.TextContent)
		if !ok {
			continue
		}
		var hit ScholarHit
		if err := json.Unmarshal([]byte(tc.Text), &hit); err != nil {
			log.Printf("Failed to unmarshall tc")
			continue
		}
		// log.Printf("Appending: \n Title: %s\nAbstract: %s\nAuthors :%s\nURL: %s\n",
		// hit.Title, hit.Abstract, hit.Authors, hit.Url)
		scholarHits = append(scholarHits, hit)
	}
	return scholarHits, nil
}
