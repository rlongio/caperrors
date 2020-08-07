package filters

import (
	"os"
	"testing"
	"time"
)

type BeforeResult struct {
	file     os.FileInfo
	time     time.Time
	expected bool
}

var beforeResults = []BeforeResult{
	{
		file:     NewModTimeMock(time.Date(2021, time.January, 1, 0, 0, 0, 0, time.Local)),
		time:     time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local),
		expected: false,
	},
	{
		file:     NewModTimeMock(time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local)),
		time:     time.Date(2021, time.January, 1, 0, 0, 0, 0, time.Local),
		expected: true,
	},
	{
		file:     NewModTimeMock(time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local)),
		time:     time.Time{},
		expected: true,
	},
}

func TestBefore(t *testing.T) {
	for _, test := range beforeResults {
		a := NewBefore(test.time)
		result := a.OK(test.file)
		if result != test.expected {
			t.Errorf("%v: Expected %v, got %v", test.file.ModTime(), test.expected, result)
		}
	}
}
