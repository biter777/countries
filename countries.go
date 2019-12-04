/*
Package countries - ISO 3166 (ISO3166-1, ISO3166, Digit code, Alpha-2 and Alpha-3), ISO 4217 countries codes and names (on eng and rus), currency designators, calling phone codes, countries capitals and regions (UN M.49 code), countries domains (ccTLD), Very FAST, NO maps[], NO slices[], NO init() func, NO external files and data, NO interface{}, NO specific dependencies, Databases compatible, Emoji countries flags and currencies support, full support ISO-3166-1, ISO-4217, Unicode CLDR and ccTLD standarts.

Usage

	func main() {
	countryJapan := countries.Japan
	fmt.Printf("Country name in english: %v\n", countryJapan)
	fmt.Printf("Country name in russian: %v\n", countryJapan.StringRus())
	fmt.Printf("Country digit code: %d\n", countryJapan)
	fmt.Printf("Country Alpha-2 code: %v\n", countryJapan.Alpha2())
	fmt.Printf("Country Alpha-3 code: %v\n", countryJapan.Alpha3())
	fmt.Printf("Country Capital: %v\n", countryJapan.Capital())
	fmt.Printf("Country call code: %v\n", countryJapan.CallCode())
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
	fmt.Printf("Country Capital: %v\n", japanInfo.Capital)
	fmt.Printf("Country call code: %v\n", japanInfo.CallCode)
	fmt.Printf("Country region name: %v\n", japanInfo.Region)
	fmt.Printf("Country region code: %d\n", japanInfo.Region)
	fmt.Printf("Country emoji: %v\n", japanInfo.Emoji)
	}

Detection usage

	// Detect by name
	country := countries.ByName("angola")
	fmt.Printf("Country digit code: %d\n", country)
	fmt.Printf("Country Alpha-2 code: %v\n", country.Alpha2())
	fmt.Printf("Country Alpha-3 code: %v\n", country.Alpha3())

Database usage

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

Contributing

 Welcome pull requests, bug fixes and issue reports.
 Before proposing a change, please discuss it first by raising an issue. */
package countries

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

// CountryCode - country code (250 countries). Three codes present, for example Russia == RU == RUS == 643.
type CountryCode int64 // int64 for database/sql/driver.Valuer compatibility

// Country - all info about country
type Country struct {
	Name     string       `json:"name"`
	Alpha2   string       `json:"cca2"`
	Alpha3   string       `json:"cca3"`
	Emoji    string       `json:"emoji"`
	Code     CountryCode  `json:"code"`
	Currency CurrencyCode `json:"currency"`
	Capital  CapitalCode  `json:"capital"`
	CallCode CallCode     `json:"callingCode"`
	Domain   DomainCode   `json:"domain"`
	Region   RegionCode   `json:"region"`
}

// Typer - typer interface, provide a name of type
type Typer interface {
	Type() string
}

// TypeCountryCode for Typer interface
const TypeCountryCode = "countries.CountryCode"

// TypeCountry for Typer interface
const TypeCountry = "countries.Country"

// UnknownMsg - unknown return message
const UnknownMsg = "Unknown"

// Digit ISO 3166-1. Three codes present, for example Russia == RU == RUS == 643.
const (
	Unknown                                CountryCode = 0
	International                          CountryCode = 999
	Albania                                CountryCode = 8
	Algeria                                CountryCode = 12
	AmericanSamoa                          CountryCode = 16
	Andorra                                CountryCode = 20
	Angola                                 CountryCode = 24
	Anguilla                               CountryCode = 660
	Antarctica                             CountryCode = 10
	AntiguaAndBarbuda                      CountryCode = 28
	Argentina                              CountryCode = 32
	Armenia                                CountryCode = 51
	Aruba                                  CountryCode = 533
	Australia                              CountryCode = 36
	Austria                                CountryCode = 40
	Azerbaijan                             CountryCode = 31
	Bahamas                                CountryCode = 44
	Bahrain                                CountryCode = 48
	Bangladesh                             CountryCode = 50
	Barbados                               CountryCode = 52
	Belarus                                CountryCode = 112
	Belgium                                CountryCode = 56
	Belize                                 CountryCode = 84
	Benin                                  CountryCode = 204
	Bermuda                                CountryCode = 60
	Bhutan                                 CountryCode = 64
	Bolivia                                CountryCode = 68
	BosniaAndHerzegovina                   CountryCode = 70
	Botswana                               CountryCode = 72
	Bouvet                                 CountryCode = 74
	Brazil                                 CountryCode = 76
	BritishIndianOceanTerritory            CountryCode = 86
	Brunei                                 CountryCode = 96
	Bulgaria                               CountryCode = 100
	BurkinaFaso                            CountryCode = 854
	Burundi                                CountryCode = 108
	Cambodia                               CountryCode = 116
	Cameroon                               CountryCode = 120
	Canada                                 CountryCode = 124
	CapeVerde                              CountryCode = 132
	CaboVerde                              CountryCode = 132
	CaymanIslands                          CountryCode = 136
	CentralAfricanRepublic                 CountryCode = 140
	Chad                                   CountryCode = 148
	ChannelIslands                         CountryCode = 830
	Chile                                  CountryCode = 152
	China                                  CountryCode = 156
	ChristmasIsland                        CountryCode = 162
	CocosIslands                           CountryCode = 166
	Colombia                               CountryCode = 170
	Comoros                                CountryCode = 174
	Congo                                  CountryCode = 178
	CongoDemocracticRepublic               CountryCode = 180
	CookIslands                            CountryCode = 184
	CostaRica                              CountryCode = 188
	CoteDIvoire                            CountryCode = 384
	IvoryCoast                             CountryCode = 384
	Croatia                                CountryCode = 191
	Cuba                                   CountryCode = 192
	Cyprus                                 CountryCode = 196
	CzechRepublic                          CountryCode = 203
	Denmark                                CountryCode = 208
	Djibouti                               CountryCode = 262
	Dominica                               CountryCode = 212
	DominicanRepublic                      CountryCode = 214
	Ecuador                                CountryCode = 218
	Egypt                                  CountryCode = 818
	ElSalvador                             CountryCode = 222
	EquatorialGuinea                       CountryCode = 226
	Eritrea                                CountryCode = 232
	Estonia                                CountryCode = 233
	Ethiopia                               CountryCode = 231
	FaroeIslands                           CountryCode = 234
	FalklandIslands                        CountryCode = 238
	Fiji                                   CountryCode = 242
	Finland                                CountryCode = 246
	France                                 CountryCode = 250
	FrenchGuiana                           CountryCode = 254
	FrenchPolynesia                        CountryCode = 258
	FrenchSouthernTerritories              CountryCode = 260
	Gabon                                  CountryCode = 266
	Gambia                                 CountryCode = 270
	Georgia                                CountryCode = 268
	Germany                                CountryCode = 276
	Ghana                                  CountryCode = 288
	Gibraltar                              CountryCode = 292
	Greece                                 CountryCode = 300
	Greenland                              CountryCode = 304
	Grenada                                CountryCode = 308
	Guadeloupe                             CountryCode = 312
	Guam                                   CountryCode = 316
	Guatemala                              CountryCode = 320
	Guinea                                 CountryCode = 324
	GuineaBissau                           CountryCode = 624
	Guyana                                 CountryCode = 328
	Haiti                                  CountryCode = 332
	HeardIslandandMcDonaldIslands          CountryCode = 334
	Honduras                               CountryCode = 340
	HongKong                               CountryCode = 344
	Hungary                                CountryCode = 348
	Iceland                                CountryCode = 352
	India                                  CountryCode = 356
	Indonesia                              CountryCode = 360
	Iran                                   CountryCode = 364
	Iraq                                   CountryCode = 368
	Ireland                                CountryCode = 372
	IsleOfMan                              CountryCode = 833
	Israel                                 CountryCode = 376
	Italy                                  CountryCode = 380
	Jamaica                                CountryCode = 388
	Japan                                  CountryCode = 392
	Jordan                                 CountryCode = 400
	Kazakhstan                             CountryCode = 398
	Kenya                                  CountryCode = 404
	Kiribati                               CountryCode = 296
	Korea                                  CountryCode = 410
	KoreaNorth                             CountryCode = 408
	Kuwait                                 CountryCode = 414
	Kyrgyzstan                             CountryCode = 417
	Laos                                   CountryCode = 418
	Latvia                                 CountryCode = 428
	Lebanon                                CountryCode = 422
	Lesotho                                CountryCode = 426
	Liberia                                CountryCode = 430
	Libya                                  CountryCode = 434
	Liechtenstein                          CountryCode = 438
	Lithuania                              CountryCode = 440
	Luxembourg                             CountryCode = 442
	Macau                                  CountryCode = 446
	Macao                                  CountryCode = 446
	Macedonia                              CountryCode = 807
	Madagascar                             CountryCode = 450
	Malawi                                 CountryCode = 454
	Malaysia                               CountryCode = 458
	Maldives                               CountryCode = 462
	Mali                                   CountryCode = 466
	Malta                                  CountryCode = 470
	MarshallIslands                        CountryCode = 584
	Martinique                             CountryCode = 474
	Mauritania                             CountryCode = 478
	Mauritius                              CountryCode = 480
	Mayotte                                CountryCode = 175
	Mexico                                 CountryCode = 484
	Micronesia                             CountryCode = 583
	Moldova                                CountryCode = 498
	Monaco                                 CountryCode = 492
	Mongolia                               CountryCode = 496
	Montserrat                             CountryCode = 500
	Morocco                                CountryCode = 504
	Mozambique                             CountryCode = 508
	Myanmar                                CountryCode = 104
	Namibia                                CountryCode = 516
	Nauru                                  CountryCode = 520
	Nepal                                  CountryCode = 524
	Netherlands                            CountryCode = 528
	NetherlandsAntilles                    CountryCode = 530
	NewCaledonia                           CountryCode = 540
	NewZealand                             CountryCode = 554
	Nicaragua                              CountryCode = 558
	Niger                                  CountryCode = 562
	Nigeria                                CountryCode = 566
	Niue                                   CountryCode = 570
	NorfolkIsland                          CountryCode = 574
	NorthernMarianaIslands                 CountryCode = 580
	Norway                                 CountryCode = 578
	Oman                                   CountryCode = 512
	Pakistan                               CountryCode = 586
	Palau                                  CountryCode = 585
	Palestine                              CountryCode = 275
	Panama                                 CountryCode = 591
	PapuaNewGuinea                         CountryCode = 598
	Paraguay                               CountryCode = 600
	Peru                                   CountryCode = 604
	Philippines                            CountryCode = 608
	Pitcairn                               CountryCode = 612
	Poland                                 CountryCode = 616
	Portugal                               CountryCode = 620
	PuertoRico                             CountryCode = 630
	Qatar                                  CountryCode = 634
	Reunion                                CountryCode = 638
	Romania                                CountryCode = 642
	Russia                                 CountryCode = 643
	Rwanda                                 CountryCode = 646
	SaintHelena                            CountryCode = 654
	SaintKittsAndNevis                     CountryCode = 659
	SaintLucia                             CountryCode = 662
	SaintPierreAndMiquelon                 CountryCode = 666
	SaintVincentAndTheGrenadines           CountryCode = 670
	Samoa                                  CountryCode = 882
	SanMarino                              CountryCode = 674
	SaoTomeAndPrincipe                     CountryCode = 678
	SaudiArabia                            CountryCode = 682
	Senegal                                CountryCode = 686
	Seychelles                             CountryCode = 690
	SierraLeone                            CountryCode = 694
	Singapore                              CountryCode = 702
	Slovakia                               CountryCode = 703
	Slovenia                               CountryCode = 705
	SolomonIslands                         CountryCode = 90
	Somalia                                CountryCode = 706
	SouthAfrica                            CountryCode = 710
	UAR                                    CountryCode = 710
	SouthGeorgiaAndTheSouthSandwichIslands CountryCode = 239
	Spain                                  CountryCode = 724
	SriLanka                               CountryCode = 144
	Sudan                                  CountryCode = 736
	Suriname                               CountryCode = 740
	SvalbardAndJanMayenIslands             CountryCode = 744
	Swaziland                              CountryCode = 748
	Sweden                                 CountryCode = 752
	Scotland                               CountryCode = 826
	Switzerland                            CountryCode = 756
	Syria                                  CountryCode = 760
	Taiwan                                 CountryCode = 158
	Tajikistan                             CountryCode = 762
	Tanzania                               CountryCode = 834
	Thailand                               CountryCode = 764
	TimorLeste                             CountryCode = 626
	Togo                                   CountryCode = 768
	Tokelau                                CountryCode = 772
	Tonga                                  CountryCode = 776
	TrinidadAndTobago                      CountryCode = 780
	Tunisia                                CountryCode = 788
	Turkey                                 CountryCode = 792
	Turkmenistan                           CountryCode = 795
	TurksAndCaicosIslands                  CountryCode = 796
	Tuvalu                                 CountryCode = 798
	Uganda                                 CountryCode = 800
	Ukraine                                CountryCode = 804
	UnitedArabEmirates                     CountryCode = 784
	UnitedKingdom                          CountryCode = 826
	UnitedStatesOfAmerica                  CountryCode = 840
	UnitedStatesMinorOutlyingIslands       CountryCode = 581
	Uruguay                                CountryCode = 858
	Wales                                  CountryCode = 826
	Uzbekistan                             CountryCode = 860
	Vanuatu                                CountryCode = 548
	HolySee                                CountryCode = 336
	Venezuela                              CountryCode = 862
	Vietnam                                CountryCode = 704
	VirginIslandsBritish                   CountryCode = 92
	VirginIslandsUS                        CountryCode = 850
	WallisandFutunaIslands                 CountryCode = 876
	WesternSahara                          CountryCode = 732
	Yemen                                  CountryCode = 887
	Yugoslavia                             CountryCode = 891
	Zambia                                 CountryCode = 894
	Zimbabwe                               CountryCode = 716
	Afghanistan                            CountryCode = 4
	Serbia                                 CountryCode = 688
	AlandIslands                           CountryCode = 248
	Bonaire                                CountryCode = 535
	Guernsey                               CountryCode = 831
	Jersey                                 CountryCode = 832
	Curacao                                CountryCode = 531
	SaintBarthelemy                        CountryCode = 652
	SaintMartinFrench                      CountryCode = 663
	SintMaartenDutch                       CountryCode = 534
	Montenegro                             CountryCode = 499
	SouthSudan                             CountryCode = 728
)

