//nolint:gocyclo
package countries

import (
	"encoding/json"
	"testing"
)

func TestCountriesCount(t *testing.T) {
	out := Total()
	allCountries := All()
	want := len(allCountries)
	if out != want {
		t.Errorf("Test All() err, want %v, got %v", want, out)
	}
	allCountriesInfo := AllInfo()
	want = len(allCountriesInfo)
	if out != want {
		t.Errorf("Test AllInfo() err, want %v, got %v", want, out)
	}
	allNonCountries := AllNonCountries()
	if len(allNonCountries) < 10 {
		t.Errorf("Test allNonCountries() err")
	}
	allCountries = append(allCountries, Unknown, International, None)
	allCountries = append(allCountries, allNonCountries...)
	for _, c := range allCountries {
		countryCodeOut := ByName(c.String())
		countryCodeWant := c
		if countryCodeOut != countryCodeWant {
			t.Errorf("Test ByName() err, want %v, got %v", countryCodeWant, countryCodeOut)
		}
		countryCodeOut = ByNumeric(int(c))
		countryCodeWant = c
		if countryCodeOut != countryCodeWant {
			t.Errorf("Test ByNumeric() err, want %v, got %v", countryCodeWant, countryCodeOut)
		}
		countryCodeOut = c.Info().Code
		if countryCodeOut != countryCodeWant {
			t.Errorf("Test Info() err, want %v, got %v", countryCodeWant, countryCodeOut)
		}
	}
	for _, c := range allCountriesInfo {
		_, err := c.Value()
		_, err2 := json.Marshal(c)
		if err != nil || err2 != nil {
			t.Errorf("Test allCountriesInfo.Value() err")
		}
		if c.Type() != TypeCountry {
			t.Errorf("Test allCountriesInfo.Type() err")
		}
		err = c.Scan(c)
		if err != nil {
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

	out = TotalCallCodes()
	callCodes := AllCallCodes()
	want = len(callCodes)
	if out != want {
		t.Errorf("Test AllCallCodes() err, want %v, got %v", want, out)
	}
	callCodesInfo := AllCallCodesInfo()
	want = len(callCodesInfo)
	if out != want {
		t.Errorf("Test AllCallCodesInfo() err, want %v, got %v", want, out)
	}
	for _, c := range callCodes {
		if c.Info().Code != c {
			t.Errorf("Test callCodes.Info() err")
		}
	}
	_, err := callCodesInfo[0].Value()
	_, err2 := json.Marshal(callCodesInfo[0])
	if err != nil || err2 != nil {
		t.Errorf("Test callCodesInfo.Value() err")
	}
	if callCodesInfo[0].Type() != TypeCallCodeInfo {
		t.Errorf("Test callCodesInfo.Type() err")
	}
	err = callCodesInfo[0].Scan(callCodesInfo[0])
	if err != nil {
		t.Errorf("Test callCodesInfo.Scan() err")
	}
	err = callCodesInfo[0].Scan(nil)
	if err == nil {
		t.Errorf("Test callCodesInfo.Scan() err")
	}
	callCodesInfo[0] = nil
	err = callCodesInfo[0].Scan(callCodesInfo[0])
	if err == nil {
		t.Errorf("Test callCodesInfo.Scan() err")
	}

	out = TotalCurrencies()
	currencies := AllCurrencies()
	want = len(currencies)
	if out != want {
		t.Errorf("Test AllCurrencies() err, want %v, got %v", want, out)
	}
	for _, c := range currencies {
		currencyCodeOut := CurrencyCodeByName(c.String())
		currencyCodeWant := c
		if currencyCodeOut != currencyCodeWant {
			t.Errorf("Test CurrencyCodeByName() err, want %v, got %v", currencyCodeWant, currencyCodeOut)
		}
		if c.NickelRounding() != false && c != CurrencyCAD && c != CurrencyDKK && c != CurrencyCHF {
			t.Errorf("Test NickelRounding() err, want %v, got %v", false, c.NickelRounding())
		}
	}
	currenciesInfo := AllCurrenciesInfo()
	want = len(currenciesInfo)
	if out != want {
		t.Errorf("Test AllCurrenciesInfo() err, want %v, got %v", want, out)
	}
	_, err = currenciesInfo[0].Value()
	_, err2 = json.Marshal(currenciesInfo[0])
	if err != nil || err2 != nil {
		t.Errorf("Test currenciesInfo.Value() err")
	}
	if currenciesInfo[0].Type() != TypeCurrency {
		t.Errorf("Test currenciesInfo.Type() err")
	}
	err = currenciesInfo[0].Scan(currenciesInfo[0])
	if err != nil {
		t.Errorf("Test currenciesInfo.Scan() err")
	}
	err = currenciesInfo[0].Scan(nil)
	if err == nil {
		t.Errorf("Test currenciesInfo.Scan() err")
	}
	currenciesInfo[0] = nil
	err = currenciesInfo[0].Scan(currenciesInfo[0])
	if err == nil {
		t.Errorf("Test currenciesInfo.Scan() err")
	}

	out = TotalDomains()
	domains := AllDomains()
	want = len(domains)
	if out != want {
		t.Errorf("Test AllDomains() err, want %v, got %v", want, out)
	}
	for _, c := range allCountries {
		domainCodeOut := DomainCodeByName(c.String())
		domainCodeWant := DomainCode(c)
		if domainCodeOut != domainCodeWant {
			t.Errorf("Test DomainCodeByName() err, want %v, got %v", domainCodeWant, domainCodeOut)
		}
	}
	domainsInfo := AllDomainsInfo()
	want = len(domainsInfo)
	if out != want {
		t.Errorf("Test AllDomainsInfo() err, want %v, got %v", want, out)
	}
	_, err = domainsInfo[0].Value()
	_, err2 = json.Marshal(domainsInfo[0])
	if err != nil || err2 != nil {
		t.Errorf("Test domainsInfo.Value() err")
	}
	if domainsInfo[0].Type() != TypeDomain {
		t.Errorf("Test domainsInfo.Type() err")
	}
	err = domainsInfo[0].Scan(domainsInfo[0])
	if err != nil {
		t.Errorf("Test domainsInfo.Scan() err")
	}
	err = domainsInfo[0].Scan(nil)
	if err == nil {
		t.Errorf("Test domainsInfo.Scan() err")
	}
	domainsInfo[0] = nil
	err = domainsInfo[0].Scan(domainsInfo[0])
	if err == nil {
		t.Errorf("Test domainsInfo.Scan() err")
	}

	out = TotalRegions()
	regions := AllRegions()
	want = len(regions)
	if out != want {
		t.Errorf("Test AllRegions() err, want %v, got %v", want, out)
	}
	for _, r := range regions {
		if !r.IsValid() {
			t.Errorf("Test r.IsValid() err")
		}
		regionCodeOut := RegionCodeByName(r.String())
		regionCodeWant := r
		if regionCodeOut != regionCodeWant {
			t.Errorf("Test RegionCodeByName() err, want %v, got %v", regionCodeWant, regionCodeOut)
		}
	}
	regionsInfo := AllRegionsInfo()
	want = len(regionsInfo)
	if out != want {
		t.Errorf("Test AllRegionsInfo() err, want %v, got %v", want, out)
	}
	_, err = regionsInfo[0].Value()
	_, err2 = json.Marshal(regionsInfo[0])
	if err != nil || err2 != nil {
		t.Errorf("Test regionsInfo.Value() err")
	}
	if regionsInfo[0].Type() != TypeRegion {
		t.Errorf("Test regionsInfo.Type() err")
	}
	err = regionsInfo[0].Scan(regionsInfo[0])
	if err != nil {
		t.Errorf("Test regionsInfo.Scan() err")
	}
	err = regionsInfo[0].Scan(nil)
	if err == nil {
		t.Errorf("Test regionsInfo.Scan() err")
	}
	regionsInfo[0] = nil
	err = regionsInfo[0].Scan(regionsInfo[0])
	if err == nil {
		t.Errorf("Test regionsInfo.Scan() err")
	}

	out = TotalCapitals()
	capitals := AllCapitals()
	want = len(capitals)
	if out != want {
		t.Errorf("Test AllCapitals() err, want %v, got %v", want, out)
	}
	capitalsInfo := AllCapitalsInfo()
	want = len(capitalsInfo)
	if out != want {
		t.Errorf("Test AllCapitalsInfo() err, want %v, got %v", want, out)
	}
	for _, c := range capitals {
		capitalCodeOut := CapitalCodeByName(c.String())
		capitalCodeWant := c
		if capitalCodeOut != capitalCodeWant && c != CapitalYU && c != CapitalAQ && c != CapitalBV && c != CapitalHM &&
			c != CapitalUM && c != CapitalTK && c != CapitalBQ {
			t.Errorf("Test CapitalCodeByName() err, want %v, got %v", capitalCodeWant, capitalCodeOut)
		}
	}
	_, err = capitalsInfo[0].Value()
	_, err2 = json.Marshal(capitals[0])
	if err != nil || err2 != nil {
		t.Errorf("Test capitalsInfo.Value() err")
	}
	if capitalsInfo[0].Type() != TypeCapital {
		t.Errorf("Test capitalsInfo.Type() err")
	}
	err = capitalsInfo[0].Scan(capitalsInfo[0])
	if err != nil {
		t.Errorf("Test capitalsInfo.Scan() err")
	}
	err = capitalsInfo[0].Scan(nil)
	if err == nil {
		t.Errorf("Test capitalsInfo.Scan() err")
	}
	capitalsInfo[0] = nil
	err = capitalsInfo[0].Scan(capitalsInfo[0])
	if err == nil {
		t.Errorf("Test capitalsInfo.Scan() err")
	}
}

func TestTypes(t *testing.T) {
	allCapitals := AllCapitals()
	typeOut := allCapitals[0].Type()
	typeWant := TypeCapitalCode
	if typeOut != typeWant {
		t.Errorf("Test capitals type err, want %v, got %v", typeWant, typeOut)
	}

	allRegions := AllRegions()
	typeOut = allRegions[0].Type()
	typeWant = TypeRegionCode
	if typeOut != typeWant {
		t.Errorf("Test regions type err, want %v, got %v", typeWant, typeOut)
	}
}

func TestCountriesInfo(t *testing.T) {
	all := All()
	allNonCountries := AllNonCountries()
	if len(allNonCountries) < 10 {
		t.Errorf("Test allNonCountries() err")
	}
	all = append(all, Unknown, International, None)
	all = append(all, allNonCountries...)

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

//nolint:gocyclo
func TestCurrenciesInfo(t *testing.T) { //nolint:gocyclo
	all := AllCurrencies()
	for i := 0; i < len(all); i++ {
		c := all[i].Info()
		if c == nil {
			t.Errorf("Test Info() err: c == nil, i: %v", i)
		}
		if c.Name == "" || c.Name == UnknownMsg {
			t.Errorf("Test c.Name err, c: %v", *c)
		}
		if c.Alpha == "" || c.Alpha == UnknownMsg {
			t.Errorf("Test c.Alpha err, c: %v", *c)
		}
		if c.Digits < 0 || c.Digits > 4 {
			t.Errorf("Test c.Digits err, c: %v", *c)
		}
		if c.Code == CurrencyUnknown {
			t.Errorf("Test c.Code err, c: %v", *c)
		}
		if len(c.Countries) < 1 {
			t.Errorf("Test c.Countries err, c: %v", *c)
		}
	}
}

func TestCapitalsInfo(t *testing.T) {
	all := AllCapitals()
	for i := 0; i < len(all); i++ {
		c := all[i].Info()
		if c == nil {
			t.Errorf("Test Info() err: c == nil, i: %v", i)
		}
		if c.Name == "" || c.Name == UnknownMsg {
			t.Errorf("Test c.Name err, c: %v", *c)
		}
		if c.Country == Unknown {
			t.Errorf("Test c.Country err, c: %v", *c)
		}
		if c.Code == CapitalUnknown {
			t.Errorf("Test c.Code err, c: %v", *c)
		}
	}
}

func TestDomainsInfo(t *testing.T) {
	all := AllDomains()
	for i := 0; i < len(all); i++ {
		c := all[i].Info()
		if c == nil {
			t.Errorf("Test Info() err: c == nil, i: %v", i)
		}
		if c.Name == "" || c.Name == UnknownMsg {
			t.Errorf("Test c.Name err, c: %v", *c)
		}
		if c.Country == Unknown {
			t.Errorf("Test c.Country err, c: %v", *c)
		}
		if c.Code == DomainUnknown {
			t.Errorf("Test c.Code err, c: %v", *c)
		}
	}
}
