package filters

import (
	"testing"

	"github.com/SaltyCatFish/caperrors/pkg/file"
	filemock "github.com/SaltyCatFish/caperrors/pkg/file/mocks"
	filtermock "github.com/SaltyCatFish/caperrors/pkg/filters/mocks"
)

type addResult struct {
	filter   Filterer
	qty      int
	expected bool
}

var addResults = []addResult{
	{qty: 2},
	{qty: 10},
	{qty: 100},
	{qty: 1000},
	{qty: 0},
}

func TestAdd(t *testing.T) {
	for _, test := range addResults {
		filter := NewFilter()
		mockFilterer := filtermock.NewMockFilter(true)
		for i := 1; i <= test.qty; i++ {
			filter.Add(mockFilterer)
		}

		if len(filter.filters) != test.qty {
			t.Fatalf("%v does not equal %v", len(filter.filters), test.qty)
		}
	}
}

type isOKResult struct {
	results  []bool
	expected bool
}

var isOKResults = []isOKResult{
	{
		results:  []bool{true, true, true},
		expected: true,
	},
	{
		results:  []bool{false, false, false},
		expected: false,
	},
	{
		results:  []bool{false, true, true},
		expected: false,
	},
}

func TestIsOK(t *testing.T) {
	for _, test := range isOKResults {
		filter := NewFilter()
		for _, result := range test.results {
			filter.Add(filtermock.NewMockFilter(result))
		}

		fileMock := file.NewFile("", filemock.NewFileInfoMock())
		if filter.OK(fileMock) != test.expected {
			t.Errorf("%v does not equal %v", filter.OK(fileMock), test.expected)
		}
	}
}
