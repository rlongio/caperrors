package main

import (
	"flag"
	"strings"
	"sync"
	"time"

	"noaa.gov/rlong/cap_errors/internal"
	"noaa.gov/rlong/cap_errors/pkg"
)

func main() {

	var filePaths []string

	searchPath := flag.String("path", "", "root path of error files.")
	logFilePath := flag.String("log_path", "", "path of the log file (typically CapHandler.log)")

	finclude := flag.String("include", "", "comma separated values that must be present in path")
	fexclude := flag.String("exclude", "", "comma separated values that cannot be present in path")
	fbefore := flag.String("before", "", "YYYY-MM-DD date that results must come before")
	fafter := flag.String("after", "", "YYYY-MM-DD date that results must come after")

	any := pkg.Any{
		Values: toSlice(*finclude),
	}

	none := pkg.None{
		Values: toSlice(*fexclude),
	}

	b, err := toDate(*fbefore)
	if err != nil {
		panic(err)
	}
	before := pkg.Before{
		Value: b,
	}

	a, err := toDate(*fafter)
	if err != nil {
		panic(err)
	}
	after := pkg.After{
		Value: a,
	}

	filters := []pkg.Filterer{
		any,
		none,
		before,
		after,
	}

	flag.Parse()

	filePaths = append(filePaths, *searchPath)

	var wg sync.WaitGroup
	var fx = pkg.Product.Print

	ch := make(chan pkg.File)
	p := internal.GetFilteredProductFilesFromDirectoriesRecursively(filePaths, filters)

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

func worker(fx pkg.Process, logFilePath string, cs chan pkg.File, wg *sync.WaitGroup) {
	for i := range cs {
		fx(pkg.CreateProduct(i, logFilePath))
		wg.Done()
	}
}

func toSlice(value string) []string {
	return strings.Split(value, ", ")
}

// DateStringToDate takes a date string in the form of YYYY-MM-DD and
// returns a time.Time instnace
func toDate(date string) (t time.Time, err error) {
	if date == "" {
		return time.Time{}, err
	}
	layout := "2006-01-02"
	t, err = time.ParseInLocation(layout, date, time.Local)
	return
}
