// Package countries - ISO 3166 (ISO3166-1, ISO3166, Digit, Alpha-2 and Alpha-3) countries codes and names (on eng and rus), ISO 4217 currency designators, ITU-T E.164 IDD calling phone codes, countries capitals, UN M.49 regions codes, ccTLD countries domains, IOC/NOC and FIFA letters codes, VERY FAST, NO maps[], NO slices[], NO external links/files/data, NO interface{}, NO specific dependencies, Databases compatible, Emoji countries flags and currencies support, full support ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and ccTLD standarts. Full support ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and ccTLD standarts.
package countries

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// CurrencyCode - currency code of country
type CurrencyCode int64 // int64 for database/sql/driver.Valuer compatibility

// Currency - all info about currency
type Currency struct {
	NickelRounding bool
	Name           string
	Alpha          string
	Digits         int
	Code           CurrencyCode
	Countries      []CountryCode
}

// TypeCurrencyCode for Typer interface
const TypeCurrencyCode string = "countries.CurrencyCode"

// TypeCurrency for Typer interface
const TypeCurrency string = "countries.Currency"

// Currencies. Two codes present, for example CurrencyUSDollar == CurrencyUSD == 840.
const (
	CurrencyUnknown                        CurrencyCode = 0
	CurrencyAfghani                        CurrencyCode = 971
	CurrencyLek                            CurrencyCode = 8
	CurrencyAlgerianDinar                  CurrencyCode = 12
	CurrencyUSDollar                       CurrencyCode = 840
	CurrencyEuro                           CurrencyCode = 978
	CurrencyKwanza                         CurrencyCode = 973
	CurrencyEastCaribbeanDollar            CurrencyCode = 951
	CurrencyArgentinePeso                  CurrencyCode = 32
	CurrencyArmenianDram                   CurrencyCode = 51
	CurrencyArubanFlorin                   CurrencyCode = 533
	CurrencyAustralianDollar               CurrencyCode = 36
	CurrencyAzerbaijanianManat             CurrencyCode = 944
	CurrencyBahamianDollar                 CurrencyCode = 44
	CurrencyBahrainiDinar                  CurrencyCode = 48
	CurrencyTaka                           CurrencyCode = 50
	CurrencyBarbadosDollar                 CurrencyCode = 52
	CurrencyBelarussianRuble               CurrencyCode = 974
	CurrencyBelizeDollar                   CurrencyCode = 84
	CurrencyCFAFrancBCEAO                  CurrencyCode = 952
	CurrencyBermudianDollar                CurrencyCode = 60
	CurrencyNgultrum                       CurrencyCode = 64
	CurrencyIndianRupee                    CurrencyCode = 356
	CurrencyBoliviano                      CurrencyCode = 68
	CurrencyConvertibleMark                CurrencyCode = 977
	CurrencyPula                           CurrencyCode = 72
	CurrencyNorwegianKrone                 CurrencyCode = 578
	CurrencyBrazilianReal                  CurrencyCode = 986
	CurrencyBruneiDollar                   CurrencyCode = 96
	CurrencyBulgarianLev                   CurrencyCode = 975
	CurrencyBurundiFranc                   CurrencyCode = 108
	CurrencyCaboVerdeEscudo                CurrencyCode = 132
	CurrencyRiel                           CurrencyCode = 116
	CurrencyCFAFrancBEAC                   CurrencyCode = 950
	CurrencyCanadianDollar                 CurrencyCode = 124
	CurrencyCaymanIslandsDollar            CurrencyCode = 136
	CurrencyUnidaddeFomento                CurrencyCode = 990
	CurrencyChileanPeso                    CurrencyCode = 152
	CurrencyYuanRenminbi                   CurrencyCode = 156
	CurrencyColombianPeso                  CurrencyCode = 170
	CurrencyUnidaddeValorReal              CurrencyCode = 970
	CurrencyComoroFranc                    CurrencyCode = 174
	CurrencyCongoleseFranc                 CurrencyCode = 976
	CurrencyNewZealandDollar               CurrencyCode = 554
	CurrencyCostaRicanColon                CurrencyCode = 188
	CurrencyKuna                           CurrencyCode = 191
	CurrencyPesoConvertible                CurrencyCode = 931
	CurrencyCubanPeso                      CurrencyCode = 192
	CurrencyNetherlandsAntilleanGuilder    CurrencyCode = 532
	CurrencyCzechKoruna                    CurrencyCode = 203
	CurrencyDanishKrone                    CurrencyCode = 208
	CurrencyDjiboutiFranc                  CurrencyCode = 262
	CurrencyDominicanPeso                  CurrencyCode = 214
	CurrencyEgyptianPound                  CurrencyCode = 818
	CurrencyElSalvadorColon                CurrencyCode = 222
	CurrencyNakfa                          CurrencyCode = 232
	CurrencyEthiopianBirr                  CurrencyCode = 230
	CurrencyFalklandIslandsPound           CurrencyCode = 238
	CurrencyFijiDollar                     CurrencyCode = 242
	CurrencyCFPFranc                       CurrencyCode = 953
	CurrencyDalasi                         CurrencyCode = 270
	CurrencyLari                           CurrencyCode = 981
	CurrencyGhanaCedi                      CurrencyCode = 936
	CurrencyGibraltarPound                 CurrencyCode = 292
	CurrencyQuetzal                        CurrencyCode = 320
	CurrencyPoundSterling                  CurrencyCode = 826
	CurrencyGuineaFranc                    CurrencyCode = 324
	CurrencyGuyanaDollar                   CurrencyCode = 328
	CurrencyGourde                         CurrencyCode = 332
	CurrencyLempira                        CurrencyCode = 340
	CurrencyHongKongDollar                 CurrencyCode = 344
	CurrencyForint                         CurrencyCode = 348
	CurrencyIcelandKrona                   CurrencyCode = 352
	CurrencyRupiah                         CurrencyCode = 360
	CurrencySDR                            CurrencyCode = 960
	CurrencyIranianRial                    CurrencyCode = 364
	CurrencyIraqiDinar                     CurrencyCode = 368
	CurrencyNewIsraeliSheqel               CurrencyCode = 376
	CurrencyJamaicanDollar                 CurrencyCode = 388
	CurrencyYen                            CurrencyCode = 392
	CurrencyJordanianDinar                 CurrencyCode = 400
	CurrencyTenge                          CurrencyCode = 398
	CurrencyKenyanShilling                 CurrencyCode = 404
	CurrencyNorthKoreanWon                 CurrencyCode = 408
	CurrencyWon                            CurrencyCode = 410
	CurrencyKuwaitiDinar                   CurrencyCode = 414
	CurrencySom                            CurrencyCode = 417
	CurrencyKip                            CurrencyCode = 418
	CurrencyLebanesePound                  CurrencyCode = 422
	CurrencyLoti                           CurrencyCode = 426
	CurrencyRand                           CurrencyCode = 710
	CurrencyLiberianDollar                 CurrencyCode = 430
	CurrencyLibyanDinar                    CurrencyCode = 434
	CurrencySwissFranc                     CurrencyCode = 756
	CurrencyPataca                         CurrencyCode = 446
	CurrencyDenar                          CurrencyCode = 807
	CurrencyMalagasyAriary                 CurrencyCode = 969
	CurrencyKwacha                         CurrencyCode = 454
	CurrencyMalaysianRinggit               CurrencyCode = 458
	CurrencyRufiyaa                        CurrencyCode = 462
	CurrencyOuguiya                        CurrencyCode = 929
	CurrencyMauritiusRupee                 CurrencyCode = 480
	CurrencyADBUnitofAccount               CurrencyCode = 965
	CurrencyMexicanPeso                    CurrencyCode = 484
	CurrencyMexicanUnidaddeInversion       CurrencyCode = 979
	CurrencyMexicanUDI                     CurrencyCode = 979
	CurrencyMoldovanLeu                    CurrencyCode = 498
	CurrencyTugrik                         CurrencyCode = 496
	CurrencyMoroccanDirham                 CurrencyCode = 504
	CurrencyMozambiqueMetical              CurrencyCode = 943
	CurrencyKyat                           CurrencyCode = 104
	CurrencyNamibiaDollar                  CurrencyCode = 516
	CurrencyNepaleseRupee                  CurrencyCode = 524
	CurrencyCordobaOro                     CurrencyCode = 558
	CurrencyNaira                          CurrencyCode = 566
	CurrencyRialOmani                      CurrencyCode = 512
	CurrencyPakistanRupee                  CurrencyCode = 586
	CurrencyBalboa                         CurrencyCode = 590
	CurrencyKina                           CurrencyCode = 598
	CurrencyGuarani                        CurrencyCode = 600
	CurrencyNuevoSol                       CurrencyCode = 604
	CurrencyPhilippinePeso                 CurrencyCode = 608
	CurrencyZloty                          CurrencyCode = 985
	CurrencyQatariRial                     CurrencyCode = 634
	CurrencyRomanianLeu                    CurrencyCode = 946
	CurrencyRussianRuble                   CurrencyCode = 643
	CurrencyRwandaFranc                    CurrencyCode = 646
	CurrencySaintHelenaPound               CurrencyCode = 654
	CurrencyTala                           CurrencyCode = 882
	CurrencyDobra                          CurrencyCode = 930
	CurrencySaudiRiyal                     CurrencyCode = 682
	CurrencySerbianDinar                   CurrencyCode = 941
	CurrencySeychellesRupee                CurrencyCode = 690
	CurrencyLeone                          CurrencyCode = 694
	CurrencySingaporeDollar                CurrencyCode = 702
	CurrencySucre                          CurrencyCode = 994
	CurrencySolomonIslandsDollar           CurrencyCode = 90
	CurrencySomaliShilling                 CurrencyCode = 706
	CurrencySouthSudanesePound             CurrencyCode = 728
	CurrencySriLankaRupee                  CurrencyCode = 144
	CurrencySudanesePound                  CurrencyCode = 938
	CurrencySurinamDollar                  CurrencyCode = 968
	CurrencyLilangeni                      CurrencyCode = 748
	CurrencySwedishKrona                   CurrencyCode = 752
	CurrencyWIREuro                        CurrencyCode = 947
	CurrencyWIRFranc                       CurrencyCode = 948
	CurrencySyrianPound                    CurrencyCode = 760
	CurrencyNewTaiwanDollar                CurrencyCode = 901
	CurrencySomoni                         CurrencyCode = 972
	CurrencyTanzanianShilling              CurrencyCode = 834
	CurrencyBaht                           CurrencyCode = 764
	CurrencyPaanga                         CurrencyCode = 776
	CurrencyTrinidadandTobagoDollar        CurrencyCode = 780
	CurrencyTunisianDinar                  CurrencyCode = 788
	CurrencyTurkishLira                    CurrencyCode = 949
	CurrencyTurkmenistanNewManat           CurrencyCode = 934
	CurrencyUgandaShilling                 CurrencyCode = 800
	CurrencyHryvnia                        CurrencyCode = 980
	CurrencyUAEDirham                      CurrencyCode = 784
	CurrencyUSDollarNextday                CurrencyCode = 997
	CurrencyUruguayPesoenUnidadesIndexadas CurrencyCode = 940
	CurrencyUruguayPUI                     CurrencyCode = 940
	CurrencyURUIURUI                       CurrencyCode = 940
	CurrencyPesoUruguayo                   CurrencyCode = 858
	CurrencyUzbekistanSum                  CurrencyCode = 860
	CurrencyVatu                           CurrencyCode = 548
	CurrencyBolivar                        CurrencyCode = 937
	CurrencyDong                           CurrencyCode = 704
	CurrencyYemeniRial                     CurrencyCode = 886
	CurrencyZambianKwacha                  CurrencyCode = 967
	CurrencyZimbabweDollar                 CurrencyCode = 932
	CurrencyYugoslavianDinar               CurrencyCode = 891
)

