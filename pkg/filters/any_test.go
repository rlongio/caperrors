package filters

import (
	"os"
	"testing"
)

type AnyResult struct {
	values   []string
	file     os.FileInfo
	expected bool
}

var anyResults = []AnyResult{
	{
		values: []string{
			"hello",
		},
		file:     NewModNameMock("hello"),
		expected: true,
	},
	{
		values: []string{
			"hello",
		},
		file:     NewModNameMock("goodbye"),
		expected: false,
	},
	{
		values:   []string{},
		file:     NewModNameMock("emptyAny"),
		expected: true,
	},
	{
		values: []string{
			"hello",
			"goodbye",
		},
		file:     NewModNameMock("hello"),
		expected: true,
	},
	{
		values: []string{
			"hello",
			"goodbye",
		},
		file:     NewModNameMock("afternoon"),
		expected: false,
	},
}

func TestAny(t *testing.T) {
	for _, test := range anyResults {
		a := NewAny(test.values)
		result := a.IsOK(test.file)
		if result != test.expected {
			t.Errorf("%v: Expected %v, got %v", test.file.Name(), test.expected, result)
		}
	}
}
