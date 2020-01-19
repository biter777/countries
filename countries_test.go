//nolint:gocyclo
package countries

import (
	"encoding/json"
	"testing"
)

func getAllCountries() []CountryCode {
	allCountries := append(All(), Unknown, International, None)
	return append(allCountries, AllNonCountries()...)
}

//nolint:gocyclo
func TestCountriesByName(t *testing.T) {
	for _, c := range getAllCountries() {
		countryCodeOut := ByName(c.String())
		if countryCodeOut != c {
			t.Errorf("Test ByName() err, want %v, got %v", c, countryCodeOut)
		}
	}
}

//nolint:gocyclo
func TestCountriesByNumeric(t *testing.T) {
	for _, c := range getAllCountries() {
		countryCodeOut := ByNumeric(int(c))
		if countryCodeOut != c {
			t.Errorf("Test ByNumeric() err, want %v, got %v", c, countryCodeOut)
		}
	}
}

//nolint:gocyclo
func TestCountriesInfo(t *testing.T) {
	for _, c := range getAllCountries() {
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
		var c2 Country
		c2 = *c
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
		var c2 CallCodeInfo
		c2 = *c
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
func TestCurrenciesCode(t *testing.T) {
	for _, c := range AllCurrencies() {
		info := c.Info()
		if info == nil {
			t.Errorf("Test Info() err: c == nil")
		}
		if info.Name == "" || info.Name == UnknownMsg {
			t.Errorf("Test info.Name err, c: %v", *info)
		}
		if info.Alpha == "" || info.Alpha == UnknownMsg {
			t.Errorf("Test info.Alpha err, c: %v", *info)
		}
		if info.Code == CurrencyUnknown || info.Code != c {
			t.Errorf("Test info.Code err, c: %v", *info)
		}
		if len(info.Countries) < 1 {
			t.Errorf("Test info.Countries err, c: %v", *info)
		}
	}
}

//nolint:gocyclo
func TestCurrenciesIsValid(t *testing.T) {
	for _, c := range AllCurrencies() {
		if !c.IsValid() {
			t.Errorf("Test urrencyCode.IsValid() err")
		}
	}
}

//nolint:gocyclo
func TestCurrenciesString(t *testing.T) {
	for _, c := range AllCurrencies() {
		if c.String() == "" || c.String() == UnknownMsg {
			t.Errorf("Test urrencyCode.String() err")
		}
	}
}

//nolint:gocyclo
func TestCurrenciesType(t *testing.T) {
	for _, c := range AllCurrencies() {
		if c.Type() != TypeCurrencyCode {
			t.Errorf("Test urrencyCode.Type() err")
		}
	}
}

//nolint:gocyclo
func TestCurrenciesInfo(t *testing.T) {
	for _, c := range AllCurrencies() {
		info := c.Info()
		if info == nil {
			t.Errorf("Test Info() err: c == nil")
		}
		if info.Name == "" || info.Name == UnknownMsg {
			t.Errorf("Test info.Name err, c: %v", *info)
		}
		if info.Alpha == "" || info.Alpha == UnknownMsg {
			t.Errorf("Test info.Alpha err, c: %v", *info)
		}
		if info.Digits < 0 || info.Digits > 4 {
			t.Errorf("Test info.Digits err, c: %v", *info)
		}
		if info.Code == CurrencyUnknown || info.Code != c {
			t.Errorf("Test info.Code err, c: %v", *info)
		}
		if len(info.Countries) < 1 {
			t.Errorf("Test info.Countries err, c: %v", *info)
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
		var c2 Currency
		c2 = *c
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
		if info == nil {
			t.Errorf("Test Info() err: c == nil")
		}
		if info.Name == "" || info.Name == UnknownMsg {
			t.Errorf("Test info.Name err, c: %v", *info)
		}
		if info.Code == DomainUnknown || info.Code != c {
			t.Errorf("Test info.Code err, c: %v", *info)
		}
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
		if info == nil {
			t.Errorf("Test Info() err: c == nil")
		}
		if info.Name == "" || info.Name == UnknownMsg {
			t.Errorf("Test info.Name err, c: %v", *info)
		}
		if info.Country == Unknown {
			t.Errorf("Test info.Country err, c: %v", *info)
		}
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
		var c2 Domain
		c2 = *c
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
	for _, c := range getAllCountries() {
		out := DomainCodeByName(c.String())
		want := DomainCode(c)
		if out != want {
			t.Errorf("Test DomainCodeByName() err, want %v, got %v", want, out)
		}
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

//nolint:gocyclo
func TestRegionsIsValid(t *testing.T) {
	for _, c := range AllRegions() {
		if !c.IsValid() {
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
		var c2 Region
		c2 = *c
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
	for _, c := range getAllCountries() {
		out := c.Type()
		if out != TypeCountryCode {
			t.Errorf("Test All.Type() err, want %v, got %v", TypeCallCodeInfo, out)
		}
	}
}

//nolint:gocyclo
func TestCountriesCode(t *testing.T) {
	for _, c := range getAllCountries() {
		out2 := c.Info().Code
		if out2 != c {
			t.Errorf("Test All.Info().Code err, want %v, got %v", c, out2)
		}
	}
}

//nolint:gocyclo
func TestCountriesIsValid(t *testing.T) {
	for _, c := range getAllCountries() {
		if !c.IsValid() && c != Unknown {
			t.Errorf("Test All.IsValid() err")
		}
	}
}

//nolint:gocyclo
func TestCountriesAlpha(t *testing.T) {
	for _, c := range getAllCountries() {
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
	for _, c := range getAllCountries() {
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
	for _, c := range getAllCountries() {
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
	for _, c := range getAllCountries() {
		if (c.FIFA() == "" || c.FIFA() == UnknownMsg) && c != Unknown {
			t.Errorf("Test All.FIFA() err")
		}
	}
}

//nolint:gocyclo
func TestCountriesIOC(t *testing.T) {
	for _, c := range getAllCountries() {
		if (c.IOC() == "" || c.IOC() == UnknownMsg) && c != Unknown {
			t.Errorf("Test All.IOC() err")
		}
	}
}

//nolint:gocyclo
func TestCountriesCapital(t *testing.T) {
	for _, c := range getAllCountries() {
		if (c.Capital() == CapitalUnknown) && c != Unknown {
			t.Errorf("Test All.Capital() err")
		}
	}
}

//nolint:gocyclo
func TestCountriesCurrency(t *testing.T) {
	for _, c := range getAllCountries() {
		if c.Currency() == CurrencyUnknown && c != ATA && c != Unknown {
			t.Errorf("Test All.Currency() err")
		}
	}
}

//nolint:gocyclo
func TestCountriesDomain(t *testing.T) {
	for _, c := range getAllCountries() {
		if c.Domain() == DomainUnknown && c != Unknown {
			t.Errorf("Test All.Domain() err")
		}
	}
}

//nolint:gocyclo
func TestCountriesRegion(t *testing.T) {
	for _, c := range getAllCountries() {
		if c.Region() == RegionUnknown && c != Unknown {
			t.Errorf("Test All.Region() err")
		}
	}
}

//nolint:gocyclo
func TestCountriesCallCodes(t *testing.T) {
	for _, c := range getAllCountries() {
		if len(c.CallCodes()) < 1 && c != None {
			t.Errorf("Test All.CallCodes() err")
		}
	}
}

/*
//nolint:gocyclo
func TestCountriesAllInfo(t *testing.T) {
	all := getAllCountries()
	for i := 0; i < len(all); i++ {
		if !all[i].IsValid() && all[i] != Unknown {
			t.Errorf("Test Country.IsValid() err: c == nil, i: %v", i)
		}
		if all[i].Type() != TypeCountryCode {
			t.Errorf("Test Country.Type() err: all[%v]: %v", i, all[i])
		}
		if (all[i].String() == "" || all[i].String() == UnknownMsg) && all[i] != Unknown {
			t.Errorf("Test Country.String() err: all[%v]: %v", i, all[i])
		}
		if (all[i].StringRus() == "" || all[i].StringRus() == UnknownMsg) && all[i] != Unknown {
			all[i].StringRus()
			t.Errorf("Test Country.StringRus() err: all[%v]: %v", i, all[i])
		}
		if all[i].Emoji3() == "" {
			t.Errorf("Test Country.Emoji3() err: all[%v]: %v", i, all[i])
		}

		c := all[i].Info()
		if c == nil {
			t.Errorf("Test Info() err: c == nil, i: %v", i)
		}
		if (c.Name == "" || c.Name == UnknownMsg || c.Name != all[i].String()) && c.Code != Unknown {
			t.Errorf("Test c.Name err, c: %v", *c)
		}
		if (c.Alpha2 == "" || c.Alpha2 == UnknownMsg || c.Alpha2 != all[i].Alpha2()) && c.Code != Unknown {
			t.Errorf("Test c.Alpha2 err, c: %v", *c)
		}
		if (c.Alpha3 == "" || c.Alpha3 == UnknownMsg || c.Alpha3 != all[i].Alpha3()) && c.Code != Unknown {
			t.Errorf("Test c.Alpha3 err, c: %v", *c)
		}
		if (c.IOC == "" || c.IOC == UnknownMsg || c.IOC != all[i].IOC()) && c.Code != Unknown {
			t.Errorf("Test c.IOC err, c: %v", *c)
		}
		if (c.FIFA == "" || c.FIFA == UnknownMsg || c.FIFA != all[i].FIFA()) && c.Code != Unknown {
			t.Errorf("Test c.FIFA err, c: %v", *c)
		}
		if c.Emoji == "" || c.Emoji == UnknownMsg || c.Emoji != all[i].Emoji() {
			t.Errorf("Test c.Emoji err, c: %v", *c)
		}
		if c.Code != Unknown && c.Code != all[i] {
			t.Errorf("Test c.Code err, c: %v", *c)
		}
		if !c.Code.IsValid() && c.Code != Unknown {
			t.Errorf("Test c.Code.IsValid() err, c: %v", *c)
		}
		if c.Code.Type() != TypeCountryCode {
			t.Errorf("Test c.Code.Type() err, c: %v", *c)
		}
		if (c.Code.String() == "" || c.Code.String() == UnknownMsg) && c.Code != Unknown {
			t.Errorf("Test c.Code.String() err, c: %v", *c)
		}
		if (c.Currency == CurrencyUnknown || c.Currency != all[i].Currency()) && c.Code != ATA && c.Code != Unknown {
			t.Errorf("Test c.Currency err, c: %v", *c)
		}
		if (!c.Currency.IsValid() && c.Code != ATA) && c.Code != Unknown {
			t.Errorf("Test c.Currency.IsValid() err, c: %v", *c)
		}
		if c.Currency.Type() != TypeCurrencyCode {
			t.Errorf("Test c.Currency.Type() err, c: %v", *c)
		}
		if (c.Currency.String() == "" || c.Currency.String() == UnknownMsg) && c.Code != ATA && c.Code != Unknown {
			t.Errorf("Test c.Currency.String() err, c: %v", *c)
		}
		if c.Capital != CapitalUnknown && c.Capital != all[i].Capital() {
			t.Errorf("Test c.Capital err, c: %v", *c)
		}
		if c.Capital.Country() != all[i] && c.Code != None && c.Code < 999 {
			t.Errorf("Test c.Capital.Country() err, c: %v", *c)
		}
		if c.Capital.Info().Code != c.Capital {
			t.Errorf("Test c.Capital.Info() err, c: %v", *c)
		}
		if !c.Capital.IsValid() && c.Code != Unknown {
			t.Errorf("Test c.Capital.IsValid() err, c: %v", *c)
		}
		if c.Capital.Type() != TypeCapitalCode {
			t.Errorf("Test c.Capital.Type() err, c: %v", *c)
		}
		if (c.Capital.String() == "" || c.Capital.String() == UnknownMsg) && c.Code != Unknown {
			t.Errorf("Test c.Capital.String() err, c: %v", *c)
		}
		if (len(c.CallCodes) < 1 || len(c.CallCodes) != len(all[i].CallCodes())) && c.Code != None {
			t.Errorf("Test c.CallCodes err, c: %v", *c)
		}
		if len(c.CallCodes) > 0 {
			if c.CallCodes[0].Info().Code != c.CallCodes[0] {
				t.Errorf("Test c.CallCodes.Info() err, c: %v", *c)
			}
			if len(c.CallCodes[0].Countries()) < 1 {
				t.Errorf("Test c.CallCodes.Countries() err, c: %v", *c)
			}
			if !c.CallCodes[0].IsValid() && c.Code != ATA && c.Code != HMD && c.Code != Unknown {
				t.Errorf("Test c.CallCodes.IsValid() err, c: %v", *c)
			}
			if c.CallCodes[0].Type() != TypeCallCode {
				t.Errorf("Test c.CallCodes.Type() err, c: %v", *c)
			}
			if c.CallCodes[0].String() == "" || c.CallCodes[0].String() == UnknownMsg {
				t.Errorf("Test c.CallCodes[0].String() err, c: %v", *c)
			}
		}
		if c.Domain != DomainUnknown && c.Domain != all[i].Domain() && c.Code != XKX {
			t.Errorf("Test c.Domain err, c: %v", *c)
		}
		if !c.Domain.IsValid() && c.Code != Unknown {
			t.Errorf("Test c.Domain.IsValid() err, c: %v", *c)
		}
		if c.Domain.Type() != TypeDomainCode {
			t.Errorf("Test c.Domain.Type() err, c: %v", *c)
		}
		if (c.Domain.String() == "" || c.Domain.String() == UnknownMsg) && c.Code != Unknown && c.Code != None {
			t.Errorf("Test c.Domain.String() err, c: %v", *c)
		}
		if (c.Region == RegionUnknown || c.Region != all[i].Region()) && c.Code != Unknown && int(c.Code) < 999 {
			t.Errorf("Test c.Region err, c: %v", *c)
		}
		if !c.Region.IsValid() && c.Code != Unknown && int(c.Code) < 999 {
			t.Errorf("Test c.Region.IsValid() err, c: %v", *c)
		}
		if c.Region.Type() != TypeRegionCode {
			t.Errorf("Test c.Region.Type() err, c: %v", *c)
		}
		if (c.Region.String() == "" || c.Region.String() == UnknownMsg) && c.Code != Unknown && int(c.Code) < 999 {
			t.Errorf("Test c.Region.String() err, c: %v", *c)
		}
	}
}
*/

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
		if info == nil {
			t.Errorf("Test Info() err: c == nil")
		}
		if info.Name == "" || info.Name == UnknownMsg {
			t.Errorf("Test info.Name err, c: %v", *info)
		}
		if info.Code != CapitalUnknown && info.Code != c {
			t.Errorf("Test info.Code err, c: %v", *info)
		}
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
		var c2 Capital
		c2 = *c
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
