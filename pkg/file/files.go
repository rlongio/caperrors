package file

import "sort"

// Files holds a slice of File
type Files []File

// NewFiles returns a slice of File
func NewFiles() Files {
	return Files{}
}

//Add adds a file
func (f *Files) Add(file File) {
	*f = append(*f, file)
}

// SortAsc returns the files sorted by ModTime descending
func (f Files) SortAsc() {
	sort.Slice(f, func(i, j int) bool { return f[i].ModTime().Before(f[j].ModTime()) })
	return
}

// SortDesc returns the files sorted by ModTime descending
func (f Files) SortDesc() {
	sort.Slice(f, func(i, j int) bool { return f[i].ModTime().After(f[j].ModTime()) })
	return
}
