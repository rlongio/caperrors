package filters

import (
	"strings"

	"github.com/SaltyCatFish/caperrors/pkg/file"
)

// None holds a slice of values
type None struct {
	values []string
}

// OK returns true if the value is not in None
func (n None) OK(file file.Filer) bool {
	if (len(n.values)) < 1 || n.values[0] == "" {
		return true
	}
	for _, needle := range n.values {
		if strings.Contains(absFilePath(file), needle) {
			return false
		}
	}
	return true
}

// NewNone returns a new instance of None
func NewNone(values []string) None {
	return None{
		values: values,
	}
}
