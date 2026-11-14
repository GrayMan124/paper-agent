package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	semanticscholar "github.com/GrayMan124/paper-agent/internal/checker/semanticScholar"
	"github.com/joho/godotenv"
)

// const MCPPATH = "../../mcp_servers.json"

func main() {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load enviorment: %v", err)
	}
	ss := semanticscholar.SemanticScholarConfig{
		API_KEY: os.Getenv("SEMANTIC_API"),
		Client:  &http.Client{},
		Muc:     sync.Mutex{},
		LastReq: time.Now(),
	}
	if err := run(ctx, &ss, os.Args[1:]); err != nil {
		log.Fatalf("Failed run with error: %v", err)
	}
}

func run(ctx context.Context, semantic *semanticscholar.SemanticScholarConfig, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Expected subcommand (e.g. 'Scholar')")
	}
	log.Print("Paper-Bot is running...")
	subcommand := args[0]
	subArgs := args[1:]
	switch subcommand {
	case "scholar":
		return runScholar(ctx, subArgs)
	// case "api":
	// 	return semantic.SemanticScholarAPI("Attention is all")
	case "search_check":
		return runCheck(ctx, subArgs, semantic)
	default:
		return fmt.Errorf("Unknown command: %q", subcommand)
	}
}
