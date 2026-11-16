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
	"github.com/GrayMan124/paper-agent/internal/sources/scholar"
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
	// Preparing google scholar
	gsc, err := scholar.BuildConfig("./mcp_servers.json")
	schol := scholar.Scholar{Config: gsc}
	if err := run(ctx, &ss, schol, os.Args[1:]); err != nil {
		log.Fatalf("Failed create google scholar with error: %v", err)
	}

}

func run(ctx context.Context, semantic *semanticscholar.SemanticScholarConfig, schol scholar.Scholar, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Expected subcommand (e.g. 'Scholar')")
	}
	log.Print("Paper-Bot is running...")
	subcommand := args[0]
	subArgs := args[1:]
	switch subcommand {
	case "scholar":
		return runScholar(ctx, subArgs, schol)
	// case "api":
	// 	return semantic.SemanticScholarAPI("Attention is all")
	case "search_check":
		return runCheck(ctx, subArgs, semantic, schol)
	default:
		return fmt.Errorf("Unknown command: %q", subcommand)
	}
}