// Alpha-2 digit ISO 3166-1. Three codes present, for example Russia == RU == RUS == 643.
const (
	AL CountryCode = 8
	DZ CountryCode = 12
	AS CountryCode = 16
	AD CountryCode = 20
	AO CountryCode = 24
	AI CountryCode = 660
	AQ CountryCode = 10
	AG CountryCode = 28
	AR CountryCode = 32
	AM CountryCode = 51
	AW CountryCode = 533
	AU CountryCode = 36
	AT CountryCode = 40
	AZ CountryCode = 31
	BS CountryCode = 44
	BH CountryCode = 48
	BD CountryCode = 50
	BB CountryCode = 52
	BY CountryCode = 112
	BE CountryCode = 56
	BZ CountryCode = 84
	BJ CountryCode = 204
	BM CountryCode = 60
	BT CountryCode = 64
	BO CountryCode = 68
	BA CountryCode = 70
	BW CountryCode = 72
	BV CountryCode = 74
	BR CountryCode = 76
	IO CountryCode = 86
	BN CountryCode = 96
	BG CountryCode = 100
	BF CountryCode = 854
	BI CountryCode = 108
	KH CountryCode = 116
	CM CountryCode = 120
	CA CountryCode = 124
	CV CountryCode = 132
	KY CountryCode = 136
	CF CountryCode = 140
	TD CountryCode = 148
	CL CountryCode = 152
	CN CountryCode = 156
	CX CountryCode = 162
	CC CountryCode = 166
	CO CountryCode = 170
	KM CountryCode = 174
	CG CountryCode = 178
	CD CountryCode = 180
	CK CountryCode = 184
	CR CountryCode = 188
	CI CountryCode = 384
	HR CountryCode = 191
	CU CountryCode = 192
	CY CountryCode = 196
	CZ CountryCode = 203
	DK CountryCode = 208
	DJ CountryCode = 262
	DM CountryCode = 212
	DO CountryCode = 214
	EC CountryCode = 218
	EG CountryCode = 818
	SV CountryCode = 222
	GQ CountryCode = 226
	ER CountryCode = 232
	EE CountryCode = 233
	ET CountryCode = 231
	FO CountryCode = 234
	FK CountryCode = 238
	FJ CountryCode = 242
	FI CountryCode = 246
	FR CountryCode = 250
	GF CountryCode = 254
	PF CountryCode = 258
	TF CountryCode = 260
	GA CountryCode = 266
	GM CountryCode = 270
	GE CountryCode = 268
	DE CountryCode = 276
	GH CountryCode = 288
	GI CountryCode = 292
	GR CountryCode = 300
	GL CountryCode = 304
	GD CountryCode = 308
	GP CountryCode = 312
	GU CountryCode = 316
	GT CountryCode = 320
	GN CountryCode = 324
	GW CountryCode = 624
	GY CountryCode = 328
	HT CountryCode = 332
	HM CountryCode = 334
	HN CountryCode = 340
	HK CountryCode = 344
	HU CountryCode = 348
	IS CountryCode = 352
	IN CountryCode = 356
	ID CountryCode = 360
	IR CountryCode = 364
	IQ CountryCode = 368
	IE CountryCode = 372
	IL CountryCode = 376
	IT CountryCode = 380
	JM CountryCode = 388
	JP CountryCode = 392
	JO CountryCode = 400
	KZ CountryCode = 398
	KE CountryCode = 404
	KI CountryCode = 296
	KR CountryCode = 410
	KP CountryCode = 408
	KW CountryCode = 414
	KG CountryCode = 417
	LA CountryCode = 418
	LV CountryCode = 428
	LB CountryCode = 422
	LS CountryCode = 426
	LR CountryCode = 430
	LY CountryCode = 434
	LI CountryCode = 438
	LT CountryCode = 440
	LU CountryCode = 442
	MO CountryCode = 446
	MK CountryCode = 807
	MG CountryCode = 450
	MW CountryCode = 454
	MY CountryCode = 458
	MV CountryCode = 462
	ML CountryCode = 466
	MT CountryCode = 470
	MH CountryCode = 584
	MQ CountryCode = 474
	MR CountryCode = 478
	MU CountryCode = 480
	YT CountryCode = 175
	MX CountryCode = 484
	FM CountryCode = 583
	MD CountryCode = 498
	MC CountryCode = 492
	MN CountryCode = 496
	MS CountryCode = 500
	MA CountryCode = 504
	MZ CountryCode = 508
	MM CountryCode = 104
	NA CountryCode = 516
	NR CountryCode = 520
	NP CountryCode = 524
	NL CountryCode = 528
	AN CountryCode = 530
	NC CountryCode = 540
	NZ CountryCode = 554
	NI CountryCode = 558
	NE CountryCode = 562
	NG CountryCode = 566
	NU CountryCode = 570
	NF CountryCode = 574
	MP CountryCode = 580
	NO CountryCode = 578
	OM CountryCode = 512
	PK CountryCode = 586
	PW CountryCode = 585
	PS CountryCode = 275
	PA CountryCode = 591
	PG CountryCode = 598
	PY CountryCode = 600
	PE CountryCode = 604
	PH CountryCode = 608
	PN CountryCode = 612
	PL CountryCode = 616
	PT CountryCode = 620
	PR CountryCode = 630
	QA CountryCode = 634
	RE CountryCode = 638
	RO CountryCode = 642
	RU CountryCode = 643
	RW CountryCode = 646
	SH CountryCode = 654
	KN CountryCode = 659
	LC CountryCode = 662
	PM CountryCode = 666
	VC CountryCode = 670
	WS CountryCode = 882
	SM CountryCode = 674
	ST CountryCode = 678
	SA CountryCode = 682
	SN CountryCode = 686
	SC CountryCode = 690
	SL CountryCode = 694
	SG CountryCode = 702
	SK CountryCode = 703
	SI CountryCode = 705
	SB CountryCode = 90
	SO CountryCode = 706
	ZA CountryCode = 710
	GS CountryCode = 239
	ES CountryCode = 724
	LK CountryCode = 144
	SD CountryCode = 736
	SR CountryCode = 740
	SJ CountryCode = 744
	SZ CountryCode = 748
	SE CountryCode = 752
	XS CountryCode = 826
	CH CountryCode = 756
	SY CountryCode = 760
	TW CountryCode = 158
	TJ CountryCode = 762
	TZ CountryCode = 834
	TH CountryCode = 764
	TL CountryCode = 626
	TG CountryCode = 768
	TK CountryCode = 772
	TO CountryCode = 776
	TT CountryCode = 780
	TN CountryCode = 788
	TR CountryCode = 792
	TM CountryCode = 795
	TC CountryCode = 796
	TV CountryCode = 798
	UG CountryCode = 800
	UA CountryCode = 804
	AE CountryCode = 784
	GB CountryCode = 826
	US CountryCode = 840
	UM CountryCode = 581
	UY CountryCode = 858
	UZ CountryCode = 860
	VU CountryCode = 548
	VA CountryCode = 336
	VE CountryCode = 862
	VN CountryCode = 704
	VG CountryCode = 92
	VI CountryCode = 850
	WF CountryCode = 876
	EH CountryCode = 732
	YE CountryCode = 887
	YU CountryCode = 891
	ZM CountryCode = 894
	ZW CountryCode = 716
	AF CountryCode = 4
	RS CountryCode = 688
	AX CountryCode = 248
	BQ CountryCode = 535
	GG CountryCode = 831
	JE CountryCode = 832
	CW CountryCode = 531
	IM CountryCode = 833
	BL CountryCode = 652
	MF CountryCode = 663
	SX CountryCode = 534
	ME CountryCode = 499
	SS CountryCode = 728
)

// Alpha-3 digit ISO 3166-1. Three codes present, for example Russia == RU == RUS == 643.
const (
	ALB CountryCode = 8
	DZA CountryCode = 12
	ASM CountryCode = 16
	AND CountryCode = 20
	AGO CountryCode = 24
	AIA CountryCode = 660
	ATA CountryCode = 10
	ATG CountryCode = 28
	ARG CountryCode = 32
	ARM CountryCode = 51
	ABW CountryCode = 533
	AUS CountryCode = 36
	AUT CountryCode = 40
	AZE CountryCode = 31
	BHS CountryCode = 44
	BHR CountryCode = 48
	BGD CountryCode = 50
	BRB CountryCode = 52
	BLR CountryCode = 112
	BEL CountryCode = 56
	BLZ CountryCode = 84
	BEN CountryCode = 204
	BMU CountryCode = 60
	BTN CountryCode = 64
	BOL CountryCode = 68
	BIH CountryCode = 70
	BWA CountryCode = 72
	BVT CountryCode = 74
	BRA CountryCode = 76
	IOT CountryCode = 86
	BRN CountryCode = 96
	BGR CountryCode = 100
	BFA CountryCode = 854
	BDI CountryCode = 108
	KHM CountryCode = 116
	CMR CountryCode = 120
	CAN CountryCode = 124
	CPV CountryCode = 132
	CYM CountryCode = 136
	CAF CountryCode = 140
	TCD CountryCode = 148
	CHL CountryCode = 152
	CHN CountryCode = 156
	CXR CountryCode = 162
	CCK CountryCode = 166
	COL CountryCode = 170
	COM CountryCode = 174
	COG CountryCode = 178
	COD CountryCode = 180
	COK CountryCode = 184
	CRI CountryCode = 188
	CIV CountryCode = 384
	HRV CountryCode = 191
	CUB CountryCode = 192
	CYP CountryCode = 196
	CZE CountryCode = 203
	DNK CountryCode = 208
	DJI CountryCode = 262
	DMA CountryCode = 212
	DOM CountryCode = 214
	ECU CountryCode = 218
	EGY CountryCode = 818
	SLV CountryCode = 222
	GNQ CountryCode = 226
	ERI CountryCode = 232
	EST CountryCode = 233
	ETH CountryCode = 231
	FRO CountryCode = 234
	FLK CountryCode = 238
	FJI CountryCode = 242
	FIN CountryCode = 246
	FRA CountryCode = 250
	GUF CountryCode = 254
	PYF CountryCode = 258
	ATF CountryCode = 260
	GAB CountryCode = 266
	GMB CountryCode = 270
	GEO CountryCode = 268
	DEU CountryCode = 276
	GHA CountryCode = 288
	GIB CountryCode = 292
	GRC CountryCode = 300
	GRL CountryCode = 304
	GRD CountryCode = 308
	GLP CountryCode = 312
	GUM CountryCode = 316
	GTM CountryCode = 320
	GIN CountryCode = 324
	GNB CountryCode = 624
	GUY CountryCode = 328
	HTI CountryCode = 332
	HMD CountryCode = 334
	HND CountryCode = 340
	HKG CountryCode = 344
	HUN CountryCode = 348
	ISL CountryCode = 352
	IND CountryCode = 356
	IDN CountryCode = 360
	IRN CountryCode = 364
	IRQ CountryCode = 368
	IRL CountryCode = 372
	ISR CountryCode = 376
	ITA CountryCode = 380
	JAM CountryCode = 388
	JPN CountryCode = 392
	JOR CountryCode = 400
	KAZ CountryCode = 398
	KEN CountryCode = 404
	KIR CountryCode = 296
	KOR CountryCode = 410
	PRK CountryCode = 408
	KWT CountryCode = 414
	KGZ CountryCode = 417
	LAO CountryCode = 418
	LVA CountryCode = 428
	LBN CountryCode = 422
	LSO CountryCode = 426
	LBR CountryCode = 430
	LBY CountryCode = 434
	LIE CountryCode = 438
	LTU CountryCode = 440
	LUX CountryCode = 442
	MAC CountryCode = 446
	MKD CountryCode = 807
	MDG CountryCode = 450
	MWI CountryCode = 454
	MYS CountryCode = 458
	MDV CountryCode = 462
	MLI CountryCode = 466
	MLT CountryCode = 470
	MHL CountryCode = 584
	MTQ CountryCode = 474
	MRT CountryCode = 478
	MUS CountryCode = 480
	MYT CountryCode = 175
	MEX CountryCode = 484
	FSM CountryCode = 583
	MDA CountryCode = 498
	MCO CountryCode = 492
	MNG CountryCode = 496
	MSR CountryCode = 500
	MAR CountryCode = 504
	MOZ CountryCode = 508
	MMR CountryCode = 104
	NAM CountryCode = 516
	NRU CountryCode = 520
	NPL CountryCode = 524
	NLD CountryCode = 528
	ANT CountryCode = 530
	NCL CountryCode = 540
	NZL CountryCode = 554
	NIC CountryCode = 558
	NER CountryCode = 562
	NGA CountryCode = 566
	NIU CountryCode = 570
	NFK CountryCode = 574
	MNP CountryCode = 580
	NOR CountryCode = 578
	OMN CountryCode = 512
	PAK CountryCode = 586
	PLW CountryCode = 585
	PSE CountryCode = 275
	PAN CountryCode = 591
	PNG CountryCode = 598
	PRY CountryCode = 600
	PER CountryCode = 604
	PHL CountryCode = 608
	PCN CountryCode = 612
	POL CountryCode = 616
	PRT CountryCode = 620
	PRI CountryCode = 630
	QAT CountryCode = 634
	REU CountryCode = 638
	ROU CountryCode = 642
	RUS CountryCode = 643
	RWA CountryCode = 646
	SHN CountryCode = 654
	KNA CountryCode = 659
	LCA CountryCode = 662
	SPM CountryCode = 666
	VCT CountryCode = 670
	WSM CountryCode = 882
	SMR CountryCode = 674
	STP CountryCode = 678
	SAU CountryCode = 682
	SEN CountryCode = 686
	SYC CountryCode = 690
	SLE CountryCode = 694
	SGP CountryCode = 702
	SVK CountryCode = 703
	SVN CountryCode = 705
	SLB CountryCode = 90
	SOM CountryCode = 706
	ZAF CountryCode = 710
	SGS CountryCode = 239
	ESP CountryCode = 724
	LKA CountryCode = 144
	SDN CountryCode = 736
	SUR CountryCode = 740
	SJM CountryCode = 744
	SWZ CountryCode = 748
	SWE CountryCode = 752
	XSC CountryCode = 826
	CHE CountryCode = 756
	SYR CountryCode = 760
	TWN CountryCode = 158
	TJK CountryCode = 762
	TZA CountryCode = 834
	THA CountryCode = 764
	TLS CountryCode = 626
	TGO CountryCode = 768
	TKL CountryCode = 772
	TON CountryCode = 776
	TTO CountryCode = 780
	TUN CountryCode = 788
	TUR CountryCode = 792
	TKM CountryCode = 795
	TCA CountryCode = 796
	TUV CountryCode = 798
	UGA CountryCode = 800
	UKR CountryCode = 804
	ARE CountryCode = 784
	GBR CountryCode = 826
	USA CountryCode = 840
	UMI CountryCode = 581
	URY CountryCode = 858
	XWA CountryCode = 826
	UZB CountryCode = 860
	VUT CountryCode = 548
	VAT CountryCode = 336
	VEN CountryCode = 862
	VNM CountryCode = 704
	VGB CountryCode = 92
	VIR CountryCode = 850
	WLF CountryCode = 876
	ESH CountryCode = 732
	YEM CountryCode = 887
	YUG CountryCode = 891
	ZMB CountryCode = 894
	ZWE CountryCode = 716
	AFG CountryCode = 4
	SRB CountryCode = 688
	BES CountryCode = 535
	ALA CountryCode = 248
	JEY CountryCode = 832
	GGY CountryCode = 831
	CUW CountryCode = 531
	IMN CountryCode = 833
	BLM CountryCode = 652
	MAF CountryCode = 663
	SXM CountryCode = 534
	MNE CountryCode = 499
	SSD CountryCode = 728
)

// Total - returns number of codes in the package, countries.Total() == len(countries.All()) but static value for perfomance
func Total() int {
	return 253
}

// Emoji - return a country Alpha-2 (ISO2) as Emoji flag (example "RU" as "🇷🇺")
func (c CountryCode) Emoji() string {
	iso2 := c.Alpha3()
	buf := [...]byte{240, 159, 135, 0, 240, 159, 135, 0}
	buf[3] = iso2[0] + (166 - 'A')
	buf[7] = iso2[1] + (166 - 'A')
	return string(buf[:])
}

// Emoji3 - return a country Alpha-3 (ISO3) as Emoji (example "RUS" as "🇷🇺🇸")
func (c CountryCode) Emoji3() string {
	iso3 := c.Alpha3()
	buf := [...]byte{240, 159, 135, 0, 240, 159, 135, 0, 240, 159, 135, 0}
	buf[3] = iso3[0] + (166 - 'A')
	buf[7] = iso3[1] + (166 - 'A')
	buf[11] = iso3[2] + (166 - 'A')
	return string(buf[:])
}

// Type implements Typer interface.
func (c CountryCode) Type() string {
	return TypeCountryCode
}

