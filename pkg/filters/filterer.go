package filters

import "os"

// Filterer defines IsOK that represent if the underlying filter passes or fails
type Filterer interface {
	OK(file os.FileInfo) bool
}
