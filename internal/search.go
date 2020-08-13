package internal

import (
	"log"
	"os"
	"path/filepath"

	"github.com/SaltyCatFish/caperrors/pkg/file"
	"github.com/SaltyCatFish/caperrors/pkg/filters"
	"github.com/SaltyCatFish/caperrors/pkg/product"
)

// ProductFilesFromDirectoriesRecursively returns all files from
// a directory recursively and reports the filtered results.
func ProductFilesFromDirectoriesRecursively(paths []string, filter filters.Filter) (results file.Files) {
	for _, path := range paths {
		err := filepath.Walk(path, func(p string, f os.FileInfo, err error) error {
			if err != nil {
				log.Printf(err.Error())
				return filepath.SkipDir
			}
			results.Add(file.NewFile(p, f))
			return nil
		})
		if err != nil {
			log.Printf(err.Error())
		}
	}
	results = sortAndFilterResults(results, filter)
	return
}

func sortAndFilterResults(files file.Files, filter filters.Filter) (results file.Files) {
	files.SortDesc()
	for _, file := range files {
		if ok(file, filter) {
			results = append(results, file)
		}
	}
	return
}

func toProductFile(f file.File) (p product.File) {
	p = product.NewFile(f)
	return
}

// isOK returns true if requirements are met and filter is true
func ok(file file.File, filter filters.Filter) bool {
	return !file.IsDir() && filter.OK(file)
}

// absdir returns absolute directory combining a relative path and file
//
// os.walk path only returns a relative path.  While there is context,
// grab the absolute path.
func absdir(path string, file os.FileInfo) (result string, err error) {
	result = ""
	absfilePath, err := filepath.Abs(file.Name())
	if err != nil {
		return
	}
	result = filepath.Dir(filepath.Join(filepath.Dir(absfilePath), path))
	return
}
