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
)

// Filer contains methods to retrieve a product ID and error message from a log file
type Filer interface {
	ID(string) (string, error)
	ErrorMessage(string) (string, error)
}

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
	return "", fmt.Errorf("Unknown file extension: %v", f.base())
}

// ErrorMessage returns the error information found for the product in logFilePath
func (f File) ErrorMessage(logFilePath string) (message string, err error) {

	regexSearchString := fmt.Sprintf("Attempting to move.*/caphandler/error.*%v$", f.base())
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
			message = previousLine
			return
		}
		previousLine = scanner.Text()
	}
	err = scanner.Err()
	if err != nil {
		return
	}
	return
}

// Ext returns the file extension
func (f File) ext() string {
	return filepath.Ext(f.abspath())
}

// abspath returns the relative filepath and name
func (f File) abspath() (path string) {
	return filepath.Join(f.path, f.file.Name())
}

// base returns the name of the file without the path
func (f File) base() string {
	return f.file.Name()
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
		if strings.Contains(scanner.Text(), f.base()) && strings.Contains(scanner.Text(), "The ID of the text product in") {
			return regex.FindStringSubmatch(scanner.Text())[2], err
		}
	}
	err = scanner.Err()
	return "not found", err
}

// idFromXMLFile returns product ID parsed from the XML file
func (f File) idFromXMLFile(logFilePath string) (id string, err error) {
	id = "not found"
	file, err := os.Open(f.abspath())
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
	if strings.ToLower(f.ext()) == strings.ToLower(extension) {
		return true
	}
	return false
}
