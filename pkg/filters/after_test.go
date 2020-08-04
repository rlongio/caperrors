package filters

import (
	"os"
	"testing"
	"time"
)

type AfterResult struct {
	file     os.FileInfo
	time     time.Time
	expected bool
}

var afterResults = []AfterResult{
	{
		file:     NewModTimeMock(time.Date(2021, time.January, 1, 0, 0, 0, 0, time.Local)),
		time:     time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local),
		expected: true,
	},
	{
		file:     NewModTimeMock(time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local)),
		time:     time.Date(2021, time.January, 1, 0, 0, 0, 0, time.Local),
		expected: false,
	},
	{
		file:     NewModTimeMock(time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local)),
		time:     time.Time{},
		expected: true,
	},
}

func TestAfter(t *testing.T) {
	for _, test := range afterResults {
		a := NewAfter(test.time)
		result := a.IsOK(test.file)
		if result != test.expected {
			t.Errorf("%v: Expected %v, got %v", test.file.ModTime(), test.expected, result)
		}
	}
}
