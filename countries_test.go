//nolint:gocyclo
package countries

import (
	"encoding/json"
	"testing"
)

func getAllCountries(t *testing.T) []CountryCode {
	t.Helper()

	allCountries := append(All(), Unknown, International, None)
	return append(allCountries, AllNonCountries()...)
}

//nolint:gocyclo
func TestCountriesByName(t *testing.T) {
	for _, c := range getAllCountries(t) {
		countryCodeOut := ByName(c.String())
		if countryCodeOut != c {
			t.Errorf("Test ByName() err, want %v, got %v", c, countryCodeOut)
		}
	}
	countryCodeOut := ByName("ENGLAND")
	if countryCodeOut != GBR {
		t.Errorf("Test ByName() err, want %v, got %v", GBR, countryCodeOut)
	}
	countryCodeOut = ByName("WALES")
	if countryCodeOut != XWA {
		t.Errorf("Test ByName() err, want %v, got %v", XWA, countryCodeOut)
	}
	countryCodeOut = ByName("SCOTLAND")
	if countryCodeOut != XSC {
		t.Errorf("Test ByName() err, want %v, got %v", XSC, countryCodeOut)
	}
	countryCodeOut = ByName("BRAZIL")
	if countryCodeOut != BRA {
		t.Errorf("Test ByName() err, want %v, got %v", BRA, countryCodeOut)
	}
	countryCodeOut = ByName("pupok")
	if countryCodeOut != Unknown {
		t.Errorf("Test ByName() err, want %v, got %v", Unknown, countryCodeOut)
	}
	countryCodeOut = ByName(None.String())
	if countryCodeOut != None {
		t.Errorf("Test ByName() err, want %v, got %v", None, countryCodeOut)
	}
}

//nolint:gocyclo
func TestCountriesByNumeric(t *testing.T) {
	for _, c := range getAllCountries(t) {
		countryCodeOut := ByNumeric(int(c))
		if countryCodeOut != c {
			t.Errorf("Test ByNumeric() err, want %v, got %v", c, countryCodeOut)
		}
	}
	countryCodeOut := ByNumeric(100500999)
	if countryCodeOut != Unknown {
		t.Errorf("Test ByNumeric() err, want %v, got %v", Unknown, countryCodeOut)
	}

	countryCodeBRA := ByNumeric(986)
	if countryCodeBRA != Brazil {
		t.Errorf("Test ByNumeric() err, want %v, got %v", Brazil, countryCodeBRA)
	}
}

//nolint:gocyclo
func TestCountriesInfo(t *testing.T) {
	for _, c := range getAllCountries(t) {
		countryCodeOut := c.Info().Code
		if countryCodeOut != c {
			t.Errorf("Test Info() err, want %v, got %v", c, countryCodeOut)
		}
	}
}

//nolint:gocyclo
func TestCountriesInfoValue(t *testing.T) {
	for _, c := range AllInfo() {
		_, err := c.Value()
		_, err2 := json.Marshal(c)
		if err != nil || err2 != nil {
			t.Errorf("Test allCountriesInfo.Value() err")
		}
	}
}

//nolint:gocyclo
func TestCountriesInfoType(t *testing.T) {
	for _, c := range AllInfo() {
		out := c.Type()
		if out != TypeCountry {
			t.Errorf("Test allCountriesInfo.Type() err, want %v, got %v", TypeCountry, out)
		}
	}
}

//nolint:gocyclo
func TestCountriesInfoScan(t *testing.T) {
	for _, c := range AllInfo() {
		err := c.Scan(c)
		if err != nil {
			t.Errorf("Test allCountriesInfo.Scan() err")
		}
		c2 := *c
		err = c.Scan(c2)
		if err != nil {
			t.Errorf("Test allCountriesInfo.Scan() err")
		}
		if c.Name != c2.Name || c.Alpha2 != c2.Alpha2 || c.Alpha3 != c2.Alpha3 || c.IOC != c2.IOC || c.Code != c2.Code {
			t.Errorf("Test allCountriesInfo.Scan() err")
		}
		err = c.Scan(nil)
		if err == nil {
			t.Errorf("Test allCountriesInfo.Scan() err")
		}
		c = nil
		err = c.Scan(c)
		if err == nil {
			t.Errorf("Test allCountriesInfo.Scan() err")
		}
	}
}