// String - implements fmt.Stringer, returns a english name of country
func (c CountryCode) String() string {
	switch c {
	case 8:
		return "Albania"
	case 12:
		return "Algeria"
	case 16:
		return "American Samoa"
	case 20:
		return "Andorra"
	case 24:
		return "Angola"
	case 660:
		return "Anguilla"
	case 10:
		return "Antarctica"
	case 28:
		return "Antigua and Barbuda"
	case 32:
		return "Argentina"
	case 51:
		return "Armenia"
	case 533:
		return "Aruba"
	case 36:
		return "Australia"
	case 40:
		return "Austria"
	case 31:
		return "Azerbaijan"
	case 44:
		return "Bahamas"
	case 48:
		return "Bahrain"
	case 50:
		return "Bangladesh"
	case 52:
		return "Barbados"
	case 112:
		return "Belarus"
	case 56:
		return "Belgium"
	case 84:
		return "Belize"
	case 204:
		return "Benin"
	case 60:
		return "Bermuda"
	case 64:
		return "Bhutan"
	case 68:
		return "Bolivia"
	case 70:
		return "Bosnia and Herzegovina"
	case 72:
		return "Botswana"
	case 74:
		return "Bouvet Island"
	case 76:
		return "Brazil"
	case 86:
		return "British Indian Ocean Territory"
	case 96:
		return "Brunei Darussalam"
	case 100:
		return "Bulgaria"
	case 854:
		return "Burkina Faso"
	case 108:
		return "Burundi"
	case 116:
		return "Cambodia"
	case 120:
		return "Cameroon"
	case 124:
		return "Canada"
	case 132:
		return "Cape Verde"
	case 136:
		return "Cayman Islands"
	case 140:
		return "Central African Republic"
	case 148:
		return "Chad"
	case 152:
		return "Chile"
	case 156:
		return "China"
	case 162:
		return "Christmas Island"
	case 166:
		return "Cocos (Keeling) Islands"
	case 170:
		return "Colombia"
	case 174:
		return "Comoros"
	case 178:
		return "Congo"
	case 180:
		return "Congo (Democratic Republic of the)"
	case 184:
		return "Cook Islands"
	case 188:
		return "Costa Rica"
	case 384:
		return "Cote d'Ivoire" // Ivory Coast
	case 191:
		return "Croatia"
	case 192:
		return "Cuba"
	case 196:
		return "Cyprus"
	case 203:
		return "Czech Republic"
	case 208:
		return "Denmark"
	case 262:
		return "Djibouti"
	case 212:
		return "Dominica"
	case 214:
		return "Dominican Republic"
	case 218:
		return "Ecuador"
	case 818:
		return "Egypt"
	case 222:
		return "El Salvador"
	case 226:
		return "Equatorial Guinea"
	case 232:
		return "Eritrea"
	case 233:
		return "Estonia"
	case 231:
		return "Ethiopia"
	case 234:
		return "Faroe Islands"
	case 238:
		return "Falkland Islands (Malvinas)"
	case 242:
		return "Fiji"
	case 246:
		return "Finland"
	case 250:
		return "France"
	case 254:
		return "French Guiana"
	case 258:
		return "French Polynesia"
	case 260:
		return "French Southern Territories"
	case 266:
		return "Gabon"
	case 270:
		return "Gambia"
	case 268:
		return "Georgia"
	case 276:
		return "Germany"
	case 288:
		return "Ghana"
	case 292:
		return "Gibraltar"
	case 300:
		return "Greece"
	case 304:
		return "Greenland"
	case 308:
		return "Grenada"
	case 312:
		return "Guadeloupe"
	case 316:
		return "Guam"
	case 320:
		return "Guatemala"
	case 324:
		return "Guinea"
	case 624:
		return "Guinea-Bissau"
	case 328:
		return "Guyana"
	case 332:
		return "Haiti"
	case 334:
		return "Heard Island and McDonald Islands"
	case 340:
		return "Honduras"
	case 344:
		return "Hong Kong (Special Administrative Region of China)"
	case 348:
		return "Hungary"
	case 352:
		return "Iceland"
	case 356:
		return "India"
	case 360:
		return "Indonesia"
	case 364:
		return "Iran (Islamic Republic of)"
	case 368:
		return "Iraq"
	case 372:
		return "Ireland"
	case 376:
		return "Israel"
	case 380:
		return "Italy"
	case 388:
		return "Jamaica"
	case 392:
		return "Japan"
	case 400:
		return "Jordan"
	case 398:
		return "Kazakhstan"
	case 404:
		return "Kenya"
	case 296:
		return "Kiribati"
	case 410:
		return "Korea (Republic of)"
	case 408:
		return "Korea (Democratic People's Republic of)"
	case 414:
		return "Kuwait"
	case 417:
		return "Kyrgyzstan"
	case 418:
		return "Lao People's Democratic Republic"
	case 428:
		return "Latvia"
	case 422:
		return "Lebanon"
	case 426:
		return "Lesotho"
	case 430:
		return "Liberia"
	case 434:
		return "Libyan Arab Jamahiriya"
	case 438:
		return "Liechtenstein"
	case 440:
		return "Lithuania"
	case 442:
		return "Luxembourg"
	case 446:
		return "Macau (Special Administrative Region of China)"
	case 807:
		return "Macedonia (The former Yugoslav Republic of)"
	case 450:
		return "Madagascar"
	case 454:
		return "Malawi"
	case 458:
		return "Malaysia"
	case 462:
		return "Maldives"
	case 466:
		return "Mali"
	case 470:
		return "Malta"
	case 584:
		return "Marshall Islands"
	case 474:
		return "Martinique"
	case 478:
		return "Mauritania"
	case 480:
		return "Mauritius"
	case 175:
		return "Mayotte"
	case 484:
		return "Mexico"
	case 583:
		return "Micronesia (Federated States of)"
	case 498:
		return "Moldova (Republic of)"
	case 492:
		return "Monaco"
	case 496:
		return "Mongolia"
	case 500:
		return "Montserrat"
	case 504:
		return "Morocco"
	case 508:
		return "Mozambique"
	case 104:
		return "Myanmar"
	case 516:
		return "Namibia"
	case 520:
		return "Nauru"
	case 524:
		return "Nepal"
	case 528:
		return "Netherlands"
	case 530:
		return "Netherlands Antilles"
	case 540:
		return "New Caledonia"
	case 554:
		return "New Zealand"
	case 558:
		return "Nicaragua"
	case 562:
		return "Niger"
	case 566:
		return "Nigeria"
	case 570:
		return "Niue"
	case 574:
		return "Norfolk Island"
	case 580:
		return "Northern Mariana Islands"
	case 578:
		return "Norway"
	case 512:
		return "Oman"
	case 586:
		return "Pakistan"
	case 585:
		return "Palau"
	case 275:
		return "Palestinian Territory (Occupied)"
	case 591:
		return "Panama"
	case 598:
		return "Papua New Guinea"
	case 600:
		return "Paraguay"
	case 604:
		return "Peru"
	case 608:
		return "Philippines"
	case 612:
		return "Pitcairn"
	case 616:
		return "Poland"
	case 620:
		return "Portugal"
	case 630:
		return "Puerto Rico"
	case 634:
		return "Qatar"
	case 638:
		return "Reunion"
	case 642:
		return "Romania"
	case 643:
		return "Russian Federation"
	case 646:
		return "Rwanda"
	case 654:
		return "Saint Helena"
	case 659:
		return "Saint Kitts and Nevis"
	case 662:
		return "Saint Lucia"
	case 666:
		return "Saint Pierre and Miquelon"
	case 670:
		return "Saint Vincent and the Grenadines"
	case 882:
		return "Samoa"
	case 674:
		return "San Marino"
	case 678:
		return "Sao Tome and Principe"
	case 682:
		return "Saudi Arabia"
	case 686:
		return "Senegal"
	case 690:
		return "Seychelles"
	case 694:
		return "Sierra Leone"
	case 702:
		return "Singapore"
	case 703:
		return "Slovakia"
	case 705:
		return "Slovenia"
	case 90:
		return "Solomon Islands"
	case 706:
		return "Somalia"
	case 710:
		return "South Africa"
	case 239:
		return "South Georgia and The South Sandwich Islands"
	case 724:
		return "Spain"
	case 144:
		return "Sri Lanka"
	case 736:
		return "Sudan"
	case 740:
		return "Suriname"
	case 744:
		return "Svalbard and Jan Mayen Islands"
	case 748:
		return "Swaziland"
	case 752:
		return "Sweden"
	case 756:
		return "Switzerland"
	case 760:
		return "Syrian Arab Republic"
	case 158:
		return "Taiwan (Province of China)"
	case 762:
		return "Tajikistan"
	case 834:
		return "Tanzania (United Republic of)"
	case 764:
		return "Thailand"
	case 626:
		return "Timor-Leste (East Timor)"
	case 768:
		return "Togo"
	case 772:
		return "Tokelau"
	case 776:
		return "Tonga"
	case 780:
		return "Trinidad and Tobago"
	case 788:
		return "Tunisia"
	case 792:
		return "Turkey"
	case 795:
		return "Turkmenistan"
	case 796:
		return "Turks and Caicos Islands"
	case 798:
		return "Tuvalu"
	case 800:
		return "Uganda"
	case 804:
		return "Ukraine"
	case 784:
		return "United Arab Emirates"
	case 826:
		return "United Kingdom"
	case 840:
		return "United States"
	case 581:
		return "United States Minor Outlying Islands"
	case 858:
		return "Uruguay"
	case 860:
		return "Uzbekistan"
	case 548:
		return "Vanuatu"
	case 336:
		return "Holy See (Vatican City State)"
	case 862:
		return "Venezuela"
	case 704:
		return "Viet Nam"
	case 92:
		return "Virgin Islands (British)"
	case 850:
		return "Virgin Islands (US)"
	case 876:
		return "Wallis and Futuna Islands"
	case 732:
		return "Western Sahara"
	case 887:
		return "Yemen"
	case 891:
		return "Yugoslavia"
	case 894:
		return "Zambia"
	case 716:
		return "Zimbabwe"
	case 4:
		return "Afghanistan"
	case 688:
		return "Serbia"
	case 248:
		return "Aland Islands"
	case 535:
		return "Bonaire, Sint Eustatius And Saba"
	case 831:
		return "Guernsey"
	case 832:
		return "Jersey"
	case 531:
		return "Curacao"
	case 833:
		return "Isle Of Man"
	case 652:
		return "Saint Barthelemy"
	case 663:
		return "Saint Martin (French)"
	case 534:
		return "Sint Maarten (Dutch)"
	case 499:
		return "Montenegro"
	case 728:
		return "South Sudan"
	}
	return UnknownMsg
}

// StringRus - returns a russian name of country
func (c CountryCode) StringRus() string {
	switch c {
	case 8:
		return "Албания"
	case 12:
		return "Алжир"
	case 16:
		return "Американский Самоа"
	case 20:
		return "Андорра"
	case 24:
		return "Ангола"
	case 660:
		return "Ангилья"
	case 10:
		return "Антарктика"
	case 28:
		return "Антигуа и Барбуда"
	case 32:
		return "Аргентина"
	case 51:
		return "Армения"
	case 533:
		return "Аруба"
	case 36:
		return "Австралия"
	case 40:
		return "Австрия"
	case 31:
		return "Азербайджан"
	case 44:
		return "Багамские острова"
	case 48:
		return "Бахрейн"
	case 50:
		return "Бангладеш"
	case 52:
		return "Барбадос"
	case 112:
		return "Беларусь"
	case 56:
		return "Бельгия"
	case 84:
		return "Белиз"
	case 204:
		return "Бенин"
	case 60:
		return "Бермуды"
	case 64:
		return "Бутан"
	case 68:
		return "Боливия"
	case 70:
		return "Босния и Герцеговина"
	case 72:
		return "Ботсвана"
	case 74:
		return "остров Буве"
	case 76:
		return "Бразилия"
	case 86:
		return "Британские территории Индийского океана"
	case 96:
		return "Бруней"
	case 100:
		return "Болгария"
	case 854:
		return "Буркина Фасо"
	case 108:
		return "Бурунди"
	case 116:
		return "Камбоджа"
	case 120:
		return "Камерун"
	case 124:
		return "Канада"
	case 132:
		return "острова Зеленого Мыса"
	case 136:
		return "Каймановы острова"
	case 140:
		return "Центральная Африканская Республика"
	case 148:
		return "Чад"
	case 152:
		return "Чили"
	case 156:
		return "Китайская Народная Республика"
	case 162:
		return "остров Рождества"
	case 166:
		return "Кокосовые острова"
	case 170:
		return "Колумбия"
	case 174:
		return "Коморские острова"
	case 178:
		return "Конго"
	case 180:
		return "Демократическая республика Конго"
	case 184:
		return "острова Кука"
	case 188:
		return "Коста Рика"
	case 384:
		return "Кот-д'Ивуар"
	case 191:
		return "Хорватия"
	case 192:
		return "Куба"
	case 196:
		return "Кипр"
	case 203:
		return "Чехия"
	case 208:
		return "Дания"
	case 262:
		return "Джибути"
	case 212:
		return "Доминика"
	case 214:
		return "Доминиканская республика"
	case 218:
		return "Эквадор"
	case 818:
		return "Египет"
	case 222:
		return "Сальвадор"
	case 226:
		return "Экваториальная Гвинея"
	case 232:
		return "Эритрея"
	case 233:
		return "Эстония"
	case 231:
		return "Эфиопия"
	case 234:
		return "Фарерские острова"
	case 238:
		return "Фолклендские (Мальвинские) острова"
	case 242:
		return "Фиджи"
	case 246:
		return "Финляндия"
	case 250:
		return "Франция"
	case 254:
		return "Французская Гвиана"
	case 258:
		return "Французская Полинезия"
	case 260:
		return "Французские Южные Территории"
	case 266:
		return "Габон"
	case 270:
		return "Гамбия"
	case 268:
		return "Грузия"
	case 276:
		return "Германия"
	case 288:
		return "Гана"
	case 292:
		return "Гибралтар"
	case 300:
		return "Греция"
	case 304:
		return "Гренландия"
	case 308:
		return "Гренада"
	case 312:
		return "Гваделупа"
	case 316:
		return "Гуам"
	case 320:
		return "Гватемала"
	case 324:
		return "Гвинея"
	case 624:
		return "Гвинея-Бисау"
	case 328:
		return "Гайана"
	case 332:
		return "Гаити"
	case 334:
		return "острова Герда и МакДональда"
	case 340:
		return "Гондурас"
	case 344:
		return "Гонконг (Китай)"
	case 348:
		return "Венгрия"
	case 352:
		return "Исландия"
	case 356:
		return "Индия"
	case 360:
		return "Индонезия"
	case 364:
		return "Иран"
	case 368:
		return "Ирак"
	case 372:
		return "Ирландия"
	case 376:
		return "Израиль"
	case 380:
		return "Италия"
	case 388:
		return "Ямайка"
	case 392:
		return "Япония"
	case 400:
		return "Иордания"
	case 398:
		return "Казахстан"
	case 404:
		return "Кения"
	case 296:
		return "Кирибати"
	case 410:
		return "Корея"
	case 408:
		return "Корейская Народная Демократическая республика"
	case 414:
		return "Кувейт"
	case 417:
		return "Кыргызстан" // Киргизия
	case 418:
		return "Лаос"
	case 428:
		return "Латвия"
	case 422:
		return "Ливан"
	case 426:
		return "Лесото"
	case 430:
		return "Либерия"
	case 434:
		return "Ливия"
	case 438:
		return "Лихтенштейн"
	case 440:
		return "Литва"
	case 442:
		return "Люксембург"
	case 446:
		return "Макао (Китай)"
	case 807:
		return "Македония"
	case 450:
		return "Мадагаскар"
	case 454:
		return "Малави"
	case 458:
		return "Малайзия"
	case 462:
		return "Мальдивские острова"
	case 466:
		return "Мали"
	case 470:
		return "Мальта"
	case 584:
		return "Маршалловы острова"
	case 474:
		return "Мартиника"
	case 478:
		return "Мавритания"
	case 480:
		return "Маврикий"
	case 175:
		return "Майотта"
	case 484:
		return "Мексика"
	case 583:
		return "Микронезия"
	case 498:
		return "Молдова"
	case 492:
		return "Монако"
	case 496:
		return "Монголия"
	case 500:
		return "Монтсеррат"
	case 504:
		return "Марокко"
	case 508:
		return "Мозамбик"
	case 104:
		return "Мьянма"
	case 516:
		return "Намибия"
	case 520:
		return "Науру"
	case 524:
		return "Непал"
	case 528:
		return "Нидерланды"
	case 530:
		return "Антильские острова"
	case 540:
		return "Новая Каледония"
	case 554:
		return "Новая Зеландия"
	case 558:
		return "Никарагуа"
	case 562:
		return "Нигер"
	case 566:
		return "Нигерия"
	case 570:
		return "Ниуэ"
	case 574:
		return "остров Норфолк"
	case 580:
		return "Марианские острова"
	case 578:
		return "Норвегия"
	case 512:
		return "Оман"
	case 586:
		return "Пакистан"
	case 585:
		return "Палау"
	case 275:
		return "Палестина"
	case 591:
		return "Панама"
	case 598:
		return "Папуа - Новая Гвинея"
	case 600:
		return "Парагвай"
	case 604:
		return "Перу"
	case 608:
		return "Филиппины"
	case 612:
		return "остров Питкэрн"
	case 616:
		return "Польша"
	case 620:
		return "Португалия"
	case 630:
		return "Пуэрто-Рико"
	case 634:
		return "Катар"
	case 638:
		return "Реюньон"
	case 642:
		return "Румыния"
	case 643:
		return "Россия"
	case 646:
		return "Руанда"
	case 654:
		return "остров Святой Елены"
	case 659:
		return "Сент-Китс и Невис"
	case 662:
		return "Сент-Люсия"
	case 666:
		return "Сен-Пьер и Микелон"
	case 670:
		return "Сент-Винсент и Гренадины"
	case 882:
		return "острова Самоа"
	case 674:
		return "Сан-Марино"
	case 678:
		return "Сан-Томе и Принсипи"
	case 682:
		return "Саудовская Аравия"
	case 686:
		return "Сенегал"
	case 690:
		return "Сейшельские острова"
	case 694:
		return "Сьерра-Леоне"
	case 702:
		return "Сингапур"
	case 703:
		return "Словакия"
	case 705:
		return "Словения"
	case 90:
		return "Соломоновы острова"
	case 706:
		return "Сомали"
	case 710:
		return "ЮАР"
	case 239:
		return "Южная Георгия и Южные Сандвичевы острова"
	case 724:
		return "Испания"
	case 144:
		return "Шри Ланка"
	case 736:
		return "Судан"
	case 740:
		return "Суринам"
	case 744:
		return "острова Свалбард и Ян Майен"
	case 748:
		return "Свазиленд"
	case 752:
		return "Швеция"
	case 756:
		return "Швейцария"
	case 760:
		return "Сирия"
	case 158:
		return "Тайвань (Республика Китай)"
	case 762:
		return "Таджикистан"
	case 834:
		return "Танзания"
	case 764:
		return "Тайланд"
	case 626:
		return "Восточный Тимор"
	case 768:
		return "Того"
	case 772:
		return "Токелау"
	case 776:
		return "Тонга"
	case 780:
		return "Тринидад и Тобаго"
	case 788:
		return "Тунис"
	case 792:
		return "Турция"
	case 795:
		return "Туркменистан"
	case 796:
		return "острова Туркс и Кайкос"
	case 798:
		return "Тувалу"
	case 800:
		return "Уганда"
	case 804:
		return "Украина"
	case 784:
		return "Арабские Эмираты"
	case 826:
		return "Великобритания"
	case 840:
		return "Соединенные Штаты Америки"
	case 581:
		return "Отдаленные Острова США"
	case 858:
		return "Уругвай"
	case 860:
		return "Узбекистан"
	case 548:
		return "Вануату"
	case 336:
		return "Ватикан"
	case 862:
		return "Венесуэла"
	case 704:
		return "Вьетнам"
	case 92:
		return "Виргинские острова (Британские)"
	case 850:
		return "Виргинские острова (США)"
	case 876:
		return "острова Валлис и Футуна"
	case 732:
		return "Западная Сахара"
	case 887:
		return "Йемен"
	case 891:
		return "Югославия"
	case 894:
		return "Замбия"
	case 716:
		return "Зимбабве"
	case 4:
		return "Афганистан"
	case 688:
		return "Сербия"
	case 248:
		return "Аландские острова"
	case 535:
		return "Бонэйр, Синт-Эстатиус и Саба"
	case 831:
		return "Гернси"
	case 832:
		return "Джерси"
	case 531:
		return "Кюрасао"
	case 833:
		return "Остров Мэн"
	case 652:
		return "Сен-Бартелеми"
	case 663:
		return "Сен-Мартен"
	case 534:
		return "Синт-Мартен"
	case 499:
		return "Черногория"
	case 728:
		return "Южный Судан"
	case Internation:
		return "Internation"
	}
	return UnknownMsg
}

