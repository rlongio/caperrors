package filters

import "github.com/SaltyCatFish/caperrors/pkg/file"

// Filterer defines IsOK that represent if the underlying filter passes or fails
type Filterer interface {
	OK(file file.Filer) bool
}