//nolint:gocyclo
func TestCountriesCountInfo(t *testing.T) {
	out := Total()
	want := len(AllInfo())
	if out != want {
		t.Errorf("Test AllInfo() err, want %v, got %v", want, out)
	}
}

// Test CallCodes

//nolint:gocyclo
func TestCallCodesCount(t *testing.T) {
	out := TotalCallCodes()
	want := len(AllCallCodes())
	if out != want {
		t.Errorf("Test AllCallCodes() err, want %v, got %v", want, out)
	}
	want = len(AllCallCodesInfo())
	if out != want {
		t.Errorf("Test AllCallCodesInfo() err, want %v, got %v", want, out)
	}
}

//nolint:gocyclo
func TestCallCodesCode(t *testing.T) {
	for _, c := range AllCallCodes() {
		info := c.Info()
		if info == nil {
			t.Errorf("Test Info() err: c == nil")
			continue
		}
		if info.Code != CallCodeUnknown && info.Code != c {
			t.Errorf("Test info.Code err, c: %v", *info)
		}
		if len(info.Countries) < 1 {
			t.Errorf("Test info.Countries err, c: %v", *info)
		}
	}
}

//nolint:gocyclo
func TestCallCodesIsValid(t *testing.T) {
	for _, c := range AllCallCodes() {
		if !c.IsValid() && c != CallCodeUnknown {
			t.Errorf("Test CallCodes.IsValid() err")
		}
	}
	for _, c := range getAllCountries(t) {
		callCodes := c.Info().CallCodes
		if len(callCodes) > 0 && !callCodes[0].IsValid() && c != Unknown {
			t.Errorf("Test CallCodes.IsValid() err")
		}
	}
}

//nolint:gocyclo
func TestCallCodesString(t *testing.T) {
	for _, c := range AllCallCodes() {
		if c.String() == "" || c.String() == UnknownMsg {
			t.Errorf("Test CallCodes.String() err")
		}
	}
}

//nolint:gocyclo
func TestCallCodesType(t *testing.T) {
	for _, c := range AllCallCodes() {
		if c.Type() != TypeCallCode {
			t.Errorf("Test CallCodes.Type() err")
		}
	}
}

//nolint:gocyclo
func TestCallCodes(t *testing.T) {
	for _, c := range AllCallCodes() {
		out := c.Type()
		if out != TypeCallCode {
			t.Errorf("Test AllCallCodes.Type() err, want %v, got %v", TypeCallCodeInfo, out)
		}
		out2 := c.Info().Code
		if out2 != c {
			t.Errorf("Test AllCallCodes.Info().Code err, want %v, got %v", c, out2)
		}
		if len(c.Countries()) < 1 {
			t.Errorf("Test AllCallCodes.Countries() err")
		}
		if !c.IsValid() && c != CallCodeUnknown {
			t.Errorf("Test AllCallCodes.IsValid() err")
		}
		if c.String() == "" {
			t.Errorf("Test AllCallCodes.String() err")
		}
	}
}

//nolint:gocyclo
func TestCallCodesInfo(t *testing.T) {
	for _, c := range AllCallCodes() {
		if c.Info().Code != c {
			t.Errorf("Test AllCallCodes() err, want %v, got %v", c, c.Info().Code)
		}
	}
}

//nolint:gocyclo
func TestCallCodesInfoValue(t *testing.T) {
	for _, c := range AllCallCodesInfo() {
		_, err := c.Value()
		_, err2 := json.Marshal(c)
		if err != nil || err2 != nil {
			t.Errorf("Test allCallCodesInfo.Value() err")
		}
	}
}

//nolint:gocyclo
func TestCallCodesInfoType(t *testing.T) {
	for _, c := range AllCallCodesInfo() {
		out := c.Type()
		if out != TypeCallCodeInfo {
			t.Errorf("Test allCallCodesInfo.Type() err, want %v, got %v", TypeCallCodeInfo, out)
		}
	}
}

