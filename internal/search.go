package internal

import (
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
		err := filepath.Walk(path, func(path string, file os.FileInfo, err error) error {
			if err != nil {
				log.Fatalln(err)
				return filepath.SkipDir
			}
			if !file.IsDir() && filter.IsOK(file) {
				absdir, err := absdir(path, file)
				if err != nil {
					log.Printf("ERROR: %v", err)
					return err
				}
				f := product.NewFile(absdir, file)
				log.Printf("Appending %v", filepath.Join(absdir, file.Name()))
				files = append(files, f)
			}
			return nil
		})
		if err != nil {
			log.Printf(err.Error())
		}
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