// Currencies by ISO 4217. Two codes present, for example CurrencyUSDollar == CurrencyUSD == 840.
const (
	CurrencyAFN CurrencyCode = 971
	CurrencyALL CurrencyCode = 8
	CurrencyDZD CurrencyCode = 12
	CurrencyUSD CurrencyCode = 840
	CurrencyEUR CurrencyCode = 978
	CurrencyAOA CurrencyCode = 973
	CurrencyXCD CurrencyCode = 951
	CurrencyARS CurrencyCode = 32
	CurrencyAMD CurrencyCode = 51
	CurrencyAWG CurrencyCode = 533
	CurrencyAUD CurrencyCode = 36
	CurrencyAZN CurrencyCode = 944
	CurrencyBSD CurrencyCode = 44
	CurrencyBHD CurrencyCode = 48
	CurrencyBDT CurrencyCode = 50
	CurrencyBBD CurrencyCode = 52
	CurrencyBYR CurrencyCode = 974
	CurrencyBZD CurrencyCode = 84
	CurrencyXOF CurrencyCode = 952
	CurrencyBMD CurrencyCode = 60
	CurrencyBTN CurrencyCode = 64
	CurrencyINR CurrencyCode = 356
	CurrencyBOB CurrencyCode = 68
	CurrencyBAM CurrencyCode = 977
	CurrencyBWP CurrencyCode = 72
	CurrencyNOK CurrencyCode = 578
	CurrencyBRL CurrencyCode = 986
	CurrencyBND CurrencyCode = 96
	CurrencyBGN CurrencyCode = 975
	CurrencyBIF CurrencyCode = 108
	CurrencyCVE CurrencyCode = 132
	CurrencyKHR CurrencyCode = 116
	CurrencyXAF CurrencyCode = 950
	CurrencyCAD CurrencyCode = 124
	CurrencyKYD CurrencyCode = 136
	CurrencyCLF CurrencyCode = 990
	CurrencyCLP CurrencyCode = 152
	CurrencyCNY CurrencyCode = 156
	CurrencyCOP CurrencyCode = 170
	CurrencyCOU CurrencyCode = 970
	CurrencyKMF CurrencyCode = 174
	CurrencyCDF CurrencyCode = 976
	CurrencyNZD CurrencyCode = 554
	CurrencyCRC CurrencyCode = 188
	CurrencyHRK CurrencyCode = 191
	CurrencyCUC CurrencyCode = 931
	CurrencyCUP CurrencyCode = 192
	CurrencyANG CurrencyCode = 532
	CurrencyCZK CurrencyCode = 203
	CurrencyDKK CurrencyCode = 208
	CurrencyDJF CurrencyCode = 262
	CurrencyDOP CurrencyCode = 214
	CurrencyEGP CurrencyCode = 818
	CurrencySVC CurrencyCode = 222
	CurrencyERN CurrencyCode = 232
	CurrencyETB CurrencyCode = 230
	CurrencyFKP CurrencyCode = 238
	CurrencyFJD CurrencyCode = 242
	CurrencyXPF CurrencyCode = 953
	CurrencyGMD CurrencyCode = 270
	CurrencyGEL CurrencyCode = 981
	CurrencyGHS CurrencyCode = 936
	CurrencyGIP CurrencyCode = 292
	CurrencyGTQ CurrencyCode = 320
	CurrencyGBP CurrencyCode = 826
	CurrencyGNF CurrencyCode = 324
	CurrencyGYD CurrencyCode = 328
	CurrencyHTG CurrencyCode = 332
	CurrencyHNL CurrencyCode = 340
	CurrencyHKD CurrencyCode = 344
	CurrencyHUF CurrencyCode = 348
	CurrencyISK CurrencyCode = 352
	CurrencyIDR CurrencyCode = 360
	CurrencyXDR CurrencyCode = 960
	CurrencyIRR CurrencyCode = 364
	CurrencyIQD CurrencyCode = 368
	CurrencyILS CurrencyCode = 376
	CurrencyJMD CurrencyCode = 388
	CurrencyJPY CurrencyCode = 392
	CurrencyJOD CurrencyCode = 400
	CurrencyKZT CurrencyCode = 398
	CurrencyKES CurrencyCode = 404
	CurrencyKPW CurrencyCode = 408
	CurrencyKRW CurrencyCode = 410
	CurrencyKWD CurrencyCode = 414
	CurrencyKGS CurrencyCode = 417
	CurrencyLAK CurrencyCode = 418
	CurrencyLBP CurrencyCode = 422
	CurrencyLSL CurrencyCode = 426
	CurrencyZAR CurrencyCode = 710
	CurrencyLRD CurrencyCode = 430
	CurrencyLYD CurrencyCode = 434
	CurrencyCHF CurrencyCode = 756
	CurrencyMOP CurrencyCode = 446
	CurrencyMKD CurrencyCode = 807
	CurrencyMGA CurrencyCode = 969
	CurrencyMWK CurrencyCode = 454
	CurrencyMYR CurrencyCode = 458
	CurrencyMVR CurrencyCode = 462
	CurrencyMRU CurrencyCode = 929
	CurrencyMUR CurrencyCode = 480
	CurrencyXUA CurrencyCode = 965
	CurrencyMXN CurrencyCode = 484
	CurrencyMXV CurrencyCode = 979
	CurrencyMDL CurrencyCode = 498
	CurrencyMNT CurrencyCode = 496
	CurrencyMAD CurrencyCode = 504
	CurrencyMZN CurrencyCode = 943
	CurrencyMMK CurrencyCode = 104
	CurrencyNAD CurrencyCode = 516
	CurrencyNPR CurrencyCode = 524
	CurrencyNIO CurrencyCode = 558
	CurrencyNGN CurrencyCode = 566
	CurrencyOMR CurrencyCode = 512
	CurrencyPKR CurrencyCode = 586
	CurrencyPAB CurrencyCode = 590
	CurrencyPGK CurrencyCode = 598
	CurrencyPYG CurrencyCode = 600
	CurrencyPEN CurrencyCode = 604
	CurrencyPHP CurrencyCode = 608
	CurrencyPLN CurrencyCode = 985
	CurrencyQAR CurrencyCode = 634
	CurrencyRON CurrencyCode = 946
	CurrencyRUB CurrencyCode = 643
	CurrencyRWF CurrencyCode = 646
	CurrencySHP CurrencyCode = 654
	CurrencyWST CurrencyCode = 882
	CurrencySTN CurrencyCode = 930
	CurrencySAR CurrencyCode = 682
	CurrencyRSD CurrencyCode = 941
	CurrencySCR CurrencyCode = 690
	CurrencySLL CurrencyCode = 694
	CurrencySGD CurrencyCode = 702
	CurrencyXSU CurrencyCode = 994
	CurrencySBD CurrencyCode = 90
	CurrencySOS CurrencyCode = 706
	CurrencySSP CurrencyCode = 728
	CurrencyLKR CurrencyCode = 144
	CurrencySDG CurrencyCode = 938
	CurrencySRD CurrencyCode = 968
	CurrencySZL CurrencyCode = 748
	CurrencySEK CurrencyCode = 752
	CurrencyCHE CurrencyCode = 947
	CurrencyCHW CurrencyCode = 948
	CurrencySYP CurrencyCode = 760
	CurrencyTWD CurrencyCode = 901
	CurrencyTJS CurrencyCode = 972
	CurrencyTZS CurrencyCode = 834
	CurrencyTHB CurrencyCode = 764
	CurrencyTOP CurrencyCode = 776
	CurrencyTTD CurrencyCode = 780
	CurrencyTND CurrencyCode = 788
	CurrencyTRY CurrencyCode = 949
	CurrencyTMT CurrencyCode = 934
	CurrencyUGX CurrencyCode = 800
	CurrencyUAH CurrencyCode = 980
	CurrencyAED CurrencyCode = 784
	CurrencyUSN CurrencyCode = 997
	CurrencyUYI CurrencyCode = 940
	CurrencyUYU CurrencyCode = 858
	CurrencyUZS CurrencyCode = 860
	CurrencyVUV CurrencyCode = 548
	CurrencyVEF CurrencyCode = 937
	CurrencyVND CurrencyCode = 704
	CurrencyYER CurrencyCode = 886
	CurrencyZMW CurrencyCode = 967
	CurrencyZWL CurrencyCode = 932
	CurrencyYUD CurrencyCode = 891
)

