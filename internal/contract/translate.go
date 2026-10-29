package contract

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/GrayMan124/paper-agent/internal/sources/scholar"
)

func clearTitle(title string) string {
	cleaned := strings.TrimSpace(title)

	bracketRegex := regexp.MustCompile(`\[.*?\]`)
	cleaned = bracketRegex.ReplaceAllString(cleaned, "")

	chineseRegex := regexp.MustCompile(`\p{Han}+`)
	cleaned = chineseRegex.ReplaceAllString(cleaned, "")

	cleaned = strings.Join(strings.Fields(cleaned), " ")
	return cleaned
}

func translateScholarHit(id int, scholarHit scholar.ScholarHit) (Candidate, error) {
	yearRe := regexp.MustCompile(`\b(19|20)\d{2}\b`)
	// tagRe := regexp.MustCompile(`^(\[[^\]]+\])+\s*`)

	parts := strings.Split(scholarHit.Authors, "- ")
	log.Printf("Parts len: %v", len(parts))
	for partId, part := range parts {
		log.Printf("Part %v: %s", partId, part)
	}
	year := yearRe.FindString(parts[1])
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		return Candidate{}, err
	}
	cand := Candidate{
		ResultId:   strconv.Itoa(id),
		Title:      clearTitle(scholarHit.Title),
		AuthorsRaw: parts[0],
		Year:       &yearInt,
		Link:       &scholarHit.Url,
	}
	return cand, nil
}

func TranslateScholarHits(scholarResults []scholar.ScholarHit) ([]Candidate, error) {
	var translatedResults []Candidate
	for id, sr := range scholarResults {
		translation, err := translateScholarHit(id, sr)
		if err != nil {
			log.Fatalf("Failed to convert ScholarHit into Candidate due to error: %v", err)
		}
		translatedResults = append(translatedResults, translation)
	}
	return translatedResults, nil
}
