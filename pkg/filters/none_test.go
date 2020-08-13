package filters

import (
	"testing"

	"github.com/SaltyCatFish/caperrors/pkg/file"
	"github.com/SaltyCatFish/caperrors/pkg/file/mocks"
)

type NoneResult struct {
	values   []string
	file     file.File
	expected bool
}

var noneResults = []NoneResult{
	{
		values: []string{
			"hello",
		},
		file:     file.NewFile("path", mocks.NewModNameMock("hello")),
		expected: false,
	},
	{
		values: []string{
			"hello",
		},
		file:     file.NewFile("path", mocks.NewModNameMock("goodbye")),
		expected: true,
	},
	{
		values:   []string{},
		file:     file.NewFile("path", mocks.NewModNameMock("emptyNone")),
		expected: true,
	},
	{
		values: []string{
			"hello",
			"goodbye",
		},
		file:     file.NewFile("path", mocks.NewModNameMock("hello")),
		expected: false,
	},
	{
		values: []string{
			"hello",
			"goodbye",
		},
		file:     file.NewFile("path", mocks.NewModNameMock("afternoon")),
		expected: true,
	},
}

func TestNone(t *testing.T) {
	for _, test := range noneResults {
		a := NewNone(test.values)
		result := a.OK(test.file)
		if result != test.expected {
			t.Errorf("%v: Expected %v, got %v", test.file.Name(), test.expected, result)
		}
	}
}