// Emoji - return a currency as Emoji (only for USD, EUR, JPY and GBP)
func (c CurrencyCode) Emoji() string {
	switch c {
	case CurrencyUSD:
		return "💵"
	case CurrencyEUR:
		return "💶"
	case CurrencyJPY:
		return "💴"
	case CurrencyGBP:
		return "💷"
	}
	return c.Alpha()
}

// TotalCurrencies - returns number of currencies in the package, countries.TotalCurrencies() == len(countries.AllCurrencies()) but static value for performance
func TotalCurrencies() int {
	return 168
}

// Type implements Typer interface
func (c CurrencyCode) Type() string {
	return TypeCurrencyCode
}

// String - implements fmt.Stringer, returns a english name of currency
func (c CurrencyCode) String() string {
	switch c {
	case 840:
		return "US Dollar"
	case 978:
		return "Euro"
	case 971:
		return "Afghani"
	case 8:
		return "Lek"
	case 12:
		return "Algerian Dinar"
	case 973:
		return "Kwanza"
	case 951:
		return "East Caribbean Dollar"
	case 32:
		return "Argentine Peso"
	case 51:
		return "Armenian Dram"
	case 533:
		return "Aruban Florin"
	case 36:
		return "Australian Dollar"
	case 944:
		return "Azerbaijanian Manat"
	case 44:
		return "Bahamian Dollar"
	case 48:
		return "Bahraini Dinar"
	case 50:
		return "Taka"
	case 52:
		return "Barbados Dollar"
	case 974:
		return "Belarussian Ruble"
	case 84:
		return "Belize Dollar"
	case 952:
		return "CFA Franc BCEAO"
	case 60:
		return "Bermudian Dollar"
	case 64:
		return "Ngultrum"
	case 356:
		return "Indian Rupee"
	case 68:
		return "Boliviano"
	case 977:
		return "Convertible Mark"
	case 72:
		return "Pula"
	case 578:
		return "Norwegian Krone"
	case 986:
		return "Brazilian Real"
	case 96:
		return "Brunei Dollar"
	case 975:
		return "Bulgarian Lev"
	case 108:
		return "Burundi Franc"
	case 132:
		return "Cabo Verde Escudo"
	case 116:
		return "Riel"
	case 950:
		return "CFA Franc BEAC"
	case 124:
		return "Canadian Dollar"
	case 136:
		return "Cayman Islands Dollar"
	case 990:
		return "Unidad de Fomento"
	case 152:
		return "Chilean Peso"
	case 156:
		return "Yuan Renminbi"
	case 170:
		return "Colombian Peso"
	case 970:
		return "Unidad de Valor Real"
	case 174:
		return "Comoro Franc"
	case 976:
		return "Congolese Franc"
	case 554:
		return "New Zealand Dollar"
	case 188:
		return "Costa Rican Colon"
	case 191:
		return "Kuna"
	case 931:
		return "Peso Convertible"
	case 192:
		return "Cuban Peso"
	case 532:
		return "Netherlands Antillean Guilder"
	case 203:
		return "Czech Koruna"
	case 208:
		return "Danish Krone"
	case 262:
		return "Djibouti Franc"
	case 214:
		return "Dominican Peso"
	case 818:
		return "Egyptian Pound"
	case 222:
		return "El Salvador Colon"
	case 232:
		return "Nakfa"
	case 230:
		return "Ethiopian Birr"
	case 238:
		return "Falkland Islands Pound"
	case 242:
		return "Fiji Dollar"
	case 953:
		return "CFP Franc"
	case 270:
		return "Dalasi"
	case 981:
		return "Lari"
	case 936:
		return "Ghana Cedi"
	case 292:
		return "Gibraltar Pound"
	case 320:
		return "Quetzal"
	case 826:
		return "Pound Sterling"
	case 324:
		return "Guinea Franc"
	case 328:
		return "Guyana Dollar"
	case 332:
		return "Gourde"
	case 340:
		return "Lempira"
	case 344:
		return "Hong Kong Dollar"
	case 348:
		return "Forint"
	case 352:
		return "Iceland Krona"
	case 360:
		return "Rupiah"
	case 960:
		return "SDR (Special Drawing Right)"
	case 364:
		return "Iranian Rial"
	case 368:
		return "Iraqi Dinar"
	case 376:
		return "New Israeli Sheqel"
	case 388:
		return "Jamaican Dollar"
	case 392:
		return "Yen"
	case 400:
		return "Jordanian Dinar"
	case 398:
		return "Tenge"
	case 404:
		return "Kenyan Shilling"
	case 408:
		return "North Korean Won"
	case 410:
		return "Won"
	case 414:
		return "Kuwaiti Dinar"
	case 417:
		return "Som"
	case 418:
		return "Kip"
	case 422:
		return "Lebanese Pound"
	case 426:
		return "Loti"
	case 710:
		return "Rand"
	case 430:
		return "Liberian Dollar"
	case 434:
		return "Libyan Dinar"
	case 756:
		return "Swiss Franc"
	case 446:
		return "Pataca"
	case 807:
		return "Denar"
	case 969:
		return "Malagasy Ariary"
	case 454:
		return "Kwacha"
	case 458:
		return "Malaysian Ringgit"
	case 462:
		return "Rufiyaa"
	case 929:
		return "Ouguiya"
	case 480:
		return "Mauritius Rupee"
	case 965:
		return "ADB Unit of Account"
	case 484:
		return "Mexican Peso"
	case 979:
		return "Mexican Unidad de Inversion (UDI)"
	case 498:
		return "Moldovan Leu"
	case 496:
		return "Tugrik"
	case 504:
		return "Moroccan Dirham"
	case 943:
		return "Mozambique Metical"
	case 104:
		return "Kyat"
	case 516:
		return "Namibia Dollar"
	case 524:
		return "Nepalese Rupee"
	case 558:
		return "Cordoba Oro"
	case 566:
		return "Naira"
	case 512:
		return "Rial Omani"
	case 586:
		return "Pakistan Rupee"
	case 590:
		return "Balboa"
	case 598:
		return "Kina"
	case 600:
		return "Guarani"
	case 604:
		return "Nuevo Sol"
	case 608:
		return "Philippine Peso"
	case 985:
		return "Zloty"
	case 634:
		return "Qatari Rial"
	case 946:
		return "Romanian Leu"
	case 643:
		return "Russian Ruble"
	case 646:
		return "Rwanda Franc"
	case 654:
		return "Saint Helena Pound"
	case 882:
		return "Tala"
	case 930:
		return "Dobra"
	case 682:
		return "Saudi Riyal"
	case 941:
		return "Serbian Dinar"
	case 690:
		return "Seychelles Rupee"
	case 694:
		return "Leone"
	case 702:
		return "Singapore Dollar"
	case 994:
		return "Sucre"
	case 90:
		return "Solomon Islands Dollar"
	case 706:
		return "Somali Shilling"
	case 728:
		return "South Sudanese Pound"
	case 144:
		return "Sri Lanka Rupee"
	case 938:
		return "Sudanese Pound"
	case 968:
		return "Surinam Dollar"
	case 748:
		return "Lilangeni"
	case 752:
		return "Swedish Krona"
	case 947:
		return "WIR Euro"
	case 948:
		return "WIR Franc"
	case 760:
		return "Syrian Pound"
	case 901:
		return "New Taiwan Dollar"
	case 972:
		return "Somoni"
	case 834:
		return "Tanzanian Shilling"
	case 764:
		return "Baht"
	case 776:
		return "Pa’anga"
	case 780:
		return "Trinidad and Tobago Dollar"
	case 788:
		return "Tunisian Dinar"
	case 949:
		return "Turkish Lira"
	case 934:
		return "Turkmenistan New Manat"
	case 800:
		return "Uganda Shilling"
	case 980:
		return "Hryvnia"
	case 784:
		return "UAE Dirham"
	case 997:
		return "US Dollar (Next day)"
	case 940:
		return "Uruguay Peso en Unidades Indexadas (URUIURUI)"
	case 858:
		return "Peso Uruguayo"
	case 860:
		return "Uzbekistan Sum"
	case 548:
		return "Vatu"
	case 937:
		return "Bolivar"
	case 704:
		return "Dong"
	case 886:
		return "Yemeni Rial"
	case 967:
		return "Zambian Kwacha"
	case 932:
		return "Zimbabwe Dollar"
	case 891:
		return "Yugoslavian Dinar"
	}
	return UnknownMsg
}

