countries
=======

Countries - ISO 3166 (ISO3166-1, ISO3166, Digit, Alpha-2 and Alpha-3) countries codes and names (on eng and rus), ISO 4217 currency designators, ITU-T E.164 IDD calling phone codes, countries capitals, UN M.49 regions codes, ccTLD countries domains, IOC/NOC and FIFA letters codes, VERY FAST, NO maps[], NO slices[], NO external links/files/data, NO interface{}, NO specific dependencies, Databases compatible, Emoji countries flags and currencies support, full support ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and ccTLD standarts.

Full support ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and ccTLD standarts.

[![GoDoc](http://godoc.org/github.com/biter777/countries?status.svg)](http://godoc.org/github.com/biter777/countries)
[![License](https://img.shields.io/badge/License-BSD%202--Clause-brightgreen.svg)](https://opensource.org/licenses/BSD-2-Clause)
[![Build Status](https://travis-ci.org/biter777/countries.svg?branch=master)](https://travis-ci.org/biter777/countries)
[![Codeship Status](https://codeship.com/projects/00db4400-1803-0138-1132-7ab932dd1523/status?branch=master)](https://app.codeship.com/projects/381056)

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
	fmt.Printf("Country name in english: %v\n", countryJapan)                   // Japan
	fmt.Printf("Country name in russian: %v\n", countryJapan.StringRus())       // Япония
	fmt.Printf("Country ISO-3166 digit code: %d\n", countryJapan)               // 392
	fmt.Printf("Country ISO-3166 Alpha-2 code: %v\n", countryJapan.Alpha2())    // JP
	fmt.Printf("Country ISO-3166 Alpha-3 code: %v\n", countryJapan.Alpha3())    // JPN
	fmt.Printf("Country IOC/NOC code: %v\n", countryJapan.IOC())                // JPN
	fmt.Printf("Country FIFA code: %v\n", countryJapan.FIFA())                  // JPN
	fmt.Printf("Country Capital: %v\n", countryJapan.Capital())                 // Tokyo
	fmt.Printf("Country ITU-T E.164 call code: %v\n", countryJapan.CallCodes()) // +81
	fmt.Printf("Country ccTLD domain: %v\n", countryJapan.Domain())             // .jp
	fmt.Printf("Country UN M.49 region name: %v\n", countryJapan.Region())      // Asia
	fmt.Printf("Country UN M.49 region code: %d\n", countryJapan.Region())      // 142
	fmt.Printf("Country emoji/flag: %v\n\n", countryJapan.Emoji())              // 🇯🇵

	currencyJapan := countryJapan.Currency()
	fmt.Printf("Country ISO-4217 Currency name in english: %v\n", currencyJapan)           // Yen
	fmt.Printf("Country ISO-4217 Currency digit code: %d\n", currencyJapan)                // 392
	fmt.Printf("Country ISO-4217 Currency Alpha code: %v\n", currencyJapan.Alpha())        // JPY
	fmt.Printf("Country Currency emoji: %v\n", currencyJapan.Emoji())                      // 💴
	fmt.Printf("Country of Currency %v: %v\n\n", currencyJapan, currencyJapan.Countries()) // Japan

	// OR you can alternative use:
	japanInfo := countries.Japan.Info()
	fmt.Printf("Country name in english: %v\n", japanInfo.Name)                          // Japan
	fmt.Printf("Country ISO-3166 digit code: %d\n", japanInfo.Code)                      // 392
	fmt.Printf("Country ISO-3166 Alpha-2 code: %v\n", japanInfo.Alpha2)                  // JP
	fmt.Printf("Country ISO-3166 Alpha-3 code: %v\n", japanInfo.Alpha3)                  // JPN
	fmt.Printf("Country IOC/NOC code: %v\n", japanInfo.IOC)                              // JPN
	fmt.Printf("Country FIFA code: %v\n", japanInfo.FIFA)                                // JPN
	fmt.Printf("Country Capital: %v\n", japanInfo.Capital)                               // Tokyo
	fmt.Printf("Country ITU-T E.164 call code: %v\n", japanInfo.CallCodes)               // +81
	fmt.Printf("Country ccTLD domain: %v\n", japanInfo.Domain)                           // .jp
	fmt.Printf("Country UN M.49 region name: %v\n", japanInfo.Region)                    // Asia
	fmt.Printf("Country UN M.49 region code: %d\n", japanInfo.Region)                    // 142
	fmt.Printf("Country emoji/flag: %v\n", japanInfo.Emoji)                              // 🇯🇵
	fmt.Printf("Country ISO-4217 Currency name in english: %v\n", japanInfo.Currency)    // Yen
	fmt.Printf("Country ISO-4217 Currency digit code: %d\n", japanInfo.Currency)         // 392
	fmt.Printf("Country ISO-4217 Currency Alpha code: %v\n", japanInfo.Currency.Alpha()) // JPY

	// Detection usage
	// Detect by name
	country := countries.ByName("angola")
	fmt.Printf("Country name in english: %v\n", country)                // Angola
	fmt.Printf("Country ISO-3166 digit code: %d\n", country)            // 24
	fmt.Printf("Country ISO-3166 Alpha-2 code: %v\n", country.Alpha2()) // AO
	fmt.Printf("Country ISO-3166 Alpha-3 code: %v\n", country.Alpha3()) // AGO
	// Detect by code/numeric
	country = countries.ByNumeric(24)
	fmt.Printf("Country name in english: %v\n", country)                // Angola
	fmt.Printf("Country ISO-3166 digit code: %d\n", country)            // 24
	fmt.Printf("Country ISO-3166 Alpha-2 code: %v\n", country.Alpha2()) // AO
	fmt.Printf("Country ISO-3166 Alpha-3 code: %v\n", country.Alpha3()) // AGO

	// Comparing usage
	// Compare by code/numeric
	if countries.ByName("angola") == countries.AGO {
		fmt.Println("Yes! It's Angola!") // Yes! It's Angola!
	}
	// Compare by name
	if strings.EqualFold("angola", countries.AGO.String()) {
		fmt.Println("Yes! It's Angola!") // Yes! It's Angola!
	}

	// Database usage
	type User struct {
		gorm.Model
		Name     string
		Country  countries.CountryCode
		Currency countries.CurrencyCode
	}
	user := &User{Name: "Helen", Country: countries.Slovenia, Currency: countries.CurrencyEUR}
	db, err := gorm.Open("postgres", 500, "host=127.0.0.2 port=5432 user=usr password=1234567 dbname=db")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.Create(user)
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
@biter777 (https://github.com/biter777) 
@gavincarr (https://github.com/gavincarr) 
@benja-M-1 (https://github.com/benja-M-1) 
