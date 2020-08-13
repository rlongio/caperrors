package filters

import (
	"testing"
	"time"

	"github.com/SaltyCatFish/caperrors/pkg/file"
	"github.com/SaltyCatFish/caperrors/pkg/file/mocks"
)

type AfterResult struct {
	file     file.File
	time     time.Time
	expected bool
}

var afterResults = []AfterResult{
	{
		file:     file.NewFile("", mocks.NewModTimeMock(time.Date(2021, time.January, 1, 0, 0, 0, 0, time.Local))),
		time:     time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local),
		expected: true,
	},
	{
		file:     file.NewFile("", mocks.NewModTimeMock(time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local))),
		time:     time.Date(2021, time.January, 1, 0, 0, 0, 0, time.Local),
		expected: false,
	},
	{
		file:     file.NewFile("", mocks.NewModTimeMock(time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local))),
		time:     time.Time{},
		expected: true,
	},
}

func TestAfter(t *testing.T) {
	for _, test := range afterResults {
		a := NewAfter(test.time)
		result := a.OK(test.file)
		if result != test.expected {
			t.Errorf("%v: Expected %v, got %v", test.file.ModTime(), test.expected, result)
		}
	}
}
