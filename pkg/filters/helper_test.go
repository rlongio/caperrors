package filters

import (
	"os"
	"time"
)

type FileInfoMock struct {
	name    string
	size    int64
	mode    os.FileMode
	modtime time.Time
	isDir   bool
	sys     interface{}
}

func (f FileInfoMock) Name() string {
	return f.name
}

func (f FileInfoMock) Size() int64 {
	return f.size
}

func (f FileInfoMock) Mode() os.FileMode {
	return f.mode
}

func (f FileInfoMock) ModTime() time.Time {
	return f.modtime
}

func (f FileInfoMock) IsDir() bool {
	return f.isDir
}

func (f FileInfoMock) Sys() interface{} {
	return f.sys
}

func NewModTimeMock(time time.Time) FileInfoMock {
	return FileInfoMock{
		modtime: time,
	}
}

func NewModNameMock(name string) FileInfoMock {
	return FileInfoMock{
		name: name,
	}
}

type MockFilter struct {
	ok bool
}

func NewMockFilter(ok bool) MockFilter {
	return MockFilter{
		ok: ok,
	}
}

func (m MockFilter) OK(file os.FileInfo) bool {
	return m.ok
}
