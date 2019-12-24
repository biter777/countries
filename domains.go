package countries

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

// DomainCode - domain code
type DomainCode int64 // int64 for database/sql/driver.Valuer compatibility

// Domain - capital info
type Domain struct {
	Name    string
	Code    DomainCode
	Country CountryCode
}

// TypeDomain for Typer interface
const TypeDomain = "countries.Domain"

// TypeDomainCode for Typer interface
const TypeDomainCode = "countries.DomainCode"

// Domains codes
const (
	DomainUnknown DomainCode = 0
	DomainArpa    DomainCode = DomainCode(Internation)
	DomainCom     DomainCode = DomainCode(Internation + 1)
	DomainOrg     DomainCode = DomainCode(Internation + 2)
	DomainNet     DomainCode = DomainCode(Internation + 3)
	DomainEdu     DomainCode = DomainCode(Internation + 4)
	DomainGov     DomainCode = DomainCode(Internation + 5)
	DomainMil     DomainCode = DomainCode(Internation + 6)
	DomainTest    DomainCode = DomainCode(Internation + 7)
	DomainBiz     DomainCode = DomainCode(Internation + 8)
	DomainInfo    DomainCode = DomainCode(Internation + 9)
	DomainName    DomainCode = DomainCode(Internation + 10)
	DomainAU      DomainCode = DomainCode(AU)
	DomainAT      DomainCode = DomainCode(AT)
	DomainAZ      DomainCode = DomainCode(AZ)
	DomainAL      DomainCode = DomainCode(AL)
	DomainDZ      DomainCode = DomainCode(DZ)
	DomainAS      DomainCode = DomainCode(AS)
	DomainAI      DomainCode = DomainCode(AI)
	DomainAO      DomainCode = DomainCode(AO)
	DomainAD      DomainCode = DomainCode(AD)
	DomainAQ      DomainCode = DomainCode(AQ)
	DomainAG      DomainCode = DomainCode(AG)
	DomainAN      DomainCode = DomainCode(AN)
	DomainAE      DomainCode = DomainCode(AE)
	DomainAR      DomainCode = DomainCode(AR)
	DomainAM      DomainCode = DomainCode(AM)
	DomainAW      DomainCode = DomainCode(AW)
	DomainAF      DomainCode = DomainCode(AF)
	DomainBS      DomainCode = DomainCode(BS)
	DomainBD      DomainCode = DomainCode(BD)
	DomainBB      DomainCode = DomainCode(BB)
	DomainBH      DomainCode = DomainCode(BH)
	DomainBY      DomainCode = DomainCode(BY)
	DomainBZ      DomainCode = DomainCode(BZ)
	DomainBE      DomainCode = DomainCode(BE)
	DomainBJ      DomainCode = DomainCode(BJ)
	DomainBM      DomainCode = DomainCode(BM)
	DomainBG      DomainCode = DomainCode(BG)
	DomainBO      DomainCode = DomainCode(BO)
	DomainBA      DomainCode = DomainCode(BA)
	DomainBW      DomainCode = DomainCode(BW)
	DomainBR      DomainCode = DomainCode(BR)
	DomainIO      DomainCode = DomainCode(IO)
	DomainBN      DomainCode = DomainCode(BN)
	DomainBF      DomainCode = DomainCode(BF)
	DomainBI      DomainCode = DomainCode(BI)
	DomainBT      DomainCode = DomainCode(BT)
	DomainVU      DomainCode = DomainCode(VU)
	DomainVA      DomainCode = DomainCode(VA)
	DomainGB      DomainCode = DomainCode(GB)
	DomainHU      DomainCode = DomainCode(HU)
	DomainVE      DomainCode = DomainCode(VE)
	DomainVG      DomainCode = DomainCode(VG)
	DomainVI      DomainCode = DomainCode(VI)
	DomainTL      DomainCode = DomainCode(TL)
	DomainVN      DomainCode = DomainCode(VN)
	DomainGA      DomainCode = DomainCode(GA)
	DomainHT      DomainCode = DomainCode(HT)
	DomainGY      DomainCode = DomainCode(GY)
	DomainGM      DomainCode = DomainCode(GM)
	DomainGH      DomainCode = DomainCode(GH)
	DomainGP      DomainCode = DomainCode(GP)
	DomainGT      DomainCode = DomainCode(GT)
	DomainGN      DomainCode = DomainCode(GN)
	DomainGW      DomainCode = DomainCode(GW)
	DomainDE      DomainCode = DomainCode(DE)
	DomainGI      DomainCode = DomainCode(GI)
	DomainHN      DomainCode = DomainCode(HN)
	DomainHK      DomainCode = DomainCode(HK)
	DomainGD      DomainCode = DomainCode(GD)
	DomainGL      DomainCode = DomainCode(GL)
	DomainGR      DomainCode = DomainCode(GR)
	DomainGE      DomainCode = DomainCode(GE)
	DomainGU      DomainCode = DomainCode(GU)
	DomainDK      DomainCode = DomainCode(DK)
	DomainCD      DomainCode = DomainCode(CD)
	DomainDJ      DomainCode = DomainCode(DJ)
	DomainDM      DomainCode = DomainCode(DM)
	DomainDO      DomainCode = DomainCode(DO)
	DomainEG      DomainCode = DomainCode(EG)
	DomainZM      DomainCode = DomainCode(ZM)
	DomainEH      DomainCode = DomainCode(EH)
	DomainZW      DomainCode = DomainCode(ZW)
	DomainIL      DomainCode = DomainCode(IL)
	DomainIN      DomainCode = DomainCode(IN)
	DomainID      DomainCode = DomainCode(ID)
	DomainJO      DomainCode = DomainCode(JO)
	DomainIQ      DomainCode = DomainCode(IQ)
	DomainIR      DomainCode = DomainCode(IR)
	DomainIE      DomainCode = DomainCode(IE)
	DomainIS      DomainCode = DomainCode(IS)
	DomainES      DomainCode = DomainCode(ES)
	DomainIT      DomainCode = DomainCode(IT)
	DomainYE      DomainCode = DomainCode(YE)
	DomainKZ      DomainCode = DomainCode(KZ)
	DomainKY      DomainCode = DomainCode(KY)
	DomainKH      DomainCode = DomainCode(KH)
	DomainCM      DomainCode = DomainCode(CM)
	DomainCA      DomainCode = DomainCode(CA)
	DomainQA      DomainCode = DomainCode(QA)
	DomainKE      DomainCode = DomainCode(KE)
	DomainCY      DomainCode = DomainCode(CY)
	DomainKI      DomainCode = DomainCode(KI)
	DomainCN      DomainCode = DomainCode(CN)
	DomainCC      DomainCode = DomainCode(CC)
	DomainCO      DomainCode = DomainCode(CO)
	DomainKM      DomainCode = DomainCode(KM)
	DomainCG      DomainCode = DomainCode(CG)
	DomainKP      DomainCode = DomainCode(KP)
	DomainKR      DomainCode = DomainCode(KR)
	DomainCR      DomainCode = DomainCode(CR)
	DomainCI      DomainCode = DomainCode(CI)
	DomainCU      DomainCode = DomainCode(CU)
	DomainKW      DomainCode = DomainCode(KW)
	DomainKG      DomainCode = DomainCode(KG)
	DomainLA      DomainCode = DomainCode(LA)
	DomainLV      DomainCode = DomainCode(LV)
	DomainLS      DomainCode = DomainCode(LS)
	DomainLR      DomainCode = DomainCode(LR)
	DomainLB      DomainCode = DomainCode(LB)
	DomainLY      DomainCode = DomainCode(LY)
	DomainLT      DomainCode = DomainCode(LT)
	DomainLI      DomainCode = DomainCode(LI)
	DomainLU      DomainCode = DomainCode(LU)
	DomainMU      DomainCode = DomainCode(MU)
	DomainMR      DomainCode = DomainCode(MR)
	DomainMG      DomainCode = DomainCode(MG)
	DomainYT      DomainCode = DomainCode(YT)
	DomainMO      DomainCode = DomainCode(MO)
	DomainMK      DomainCode = DomainCode(MK)
	DomainMW      DomainCode = DomainCode(MW)
	DomainMY      DomainCode = DomainCode(MY)
	DomainML      DomainCode = DomainCode(ML)
	DomainMV      DomainCode = DomainCode(MV)
	DomainMT      DomainCode = DomainCode(MT)
	DomainMP      DomainCode = DomainCode(MP)
	DomainMA      DomainCode = DomainCode(MA)
	DomainMQ      DomainCode = DomainCode(MQ)
	DomainMH      DomainCode = DomainCode(MH)
	DomainMX      DomainCode = DomainCode(MX)
	DomainFM      DomainCode = DomainCode(FM)
	DomainMZ      DomainCode = DomainCode(MZ)
	DomainMD      DomainCode = DomainCode(MD)
	DomainMC      DomainCode = DomainCode(MC)
	DomainMN      DomainCode = DomainCode(MN)
	DomainMS      DomainCode = DomainCode(MS)
	DomainMM      DomainCode = DomainCode(MM)
	DomainNA      DomainCode = DomainCode(NA)
	DomainNR      DomainCode = DomainCode(NR)
	DomainNP      DomainCode = DomainCode(NP)
	DomainNE      DomainCode = DomainCode(NE)
	DomainNG      DomainCode = DomainCode(NG)
	DomainNL      DomainCode = DomainCode(NL)
	DomainNI      DomainCode = DomainCode(NI)
	DomainNU      DomainCode = DomainCode(NU)
	DomainNZ      DomainCode = DomainCode(NZ)
	DomainNC      DomainCode = DomainCode(NC)
	DomainNO      DomainCode = DomainCode(NO)
	DomainOM      DomainCode = DomainCode(OM)
	DomainBV      DomainCode = DomainCode(BV)
	DomainIM      DomainCode = DomainCode(IM)
	DomainNF      DomainCode = DomainCode(NF)
	DomainPN      DomainCode = DomainCode(PN)
	DomainCX      DomainCode = DomainCode(CX)
	DomainSH      DomainCode = DomainCode(SH)
	DomainWF      DomainCode = DomainCode(WF)
	DomainHM      DomainCode = DomainCode(HM)
	DomainCV      DomainCode = DomainCode(CV)
	DomainCK      DomainCode = DomainCode(CK)
	DomainWS      DomainCode = DomainCode(WS)
	DomainSJ      DomainCode = DomainCode(SJ)
	DomainTC      DomainCode = DomainCode(TC)
	DomainUM      DomainCode = DomainCode(UM)
	DomainPK      DomainCode = DomainCode(PK)
	DomainPW      DomainCode = DomainCode(PW)
	DomainPS      DomainCode = DomainCode(PS)
	DomainPA      DomainCode = DomainCode(PA)
	DomainPG      DomainCode = DomainCode(PG)
	DomainPY      DomainCode = DomainCode(PY)
	DomainPE      DomainCode = DomainCode(PE)
	DomainPL      DomainCode = DomainCode(PL)
	DomainPT      DomainCode = DomainCode(PT)
	DomainPR      DomainCode = DomainCode(PR)
	DomainRE      DomainCode = DomainCode(RE)
	DomainRU      DomainCode = DomainCode(RU)
	DomainRW      DomainCode = DomainCode(RW)
	DomainRO      DomainCode = DomainCode(RO)
	DomainSV      DomainCode = DomainCode(SV)
	DomainSM      DomainCode = DomainCode(SM)
	DomainST      DomainCode = DomainCode(ST)
	DomainSA      DomainCode = DomainCode(SA)
	DomainSZ      DomainCode = DomainCode(SZ)
	DomainSC      DomainCode = DomainCode(SC)
	DomainSN      DomainCode = DomainCode(SN)
	DomainPM      DomainCode = DomainCode(PM)
	DomainVC      DomainCode = DomainCode(VC)
	DomainKN      DomainCode = DomainCode(KN)
	DomainLC      DomainCode = DomainCode(LC)
	DomainSG      DomainCode = DomainCode(SG)
	DomainSY      DomainCode = DomainCode(SY)
	DomainSK      DomainCode = DomainCode(SK)
	DomainSI      DomainCode = DomainCode(SI)
	DomainUS      DomainCode = DomainCode(US)
	DomainSB      DomainCode = DomainCode(SB)
	DomainSO      DomainCode = DomainCode(SO)
	DomainSD      DomainCode = DomainCode(SD)
	DomainSR      DomainCode = DomainCode(SR)
	DomainSL      DomainCode = DomainCode(SL)
	DomainTJ      DomainCode = DomainCode(TJ)
	DomainTW      DomainCode = DomainCode(TW)
	DomainTH      DomainCode = DomainCode(TH)
	DomainTZ      DomainCode = DomainCode(TZ)
	DomainTG      DomainCode = DomainCode(TG)
	DomainTK      DomainCode = DomainCode(TK)
	DomainTO      DomainCode = DomainCode(TO)
	DomainTT      DomainCode = DomainCode(TT)
	DomainTV      DomainCode = DomainCode(TV)
	DomainTN      DomainCode = DomainCode(TN)
	DomainTM      DomainCode = DomainCode(TM)
	DomainTR      DomainCode = DomainCode(TR)
	DomainUG      DomainCode = DomainCode(UG)
	DomainUZ      DomainCode = DomainCode(UZ)
	DomainUA      DomainCode = DomainCode(UA)
	DomainUY      DomainCode = DomainCode(UY)
	DomainFO      DomainCode = DomainCode(FO)
	DomainFJ      DomainCode = DomainCode(FJ)
	DomainPH      DomainCode = DomainCode(PH)
	DomainFI      DomainCode = DomainCode(FI)
	DomainFK      DomainCode = DomainCode(FK)
	DomainFR      DomainCode = DomainCode(FR)
	DomainGF      DomainCode = DomainCode(GF)
	DomainPF      DomainCode = DomainCode(PF)
	DomainTF      DomainCode = DomainCode(TF)
	DomainHR      DomainCode = DomainCode(HR)
	DomainCF      DomainCode = DomainCode(CF)
	DomainTD      DomainCode = DomainCode(TD)
	DomainCZ      DomainCode = DomainCode(CZ)
	DomainCL      DomainCode = DomainCode(CL)
	DomainCH      DomainCode = DomainCode(CH)
	DomainSE      DomainCode = DomainCode(SE)
	DomainXS      DomainCode = DomainCode(XS)
	DomainLK      DomainCode = DomainCode(LK)
	DomainEC      DomainCode = DomainCode(EC)
	DomainGQ      DomainCode = DomainCode(GQ)
	DomainER      DomainCode = DomainCode(ER)
	DomainEE      DomainCode = DomainCode(EE)
	DomainET      DomainCode = DomainCode(ET)
	DomainZA      DomainCode = DomainCode(ZA)
	DomainYU      DomainCode = DomainCode(YU)
	DomainGS      DomainCode = DomainCode(GS)
	DomainJM      DomainCode = DomainCode(JM)
	DomainME      DomainCode = DomainCode(ME)
	DomainBL      DomainCode = DomainCode(BL)
	DomainSX      DomainCode = DomainCode(SX)
	DomainRS      DomainCode = DomainCode(RS)
	DomainAX      DomainCode = DomainCode(AX)
	DomainBQ      DomainCode = DomainCode(BQ)
	DomainGG      DomainCode = DomainCode(GG)
	DomainJE      DomainCode = DomainCode(JE)
	DomainCW      DomainCode = DomainCode(CW)
	DomainMF      DomainCode = DomainCode(MF)
	DomainSS      DomainCode = DomainCode(SS)
	DomainJP      DomainCode = DomainCode(JP)
)

