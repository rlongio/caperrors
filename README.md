# Cap Errors

![Badge](https://github.com/SaltyCatFish/caperrors/workflows/Go/badge.svg)

Cap Errors is a library and command line utility for collecting and parsing Cap Handler errors.

## Command Line Utilities

### caperrors

`caperrors` is a utility for find Cap Handler errors. It works by taking the product file (xml or text) and matching it with an entry in the log file to determine its ID and Error message.

The utility uses go routines to significanly increase the speed of searching the often large log file.

Help can be found by running `caperrors -h`. This will show a list of the current flags available.

The required flags are:

`- path=<path to error files>`

`- log_path=<path to log file>`

## Packages

### filters

The `filters` package is used to filter the results of a `filepath.Walk` by incrementally building different filters based on time.Time values and strings.

For example, if you wanted to filter results to only be after 2020-01-01 that include the string 'new' in the name:

```go
any := filters.NewAny([]string{"new",})

after := filters.NewAfter(time.Date(
    2020, time.January, 1, 0, 0, 0, 0, time.UTC
))

filter := filters.NewFilter()
filter.Add(any)
filter.Add(after)
```

Each `filter` implements the `filterer` interface, which requires one method.

```go
type Filterer interface {
    OK(file os.FileInfo) bool
}
```

An example of filtering search results, continue from above:

```go
var f = []os.FileInfo{}

// put search results in f...

results := []os.FileInfo{}
for x := range f {
    if filter.OK(x) {
        results = append(results, x)
    }
}
```

### product

The `product` package is used for product development. It provides a base for current and future development.

For more information, look at the source code documentation.
