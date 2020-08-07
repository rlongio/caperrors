package filters

import (
	"os"
	"testing"
)

type NoneResult struct {
	values   []string
	file     os.FileInfo
	expected bool
}

var noneResults = []NoneResult{
	{
		values: []string{
			"hello",
		},
		file:     NewModNameMock("hello"),
		expected: false,
	},
	{
		values: []string{
			"hello",
		},
		file:     NewModNameMock("goodbye"),
		expected: true,
	},
	{
		values:   []string{},
		file:     NewModNameMock("emptyNone"),
		expected: true,
	},
	{
		values: []string{
			"hello",
			"goodbye",
		},
		file:     NewModNameMock("hello"),
		expected: false,
	},
	{
		values: []string{
			"hello",
			"goodbye",
		},
		file:     NewModNameMock("afternoon"),
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