//nolint:gocyclo
func TestCallCodesInfoScan(t *testing.T) {
	for _, c := range AllCallCodesInfo() {
		err := c.Scan(c)
		if err != nil {
			t.Errorf("Test allCallCodesInfo.Scan() err")
		}
		c2 := *c
		err = c.Scan(c2)
		if err != nil {
			t.Errorf("Test allCallCodesInfo.Scan() err")
		}
		if len(c.Countries) != len(c2.Countries) || c.Code != c2.Code {
			t.Errorf("Test allCallCodesInfo.Scan() err")
		}
		if len(c.Countries) > 0 && c.Countries[0] != c2.Countries[0] {
			t.Errorf("Test allCallCodesInfo.Scan() err")
		}
		if len(c.Countries) > 1 && c.Countries[1] != c2.Countries[1] {
			t.Errorf("Test allCallCodesInfo.Scan() err")
		}
		if len(c.Countries) > 2 && c.Countries[2] != c2.Countries[2] {
			t.Errorf("Test allCallCodesInfo.Scan() err")
		}
		err = c.Scan(nil)
		if err == nil {
			t.Errorf("Test allCallCodesInfo.Scan() err")
		}
		c = nil
		err = c.Scan(c)
		if err == nil {
			t.Errorf("Test allCallCodesInfo.Scan() err")
		}
	}
}

// Test Currencies
func TestCurrenciesCount(t *testing.T) {
	out := TotalCurrencies()
	want := len(AllCurrencies())
	if out != want {
		t.Errorf("Test AllCurrencies() err, want %v, got %v", want, out)
	}
	want = len(AllCurrenciesInfo())
	if out != want {
		t.Errorf("Test AllCurrenciesInfo() err, want %v, got %v", want, out)
	}
}

//nolint:gocyclo
func TestCurrenciesIsValid(t *testing.T) {
	for _, c := range AllCurrencies() {
		if !c.IsValid() {
			t.Errorf("Test CurrencyCode.IsValid() err")
		}
	}
	for _, c := range getAllCountries(t) {
		if !c.Info().Currency.IsValid() && c != ATA && c != Unknown {
			t.Errorf("Test CurrencyCode.IsValid() err")
		}
	}
}

//nolint:gocyclo
func TestCurrenciesString(t *testing.T) {
	for _, c := range AllCurrencies() {
		if c.String() == "" || c.String() == UnknownMsg {
			t.Errorf("Test CurrencyCode.String() err")
		}
	}
	for _, c := range getAllCountries(t) {
		currency := c.Info().Currency
		if (currency.String() == "" || currency.String() == UnknownMsg) && c != ATA && c != Unknown {
			t.Errorf("Test CurrencyCode.String() err, c: %v", c)
		}
	}
}

//nolint:gocyclo
func TestCurrenciesType(t *testing.T) {
	for _, c := range AllCurrencies() {
		if c.Type() != TypeCurrencyCode {
			t.Errorf("Test CurrencyCode.Type() err")
		}
	}
}

//nolint:gocyclo
func TestCurrenciesAlpha(t *testing.T) {
	for _, c := range AllCurrencies() {
		if c.Alpha() == "" || c.Alpha() == UnknownMsg {
			t.Errorf("Test CurrencyCode.Alpha() err")
		}
	}
}

//nolint:gocyclo
func TestCurrenciesCountries(t *testing.T) {
	for _, c := range AllCurrencies() {
		if len(c.Countries()) < 1 {
			t.Errorf("Test CurrencyCode.Countries() err")
		}
	}
	if len(CurrencyNone.Countries()) != 1 || CurrencyNone.Countries()[0] != None {
		t.Errorf("Test CurrencyCode.Countries() err")
	}
}

//nolint:gocyclo
func TestCurrenciesDigits(t *testing.T) {
	for _, c := range AllCurrencies() {
		digits := c.Digits()
		if digits != 0 && digits != 2 && digits != 3 && digits != 4 {
			t.Errorf("Test CurrencyCode.Digits() err")
		}
	}
	digits := CurrencyNone.Digits()
	if digits != 0 {
		t.Errorf("Test CurrencyCode.Digits() err")
	}
	unknown := CurrencyCode(100500999)
	digits = unknown.Digits()
	if digits != -1 {
		t.Errorf("Test CurrencyCode.Digits() err")
	}
}

//nolint:gocyclo
func TestCurrenciesEmoji(t *testing.T) {
	for _, c := range AllCurrencies() {
		emoji := c.Emoji()
		if emoji == "" || emoji == UnknownMsg {
			t.Errorf("Test CurrencyCode.Emoji() err")
		}
	}
}

