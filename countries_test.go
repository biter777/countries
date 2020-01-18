//nolint:gocyclo
package countries

import (
	"testing"
)

func TestCountriesCount(t *testing.T) {
	out := Total()
	want := len(All())
	if out != want {
		t.Errorf("Test All() err, want %v, got %v", want, out)
	}

	out = TotalCallCodes()
	want = len(AllCallCodes())
	if out != want {
		t.Errorf("Test AllCallCodes() err, want %v, got %v", want, out)
	}

	out = TotalCurrencies()
	want = len(AllCurrencies())
	if out != want {
		t.Errorf("Test AllCurrencies() err, want %v, got %v", want, out)
	}

	out = TotalDomains()
	want = len(AllDomains())
	if out != want {
		t.Errorf("Test AllDomains() err, want %v, got %v", want, out)
	}

	out = TotalRegions()
	want = len(AllRegions())
	if out != want {
		t.Errorf("Test AllRegions() err, want %v, got %v", want, out)
	}
}

func TestCountriesInfo(t *testing.T) {
	all := All()
	for i := 0; i < len(all); i++ {
		c := all[i].Info()
		if c == nil {
			t.Errorf("Test Info() err: c == nil, i: %v", i)
		}
		if c.Name == "" || c.Name == UnknownMsg {
			t.Errorf("Test c.Name err, c: %v", *c)
		}
		if c.Alpha2 == "" || c.Alpha2 == UnknownMsg {
			t.Errorf("Test c.Alpha2 err, c: %v", *c)
		}
		if c.Alpha3 == "" || c.Alpha3 == UnknownMsg {
			t.Errorf("Test c.Alpha3 err, c: %v", *c)
		}
		if c.IOC == "" || c.IOC == UnknownMsg {
			t.Errorf("Test c.IOC err, c: %v", *c)
		}
		if c.FIFA == "" || c.FIFA == UnknownMsg {
			t.Errorf("Test c.FIFA err, c: %v", *c)
		}
		if c.Emoji == "" || c.Emoji == UnknownMsg {
			t.Errorf("Test c.Emoji err, c: %v", *c)
		}
		if c.Code == Unknown {
			t.Errorf("Test c.Code err, c: %v", *c)
		}
		if c.Currency == CurrencyUnknown && c.Code != ATA {
			t.Errorf("Test c.Currency err, c: %v", *c)
		}
		if c.Capital == CapitalUnknown {
			t.Errorf("Test c.Capital err, c: %v", *c)
		}
		if len(c.CallCodes) < 1 {
			t.Errorf("Test c.CallCodes err, c: %v", *c)
		}
		if c.Domain == DomainUnknown && c.Code != XKX {
			t.Errorf("Test c.Domain err, c: %v", *c)
		}
		if c.Region == RegionUnknown {
			t.Errorf("Test c.Region err, c: %v", *c)
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
