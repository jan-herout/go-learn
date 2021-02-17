package main

import (
	"fmt"
	"strings"
	"sync"
)

// ErrGroup is an error group which is safe for concurrent writes.
type ErrGroup struct {
	mu sync.Mutex
	es []error
}

// Append adds an error to the group.
func (e *ErrGroup) Append(err error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	if err == nil {
		return
	}
	e.es = append(e.es, err)
}

// Errors returns a new slice, which contains slice of errors encountered so far.
// If no errors were encountered, returns nil.
func (e *ErrGroup) Errors() []error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.es == nil {
		return nil
	}
	s := make([]error, len(e.es))
	copy(s, e.es)
	return s
}

// Error implementes the error interface
func (e *ErrGroup) Error() string {
	e.mu.Lock()
	defer e.mu.Unlock()

	if e.es == nil {
		return ""
	}
	
	if len(e.es) == 1 {
		return e.es[0].Error()
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Total #%d errors:\n", len(e.es)))
	for _, err := range e.es {
		sb.WriteString(fmt.Sprintf(" * %s\n", err))
	}
	return sb.String()
}

func main() {
	// START,OMIT
	var eg ErrGroup
	eg.Append(fmt.Errorf("první chyba"))	
	eg.Append(fmt.Errorf("druhá chyba"))		
	fmt.Println("------- BAD ----------------")
	fmt.Println(eg)
	fmt.Println("------- GOOD ---------------")
	fmt.Println(&eg)
	// END,OMIT
}