//nolint:gocyclo
func TestCurrenciesNickelRounding(t *testing.T) {
	for _, c := range AllCurrencies() {
		nr := c.NickelRounding()
		if nr != false && c != CurrencyCAD && c != CurrencyDKK && c != CurrencyCHF {
			t.Errorf("Test CurrencyCode.NickelRounding() err")
		}
		if nr != true && (c == CurrencyCAD || c == CurrencyDKK || c == CurrencyCHF) {
			t.Errorf("Test CurrencyCode.NickelRounding() err")
		}
	}
}

//nolint:gocyclo
func TestCurrenciesInfoValue(t *testing.T) {
	for _, c := range AllCurrenciesInfo() {
		_, err := c.Value()
		_, err2 := json.Marshal(c)
		if err != nil || err2 != nil {
			t.Errorf("Test allCurrenciesInfo.Value() err")
		}
	}
}

//nolint:gocyclo
func TestCurrenciesInfoType(t *testing.T) {
	for _, c := range AllCurrenciesInfo() {
		out := c.Type()
		if out != TypeCurrency {
			t.Errorf("Test allCallCodesInfo.Type() err, want %v, got %v", TypeCurrency, out)
		}
	}
}

//nolint:gocyclo
func TestCurrenciesInfoScan(t *testing.T) {
	for _, c := range AllCurrenciesInfo() {
		err := c.Scan(c)
		if err != nil {
			t.Errorf("Test allCurrenciesInfo.Scan() err")
		}
		c2 := *c
		err = c.Scan(c2)
		if err != nil {
			t.Errorf("Test allCurrenciesInfo.Scan() err")
		}
		if len(c.Countries) != len(c2.Countries) || c.Code != c2.Code {
			t.Errorf("Test allCurrenciesInfo.Scan() err")
		}
		if len(c.Countries) > 0 && c.Countries[0] != c2.Countries[0] {
			t.Errorf("Test allCurrenciesInfo.Scan() err")
		}
		if len(c.Countries) > 1 && c.Countries[1] != c2.Countries[1] {
			t.Errorf("Test allCurrenciesInfo.Scan() err")
		}
		if len(c.Countries) > 2 && c.Countries[2] != c2.Countries[2] {
			t.Errorf("Test allCurrenciesInfo.Scan() err")
		}
		err = c.Scan(nil)
		if err == nil {
			t.Errorf("Test allCurrenciesInfo.Scan() err")
		}
		c = nil
		err = c.Scan(c)
		if err == nil {
			t.Errorf("Test allCurrenciesInfo.Scan() err")
		}
	}
}

//nolint:gocyclo
func TestCurrencyCodeByName(t *testing.T) {
	for _, c := range AllCurrencies() {
		out := CurrencyCodeByName(c.Alpha())
		if out != c {
			t.Errorf("Test CurrencyCodeByName() err, want %v, got %v", c, out)
		}
	}
	out := CurrencyCodeByName("pupok")
	if out != CurrencyUnknown {
		t.Errorf("Test CurrencyCodeByName() err, want %v, got %v", CurrencyUnknown, out)
	}
	out = CurrencyCodeByName(None.Alpha2())
	if out != CurrencyNON {
		t.Errorf("Test CurrencyCodeByName() err, want %v, got %v", CurrencyNON, out)
	}
	if CurrencyCodeByName(Palestine.String()) != CurrencyILS {
		t.Errorf("Test CurrencyCodeByName() err, want %v, got %v", CurrencyILS, out)
	}
}

// Test Domains
func TestDomainsCount(t *testing.T) {
	out := TotalDomains()
	want := len(AllDomains())
	if out != want {
		t.Errorf("Test AllDomains() err, want %v, got %v", want, out)
	}
	want = len(AllDomainsInfo())
	if out != want {
		t.Errorf("Test AllDomains() err, want %v, got %v", want, out)
	}
}

//nolint:gocyclo
func TestDomainsCode(t *testing.T) {
	for _, c := range AllDomains() {
		info := c.Info()
		if info.Country == Unknown {
			t.Errorf("Test info.Code err, c: %v", *info)
		}
	}
}

//nolint:gocyclo
func TestDomainsIsValid(t *testing.T) {
	for _, c := range AllDomains() {
		if !c.IsValid() {
			t.Errorf("Test DomainCode.IsValid() err")
		}
	}
	for _, c := range getAllCountries(t) {
		if !c.Info().Domain.IsValid() && c != Unknown && int(c) < 999 {
			t.Errorf("Test DomainCode.IsValid() err")
		}
	}
}

