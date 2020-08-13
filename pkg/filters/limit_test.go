package filters

import (
	"testing"

	"github.com/SaltyCatFish/caperrors/pkg/file"
	"github.com/SaltyCatFish/caperrors/pkg/file/mocks"
)

type LimitResult struct {
	limit    uint
	actual   uint
	expected bool
}

var dummy file.File = file.NewFile("", mocks.NewDefaultFileInfoMock())

var limitResults = []LimitResult{
	{
		limit:    10,
		actual:   9,
		expected: true,
	},
	{
		limit:    5,
		actual:   9,
		expected: false,
	},
	{
		limit:    0,
		actual:   100,
		expected: true,
	},
}

func TestLimit(t *testing.T) {
	for _, test := range limitResults {
		results := []bool{}
		a := NewLimit(test.limit)
		for i := uint(1); i <= test.actual; i++ {
			results = append(results, a.OK(dummy))
		}
		if results[test.actual-1] != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, results[test.actual-1])
		}
	}
}
