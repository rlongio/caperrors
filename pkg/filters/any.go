package filters

import (
	"strings"

	"github.com/SaltyCatFish/caperrors/pkg/file"
)

// Any holds a slice of values
type Any struct {
	values []string
}

// OK returns true if the value is in Any
func (a Any) OK(f file.Filer) bool {
	if (len(a.values)) < 1 || a.values[0] == "" {
		return true
	}
	for _, needle := range a.values {
		if strings.Contains(f.Path(), needle) {
			return true
		}
	}
	return false
}

// NewAny returns a new instance of Any
func NewAny(values []string) Any {
	return Any{
		values: values,
	}
}
