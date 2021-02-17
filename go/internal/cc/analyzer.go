package cc

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

// Analyzer opens all files from input channel, and identifies DDL scripts.
type Analyzer struct {
	in     <-chan string
	out    chan Metadata
	errors []error
}

// NewAnalyzer is an analyzer constructor. Channel c is an input channel, containing paths.
func NewAnalyzer(c <-chan string) *Analyzer {
	a := &Analyzer{}
	a.in = c
	a.out = make(chan Metadata)
	return a
}

// Analyzer opens files from input channel for reading, identifies DDL scripts, and finds
// name of created table. It then sends Metadata to the output channel.
// The processing runs in the background, function returns immediately, and returns the channel.
func (s *Analyzer) Analyze() <-chan Metadata {
	go s.analyze()
	return s.out
}

// goroutine which reads the data, and decides
func (s *Analyzer) analyze() {
	defer close(s.out)
	re := regexp.MustCompile(`(?i)create\s+(set\s+|multiset\s+)?table\s+(?P<db>[a-z0-9_]+[.])?(?P<tab>[a-z0-9_]+)`)
	reCmt := regexp.MustCompile(`--conditional_create--`)
	for f := range s.in {
		// open and read the file
		h, err := os.Open(f)
		if s.checkErr(err) {
			continue
		}
		data, err := io.ReadAll(h)
		h.Close() // no defer, jsme v cyklu!
		if s.checkErr(err) {
			continue
		}
		
		// do not touch already changed files
		dataString := string(data)
		if reCmt.MatchString(dataString) {
			continue
		}
		
		// test if it is a DDL script
		match := re.FindStringSubmatch(string(dataString))
		names := re.SubexpNames()
		result := make(map[string]string)
		for i, name := range names {
			if i == len(match) {
				break
			} // no luck...
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}
		if name, ok := result["tab"]; ok {
			meta := Metadata{
				Path:      f,
				TableName: name,
				Content:   data,
			}
			s.out <- meta
		}
	}

}

func (s *Analyzer) checkErr(err error) bool {
	if err == nil {
		return false
	}
	s.errors = append(s.errors, err)
	return true
}

// Err returns an error, if some were encountered.
func (s *Analyzer) Err() error {
	if s.errors == nil {
		return nil
	}

	if len(s.errors) == 1 {
		return s.errors[0]
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("sniffer has %d errors:\n", len(s.errors)))
	for _, err := range s.errors {
		sb.WriteString(fmt.Sprintf(" * %s", err))
	}
	return fmt.Errorf("%s", sb.String())
}
