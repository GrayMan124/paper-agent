package scholar

import (
	"bytes"
	"context"
	"encoding/json"
	"os"
	"testing"
)

func decodeJson(t *testing.T, path string, v any) {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read %s: %v", path, err)
	}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(v); err != nil {
		t.Fatalf("decode %s: %v", path, err)
	}

}

func TestConfigFromFile(t *testing.T) {
	var scholarConfig ScholarConfig
	decodeJson(t, "../../../mcp_servers.json", &scholarConfig)

}

func TestScholarSearch(t *testing.T) {
	var scholarConfig ScholarConfig
	decodeJson(t, "../../../mcp_servers.json", &scholarConfig)
	scholar := Scholar{Config: scholarConfig}
	err := scholar.Search(context.Background(), "LLM robustness geometry", 5)
	if err != nil {
		t.Fatalf("Search Failed with error: %v", err)
	}
}
