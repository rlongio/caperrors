package internal

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	filters "github.com/SaltyCatFish/caperrors/pkg/filters"
	product "github.com/SaltyCatFish/caperrors/pkg/product"
)

// ProductFilesFromDirectoriesRecursively returns all files from
// a directory recursively and reports the filtered results.
func ProductFilesFromDirectoriesRecursively(paths []string, filter filters.Filter) (files []product.File) {
	for _, path := range paths {
		fmt.Printf("filter: %v", filter)
		filepath.Walk(path, func(path string, file os.FileInfo, err error) error {
			if err != nil {
				log.Fatalln(err)
				return nil
			}
			if file.IsDir() {
				return nil
			}
			if !filter.IsOK(file) {
				return nil
			}
			absdir, err := absdir(path, file)
			if err != nil {
				log.Println(err)
				return nil
			}
			f := product.NewFile(absdir, file)
			files = append(files, f)
			return nil
		})
	}
	return
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
