package file

import (
	"testing"
	"time"

	"github.com/SaltyCatFish/caperrors/pkg/file/mocks"
)

type FileResults struct {
	file File
}

var file File = NewFile("path", mocks.NewDefaultFileInfoMock())

var newFilesResults = []FileResults{
	{
		file: file,
	},
}

type SortFileResults struct {
	files Files
}

var sortFileResults = []SortFileResults{
	{
		files: Files{
			NewFile("second", mocks.NewModTimeMock(time.Date(2021, time.June, 1, 0, 0, 0, 0, time.Local))),
			NewFile("first", mocks.NewModTimeMock(time.Date(2021, time.January, 1, 0, 0, 0, 0, time.Local))),
			NewFile("third", mocks.NewModTimeMock(time.Date(2021, time.December, 1, 0, 0, 0, 0, time.Local))),
		},
	},
}

func TestNewFiles(t *testing.T) {
	for _, test := range newFilesResults {
		files := NewFiles()
		files.Add(test.file)
		if len(files) != 1 {
			t.Errorf("Expected %v, got %v", len(files), 1)
		}
	}
}

func TestFilesSortAsc(t *testing.T) {
	for _, test := range sortFileResults {
		test.files.SortAsc()
		if test.files[0].path != "first" {
			t.Errorf("Expected %v, got %v", test.files[0].path, "first")
		}
		if test.files[1].path != "second" {
			t.Errorf("Expected %v, got %v", test.files[1].path, "second")
		}
		if test.files[2].path != "third" {
			t.Errorf("Expected %v, got %v", test.files[2].path, "third")
		}
	}
}

func TestFilesSortDesc(t *testing.T) {
	for _, test := range sortFileResults {
		test.files.SortDesc()
		if test.files[0].path != "third" {
			t.Errorf("Expected %v, got %v", test.files[2].path, "third")
		}
		if test.files[1].path != "second" {
			t.Errorf("Expected %v, got %v", test.files[1].path, "second")
		}
		if test.files[2].path != "first" {
			t.Errorf("Expected %v, got %v", test.files[0].path, "first")
		}
	}
}
