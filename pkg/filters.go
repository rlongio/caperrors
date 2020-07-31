package pkg

import (
	"os"
	"path/filepath"
	"strings"
	"time"
)

// After holds a time.Time
type After struct {
	Value time.Time
}

// IsOK returns true if value occured after After
func (a After) IsOK(file os.FileInfo) bool {
	if (a.Value == time.Time{}) {
		return true
	}
	return file.ModTime().After(a.Value)
}

// Any holds a slice of values
type Any struct {
	Values []string
}

// IsOK returns true if the value is in Any
func (a Any) IsOK(file os.FileInfo) bool {
	if (len(a.Values)) <= 1 {
		return true
	}
	for _, needle := range a.Values {
		if strings.Contains(absFilePath(file), needle) {
			return true
		}
	}
	return false
}

// Before holds a time.Time
type Before struct {
	Value time.Time
}

// IsOK returns true if value occured before Before
func (b Before) IsOK(file os.FileInfo) bool {
	if (b.Value == time.Time{}) {
		return true
	}
	return file.ModTime().Before(b.Value)
}

// None holds a slice of values
type None struct {
	Values []string
}

// IsOK returns true if the value is not in None
func (n None) IsOK(file os.FileInfo) bool {
	if (len(n.Values)) <= 1 {
		return true
	}
	for _, needle := range n.Values {
		if strings.Contains(absFilePath(file), needle) {
			return false
		}
	}
	return true
}

func absFilePath(file os.FileInfo) string {
	path, _ := filepath.Abs(file.Name())
	return path
}