// Alpha2 - returns a default Alpha (Alpha-2/ISO2, 2 chars) code of country
func (c CountryCode) Alpha2() string {
	switch c {
	case 8:
		return "AL"
	case 12:
		return "DZ"
	case 16:
		return "AS"
	case 20:
		return "AD"
	case 24:
		return "AO"
	case 660:
		return "AI"
	case 10:
		return "AQ"
	case 28:
		return "AG"
	case 32:
		return "AR"
	case 51:
		return "AM"
	case 533:
		return "AW"
	case 36:
		return "AU"
	case 40:
		return "AT"
	case 31:
		return "AZ"
	case 44:
		return "BS"
	case 48:
		return "BH"
	case 50:
		return "BD"
	case 52:
		return "BB"
	case 112:
		return "BY"
	case 56:
		return "BE"
	case 84:
		return "BZ"
	case 204:
		return "BJ"
	case 60:
		return "BM"
	case 64:
		return "BT"
	case 68:
		return "BO"
	case 70:
		return "BA"
	case 72:
		return "BW"
	case 74:
		return "BV"
	case 76:
		return "BR"
	case 86:
		return "IO"
	case 96:
		return "BN"
	case 100:
		return "BG"
	case 854:
		return "BF"
	case 108:
		return "BI"
	case 116:
		return "KH"
	case 120:
		return "CM"
	case 124:
		return "CA"
	case 132:
		return "CV"
	case 136:
		return "KY"
	case 140:
		return "CF"
	case 148:
		return "TD"
	case 152:
		return "CL"
	case 156:
		return "CN"
	case 162:
		return "CX"
	case 166:
		return "CC"
	case 170:
		return "CO"
	case 174:
		return "KM"
	case 178:
		return "CG"
	case 180:
		return "CD"
	case 184:
		return "CK"
	case 188:
		return "CR"
	case 384:
		return "CI"
	case 191:
		return "HR"
	case 192:
		return "CU"
	case 196:
		return "CY"
	case 203:
		return "CZ"
	case 208:
		return "DK"
	case 262:
		return "DJ"
	case 212:
		return "DM"
	case 214:
		return "DO"
	case 218:
		return "EC"
	case 818:
		return "EG"
	case 222:
		return "SV"
	case 226:
		return "GQ"
	case 232:
		return "ER"
	case 233:
		return "EE"
	case 231:
		return "ET"
	case 238:
		return "FK"
	case 242:
		return "FJ"
	case 246:
		return "FI"
	case 250:
		return "FR"
	case 254:
		return "GF"
	case 258:
		return "PF"
	case 260:
		return "TF"
	case 266:
		return "GA"
	case 270:
		return "GM"
	case 268:
		return "GE"
	case 276:
		return "DE"
	case 288:
		return "GH"
	case 292:
		return "GI"
	case 300:
		return "GR"
	case 304:
		return "GL"
	case 308:
		return "GD"
	case 312:
		return "GP"
	case 316:
		return "GU"
	case 320:
		return "GT"
	case 324:
		return "GN"
	case 624:
		return "GW"
	case 328:
		return "GY"
	case 332:
		return "HT"
	case 334:
		return "HM"
	case 340:
		return "HN"
	case 344:
		return "HK"
	case 348:
		return "HU"
	case 352:
		return "IS"
	case 356:
		return "IN"
	case 360:
		return "ID"
	case 364:
		return "IR"
	case 368:
		return "IQ"
	case 372:
		return "IE"
	case 376:
		return "IL"
	case 380:
		return "IT"
	case 388:
		return "JM"
	case 392:
		return "JP"
	case 400:
		return "JO"
	case 398:
		return "KZ"
	case 404:
		return "KE"
	case 296:
		return "KI"
	case 410:
		return "KR"
	case 408:
		return "KP"
	case 414:
		return "KW"
	case 417:
		return "KG"
	case 418:
		return "LA"
	case 428:
		return "LV"
	case 422:
		return "LB"
	case 426:
		return "LS"
	case 430:
		return "LR"
	case 434:
		return "LY"
	case 438:
		return "LI"
	case 440:
		return "LT"
	case 442:
		return "LU"
	case 446:
		return "MO"
	case 807:
		return "MK"
	case 450:
		return "MG"
	case 454:
		return "MW"
	case 458:
		return "MY"
	case 462:
		return "MV"
	case 466:
		return "ML"
	case 470:
		return "MT"
	case 584:
		return "MH"
	case 474:
		return "MQ"
	case 478:
		return "MR"
	case 480:
		return "MU"
	case 175:
		return "YT"
	case 484:
		return "MX"
	case 583:
		return "FM"
	case 498:
		return "MD"
	case 492:
		return "MC"
	case 496:
		return "MN"
	case 500:
		return "MS"
	case 504:
		return "MA"
	case 508:
		return "MZ"
	case 104:
		return "MM"
	case 516:
		return "NA"
	case 520:
		return "NR"
	case 524:
		return "NP"
	case 528:
		return "NL"
	case 530:
		return "AN"
	case 540:
		return "NC"
	case 554:
		return "NZ"
	case 558:
		return "NI"
	case 562:
		return "NE"
	case 566:
		return "NG"
	case 570:
		return "NU"
	case 574:
		return "NF"
	case 580:
		return "MP"
	case 578:
		return "NO"
	case 512:
		return "OM"
	case 586:
		return "PK"
	case 585:
		return "PW"
	case 275:
		return "PS"
	case 591:
		return "PA"
	case 598:
		return "PG"
	case 600:
		return "PY"
	case 604:
		return "PE"
	case 608:
		return "PH"
	case 612:
		return "PN"
	case 616:
		return "PL"
	case 620:
		return "PT"
	case 630:
		return "PR"
	case 634:
		return "QA"
	case 638:
		return "RE"
	case 642:
		return "RO"
	case 643:
		return "RU"
	case 646:
		return "RW"
	case 654:
		return "SH"
	case 659:
		return "KN"
	case 662:
		return "LC"
	case 666:
		return "PM"
	case 670:
		return "VC"
	case 882:
		return "WS"
	case 674:
		return "SM"
	case 678:
		return "ST"
	case 682:
		return "SA"
	case 686:
		return "SN"
	case 690:
		return "SC"
	case 694:
		return "SL"
	case 702:
		return "SG"
	case 703:
		return "SK"
	case 705:
		return "SI"
	case 90:
		return "SB"
	case 706:
		return "SO"
	case 710:
		return "ZA"
	case 239:
		return "GS"
	case 724:
		return "ES"
	case 144:
		return "LK"
	case 736:
		return "SD"
	case 740:
		return "SR"
	case 744:
		return "SJ"
	case 748:
		return "SZ"
	case 752:
		return "SE"
	case 756:
		return "CH"
	case 760:
		return "SY"
	case 158:
		return "TW"
	case 762:
		return "TJ"
	case 834:
		return "TZ"
	case 764:
		return "TH"
	case 626:
		return "TL"
	case 768:
		return "TG"
	case 772:
		return "TK"
	case 776:
		return "TO"
	case 780:
		return "TT"
	case 788:
		return "TN"
	case 792:
		return "TR"
	case 795:
		return "TM"
	case 796:
		return "TC"
	case 798:
		return "TV"
	case 800:
		return "UG"
	case 804:
		return "UA"
	case 784:
		return "AE"
	case 826:
		return "GB"
	case 840:
		return "US"
	case 581:
		return "UM"
	case 858:
		return "UY"
	case 860:
		return "UZ"
	case 548:
		return "VU"
	case 336:
		return "VA"
	case 862:
		return "VE"
	case 704:
		return "VN"
	case 92:
		return "VG"
	case 850:
		return "VI"
	case 876:
		return "WF"
	case 732:
		return "EH"
	case 887:
		return "YE"
	case 891:
		return "YU"
	case 894:
		return "ZM"
	case 716:
		return "ZW"
	case 4:
		return "AF"
	case 688:
		return "RS"
	case 248:
		return "AX"
	case 535:
		return "BQ"
	case 831:
		return "GG"
	case 832:
		return "JE"
	case 531:
		return "CW"
	case 833:
		return "IM"
	case 652:
		return "BL"
	case 663:
		return "MF"
	case 534:
		return "SX"
	case 499:
		return "ME"
	case 728:
		return "SS"
	}
	return UnknownMsg
}

// Alpha3 - returns a Alpha-3 (ISO3, 3 chars) code of country
func (c CountryCode) Alpha3() string {
	switch c {
	case 8:
		return "ALB"
	case 12:
		return "DZA"
	case 16:
		return "ASM"
	case 20:
		return "AND"
	case 24:
		return "AGO"
	case 660:
		return "AIA"
	case 10:
		return "ATA"
	case 28:
		return "ATG"
	case 32:
		return "ARG"
	case 51:
		return "ARM"
	case 533:
		return "ABW"
	case 36:
		return "AUS"
	case 40:
		return "AUT"
	case 31:
		return "AZE"
	case 44:
		return "BHS"
	case 48:
		return "BHR"
	case 50:
		return "BGD"
	case 52:
		return "BRB"
	case 112:
		return "BLR"
	case 56:
		return "BEL"
	case 84:
		return "BLZ"
	case 204:
		return "BEN"
	case 60:
		return "BMU"
	case 64:
		return "BTN"
	case 68:
		return "BOL"
	case 70:
		return "BIH"
	case 72:
		return "BWA"
	case 74:
		return "BVT"
	case 76:
		return "BRA"
	case 86:
		return "IOT"
	case 96:
		return "BRN"
	case 100:
		return "BGR"
	case 854:
		return "BFA"
	case 108:
		return "BDI"
	case 116:
		return "KHM"
	case 120:
		return "CMR"
	case 124:
		return "CAN"
	case 132:
		return "CPV"
	case 136:
		return "CYM"
	case 140:
		return "CAF"
	case 148:
		return "TCD"
	case 152:
		return "CHL"
	case 156:
		return "CHN"
	case 162:
		return "CXR"
	case 166:
		return "CCK"
	case 170:
		return "COL"
	case 174:
		return "COM"
	case 178:
		return "COG"
	case 180:
		return "COD"
	case 184:
		return "COK"
	case 188:
		return "CRI"
	case 384:
		return "CIV"
	case 191:
		return "HRV"
	case 192:
		return "CUB"
	case 196:
		return "CYP"
	case 203:
		return "CZE"
	case 208:
		return "DNK"
	case 262:
		return "DJI"
	case 212:
		return "DMA"
	case 214:
		return "DOM"
	case 218:
		return "ECU"
	case 818:
		return "EGY"
	case 222:
		return "SLV"
	case 226:
		return "GNQ"
	case 232:
		return "ERI"
	case 233:
		return "EST"
	case 231:
		return "ETH"
	case 238:
		return "FLK"
	case 242:
		return "FJI"
	case 246:
		return "FIN"
	case 250:
		return "FRA"
	case 254:
		return "GUF"
	case 258:
		return "PYF"
	case 260:
		return "ATF"
	case 266:
		return "GAB"
	case 270:
		return "GMB"
	case 268:
		return "GEO"
	case 276:
		return "DEU"
	case 288:
		return "GHA"
	case 292:
		return "GIB"
	case 300:
		return "GRC"
	case 304:
		return "GRL"
	case 308:
		return "GRD"
	case 312:
		return "GLP"
	case 316:
		return "GUM"
	case 320:
		return "GTM"
	case 324:
		return "GIN"
	case 624:
		return "GNB"
	case 328:
		return "GUY"
	case 332:
		return "HTI"
	case 334:
		return "HMD"
	case 340:
		return "HND"
	case 344:
		return "HKG"
	case 348:
		return "HUN"
	case 352:
		return "ISL"
	case 356:
		return "IND"
	case 360:
		return "IDN"
	case 364:
		return "IRN"
	case 368:
		return "IRQ"
	case 372:
		return "IRL"
	case 376:
		return "ISR"
	case 380:
		return "ITA"
	case 388:
		return "JAM"
	case 392:
		return "JPN"
	case 400:
		return "JOR"
	case 398:
		return "KAZ"
	case 404:
		return "KEN"
	case 296:
		return "KIR"
	case 410:
		return "KOR"
	case 408:
		return "PRK"
	case 414:
		return "KWT"
	case 417:
		return "KGZ"
	case 418:
		return "LAO"
	case 428:
		return "LVA"
	case 422:
		return "LBN"
	case 426:
		return "LSO"
	case 430:
		return "LBR"
	case 434:
		return "LBY"
	case 438:
		return "LIE"
	case 440:
		return "LTU"
	case 442:
		return "LUX"
	case 446:
		return "MAC"
	case 807:
		return "MKD"
	case 450:
		return "MDG"
	case 454:
		return "MWI"
	case 458:
		return "MYS"
	case 462:
		return "MDV"
	case 466:
		return "MLI"
	case 470:
		return "MLT"
	case 584:
		return "MHL"
	case 474:
		return "MTQ"
	case 478:
		return "MRT"
	case 480:
		return "MUS"
	case 175:
		return "MYT"
	case 484:
		return "MEX"
	case 583:
		return "FSM"
	case 498:
		return "MDA"
	case 492:
		return "MCO"
	case 496:
		return "MNG"
	case 500:
		return "MSR"
	case 504:
		return "MAR"
	case 508:
		return "MOZ"
	case 104:
		return "MMR"
	case 516:
		return "NAM"
	case 520:
		return "NRU"
	case 524:
		return "NPL"
	case 528:
		return "NLD"
	case 530:
		return "ANT"
	case 540:
		return "NCL"
	case 554:
		return "NZL"
	case 558:
		return "NIC"
	case 562:
		return "NER"
	case 566:
		return "NGA"
	case 570:
		return "NIU"
	case 574:
		return "NFK"
	case 580:
		return "MNP"
	case 578:
		return "NOR"
	case 512:
		return "OMN"
	case 586:
		return "PAK"
	case 585:
		return "PLW"
	case 275:
		return "PSE"
	case 591:
		return "PAN"
	case 598:
		return "PNG"
	case 600:
		return "PRY"
	case 604:
		return "PER"
	case 608:
		return "PHL"
	case 612:
		return "PCN"
	case 616:
		return "POL"
	case 620:
		return "PRT"
	case 630:
		return "PRI"
	case 634:
		return "QAT"
	case 638:
		return "REU"
	case 642:
		return "ROU"
	case 643:
		return "RUS"
	case 646:
		return "RWA"
	case 654:
		return "SHN"
	case 659:
		return "KNA"
	case 662:
		return "LCA"
	case 666:
		return "SPM"
	case 670:
		return "VCT"
	case 882:
		return "WSM"
	case 674:
		return "SMR"
	case 678:
		return "STP"
	case 682:
		return "SAU"
	case 686:
		return "SEN"
	case 690:
		return "SYC"
	case 694:
		return "SLE"
	case 702:
		return "SGP"
	case 703:
		return "SVK"
	case 705:
		return "SVN"
	case 90:
		return "SLB"
	case 706:
		return "SOM"
	case 710:
		return "ZAF"
	case 239:
		return "SGS"
	case 724:
		return "ESP"
	case 144:
		return "LKA"
	case 736:
		return "SDN"
	case 740:
		return "SUR"
	case 744:
		return "SJM"
	case 748:
		return "SWZ"
	case 752:
		return "SWE"
	case 756:
		return "CHE"
	case 760:
		return "SYR"
	case 158:
		return "TWN"
	case 762:
		return "TJK"
	case 834:
		return "TZA"
	case 764:
		return "THA"
	case 626:
		return "TLS"
	case 768:
		return "TGO"
	case 772:
		return "TKL"
	case 776:
		return "TON"
	case 780:
		return "TTO"
	case 788:
		return "TUN"
	case 792:
		return "TUR"
	case 795:
		return "TKM"
	case 796:
		return "TCA"
	case 798:
		return "TUV"
	case 800:
		return "UGA"
	case 804:
		return "UKR"
	case 784:
		return "ARE"
	case 826:
		return "GBR"
	case 840:
		return "USA"
	case 581:
		return "UMI"
	case 858:
		return "URY"
	case 860:
		return "UZB"
	case 548:
		return "VUT"
	case 336:
		return "VAT"
	case 862:
		return "VEN"
	case 704:
		return "VNM"
	case 92:
		return "VGB"
	case 850:
		return "VIR"
	case 876:
		return "WLF"
	case 732:
		return "ESH"
	case 887:
		return "YEM"
	case 891:
		return "YUG"
	case 894:
		return "ZMB"
	case 716:
		return "ZWE"
	case 4:
		return "AFG"
	case 688:
		return "SRB"
	case 248:
		return "ALA"
	case 535:
		return "BES"
	case 831:
		return "GGY"
	case 832:
		return "JEY"
	case 531:
		return "CUW"
	case 833:
		return "IMN"
	case 652:
		return "BLM"
	case 663:
		return "MAF"
	case 534:
		return "SXM"
	case 499:
		return "MNE"
	case 728:
		return "SSD"
	}
	return UnknownMsg
}