//nolint:gocyclo
func TestDomainsCountry(t *testing.T) {
	for _, c := range AllDomains() {
		if c.Country() != CountryCode(c) {
			t.Errorf("Test DomainCode.Country() err")
		}
	}
	unknown := DomainCode(990)
	if unknown.Country() != Unknown {
		t.Errorf("Test DomainCode.Country() err")
	}
}

//nolint:gocyclo
func TestDomainsString(t *testing.T) {
	for _, c := range AllDomains() {
		if c.String() == "" || c.String() == UnknownMsg {
			t.Errorf("Test DomainCode.String() err")
		}
	}
}

//nolint:gocyclo
func TestDomainsType(t *testing.T) {
	for _, c := range AllDomains() {
		if c.Type() != TypeDomainCode {
			t.Errorf("Test DomainCode.Type() err")
		}
	}
}

//nolint:gocyclo
func TestDomainsInfo(t *testing.T) {
	for _, c := range AllDomains() {
		info := c.Info()
		if info.Code == DomainUnknown || info.Code != c {
			t.Errorf("Test info.Code err, c: %v", *info)
		}
	}
}

//nolint:gocyclo
func TestDomainsInfoValue(t *testing.T) {
	for _, c := range AllDomainsInfo() {
		_, err := c.Value()
		_, err2 := json.Marshal(c)
		if err != nil || err2 != nil {
			t.Errorf("Test allDomainsInfo.Value() err")
		}
	}
}

//nolint:gocyclo
func TestDomainsInfoType(t *testing.T) {
	for _, c := range AllDomainsInfo() {
		out := c.Type()
		if out != TypeDomain {
			t.Errorf("Test allDomainsInfo.Type() err, want %v, got %v", TypeDomain, out)
		}
	}
}

//nolint:gocyclo
func TestDomainsInfoScan(t *testing.T) {
	for _, c := range AllDomainsInfo() {
		err := c.Scan(c)
		if err != nil {
			t.Errorf("Test allDomainsInfo.Scan() err")
		}
		c2 := *c
		err = c.Scan(c2)
		if err != nil {
			t.Errorf("Test allDomainsInfo.Scan() err")
		}
		if c.Name != c2.Name || c.Code != c2.Code || c.Country != c2.Country {
			t.Errorf("Test allDomainsInfo.Scan() err")
		}
		err = c.Scan(nil)
		if err == nil {
			t.Errorf("Test allDomainsInfo.Scan() err")
		}
		c = nil
		err = c.Scan(c)
		if err == nil {
			t.Errorf("Test allDomainsInfo.Scan() err")
		}
	}
}

//nolint:gocyclo
func TestDomainCodeByName(t *testing.T) {
	for _, c := range getAllCountries(t) {
		out := DomainCodeByName(c.String())
		want := DomainCode(c)
		if out != want {
			t.Errorf("Test DomainCodeByName() err, want %v, got %v", want, out)
		}
	}
	out := DomainCodeByName("pupok")
	if out != DomainUnknown {
		t.Errorf("Test DomainCodeByName() err, want %v, got %v", DomainUnknown, out)
	}
	out = DomainCodeByName(None.String())
	if out != DomainXX {
		t.Errorf("Test DomainCodeByName() err, want %v, got %v", DomainXX, out)
	}
}

// Test Regions

//nolint:gocyclo
func TestRegionsCount(t *testing.T) {
	out := TotalRegions()
	want := len(AllRegions())
	if out != want {
		t.Errorf("Test AllRegions() err, want %v, got %v", want, out)
	}
	want = len(AllRegionsInfo())
	if out != want {
		t.Errorf("Test AllRegionsInfo() err, want %v, got %v", want, out)
	}
}

/*
//nolint:gocyclo
func TestRegionsCode(t *testing.T) {
	for _, c := range AllRegions() {
		info := c.Info()
		if info == nil {
			t.Errorf("Test Info() err: c == nil")
		}
		if info.Name == "" || info.Name == UnknownMsg {
			t.Errorf("Test info.Name err, c: %v", *info)
		}
		if info.Code == RegionUnknown || info.Code != c {
			t.Errorf("Test info.Code err, c: %v", *info)
		}
	}
}
*/

