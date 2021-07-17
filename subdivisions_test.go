package countries

import (
	"encoding/json"
	"testing"
)

// Test Subdivisions

//nolint:gocyclo
func TestSubdivisionsCount(t *testing.T) {
	got := TotalSubdivisions()

	want := len(AllSubdivisions())
	if got != want {
		t.Errorf("Test AllSubdivisions() err, want %v, got %v", want, got)
	}

	want = len(AllSubdivisionsInfo())
	if got != want {
		t.Errorf("Test AllSubdivisionsInfo() err, want %v, got %v", want, got)
	}
}

//nolint:gocyclo
func TestSubdivisionsCode(t *testing.T) {
	for _, s := range AllSubdivisions() {
		if s == SubdivisionUnknown {
			continue
		}

		if info := s.Info(); info.Country == Unknown {
			t.Errorf("Test info.Country err, s: %v", *info)
		}
	}
}

//nolint:gocyclo
func TestSubdivisionsIsValid(t *testing.T) {
	for _, s := range AllSubdivisions() {
		if s == SubdivisionUnknown {
			continue
		}

		if !s.IsValid() && s != SubdivisionUnknown {
			t.Errorf("Test SubdivisionCode.IsValid() err")
		}
	}
}

//nolint:gocyclo
func TestSubdivisionsString(t *testing.T) {
	for _, s := range AllSubdivisions() {
		if s == SubdivisionUnknown {
			continue
		}
		if s.String() == "" || s.String() == UnknownMsg {
			t.Errorf("Test SubdivisionCode.String() err, s: %v", s)
		}
	}
}

//nolint:gocyclo
func TestSubdivisionsType(t *testing.T) {
	for _, s := range AllSubdivisions() {
		if s.Type() != TypeSubdivisionCode {
			t.Errorf("Test SubdivisionCode.String() err, s: %v", s)
		}
	}
}

//nolint:gocyclo
func TestSubdivisionsCountry(t *testing.T) { //nolint:gocyclo
	for _, s := range AllSubdivisions() {
		if s == SubdivisionUnknown {
			continue
		}

		if s.Country() == Unknown {
			t.Errorf("Test SubdivisionCountry.String() err, s: %v", s)
		}
	}
}

//nolint:gocyclo
func TestSubdivisionsInfo(t *testing.T) {
	for _, s := range AllSubdivisions() {
		if s == SubdivisionUnknown {
			continue
		}

		if got := s.Info().Code; got != s {
			t.Errorf("Test AllSubdivisions() err, want %v, got %v", s, got)
		}
	}

	for _, c := range getAllCountries(t) {
		for _, s := range c.Subdivisions() {
			if got := s.Info().Code; got != s {
				t.Errorf("Test c.Subdivision[*].Info() err, want %v, got %v for c: %v", s, got, c)
			}
		}
	}
}

//nolint:gocyclo
func TestSubdivisionsInfoValue(t *testing.T) {
	for _, s := range AllSubdivisionsInfo() {
		if _, err := s.Value(); err != nil {
			t.Errorf("Test SubdivisionCode.String() Value err, s: %v", s)
		}

		if _, err := json.Marshal(s); err != nil {
			t.Errorf("Test SubdivisionCode.String() Marshal err, s: %v", s)
		}
	}
}

//nolint:gocyclo
func TestSubdivisionsInfoType(t *testing.T) {
	for _, s := range AllSubdivisionsInfo() {
		if out := s.Type(); out != TypeSubdivision {
			t.Errorf("Test AllSubdivisionsInfo.Type() err, want %v, got %v", TypeCallCodeInfo, out)
		}
	}
}

//nolint:gocyclo
func TestSubdivisionsInfoScan(t *testing.T) {
	for _, s := range AllSubdivisionsInfo() {
		err := s.Scan(s)
		if err != nil {
			t.Errorf("Test AllSubdivisionsInfo.Scan() err")
		}

		s2 := *s
		err = s.Scan(s2)
		if err != nil {
			t.Errorf("Test AllSubdivisionsInfo.Scan() err")
		}

		if s.Name != s2.Name || s.Code != s2.Code || s.Country != s2.Country {
			t.Errorf("Test AllSubdivisionsInfo.Scan() err")
		}

		err = s.Scan(nil)
		if err == nil {
			t.Errorf("Test AllSubdivisionsInfo.Scan() err")
		}

		s = nil
		err = s.Scan(s)
		if err == nil {
			t.Errorf("Test AllSubdivisionsInfo.Scan() err")
		}
	}
}

//nolint:gocyclo
func TestSubdivisionSubdivisionType(t *testing.T) { //nolint:gocyclo
	for _, s := range AllSubdivisions() {
		if s == SubdivisionUnknown {
			continue
		}

		if s.SubdivisionType() == SubdivisionTypeUnknown {
			t.Errorf("Test SubdivisionType.String() err, s: %v", s)
		}
	}
}

//nolint:gocyclo
func TestSubdivisionsByCountryCode(t *testing.T) {
	got := SubdivisionsByCountryCode(ByName("AU"))

	if count := len(got); count != 8 {
		t.Errorf("Australia has 8 subdivisions but got %d", count)
	}
}
