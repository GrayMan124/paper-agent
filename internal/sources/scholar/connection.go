package scholar

import (
	"context"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"log"
	"os/exec"
)

type Scholar struct {
	Config ScholarConfig
}

func (s Scholar) Search(ctx context.Context, query string, n int) error {
	client := mcp.NewClient(&mcp.Implementation{
		Name:    "paperbot",
		Version: "0.0.1",
	}, nil)
	cmd := exec.Command(s.Config.McpServers.GoogleScholar.Command, s.Config.McpServers.GoogleScholar.Args...)
	session, err := client.Connect(ctx, &mcp.CommandTransport{Command: cmd}, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer session.Close()
	params := &mcp.CallToolParams{
		Name:      "search_google_scholar_key_words",
		Arguments: map[string]any{"query": query, "num_results": n},
	}
	res, err := session.CallTool(ctx, params)
	if err != nil {
		log.Fatal(err)
		return err
	}
	if res.IsError {
		log.Fatalf("CallTool Failed with error: %v", err)
		return err
	}
	for _, c := range res.Content {
		log.Print(c.(*mcp.TextContent).Text)
	}
	return nil

}