//nolint:gocyclo
func TestRegionsIsValid(t *testing.T) {
	for _, c := range AllRegions() {
		if !c.IsValid() {
			t.Errorf("Test RegionCode.IsValid() err")
		}
	}
	for _, c := range getAllCountries(t) {
		if !c.Info().Region.IsValid() && c != Unknown && int(c) < 999 {
			t.Errorf("Test RegionCode.IsValid() err")
		}
	}
}

//nolint:gocyclo
func TestRegionsString(t *testing.T) {
	for _, c := range AllRegions() {
		if c.String() == "" || c.String() == UnknownMsg {
			t.Errorf("Test RegionCode.String() err")
		}
		if c.StringRus() == "" || c.StringRus() == UnknownMsg {
			t.Errorf("Test RegionCode.StringRus() err")
		}
	}
	unknown := RegionCode(100500999)
	if unknown.String() != UnknownMsg {
		t.Errorf("Test RegionCode.String() err")
	}
	if unknown.StringRus() != UnknownMsg {
		t.Errorf("Test RegionCode.StringRus() err")
	}
	if RegionNone.String() == "" || RegionNone.String() == UnknownMsg {
		t.Errorf("Test RegionCode.String() err")
	}
	if RegionNone.StringRus() == "" || RegionNone.StringRus() == UnknownMsg {
		t.Errorf("Test RegionCode.StringRus() err")
	}
}

//nolint:gocyclo
func TestRegionsType(t *testing.T) {
	for _, c := range AllRegions() {
		if c.Type() != TypeRegionCode {
			t.Errorf("Test RegionCode.Type() err")
		}
	}
}

//nolint:gocyclo
func TestRegionsInfoValue(t *testing.T) {
	for _, c := range AllRegionsInfo() {
		_, err := c.Value()
		_, err2 := json.Marshal(c)
		if err != nil || err2 != nil {
			t.Errorf("Test AllRegionsInfo.Value() err")
		}
	}
}

//nolint:gocyclo
func TestRegionsInfoType(t *testing.T) {
	for _, c := range AllRegionsInfo() {
		out := c.Type()
		if out != TypeRegion {
			t.Errorf("Test AllRegionsInfo.Type() err, want %v, got %v", TypeRegion, out)
		}
	}
}

//nolint:gocyclo
func TestRegionsInfoScan(t *testing.T) {
	for _, c := range AllRegionsInfo() {
		err := c.Scan(c)
		if err != nil {
			t.Errorf("Test AllRegionsInfo.Scan() err")
		}
		c2 := *c
		err = c.Scan(c2)
		if err != nil {
			t.Errorf("Test AllRegionsInfo.Scan() err")
		}
		if c.Name != c2.Name || c.Code != c2.Code {
			t.Errorf("Test AllRegionsInfo.Scan() err")
		}
		err = c.Scan(nil)
		if err == nil {
			t.Errorf("Test AllRegionsInfo.Scan() err")
		}
		c = nil
		err = c.Scan(c)
		if err == nil {
			t.Errorf("Test AllRegionsInfo.Scan() err")
		}
	}
}

//nolint:gocyclo
func TestRegionCodeByName(t *testing.T) {
	for _, c := range AllRegions() {
		out := RegionCodeByName(c.String())
		want := RegionCode(c)
		if out != want {
			t.Errorf("Test RegionCodeByName() err, want %v, got %v", want, out)
		}
	}
	regionCodeOut := RegionCodeByName(None.String())
	regionCodeWant := None.Region()
	if regionCodeOut != regionCodeWant {
		t.Errorf("Test RegionCodeByName() err, want %v, got %v", regionCodeWant, regionCodeOut)
	}
	regionCodeOut = RegionCodeByName("pupok")
	regionCodeWant = RegionUnknown
	if regionCodeOut != regionCodeWant {
		t.Errorf("Test RegionCodeByName() err, want %v, got %v", regionCodeWant, regionCodeOut)
	}
}

