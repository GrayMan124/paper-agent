package contract

type Candidate struct {
	ResultId   string  `json:"result_id"`
	Title      string  `json:"title"`
	AuthorsRaw string  `json:"authors_raw"`
	Year       *int    `json:"year"`
	Link       *string `json:"link"`
	Doi        *string `json:"doi"`
	ArxivId    *string `json:"arxiv_id"`
}

type Result struct {
	ResultID  string `json:"result_id"`
	Status    string `json:"status"`
	Reason    string `json:"reason"`
	Canonical *struct {
		Title    string   `json:"title"`
		Authors  []string `json:"authors"`
		Year     *int     `json:"year"`
		Doi      *string  `json:"doi"`
		ArxivID  *string  `json:"arxiv_id"`
		Venue    *string  `json:"venue"`
		Abstract *string  `json:"abstract"`
	} `json:"canonical"`
}
