package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/GrayMan124/paper-agent/internal/contract"
	"github.com/GrayMan124/paper-agent/internal/sources/scholar"
	"log"
	"os"
)

// const MCPPATH = "../../mcp_servers.json"
const MCPPATH = "./mcp_servers.json"

func main() {
	ctx := context.Background()

	if err := run(ctx, os.Args[1:]); err != nil {
		log.Fatalf("Failed run with error: %v", err)
	}
}

func run(ctx context.Context, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("Expected subcommand (e.g. 'Scholar')")
	}
	log.Print("Paper-Bot is running...")
	subcommand := args[0]
	subArgs := args[1:]
	switch subcommand {
	case "scholar":
		return runScholar(ctx, subArgs)
	default:
		return fmt.Errorf("Unknown command: %q", subcommand)
	}
	return nil
}

func runScholar(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("scholar", flag.ContinueOnError)

	query := fs.String("q", "", "Search Query")
	limit := fs.Int("limit", 5, "Max results")
	fixture := fs.String("fixture", "", "Path to local fixture JSON")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if *fixture != "" {
		//NOTE: Temporary solution, data is skipped for now it doesn't support fixtures
		_, err := os.ReadFile(*fixture)
		if err != nil {
			return err
		}
	}

	if *query == "" {
		return fmt.Errorf("Missing required flag -q (Query)")
	}

	log.Printf("Searching for %s...", *query)
	scholarConf, err := scholar.BuildConfig(MCPPATH)
	if err != nil {
		return err
	}
	schol := scholar.Scholar{Config: scholarConf}
	results, err := schol.Search(ctx, *query, *limit)
	if err != nil {
		return err
	}
	for i, res := range results {

		log.Printf("RESULT %v\nTitle:%s\nAuthors:%s\nAbstract:%s\nURL:%s\n", i, res.Title, res.Authors, res.Abstract, res.Url)
		cand, err := contract.TranslateScholarHit(i, res)
		if err != nil {
			log.Fatalf("Failed to create candidate with error: %v", err)
		}
		log.Printf("Candidate %v\nTitle:%s\nAuthors:%s\nYear:%v\nURL:%s\n\n", i, cand.Title, cand.AuthorsRaw, *cand.Year, *cand.Link)

	}
	return nil

}
