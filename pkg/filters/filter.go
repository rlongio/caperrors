// Package filters provides ways of filtering files.
package filters

import (
	"path/filepath"

	"github.com/SaltyCatFish/caperrors/pkg/file"
)

// Filter holds a slice of Filterers
type Filter struct {
	filters []Filterer
}

// Add adds a Filterer to Filter
func (f *Filter) Add(filter Filterer) {
	f.filters = append(f.filters, filter)
}

// OK returns true if all Filterers OK's return true
func (f Filter) OK(file file.Filer) bool {
	for _, f := range f.filters {
		if !f.OK(file) {
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
func absFilePath(file file.Filer) string {
	path, _ := filepath.Abs(file.Name())
	return filepath.Join(path, file.Name())
}
