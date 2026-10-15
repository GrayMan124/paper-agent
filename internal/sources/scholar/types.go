package scholar

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
