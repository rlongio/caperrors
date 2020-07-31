package internal

import (
	"os"
	"path/filepath"

	"noaa.gov/rlong/cap_errors/pkg"
)

// GetFilteredProductFilesFromDirectoriesRecursively returns all files from
// a directory recursively
func GetFilteredProductFilesFromDirectoriesRecursively(paths []string, filters []pkg.Filterer) (productFiles []pkg.File) {
	for _, path := range paths {
		filepath.Walk(path, func(path string, file os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if file.IsDir() {
				return nil
			}
			for _, f := range filters {
				if !f.IsOK(file) {
					return nil
				}
			}
			var productFile pkg.File
			productFile.Value = file
			productFiles = append(productFiles, productFile)
			return nil
		})
	}
	return
}
