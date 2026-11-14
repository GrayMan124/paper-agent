package semanticscholar

type SemanticResult struct {
	Total  int `json:"total"`
	Offset int `json:"offset"`
	Next   int `json:"next"`
	Data   []struct {
		PaperID string `json:"paperId"`
		URL     string `json:"url"`
		Title   string `json:"title"`
		Year    int    `json:"year"`
		Authors []struct {
			AuthorID string `json:"authorId"`
			Name     string `json:"name"`
		} `json:"authors"`
	} `json:"data"`
}