// Currency - returns a currency of the country
func (c CountryCode) Currency() CurrencyCode {
	switch c {
	case AUS:
		return CurrencyAUD
	case AUT, AND, MAF:
		return CurrencyEUR
	case AZE:
		return CurrencyAZN
	case ALB:
		return CurrencyALL
	case DZA:
		return CurrencyDZD
	case ASM, BES:
		return CurrencyUSD
	case AIA:
		return CurrencyXCD
	case AGO:
		return CurrencyAOA
	case ATG:
		return CurrencyXCD
	case ANT, CUW:
		return CurrencyANG
	case ARE:
		return CurrencyAED
	case ARG:
		return CurrencyARS
	case ARM:
		return CurrencyAMD
	case ABW:
		return CurrencyAWG
	case AFG:
		return CurrencyAFN
	case BHS:
		return CurrencyBSD
	case BGD:
		return CurrencyBDT
	case BRB:
		return CurrencyBBD
	case BHR:
		return CurrencyBHD
	case BLR:
		return CurrencyBYR
	case BLZ:
		return CurrencyBZD
	case BEL:
		return CurrencyEUR
	case BEN:
		return CurrencyXOF
	case BMU:
		return CurrencyBMD
	case BGR:
		return CurrencyBGN
	case BOL:
		return CurrencyBOB
	case BIH:
		return CurrencyBAM
	case BWA:
		return CurrencyBWP
	case BRA:
		return CurrencyBRL
	case IOT:
		return CurrencyUSD
	case BRN:
		return CurrencyBND
	case BFA:
		return CurrencyXOF
	case BDI:
		return CurrencyBIF
	case BTN:
		return CurrencyBTN
	case VUT:
		return CurrencyVUV
	case VAT:
		return CurrencyEUR
	case GBR, GGY, JEY, IMN:
		return CurrencyGBP
	case HUN:
		return CurrencyHUF
	case VEN:
		return CurrencyVEF
	case VGB:
		return CurrencyUSD
	case VIR:
		return CurrencyUSD
	case TLS:
		return CurrencyUSD
	case VNM:
		return CurrencyVND
	case GAB:
		return CurrencyXAF
	case HTI:
		return CurrencyHTG
	case GUY:
		return CurrencyGYD
	case GMB:
		return CurrencyGMD
	case GHA:
		return CurrencyGHS
	case GLP:
		return CurrencyEUR
	case GTM:
		return CurrencyGTQ
	case GIN:
		return CurrencyGNF
	case GNB:
		return CurrencyXOF
	case DEU:
		return CurrencyEUR
	case GIB:
		return CurrencyGIP
	case HND:
		return CurrencyHNL
	case HKG:
		return CurrencyHKD
	case GRD:
		return CurrencyXCD
	case GRL:
		return CurrencyDKK
	case GRC:
		return CurrencyEUR
	case GEO:
		return CurrencyGEL
	case GUM:
		return CurrencyUSD
	case DNK:
		return CurrencyDKK
	case COD:
		return CurrencyCDF
	case DJI:
		return CurrencyDJF
	case DMA:
		return CurrencyXCD
	case DOM:
		return CurrencyDOP
	case EGY:
		return CurrencyEGP
	case ZMB:
		return CurrencyZMW
	case ESH:
		return CurrencyMAD
	case ZWE:
		return CurrencyZWL
	case ISR:
		return CurrencyILS
	case IND:
		return CurrencyINR
	case IDN:
		return CurrencyIDR
	case JOR:
		return CurrencyJOD
	case IRQ:
		return CurrencyIQD
	case IRN:
		return CurrencyIRR
	case IRL:
		return CurrencyEUR
	case ISL:
		return CurrencyISK
	case ESP:
		return CurrencyEUR
	case ITA:
		return CurrencyEUR
	case YEM:
		return CurrencyYER
	case KAZ:
		return CurrencyKZT
	case CYM:
		return CurrencyKYD
	case KHM:
		return CurrencyKHR
	case CMR:
		return CurrencyXAF
	case CAN:
		return CurrencyCAD
	case QAT:
		return CurrencyQAR
	case KEN:
		return CurrencyKES
	case CYP:
		return CurrencyEUR
	case KIR:
		return CurrencyAUD
	case CHN:
		return CurrencyCNY
	case CCK:
		return CurrencyAUD
	case COL:
		return CurrencyCOP
	case COM:
		return CurrencyKMF
	case COG:
		return CurrencyXAF
	case PRK:
		return CurrencyKPW
	case KOR:
		return CurrencyKRW
	case CRI:
		return CurrencyCRC
	case CIV:
		return CurrencyXOF
	case CUB:
		return CurrencyCUC
	case KWT:
		return CurrencyKWD
	case KGZ:
		return CurrencyKGS
	case LAO:
		return CurrencyLAK
	case LVA:
		return CurrencyEUR
	case LSO:
		return CurrencyLSL
	case LBR:
		return CurrencyLRD
	case LBN:
		return CurrencyLBP
	case LBY:
		return CurrencyLYD
	case LTU:
		return CurrencyEUR
	case LIE:
		return CurrencyCHF
	case LUX:
		return CurrencyEUR
	case MUS:
		return CurrencyMUR
	case MRT:
		return CurrencyMRU
	case MDG:
		return CurrencyMGA
	case MYT:
		return CurrencyEUR
	case MAC:
		return CurrencyMOP
	case MKD:
		return CurrencyMKD
	case MWI:
		return CurrencyMWK
	case MYS:
		return CurrencyMYR
	case MLI:
		return CurrencyXOF
	case MDV:
		return CurrencyMVR
	case MLT:
		return CurrencyEUR
	case MNP:
		return CurrencyUSD
	case MAR:
		return CurrencyMAD
	case MTQ:
		return CurrencyEUR
	case MHL:
		return CurrencyUSD
	case MEX:
		return CurrencyMXN
	case FSM:
		return CurrencyUSD
	case MOZ:
		return CurrencyMZN
	case MDA:
		return CurrencyMDL
	case MCO:
		return CurrencyEUR
	case MNG:
		return CurrencyMNT
	case MSR:
		return CurrencyXCD
	case MMR:
		return CurrencyMMK
	case NAM:
		return CurrencyNAD
	case NRU:
		return CurrencyAUD
	case NPL:
		return CurrencyNPR
	case NER:
		return CurrencyXOF
	case NGA:
		return CurrencyNGN
	case NLD:
		return CurrencyEUR
	case NIC:
		return CurrencyNIO
	case NIU:
		return CurrencyNZD
	case NZL:
		return CurrencyNZD
	case NCL:
		return CurrencyXPF
	case NOR:
		return CurrencyNOK
	case ChannelIslands:
		return CurrencyEUR
	case OMN:
		return CurrencyOMR
	case BVT:
		return CurrencyNOK
	case NFK:
		return CurrencyAUD
	case PCN:
		return CurrencyNZD
	case CXR:
		return CurrencyAUD
	case SHN:
		return CurrencySHP
	case WLF:
		return CurrencyXPF
	case HMD:
		return CurrencyAUD
	case CPV:
		return CurrencyCVE
	case COK:
		return CurrencyNZD
	case WSM:
		return CurrencyWST
	case SJM:
		return CurrencyNOK
	case TCA:
		return CurrencyUSD
	case UMI:
		return CurrencyUSD
	case PAK:
		return CurrencyPKR
	case PLW:
		return CurrencyUSD
	case PSE:
		return CurrencyILS
	case PAN:
		return CurrencyPAB
	case PNG:
		return CurrencyPGK
	case PRY:
		return CurrencyPYG
	case PER:
		return CurrencyPEN
	case POL:
		return CurrencyPLN
	case PRT:
		return CurrencyEUR
	case PRI:
		return CurrencyUSD
	case REU:
		return CurrencyEUR
	case RUS:
		return CurrencyRUB
	case RWA:
		return CurrencyRWF
	case ROU:
		return CurrencyRON
	case SLV:
		return CurrencySVC
	case SMR:
		return CurrencyEUR
	case STP:
		return CurrencySTN
	case SAU:
		return CurrencySAR
	case SWZ:
		return CurrencySZL
	case SYC:
		return CurrencySCR
	case SEN:
		return CurrencyXOF
	case SPM:
		return CurrencyEUR
	case VCT:
		return CurrencyXCD
	case KNA:
		return CurrencyXCD
	case LCA:
		return CurrencyXCD
	case SGP:
		return CurrencySGD
	case SYR:
		return CurrencySYP
	case SVK:
		return CurrencyEUR
	case SVN:
		return CurrencyEUR
	case USA:
		return CurrencyUSN
	case SLB:
		return CurrencySBD
	case SOM:
		return CurrencySOS
	case SDN:
		return CurrencySDG
	case SUR:
		return CurrencySRD
	case SLE:
		return CurrencySLL
	case TJK:
		return CurrencyTJS
	case TWN:
		return CurrencyTWD
	case THA:
		return CurrencyTHB
	case TZA:
		return CurrencyTZS
	case TGO:
		return CurrencyXOF
	case TKL:
		return CurrencyNZD
	case TON:
		return CurrencyTOP
	case TTO:
		return CurrencyTTD
	case TUV:
		return CurrencyAUD
	case TUN:
		return CurrencyTND
	case TKM:
		return CurrencyTMT
	case TUR:
		return CurrencyTRY
	case UGA:
		return CurrencyUGX
	case UZB:
		return CurrencyUZS
	case UKR:
		return CurrencyUAH
	case URY:
		return CurrencyUYI
	case FRO:
		return CurrencyDKK
	case FJI:
		return CurrencyFJD
	case PHL:
		return CurrencyPHP
	case FIN:
		return CurrencyEUR
	case FLK:
		return CurrencyFKP
	case FRA:
		return CurrencyEUR
	case GUF:
		return CurrencyEUR
	case PYF:
		return CurrencyXPF
	case ATF:
		return CurrencyEUR
	case HRV:
		return CurrencyHRK
	case CAF:
		return CurrencyXAF
	case TCD:
		return CurrencyXAF
	case CZE:
		return CurrencyCZK
	case CHL:
		return CurrencyCLF
	case CHE:
		return CurrencyCHE
	case SWE:
		return CurrencySEK
	case LKA:
		return CurrencyLKR
	case ECU:
		return CurrencyUSD
	case GNQ:
		return CurrencyXAF
	case ERI:
		return CurrencyERN
	case EST:
		return CurrencyEUR
	case ETH:
		return CurrencyETB
	case ZAF:
		return CurrencyZAR
	case YUG:
		return CurrencyYUD
	case SGS:
		return CurrencyGBP
	case JAM:
		return CurrencyJMD
	case JPN:
		return CurrencyJPY
	case BLM, MNE, ALA:
		return CurrencyEUR
	case SXM:
		return CurrencyANG
	case SRB:
		return CurrencyRSD
	case SSD:
		return CurrencySSP
	}
	return CurrencyUnknown
}

// All - return all country codes
func All() []CountryCode {
	return []CountryCode{
		AUS,
		AUT,
		AZE,
		ALB,
		DZA,
		ASM,
		AIA,
		AGO,
		AND,
		ATA,
		ATG,
		ANT,
		ARE,
		ARG,
		ARM,
		ABW,
		AFG,
		BHS,
		BGD,
		BRB,
		BHR,
		BLR,
		BLZ,
		BEL,
		BEN,
		BMU,
		BGR,
		BOL,
		BIH,
		BWA,
		BRA,
		IOT,
		BRN,
		BFA,
		BDI,
		BTN,
		VUT,
		VAT,
		GBR,
		HUN,
		VEN,
		VGB,
		VIR,
		TLS,
		VNM,
		GAB,
		HTI,
		GUY,
		GMB,
		GHA,
		GLP,
		GTM,
		GIN,
		GNB,
		DEU,
		GIB,
		HND,
		HKG,
		GRD,
		GRL,
		GRC,
		GEO,
		GUM,
		DNK,
		COD,
		DJI,
		DMA,
		DOM,
		EGY,
		ZMB,
		ESH,
		ZWE,
		ISR,
		IND,
		IDN,
		JOR,
		IRQ,
		IRN,
		IRL,
		ISL,
		ESP,
		ITA,
		YEM,
		KAZ,
		CYM,
		KHM,
		CMR,
		CAN,
		QAT,
		KEN,
		CYP,
		KIR,
		CHN,
		CCK,
		COL,
		COM,
		COG,
		PRK,
		KOR,
		CRI,
		CIV,
		CUB,
		KWT,
		KGZ,
		LAO,
		LVA,
		LSO,
		LBR,
		LBN,
		LBY,
		LTU,
		LIE,
		LUX,
		MUS,
		MRT,
		MDG,
		MYT,
		MAC,
		MKD,
		MWI,
		MYS,
		MLI,
		MDV,
		MLT,
		MNP,
		MAR,
		MTQ,
		MHL,
		MEX,
		FSM,
		MOZ,
		MDA,
		MCO,
		MNG,
		MSR,
		MMR,
		NAM,
		NRU,
		NPL,
		NER,
		NGA,
		NLD,
		NIC,
		NIU,
		NZL,
		NCL,
		NOR,
		OMN,
		BVT,
		IMN,
		NFK,
		PCN,
		CXR,
		SHN,
		WLF,
		HMD,
		CPV,
		COK,
		WSM,
		SJM,
		TCA,
		UMI,
		PAK,
		PLW,
		PSE,
		PAN,
		PNG,
		PRY,
		PER,
		POL,
		PRT,
		PRI,
		REU,
		RUS,
		RWA,
		ROU,
		SLV,
		SMR,
		STP,
		SAU,
		SWZ,
		SYC,
		SEN,
		SPM,
		VCT,
		KNA,
		LCA,
		SGP,
		SYR,
		SVK,
		SVN,
		USA,
		SLB,
		SOM,
		SDN,
		SUR,
		SLE,
		TJK,
		TWN,
		THA,
		TZA,
		TGO,
		TKL,
		TON,
		TTO,
		TUV,
		TUN,
		TKM,
		TUR,
		UGA,
		UZB,
		UKR,
		URY,
		XWA,
		FRO,
		FJI,
		PHL,
		FIN,
		FLK,
		FRA,
		GUF,
		PYF,
		ATF,
		HRV,
		CAF,
		TCD,
		CZE,
		CHL,
		CHE,
		SWE,
		XSC,
		LKA,
		ECU,
		GNQ,
		ERI,
		EST,
		ETH,
		ZAF,
		YUG,
		SGS,
		JAM,
		MNE,
		BLM,
		SXM,
		SRB,
		ALA,
		BES,
		GGY,
		JEY,
		CUW,
		MAF,
		SSD,
		JPN,
	}
}

