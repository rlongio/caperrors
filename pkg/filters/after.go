package filters

import (
	"time"

	"github.com/SaltyCatFish/caperrors/pkg/file"
)

// After holds a time.Time
type After struct {
	value time.Time
}

// OK returns true if value occurred after After
func (a After) OK(f file.Filer) bool {
	if (a.value == time.Time{}) {
		return true
	}
	return f.ModTime().After(a.value)
}

// NewAfter returns a new instance of After
func NewAfter(value time.Time) After {
	return After{
		value: value,
	}
}
