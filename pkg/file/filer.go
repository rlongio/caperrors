package file

import (
	"time"
)

// Filer contains file operations
type Filer interface {
	Ext() string
	Path() string
	Base() string
	ModTime() time.Time
	HasExtension(string) bool
	Name() string
}
