package file

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/SaltyCatFish/caperrors/pkg/file/mocks"
)

type NewFileResult struct {
	path     string
	fileInfo os.FileInfo
}

var newFileResults = []NewFileResult{
	{
		path:     "/home/derp",
		fileInfo: mocks.NewDefaultFileInfoMock(),
	},
}

func TestNewFile(t *testing.T) {
	for _, test := range newFileResults {
		a := NewFile(test.path, test.fileInfo)

		if a.Path() != filepath.Join(test.path, a.Name()) {
			t.Errorf("Expected %v, got %v", filepath.Join(test.path, a.Name()), a.Path())
		}

		if a.Base() != test.fileInfo.Name() {
			t.Errorf("Expected %v, got %v", a.Base(), test.fileInfo.Name())
		}

		if a.Ext() != ".txt" {
			t.Errorf("Expected %v, got %v", a.Ext(), ".txt")
		}

		if a.HasExtension(".txt") != true {
			t.Errorf("Expected %v, got %v", a.HasExtension(".txt"), true)
		}

		if a.HasExtension(".TXT") != true {
			t.Errorf("Expected %v, got %v", a.HasExtension(".txt"), true)
		}

		if a.HasExtension(".doc") != false {
			t.Errorf("Expected %v, got %v", a.HasExtension(".txt"), false)
		}

	}
}