// CallCode - return calling code of country
func (c CountryCode) CallCode() CallCode {
	switch c {
	case AUS:
		return CallCode(61)
	case AUT:
		return CallCode(43)
	case AZE:
		return CallCode(994)
	case ALB:
		return CallCode(355)
	case DZA:
		return CallCode(213)
	case ASM:
		return CallCode(1684)
	case AIA:
		return CallCode(1264)
	case AGO:
		return CallCode(244)
	case AND:
		return CallCode(376)
	case ATA:
		return CallCode(672)
	case ATG:
		return CallCode(1268)
	case ANT:
		return CallCode(599)
	case ARE:
		return CallCode(971)
	case ARG:
		return CallCode(54)
	case ARM:
		return CallCode(374)
	case ABW:
		return CallCode(297)
	case AFG:
		return CallCode(93)
	case BHS:
		return CallCode(1242)
	case BGD:
		return CallCode(880)
	case BRB:
		return CallCode(1246)
	case BHR:
		return CallCode(973)
	case BLR:
		return CallCode(375)
	case BLZ:
		return CallCode(501)
	case BEL:
		return CallCode(32)
	case BEN:
		return CallCode(229)
	case BMU:
		return CallCode(1441)
	case BGR:
		return CallCode(359)
	case BOL:
		return CallCode(591)
	case BIH:
		return CallCode(387)
	case BWA:
		return CallCode(267)
	case BRA:
		return CallCode(55)
	case IOT:
		return CallCode(246)
	case BRN:
		return CallCode(673)
	case BFA:
		return CallCode(226)
	case BDI:
		return CallCode(257)
	case BTN:
		return CallCode(975)
	case VUT:
		return CallCode(678)
	case VAT:
		return CallCode(39)
	case GBR:
		return CallCode(44)
	case HUN:
		return CallCode(36)
	case VEN:
		return CallCode(58)
	case VGB:
		return CallCode(1284)
	case VIR:
		return CallCode(1340)
	case TLS:
		return CallCode(670)
	case VNM:
		return CallCode(84)
	case GAB:
		return CallCode(241)
	case HTI:
		return CallCode(509)
	case GUY:
		return CallCode(592)
	case GMB:
		return CallCode(220)
	case GHA:
		return CallCode(233)
	case GLP:
		return CallCode(590)
	case GTM:
		return CallCode(502)
	case GIN:
		return CallCode(224)
	case GNB:
		return CallCode(245)
	case DEU:
		return CallCode(49)
	case GIB:
		return CallCode(350)
	case HND:
		return CallCode(504)
	case HKG:
		return CallCode(852)
	case GRD:
		return CallCode(1473)
	case GRL:
		return CallCode(299)
	case GRC:
		return CallCode(30)
	case GEO:
		return CallCode(995)
	case GUM:
		return CallCode(1671)
	case DNK:
		return CallCode(45)
	case COD:
		return CallCode(243)
	case DJI:
		return CallCode(253)
	case DMA:
		return CallCode(1767)
	case DOM:
		return CallCode(1809)
	case EGY:
		return CallCode(20)
	case ZMB:
		return CallCode(260)
	case ESH:
		return CallCode(212)
	case ZWE:
		return CallCode(263)
	case ISR:
		return CallCode(972)
	case IND:
		return CallCode(91)
	case IDN:
		return CallCode(62)
	case JOR:
		return CallCode(962)
	case IRQ:
		return CallCode(964)
	case IRN:
		return CallCode(98)
	case IRL:
		return CallCode(353)
	case ISL:
		return CallCode(354)
	case ESP:
		return CallCode(34)
	case ITA:
		return CallCode(39)
	case YEM:
		return CallCode(967)
	case KAZ:
		return CallCode(7)
	case CYM:
		return CallCode(1345)
	case KHM:
		return CallCode(855)
	case CMR:
		return CallCode(237)
	case CAN:
		return CallCode(1)
	case QAT:
		return CallCode(974)
	case KEN:
		return CallCode(254)
	case CYP:
		return CallCode(357)
	case KIR:
		return CallCode(686)
	case CHN:
		return CallCode(86)
	case CCK:
		return CallCode(672)
	case COL:
		return CallCode(57)
	case COM:
		return CallCode(269)
	case COG:
		return CallCode(242)
	case PRK:
		return CallCode(850)
	case KOR:
		return CallCode(82)
	case CRI:
		return CallCode(506)
	case CIV:
		return CallCode(225)
	case CUB:
		return CallCode(53)
	case KWT:
		return CallCode(965)
	case KGZ:
		return CallCode(996)
	case LAO:
		return CallCode(856)
	case LVA:
		return CallCode(371)
	case LSO:
		return CallCode(266)
	case LBR:
		return CallCode(231)
	case LBN:
		return CallCode(961)
	case LBY:
		return CallCode(218)
	case LTU:
		return CallCode(370)
	case LIE:
		return CallCode(423)
	case LUX:
		return CallCode(352)
	case MUS:
		return CallCode(230)
	case MRT:
		return CallCode(222)
	case MDG:
		return CallCode(261)
	case MYT:
		return CallCode(262)
	case MAC:
		return CallCode(853)
	case MKD:
		return CallCode(389)
	case MWI:
		return CallCode(265)
	case MYS:
		return CallCode(60)
	case MLI:
		return CallCode(223)
	case MDV:
		return CallCode(960)
	case MLT:
		return CallCode(356)
	case MNP:
		return CallCode(1670)
	case MAR:
		return CallCode(212)
	case MTQ:
		return CallCode(596)
	case MHL:
		return CallCode(692)
	case MEX:
		return CallCode(52)
	case FSM:
		return CallCode(691)
	case MOZ:
		return CallCode(258)
	case MDA:
		return CallCode(373)
	case MCO:
		return CallCode(377)
	case MNG:
		return CallCode(976)
	case MSR:
		return CallCode(1664)
	case MMR:
		return CallCode(95)
	case NAM:
		return CallCode(264)
	case NRU:
		return CallCode(674)
	case NPL:
		return CallCode(977)
	case NER:
		return CallCode(227)
	case NGA:
		return CallCode(234)
	case NLD:
		return CallCode(31)
	case NIC:
		return CallCode(505)
	case NIU:
		return CallCode(683)
	case NZL:
		return CallCode(64)
	case NCL:
		return CallCode(687)
	case NOR:
		return CallCode(47)
	case OMN:
		return CallCode(968)
	case BVT:
		return CallCode(47)
	case IMN:
		return CallCode(44)
	case NFK:
		return CallCode(672)
	case PCN:
		return CallCode(870)
	case CXR:
		return CallCode(61)
	case SHN:
		return CallCode(290)
	case WLF:
		return CallCode(681)
	case HMD:
		return CallCode(0)
	case CPV:
		return CallCode(238)
	case COK:
		return CallCode(682)
	case WSM:
		return CallCode(684)
	case SJM:
		return CallCode(47)
	case TCA:
		return CallCode(1649)
	case UMI:
		return CallCode(1)
	case PAK:
		return CallCode(92)
	case PLW:
		return CallCode(680)
	case PSE:
		return CallCode(972)
	case PAN:
		return CallCode(507)
	case PNG:
		return CallCode(675)
	case PRY:
		return CallCode(595)
	case PER:
		return CallCode(51)
	case POL:
		return CallCode(48)
	case PRT:
		return CallCode(351)
	case PRI:
		return CallCode(1787)
	case REU:
		return CallCode(262)
	case RUS:
		return CallCode(7)
	case RWA:
		return CallCode(250)
	case ROU:
		return CallCode(40)
	case SLV:
		return CallCode(503)
	case SMR:
		return CallCode(378)
	case STP:
		return CallCode(239)
	case SAU:
		return CallCode(966)
	case SWZ:
		return CallCode(268)
	case SYC:
		return CallCode(248)
	case SEN:
		return CallCode(221)
	case SPM:
		return CallCode(508)
	case VCT:
		return CallCode(1784)
	case KNA:
		return CallCode(1869)
	case LCA:
		return CallCode(1758)
	case SGP:
		return CallCode(65)
	case SYR:
		return CallCode(963)
	case SVK:
		return CallCode(421)
	case SVN:
		return CallCode(386)
	case USA:
		return CallCode(1)
	case SLB:
		return CallCode(677)
	case SOM:
		return CallCode(252)
	case SDN:
		return CallCode(249)
	case SUR:
		return CallCode(597)
	case SLE:
		return CallCode(232)
	case TJK:
		return CallCode(992)
	case TWN:
		return CallCode(886)
	case THA:
		return CallCode(66)
	case TZA:
		return CallCode(255)
	case TGO:
		return CallCode(228)
	case TKL:
		return CallCode(690)
	case TON:
		return CallCode(676)
	case TTO:
		return CallCode(1868)
	case TUV:
		return CallCode(688)
	case TUN:
		return CallCode(216)
	case TKM:
		return CallCode(7370)
	case TUR:
		return CallCode(90)
	case UGA:
		return CallCode(256)
	case UZB:
		return CallCode(998)
	case UKR:
		return CallCode(380)
	case URY:
		return CallCode(598)
	case FRO:
		return CallCode(298)
	case FJI:
		return CallCode(679)
	case PHL:
		return CallCode(63)
	case FIN:
		return CallCode(358)
	case FLK:
		return CallCode(500)
	case FRA:
		return CallCode(33)
	case GUF:
		return CallCode(594)
	case PYF:
		return CallCode(689)
	case ATF:
		return CallCode(1)
	case HRV:
		return CallCode(385)
	case CAF:
		return CallCode(236)
	case TCD:
		return CallCode(235)
	case CZE:
		return CallCode(420)
	case CHL:
		return CallCode(56)
	case CHE:
		return CallCode(41)
	case SWE:
		return CallCode(46)
	case LKA:
		return CallCode(94)
	case ECU:
		return CallCode(593)
	case GNQ:
		return CallCode(240)
	case ERI:
		return CallCode(291)
	case EST:
		return CallCode(372)
	case ETH:
		return CallCode(251)
	case ZAF:
		return CallCode(27)
	case YUG:
		return CallCode(38)
	case SGS:
		return CallCode(500)
	case JAM:
		return CallCode(1876)
	case MNE:
		return CallCode(382)
	case BLM:
		return CallCode(590)
	case SXM:
		return CallCode(721)
	case SRB:
		return CallCode(381)
	case ALA:
		return CallCode(358)
	case BES:
		return CallCode(599)
	case GGY:
		return CallCode(44)
	case JEY:
		return CallCode(44)
	case CUW:
		return CallCode(599)
	case MAF:
		return CallCode(590)
	case SSD:
		return CallCode(211)
	case JPN:
		return CallCode(81)
	}
	return 0
}

// Domain - return domain code of country
func (c CountryCode) Domain() DomainCode {
	domain := DomainCode(c)
	if domain.IsValid() {
		return domain
	}
	return DomainUnknown
}

// Region - return Region code ot the country
func (c CountryCode) Region() RegionCode {
	switch c {
	case AUS:
		return RegionOC
	case AUT:
		return RegionEU
	case AZE:
		return RegionAS
	case ALB:
		return RegionEU
	case DZA:
		return RegionAF
	case ASM:
		return RegionOC
	case AIA:
		return RegionNA
	case AGO:
		return RegionAF
	case AND:
		return RegionEU
	case ATA:
		return RegionAN
	case ATG:
		return RegionNA
	case ANT:
		return RegionNA
	case ARE:
		return RegionAS
	case ARG:
		return RegionSA
	case ARM:
		return RegionAS
	case ABW:
		return RegionNA
	case AFG:
		return RegionAS
	case BHS:
		return RegionNA
	case BGD:
		return RegionAS
	case BRB:
		return RegionNA
	case BHR:
		return RegionAS
	case BLR:
		return RegionEU
	case BLZ:
		return RegionNA
	case BEL:
		return RegionEU
	case BEN:
		return RegionAF
	case BMU:
		return RegionNA
	case BGR:
		return RegionEU
	case BOL:
		return RegionSA
	case BIH:
		return RegionEU
	case BWA:
		return RegionAF
	case BRA:
		return RegionSA
	case IOT:
		return RegionAS
	case BRN:
		return RegionAS
	case BFA:
		return RegionAF
	case BDI:
		return RegionAF
	case BTN:
		return RegionAS
	case VUT:
		return RegionOC
	case VAT:
		return RegionEU
	case GBR:
		return RegionEU
	case HUN:
		return RegionEU
	case VEN:
		return RegionSA
	case VGB:
		return RegionNA
	case VIR:
		return RegionNA
	case TLS:
		return RegionAS
	case VNM:
		return RegionAS
	case GAB:
		return RegionAF
	case HTI:
		return RegionNA
	case GUY:
		return RegionSA
	case GMB:
		return RegionAF
	case GHA:
		return RegionAF
	case GLP:
		return RegionNA
	case GTM:
		return RegionNA
	case GIN:
		return RegionAF
	case GNB:
		return RegionAF
	case DEU:
		return RegionEU
	case GIB:
		return RegionEU
	case HND:
		return RegionNA
	case HKG:
		return RegionAS
	case GRD:
		return RegionNA
	case GRL:
		return RegionNA
	case GRC:
		return RegionEU
	case GEO:
		return RegionAS
	case GUM:
		return RegionOC
	case DNK:
		return RegionEU
	case COD:
		return RegionAF
	case DJI:
		return RegionAF
	case DMA:
		return RegionNA
	case DOM:
		return RegionNA
	case EGY:
		return RegionAF
	case ZMB:
		return RegionAF
	case ESH:
		return RegionAF
	case ZWE:
		return RegionAF
	case ISR:
		return RegionAS
	case IND:
		return RegionAS
	case IDN:
		return RegionAS
	case JOR:
		return RegionAS
	case IRQ:
		return RegionAS
	case IRN:
		return RegionAS
	case IRL:
		return RegionEU
	case ISL:
		return RegionEU
	case ESP:
		return RegionEU
	case ITA:
		return RegionEU
	case YEM:
		return RegionAS
	case KAZ:
		return RegionAS
	case CYM:
		return RegionNA
	case KHM:
		return RegionAS
	case CMR:
		return RegionAF
	case CAN:
		return RegionNA
	case QAT:
		return RegionAS
	case KEN:
		return RegionAF
	case CYP:
		return RegionAS
	case KIR:
		return RegionOC
	case CHN:
		return RegionAS
	case CCK:
		return RegionAS
	case COL:
		return RegionSA
	case COM:
		return RegionAF
	case COG:
		return RegionAF
	case PRK:
		return RegionAS
	case KOR:
		return RegionAS
	case CRI:
		return RegionNA
	case CIV:
		return RegionAF
	case CUB:
		return RegionNA
	case KWT:
		return RegionAS
	case KGZ:
		return RegionAS
	case LAO:
		return RegionAS
	case LVA:
		return RegionEU
	case LSO:
		return RegionAF
	case LBR:
		return RegionAF
	case LBN:
		return RegionAS
	case LBY:
		return RegionAF
	case LTU:
		return RegionEU
	case LIE:
		return RegionEU
	case LUX:
		return RegionEU
	case MUS:
		return RegionAF
	case MRT:
		return RegionAF
	case MDG:
		return RegionAF
	case MYT:
		return RegionAF
	case MAC:
		return RegionAS
	case MKD:
		return RegionEU
	case MWI:
		return RegionAF
	case MYS:
		return RegionAS
	case MLI:
		return RegionAF
	case MDV:
		return RegionAS
	case MLT:
		return RegionEU
	case MNP:
		return RegionOC
	case MAR:
		return RegionAF
	case MTQ:
		return RegionNA
	case MHL:
		return RegionOC
	case MEX:
		return RegionNA
	case FSM:
		return RegionOC
	case MOZ:
		return RegionAF
	case MDA:
		return RegionEU
	case MCO:
		return RegionEU
	case MNG:
		return RegionAS
	case MSR:
		return RegionNA
	case MMR:
		return RegionAS
	case NAM:
		return RegionAF
	case NRU:
		return RegionOC
	case NPL:
		return RegionAS
	case NER:
		return RegionAF
	case NGA:
		return RegionAF
	case NLD:
		return RegionEU
	case NIC:
		return RegionNA
	case NIU:
		return RegionOC
	case NZL:
		return RegionOC
	case NCL:
		return RegionOC
	case NOR:
		return RegionEU
	case OMN:
		return RegionAS
	case BVT:
		return RegionAN
	case IMN:
		return RegionEU
	case NFK:
		return RegionOC
	case PCN:
		return RegionOC
	case CXR:
		return RegionAS
	case SHN:
		return RegionAF
	case WLF:
		return RegionOC
	case HMD:
		return RegionAN
	case CPV:
		return RegionAF
	case COK:
		return RegionOC
	case WSM:
		return RegionOC
	case SJM:
		return RegionEU
	case TCA:
		return RegionNA
	case UMI:
		return RegionOC
	case PAK:
		return RegionAS
	case PLW:
		return RegionOC
	case PSE:
		return RegionAS
	case PAN:
		return RegionNA
	case PNG:
		return RegionOC
	case PRY:
		return RegionSA
	case PER:
		return RegionSA
	case POL:
		return RegionEU
	case PRT:
		return RegionEU
	case PRI:
		return RegionNA
	case REU:
		return RegionAF
	case RUS:
		return RegionEU
	case RWA:
		return RegionAF
	case ROU:
		return RegionEU
	case SLV:
		return RegionNA
	case SMR:
		return RegionEU
	case STP:
		return RegionAF
	case SAU:
		return RegionAS
	case SWZ:
		return RegionAF
	case SYC:
		return RegionAF
	case SEN:
		return RegionAF
	case SPM:
		return RegionNA
	case VCT:
		return RegionNA
	case KNA:
		return RegionNA
	case LCA:
		return RegionNA
	case SGP:
		return RegionAS
	case SYR:
		return RegionAS
	case SVK:
		return RegionEU
	case SVN:
		return RegionEU
	case USA:
		return RegionNA
	case SLB:
		return RegionOC
	case SOM:
		return RegionAF
	case SDN:
		return RegionAF
	case SUR:
		return RegionSA
	case SLE:
		return RegionAF
	case TJK:
		return RegionAS
	case TWN:
		return RegionAS
	case THA:
		return RegionAS
	case TZA:
		return RegionAF
	case TGO:
		return RegionAF
	case TKL:
		return RegionOC
	case TON:
		return RegionOC
	case TTO:
		return RegionNA
	case TUV:
		return RegionOC
	case TUN:
		return RegionAF
	case TKM:
		return RegionAS
	case TUR:
		return RegionEU
	case UGA:
		return RegionAF
	case UZB:
		return RegionAS
	case UKR:
		return RegionEU
	case URY:
		return RegionSA
	case FRO:
		return RegionEU
	case FJI:
		return RegionOC
	case PHL:
		return RegionAS
	case FIN:
		return RegionEU
	case FLK:
		return RegionSA
	case FRA:
		return RegionEU
	case GUF:
		return RegionSA
	case PYF:
		return RegionOC
	case ATF:
		return RegionAN
	case HRV:
		return RegionEU
	case CAF:
		return RegionAF
	case TCD:
		return RegionAF
	case CZE:
		return RegionEU
	case CHL:
		return RegionSA
	case CHE:
		return RegionEU
	case SWE:
		return RegionEU
	case LKA:
		return RegionAS
	case ECU:
		return RegionSA
	case GNQ:
		return RegionAF
	case ERI:
		return RegionAF
	case EST:
		return RegionEU
	case ETH:
		return RegionAF
	case ZAF:
		return RegionAF
	case YUG:
		return RegionEU
	case SGS:
		return RegionAN
	case JAM:
		return RegionNA
	case MNE:
		return RegionEU
	case BLM:
		return RegionNA
	case SXM:
		return RegionNA
	case SRB:
		return RegionEU
	case ALA:
		return RegionEU
	case BES:
		return RegionNA
	case GGY:
		return RegionEU
	case JEY:
		return RegionEU
	case CUW:
		return RegionOC
	case MAF:
		return RegionNA
	case SSD:
		return RegionAF
	case JPN:
		return RegionAS
	}
	return RegionUnknown
}