// Alpha - returns a default Alpha (3 chars) code of currency
func (c CurrencyCode) Alpha() string {
	switch c {
	case 840:
		return "USD"
	case 978:
		return "EUR"
	case 971:
		return "AFN"
	case 8:
		return "ALL"
	case 12:
		return "DZD"
	case 973:
		return "AOA"
	case 951:
		return "XCD"
	case 32:
		return "ARS"
	case 51:
		return "AMD"
	case 533:
		return "AWG"
	case 36:
		return "AUD"
	case 944:
		return "AZN"
	case 44:
		return "BSD"
	case 48:
		return "BHD"
	case 50:
		return "BDT"
	case 52:
		return "BBD"
	case 974:
		return "BYR"
	case 84:
		return "BZD"
	case 952:
		return "XOF"
	case 60:
		return "BMD"
	case 64:
		return "BTN"
	case 356:
		return "INR"
	case 68:
		return "BOB"
	case 977:
		return "BAM"
	case 72:
		return "BWP"
	case 578:
		return "NOK"
	case 986:
		return "BRL"
	case 96:
		return "BND"
	case 975:
		return "BGN"
	case 108:
		return "BIF"
	case 132:
		return "CVE"
	case 116:
		return "KHR"
	case 950:
		return "XAF"
	case 124:
		return "CAD"
	case 136:
		return "KYD"
	case 990:
		return "CLF"
	case 152:
		return "CLP"
	case 156:
		return "CNY"
	case 170:
		return "COP"
	case 970:
		return "COU"
	case 174:
		return "KMF"
	case 976:
		return "CDF"
	case 554:
		return "NZD"
	case 188:
		return "CRC"
	case 191:
		return "HRK"
	case 931:
		return "CUC"
	case 192:
		return "CUP"
	case 532:
		return "ANG"
	case 203:
		return "CZK"
	case 208:
		return "DKK"
	case 262:
		return "DJF"
	case 214:
		return "DOP"
	case 818:
		return "EGP"
	case 222:
		return "SVC"
	case 232:
		return "ERN"
	case 230:
		return "ETB"
	case 238:
		return "FKP"
	case 242:
		return "FJD"
	case 953:
		return "XPF"
	case 270:
		return "GMD"
	case 981:
		return "GEL"
	case 936:
		return "GHS"
	case 292:
		return "GIP"
	case 320:
		return "GTQ"
	case 826:
		return "GBP"
	case 324:
		return "GNF"
	case 328:
		return "GYD"
	case 332:
		return "HTG"
	case 340:
		return "HNL"
	case 344:
		return "HKD"
	case 348:
		return "HUF"
	case 352:
		return "ISK"
	case 360:
		return "IDR"
	case 960:
		return "XDR"
	case 364:
		return "IRR"
	case 368:
		return "IQD"
	case 376:
		return "ILS"
	case 388:
		return "JMD"
	case 392:
		return "JPY"
	case 400:
		return "JOD"
	case 398:
		return "KZT"
	case 404:
		return "KES"
	case 408:
		return "KPW"
	case 410:
		return "KRW"
	case 414:
		return "KWD"
	case 417:
		return "KGS"
	case 418:
		return "LAK"
	case 422:
		return "LBP"
	case 426:
		return "LSL"
	case 710:
		return "ZAR"
	case 430:
		return "LRD"
	case 434:
		return "LYD"
	case 756:
		return "CHF"
	case 446:
		return "MOP"
	case 807:
		return "MKD"
	case 969:
		return "MGA"
	case 454:
		return "MWK"
	case 458:
		return "MYR"
	case 462:
		return "MVR"
	case 929:
		return "MRU"
	case 480:
		return "MUR"
	case 965:
		return "XUA"
	case 484:
		return "MXN"
	case 979:
		return "MXV"
	case 498:
		return "MDL"
	case 496:
		return "MNT"
	case 504:
		return "MAD"
	case 943:
		return "MZN"
	case 104:
		return "MMK"
	case 516:
		return "NAD"
	case 524:
		return "NPR"
	case 558:
		return "NIO"
	case 566:
		return "NGN"
	case 512:
		return "OMR"
	case 586:
		return "PKR"
	case 590:
		return "PAB"
	case 598:
		return "PGK"
	case 600:
		return "PYG"
	case 604:
		return "PEN"
	case 608:
		return "PHP"
	case 985:
		return "PLN"
	case 634:
		return "QAR"
	case 946:
		return "RON"
	case 643:
		return "RUB"
	case 646:
		return "RWF"
	case 654:
		return "SHP"
	case 882:
		return "WST"
	case 930:
		return "STN"
	case 682:
		return "SAR"
	case 941:
		return "RSD"
	case 690:
		return "SCR"
	case 694:
		return "SLL"
	case 702:
		return "SGD"
	case 994:
		return "XSU"
	case 90:
		return "SBD"
	case 706:
		return "SOS"
	case 728:
		return "SSP"
	case 144:
		return "LKR"
	case 938:
		return "SDG"
	case 968:
		return "SRD"
	case 748:
		return "SZL"
	case 752:
		return "SEK"
	case 947:
		return "CHE"
	case 948:
		return "CHW"
	case 760:
		return "SYP"
	case 901:
		return "TWD"
	case 972:
		return "TJS"
	case 834:
		return "TZS"
	case 764:
		return "THB"
	case 776:
		return "TOP"
	case 780:
		return "TTD"
	case 788:
		return "TND"
	case 949:
		return "TRY"
	case 934:
		return "TMT"
	case 800:
		return "UGX"
	case 980:
		return "UAH"
	case 784:
		return "AED"
	case 997:
		return "USN"
	case 940:
		return "UYI"
	case 858:
		return "UYU"
	case 860:
		return "UZS"
	case 548:
		return "VUV"
	case 937:
		return "VEF"
	case 704:
		return "VND"
	case 886:
		return "YER"
	case 967:
		return "ZMW"
	case 932:
		return "ZWL"
	case 891:
		return "YUD"
	}
	return UnknownMsg
}

// IsValid - returns true, if code is correct
func (c CurrencyCode) IsValid() bool {
	return c.Alpha() != UnknownMsg
}

