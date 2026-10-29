package contract

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

func SaveResultsToJSON(path string, candidates []Candidate) error {
	fullPath := filepath.Join(path, "candidates.json")
	file, err := os.Create(fullPath)
	if err != nil {
		log.Fatalf("Failed to create file with error: %v", err)
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "	")
	err = encoder.Encode(candidates)
	if err != nil {
		log.Fatalf("Encode candidates into a file with error: %v", err)
		return err
	}
	return err

}
