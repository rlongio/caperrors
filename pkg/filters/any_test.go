package filters

import (
	"testing"

	"github.com/SaltyCatFish/caperrors/pkg/file"
	"github.com/SaltyCatFish/caperrors/pkg/file/mocks"
)

type AnyResult struct {
	values   []string
	file     file.File
	expected bool
}

var anyResults = []AnyResult{
	{
		values: []string{
			"hello",
		},
		file:     file.NewFile("hello", mocks.NewModNameMock("world")),
		expected: true,
	},
	{
		values: []string{
			"hello",
		},
		file:     file.NewFile("goodbye", mocks.NewModNameMock("world")),
		expected: false,
	},
	{
		values:   []string{},
		file:     file.NewFile("empty", mocks.NewModNameMock("any")),
		expected: true,
	},
	{
		values: []string{
			"hello",
			"goodbye",
		},
		file:     file.NewFile("hello", mocks.NewModNameMock("howdy")),
		expected: true,
	},
	{
		values: []string{
			"hello",
			"goodbye",
		},
		file:     file.NewFile("howdy", mocks.NewModNameMock("partner")),
		expected: false,
	},
}

func TestAny(t *testing.T) {
	for _, test := range anyResults {
		a := NewAny(test.values)
		result := a.OK(test.file)
		if result != test.expected {
			t.Errorf("%v: Expected %v, got %v", test.file.Name(), test.expected, result)
		}
	}
}
