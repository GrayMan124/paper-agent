package scholar

import (
	"bytes"
	"encoding/json"
	"os"
)

type ScholarConfig struct {
	McpServers struct {
		GoogleScholar struct {
			Command string   `json:"command"`
			Args    []string `json:"args"`
			Env     struct {
			} `json:"env"`
		} `json:"google-scholar"`
	} `json:"mcpServers"`
}

type ScholarHit struct {
	Title, Authors, Abstract, Url string
}

func configFromJSON(path string, v any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(v); err != nil {
		return err
	}
	return nil
}

func BuildConfig(path string) (ScholarConfig, error) {
	var scholarConfig ScholarConfig
	err := configFromJSON(path, &scholarConfig)
	if err != nil {
		return scholarConfig, err
	}
	return scholarConfig, err

}
