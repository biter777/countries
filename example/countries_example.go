package main

import (
	"fmt"
	"strings"

	"github.com/royshahaf/countries"
	// "github.com/jinzhu/gorm"
)

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
	angola := countries.ByName("angola")
	fmt.Printf("Country name in english: %v\n", angola)                // Angola
	fmt.Printf("Country ISO-3166 digit code: %d\n", angola)            // 24
	fmt.Printf("Country ISO-3166 Alpha-2 code: %v\n", angola.Alpha2()) // AO
	fmt.Printf("Country ISO-3166 Alpha-3 code: %v\n", angola.Alpha3()) // AGO
	// Detect by code/numeric
	country := countries.ByNumeric(40)
	fmt.Printf("Country name in english: %v\n", country)                // Austria
	fmt.Printf("Country ISO-3166 digit code: %d\n", country)            // 40
	fmt.Printf("Country ISO-3166 Alpha-2 code: %v\n", country.Alpha2()) // AT
	fmt.Printf("Country ISO-3166 Alpha-3 code: %v\n", country.Alpha3()) // AUT

	// Comparing usage
	// Compare by code/numeric
	if countries.ByName("angola") == countries.AGO {
		fmt.Println("Yes! It's Angola!") // Yes! It's Angola!
	}
	// Compare by name
	if strings.EqualFold("angola", countries.AGO.String()) {
		fmt.Println("Yes! It's Angola!") // Yes! It's Angola!
	}

	if countries.ByName("Saint Martin") == countries.MAF {
		fmt.Println("Yes! It's Saint Martin!") // Yes! It's Saint Martin!
	}

	if countries.ByName("Iran, Islamic Republic Of") == countries.IRN {
		fmt.Println("Yes! It's Iran!") // Yes! It's Iran!
	}
	// Database usage
	/*
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
	*/
}