// Capital - return a capital of country
func (c CountryCode) Capital() CapitalCode {
	switch c {
	case AUS:
		return CapitalAU
	case AUT:
		return CapitalAT
	case AZE:
		return CapitalAZ
	case ALB:
		return CapitalAL
	case DZA:
		return CapitalDZ
	case ASM:
		return CapitalAS
	case AIA:
		return CapitalAI
	case AGO:
		return CapitalAO
	case AND:
		return CapitalAD
	case ATA:
		return CapitalAQ
	case ATG:
		return CapitalAG
	case ANT:
		return CapitalAN
	case ARE:
		return CapitalAE
	case ARG:
		return CapitalAR
	case ARM:
		return CapitalAM
	case ABW:
		return CapitalAW
	case AFG:
		return CapitalAF
	case BHS:
		return CapitalBS
	case BGD:
		return CapitalBD
	case BRB:
		return CapitalBB
	case BHR:
		return CapitalBH
	case BLR:
		return CapitalBY
	case BLZ:
		return CapitalBZ
	case BEL:
		return CapitalBE
	case BEN:
		return CapitalBJ
	case BMU:
		return CapitalBM
	case BGR:
		return CapitalBG
	case BOL:
		return CapitalBO
	case BIH:
		return CapitalBA
	case BWA:
		return CapitalBW
	case BRA:
		return CapitalBR
	case IOT:
		return CapitalIO
	case BRN:
		return CapitalBN
	case BFA:
		return CapitalBF
	case BDI:
		return CapitalBI
	case BTN:
		return CapitalBT
	case VUT:
		return CapitalVU
	case VAT:
		return CapitalVA
	case GBR:
		return CapitalGB
	case HUN:
		return CapitalHU
	case VEN:
		return CapitalVE
	case VGB:
		return CapitalVG
	case VIR:
		return CapitalVI
	case TLS:
		return CapitalTL
	case VNM:
		return CapitalVN
	case GAB:
		return CapitalGA
	case HTI:
		return CapitalHT
	case GUY:
		return CapitalGY
	case GMB:
		return CapitalGM
	case GHA:
		return CapitalGH
	case GLP:
		return CapitalGP
	case GTM:
		return CapitalGT
	case GIN:
		return CapitalGN
	case GNB:
		return CapitalGW
	case DEU:
		return CapitalDE
	case GIB:
		return CapitalGI
	case HND:
		return CapitalHN
	case HKG:
		return CapitalHK
	case GRD:
		return CapitalGD
	case GRL:
		return CapitalGL
	case GRC:
		return CapitalGR
	case GEO:
		return CapitalGE
	case GUM:
		return CapitalGU
	case DNK:
		return CapitalDK
	case COD:
		return CapitalCD
	case DJI:
		return CapitalDJ
	case DMA:
		return CapitalDM
	case DOM:
		return CapitalDO
	case EGY:
		return CapitalEG
	case ZMB:
		return CapitalZM
	case ESH:
		return CapitalEH
	case ZWE:
		return CapitalZW
	case ISR:
		return CapitalIL
	case IND:
		return CapitalIN
	case IDN:
		return CapitalID
	case JOR:
		return CapitalJO
	case IRQ:
		return CapitalIQ
	case IRN:
		return CapitalIR
	case IRL:
		return CapitalIE
	case ISL:
		return CapitalIS
	case ESP:
		return CapitalES
	case ITA:
		return CapitalIT
	case YEM:
		return CapitalYE
	case KAZ:
		return CapitalKZ
	case CYM:
		return CapitalKY
	case KHM:
		return CapitalKH
	case CMR:
		return CapitalCM
	case CAN:
		return CapitalCA
	case QAT:
		return CapitalQA
	case KEN:
		return CapitalKE
	case CYP:
		return CapitalCY
	case KIR:
		return CapitalKI
	case CHN:
		return CapitalCN
	case CCK:
		return CapitalCC
	case COL:
		return CapitalCO
	case COM:
		return CapitalKM
	case COG:
		return CapitalCG
	case PRK:
		return CapitalKP
	case KOR:
		return CapitalKR
	case CRI:
		return CapitalCR
	case CIV:
		return CapitalCI
	case CUB:
		return CapitalCU
	case KWT:
		return CapitalKW
	case KGZ:
		return CapitalKG
	case LAO:
		return CapitalLA
	case LVA:
		return CapitalLV
	case LSO:
		return CapitalLS
	case LBR:
		return CapitalLR
	case LBN:
		return CapitalLB
	case LBY:
		return CapitalLY
	case LTU:
		return CapitalLT
	case LIE:
		return CapitalLI
	case LUX:
		return CapitalLU
	case MUS:
		return CapitalMU
	case MRT:
		return CapitalMR
	case MDG:
		return CapitalMG
	case MYT:
		return CapitalYT
	case MAC:
		return CapitalMO
	case MKD:
		return CapitalMK
	case MWI:
		return CapitalMW
	case MYS:
		return CapitalMY
	case MLI:
		return CapitalML
	case MDV:
		return CapitalMV
	case MLT:
		return CapitalMT
	case MNP:
		return CapitalMP
	case MAR:
		return CapitalMA
	case MTQ:
		return CapitalMQ
	case MHL:
		return CapitalMH
	case MEX:
		return CapitalMX
	case FSM:
		return CapitalFM
	case MOZ:
		return CapitalMZ
	case MDA:
		return CapitalMD
	case MCO:
		return CapitalMC
	case MNG:
		return CapitalMN
	case MSR:
		return CapitalMS
	case MMR:
		return CapitalMM
	case NAM:
		return CapitalNA
	case NRU:
		return CapitalNR
	case NPL:
		return CapitalNP
	case NER:
		return CapitalNE
	case NGA:
		return CapitalNG
	case NLD:
		return CapitalNL
	case NIC:
		return CapitalNI
	case NIU:
		return CapitalNU
	case NZL:
		return CapitalNZ
	case NCL:
		return CapitalNC
	case NOR:
		return CapitalNO
	case OMN:
		return CapitalOM
	case BVT:
		return CapitalBV
	case IMN:
		return CapitalIM
	case NFK:
		return CapitalNF
	case PCN:
		return CapitalPN
	case CXR:
		return CapitalCX
	case SHN:
		return CapitalSH
	case WLF:
		return CapitalWF
	case HMD:
		return CapitalHM
	case CPV:
		return CapitalCV
	case COK:
		return CapitalCK
	case WSM:
		return CapitalWS
	case SJM:
		return CapitalSJ
	case TCA:
		return CapitalTC
	case UMI:
		return CapitalUM
	case PAK:
		return CapitalPK
	case PLW:
		return CapitalPW
	case PSE:
		return CapitalPS
	case PAN:
		return CapitalPA
	case PNG:
		return CapitalPG
	case PRY:
		return CapitalPY
	case PER:
		return CapitalPE
	case POL:
		return CapitalPL
	case PRT:
		return CapitalPT
	case PRI:
		return CapitalPR
	case REU:
		return CapitalRE
	case RUS:
		return CapitalRU
	case RWA:
		return CapitalRW
	case ROU:
		return CapitalRO
	case SLV:
		return CapitalSV
	case SMR:
		return CapitalSM
	case STP:
		return CapitalST
	case SAU:
		return CapitalSA
	case SWZ:
		return CapitalSZ
	case SYC:
		return CapitalSC
	case SEN:
		return CapitalSN
	case SPM:
		return CapitalPM
	case VCT:
		return CapitalVC
	case KNA:
		return CapitalKN
	case LCA:
		return CapitalLC
	case SGP:
		return CapitalSG
	case SYR:
		return CapitalSY
	case SVK:
		return CapitalSK
	case SVN:
		return CapitalSI
	case USA:
		return CapitalUS
	case SLB:
		return CapitalSB
	case SOM:
		return CapitalSO
	case SDN:
		return CapitalSD
	case SUR:
		return CapitalSR
	case SLE:
		return CapitalSL
	case TJK:
		return CapitalTJ
	case TWN:
		return CapitalTW
	case THA:
		return CapitalTH
	case TZA:
		return CapitalTZ
	case TGO:
		return CapitalTG
	case TKL:
		return CapitalTK
	case TON:
		return CapitalTO
	case TTO:
		return CapitalTT
	case TUV:
		return CapitalTV
	case TUN:
		return CapitalTN
	case TKM:
		return CapitalTM
	case TUR:
		return CapitalTR
	case UGA:
		return CapitalUG
	case UZB:
		return CapitalUZ
	case UKR:
		return CapitalUA
	case URY:
		return CapitalUY
	case FRO:
		return CapitalFO
	case FJI:
		return CapitalFJ
	case PHL:
		return CapitalPH
	case FIN:
		return CapitalFI
	case FLK:
		return CapitalFK
	case FRA:
		return CapitalFR
	case GUF:
		return CapitalGF
	case PYF:
		return CapitalPF
	case ATF:
		return CapitalTF
	case HRV:
		return CapitalHR
	case CAF:
		return CapitalCF
	case TCD:
		return CapitalTD
	case CZE:
		return CapitalCZ
	case CHL:
		return CapitalCL
	case CHE:
		return CapitalCH
	case SWE:
		return CapitalSE
	case LKA:
		return CapitalLK
	case ECU:
		return CapitalEC
	case GNQ:
		return CapitalGQ
	case ERI:
		return CapitalER
	case EST:
		return CapitalEE
	case ETH:
		return CapitalET
	case ZAF:
		return CapitalZA
	case YUG:
		return CapitalYU
	case SGS:
		return CapitalGS
	case JAM:
		return CapitalJM
	case MNE:
		return CapitalME
	case BLM:
		return CapitalBL
	case SXM:
		return CapitalSX
	case SRB:
		return CapitalRS
	case ALA:
		return CapitalAX
	case BES:
		return CapitalBQ
	case GGY:
		return CapitalGG
	case JEY:
		return CapitalJE
	case CUW:
		return CapitalCW
	case MAF:
		return CapitalMF
	case SSD:
		return CapitalSS
	case JPN:
		return CapitalJP
	}
	return CapitalUnknown
}

// Info - return all info about country as Country struct
func (c CountryCode) Info() *Country {
	return &Country{
		Name:     c.String(),
		Alpha2:   c.Alpha2(),
		Alpha3:   c.Alpha3(),
		Emoji:    c.Emoji(),
		Code:     c,
		Capital:  c.Capital(),
		Currency: c.Currency(),
		CallCode: c.CallCode(),
		Domain:   c.Domain(),
		Region:   c.Region(),
	}
}

// Type implements Typer interface.
func (country *Country) Type() string {
	return TypeCountry
}

// Value implements database/sql/driver.Valuer
func (country Country) Value() (driver.Value, error) {
	return json.Marshal(country)
}

// Scan implements database/sql.Scanner
func (country *Country) Scan(src interface{}) error {
	if country == nil {
		return fmt.Errorf("countries::Scan: Country scan err: country == nil")
	}
	switch src := src.(type) {
	case *Country:
		*country = *src
	case Country:
		*country = src
	case nil:
		country = nil
	default:
		return fmt.Errorf("countries::Scan: Country scan err: unexpected value of type %T for %T", src, *country)
	}
	return nil
}

//
// AllInfo - return all currencies as []Currency
func AllInfo() []*Country {
	all := All()
	countries := make([]*Country, 0, len(all))
	for _, v := range all {
		countries = append(countries, v.Info())
	}
	return countries
}

func textPrepare(text string) string {
	text = strings.Replace(text, " ", "", -1)
	text = strings.Replace(text, "-", "", -1)
	text = strings.Replace(text, "‐", "", -1)
	text = strings.Replace(text, "‑", "", -1)
	text = strings.Replace(text, "‒", "", -1)
	text = strings.Replace(text, "―", "", -1)
	text = strings.Replace(text, "—", "", -1)
	text = strings.Replace(text, "–", "", -1)
	text = strings.Replace(text, "_", "", -1)
	text = strings.Replace(text, ".", "", -1)
	text = strings.Replace(text, ",", "", -1)
	text = strings.Replace(text, ":", "", -1)
	text = strings.Replace(text, ";", "", -1)
	text = strings.Replace(text, "+", "", -1)
	text = strings.Replace(text, "'", "", -1)
	text = strings.Replace(text, "(", "", -1)
	text = strings.Replace(text, ")", "", -1)
	text = strings.Replace(text, "[", "", -1)
	text = strings.Replace(text, "]", "", -1)
	return strings.ToUpper(text)
}

