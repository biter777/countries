countries
=======

Countries - ISO 3166 (ISO3166-1, ISO3166, Digit code, Alpha-2 and Alpha-3), ISO 4217 countries codes and names (on eng and rus), currency designators, calling phone codes, countries capitals and regions (UN M.49 code), countries domains (ccTLD), IOC and FIFA letters codes, Very FAST, NO maps[], NO slices[], NO external files and data, NO interface{}, NO specific dependencies, Databases compatible, Emoji countries flags and currencies support, full support ISO-3166-1, ISO-4217, Unicode CLDR and ccTLD standarts.

Full support ISO-3166-1, ISO-4217, Unicode CLDR and ccTLD standarts.

[![GoDoc](http://godoc.org/github.com/biter777/countries?status.svg)](http://godoc.org/github.com/biter777/countries)


installation
------------

    go get github.com/biter777/countries

usage
-----

```go
import "fmt"
import "github.com/biter777/countries"

func main() {
	countryJapan := countries.Japan
	fmt.Printf("Country name in english: %v\n", countryJapan)
	fmt.Printf("Country name in russian: %v\n", countryJapan.StringRus())
	fmt.Printf("Country digit code: %d\n", countryJapan)
	fmt.Printf("Country Alpha-2 code: %v\n", countryJapan.Alpha2())
	fmt.Printf("Country Alpha-3 code: %v\n", countryJapan.Alpha3())
	fmt.Printf("Country IOC code: %v\n", countryJapan.IOC())
	fmt.Printf("Country FIFA code: %v\n", countryJapan.FIFA())
	fmt.Printf("Country Capital: %v\n", countryJapan.Capital())
	fmt.Printf("Country call code: %v\n", countryJapan.CallCodes())
	fmt.Printf("Country domain: %v\n", countryJapan.Domain())
	fmt.Printf("Country region name: %v\n", countryJapan.Region())
	fmt.Printf("Country region code: %d\n", countryJapan.Region())
	fmt.Printf("Country emoji: %v\n\n", countryJapan.Emoji())

	currencyJapan := countryJapan.Currency()
	fmt.Printf("Country Currency name in english: %v\n", currencyJapan)
	fmt.Printf("Country Currency digit code: %d\n", currencyJapan)
	fmt.Printf("Country Currency Alpha code: %v\n", currencyJapan.Alpha())
	fmt.Printf("Country Currency emoji: %v\n", currencyJapan.Emoji())
	fmt.Printf("Country of Currency %v: %v\n\n", currencyJapan, currencyJapan.Countries())

	// OR you can alternative use:
	japanInfo := countries.Japan.Info()
	fmt.Printf("Country name in english: %v\n", japanInfo.Name)
	fmt.Printf("Country digit code: %d\n", japanInfo.Code)
	fmt.Printf("Country Alpha-2 code: %v\n", japanInfo.Alpha2)
	fmt.Printf("Country Alpha-3 code: %v\n", japanInfo.Alpha3)
	fmt.Printf("Country IOC code: %v\n", japanInfo.IOC)
	fmt.Printf("Country FIFA code: %v\n", japanInfo.FIFA)
	fmt.Printf("Country Capital: %v\n", japanInfo.Capital)
	fmt.Printf("Country call code: %v\n", japanInfo.CallCodes)
	fmt.Printf("Country region name: %v\n", japanInfo.Region)
	fmt.Printf("Country region code: %d\n", japanInfo.Region)
	fmt.Printf("Country emoji: %v\n", japanInfo.Emoji)

	// Detect by name
	country := countries.ByName("angola")
	fmt.Printf("Country digit code: %d\n", country)
	fmt.Printf("Country Alpha-2 code: %v\n", country.Alpha2())
	fmt.Printf("Country Alpha-3 code: %v\n", country.Alpha3())

	// Database usage
	type User struct {
		gorm.Model
		Name    string
		Country countries.CountryCode
	}
	user := &User{Name: "Helen", Country: countries.Slovenia}
	db, err := gorm.Open("postgres", 500, "host=127.0.0.2 port=5432 user=usr password=1234567 dbname=db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Create(user).Error
	if err != nil {
		panic(err)
	}
}
```

options
-------

You can take a counties names in russian language, use StringRus(). For Emoji use Emoji(). Enjoy!

```go
import "github.com/biter777/countries"
```

For more complex options, consult the [documentation](http://godoc.org/github.com/biter777/countries).

contributing
------------

(c) Biter

Welcome pull requests, bug fixes and issue reports.
Before proposing a change, please discuss it first by raising an issue.

Contributors list:
@biter777 (https://github.com/biter777), 
@gavincarr (https://github.com/gavincarr),  
@benja-M-1 (https://github.com/benja-M-1) 
