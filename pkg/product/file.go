package product

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	file "github.com/SaltyCatFish/caperrors/pkg/file"
)

// Filer contains methods to retrieve a product ID and error message from a log file
type Filer interface {
	ID(string) (string, error)
	ErrorMessage(string) (string, error)
}

// File contains methods for retrieving product information
// from a file
type File struct {
	file.File
}

// NewFile returns a File struct
func NewFile(f file.File) File {
	return File{f}
}

// ID retrieves the product ID from either a .txt or .xml file
func (f File) ID(logFilePath string) (id string, err error) {
	if f.HasExtension(".txt") {
		id, err = f.idFromTextFile(logFilePath)
		return
	}
	if f.File.HasExtension(".xml") {
		id, err = f.idFromXMLFile(logFilePath)
		return
	}
	return "", fmt.Errorf("Unknown file extension: %v", f.File.Base())
}

// ErrorMessage returns the error information found for the product in logFilePath
func (f File) ErrorMessage(logFilePath string) (message string, err error) {

	regexSearchString := fmt.Sprintf("Attempting to move.*/caphandler/error.*%v$", f.File.Base())
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
		if strings.Contains(scanner.Text(), f.File.Base()) && strings.Contains(scanner.Text(), "The ID of the text product in") {
			return regex.FindStringSubmatch(scanner.Text())[2], err
		}
	}
	err = scanner.Err()
	return "not found", err
}

// idFromXMLFile returns product ID parsed from the XML file
func (f File) idFromXMLFile(logFilePath string) (id string, err error) {
	id = "not found"
	file, err := os.Open(f.File.Abspath())
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
