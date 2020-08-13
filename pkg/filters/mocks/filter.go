package mocks

import (
	"github.com/SaltyCatFish/caperrors/pkg/file"
)

// MockFilter is a mock of Filter
type MockFilter struct {
	ok bool
}

// NewMockFilter returns a new MockFilter
func NewMockFilter(ok bool) MockFilter {
	return MockFilter{
		ok: ok,
	}
}

// OK returns true is MockFilter.ok is true
func (m MockFilter) OK(file file.Filer) bool {
	return m.ok
}
