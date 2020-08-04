package main

import (
	"flag"
	"strings"
	"sync"
	"time"

	"noaa.gov/rlong/cap_errors/internal"
	filters "noaa.gov/rlong/cap_errors/pkg/filters"
	product "noaa.gov/rlong/cap_errors/pkg/product"
)

func main() {

	var filePaths []string

	searchPath := flag.String("path", "", "root path of error files.")
	logFilePath := flag.String("log_path", "", "path of the log file (typically CapHandler.log)")

	finclude := flag.String("include", "", "comma separated values that must be present in path")
	fexclude := flag.String("exclude", "", "comma separated values that cannot be present in path")
	fbefore := flag.String("before", "", "YYYY-MM-DD date that results must come before")
	fafter := flag.String("after", "", "YYYY-MM-DD date that results must come after")

	any := filters.NewAny(toSlice(*finclude))
	none := filters.NewNone(toSlice(*fexclude))
	before := filters.NewBefore(toDate(*fbefore))
	after := filters.NewAfter(toDate(*fafter))
	extensions := filters.NewAny([]string{
		".txt",
		".xml",
	})

	filter := filters.NewFilter()
	filter.Add(any)
	filter.Add(none)
	filter.Add(before)
	filter.Add(after)
	filter.Add(extensions)

	flag.Parse()

	filePaths = append(filePaths, *searchPath)

	var wg sync.WaitGroup
	var fx = product.Product.Print

	ch := make(chan product.File)
	p := internal.ProductFilesFromDirectoriesRecursively(filePaths, filter)

	for i := 1; i <= 5; i++ { // worker goroutines
		go worker(fx, *logFilePath, ch, &wg)
	}
	for _, productFile := range p {
		wg.Add(1)
		ch <- productFile
	}
	wg.Wait()
	close(ch)
}

// worker creates a product from a file and logfile
func worker(fx product.Process, logFilePath string, cs chan product.File, wg *sync.WaitGroup) {
	for i := range cs {
		fx(product.CreateProduct(i, logFilePath))
		wg.Done()
	}
}

// toSlice converts string to a slice using "," as delimiter
func toSlice(value string) []string {
	return strings.Split(value, ", ")
}

// toDate takes a date string in the form of YYYY-MM-DD and
// returns a time.Time instnace
func toDate(date string) (t time.Time) {
	if date == "" {
		return time.Time{}
	}
	layout := "2006-01-02"
	t, err := time.ParseInLocation(layout, date, time.Local)
	if err != nil {
		panic(err)
	}
	return t
}
