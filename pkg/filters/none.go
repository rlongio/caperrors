package filters

import (
	"os"
	"strings"
)

// None holds a slice of values
type None struct {
	values []string
}

// IsOK returns true if the value is not in None
func (n None) IsOK(file os.FileInfo) bool {
	if (len(n.values)) <= 1 {
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
