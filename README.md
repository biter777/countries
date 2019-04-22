countries
=======

Package countries - ISO 3166 (3166-1) countries codes and countries names (on eng and rus), very light and quick, NO maps[], NO slices[], NO any external packages, NO init() func, Databases compatible

[![GoDoc](http://godoc.org/github.com/biter777/countries?status.svg)](http://godoc.org/github.com/biter777/countries)


installation
------------

    go get github.com/biter777/countries

usage
-----

```go
import "github.com/biter777/countries"

func main() {
	countryJapan := countries.Japan
    fmt.Printf("Country name: %v\n", countryJapan)
    fmt.Printf("Country name in russian: %v\n", countryJapan.StringRus())
    fmt.Printf("Country ISO2: %v\n", countryJapan.ISO2())
    fmt.Printf("Country ISO3: %v\n", countryJapan.ISO3())
    ...
}
```

options
-------

You can take a counties names in russian language, use StringRus()

```go
import "github.com/biter777/countries"
```

For more complex options, consult the [documentation](http://godoc.org/github.com/biter777/countries).

contributing
------------

(c) Biter

We welcome pull requests, bug fixes and issue reports.
Before proposing a change, please discuss it first by raising an issue.
