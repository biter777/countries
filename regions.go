package countries

// RegionCode - Region code (UN M.49 code standart)
type RegionCode int64 // int64 for database/sql/driver.Valuer compatibility

// TypeRegionCode for Typer interface
const TypeRegionCode = "countries.RegionCode"

const (
	RegionUnknown RegionCode = 0
	RegionAF      RegionCode = 2
	RegionNA      RegionCode = 3
	RegionSA      RegionCode = 5
	RegionOC      RegionCode = 9
	RegionAN      RegionCode = 999
	RegionAS      RegionCode = 142
	RegionEU      RegionCode = 150
)

const (
	RegionAfrica       RegionCode = 2
	RegionNorthAmerica RegionCode = 3
	RegionSouthAmerica RegionCode = 5
	RegionOceania      RegionCode = 9
	RegionAntarctica   RegionCode = 999
	RegionAsia         RegionCode = 142
	RegionEurope       RegionCode = 150
)

// Type implements Typer interface
func (c RegionCode) Type() string {
	return TypeRegionCode
}

// String - implements fmt.Stringer, returns a Region name in english
func (c RegionCode) String() string {
	switch c {
	case RegionAF:
		return "Africa"
	case RegionNA:
		return "North America"
	case RegionOC:
		return "Oceania"
	case RegionAN:
		return "Antarctica"
	case RegionAS:
		return "Asia"
	case RegionEU:
		return "Europe"
	case RegionSA:
		return "South America"
	}
	return UnknownMsg
}

// String - implements fmt.Stringer, returns a Region name in russian
func (c RegionCode) StringRus() string {
	switch c {
	case RegionAF:
		return "Африка"
	case RegionNA:
		return "Северная Америка"
	case RegionOC:
		return "Океания"
	case RegionAN:
		return "Антарктика"
	case RegionAS:
		return "Азия"
	case RegionEU:
		return "Европа"
	case RegionSA:
		return "Южная Америка"
	}
	return UnknownMsg
}

// TotalRegions - returns number of Regions codes in the package
func TotalRegions() int {
	return 7
}

// AllRegions - returns all Regions
func AllRegions() []RegionCode {
	return []RegionCode{
		RegionAF,
		RegionNA,
		RegionOC,
		RegionAN,
		RegionAS,
		RegionEU,
		RegionSA,
	}
}

//