// ByName - return CountryCode by country Alpha-2 / Alpha-3 / name, case-insensitive, example: rus := ByName("Ru") OR rus := ByName("russia")
func ByName(name string) CountryCode {
	switch textPrepare(name) {
	case "AU", "AUS", "AUSTRALIA":
		return AUS
	case "AT", "AUT", "AUSTRIA":
		return AUT
	case "AZ", "AZE", "AZERBAIJAN":
		return AZE
	case "AL", "ALB", "ALBANIA":
		return ALB
	case "DZ", "DZA", "ALGERIA":
		return DZA
	case "AS", "ASM", "AMERICANSAMOA":
		return ASM
	case "AI", "AIA", "ANGUILLA":
		return AIA
	case "XEN", "ENGLAND":
		return GBR
	case "AO", "AGO", "ANGOLA":
		return AGO
	case "AD", "AND", "ANDORRA":
		return AND
	case "AQ", "ATA", "ANTARCTICA":
		return ATA
	case "AG", "ATG", "ANTIGUAANDBARBUDA", "ANTIGUA":
		return ATG
	case "AN", "ANT", "NETHERLANDSANTILLES":
		return ANT
	case "AE", "ARE", "UNITEDARABEMIRATES", "ARABEMIRATES":
		return ARE
	case "AR", "ARG", "ARGENTINA":
		return ARG
	case "AM", "ARM", "ARMENIA":
		return ARM
	case "AW", "ABW", "ARUBA":
		return ABW
	case "AF", "AFG", "AFGHANISTAN":
		return AFG
	case "BS", "BHS", "BAHAMAS":
		return BHS
	case "BD", "BGD", "BANGLADESH":
		return BGD
	case "BB", "BRB", "BARBADOS":
		return BRB
	case "BH", "BHR", "BAHRAIN":
		return BHR
	case "BY", "BLR", "BELARUS":
		return BLR
	case "BZ", "BLZ", "BELIZE":
		return BLZ
	case "BE", "BEL", "BELGIUM":
		return BEL
	case "BJ", "BEN", "BENIN":
		return BEN
	case "BM", "BMU", "BERMUDA":
		return BMU
	case "BG", "BGR", "BULGARIA":
		return BGR
	case "BO", "BOL", "BOLIVIA":
		return BOL
	case "BA", "BIH", "BOSNIAANDHERZEGOVINA", "BOSNIA":
		return BIH
	case "BW", "BWA", "BOTSWANA":
		return BWA
	case "BR", "BRA", "BRAZIL":
		return BRA
	case "IO", "IOT", "BRITISHINDIANOCEANTERRITORY":
		return IOT
	case "BN", "BRN", "BRUNEI":
		return BRN
	case "BF", "BFA", "BURKINAFASO":
		return BFA
	case "BI", "BDI", "BURUNDI":
		return BDI
	case "BT", "BTN", "BHUTAN":
		return BTN
	case "VU", "VUT", "VANUATU":
		return VUT
	case "VA", "VAT", "HOLYSEEVATICAN", "HOLYSEE", "VATICAN":
		return VAT
	case "GB", "GBR", "UNITEDKINGDOM", "UK", "GREATBRITAN", "GREATBRITAIN":
		return GBR
	case "HU", "HUN", "HUNGARY":
		return HUN
	case "VE", "VEN", "VENEZUELA":
		return VEN
	case "VG", "VGB", "VIRGINISLANDSBRITISH", "VIRGINISLANDSGB":
		return VGB
	case "VI", "VIR", "VIRGINISLANDSUS":
		return VIR
	case "TL", "TLS", "TIMORLESTE":
		return TLS
	case "VN", "VNM", "VIETNAM":
		return VNM
	case "GA", "GAB", "GABON":
		return GAB
	case "HT", "HTI", "HAITI":
		return HTI
	case "GY", "GUY", "GUYANA":
		return GUY
	case "GM", "GMB", "GAMBIA":
		return GMB
	case "GH", "GHA", "GHANA":
		return GHA
	case "GP", "GLP", "GUADELOUPE":
		return GLP
	case "GT", "GTM", "GUATEMALA":
		return GTM
	case "GN", "GIN", "GUINEA":
		return GIN
	case "GW", "GNB", "GUINEABISSAU":
		return GNB
	case "DE", "DEU", "GERMANY", "DEUTSCHLAND":
		return DEU
	case "GI", "GIB", "GIBRALTAR", "HIBRALTAR":
		return GIB
	case "HN", "HND", "HONDURAS":
		return HND
	case "HK", "HKG", "HONGKONG":
		return HKG
	case "GD", "GRD", "GRENADA":
		return GRD
	case "GL", "GRL", "GREENLAND":
		return GRL
	case "GR", "GRC", "GREECE":
		return GRC
	case "GE", "GEO", "GEORGIA":
		return GEO
	case "GU", "GUM", "GUAM":
		return GUM
	case "DK", "DNK", "DENMARK":
		return DNK
	case "CD", "COD", "CONGODEMOCRACTICREPUBLIC", "CONGODEMOCRACTIC":
		return COD
	case "DJ", "DJI", "DJIBOUTI":
		return DJI
	case "DM", "DMA", "DOMINICA":
		return DMA
	case "DO", "DOM", "DOMINICANREPUBLIC", "DOMINICANA":
		return DOM
	case "EG", "EGY", "EGYPT":
		return EGY
	case "ZM", "ZMB", "ZAMBIA":
		return ZMB
	case "EH", "ESH", "WESTERNSAHARA":
		return ESH
	case "ZW", "ZWE", "ZIMBABWE":
		return ZWE
	case "IL", "ISR", "ISRAEL":
		return ISR
	case "IN", "IND", "INDIA":
		return IND
	case "ID", "IDN", "INDONESIA":
		return IDN
	case "JO", "JOR", "JORDAN":
		return JOR
	case "IQ", "IRQ", "IRAQ":
		return IRQ
	case "IR", "IRN", "IRAN":
		return IRN
	case "IE", "IRL", "IRELAND":
		return IRL
	case "IS", "ISL", "ICELAND":
		return ISL
	case "ES", "ESP", "SPAIN", "SPANIEN", "ISPANIA":
		return ESP
	case "IT", "ITA", "ITALY":
		return ITA
	case "YE", "YEM", "YEMEN":
		return YEM
	case "KZ", "KAZ", "KAZAKHSTAN", "KAZAHSTAN":
		return KAZ
	case "KY", "CYM", "CAYMANISLANDS", "KAYMANISLANDS":
		return CYM
	case "KH", "KHM", "CAMBODIA":
		return KHM
	case "CM", "CMR", "CAMEROON":
		return CMR
	case "CA", "CAN", "CANADA":
		return CAN
	case "QA", "QAT", "QATAR":
		return QAT
	case "KE", "KEN", "KENYA":
		return KEN
	case "CY", "CYP", "CYPRUS":
		return CYP
	case "KI", "KIR", "KIRIBATI":
		return KIR
	case "CN", "CHN", "CHINA":
		return CHN
	case "CC", "CCK", "KEELING", "COCOSKEELINGISLANDS", "COCOSISLANDS":
		return CCK
	case "CO", "COL", "COLOMBIA":
		return COL
	case "KM", "COM", "COMOROS":
		return COM
	case "CG", "COG", "CONGO":
		return COG
	case "KP", "PRK", "KOREANORTH", "NORTHKOREA":
		return PRK
	case "KR", "KOR", "KOREA":
		return KOR
	case "CR", "CRI", "COSTARICA":
		return CRI
	case "CI", "CIV", "COTEDIVOIRE":
		return CIV
	case "CU", "CUB", "CUBA":
		return CUB
	case "KW", "KWT", "KUWAIT":
		return KWT
	case "KG", "KGZ", "KYRGYZSTAN", "KIRGISISTAN":
		return KGZ
	case "LA", "LAO", "LAOS":
		return LAO
	case "LV", "LVA", "LATVIA", "LATVIYA", "LETTLAND":
		return LVA
	case "LS", "LSO", "LESOTHO":
		return LSO
	case "LR", "LBR", "LIBERIA":
		return LBR
	case "LB", "LBN", "LEBANON":
		return LBN
	case "LY", "LBY", "LIBYA":
		return LBY
	case "LT", "LTU", "LITHUANIA", "LITAUEN", "LITVA":
		return LTU
	case "LI", "LIE", "LIECHTENSTEIN":
		return LIE
	case "LU", "LUX", "LUXEMBOURG":
		return LUX
	case "MU", "MUS", "MAURITIUS":
		return MUS
	case "MR", "MRT", "MAURITANIA":
		return MRT
	case "MG", "MDG", "MADAGASCAR":
		return MDG
	case "YT", "MYT", "MAYOTTE":
		return MYT
	case "MO", "MAC", "MACAUCHINA", "MACAU", "MACAO":
		return MAC
	case "MK", "MKD", "MACEDONIA":
		return MKD
	case "MW", "MWI", "MALAWI":
		return MWI
	case "MY", "MYS", "MALAYSIA":
		return MYS
	case "ML", "MLI", "MALI":
		return MLI
	case "MV", "MDV", "MALDIVES":
		return MDV
	case "MT", "MLT", "MALTA":
		return MLT
	case "MP", "MNP", "NORTHERNMARIANAISLANDS", "MARIANAISLANDS":
		return MNP
	case "MA", "MAR", "MOROCCO", "MOROCO":
		return MAR
	case "MQ", "MTQ", "MARTINIQUE":
		return MTQ
	case "MH", "MHL", "MARSHALLISLANDS", "MARSHALL":
		return MHL
	case "MX", "MEX", "MEXICO":
		return MEX
	case "FM", "FSM", "MICRONESIA":
		return FSM
	case "MZ", "MOZ", "MOZAMBIQUE":
		return MOZ
	case "MD", "MDA", "MOLDOVA":
		return MDA
	case "MC", "MCO", "MONACO":
		return MCO
	case "MN", "MNG", "MONGOLIA":
		return MNG
	case "MS", "MSR", "MONTSERRAT":
		return MSR
	case "MM", "MMR", "MYANMAR":
		return MMR
	case "NA", "NAM", "NAMIBIA":
		return NAM
	case "NR", "NRU", "NAURU":
		return NRU
	case "NP", "NPL", "NEPAL":
		return NPL
	case "NE", "NER", "NIGER":
		return NER
	case "NG", "NGA", "NIGERIA":
		return NGA
	case "NL", "NLD", "NETHERLANDS", "HOLLAND", "HOLLANDIA":
		return NLD
	case "NI", "NIC", "NICARAGUA":
		return NIC
	case "NU", "NIU", "NIUE":
		return NIU
	case "NZ", "NZL", "NEWZEALAND", "NEWZELANDIA", "NEWZELAND":
		return NZL
	case "NC", "NCL", "NEWCALEDONIA":
		return NCL
	case "NO", "NOR", "NORWAY":
		return NOR
	case "OM", "OMN", "OMAN":
		return OMN
	case "BV", "BVT", "BOUVET":
		return BVT
	case "IM", "IMN", "ISLEOFMAN":
		return IMN
	case "NF", "NFK", "NORFOLKISLAND", "NORFOLK":
		return NFK
	case "PN", "PCN", "PITCAIRN":
		return PCN
	case "CX", "CXR", "CHRISTMASISLAND":
		return CXR
	case "SH", "SHN", "SAINTHELENA", "SAINTELENA":
		return SHN
	case "WF", "WLF", "WALLISANDFUTUNAISLANDS":
		return WLF
	case "HM", "HMD", "HEARDISLANDANDMCDONALDISLANDS", "HEARDISLAND":
		return HMD
	case "CV", "CPV", "CAPEVERDE":
		return CPV
	case "CK", "COK", "COOKISLANDS":
		return COK
	case "WS", "WSM", "SAMOA":
		return WSM
	case "SJ", "SJM", "SVALBARDANDJANMAYENISLANDS", "SVALBARD":
		return SJM
	case "TC", "TCA", "TURKSANDCAICOSISLANDS", "CAICOSISLANDS", "CACOSISLANDS":
		return TCA
	case "UM", "UMI", "UNITEDSTATESMINOROUTLYINGISLANDS", "MINOROUTLYINGISLANDS", "MINOROUTLYING", "USMI":
		return UMI
	case "PK", "PAK", "PAKISTAN":
		return PAK
	case "PW", "PLW", "PALAU":
		return PLW
	case "PS", "PSE", "PALESTINE":
		return PSE
	case "PA", "PAN", "PANAMA":
		return PAN
	case "PG", "PNG", "PAPUANEWGUINEA", "PAPUA":
		return PNG
	case "PY", "PRY", "PARAGUAY":
		return PRY
	case "PE", "PER", "PERU":
		return PER
	case "PL", "POL", "POLAND":
		return POL
	case "PT", "PRT", "PORTUGAL":
		return PRT
	case "PR", "PRI", "PUERTORICO":
		return PRI
	case "RE", "REU", "REUNION":
		return REU
	case "RU", "RUS", "RUSSIA", "RUSSO", "RUSSISH", "RUSSLAND", "RUSIA", "ROSSIA", "RUSSIAN", "RUSSIANFEDERATION":
		return RUS
	case "RW", "RWA", "RWANDA", "RUANDA":
		return RWA
	case "RO", "ROU", "ROMANIA", "RUMINIA":
		return ROU
	case "SV", "SLV", "ELSALVADOR":
		return SLV
	case "SM", "SMR", "SANMARINO":
		return SMR
	case "ST", "STP", "SAOTOMEANDPRINCIPE", "SAOTOME":
		return STP
	case "SA", "SAU", "SAUDIARABIA", "SAUDI":
		return SAU
	case "SZ", "SWZ", "SWAZILAND":
		return SWZ
	case "SC", "SYC", "SEYCHELLES":
		return SYC
	case "SN", "SEN", "SENEGAL":
		return SEN
	case "PM", "SPM", "SAINTPIERREANDMIQUELON", "SAINTPIERRE":
		return SPM
	case "VC", "VCT", "SAINTVINCENTANDTHEGRENADINES", "SAINTVINCENT":
		return VCT
	case "KN", "KNA", "SAINTKITTSANDNEVIS", "SAINTKITTS":
		return KNA
	case "LC", "LCA", "SAINTLUCIA", "LUCIA":
		return LCA
	case "SG", "SGP", "SINGAPORE":
		return SGP
	case "SY", "SYR", "SYRIA":
		return SYR
	case "SK", "SVK", "SLOVAKIA":
		return SVK
	case "SI", "SVN", "SLOVENIA":
		return SVN
	case "US", "USA", "UNITEDSTATES", "UNITEDSTATESOFAMERICA":
		return USA
	case "SB", "SLB", "SOLOMONISLANDS", "SOLOMON":
		return SLB
	case "SO", "SOM", "SOMALIA":
		return SOM
	case "SD", "SDN", "SUDAN":
		return SDN
	case "SR", "SUR", "SURINAME":
		return SUR
	case "SL", "SLE", "SIERRALEONE":
		return SLE
	case "TJ", "TJK", "TAJIKISTAN", "TADJIKISTAN":
		return TJK
	case "TW", "TWN", "TAIWAN":
		return TWN
	case "TH", "THA", "THAILAND", "THAI":
		return THA
	case "TZ", "TZA", "TANZANIA":
		return TZA
	case "TG", "TGO", "TOGO":
		return TGO
	case "TK", "TKL", "TOKELAU":
		return TKL
	case "TO", "TON", "TONGA":
		return TON
	case "TT", "TTO", "TRINIDADANDTOBAGO", "TRINIDAD":
		return TTO
	case "TV", "TUV", "TUVALU":
		return TUV
	case "TN", "TUN", "TUNISIA":
		return TUN
	case "TM", "TKM", "TURKMENISTAN", "TURKMENISTON":
		return TKM
	case "TR", "TUR", "TURKEY", "TURCIA":
		return TUR
	case "UG", "UGA", "UGANDA":
		return UGA
	case "UZ", "UZB", "UZBEKISTAN", "UZBEKISTON":
		return UZB
	case "UA", "UKR", "UKRAINE", "UKRAINA":
		return UKR
	case "UY", "URY", "URUGUAY":
		return URY
	case "XW", "XWA", "WALES":
		return XWA
	case "FO", "FRO", "FAROEISLANDS", "FAROE":
		return FRO
	case "FJ", "FJI", "FIJI":
		return FJI
	case "PH", "PHL", "PHILIPPINES", "PHILIPINES":
		return PHL
	case "FI", "FIN", "FINLAND", "FINNISH":
		return FIN
	case "FK", "FLK", "FALKLANDISLANDSMALVINAS", "MALVINAS", "FALKLANDISLANDS", "FALKLAND":
		return FLK
	case "FR", "FRA", "FRANCE":
		return FRA
	case "GF", "GUF", "FRENCHGUIANA", "GUIANA":
		return GUF
	case "PF", "PYF", "FRENCHPOLYNESIA", "POLYNESIA":
		return PYF
	case "TF", "ATF", "FRENCHSOUTHERNTERRITORIES", "SOUTHERNTERRITORIESFRENCH":
		return ATF
	case "HR", "HRV", "CROATIA":
		return HRV
	case "CF", "CAF", "CENTRALAFRICANREPUBLIC", "CENTRALAFRICAN":
		return CAF
	case "TD", "TCD", "CHAD":
		return TCD
	case "CZ", "CZE", "CZECHREPUBLIC", "CZECH":
		return CZE
	case "CL", "CHL", "CHILE":
		return CHL
	case "CH", "CHE", "SWITZERLAND", "SWISS":
		return CHE
	case "SE", "SWE", "SWEDEN":
		return SWE
	case "XS", "XSC", "SCOTLAND":
		return XSC
	case "LK", "LKA", "SRILANKA":
		return LKA
	case "EC", "ECU", "ECUADOR":
		return ECU
	case "GQ", "GNQ", "EQUATORIALGUINEA":
		return GNQ
	case "ER", "ERI", "ERITREA":
		return ERI
	case "EE", "EST", "ESTONIA":
		return EST
	case "ET", "ETH", "ETHIOPIA":
		return ETH
	case "ZA", "ZAF", "SOUTHAFRICA":
		return ZAF
	case "YU", "YUG", "YUGOSLAVIA":
		return YUG
	case "GS", "SGS", "SOUTHGEORGIAANDTHESOUTHSANDWICHISLANDS", "SOUTHGEORGIA":
		return SGS
	case "JM", "JAM", "JAMAICA", "JAMAIKA", "YAMAICA", "YAMAIKA":
		return JAM
	case "ME", "MNE", "MONTENEGRO":
		return MNE
	case "BL", "BLM", "SAINTBARTHELEMY":
		return BLM
	case "SX", "SXM", "SINTMAARTENDUTCH":
		return SXM
	case "RS", "SRB", "SERBIA":
		return SRB
	case "AX", "ALA", "ALANDISLANDS":
		return ALA
	case "BQ", "BES", "BONAIRE":
		return BES
	case "GG", "GGY", "GUERNSEY":
		return GGY
	case "JE", "JEY", "JERSEY":
		return JEY
	case "CW", "CUW", "CURACAO", "CURAKAO", "KURACAO", "KURAKAO":
		return CUW
	case "MF", "MAF", "SAINTMARTINFRENCH":
		return MAF
	case "SS", "SSD", "SOUTHSUDAN":
		return SSD
	case "JP", "JPN", "JAPAN":
		return JPN
	}
	return Unknown
}

// IsValid - returns true, if code is correct
func (c CountryCode) IsValid() bool {
	return c.Alpha2() != UnknownMsg
}