//nolint:gocyclo
func TestCapitalCodeByName(t *testing.T) {
	for _, c := range AllCapitals() {
		capitalCodeOut := CapitalCodeByName(c.String())
		if capitalCodeOut != c && c != CapitalYU && c != CapitalAQ && c != CapitalBV && c != CapitalHM &&
			c != CapitalUM && c != CapitalTK && c != CapitalBQ {
			t.Errorf("Test CapitalCodeByName() err, want %v, got %v", c, capitalCodeOut)
		}
	}
	capitalCodeOut := CapitalCodeByName("pupok")
	if capitalCodeOut != CapitalUnknown {
		t.Errorf("Test CapitalCodeByName() err, want %v, got %v", CapitalUnknown, capitalCodeOut)
	}
	capitalCodeOut = CapitalCodeByName(None.Alpha2())
	if capitalCodeOut != CapitalXX {
		t.Errorf("Test CapitalCodeByName() err, want %v, got %v", CapitalXX, capitalCodeOut)
	}
}

// Test Countries

//nolint:gocyclo
func TestCountriesCount(t *testing.T) {
	out := Total()
	want := len(All())
	if out != want {
		t.Errorf("Test All() err, want %v, got %v", want, out)
	}
	if len(AllNonCountries()) < 10 {
		t.Errorf("Test allNonCountries() err")
	}
}

//nolint:gocyclo
func TestCountriesType(t *testing.T) {
	for _, c := range getAllCountries(t) {
		out := c.Type()
		if out != TypeCountryCode {
			t.Errorf("Test All.Type() err, want %v, got %v", TypeCallCodeInfo, out)
		}
	}
}

//nolint:gocyclo
func TestCountriesCode(t *testing.T) {
	for _, c := range getAllCountries(t) {
		out2 := c.Info().Code
		if out2 != c {
			t.Errorf("Test All.Info().Code err, want %v, got %v", c, out2)
		}
	}
}

//nolint:gocyclo
func TestCountriesIsValid(t *testing.T) {
	for _, c := range getAllCountries(t) {
		if !c.IsValid() && c != Unknown {
			t.Errorf("Test All.IsValid() err")
		}
	}
}

//nolint:gocyclo
func TestCountriesAlpha(t *testing.T) {
	for _, c := range getAllCountries(t) {
		if c.Alpha2() == "" || c.Alpha2() == UnknownMsg && c != Unknown {
			t.Errorf("Test All.Alpha2() err")
		}
		if c.Alpha3() == "" || c.Alpha3() == UnknownMsg && c != Unknown {
			t.Errorf("Test All.Alpha2() err")
		}
	}
}

//nolint:gocyclo
func TestCountriesString(t *testing.T) {
	for _, c := range getAllCountries(t) {
		if (c.String() == "" || c.String() == UnknownMsg) && c != Unknown {
			t.Errorf("Test All.String() err")
		}
		if (c.StringRus() == "" || c.StringRus() == UnknownMsg) && c != Unknown {
			t.Errorf("Test All.StringRus() err")
		}
	}
}

//nolint:gocyclo
func TestCountriesEmoji(t *testing.T) {
	for _, c := range getAllCountries(t) {
		if c.Emoji() == "" || c.Emoji() == UnknownMsg {
			t.Errorf("Test All.Emoji() err")
		}
		if c.Emoji3() == "" || c.Emoji3() == UnknownMsg {
			t.Errorf("Test All.Emoji3() err")
		}
	}
}

//nolint:gocyclo
func TestCountriesFIFA(t *testing.T) {
	for _, c := range getAllCountries(t) {
		if (c.FIFA() == "" || c.FIFA() == UnknownMsg) && c != Unknown {
			t.Errorf("Test All.FIFA() err")
		}
	}
}

//nolint:gocyclo
func TestCountriesIOC(t *testing.T) {
	for _, c := range getAllCountries(t) {
		if (c.IOC() == "" || c.IOC() == UnknownMsg) && c != Unknown {
			t.Errorf("Test All.IOC() err")
		}
	}
}

//nolint:gocyclo
func TestCountriesCapital(t *testing.T) {
	for _, c := range getAllCountries(t) {
		if (c.Capital() == CapitalUnknown) && c != Unknown {
			t.Errorf("Test All.Capital() err")
		}
	}
}

//nolint:gocyclo
func TestCountriesCurrency(t *testing.T) {
	for _, c := range getAllCountries(t) {
		if c.Currency() == CurrencyUnknown && c != ATA && c != Unknown {
			t.Errorf("Test All.Currency() err")
		}
	}
	if ChannelIslands.Currency() != CurrencyEUR {
		t.Errorf("Test All.Currency() err")
	}
}

//nolint:gocyclo
func TestCountriesDomain(t *testing.T) {
	for _, c := range getAllCountries(t) {
		if c.Domain() == DomainUnknown && c != Unknown {
			t.Errorf("Test All.Domain() err")
		}
	}
}

