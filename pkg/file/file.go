package file

import (
	"os"
	"path/filepath"
	"strings"
)

// File represents a file that contains product information
type File struct {
	path string
	os.FileInfo
}

// NewFile returns a new instance of File
func NewFile(path string, file os.FileInfo) File {
	return File{
		path:     path,
		FileInfo: file,
	}
}

// Ext returns the file extension
func (f File) Ext() string {
	return filepath.Ext(f.Path())
}

// Path returns the relative filepath and name
func (f File) Path() string {
	return f.path
}

// Base returns the name of the file without the path
func (f File) Base() string {
	return f.Name()
}

// HasExtension returns true if the file has the passes extension in its path
// false otherwise
func (f File) HasExtension(extension string) bool {
	if strings.ToLower(f.Ext()) == strings.ToLower(extension) {
		return true
	}
	return false
}