// Type implements Typer interface
func (c DomainCode) Type() string {
	return TypeDomainCode
}

// String - implements fmt.Stringer, returns a domain (internet ccTDL)
func (c DomainCode) String() string {
	switch c {
	case 	DomainArpa:
		return ".arpa"
	case DomainCom:
		return ".com"
	case DomainOrg:
		return ".org"
	case DomainNet:
		return ".net"
	case DomainEdu:
		return ".edu"
	case DomainGov:
		return ".gov"
	case DomainMil:
		return ".mil"
	case DomainTest:
		return ".test"
	case DomainBiz:
		return ".biz"
	case DomainInfo:
		return ".info"
	case DomainName:
		return ".name"
	case DomainBV, DomainSJ:
		return ".no"
	case DomainGB:
		return ".uk"
	}

	if c >= 999 {
		c = DomainCode(999)
	}

	a2 := CountryCode(c).Alpha2()
	if a2 == UnknownMsg {
		return UnknownMsg
	}

	return "." + strings.ToLower(a2)
}

// IsValid - returns true, if code is correct
func (c DomainCode) IsValid() bool {
	return c.String() != UnknownMsg
}

// Country - returns a country of domain
func (c DomainCode) Country() CountryCode {
	if !c.IsValid() {
		return Unknown
	}
	return CountryCode(c)
}

