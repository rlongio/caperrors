package product

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// File represents a file that contains product information
type File struct {
	path string
	file os.FileInfo
}

// NewFile returns a new instance of File
func NewFile(path string, file os.FileInfo) File {
	return File{
		path: path,
		file: file,
	}
}

// Ext returns the file extension
func (f File) Ext() string {
	return filepath.Ext(f.Path())
}

// Path returns the relative filepath and name
func (f File) Path() (path string) {
	return filepath.Join(f.path, f.file.Name())
}

// ModTime returns the last time the file was modified
func (f File) ModTime() (modTime time.Time) {
	return f.ModTime()
}

// Base returns the name of the file without the path
func (f File) Base() string {
	return f.file.Name()
}

// ID retrieves the product ID from either a .txt or .xml file
func (f File) ID(logFilePath string) (id string, err error) {
	if f.hasExtension(".txt") {
		id, err = f.idFromTextFile(logFilePath)
		return
	}
	if f.hasExtension(".xml") {
		id, err = f.idFromXMLFile(logFilePath)
		return
	}
	return "", fmt.Errorf("Unknown file extension: %v", f.Base())
}

// ErrorMessage returns the error information found for the product in logFilePath
func (f File) ErrorMessage(logFilePath string) (errorMessage string, err error) {

	regexSearchString := fmt.Sprintf("Attempting to move.*/caphandler/error.*%v$", f.Base())
	regex := regexp.MustCompile(regexSearchString)

	logFile, err := os.Open(logFilePath)
	if err != nil {
		return
	}
	defer logFile.Close()

	previousLine := ""
	scanner := bufio.NewScanner(logFile)
	for scanner.Scan() {
		if regex.FindString(scanner.Text()) != "" {
			errorMessage = previousLine
			return
		}
		previousLine = scanner.Text()
	}
	err = scanner.Err()
	return "not found", err
}

// idFromTextFile returns the matching id for product in logFilePath
func (f File) idFromTextFile(logFilePath string) (id string, err error) {
	regex := regexp.MustCompile(`(=\s)(.*$)`)

	logFile, err := os.Open(logFilePath)
	if err != nil {
		return
	}
	defer logFile.Close()

	scanner := bufio.NewScanner(logFile)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), f.Base()) && strings.Contains(scanner.Text(), "The ID of the text product in") {
			return regex.FindStringSubmatch(scanner.Text())[2], err
		}
	}
	err = scanner.Err()
	return "not found", err
}

// idFromXMLFile returns product ID parsed from the XML file
func (f File) idFromXMLFile(logFilePath string) (id string, err error) {
	id = "not found"
	file, err := os.Open(f.Path())
	if err != nil {
		return
	}
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	defer file.Close()

	var identifier struct {
		XMLName xml.Name `xml:"alerts"`
		ID      string   `xml:"alert>identifier"`
	}

	err = xml.Unmarshal([]byte(content), &identifier)
	if err != nil {
		return
	}
	id = identifier.ID
	return
}

// hasExtension returns true if the file has the passes extension in its path
// false otherwise
func (f File) hasExtension(extension string) bool {
	if strings.ToLower(f.Ext()) == strings.ToLower(extension) {
		return true
	}
	return false
}