// Countries - returns a country codes of currency using
func (c CurrencyCode) Countries() []CountryCode {
	switch c {
	case CurrencyAFN:
		return []CountryCode{AFG}
	case CurrencyDZD:
		return []CountryCode{DZA}
	case CurrencyARS:
		return []CountryCode{ARG}
	case CurrencyAMD:
		return []CountryCode{ARM}
	case CurrencyAWG:
		return []CountryCode{ABW}
	case CurrencyAUD:
		return []CountryCode{AUS, CXR, CCK, HMD, KIR, NRU, NFK, TUV}
	case CurrencyAZN:
		return []CountryCode{AZE}
	case CurrencyBSD:
		return []CountryCode{BHS}
	case CurrencyBHD:
		return []CountryCode{BHR}
	case CurrencyTHB:
		return []CountryCode{THA}
	case CurrencyPAB:
		return []CountryCode{PAN}
	case CurrencyBBD:
		return []CountryCode{BRB}
	case CurrencyBYR:
		return []CountryCode{BLR}
	case CurrencyBZD:
		return []CountryCode{BLZ}
	case CurrencyBMD:
		return []CountryCode{BMU}
	case CurrencyVEF:
		return []CountryCode{VEN}
	case CurrencyBOB:
		return []CountryCode{BOL}
	case CurrencyBRL:
		return []CountryCode{BRA}
	case CurrencyBND:
		return []CountryCode{BRN}
	case CurrencyBGN:
		return []CountryCode{BGR}
	case CurrencyBIF:
		return []CountryCode{BDI}
	case CurrencyCVE:
		return []CountryCode{CPV}
	case CurrencyCAD:
		return []CountryCode{CAN}
	case CurrencyKYD:
		return []CountryCode{CYM}
	case CurrencyXOF:
		return []CountryCode{BEN, BFA, CIV, GNB, MLI, NER, SEN, TGO}
	case CurrencyXAF:
		return []CountryCode{CMR, CAF, TCD, COG, GNQ, GAB}
	case CurrencyXPF:
		return []CountryCode{PYF, NCL, WLF}
	case CurrencyCLP:
		return []CountryCode{CHL}
	case CurrencyCOP:
		return []CountryCode{COL}
	case CurrencyKMF:
		return []CountryCode{COM}
	case CurrencyCDF:
		return []CountryCode{COD}
	case CurrencyBAM:
		return []CountryCode{BIH}
	case CurrencyNIO:
		return []CountryCode{NIC}
	case CurrencyCRC:
		return []CountryCode{CRI}
	case CurrencyCUP:
		return []CountryCode{CUB}
	case CurrencyCZK:
		return []CountryCode{CZE}
	case CurrencyGMD:
		return []CountryCode{GMB}
	case CurrencyDKK:
		return []CountryCode{DNK, FRO, GRL}
	case CurrencyMKD:
		return []CountryCode{MKD}
	case CurrencyDJF:
		return []CountryCode{DJI}
	case CurrencySTN:
		return []CountryCode{STP}
	case CurrencyDOP:
		return []CountryCode{DOM}
	case CurrencyVND:
		return []CountryCode{VNM}
	case CurrencyXCD:
		return []CountryCode{AIA, ATG, DMA, GRD, MSR, KNA, LCA, VCT}
	case CurrencyEGP:
		return []CountryCode{EGY}
	case CurrencySVC:
		return []CountryCode{SLV}
	case CurrencyETB:
		return []CountryCode{ETH}
	case CurrencyEUR:
		return []CountryCode{AND, AUT, BEL, CYP, EST, FIN, FRA, GUF, ATF, DEU, GRC, GLP, VAT, IRL, ITA, LVA, LTU, LUX, MLT, MTQ, MYT, MCO, MNE, NLD, PRT, REU, BLM, MAF, SPM, SMR, SVK, SVN, ESP, ALA}
	case CurrencyFKP:
		return []CountryCode{FLK}
	case CurrencyFJD:
		return []CountryCode{FJI}
	case CurrencyHUF:
		return []CountryCode{HUN}
	case CurrencyGHS:
		return []CountryCode{GHA}
	case CurrencyGIP:
		return []CountryCode{GIB}
	case CurrencyHTG:
		return []CountryCode{HTI}
	case CurrencyPYG:
		return []CountryCode{PRY}
	case CurrencyGNF:
		return []CountryCode{GIN}
	case CurrencyGYD:
		return []CountryCode{GUY}
	case CurrencyHKD:
		return []CountryCode{HKG}
	case CurrencyUAH:
		return []CountryCode{UKR}
	case CurrencyISK:
		return []CountryCode{ISL}
	case CurrencyINR:
		return []CountryCode{BTN, IND}
	case CurrencyIRR:
		return []CountryCode{IRN}
	case CurrencyIQD:
		return []CountryCode{IRQ}
	case CurrencyJMD:
		return []CountryCode{JAM}
	case CurrencyJOD:
		return []CountryCode{JOR}
	case CurrencyKES:
		return []CountryCode{KEN}
	case CurrencyPGK:
		return []CountryCode{PNG}
	case CurrencyLAK:
		return []CountryCode{LAO}
	case CurrencyHRK:
		return []CountryCode{HRV}
	case CurrencyKWD:
		return []CountryCode{KWT}
	case CurrencyMWK:
		return []CountryCode{MWI}
	case CurrencyAOA:
		return []CountryCode{AGO}
	case CurrencyMMK:
		return []CountryCode{MMR}
	case CurrencyGEL:
		return []CountryCode{GEO}
	case CurrencyLBP:
		return []CountryCode{LBN}
	case CurrencyALL:
		return []CountryCode{ALB}
	case CurrencyHNL:
		return []CountryCode{HND}
	case CurrencySLL:
		return []CountryCode{SLE}
	case CurrencyLRD:
		return []CountryCode{LBR}
	case CurrencyLYD:
		return []CountryCode{LBY}
	case CurrencySZL:
		return []CountryCode{SWZ}
	case CurrencyLSL:
		return []CountryCode{LSO}
	case CurrencyMGA:
		return []CountryCode{MDG}
	case CurrencyMYR:
		return []CountryCode{MYS}
	case CurrencyMUR:
		return []CountryCode{MUS}
	case CurrencyMXN:
		return []CountryCode{MEX}
	case CurrencyMXV:
		return []CountryCode{MEX}
	case CurrencyMDL:
		return []CountryCode{MDA}
	case CurrencyMAD:
		return []CountryCode{MAR, ESH}
	case CurrencyMZN:
		return []CountryCode{MOZ}
	case CurrencyNGN:
		return []CountryCode{NGA}
	case CurrencyERN:
		return []CountryCode{ERI}
	case CurrencyNAD:
		return []CountryCode{NAM}
	case CurrencyNPR:
		return []CountryCode{NPL}
	case CurrencyANG:
		return []CountryCode{CUW, SXM, ANT}
	case CurrencyILS:
		return []CountryCode{ISR, PSE}
	case CurrencyTWD:
		return []CountryCode{TWN}
	case CurrencyNZD:
		return []CountryCode{COK, NZL, NIU, PCN, TKL}
	case CurrencyBTN:
		return []CountryCode{BTN}
	case CurrencyKPW:
		return []CountryCode{PRK}
	case CurrencyNOK:
		return []CountryCode{BVT, NOR, SJM}
	case CurrencyPEN:
		return []CountryCode{PER}
	case CurrencyMRU:
		return []CountryCode{MRT}
	case CurrencyTOP:
		return []CountryCode{TON}
	case CurrencyPKR:
		return []CountryCode{PAK}
	case CurrencyMOP:
		return []CountryCode{MAC}
	case CurrencyCUC:
		return []CountryCode{CUB}
	case CurrencyUYU:
		return []CountryCode{URY}
	case CurrencyPHP:
		return []CountryCode{PHL}
	case CurrencyGBP:
		return []CountryCode{GGY, IMN, JEY, GBR, XWA, SGS, XSC}
	case CurrencyBWP:
		return []CountryCode{BWA}
	case CurrencyQAR:
		return []CountryCode{QAT}
	case CurrencyGTQ:
		return []CountryCode{GTM}
	case CurrencyZAR:
		return []CountryCode{LSO, NAM, ZAF}
	case CurrencyOMR:
		return []CountryCode{OMN}
	case CurrencyKHR:
		return []CountryCode{KHM}
	case CurrencyRON:
		return []CountryCode{ROU}
	case CurrencyMVR:
		return []CountryCode{MDV}
	case CurrencyIDR:
		return []CountryCode{IDN}
	case CurrencyRUB:
		return []CountryCode{RUS}
	case CurrencyRWF:
		return []CountryCode{RWA}
	case CurrencySHP:
		return []CountryCode{SHN}
	case CurrencySAR:
		return []CountryCode{SAU}
	case CurrencyRSD:
		return []CountryCode{SRB}
	case CurrencySCR:
		return []CountryCode{SYC}
	case CurrencySGD:
		return []CountryCode{SGP}
	case CurrencySBD:
		return []CountryCode{SLB}
	case CurrencyKGS:
		return []CountryCode{KGZ}
	case CurrencySOS:
		return []CountryCode{SOM}
	case CurrencyTJS:
		return []CountryCode{TJK}
	case CurrencySSP:
		return []CountryCode{SSD}
	case CurrencyLKR:
		return []CountryCode{LKA}
	case CurrencySDG:
		return []CountryCode{SDN}
	case CurrencySRD:
		return []CountryCode{SUR}
	case CurrencySEK:
		return []CountryCode{SWE}
	case CurrencyCHF:
		return []CountryCode{LIE, CHE}
	case CurrencySYP:
		return []CountryCode{SYR}
	case CurrencyBDT:
		return []CountryCode{BGD}
	case CurrencyWST:
		return []CountryCode{WSM}
	case CurrencyTZS:
		return []CountryCode{TZA}
	case CurrencyKZT:
		return []CountryCode{KAZ}
	case CurrencyTTD:
		return []CountryCode{TTO}
	case CurrencyMNT:
		return []CountryCode{MNG}
	case CurrencyTND:
		return []CountryCode{TUN}
	case CurrencyTRY:
		return []CountryCode{TUR}
	case CurrencyTMT:
		return []CountryCode{TKM}
	case CurrencyAED:
		return []CountryCode{ARE}
	case CurrencyUGX:
		return []CountryCode{UGA}
	case CurrencyCLF:
		return []CountryCode{CHL}
	case CurrencyCOU:
		return []CountryCode{COL}
	case CurrencyUYI:
		return []CountryCode{URY}
	case CurrencyUSD:
		return []CountryCode{ASM, BES, IOT, ECU, SLV, GUM, HTI, MHL, FSM, MNP, PLW, PAN, PRI, TLS, TCA, UMI, USA, VGB, VIR}
	case CurrencyUSN:
		return []CountryCode{USA}
	case CurrencyUZS:
		return []CountryCode{UZB}
	case CurrencyVUV:
		return []CountryCode{VUT}
	case CurrencyCHE:
		return []CountryCode{CHE}
	case CurrencyCHW:
		return []CountryCode{CHE}
	case CurrencyKRW:
		return []CountryCode{KOR}
	case CurrencyYER:
		return []CountryCode{YEM}
	case CurrencyJPY:
		return []CountryCode{JPN}
	case CurrencyCNY:
		return []CountryCode{CHN}
	case CurrencyYUD:
		return []CountryCode{YUG}
	case CurrencyZMW:
		return []CountryCode{ZMB}
	case CurrencyZWL:
		return []CountryCode{ZWE}
	case CurrencyPLN:
		return []CountryCode{POL}
	}
	return []CountryCode{Unknown}
}

