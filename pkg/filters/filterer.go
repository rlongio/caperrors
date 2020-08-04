package filters

import "os"

// Filterer defines IsOK that represent if the underlying filter passes or fails
type Filterer interface {
	IsOK(file os.FileInfo) bool
}