// Info - returns domain information as Domain
func (c DomainCode) Info() *Domain {
	return &Domain{
		Name:    c.String(),
		Code:    c,
		Country: c.Country(),
	}
}

// Type implements Typer interface
func (c Domain) Type() string {
	return TypeDomain
}

// Value implements database/sql/driver.Valuer
func (c Domain) Value() (driver.Value, error) {
	return json.Marshal(c)
}

// Scan implements database/sql.Scanner
func (c *Domain) Scan(src interface{}) error {
	if c == nil {
		return fmt.Errorf("countries::Scan: Domain scan err: domain == nil")
	}
	switch src := src.(type) {
	case *Domain:
		*c = *src
	case Domain:
		*c = src
	case nil:
		c = nil
	default:
		return fmt.Errorf("countries::Scan: domain scan err: unexpected value of type %T for %T", src, *c)
	}
	return nil
}

// DomainCodeByName - return DomainCode by name, case-insensitive, example: domainAE := DomainCodeByName(".ae") OR capitalAE := domainAE("ae")
func DomainCodeByName(name string) DomainCode {
	country := ByName(name)
	if country == Unknown {
		return DomainUnknown
	}
	return DomainCode(country)
}

// AllDomains - returns all domains codes
func AllDomains() []DomainCode {
	return []DomainCode{
		DomainArpa,
		DomainCom,
		DomainOrg,
		DomainNet,
		DomainEdu,
		DomainGov,
		DomainMil,
		DomainTest,
		DomainBiz,
		DomainInfo,
		DomainName,
		DomainAU,
		DomainAT,
		DomainAZ,
		DomainAL,
		DomainDZ,
		DomainAS,
		DomainAI,
		DomainAO,
		DomainAD,
		DomainAQ,
		DomainAG,
		DomainAN,
		DomainAE,
		DomainAR,
		DomainAM,
		DomainAW,
		DomainAF,
		DomainBS,
		DomainBD,
		DomainBB,
		DomainBH,
		DomainBY,
		DomainBZ,
		DomainBE,
		DomainBJ,
		DomainBM,
		DomainBG,
		DomainBO,
		DomainBA,
		DomainBW,
		DomainBR,
		DomainIO,
		DomainBN,
		DomainBF,
		DomainBI,
		DomainBT,
		DomainVU,
		DomainVA,
		DomainGB,
		DomainHU,
		DomainVE,
		DomainVG,
		DomainVI,
		DomainTL,
		DomainVN,
		DomainGA,
		DomainHT,
		DomainGY,
		DomainGM,
		DomainGH,
		DomainGP,
		DomainGT,
		DomainGN,
		DomainGW,
		DomainDE,
		DomainGI,
		DomainHN,
		DomainHK,
		DomainGD,
		DomainGL,
		DomainGR,
		DomainGE,
		DomainGU,
		DomainDK,
		DomainCD,
		DomainDJ,
		DomainDM,
		DomainDO,
		DomainEG,
		DomainZM,
		DomainEH,
		DomainZW,
		DomainIL,
		DomainIN,
		DomainID,
		DomainJO,
		DomainIQ,
		DomainIR,
		DomainIE,
		DomainIS,
		DomainES,
		DomainIT,
		DomainYE,
		DomainKZ,
		DomainKY,
		DomainKH,
		DomainCM,
		DomainCA,
		DomainQA,
		DomainKE,
		DomainCY,
		DomainKI,
		DomainCN,
		DomainCC,
		DomainCO,
		DomainKM,
		DomainCG,
		DomainKP,
		DomainKR,
		DomainCR,
		DomainCI,
		DomainCU,
		DomainKW,
		DomainKG,
		DomainLA,
		DomainLV,
		DomainLS,
		DomainLR,
		DomainLB,
		DomainLY,
		DomainLT,
		DomainLI,
		DomainLU,
		DomainMU,
		DomainMR,
		DomainMG,
		DomainYT,
		DomainMO,
		DomainMK,
		DomainMW,
		DomainMY,
		DomainML,
		DomainMV,
		DomainMT,
		DomainMP,
		DomainMA,
		DomainMQ,
		DomainMH,
		DomainMX,
		DomainFM,
		DomainMZ,
		DomainMD,
		DomainMC,
		DomainMN,
		DomainMS,
		DomainMM,
		DomainNA,
		DomainNR,
		DomainNP,
		DomainNE,
		DomainNG,
		DomainNL,
		DomainNI,
		DomainNU,
		DomainNZ,
		DomainNC,
		DomainNO,
		DomainOM,
		DomainBV,
		DomainIM,
		DomainNF,
		DomainPN,
		DomainCX,
		DomainSH,
		DomainWF,
		DomainHM,
		DomainCV,
		DomainCK,
		DomainWS,
		DomainSJ,
		DomainTC,
		DomainUM,
		DomainPK,
		DomainPW,
		DomainPS,
		DomainPA,
		DomainPG,
		DomainPY,
		DomainPE,
		DomainPL,
		DomainPT,
		DomainPR,
		DomainRE,
		DomainRU,
		DomainRW,
		DomainRO,
		DomainSV,
		DomainSM,
		DomainST,
		DomainSA,
		DomainSZ,
		DomainSC,
		DomainSN,
		DomainPM,
		DomainVC,
		DomainKN,
		DomainLC,
		DomainSG,
		DomainSY,
		DomainSK,
		DomainSI,
		DomainUS,
		DomainSB,
		DomainSO,
		DomainSD,
		DomainSR,
		DomainSL,
		DomainTJ,
		DomainTW,
		DomainTH,
		DomainTZ,
		DomainTG,
		DomainTK,
		DomainTO,
		DomainTT,
		DomainTV,
		DomainTN,
		DomainTM,
		DomainTR,
		DomainUG,
		DomainUZ,
		DomainUA,
		DomainUY,
		DomainFO,
		DomainFJ,
		DomainPH,
		DomainFI,
		DomainFK,
		DomainFR,
		DomainGF,
		DomainPF,
		DomainTF,
		DomainHR,
		DomainCF,
		DomainTD,
		DomainCZ,
		DomainCL,
		DomainCH,
		DomainSE,
		DomainXS,
		DomainLK,
		DomainEC,
		DomainGQ,
		DomainER,
		DomainEE,
		DomainET,
		DomainZA,
		DomainYU,
		DomainGS,
		DomainJM,
		DomainME,
		DomainBL,
		DomainSX,
		DomainRS,
		DomainAX,
		DomainBQ,
		DomainGG,
		DomainJE,
		DomainCW,
		DomainMF,
		DomainSS,
		DomainJP,
	}
}

// AllDomainsInfo - return all domains as []*Domain
func AllDomainsInfo() []*Domain {
	all := AllDomains()
	domains := make([]*Domain, 0, len(all))
	for _, v := range all {
		domains = append(domains, v.Info())
	}
	return domains
}

// TotalDomains - returns number of domains in the package, countries.TotalDomains() == len(countries.AllDomains()) but static value for perfomance
func TotalDomains() int {
	return 263
}