// AllCurrencies - return all currencies codes
func AllCurrencies() []CurrencyCode {
	return []CurrencyCode{
		CurrencyAFN,
		CurrencyALL,
		CurrencyDZD,
		CurrencyUSD,
		CurrencyEUR,
		CurrencyAOA,
		CurrencyXCD,
		CurrencyARS,
		CurrencyAMD,
		CurrencyAWG,
		CurrencyAUD,
		CurrencyAZN,
		CurrencyBSD,
		CurrencyBHD,
		CurrencyBDT,
		CurrencyBBD,
		CurrencyBYR,
		CurrencyBZD,
		CurrencyXOF,
		CurrencyBMD,
		CurrencyBTN,
		CurrencyINR,
		CurrencyBOB,
		CurrencyBAM,
		CurrencyBWP,
		CurrencyNOK,
		CurrencyBRL,
		CurrencyBND,
		CurrencyBGN,
		CurrencyBIF,
		CurrencyCVE,
		CurrencyKHR,
		CurrencyXAF,
		CurrencyCAD,
		CurrencyKYD,
		CurrencyCLF,
		CurrencyCLP,
		CurrencyCNY,
		CurrencyCOP,
		CurrencyCOU,
		CurrencyKMF,
		CurrencyCDF,
		CurrencyNZD,
		CurrencyCRC,
		CurrencyHRK,
		CurrencyCUC,
		CurrencyCUP,
		CurrencyANG,
		CurrencyCZK,
		CurrencyDKK,
		CurrencyDJF,
		CurrencyDOP,
		CurrencyEGP,
		CurrencySVC,
		CurrencyERN,
		CurrencyETB,
		CurrencyFKP,
		CurrencyFJD,
		CurrencyXPF,
		CurrencyGMD,
		CurrencyGEL,
		CurrencyGHS,
		CurrencyGIP,
		CurrencyGTQ,
		CurrencyGBP,
		CurrencyGNF,
		CurrencyGYD,
		CurrencyHTG,
		CurrencyHNL,
		CurrencyHKD,
		CurrencyHUF,
		CurrencyISK,
		CurrencyIDR,
		CurrencyXDR,
		CurrencyIRR,
		CurrencyIQD,
		CurrencyILS,
		CurrencyJMD,
		CurrencyJPY,
		CurrencyJOD,
		CurrencyKZT,
		CurrencyKES,
		CurrencyKPW,
		CurrencyKRW,
		CurrencyKWD,
		CurrencyKGS,
		CurrencyLAK,
		CurrencyLBP,
		CurrencyLSL,
		CurrencyZAR,
		CurrencyLRD,
		CurrencyLYD,
		CurrencyCHF,
		CurrencyMOP,
		CurrencyMKD,
		CurrencyMGA,
		CurrencyMWK,
		CurrencyMYR,
		CurrencyMVR,
		CurrencyMRU,
		CurrencyMUR,
		CurrencyXUA,
		CurrencyMXN,
		CurrencyMXV,
		CurrencyMDL,
		CurrencyMNT,
		CurrencyMAD,
		CurrencyMZN,
		CurrencyMMK,
		CurrencyNAD,
		CurrencyNPR,
		CurrencyNIO,
		CurrencyNGN,
		CurrencyOMR,
		CurrencyPKR,
		CurrencyPAB,
		CurrencyPGK,
		CurrencyPYG,
		CurrencyPEN,
		CurrencyPHP,
		CurrencyPLN,
		CurrencyQAR,
		CurrencyRON,
		CurrencyRUB,
		CurrencyRWF,
		CurrencySHP,
		CurrencyWST,
		CurrencySTN,
		CurrencySAR,
		CurrencyRSD,
		CurrencySCR,
		CurrencySLL,
		CurrencySGD,
		CurrencyXSU,
		CurrencySBD,
		CurrencySOS,
		CurrencySSP,
		CurrencyLKR,
		CurrencySDG,
		CurrencySRD,
		CurrencySZL,
		CurrencySEK,
		CurrencyCHE,
		CurrencyCHW,
		CurrencySYP,
		CurrencyTWD,
		CurrencyTJS,
		CurrencyTZS,
		CurrencyTHB,
		CurrencyTOP,
		CurrencyTTD,
		CurrencyTND,
		CurrencyTRY,
		CurrencyTMT,
		CurrencyUGX,
		CurrencyUAH,
		CurrencyAED,
		CurrencyUSN,
		CurrencyUYI,
		CurrencyUYU,
		CurrencyUZS,
		CurrencyVUV,
		CurrencyVEF,
		CurrencyVND,
		CurrencyYER,
		CurrencyZMW,
		CurrencyYUD,
		CurrencyZWL,
	}
}

