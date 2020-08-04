package filters

import (
	"os"
	"time"
)

// Before holds a time.Time
type Before struct {
	value time.Time
}

// IsOK returns true if value occured before Before
func (b Before) IsOK(file os.FileInfo) bool {
	if (b.value == time.Time{}) {
		return true
	}
	return file.ModTime().Before(b.value)
}

// NewBefore returns a new instance of Before
func NewBefore(value time.Time) Before {
	return Before{
		value: value,
	}
}
