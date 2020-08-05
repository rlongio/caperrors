package product

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
	"testing"
)

func decompress(path string) (*bufio.Scanner, error) {
	fmt.Println("Decompressing")
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	gz, err := gzip.NewReader(file)
	if err != nil {
		return nil, err
	}

	return bufio.NewScanner(gz), nil
}

func TestThing(t *testing.T) {
	scanner, err := decompress("../../testdata/files/dummy_test.xml.gz")
	if err != nil {
		t.Fatal("f'd up")
	}
	t.Log("Results")
	fmt.Println("Results")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
