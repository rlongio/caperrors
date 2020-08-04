package filters

import (
	"os"
	"time"
)

// After holds a time.Time
type After struct {
	value time.Time
}

// IsOK returns true if value occured after After
func (a After) IsOK(file os.FileInfo) bool {
	if (a.value == time.Time{}) {
		return true
	}
	return file.ModTime().After(a.value)
}

// NewAfter returns a new instance of After
func NewAfter(value time.Time) After {
	return After{
		value: value,
	}
}
