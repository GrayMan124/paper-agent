package contract

import (
	"bytes"
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

func TestCandidateDecode(t *testing.T) {
	var cand Candidate
	decodeJson(t, "../../schema/candidate_example.json", &cand)

}

func TestResultCorrDecode(t *testing.T) {
	var res Result
	decodeJson(t, "../../schema/result_example_corr.json", &res)

}
func TestResultFailDecode(t *testing.T) {
	var res Result
	decodeJson(t, "../../schema/result_failed_nf.json", &res)
}
