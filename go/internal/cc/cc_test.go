package cc

import (
	"testing"
)

func Test_cc(t *testing.T) {
	prepFixtures("testData")
	sniffer := NewSniffer("testData")
	analyzer := NewAnalyzer(sniffer.Sniff([]string{"sql"}))

	files := make([]Metadata, 0)
	for f := range analyzer.Analyze() {
		files = append(files, f)
	}

	if len(files) != 1 {
		t.Errorf("expected to see 1 ddl script, got %d", len(files))
	}

	if err := sniffer.Err(); err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if err := analyzer.Err(); err != nil {
		t.Errorf("unexpected error: %s", err)
	}
}
