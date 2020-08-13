package mocks

import (
	"os"
	"time"
)

// FileInfoMock returns a mock of os.FileInfo
type FileInfoMock struct {
	name    string
	size    int64
	mode    os.FileMode
	modtime time.Time
	isDir   bool
	sys     interface{}
}

// NewDefaultFileInfoMock returns a FileInfoMock with default values
// for testing
func NewDefaultFileInfoMock() FileInfoMock {
	return FileInfoMock{
		name:    "file.txt",
		size:    100,
		mode:    755,
		modtime: time.Date(2020, time.January, 1, 0, 0, 0, 0, time.Local),
		isDir:   false,
		sys:     nil,
	}
}

// Name returns name
func (f FileInfoMock) Name() string {
	return f.name
}

// Size returns size
func (f FileInfoMock) Size() int64 {
	return f.size
}

// Mode returns size (unint32)
func (f FileInfoMock) Mode() os.FileMode {
	return f.mode
}

// ModTime returns modification time as time.Time
func (f FileInfoMock) ModTime() time.Time {
	return f.modtime
}

// IsDir returns true if mock is a directory
func (f FileInfoMock) IsDir() bool {
	return f.isDir
}

// Sys returns an empty interface
func (f FileInfoMock) Sys() interface{} {
	return f.sys
}

// NewFileInfoMock returns a new instance of FileInfoMock
func NewFileInfoMock() FileInfoMock {
	return FileInfoMock{}
}

// NewModTimeMock creates a new FileInfoMock with a modTime
func NewModTimeMock(time time.Time) FileInfoMock {
	return FileInfoMock{
		modtime: time,
	}
}

// NewModNameMock creates a new FileInfoMock with a name
func NewModNameMock(name string) FileInfoMock {
	return FileInfoMock{
		name: name,
	}
}
