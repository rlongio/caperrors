// Package filters provides ways of filtering files.
package filters

import (
	"os"
	"path/filepath"
)

// Filter holds a slice of Filterers
type Filter struct {
	filters []Filterer
}

// Add adds a Filterer to Filter
func (f *Filter) Add(filter Filterer) {
	f.filters = append(f.filters, filter)
}

// IsOK returns true if all Filterers IsOK's return true
func (f Filter) IsOK(file os.FileInfo) bool {
	for _, f := range f.filters {
		if !f.IsOK(file) {
			return false
		}
	}
	return true
}

// NewFilter returns a new instance of Filter
func NewFilter() Filter {
	return Filter{}
}

// absFilePath returns the absolute file path of a file
func absFilePath(file os.FileInfo) string {
	path, _ := filepath.Abs(file.Name())
	return path
}
