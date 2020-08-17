package filters

import "github.com/SaltyCatFish/caperrors/pkg/file"

// Limit holds a max and current uint32
type Limit struct {
	max     uint
	current uint
}

// OK returns true if value occured before Before
func (l *Limit) OK(file file.Filer) bool {
	if l.max == 0 {
		return true
	}
	l.current++
	if l.current <= l.max {
		return true
	}
	return false
}

// NewLimit returns a reference to an instance of Limit
func NewLimit(value uint) *Limit {
	return &Limit{
		max:     value,
		current: 0,
	}
}
