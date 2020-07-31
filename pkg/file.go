package pkg

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// File represents a file that contains product information
type File struct {
	Value os.FileInfo
}

// Path returns the relative filepath and name
func (f File) Path() (path string, err error) {
	path, err = filepath.Abs(f.Base())
	return
}

// ModTime returns the last time the file was modified
func (f File) ModTime() (modTime time.Time) {
	return f.ModTime()
}

// Base returns the name of the file without the path
func (f File) Base() string {
	return f.Value.Name()
}

// GetProductID retrieves the product ID from either a .txt or .xml file
func (f File) GetProductID(logFilePath string) (productID string, err error) {
	if hasExtension(f.Base(), ".txt") {
		productID, err = f.getProductIDFromTextFile(logFilePath)
		return
	}
	if hasExtension(f.Base(), ".xml") {
		productID, err = f.getProductIDFromXMLFile(logFilePath)
		return
	}
	log.Printf("Warning: unrecognized file extension [%v]", f.Base())
	return "", err
}

// GetProductErrorInformation returns the error information found for the product in the log file
func (f File) GetProductErrorInformation(logFilePath string) (errorInformation string, err error) {

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
			errorInformation = previousLine
			return
		}
		previousLine = scanner.Text()
	}
	err = scanner.Err()
	return "not found", err
}

func (f File) getProductIDFromTextFile(logFilePath string) (productID string, err error) {
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

func (f File) getProductIDFromXMLFile(logFilePath string) (productID string, err error) {
	logFile, err := os.Open(logFilePath)
	if err != nil {
		log.Fatal(err)
	}
	content, err := ioutil.ReadAll(logFile)
	defer logFile.Close()

	var identifier struct {
		ID string `xml:"identifier"`
	}

	err = xml.Unmarshal([]byte(content), &identifier)
	return identifier.ID, err
}

// HasExtension returns true if the file has the passes extension in its path
// false otherwise
func hasExtension(filePath string, extension string) bool {
	if strings.ToLower(filepath.Ext(filePath)) == strings.ToLower(extension) {
		return true
	}
	return false
}
