package internal

import (
	"io/ioutil"
	"os"
)

// CreateTempDirectory creates a tempoarary directory
func CreateTempDirectory(name string) (dir string, err error) {
	dir, err = ioutil.TempDir("", "")
	return
}

// CreateTempFile creates a temporary file
func CreateTempFile(path string, contentString string) (tmpfile *os.File, err error) {
	content := []byte(contentString)
	tmpfile, err = ioutil.TempFile(path, "")
	if err != nil {
		return
	}
	if _, err := tmpfile.Write(content); err != nil {
		return nil, err
	}
	return
}