//nolint:gocyclo
func TestCountriesRegion(t *testing.T) {
	for _, c := range getAllCountries(t) {
		if c.Region() == RegionUnknown && c != Unknown {
			t.Errorf("Test All.Region() err")
		}
	}
}

//nolint:gocyclo
func TestCountriesCallCodes(t *testing.T) {
	for _, c := range getAllCountries(t) {
		if len(c.CallCodes()) < 1 && c != None {
			t.Errorf("Test All.CallCodes() err")
		}
	}
}

// Test Capitals

//nolint:gocyclo
func TestCapitalsCount(t *testing.T) {
	out := TotalCapitals()
	want := len(AllCapitals())
	if out != want {
		t.Errorf("Test AllCapitals() err, want %v, got %v", want, out)
	}
	want = len(AllCapitalsInfo())
	if out != want {
		t.Errorf("Test AllCapitalsInfo() err, want %v, got %v", want, out)
	}
}

//nolint:gocyclo
func TestCapitalsCode(t *testing.T) {
	for _, c := range AllCapitals() {
		info := c.Info()
		if info.Country == Unknown {
			t.Errorf("Test info.Country err, c: %v", *info)
		}
	}
}

//nolint:gocyclo
func TestCapitalsIsValid(t *testing.T) {
	for _, c := range AllCapitals() {
		if !c.IsValid() && c != CapitalUnknown {
			t.Errorf("Test CapitalCode.IsValid() err")
		}
	}
}

//nolint:gocyclo
func TestCapitalsString(t *testing.T) {
	for _, c := range AllCapitals() {
		if c.String() == "" || c.String() == UnknownMsg {
			t.Errorf("Test CapitalCode.String() err")
		}
	}
}

//nolint:gocyclo
func TestCapitalsType(t *testing.T) {
	for _, c := range AllCapitals() {
		if c.Type() != TypeCapitalCode {
			t.Errorf("Test CapitalsCode.Type() err")
		}
	}
}

//nolint:gocyclo
func TestCapitalsCountry(t *testing.T) { //nolint:gocyclo
	for _, c := range AllCapitals() {
		if c.Country() == Unknown {
			t.Errorf("Test CapitalsCode.Country() err")
		}
	}
}

//nolint:gocyclo
func TestCapitalsInfo(t *testing.T) {
	for _, c := range AllCapitals() {
		if c.Info().Code != c {
			t.Errorf("Test AllCapitals() err, want %v, got %v", c, c.Info().Code)
		}
	}
	for _, c := range getAllCountries(t) {
		if c.Info().Capital.Info().Code != c.Info().Capital {
			t.Errorf("Test c.Capital.Info() err, c: %v", c)
		}
	}
}

//nolint:gocyclo
func TestCapitalsInfoValue(t *testing.T) {
	for _, c := range AllCapitalsInfo() {
		_, err := c.Value()
		_, err2 := json.Marshal(c)
		if err != nil || err2 != nil {
			t.Errorf("Test AllCapitalsInfo.Value() err")
		}
	}
}

//nolint:gocyclo
func TestCapitalsInfoType(t *testing.T) {
	for _, c := range AllCapitalsInfo() {
		out := c.Type()
		if out != TypeCapital {
			t.Errorf("Test AllCapitalsInfo.Type() err, want %v, got %v", TypeCallCodeInfo, out)
		}
	}
}

//nolint:gocyclo
func TestCapitalsInfoScan(t *testing.T) {
	for _, c := range AllCapitalsInfo() {
		err := c.Scan(c)
		if err != nil {
			t.Errorf("Test AllCapitalsInfo.Scan() err")
		}
		c2 := *c
		err = c.Scan(c2)
		if err != nil {
			t.Errorf("Test AllCapitalsInfo.Scan() err")
		}
		if c.Name != c2.Name || c.Code != c2.Code || c.Country != c2.Country {
			t.Errorf("Test AllCapitalsInfo.Scan() err")
		}
		err = c.Scan(nil)
		if err == nil {
			t.Errorf("Test AllCapitalsInfo.Scan() err")
		}
		c = nil
		err = c.Scan(c)
		if err == nil {
			t.Errorf("Test AllCapitalsInfo.Scan() err")
		}
	}
}
