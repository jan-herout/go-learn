package cc

import (
	"io/fs"
	"fmt"
	"path/filepath"
	"strings"
)

// Sniffer prepares list of SQL files, recursively, in a given directory.
type Sniffer struct {
	dir     string
	outChan chan string
	errors []error
}

// NewSniffer is a sniffer constructor. Each sniffer is bound to a directory.
func NewSniffer(dir string) *Sniffer {
	return &Sniffer{dir: dir}
}

// Sniff sends list of files, which have one of the provided extensions, down the channel.
// The processing runs in the background, function returns immediately, and returns the channel.
func (s *Sniffer) Sniff(ext []string) <-chan string {
	s.outChan = make(chan string)
	go s.sniff(ext)
	return s.outChan
}

func (s *Sniffer) sniff(ext []string) {
	defer close(s.outChan)
	filepath.Walk(s.dir, func(path string, info fs.FileInfo, err error) error {
		// chyba na vstupu?
		if err != nil {
			s.errors = append(s.errors, err)
		}
		// je to adresář? pokud ano, skip
		if info.IsDir() {
			return nil
		}
		// je to jeden ze "chtěných" souborů?
		for _, e := range ext {
			extUp := strings.ToUpper(e)
			nameUp := strings.ToUpper(path)
			if strings.HasSuffix(nameUp, extUp) {
				s.outChan <- path // ano, pošli ho dál
				return nil
			}
		}
		return nil
	})
}

// Err returns an error, if some were encountered.
func (s *Sniffer) Err() error {
	if s.errors == nil {
		return nil		
	}
	
	if len(s.errors) == 1 {
		return s.errors[0]
	}
	
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("sniffer has %d errors:\n", len(s.errors)))
	for _,err := range s.errors {
		sb.WriteString(fmt.Sprintf(" * %s", err))
	}
	return fmt.Errorf("%s", sb.String())
}
