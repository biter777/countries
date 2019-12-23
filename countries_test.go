package countries

import (
	"testing"
)

func TestCount(t *testing.T) {
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