// Digits - returns a number of digits used for each currency
func (c CurrencyCode) Digits() int {
	switch c {
	case CurrencyAFN:
		return 0
	case CurrencyALL:
		return 0
	case CurrencyDZD:
		return 2
	case CurrencyUSD:
		return 2
	case CurrencyEUR:
		return 2
	case CurrencyAOA:
		return 2
	case CurrencyXCD:
		return 2
	case CurrencyARS:
		return 2
	case CurrencyAMD:
		return 0
	case CurrencyAWG:
		return 2
	case CurrencyAUD:
		return 2
	case CurrencyAZN:
		return 2
	case CurrencyBSD:
		return 2
	case CurrencyBHD:
		return 3
	case CurrencyBDT:
		return 2
	case CurrencyBBD:
		return 2
	case CurrencyBYR:
		return 0
	case CurrencyBZD:
		return 2
	case CurrencyXOF:
		return 0
	case CurrencyBMD:
		return 2
	case CurrencyBTN:
		return 2
	case CurrencyINR:
		return 2
	case CurrencyBOB:
		return 2
	case CurrencyBAM:
		return 2
	case CurrencyBWP:
		return 2
	case CurrencyNOK:
		return 2
	case CurrencyBRL:
		return 2
	case CurrencyBND:
		return 2
	case CurrencyBGN:
		return 2
	case CurrencyBIF:
		return 0
	case CurrencyCVE:
		return 2
	case CurrencyKHR:
		return 2
	case CurrencyXAF:
		return 0
	case CurrencyCAD:
		return 2
	case CurrencyKYD:
		return 2
	case CurrencyCLF:
		return 4
	case CurrencyCLP:
		return 0
	case CurrencyCNY:
		return 2
	case CurrencyCOP:
		return 0
	case CurrencyCOU:
		return 2
	case CurrencyKMF:
		return 0
	case CurrencyCDF:
		return 2
	case CurrencyNZD:
		return 2
	case CurrencyCRC:
		return 2
	case CurrencyHRK:
		return 2
	case CurrencyCUC:
		return 2
	case CurrencyCUP:
		return 2
	case CurrencyANG:
		return 2
	case CurrencyCZK:
		return 2
	case CurrencyDKK:
		return 2
	case CurrencyDJF:
		return 0
	case CurrencyDOP:
		return 2
	case CurrencyEGP:
		return 2
	case CurrencySVC:
		return 2
	case CurrencyERN:
		return 2
	case CurrencyETB:
		return 2
	case CurrencyFKP:
		return 2
	case CurrencyFJD:
		return 2
	case CurrencyXPF:
		return 0
	case CurrencyGMD:
		return 2
	case CurrencyGEL:
		return 2
	case CurrencyGHS:
		return 2
	case CurrencyGIP:
		return 2
	case CurrencyGTQ:
		return 2
	case CurrencyGBP:
		return 2
	case CurrencyGNF:
		return 0
	case CurrencyGYD:
		return 0
	case CurrencyHTG:
		return 2
	case CurrencyHNL:
		return 2
	case CurrencyHKD:
		return 2
	case CurrencyHUF:
		return 2
	case CurrencyISK:
		return 0
	case CurrencyIDR:
		return 0
	case CurrencyXDR:
		return 2
	case CurrencyIRR:
		return 0
	case CurrencyIQD:
		return 0
	case CurrencyILS:
		return 2
	case CurrencyJMD:
		return 2
	case CurrencyJPY:
		return 0
	case CurrencyJOD:
		return 3
	case CurrencyKZT:
		return 2
	case CurrencyKES:
		return 2
	case CurrencyKPW:
		return 0
	case CurrencyKRW:
		return 0
	case CurrencyKWD:
		return 3
	case CurrencyKGS:
		return 2
	case CurrencyLAK:
		return 0
	case CurrencyLBP:
		return 0
	case CurrencyLSL:
		return 2
	case CurrencyZAR:
		return 2
	case CurrencyLRD:
		return 2
	case CurrencyLYD:
		return 3
	case CurrencyCHF:
		return 2
	case CurrencyMOP:
		return 2
	case CurrencyMKD:
		return 2
	case CurrencyMGA:
		return 0
	case CurrencyMWK:
		return 2
	case CurrencyMYR:
		return 2
	case CurrencyMVR:
		return 2
	case CurrencyMRU:
		return 2
	case CurrencyMUR:
		return 0
	case CurrencyXUA:
		return 2
	case CurrencyMXN:
		return 2
	case CurrencyMXV:
		return 2
	case CurrencyMDL:
		return 2
	case CurrencyMNT:
		return 0
	case CurrencyMAD:
		return 2
	case CurrencyMZN:
		return 2
	case CurrencyMMK:
		return 0
	case CurrencyNAD:
		return 2
	case CurrencyNPR:
		return 2
	case CurrencyNIO:
		return 2
	case CurrencyNGN:
		return 2
	case CurrencyOMR:
		return 3
	case CurrencyPKR:
		return 0
	case CurrencyPAB:
		return 2
	case CurrencyPGK:
		return 2
	case CurrencyPYG:
		return 0
	case CurrencyPEN:
		return 2
	case CurrencyPHP:
		return 2
	case CurrencyPLN:
		return 2
	case CurrencyQAR:
		return 2
	case CurrencyRON:
		return 2
	case CurrencyRUB:
		return 2
	case CurrencyRWF:
		return 0
	case CurrencySHP:
		return 2
	case CurrencyWST:
		return 2
	case CurrencySTN:
		return 2
	case CurrencySAR:
		return 2
	case CurrencyRSD:
		return 0
	case CurrencySCR:
		return 2
	case CurrencySLL:
		return 0
	case CurrencySGD:
		return 2
	case CurrencyXSU:
		return 2
	case CurrencySBD:
		return 2
	case CurrencySOS:
		return 0
	case CurrencySSP:
		return 2
	case CurrencyLKR:
		return 2
	case CurrencySDG:
		return 2
	case CurrencySRD:
		return 2
	case CurrencySZL:
		return 2
	case CurrencySEK:
		return 2
	case CurrencyCHE:
		return 2
	case CurrencyCHW:
		return 2
	case CurrencySYP:
		return 0
	case CurrencyTWD:
		return 2
	case CurrencyTJS:
		return 2
	case CurrencyTZS:
		return 0
	case CurrencyTHB:
		return 2
	case CurrencyTOP:
		return 2
	case CurrencyTTD:
		return 2
	case CurrencyTND:
		return 3
	case CurrencyTRY:
		return 2
	case CurrencyTMT:
		return 2
	case CurrencyUGX:
		return 0
	case CurrencyUAH:
		return 2
	case CurrencyAED:
		return 2
	case CurrencyUSN:
		return 2
	case CurrencyUYI:
		return 0
	case CurrencyUYU:
		return 2
	case CurrencyUZS:
		return 0
	case CurrencyVUV:
		return 0
	case CurrencyVEF:
		return 2
	case CurrencyVND:
		return 0
	case CurrencyYER:
		return 0
	case CurrencyZMW:
		return 2
	case CurrencyYUD:
		return 2
	case CurrencyZWL:
		return 2
	}
	return -1 // never gone here
}

// NickelRounding - returns true, if the currency uses ‘nickel rounding’ in transactions
func (c CurrencyCode) NickelRounding() bool {
	switch c {
	case CurrencyCAD, CurrencyDKK, CurrencyCHF:
		return true
	}
	return false
}

// Info - return all info about currency as Currency struct
func (c CurrencyCode) Info() *Currency {
	return &Currency{
		NickelRounding: c.NickelRounding(),
		Name:           c.String(),
		Alpha:          c.Alpha(),
		Digits:         c.Digits(),
		Code:           c,
		Countries:      c.Countries(),
	}
}

// Type implements Typer interface.
func (currency *Currency) Type() string {
	return TypeCurrency
}

// Value implements database/sql/driver.Valuer
func (currency Currency) Value() (driver.Value, error) {
	return json.Marshal(currency)
}

// Scan implements database/sql.Scanner
func (currency *Currency) Scan(src interface{}) error {
	if currency == nil {
		return fmt.Errorf("countries::Scan: Currency scan err: currency == nil")
	}
	switch src := src.(type) {
	case *Currency:
		*currency = *src
	case Currency:
		*currency = src
	default:
		return fmt.Errorf("countries::Scan: Currency scan err: unexpected value of type %T for %T", src, *currency)
	}
	return nil
}

// AllCurrenciesInfo - return all currencies as []Currency
func AllCurrenciesInfo() []*Currency {
	all := AllCurrencies()
	currencies := make([]*Currency, 0, len(all))
	for _, v := range all {
		currencies = append(currencies, v.Info())
	}
	return currencies
}

