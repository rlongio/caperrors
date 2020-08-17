// Package cap_errors is a command line utility for gathering NOAA product error information.
package main

import (
	"flag"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/SaltyCatFish/caperrors/internal"
	"github.com/SaltyCatFish/caperrors/pkg/filters"
	"github.com/SaltyCatFish/caperrors/pkg/product"
)

func main() {

	var filePaths []string

	searchPath := flag.String("path", "", "root path of error files.")
	logFilePath := flag.String("log_path", "", "path of the log file (typically CapHandler.log)")

	finclude := flag.String("include", "", "comma separated values that must be present in path")
	fexclude := flag.String("exclude", "", "comma separated values that cannot be present in path")
	fbefore := flag.String("before", "", "YYYY-MM-DD date that results must come before")
	fafter := flag.String("after", "", "YYYY-MM-DD date that results must come after")
	flimit := flag.Uint("limit", 0, "limit the number of results to return")
	fdiagnostic := flag.Bool("diagnostic", false, "Set to true for use as Debugger Diagnostic")

	flag.Parse()

	any := filters.NewAny(toSlice(*finclude))
	none := filters.NewNone(toSlice(*fexclude))
	before := filters.NewBefore(toDate(*fbefore))
	after := filters.NewAfter(toDate(*fafter))
	extensions := filters.NewAny([]string{
		".txt",
		".xml",
	})
	noextensions := filters.NewNone([]string{ //TODO Dev only, prevents from including gz files
		".gz",
	})
	limit := filters.NewLimit(*flimit)

	filter := filters.NewFilter()
	filter.Add(any)
	filter.Add(none)
	filter.Add(before)
	filter.Add(after)
	filter.Add(extensions)
	filter.Add(noextensions)
	filter.Add(limit)

	filePaths = append(filePaths, *searchPath)

	var wg sync.WaitGroup
	var fx = product.Product.Print
	if *fdiagnostic {
		fx = product.Product.PrintDiagnosticOutput
	}

	ch := make(chan product.File)
	log.Printf("gathering product files..")
	p := internal.ProductFilesFromDirectoriesRecursively(filePaths, filter)
	log.Printf("found %v product files", len(p))

	for i := 1; i <= 5; i++ { // worker goroutines
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(fx, *logFilePath, ch)
		}()
	}
	for _, file := range p {
		ch <- product.NewFile(file)
	}
	close(ch)
	wg.Wait()
	log.Printf("finished")
}

// worker creates a product from a file and logfile
func worker(fx product.Process, logFilePath string, cs <-chan product.File) {
	for i := range cs {
		fx(product.CreateProduct(i, logFilePath))
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