// CurrencyCodeByName - return CurrencyCode by currencyCode Alph name, case-insensitive, example: currencyUSD := CurrencyCodeByName("usd") OR currencyUSD := CurrencyCodeByName("USD")
func CurrencyCodeByName(name string) CurrencyCode {
	switch textPrepare(name) {
	case "AFN", "AFGHANI":
		return CurrencyAFN
	case "ALL", "LEK":
		return CurrencyALL
	case "DZD", "ALGERIANDINAR":
		return CurrencyDZD
	case "USD", "USDOLLAR":
		return CurrencyUSD
	case "EUR", "EURO":
		return CurrencyEUR
	case "AOA", "KWANZA":
		return CurrencyAOA
	case "XCD", "EASTCARIBBEANDOLLAR":
		return CurrencyXCD
	case "ARS", "ARGENTINEPESO":
		return CurrencyARS
	case "AMD", "ARMENIANDRAM":
		return CurrencyAMD
	case "AWG", "ARUBANFLORIN":
		return CurrencyAWG
	case "AUD", "AUSTRALIANDOLLAR":
		return CurrencyAUD
	case "AZN", "AZERBAIJANIANMANAT":
		return CurrencyAZN
	case "BSD", "BAHAMIANDOLLAR":
		return CurrencyBSD
	case "BHD", "BAHRAINIDINAR":
		return CurrencyBHD
	case "BDT", "TAKA":
		return CurrencyBDT
	case "BBD", "BARBADOSDOLLAR":
		return CurrencyBBD
	case "BYR", "BELARUSSIANRUBLE":
		return CurrencyBYR
	case "BZD", "BELIZEDOLLAR":
		return CurrencyBZD
	case "XOF", "CFAFRANCBCEAO":
		return CurrencyXOF
	case "BMD", "BERMUDIANDOLLAR":
		return CurrencyBMD
	case "BTN", "NGULTRUM":
		return CurrencyBTN
	case "INR", "INDIANRUPEE":
		return CurrencyINR
	case "BOB", "BOLIVIANO":
		return CurrencyBOB
	case "BAM", "CONVERTIBLEMARK":
		return CurrencyBAM
	case "BWP", "PULA":
		return CurrencyBWP
	case "NOK", "NORWEGIANKRONE":
		return CurrencyNOK
	case "BRL", "BRAZILIANREAL":
		return CurrencyBRL
	case "BND", "BRUNEIDOLLAR":
		return CurrencyBND
	case "BGN", "BULGARIANLEV":
		return CurrencyBGN
	case "BIF", "BURUNDIFRANC":
		return CurrencyBIF
	case "CVE", "CABOVERDEESCUDO":
		return CurrencyCVE
	case "KHR", "RIEL":
		return CurrencyKHR
	case "XAF", "CFAFRANCBEAC":
		return CurrencyXAF
	case "CAD", "CANADIANDOLLAR":
		return CurrencyCAD
	case "KYD", "CAYMANISLANDSDOLLAR":
		return CurrencyKYD
	case "CLF", "UNIDADDEFOMENTO":
		return CurrencyCLF
	case "CLP", "CHILEANPESO":
		return CurrencyCLP
	case "CNY", "YUANRENMINBI":
		return CurrencyCNY
	case "COP", "COLOMBIANPESO":
		return CurrencyCOP
	case "COU", "UNIDADDEVALORREAL":
		return CurrencyCOU
	case "KMF", "COMOROFRANC":
		return CurrencyKMF
	case "CDF", "CONGOLESEFRANC":
		return CurrencyCDF
	case "NZD", "NEWZEALANDDOLLAR":
		return CurrencyNZD
	case "CRC", "COSTARICANCOLON":
		return CurrencyCRC
	case "HRK", "KUNA":
		return CurrencyHRK
	case "CUC", "PESOCONVERTIBLE":
		return CurrencyCUC
	case "CUP", "CUBANPESO":
		return CurrencyCUP
	case "ANG", "NETHERLANDSANTILLEANGUILDER":
		return CurrencyANG
	case "CZK", "CZECHKORUNA":
		return CurrencyCZK
	case "DKK", "DANISHKRONE":
		return CurrencyDKK
	case "DJF", "DJIBOUTIFRANC":
		return CurrencyDJF
	case "DOP", "DOMINICANPESO":
		return CurrencyDOP
	case "EGP", "EGYPTIANPOUND":
		return CurrencyEGP
	case "SVC", "ELSALVADORCOLON":
		return CurrencySVC
	case "ERN", "NAKFA":
		return CurrencyERN
	case "ETB", "ETHIOPIANBIRR":
		return CurrencyETB
	case "FKP", "FALKLANDISLANDSPOUND":
		return CurrencyFKP
	case "FJD", "FIJIDOLLAR":
		return CurrencyFJD
	case "XPF", "CFPFRANC":
		return CurrencyXPF
	case "GMD", "DALASI":
		return CurrencyGMD
	case "GEL", "LARI":
		return CurrencyGEL
	case "GHS", "GHANACEDI":
		return CurrencyGHS
	case "GIP", "GIBRALTARPOUND":
		return CurrencyGIP
	case "GTQ", "QUETZAL":
		return CurrencyGTQ
	case "GBP", "POUNDSTERLING":
		return CurrencyGBP
	case "GNF", "GUINEAFRANC":
		return CurrencyGNF
	case "GYD", "GUYANADOLLAR":
		return CurrencyGYD
	case "HTG", "GOURDE":
		return CurrencyHTG
	case "HNL", "LEMPIRA":
		return CurrencyHNL
	case "HKD", "HONGKONGDOLLAR":
		return CurrencyHKD
	case "HUF", "FORINT":
		return CurrencyHUF
	case "ISK", "ICELANDKRONA":
		return CurrencyISK
	case "IDR", "RUPIAH":
		return CurrencyIDR
	case "XDR", "SDR(SPECIALDRAWINGRIGHT)":
		return CurrencyXDR
	case "IRR", "IRANIANRIAL":
		return CurrencyIRR
	case "IQD", "IRAQIDINAR":
		return CurrencyIQD
	case "ILS", "NEWISRAELISHEQEL":
		return CurrencyILS
	case "PALESTINE":
		return CurrencyILS
	case "JMD", "JAMAICANDOLLAR":
		return CurrencyJMD
	case "JPY", "YEN":
		return CurrencyJPY
	case "JOD", "JORDANIANDINAR":
		return CurrencyJOD
	case "KZT", "TENGE":
		return CurrencyKZT
	case "KES", "KENYANSHILLING":
		return CurrencyKES
	case "KPW", "NORTHKOREANWON":
		return CurrencyKPW
	case "KRW", "WON":
		return CurrencyKRW
	case "KWD", "KUWAITIDINAR":
		return CurrencyKWD
	case "KGS", "SOM":
		return CurrencyKGS
	case "LAK", "KIP":
		return CurrencyLAK
	case "LBP", "LEBANESEPOUND":
		return CurrencyLBP
	case "LSL", "LOTI":
		return CurrencyLSL
	case "ZAR", "RAND":
		return CurrencyZAR
	case "LRD", "LIBERIANDOLLAR":
		return CurrencyLRD
	case "LYD", "LIBYANDINAR":
		return CurrencyLYD
	case "CHF", "SWISSFRANC":
		return CurrencyCHF
	case "MOP", "PATACA":
		return CurrencyMOP
	case "MKD", "DENAR":
		return CurrencyMKD
	case "MGA", "MALAGASYARIARY":
		return CurrencyMGA
	case "MWK", "KWACHA":
		return CurrencyMWK
	case "MYR", "MALAYSIANRINGGIT":
		return CurrencyMYR
	case "MVR", "RUFIYAA":
		return CurrencyMVR
	case "MRU", "OUGUIYA":
		return CurrencyMRU
	case "MUR", "MAURITIUSRUPEE":
		return CurrencyMUR
	case "XUA", "ADBUNITOFACCOUNT":
		return CurrencyXUA
	case "MXN", "MEXICANPESO":
		return CurrencyMXN
	case "MXV", "MEXICANUNIDADDEINVERSION(UDI)":
		return CurrencyMXV
	case "MDL", "MOLDOVANLEU":
		return CurrencyMDL
	case "MNT", "TUGRIK":
		return CurrencyMNT
	case "MAD", "MOROCCANDIRHAM":
		return CurrencyMAD
	case "MZN", "MOZAMBIQUEMETICAL":
		return CurrencyMZN
	case "MMK", "KYAT":
		return CurrencyMMK
	case "NAD", "NAMIBIADOLLAR":
		return CurrencyNAD
	case "NPR", "NEPALESERUPEE":
		return CurrencyNPR
	case "NIO", "CORDOBAORO":
		return CurrencyNIO
	case "NGN", "NAIRA":
		return CurrencyNGN
	case "OMR", "RIALOMANI":
		return CurrencyOMR
	case "PKR", "PAKISTANRUPEE":
		return CurrencyPKR
	case "PAB", "BALBOA":
		return CurrencyPAB
	case "PGK", "KINA":
		return CurrencyPGK
	case "PYG", "GUARANI":
		return CurrencyPYG
	case "PEN", "NUEVOSOL":
		return CurrencyPEN
	case "PHP", "PHILIPPINEPESO":
		return CurrencyPHP
	case "PLN", "ZLOTY":
		return CurrencyPLN
	case "QAR", "QATARIRIAL":
		return CurrencyQAR
	case "RON", "ROMANIANLEU":
		return CurrencyRON
	case "RUB", "RUSSIANRUBLE":
		return CurrencyRUB
	case "RWF", "RWANDAFRANC":
		return CurrencyRWF
	case "SHP", "SAINTHELENAPOUND":
		return CurrencySHP
	case "WST", "TALA":
		return CurrencyWST
	case "STN", "DOBRA":
		return CurrencySTN
	case "SAR", "SAUDIRIYAL":
		return CurrencySAR
	case "RSD", "SERBIANDINAR":
		return CurrencyRSD
	case "SCR", "SEYCHELLESRUPEE":
		return CurrencySCR
	case "SLL", "LEONE":
		return CurrencySLL
	case "SGD", "SINGAPOREDOLLAR":
		return CurrencySGD
	case "XSU", "SUCRE":
		return CurrencyXSU
	case "SBD", "SOLOMONISLANDSDOLLAR":
		return CurrencySBD
	case "SOS", "SOMALISHILLING":
		return CurrencySOS
	case "SSP", "SOUTHSUDANESEPOUND":
		return CurrencySSP
	case "LKR", "SRILANKARUPEE":
		return CurrencyLKR
	case "SDG", "SUDANESEPOUND":
		return CurrencySDG
	case "SRD", "SURINAMDOLLAR":
		return CurrencySRD
	case "SZL", "LILANGENI":
		return CurrencySZL
	case "SEK", "SWEDISHKRONA":
		return CurrencySEK
	case "CHE", "WIREURO":
		return CurrencyCHE
	case "CHW", "WIRFRANC":
		return CurrencyCHW
	case "SYP", "SYRIANPOUND":
		return CurrencySYP
	case "TWD", "NEWTAIWANDOLLAR":
		return CurrencyTWD
	case "TJS", "SOMONI":
		return CurrencyTJS
	case "TZS", "TANZANIANSHILLING":
		return CurrencyTZS
	case "THB", "BAHT":
		return CurrencyTHB
	case "TOP", "PA’ANGA":
		return CurrencyTOP
	case "TTD", "TRINIDADANDTOBAGODOLLAR":
		return CurrencyTTD
	case "TND", "TUNISIANDINAR":
		return CurrencyTND
	case "TRY", "TURKISHLIRA":
		return CurrencyTRY
	case "TMT", "TURKMENISTANNEWMANAT":
		return CurrencyTMT
	case "UGX", "UGANDASHILLING":
		return CurrencyUGX
	case "UAH", "HRYVNIA":
		return CurrencyUAH
	case "AED", "UAEDIRHAM":
		return CurrencyAED
	case "USN", "USDOLLAR(NEXTDAY)":
		return CurrencyUSN
	case "UYI", "URUGUAYPESOENUNIDADESINDEXADAS(URUIURUI)":
		return CurrencyUYI
	case "UYU", "PESOURUGUAYO":
		return CurrencyUYU
	case "UZS", "UZBEKISTANSUM":
		return CurrencyUZS
	case "VUV", "VATU":
		return CurrencyVUV
	case "VEF", "BOLIVAR":
		return CurrencyVEF
	case "VND", "DONG":
		return CurrencyVND
	case "YER", "YEMENIRIAL":
		return CurrencyYER
	case "ZMW", "ZAMBIANKWACHA":
		return CurrencyZMW
	case "YUD", "YUGOSLAVIANDINAR":
		return CurrencyYUD
	case "ZWL", "ZIMBABWEDOLLAR":
		return CurrencyZWL
	}
	return CurrencyUnknown
}
