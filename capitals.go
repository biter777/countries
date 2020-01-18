// Package countries - ISO 3166 (ISO3166-1, ISO3166, Digit, Alpha-2 and Alpha-3) countries codes and names (on eng and rus), ISO 4217 currency designators, ITU-T E.164 IDD calling phone codes, countries capitals, UN M.49 regions codes, ccTLD countries domains, IOC/NOC and FIFA letters codes, VERY FAST, NO maps[], NO slices[], NO external links/files/data, NO interface{}, NO specific dependencies, Databases compatible, Emoji countries flags and currencies support, full support ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and ccTLD standarts. Full support ISO-3166-1, ISO-4217, ITU-T E.164, Unicode CLDR and ccTLD standarts.
package countries

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// CapitalCode - capital code
type CapitalCode int64 // int64 for database/sql/driver.Valuer compatibility

// Capital - capital info
type Capital struct {
	Name    string
	Code    CapitalCode
	Country CountryCode
}

// TypeCapitalCode for Typer interface
const TypeCapitalCode string = "countries.CapitalCode"

// TypeCapital for Typer interface
const TypeCapital string = "countries.Capital"

const (
	// CapitalUnknown CapitalCode = 0
	CapitalUnknown CapitalCode = 0
	// CapitalAU      CapitalCode = CapitalCode(AU)
	CapitalAU CapitalCode = CapitalCode(AU)
	// CapitalAT      CapitalCode = CapitalCode(AT)
	CapitalAT CapitalCode = CapitalCode(AT)
	// CapitalAZ      CapitalCode = CapitalCode(AZ)
	CapitalAZ CapitalCode = CapitalCode(AZ)
	// CapitalAL      CapitalCode = CapitalCode(AL)
	CapitalAL CapitalCode = CapitalCode(AL)
	// CapitalDZ      CapitalCode = CapitalCode(DZ)
	CapitalDZ CapitalCode = CapitalCode(DZ)
	// CapitalAS      CapitalCode = CapitalCode(AS)
	CapitalAS CapitalCode = CapitalCode(AS)
	// CapitalAI      CapitalCode = CapitalCode(AI)
	CapitalAI CapitalCode = CapitalCode(AI)
	// CapitalAO      CapitalCode = CapitalCode(AO)
	CapitalAO CapitalCode = CapitalCode(AO)
	// CapitalAD      CapitalCode = CapitalCode(AD)
	CapitalAD CapitalCode = CapitalCode(AD)
	// CapitalAQ      CapitalCode = CapitalCode(AQ)
	CapitalAQ CapitalCode = CapitalCode(AQ)
	// CapitalAG      CapitalCode = CapitalCode(AG)
	CapitalAG CapitalCode = CapitalCode(AG)
	// CapitalAN      CapitalCode = CapitalCode(AN)
	CapitalAN CapitalCode = CapitalCode(AN)
	// CapitalAE      CapitalCode = CapitalCode(AE)
	CapitalAE CapitalCode = CapitalCode(AE)
	// CapitalAR      CapitalCode = CapitalCode(AR)
	CapitalAR CapitalCode = CapitalCode(AR)
	// CapitalAM      CapitalCode = CapitalCode(AM)
	CapitalAM CapitalCode = CapitalCode(AM)
	// CapitalAW      CapitalCode = CapitalCode(AW)
	CapitalAW CapitalCode = CapitalCode(AW)
	// CapitalAF      CapitalCode = CapitalCode(AF)
	CapitalAF CapitalCode = CapitalCode(AF)
	// CapitalBS      CapitalCode = CapitalCode(BS)
	CapitalBS CapitalCode = CapitalCode(BS)
	// CapitalBD      CapitalCode = CapitalCode(BD)
	CapitalBD CapitalCode = CapitalCode(BD)
	// CapitalBB      CapitalCode = CapitalCode(BB)
	CapitalBB CapitalCode = CapitalCode(BB)
	// CapitalBH      CapitalCode = CapitalCode(BH)
	CapitalBH CapitalCode = CapitalCode(BH)
	// CapitalBY      CapitalCode = CapitalCode(BY)
	CapitalBY CapitalCode = CapitalCode(BY)
	// CapitalBZ      CapitalCode = CapitalCode(BZ)
	CapitalBZ CapitalCode = CapitalCode(BZ)
	// CapitalBE      CapitalCode = CapitalCode(BE)
	CapitalBE CapitalCode = CapitalCode(BE)
	// CapitalBJ      CapitalCode = CapitalCode(BJ)
	CapitalBJ CapitalCode = CapitalCode(BJ)
	// CapitalBM      CapitalCode = CapitalCode(BM)
	CapitalBM CapitalCode = CapitalCode(BM)
	// CapitalBG      CapitalCode = CapitalCode(BG)
	CapitalBG CapitalCode = CapitalCode(BG)
	// CapitalBO      CapitalCode = CapitalCode(BO)
	CapitalBO CapitalCode = CapitalCode(BO)
	// CapitalBA      CapitalCode = CapitalCode(BA)
	CapitalBA CapitalCode = CapitalCode(BA)
	// CapitalBW      CapitalCode = CapitalCode(BW)
	CapitalBW CapitalCode = CapitalCode(BW)
	// CapitalBR      CapitalCode = CapitalCode(BR)
	CapitalBR CapitalCode = CapitalCode(BR)
	// CapitalIO      CapitalCode = CapitalCode(IO)
	CapitalIO CapitalCode = CapitalCode(IO)
	// CapitalBN      CapitalCode = CapitalCode(BN)
	CapitalBN CapitalCode = CapitalCode(BN)
	// CapitalBF      CapitalCode = CapitalCode(BF)
	CapitalBF CapitalCode = CapitalCode(BF)
	// CapitalBI      CapitalCode = CapitalCode(BI)
	CapitalBI CapitalCode = CapitalCode(BI)
	// CapitalBT      CapitalCode = CapitalCode(BT)
	CapitalBT CapitalCode = CapitalCode(BT)
	// CapitalVU      CapitalCode = CapitalCode(VU)
	CapitalVU CapitalCode = CapitalCode(VU)
	// CapitalVA      CapitalCode = CapitalCode(VA)
	CapitalVA CapitalCode = CapitalCode(VA)
	// CapitalGB      CapitalCode = CapitalCode(GB)
	CapitalGB CapitalCode = CapitalCode(GB)
	// CapitalHU      CapitalCode = CapitalCode(HU)
	CapitalHU CapitalCode = CapitalCode(HU)
	// CapitalVE      CapitalCode = CapitalCode(VE)
	CapitalVE CapitalCode = CapitalCode(VE)
	// CapitalVG      CapitalCode = CapitalCode(VG)
	CapitalVG CapitalCode = CapitalCode(VG)
	// CapitalVI      CapitalCode = CapitalCode(VI)
	CapitalVI CapitalCode = CapitalCode(VI)
	// CapitalTL      CapitalCode = CapitalCode(TL)
	CapitalTL CapitalCode = CapitalCode(TL)
	// CapitalVN      CapitalCode = CapitalCode(VN)
	CapitalVN CapitalCode = CapitalCode(VN)
	// CapitalGA      CapitalCode = CapitalCode(GA)
	CapitalGA CapitalCode = CapitalCode(GA)
	// CapitalHT      CapitalCode = CapitalCode(HT)
	CapitalHT CapitalCode = CapitalCode(HT)
	// CapitalGY      CapitalCode = CapitalCode(GY)
	CapitalGY CapitalCode = CapitalCode(GY)
	// CapitalGM      CapitalCode = CapitalCode(GM)
	CapitalGM CapitalCode = CapitalCode(GM)
	// CapitalGH      CapitalCode = CapitalCode(GH)
	CapitalGH CapitalCode = CapitalCode(GH)
	// CapitalGP      CapitalCode = CapitalCode(GP)
	CapitalGP CapitalCode = CapitalCode(GP)
	// CapitalGT      CapitalCode = CapitalCode(GT)
	CapitalGT CapitalCode = CapitalCode(GT)
	// CapitalGN      CapitalCode = CapitalCode(GN)
	CapitalGN CapitalCode = CapitalCode(GN)
	// CapitalGW      CapitalCode = CapitalCode(GW)
	CapitalGW CapitalCode = CapitalCode(GW)
	// CapitalDE      CapitalCode = CapitalCode(DE)
	CapitalDE CapitalCode = CapitalCode(DE)
	// CapitalGI      CapitalCode = CapitalCode(GI)
	CapitalGI CapitalCode = CapitalCode(GI)
	// CapitalHN      CapitalCode = CapitalCode(HN)
	CapitalHN CapitalCode = CapitalCode(HN)
	// CapitalHK      CapitalCode = CapitalCode(HK)
	CapitalHK CapitalCode = CapitalCode(HK)
	// CapitalGD      CapitalCode = CapitalCode(GD)
	CapitalGD CapitalCode = CapitalCode(GD)
	// CapitalGL      CapitalCode = CapitalCode(GL)
	CapitalGL CapitalCode = CapitalCode(GL)
	// CapitalGR      CapitalCode = CapitalCode(GR)
	CapitalGR CapitalCode = CapitalCode(GR)
	// CapitalGE      CapitalCode = CapitalCode(GE)
	CapitalGE CapitalCode = CapitalCode(GE)
	// CapitalGU      CapitalCode = CapitalCode(GU)
	CapitalGU CapitalCode = CapitalCode(GU)
	// CapitalDK      CapitalCode = CapitalCode(DK)
	CapitalDK CapitalCode = CapitalCode(DK)
	// CapitalCD      CapitalCode = CapitalCode(CD)
	CapitalCD CapitalCode = CapitalCode(CD)
	// CapitalDJ      CapitalCode = CapitalCode(DJ)
	CapitalDJ CapitalCode = CapitalCode(DJ)
	// CapitalDM      CapitalCode = CapitalCode(DM)
	CapitalDM CapitalCode = CapitalCode(DM)
	// CapitalDO      CapitalCode = CapitalCode(DO)
	CapitalDO CapitalCode = CapitalCode(DO)
	// CapitalEG      CapitalCode = CapitalCode(EG)
	CapitalEG CapitalCode = CapitalCode(EG)
	// CapitalZM      CapitalCode = CapitalCode(ZM)
	CapitalZM CapitalCode = CapitalCode(ZM)
	// CapitalEH      CapitalCode = CapitalCode(EH)
	CapitalEH CapitalCode = CapitalCode(EH)
	// CapitalZW      CapitalCode = CapitalCode(ZW)
	CapitalZW CapitalCode = CapitalCode(ZW)
	// CapitalIL      CapitalCode = CapitalCode(IL)
	CapitalIL CapitalCode = CapitalCode(IL)
	// CapitalIN      CapitalCode = CapitalCode(IN)
	CapitalIN CapitalCode = CapitalCode(IN)
	// CapitalID      CapitalCode = CapitalCode(ID)
	CapitalID CapitalCode = CapitalCode(ID)
	// CapitalJO      CapitalCode = CapitalCode(JO)
	CapitalJO CapitalCode = CapitalCode(JO)
	// CapitalIQ      CapitalCode = CapitalCode(IQ)
	CapitalIQ CapitalCode = CapitalCode(IQ)
	// CapitalIR      CapitalCode = CapitalCode(IR)
	CapitalIR CapitalCode = CapitalCode(IR)
	// CapitalIE      CapitalCode = CapitalCode(IE)
	CapitalIE CapitalCode = CapitalCode(IE)
	// CapitalIS      CapitalCode = CapitalCode(IS)
	CapitalIS CapitalCode = CapitalCode(IS)
	// CapitalES      CapitalCode = CapitalCode(ES)
	CapitalES CapitalCode = CapitalCode(ES)
	// CapitalIT      CapitalCode = CapitalCode(IT)
	CapitalIT CapitalCode = CapitalCode(IT)
	// CapitalYE      CapitalCode = CapitalCode(YE)
	CapitalYE CapitalCode = CapitalCode(YE)
	// CapitalKZ      CapitalCode = CapitalCode(KZ)
	CapitalKZ CapitalCode = CapitalCode(KZ)
	// CapitalKY      CapitalCode = CapitalCode(KY)
	CapitalKY CapitalCode = CapitalCode(KY)
	// CapitalKH      CapitalCode = CapitalCode(KH)
	CapitalKH CapitalCode = CapitalCode(KH)
	// CapitalCM      CapitalCode = CapitalCode(CM)
	CapitalCM CapitalCode = CapitalCode(CM)
	// CapitalCA      CapitalCode = CapitalCode(CA)
	CapitalCA CapitalCode = CapitalCode(CA)
	// CapitalQA      CapitalCode = CapitalCode(QA)
	CapitalQA CapitalCode = CapitalCode(QA)
	// CapitalKE      CapitalCode = CapitalCode(KE)
	CapitalKE CapitalCode = CapitalCode(KE)
	// CapitalCY      CapitalCode = CapitalCode(CY)
	CapitalCY CapitalCode = CapitalCode(CY)
	// CapitalKI      CapitalCode = CapitalCode(KI)
	CapitalKI CapitalCode = CapitalCode(KI)
	// CapitalCN      CapitalCode = CapitalCode(CN)
	CapitalCN CapitalCode = CapitalCode(CN)
	// CapitalCC      CapitalCode = CapitalCode(CC)
	CapitalCC CapitalCode = CapitalCode(CC)
	// CapitalCO      CapitalCode = CapitalCode(CO)
	CapitalCO CapitalCode = CapitalCode(CO)
	// CapitalKM      CapitalCode = CapitalCode(KM)
	CapitalKM CapitalCode = CapitalCode(KM)
	// CapitalCG      CapitalCode = CapitalCode(CG)
	CapitalCG CapitalCode = CapitalCode(CG)
	// CapitalKP      CapitalCode = CapitalCode(KP)
	CapitalKP CapitalCode = CapitalCode(KP)
	// CapitalKR      CapitalCode = CapitalCode(KR)
	CapitalKR CapitalCode = CapitalCode(KR)
	// CapitalCR      CapitalCode = CapitalCode(CR)
	CapitalCR CapitalCode = CapitalCode(CR)
	// CapitalCI      CapitalCode = CapitalCode(CI)
	CapitalCI CapitalCode = CapitalCode(CI)
	// CapitalCU      CapitalCode = CapitalCode(CU)
	CapitalCU CapitalCode = CapitalCode(CU)
	// CapitalKW      CapitalCode = CapitalCode(KW)
	CapitalKW CapitalCode = CapitalCode(KW)
	// CapitalKG      CapitalCode = CapitalCode(KG)
	CapitalKG CapitalCode = CapitalCode(KG)
	// CapitalLA      CapitalCode = CapitalCode(LA)
	CapitalLA CapitalCode = CapitalCode(LA)
	// CapitalLV      CapitalCode = CapitalCode(LV)
	CapitalLV CapitalCode = CapitalCode(LV)
	// CapitalLS      CapitalCode = CapitalCode(LS)
	CapitalLS CapitalCode = CapitalCode(LS)
	// CapitalLR      CapitalCode = CapitalCode(LR)
	CapitalLR CapitalCode = CapitalCode(LR)
	// CapitalLB      CapitalCode = CapitalCode(LB)
	CapitalLB CapitalCode = CapitalCode(LB)
	// CapitalLY      CapitalCode = CapitalCode(LY)
	CapitalLY CapitalCode = CapitalCode(LY)
	// CapitalLT      CapitalCode = CapitalCode(LT)
	CapitalLT CapitalCode = CapitalCode(LT)
	// CapitalLI      CapitalCode = CapitalCode(LI)
	CapitalLI CapitalCode = CapitalCode(LI)
	// CapitalLU      CapitalCode = CapitalCode(LU)
	CapitalLU CapitalCode = CapitalCode(LU)
	// CapitalMU      CapitalCode = CapitalCode(MU)
	CapitalMU CapitalCode = CapitalCode(MU)
	// CapitalMR      CapitalCode = CapitalCode(MR)
	CapitalMR CapitalCode = CapitalCode(MR)
	// CapitalMG      CapitalCode = CapitalCode(MG)
	CapitalMG CapitalCode = CapitalCode(MG)
	// CapitalYT      CapitalCode = CapitalCode(YT)
	CapitalYT CapitalCode = CapitalCode(YT)
	// CapitalMO      CapitalCode = CapitalCode(MO)
	CapitalMO CapitalCode = CapitalCode(MO)
	// CapitalMK      CapitalCode = CapitalCode(MK)
	CapitalMK CapitalCode = CapitalCode(MK)
	// CapitalMW      CapitalCode = CapitalCode(MW)
	CapitalMW CapitalCode = CapitalCode(MW)
	// CapitalMY      CapitalCode = CapitalCode(MY)
	CapitalMY CapitalCode = CapitalCode(MY)
	// CapitalML      CapitalCode = CapitalCode(ML)
	CapitalML CapitalCode = CapitalCode(ML)
	// CapitalMV      CapitalCode = CapitalCode(MV)
	CapitalMV CapitalCode = CapitalCode(MV)
	// CapitalMT      CapitalCode = CapitalCode(MT)
	CapitalMT CapitalCode = CapitalCode(MT)
	// CapitalMP      CapitalCode = CapitalCode(MP)
	CapitalMP CapitalCode = CapitalCode(MP)
	// CapitalMA      CapitalCode = CapitalCode(MA)
	CapitalMA CapitalCode = CapitalCode(MA)
	// CapitalMQ      CapitalCode = CapitalCode(MQ)
	CapitalMQ CapitalCode = CapitalCode(MQ)
	// CapitalMH      CapitalCode = CapitalCode(MH)
	CapitalMH CapitalCode = CapitalCode(MH)
	// CapitalMX      CapitalCode = CapitalCode(MX)
	CapitalMX CapitalCode = CapitalCode(MX)
	// CapitalFM      CapitalCode = CapitalCode(FM)
	CapitalFM CapitalCode = CapitalCode(FM)
	// CapitalMZ      CapitalCode = CapitalCode(MZ)
	CapitalMZ CapitalCode = CapitalCode(MZ)
	// CapitalMD      CapitalCode = CapitalCode(MD)
	CapitalMD CapitalCode = CapitalCode(MD)
	// CapitalMC      CapitalCode = CapitalCode(MC)
	CapitalMC CapitalCode = CapitalCode(MC)
	// CapitalMN      CapitalCode = CapitalCode(MN)
	CapitalMN CapitalCode = CapitalCode(MN)
	// CapitalMS      CapitalCode = CapitalCode(MS)
	CapitalMS CapitalCode = CapitalCode(MS)
	// CapitalMM      CapitalCode = CapitalCode(MM)
	CapitalMM CapitalCode = CapitalCode(MM)
	// CapitalNA      CapitalCode = CapitalCode(NA)
	CapitalNA CapitalCode = CapitalCode(NA)
	// CapitalNR      CapitalCode = CapitalCode(NR)
	CapitalNR CapitalCode = CapitalCode(NR)
	// CapitalNP      CapitalCode = CapitalCode(NP)
	CapitalNP CapitalCode = CapitalCode(NP)
	// CapitalNE      CapitalCode = CapitalCode(NE)
	CapitalNE CapitalCode = CapitalCode(NE)
	// CapitalNG      CapitalCode = CapitalCode(NG)
	CapitalNG CapitalCode = CapitalCode(NG)
	// CapitalNL      CapitalCode = CapitalCode(NL)
	CapitalNL CapitalCode = CapitalCode(NL)
	// CapitalNI      CapitalCode = CapitalCode(NI)
	CapitalNI CapitalCode = CapitalCode(NI)
	// CapitalNU      CapitalCode = CapitalCode(NU)
	CapitalNU CapitalCode = CapitalCode(NU)
	// CapitalNZ      CapitalCode = CapitalCode(NZ)
	CapitalNZ CapitalCode = CapitalCode(NZ)
	// CapitalNC      CapitalCode = CapitalCode(NC)
	CapitalNC CapitalCode = CapitalCode(NC)
	// CapitalNO      CapitalCode = CapitalCode(NO)
	CapitalNO CapitalCode = CapitalCode(NO)
	// CapitalOM      CapitalCode = CapitalCode(OM)
	CapitalOM CapitalCode = CapitalCode(OM)
	// CapitalBV      CapitalCode = CapitalCode(BV)
	CapitalBV CapitalCode = CapitalCode(BV)
	// CapitalIM      CapitalCode = CapitalCode(IM)
	CapitalIM CapitalCode = CapitalCode(IM)
	// CapitalNF      CapitalCode = CapitalCode(NF)
	CapitalNF CapitalCode = CapitalCode(NF)
	// CapitalPN      CapitalCode = CapitalCode(PN)
	CapitalPN CapitalCode = CapitalCode(PN)
	// CapitalCX      CapitalCode = CapitalCode(CX)
	CapitalCX CapitalCode = CapitalCode(CX)
	// CapitalSH      CapitalCode = CapitalCode(SH)
	CapitalSH CapitalCode = CapitalCode(SH)
	// CapitalWF      CapitalCode = CapitalCode(WF)
	CapitalWF CapitalCode = CapitalCode(WF)
	// CapitalHM      CapitalCode = CapitalCode(HM)
	CapitalHM CapitalCode = CapitalCode(HM)
	// CapitalCV      CapitalCode = CapitalCode(CV)
	CapitalCV CapitalCode = CapitalCode(CV)
	// CapitalCK      CapitalCode = CapitalCode(CK)
	CapitalCK CapitalCode = CapitalCode(CK)
	// CapitalWS      CapitalCode = CapitalCode(WS)
	CapitalWS CapitalCode = CapitalCode(WS)
	// CapitalSJ      CapitalCode = CapitalCode(SJ)
	CapitalSJ CapitalCode = CapitalCode(SJ)
	// CapitalTC      CapitalCode = CapitalCode(TC)
	CapitalTC CapitalCode = CapitalCode(TC)
	// CapitalUM      CapitalCode = CapitalCode(UM)
	CapitalUM CapitalCode = CapitalCode(UM)
	// CapitalPK      CapitalCode = CapitalCode(PK)
	CapitalPK CapitalCode = CapitalCode(PK)
	// CapitalPW      CapitalCode = CapitalCode(PW)
	CapitalPW CapitalCode = CapitalCode(PW)
	// CapitalPS      CapitalCode = CapitalCode(PS)
	CapitalPS CapitalCode = CapitalCode(PS)
	// CapitalPA      CapitalCode = CapitalCode(PA)
	CapitalPA CapitalCode = CapitalCode(PA)
	// CapitalPG      CapitalCode = CapitalCode(PG)
	CapitalPG CapitalCode = CapitalCode(PG)
	// CapitalPY      CapitalCode = CapitalCode(PY)
	CapitalPY CapitalCode = CapitalCode(PY)
	// CapitalPE      CapitalCode = CapitalCode(PE)
	CapitalPE CapitalCode = CapitalCode(PE)
	// CapitalPL      CapitalCode = CapitalCode(PL)
	CapitalPL CapitalCode = CapitalCode(PL)
	// CapitalPT      CapitalCode = CapitalCode(PT)
	CapitalPT CapitalCode = CapitalCode(PT)
	// CapitalPR      CapitalCode = CapitalCode(PR)
	CapitalPR CapitalCode = CapitalCode(PR)
	// CapitalRE      CapitalCode = CapitalCode(RE)
	CapitalRE CapitalCode = CapitalCode(RE)
	// CapitalRU      CapitalCode = CapitalCode(RU)
	CapitalRU CapitalCode = CapitalCode(RU)
	// CapitalRW      CapitalCode = CapitalCode(RW)
	CapitalRW CapitalCode = CapitalCode(RW)
	// CapitalRO      CapitalCode = CapitalCode(RO)
	CapitalRO CapitalCode = CapitalCode(RO)
	// CapitalSV      CapitalCode = CapitalCode(SV)
	CapitalSV CapitalCode = CapitalCode(SV)
	// CapitalSM      CapitalCode = CapitalCode(SM)
	CapitalSM CapitalCode = CapitalCode(SM)
	// CapitalST      CapitalCode = CapitalCode(ST)
	CapitalST CapitalCode = CapitalCode(ST)
	// CapitalSA      CapitalCode = CapitalCode(SA)
	CapitalSA CapitalCode = CapitalCode(SA)
	// CapitalSZ      CapitalCode = CapitalCode(SZ)
	CapitalSZ CapitalCode = CapitalCode(SZ)
	// CapitalSC      CapitalCode = CapitalCode(SC)
	CapitalSC CapitalCode = CapitalCode(SC)
	// CapitalSN      CapitalCode = CapitalCode(SN)
	CapitalSN CapitalCode = CapitalCode(SN)
	// CapitalPM      CapitalCode = CapitalCode(PM)
	CapitalPM CapitalCode = CapitalCode(PM)
	// CapitalVC      CapitalCode = CapitalCode(VC)
	CapitalVC CapitalCode = CapitalCode(VC)
	// CapitalKN      CapitalCode = CapitalCode(KN)
	CapitalKN CapitalCode = CapitalCode(KN)
	// CapitalLC      CapitalCode = CapitalCode(LC)
	CapitalLC CapitalCode = CapitalCode(LC)
	// CapitalSG      CapitalCode = CapitalCode(SG)
	CapitalSG CapitalCode = CapitalCode(SG)
	// CapitalSY      CapitalCode = CapitalCode(SY)
	CapitalSY CapitalCode = CapitalCode(SY)
	// CapitalSK      CapitalCode = CapitalCode(SK)
	CapitalSK CapitalCode = CapitalCode(SK)
	// CapitalSI      CapitalCode = CapitalCode(SI)
	CapitalSI CapitalCode = CapitalCode(SI)
	// CapitalUS      CapitalCode = CapitalCode(US)
	CapitalUS CapitalCode = CapitalCode(US)
	// CapitalSB      CapitalCode = CapitalCode(SB)
	CapitalSB CapitalCode = CapitalCode(SB)
	// CapitalSO      CapitalCode = CapitalCode(SO)
	CapitalSO CapitalCode = CapitalCode(SO)
	// CapitalSD      CapitalCode = CapitalCode(SD)
	CapitalSD CapitalCode = CapitalCode(SD)
	// CapitalSR      CapitalCode = CapitalCode(SR)
	CapitalSR CapitalCode = CapitalCode(SR)
	// CapitalSL      CapitalCode = CapitalCode(SL)
	CapitalSL CapitalCode = CapitalCode(SL)
	// CapitalTJ      CapitalCode = CapitalCode(TJ)
	CapitalTJ CapitalCode = CapitalCode(TJ)
	// CapitalTW      CapitalCode = CapitalCode(TW)
	CapitalTW CapitalCode = CapitalCode(TW)
	// CapitalTH      CapitalCode = CapitalCode(TH)
	CapitalTH CapitalCode = CapitalCode(TH)
	// CapitalTZ      CapitalCode = CapitalCode(TZ)
	CapitalTZ CapitalCode = CapitalCode(TZ)
	// CapitalTG      CapitalCode = CapitalCode(TG)
	CapitalTG CapitalCode = CapitalCode(TG)
	// CapitalTK      CapitalCode = CapitalCode(TK)
	CapitalTK CapitalCode = CapitalCode(TK)
	// CapitalTO      CapitalCode = CapitalCode(TO)
	CapitalTO CapitalCode = CapitalCode(TO)
	// CapitalTT      CapitalCode = CapitalCode(TT)
	CapitalTT CapitalCode = CapitalCode(TT)
	// CapitalTV      CapitalCode = CapitalCode(TV)
	CapitalTV CapitalCode = CapitalCode(TV)
	// CapitalTN      CapitalCode = CapitalCode(TN)
	CapitalTN CapitalCode = CapitalCode(TN)
	// CapitalTM      CapitalCode = CapitalCode(TM)
	CapitalTM CapitalCode = CapitalCode(TM)
	// CapitalTR      CapitalCode = CapitalCode(TR)
	CapitalTR CapitalCode = CapitalCode(TR)
	// CapitalUG      CapitalCode = CapitalCode(UG)
	CapitalUG CapitalCode = CapitalCode(UG)
	// CapitalUZ      CapitalCode = CapitalCode(UZ)
	CapitalUZ CapitalCode = CapitalCode(UZ)
	// CapitalUA      CapitalCode = CapitalCode(UA)
	CapitalUA CapitalCode = CapitalCode(UA)
	// CapitalUY      CapitalCode = CapitalCode(UY)
	CapitalUY CapitalCode = CapitalCode(UY)
	// CapitalFO      CapitalCode = CapitalCode(FO)
	CapitalFO CapitalCode = CapitalCode(FO)
	// CapitalFJ      CapitalCode = CapitalCode(FJ)
	CapitalFJ CapitalCode = CapitalCode(FJ)
	// CapitalPH      CapitalCode = CapitalCode(PH)
	CapitalPH CapitalCode = CapitalCode(PH)
	// CapitalFI      CapitalCode = CapitalCode(FI)
	CapitalFI CapitalCode = CapitalCode(FI)
	// CapitalFK      CapitalCode = CapitalCode(FK)
	CapitalFK CapitalCode = CapitalCode(FK)
	// CapitalFR      CapitalCode = CapitalCode(FR)
	CapitalFR CapitalCode = CapitalCode(FR)
	// CapitalGF      CapitalCode = CapitalCode(GF)
	CapitalGF CapitalCode = CapitalCode(GF)
	// CapitalPF      CapitalCode = CapitalCode(PF)
	CapitalPF CapitalCode = CapitalCode(PF)
	// CapitalTF      CapitalCode = CapitalCode(TF)
	CapitalTF CapitalCode = CapitalCode(TF)
	// CapitalHR      CapitalCode = CapitalCode(HR)
	CapitalHR CapitalCode = CapitalCode(HR)
	// CapitalCF      CapitalCode = CapitalCode(CF)
	CapitalCF CapitalCode = CapitalCode(CF)
	// CapitalTD      CapitalCode = CapitalCode(TD)
	CapitalTD CapitalCode = CapitalCode(TD)
	// CapitalCZ      CapitalCode = CapitalCode(CZ)
	CapitalCZ CapitalCode = CapitalCode(CZ)
	// CapitalCL      CapitalCode = CapitalCode(CL)
	CapitalCL CapitalCode = CapitalCode(CL)
	// CapitalCH      CapitalCode = CapitalCode(CH)
	CapitalCH CapitalCode = CapitalCode(CH)
	// CapitalSE      CapitalCode = CapitalCode(SE)
	CapitalSE CapitalCode = CapitalCode(SE)
	// CapitalXS      CapitalCode = CapitalCode(XS)
	CapitalXS CapitalCode = CapitalCode(XS)
	// CapitalLK      CapitalCode = CapitalCode(LK)
	CapitalLK CapitalCode = CapitalCode(LK)
	// CapitalEC      CapitalCode = CapitalCode(EC)
	CapitalEC CapitalCode = CapitalCode(EC)
	// CapitalGQ      CapitalCode = CapitalCode(GQ)
	CapitalGQ CapitalCode = CapitalCode(GQ)
	// CapitalER      CapitalCode = CapitalCode(ER)
	CapitalER CapitalCode = CapitalCode(ER)
	// CapitalEE      CapitalCode = CapitalCode(EE)
	CapitalEE CapitalCode = CapitalCode(EE)
	// CapitalET      CapitalCode = CapitalCode(ET)
	CapitalET CapitalCode = CapitalCode(ET)
	// CapitalZA      CapitalCode = CapitalCode(ZA)
	CapitalZA CapitalCode = CapitalCode(ZA)
	// CapitalYU      CapitalCode = CapitalCode(YU)
	CapitalYU CapitalCode = CapitalCode(YU)
	// CapitalGS      CapitalCode = CapitalCode(GS)
	CapitalGS CapitalCode = CapitalCode(GS)
	// CapitalJM      CapitalCode = CapitalCode(JM)
	CapitalJM CapitalCode = CapitalCode(JM)
	// CapitalME      CapitalCode = CapitalCode(ME)
	CapitalME CapitalCode = CapitalCode(ME)
	// CapitalBL      CapitalCode = CapitalCode(BL)
	CapitalBL CapitalCode = CapitalCode(BL)
	// CapitalSX      CapitalCode = CapitalCode(SX)
	CapitalSX CapitalCode = CapitalCode(SX)
	// CapitalRS      CapitalCode = CapitalCode(RS)
	CapitalRS CapitalCode = CapitalCode(RS)
	// CapitalAX      CapitalCode = CapitalCode(AX)
	CapitalAX CapitalCode = CapitalCode(AX)
	// CapitalBQ      CapitalCode = CapitalCode(BQ)
	CapitalBQ CapitalCode = CapitalCode(BQ)
	// CapitalGG      CapitalCode = CapitalCode(GG)
	CapitalGG CapitalCode = CapitalCode(GG)
	// CapitalJE      CapitalCode = CapitalCode(JE)
	CapitalJE CapitalCode = CapitalCode(JE)
	// CapitalCW      CapitalCode = CapitalCode(CW)
	CapitalCW CapitalCode = CapitalCode(CW)
	// CapitalMF      CapitalCode = CapitalCode(MF)
	CapitalMF CapitalCode = CapitalCode(MF)
	// CapitalSS      CapitalCode = CapitalCode(SS)
	CapitalSS CapitalCode = CapitalCode(SS)
	// CapitalXK      CapitalCode = CapitalCode(XK)
	CapitalXK CapitalCode = CapitalCode(XK)
	// CapitalJP      CapitalCode = CapitalCode(JP)
	CapitalJP CapitalCode = CapitalCode(JP)
)

// Type implements Typer interface
func (c CapitalCode) Type() string {
	return TypeCapitalCode
}

// String - implements fmt.Stringer, returns a english name the capital of country
//nolint:gocyclo
func (c CapitalCode) String() string { //nolint:gocyclo
	switch c {
	case CapitalAU:
		return "Canberra"
	case CapitalAT:
		return "Vienna"
	case CapitalAZ:
		return "Baku"
	case CapitalAL:
		return "Tirana"
	case CapitalDZ:
		return "Algiers"
	case CapitalAS:
		return "Pago Pago"
	case CapitalAI:
		return "The Valley"
	case CapitalAO:
		return "Luanda"
	case CapitalAD:
		return "Andorra la Vella"
	case CapitalAQ:
		return "None"
	case CapitalAG:
		return "St. John's"
	case CapitalAN:
		return "Willemstad"
	case CapitalAE:
		return "Abu Dhabi"
	case CapitalAR:
		return "Buenos Aires"
	case CapitalAM:
		return "Yerevan"
	case CapitalAW:
		return "Oranjestad"
	case CapitalAF:
		return "Kabul"
	case CapitalBS:
		return "Nassau"
	case CapitalBD:
		return "Dhaka"
	case CapitalBB:
		return "Bridgetown"
	case CapitalBH:
		return "Manama"
	case CapitalBY:
		return "Minsk"
	case CapitalBZ:
		return "Belmopan"
	case CapitalBE:
		return "Brussels"
	case CapitalBJ:
		return "Porto-Novo"
	case CapitalBM:
		return "Hamilton"
	case CapitalBG:
		return "Sofia"
	case CapitalBO:
		return "Sucre"
	case CapitalBA:
		return "Sarajevo"
	case CapitalBW:
		return "Gaborone"
	case CapitalBR:
		return "Brasilia"
	case CapitalIO:
		return "Diego Garcia"
	case CapitalBN:
		return "Bandar Seri Begawan"
	case CapitalBF:
		return "Ouagadougou"
	case CapitalBI:
		return "Bujumbura"
	case CapitalBT:
		return "Thimphu"
	case CapitalVU:
		return "Port Vila"
	case CapitalVA:
		return "Vatican City"
	case CapitalGB:
		return "London"
	case CapitalHU:
		return "Budapest"
	case CapitalVE:
		return "Caracas"
	case CapitalVG:
		return "Road Town"
	case CapitalVI:
		return "Charlotte Amalie"
	case CapitalTL:
		return "Dili"
	case CapitalVN:
		return "Hanoi"
	case CapitalGA:
		return "Libreville"
	case CapitalHT:
		return "Port-au-Prince"
	case CapitalGY:
		return "Georgetown"
	case CapitalGM:
		return "Banjul"
	case CapitalGH:
		return "Accra"
	case CapitalGP:
		return "Basse-Terre"
	case CapitalGT:
		return "Guatemala City"
	case CapitalGN:
		return "Conakry"
	case CapitalGW:
		return "Bissau"
	case CapitalDE:
		return "Berlin"
	case CapitalGI:
		return "Gibraltar"
	case CapitalHN:
		return "Tegucigalpa"
	case CapitalHK:
		return "Hong Kong"
	case CapitalGD:
		return "St. George's"
	case CapitalGL:
		return "Nuuk"
	case CapitalGR:
		return "Athens"
	case CapitalGE:
		return "Tbilisi"
	case CapitalGU:
		return "Hagatna"
	case CapitalDK:
		return "Copenhagen"
	case CapitalCD:
		return "Kinshasa"
	case CapitalDJ:
		return "Djibouti"
	case CapitalDM:
		return "Roseau"
	case CapitalDO:
		return "Santo Domingo"
	case CapitalEG:
		return "Cairo"
	case CapitalZM:
		return "Lusaka"
	case CapitalEH:
		return "El-Aaiun"
	case CapitalZW:
		return "Harare"
	case CapitalIL:
		return "Jerusalem"
	case CapitalIN:
		return "New Delhi"
	case CapitalID:
		return "Jakarta"
	case CapitalJO:
		return "Amman"
	case CapitalIQ:
		return "Baghdad"
	case CapitalIR:
		return "Tehran"
	case CapitalIE:
		return "Dublin"
	case CapitalIS:
		return "Reykjavik"
	case CapitalES:
		return "Madrid"
	case CapitalIT:
		return "Rome"
	case CapitalYE:
		return "Sanaa"
	case CapitalKZ:
		return "Nur-Sultan"
	case CapitalKY:
		return "George Town"
	case CapitalKH:
		return "Phnom Penh"
	case CapitalCM:
		return "Yaounde"
	case CapitalCA:
		return "Ottawa"
	case CapitalQA:
		return "Doha"
	case CapitalKE:
		return "Nairobi"
	case CapitalCY:
		return "Nicosia"
	case CapitalKI:
		return "Tarawa"
	case CapitalCN:
		return "Beijing"
	case CapitalCC:
		return "West Island"
	case CapitalCO:
		return "Bogota"
	case CapitalKM:
		return "Moroni"
	case CapitalCG:
		return "Brazzaville"
	case CapitalKP:
		return "Pyongyang"
	case CapitalKR:
		return "Seoul"
	case CapitalCR:
		return "San Jose"
	case CapitalCI:
		return "Yamoussoukro"
	case CapitalCU:
		return "Havana"
	case CapitalKW:
		return "Kuwait City"
	case CapitalKG:
		return "Bishkek"
	case CapitalLA:
		return "Vientiane"
	case CapitalLV:
		return "Riga"
	case CapitalLS:
		return "Maseru"
	case CapitalLR:
		return "Monrovia"
	case CapitalLB:
		return "Beirut"
	case CapitalLY:
		return "Tripoli"
	case CapitalLT:
		return "Vilnius"
	case CapitalLI:
		return "Vaduz"
	case CapitalLU:
		return "Luxembourg"
	case CapitalMU:
		return "Port Louis"
	case CapitalMR:
		return "Nouakchott"
	case CapitalMG:
		return "Antananarivo"
	case CapitalYT:
		return "Mamoudzou"
	case CapitalMO:
		return "Macao"
	case CapitalMK:
		return "Skopje"
	case CapitalMW:
		return "Lilongwe"
	case CapitalMY:
		return "Kuala Lumpur"
	case CapitalML:
		return "Bamako"
	case CapitalMV:
		return "Male"
	case CapitalMT:
		return "Valletta"
	case CapitalMP:
		return "Saipan"
	case CapitalMA:
		return "Rabat"
	case CapitalMQ:
		return "Fort-de-France"
	case CapitalMH:
		return "Majuro"
	case CapitalMX:
		return "Mexico City"
	case CapitalFM:
		return "Palikir"
	case CapitalMZ:
		return "Maputo"
	case CapitalMD:
		return "Chisinau"
	case CapitalMC:
		return "Monaco"
	case CapitalMN:
		return "Ulaanbaatar"
	case CapitalMS:
		return "Plymouth"
	case CapitalMM:
		return "Nay Pyi Taw"
	case CapitalNA:
		return "Windhoek"
	case CapitalNR:
		return "Yaren"
	case CapitalNP:
		return "Kathmandu"
	case CapitalNE:
		return "Niamey"
	case CapitalNG:
		return "Abuja"
	case CapitalNL:
		return "Amsterdam"
	case CapitalNI:
		return "Managua"
	case CapitalNU:
		return "Alofi"
	case CapitalNZ:
		return "Wellington"
	case CapitalNC:
		return "Noumea"
	case CapitalNO:
		return "Oslo"
	case CapitalOM:
		return "Muscat"
	case CapitalBV:
		return "None"
	case CapitalIM:
		return "Douglas"
	case CapitalNF:
		return "Kingston"
	case CapitalPN:
		return "Adamstown"
	case CapitalCX:
		return "Flying Fish Cove"
	case CapitalSH:
		return "Jamestown"
	case CapitalWF:
		return "Mata Utu"
	case CapitalHM:
		return "None"
	case CapitalCV:
		return "Praia"
	case CapitalCK:
		return "Avarua"
	case CapitalWS:
		return "Apia"
	case CapitalSJ:
		return "Longyearbyen"
	case CapitalTC:
		return "Cockburn Town"
	case CapitalUM:
		return "None"
	case CapitalPK:
		return "Islamabad"
	case CapitalPW:
		return "Melekeok"
	case CapitalPS:
		return "East Jerusalem"
	case CapitalPA:
		return "Panama City"
	case CapitalPG:
		return "Port Moresby"
	case CapitalPY:
		return "Asuncion"
	case CapitalPE:
		return "Lima"
	case CapitalPL:
		return "Warsaw"
	case CapitalPT:
		return "Lisbon"
	case CapitalPR:
		return "San Juan"
	case CapitalRE:
		return "Saint-Denis"
	case CapitalRU:
		return "Moscow"
	case CapitalRW:
		return "Kigali"
	case CapitalRO:
		return "Bucharest"
	case CapitalSV:
		return "San Salvador"
	case CapitalSM:
		return "San Marino"
	case CapitalST:
		return "Sao Tome"
	case CapitalSA:
		return "Riyadh"
	case CapitalSZ:
		return "Mbabane"
	case CapitalSC:
		return "Victoria"
	case CapitalSN:
		return "Dakar"
	case CapitalPM:
		return "Saint-Pierre"
	case CapitalVC:
		return "Kingstown"
	case CapitalKN:
		return "Basseterre"
	case CapitalLC:
		return "Castries"
	case CapitalSG:
		return "Singapore"
	case CapitalSY:
		return "Damascus"
	case CapitalSK:
		return "Bratislava"
	case CapitalSI:
		return "Ljubljana"
	case CapitalUS:
		return "Washington"
	case CapitalSB:
		return "Honiara"
	case CapitalSO:
		return "Mogadishu"
	case CapitalSD:
		return "Khartoum"
	case CapitalSR:
		return "Paramaribo"
	case CapitalSL:
		return "Freetown"
	case CapitalTJ:
		return "Dushanbe"
	case CapitalTW:
		return "Taipei"
	case CapitalTH:
		return "Bangkok"
	case CapitalTZ:
		return "Dodoma"
	case CapitalTG:
		return "Lome"
	case CapitalTK:
		return "None"
	case CapitalTO:
		return "Nuku'alofa"
	case CapitalTT:
		return "Port of Spain"
	case CapitalTV:
		return "Funafuti"
	case CapitalTN:
		return "Tunis"
	case CapitalTM:
		return "Ashgabat"
	case CapitalTR:
		return "Ankara"
	case CapitalUG:
		return "Kampala"
	case CapitalUZ:
		return "Tashkent"
	case CapitalUA:
		return "Kyiv"
	case CapitalUY:
		return "Montevideo"
	case CapitalFO:
		return "Torshavn"
	case CapitalFJ:
		return "Suva"
	case CapitalPH:
		return "Manila"
	case CapitalFI:
		return "Helsinki"
	case CapitalFK:
		return "Stanley"
	case CapitalFR:
		return "Paris"
	case CapitalGF:
		return "Cayenne"
	case CapitalPF:
		return "Papeete"
	case CapitalTF:
		return "Port-aux-Francais"
	case CapitalHR:
		return "Zagreb"
	case CapitalCF:
		return "Bangui"
	case CapitalTD:
		return "N'Djamena"
	case CapitalCZ:
		return "Prague"
	case CapitalCL:
		return "Santiago"
	case CapitalCH:
		return "Bern"
	case CapitalSE:
		return "Stockholm"
	case CapitalLK:
		return "Colombo"
	case CapitalEC:
		return "Quito"
	case CapitalGQ:
		return "Malabo"
	case CapitalER:
		return "Asmara"
	case CapitalEE:
		return "Tallinn"
	case CapitalET:
		return "Addis Ababa"
	case CapitalZA:
		return "Pretoria"
	case CapitalYU:
		return "Belgrade"
	case CapitalGS:
		return "Grytviken"
	case CapitalJM:
		return "Kingston"
	case CapitalME:
		return "Podgorica"
	case CapitalBL:
		return "Gustavia"
	case CapitalSX:
		return "Philipsburg"
	case CapitalRS:
		return "Belgrade"
	case CapitalAX:
		return "Mariehamn"
	case CapitalBQ:
		return "None"
	case CapitalGG:
		return "St Peter Port"
	case CapitalJE:
		return "Saint Helier"
	case CapitalCW:
		return "Willemstad"
	case CapitalMF:
		return "Marigot"
	case CapitalSS:
		return "Juba"
	case CapitalXK:
		return "Pristina"
	case CapitalJP:
		return "Tokyo"
	}
	return UnknownMsg
}

// Country - returns a country of capital
//nolint:gocyclo
func (c CapitalCode) Country() CountryCode { //nolint:gocyclo
	switch c {
	case CapitalAU:
		return AU
	case CapitalAT:
		return AT
	case CapitalAZ:
		return AZ
	case CapitalAL:
		return AL
	case CapitalDZ:
		return DZ
	case CapitalAS:
		return AS
	case CapitalAI:
		return AI
	case CapitalAO:
		return AO
	case CapitalAD:
		return AD
	case CapitalAQ:
		return AQ
	case CapitalAG:
		return AG
	case CapitalAN:
		return AN
	case CapitalAE:
		return AE
	case CapitalAR:
		return AR
	case CapitalAM:
		return AM
	case CapitalAW:
		return AW
	case CapitalAF:
		return AF
	case CapitalBS:
		return BS
	case CapitalBD:
		return BD
	case CapitalBB:
		return BB
	case CapitalBH:
		return BH
	case CapitalBY:
		return BY
	case CapitalBZ:
		return BZ
	case CapitalBE:
		return BE
	case CapitalBJ:
		return BJ
	case CapitalBM:
		return BM
	case CapitalBG:
		return BG
	case CapitalBO:
		return BO
	case CapitalBA:
		return BA
	case CapitalBW:
		return BW
	case CapitalBR:
		return BR
	case CapitalIO:
		return IO
	case CapitalBN:
		return BN
	case CapitalBF:
		return BF
	case CapitalBI:
		return BI
	case CapitalBT:
		return BT
	case CapitalVU:
		return VU
	case CapitalVA:
		return VA
	case CapitalGB:
		return GB
	case CapitalHU:
		return HU
	case CapitalVE:
		return VE
	case CapitalVG:
		return VG
	case CapitalVI:
		return VI
	case CapitalTL:
		return TL
	case CapitalVN:
		return VN
	case CapitalGA:
		return GA
	case CapitalHT:
		return HT
	case CapitalGY:
		return GY
	case CapitalGM:
		return GM
	case CapitalGH:
		return GH
	case CapitalGP:
		return GP
	case CapitalGT:
		return GT
	case CapitalGN:
		return GN
	case CapitalGW:
		return GW
	case CapitalDE:
		return DE
	case CapitalGI:
		return GI
	case CapitalHN:
		return HN
	case CapitalHK:
		return HK
	case CapitalGD:
		return GD
	case CapitalGL:
		return GL
	case CapitalGR:
		return GR
	case CapitalGE:
		return GE
	case CapitalGU:
		return GU
	case CapitalDK:
		return DK
	case CapitalCD:
		return CD
	case CapitalDJ:
		return DJ
	case CapitalDM:
		return DM
	case CapitalDO:
		return DO
	case CapitalEG:
		return EG
	case CapitalZM:
		return ZM
	case CapitalEH:
		return EH
	case CapitalZW:
		return ZW
	case CapitalIL:
		return IL
	case CapitalIN:
		return IN
	case CapitalID:
		return ID
	case CapitalJO:
		return JO
	case CapitalIQ:
		return IQ
	case CapitalIR:
		return IR
	case CapitalIE:
		return IE
	case CapitalIS:
		return IS
	case CapitalES:
		return ES
	case CapitalIT:
		return IT
	case CapitalYE:
		return YE
	case CapitalKZ:
		return KZ
	case CapitalKY:
		return KY
	case CapitalKH:
		return KH
	case CapitalCM:
		return CM
	case CapitalCA:
		return CA
	case CapitalQA:
		return QA
	case CapitalKE:
		return KE
	case CapitalCY:
		return CY
	case CapitalKI:
		return KI
	case CapitalCN:
		return CN
	case CapitalCC:
		return CC
	case CapitalCO:
		return CO
	case CapitalKM:
		return KM
	case CapitalCG:
		return CG
	case CapitalKP:
		return KP
	case CapitalKR:
		return KR
	case CapitalCR:
		return CR
	case CapitalCI:
		return CI
	case CapitalCU:
		return CU
	case CapitalKW:
		return KW
	case CapitalKG:
		return KG
	case CapitalLA:
		return LA
	case CapitalLV:
		return LV
	case CapitalLS:
		return LS
	case CapitalLR:
		return LR
	case CapitalLB:
		return LB
	case CapitalLY:
		return LY
	case CapitalLT:
		return LT
	case CapitalLI:
		return LI
	case CapitalLU:
		return LU
	case CapitalMU:
		return MU
	case CapitalMR:
		return MR
	case CapitalMG:
		return MG
	case CapitalYT:
		return YT
	case CapitalMO:
		return MO
	case CapitalMK:
		return MK
	case CapitalMW:
		return MW
	case CapitalMY:
		return MY
	case CapitalML:
		return ML
	case CapitalMV:
		return MV
	case CapitalMT:
		return MT
	case CapitalMP:
		return MP
	case CapitalMA:
		return MA
	case CapitalMQ:
		return MQ
	case CapitalMH:
		return MH
	case CapitalMX:
		return MX
	case CapitalFM:
		return FM
	case CapitalMZ:
		return MZ
	case CapitalMD:
		return MD
	case CapitalMC:
		return MC
	case CapitalMN:
		return MN
	case CapitalMS:
		return MS
	case CapitalMM:
		return MM
	case CapitalNA:
		return NA
	case CapitalNR:
		return NR
	case CapitalNP:
		return NP
	case CapitalNE:
		return NE
	case CapitalNG:
		return NG
	case CapitalNL:
		return NL
	case CapitalNI:
		return NI
	case CapitalNU:
		return NU
	case CapitalNZ:
		return NZ
	case CapitalNC:
		return NC
	case CapitalNO:
		return NO
	case CapitalOM:
		return OM
	case CapitalBV:
		return BV
	case CapitalIM:
		return IM
	case CapitalNF:
		return NF
	case CapitalPN:
		return PN
	case CapitalCX:
		return CX
	case CapitalSH:
		return SH
	case CapitalWF:
		return WF
	case CapitalHM:
		return HM
	case CapitalCV:
		return CV
	case CapitalCK:
		return CK
	case CapitalWS:
		return WS
	case CapitalSJ:
		return SJ
	case CapitalTC:
		return TC
	case CapitalUM:
		return UM
	case CapitalPK:
		return PK
	case CapitalPW:
		return PW
	case CapitalPS:
		return PS
	case CapitalPA:
		return PA
	case CapitalPG:
		return PG
	case CapitalPY:
		return PY
	case CapitalPE:
		return PE
	case CapitalPL:
		return PL
	case CapitalPT:
		return PT
	case CapitalPR:
		return PR
	case CapitalRE:
		return RE
	case CapitalRU:
		return RU
	case CapitalRW:
		return RW
	case CapitalRO:
		return RO
	case CapitalSV:
		return SV
	case CapitalSM:
		return SM
	case CapitalST:
		return ST
	case CapitalSA:
		return SA
	case CapitalSZ:
		return SZ
	case CapitalSC:
		return SC
	case CapitalSN:
		return SN
	case CapitalPM:
		return PM
	case CapitalVC:
		return VC
	case CapitalKN:
		return KN
	case CapitalLC:
		return LC
	case CapitalSG:
		return SG
	case CapitalSY:
		return SY
	case CapitalSK:
		return SK
	case CapitalSI:
		return SI
	case CapitalUS:
		return US
	case CapitalSB:
		return SB
	case CapitalSO:
		return SO
	case CapitalSD:
		return SD
	case CapitalSR:
		return SR
	case CapitalSL:
		return SL
	case CapitalTJ:
		return TJ
	case CapitalTW:
		return TW
	case CapitalTH:
		return TH
	case CapitalTZ:
		return TZ
	case CapitalTG:
		return TG
	case CapitalTK:
		return TK
	case CapitalTO:
		return TO
	case CapitalTT:
		return TT
	case CapitalTV:
		return TV
	case CapitalTN:
		return TN
	case CapitalTM:
		return TM
	case CapitalTR:
		return TR
	case CapitalUG:
		return UG
	case CapitalUZ:
		return UZ
	case CapitalUA:
		return UA
	case CapitalUY:
		return UY
	case CapitalFO:
		return FO
	case CapitalFJ:
		return FJ
	case CapitalPH:
		return PH
	case CapitalFI:
		return FI
	case CapitalFK:
		return FK
	case CapitalFR:
		return FR
	case CapitalGF:
		return GF
	case CapitalPF:
		return PF
	case CapitalTF:
		return TF
	case CapitalHR:
		return HR
	case CapitalCF:
		return CF
	case CapitalTD:
		return TD
	case CapitalCZ:
		return CZ
	case CapitalCL:
		return CL
	case CapitalCH:
		return CH
	case CapitalSE:
		return SE
	case CapitalLK:
		return LK
	case CapitalEC:
		return EC
	case CapitalGQ:
		return GQ
	case CapitalER:
		return ER
	case CapitalEE:
		return EE
	case CapitalET:
		return ET
	case CapitalZA:
		return ZA
	case CapitalYU:
		return YU
	case CapitalGS:
		return GS
	case CapitalJM:
		return JM
	case CapitalME:
		return ME
	case CapitalBL:
		return BL
	case CapitalSX:
		return SX
	case CapitalRS:
		return RS
	case CapitalAX:
		return AX
	case CapitalBQ:
		return BQ
	case CapitalGG:
		return GG
	case CapitalJE:
		return JE
	case CapitalCW:
		return CW
	case CapitalMF:
		return MF
	case CapitalSS:
		return SS
	case CapitalXK:
		return XK
	case CapitalJP:
		return JP
	}
	return Unknown
}

// IsValid - returns true, if code is correct
func (c CapitalCode) IsValid() bool {
	return c.String() != UnknownMsg
}

// Info - return CapitalCode as Capital info
func (c CapitalCode) Info() *Capital {
	return &Capital{
		Name:    c.String(),
		Code:    c,
		Country: c.Country(),
	}
}

// Type implements Typer interface
func (c Capital) Type() string {
	return TypeCapital
}

// Value implements database/sql/driver.Valuer
func (c Capital) Value() (driver.Value, error) {
	return json.Marshal(c)
}

// Scan implements database/sql.Scanner
func (c *Capital) Scan(src interface{}) error {
	if c == nil {
		return fmt.Errorf("countries::Scan: Capital scan err: capital == nil")
	}
	switch src := src.(type) {
	case *Capital:
		*c = *src
	case Capital:
		*c = src
	default:
		return fmt.Errorf("countries::Scan: Capital scan err: unexpected value of type %T for %T", src, *c)
	}
	return nil
}

// AllCapitals - return all capital codes
func AllCapitals() []CapitalCode {
	return []CapitalCode{
		CapitalAU,
		CapitalAT,
		CapitalAZ,
		CapitalAL,
		CapitalDZ,
		CapitalAS,
		CapitalAI,
		CapitalAO,
		CapitalAD,
		CapitalAQ,
		CapitalAG,
		CapitalAN,
		CapitalAE,
		CapitalAR,
		CapitalAM,
		CapitalAW,
		CapitalAF,
		CapitalBS,
		CapitalBD,
		CapitalBB,
		CapitalBH,
		CapitalBY,
		CapitalBZ,
		CapitalBE,
		CapitalBJ,
		CapitalBM,
		CapitalBG,
		CapitalBO,
		CapitalBA,
		CapitalBW,
		CapitalBR,
		CapitalIO,
		CapitalBN,
		CapitalBF,
		CapitalBI,
		CapitalBT,
		CapitalVU,
		CapitalVA,
		CapitalGB,
		CapitalHU,
		CapitalVE,
		CapitalVG,
		CapitalVI,
		CapitalTL,
		CapitalVN,
		CapitalGA,
		CapitalHT,
		CapitalGY,
		CapitalGM,
		CapitalGH,
		CapitalGP,
		CapitalGT,
		CapitalGN,
		CapitalGW,
		CapitalDE,
		CapitalGI,
		CapitalHN,
		CapitalHK,
		CapitalGD,
		CapitalGL,
		CapitalGR,
		CapitalGE,
		CapitalGU,
		CapitalDK,
		CapitalCD,
		CapitalDJ,
		CapitalDM,
		CapitalDO,
		CapitalEG,
		CapitalZM,
		CapitalEH,
		CapitalZW,
		CapitalIL,
		CapitalIN,
		CapitalID,
		CapitalJO,
		CapitalIQ,
		CapitalIR,
		CapitalIE,
		CapitalIS,
		CapitalES,
		CapitalIT,
		CapitalYE,
		CapitalKZ,
		CapitalKY,
		CapitalKH,
		CapitalCM,
		CapitalCA,
		CapitalQA,
		CapitalKE,
		CapitalCY,
		CapitalKI,
		CapitalCN,
		CapitalCC,
		CapitalCO,
		CapitalKM,
		CapitalCG,
		CapitalKP,
		CapitalKR,
		CapitalCR,
		CapitalCI,
		CapitalCU,
		CapitalKW,
		CapitalKG,
		CapitalLA,
		CapitalLV,
		CapitalLS,
		CapitalLR,
		CapitalLB,
		CapitalLY,
		CapitalLT,
		CapitalLI,
		CapitalLU,
		CapitalMU,
		CapitalMR,
		CapitalMG,
		CapitalYT,
		CapitalMO,
		CapitalMK,
		CapitalMW,
		CapitalMY,
		CapitalML,
		CapitalMV,
		CapitalMT,
		CapitalMP,
		CapitalMA,
		CapitalMQ,
		CapitalMH,
		CapitalMX,
		CapitalFM,
		CapitalMZ,
		CapitalMD,
		CapitalMC,
		CapitalMN,
		CapitalMS,
		CapitalMM,
		CapitalNA,
		CapitalNR,
		CapitalNP,
		CapitalNE,
		CapitalNG,
		CapitalNL,
		CapitalNI,
		CapitalNU,
		CapitalNZ,
		CapitalNC,
		CapitalNO,
		CapitalOM,
		CapitalBV,
		CapitalIM,
		CapitalNF,
		CapitalPN,
		CapitalCX,
		CapitalSH,
		CapitalWF,
		CapitalHM,
		CapitalCV,
		CapitalCK,
		CapitalWS,
		CapitalSJ,
		CapitalTC,
		CapitalUM,
		CapitalPK,
		CapitalPW,
		CapitalPS,
		CapitalPA,
		CapitalPG,
		CapitalPY,
		CapitalPE,
		CapitalPL,
		CapitalPT,
		CapitalPR,
		CapitalRE,
		CapitalRU,
		CapitalRW,
		CapitalRO,
		CapitalSV,
		CapitalSM,
		CapitalST,
		CapitalSA,
		CapitalSZ,
		CapitalSC,
		CapitalSN,
		CapitalPM,
		CapitalVC,
		CapitalKN,
		CapitalLC,
		CapitalSG,
		CapitalSY,
		CapitalSK,
		CapitalSI,
		CapitalUS,
		CapitalSB,
		CapitalSO,
		CapitalSD,
		CapitalSR,
		CapitalSL,
		CapitalTJ,
		CapitalTW,
		CapitalTH,
		CapitalTZ,
		CapitalTG,
		CapitalTK,
		CapitalTO,
		CapitalTT,
		CapitalTV,
		CapitalTN,
		CapitalTM,
		CapitalTR,
		CapitalUG,
		CapitalUZ,
		CapitalUA,
		CapitalUY,
		CapitalFO,
		CapitalFJ,
		CapitalPH,
		CapitalFI,
		CapitalFK,
		CapitalFR,
		CapitalGF,
		CapitalPF,
		CapitalTF,
		CapitalHR,
		CapitalCF,
		CapitalTD,
		CapitalCZ,
		CapitalCL,
		CapitalCH,
		CapitalSE,
		CapitalXS,
		CapitalLK,
		CapitalEC,
		CapitalGQ,
		CapitalER,
		CapitalEE,
		CapitalET,
		CapitalZA,
		CapitalYU,
		CapitalGS,
		CapitalJM,
		CapitalME,
		CapitalBL,
		CapitalSX,
		CapitalRS,
		CapitalAX,
		CapitalBQ,
		CapitalGG,
		CapitalJE,
		CapitalCW,
		CapitalMF,
		CapitalSS,
		CapitalXK,
		CapitalJP,
	}
}

// AllCapitalsInfo - return all capital codes as []Capital
func AllCapitalsInfo() []*Capital {
	all := AllCapitals()
	capitals := make([]*Capital, 0, len(all))
	for _, v := range all {
		capitals = append(capitals, v.Info())
	}
	return capitals
}

// CapitalCodeByName - return CapitalCode by name, case-insensitive, example: capitalAE := CapitalCodeByName("Abu-Dhabi") OR capitalAE := CapitalCodeByName("abu-dhabi")
//nolint:gocyclo
func CapitalCodeByName(name string) CapitalCode { //nolint:gocyclo
	switch textPrepare(name) {
	case "ORANJESTAD":
		return CapitalAW
	case "STJOHNS":
		return CapitalAG
	case "ABUDHABI":
		return CapitalAE
	case "KABUL":
		return CapitalAF
	case "ALGIERS":
		return CapitalDZ
	case "BAKU":
		return CapitalAZ
	case "TIRANA":
		return CapitalAL
	case "YEREVAN":
		return CapitalAM
	case "ANDORRALAVELLA":
		return CapitalAD
	case "LUANDA":
		return CapitalAO
	case "PAGOPAGO":
		return CapitalAS
	case "BUENOSAIRES":
		return CapitalAR
	case "CANBERRA":
		return CapitalAU
	case "VIENNA":
		return CapitalAT
	case "THEVALLEY":
		return CapitalAI
	case "MANAMA":
		return CapitalBH
	case "BRIDGETOWN":
		return CapitalBB
	case "GABORONE":
		return CapitalBW
	case "HAMILTON":
		return CapitalBM
	case "BRUSSELS":
		return CapitalBE
	case "NASSAU":
		return CapitalBS
	case "DHAKA":
		return CapitalBD
	case "BELMOPAN":
		return CapitalBZ
	case "SARAJEVO":
		return CapitalBA
	case "SUCRE":
		return CapitalBO
	case "NAYPYITAW":
		return CapitalMM
	case "PORTO-NOVO":
		return CapitalBJ
	case "MINSK":
		return CapitalBY
	case "HONIARA":
		return CapitalSB
	case "BRASILIA":
		return CapitalBR
	case "THIMPHU":
		return CapitalBT
	case "SOFIA":
		return CapitalBG
	case "BANDARSERIBEGAWAN":
		return CapitalBN
	case "BUJUMBURA":
		return CapitalBI
	case "OTTAWA":
		return CapitalCA
	case "PHNOMPENH":
		return CapitalKH
	case "NDJAMENA":
		return CapitalTD
	case "COLOMBO":
		return CapitalLK
	case "BRAZZAVILLE":
		return CapitalCG
	case "KINSHASA":
		return CapitalCD
	case "BEIJING":
		return CapitalCN
	case "SANTIAGO":
		return CapitalCL
	case "GEORGETOWNCAYMANISLANDS", "CAYMANISLANDSGEORGETOWN":
		return CapitalKY
	case "WESTISLAND":
		return CapitalCC
	case "YAOUNDE":
		return CapitalCM
	case "MORONI":
		return CapitalKM
	case "BOGOTA":
		return CapitalCO
	case "SAIPAN":
		return CapitalMP
	case "SANJOSE":
		return CapitalCR
	case "BANGUI":
		return CapitalCF
	case "HAVANA":
		return CapitalCU
	case "PRAIA":
		return CapitalCV
	case "AVARUA":
		return CapitalCK
	case "NICOSIA":
		return CapitalCY
	case "COPENHAGEN":
		return CapitalDK
	case "DJIBOUTI":
		return CapitalDJ
	case "ROSEAU":
		return CapitalDM
	case "SANTODOMINGO":
		return CapitalDO
	case "QUITO":
		return CapitalEC
	case "CAIRO":
		return CapitalEG
	case "DUBLIN":
		return CapitalIE
	case "MALABO":
		return CapitalGQ
	case "TALLINN":
		return CapitalEE
	case "ASMARA":
		return CapitalER
	case "SANSALVADOR":
		return CapitalSV
	case "ADDISABABA":
		return CapitalET
	case "PRAGUE":
		return CapitalCZ
	case "CAYENNE":
		return CapitalGF
	case "HELSINKI":
		return CapitalFI
	case "SUVA":
		return CapitalFJ
	case "STANLEY":
		return CapitalFK
	case "PALIKIR":
		return CapitalFM
	case "TORSHAVN":
		return CapitalFO
	case "PAPEETE":
		return CapitalPF
	case "PARIS":
		return CapitalFR
	case "PORT-AUX-FRANCAIS":
		return CapitalTF
	case "BANJUL":
		return CapitalGM
	case "LIBREVILLE":
		return CapitalGA
	case "TBILISI":
		return CapitalGE
	case "ACCRA":
		return CapitalGH
	case "GIBRALTAR":
		return CapitalGI
	case "STGEORGES":
		return CapitalGD
	case "STPETERPORT":
		return CapitalGG
	case "NUUK":
		return CapitalGL
	case "BERLIN":
		return CapitalDE
	case "BASSE-TERRE":
		return CapitalGP
	case "HAGATNA":
		return CapitalGU
	case "ATHENS":
		return CapitalGR
	case "GUATEMALACITY":
		return CapitalGT
	case "CONAKRY":
		return CapitalGN
	case "GEORGETOWNGUYANA", "GUYANAGEORGETOWN":
		return CapitalGY
	case "PORT-AU-PRINCE":
		return CapitalHT
	case "HONGKONG":
		return CapitalHK
	case "TEGUCIGALPA":
		return CapitalHN
	case "ZAGREB":
		return CapitalHR
	case "BUDAPEST":
		return CapitalHU
	case "REYKJAVIK":
		return CapitalIS
	case "JAKARTA":
		return CapitalID
	case "DOUGLAS":
		return CapitalIM
	case "NEWDELHI":
		return CapitalIN
	case "DIEGOGARCIA":
		return CapitalIO
	case "TEHRAN":
		return CapitalIR
	case "JERUSALEM":
		return CapitalIL
	case "ROME":
		return CapitalIT
	case "YAMOUSSOUKRO":
		return CapitalCI
	case "BAGHDAD":
		return CapitalIQ
	case "TOKYO":
		return CapitalJP
	case "SAINTHELIER":
		return CapitalJE
	case "KINGSTON", "KINGSTONJAMAICA", "KINGSTONJAMAIKA", "JAMAICAKINGSTON", "JAMAIKAKINGSTON":
		return CapitalJM
	case "AMMAN":
		return CapitalJO
	case "NAIROBI":
		return CapitalKE
	case "BISHKEK":
		return CapitalKG
	case "PYONGYANG":
		return CapitalKP
	case "TARAWA":
		return CapitalKI
	case "SEOUL":
		return CapitalKR
	case "FLYINGFISHCOVE":
		return CapitalCX
	case "KUWAITCITY":
		return CapitalKW
	case "NUR-SULTAN":
		return CapitalKZ
	case "VIENTIANE":
		return CapitalLA
	case "BEIRUT":
		return CapitalLB
	case "RIGA":
		return CapitalLV
	case "VILNIUS":
		return CapitalLT
	case "MONROVIA":
		return CapitalLR
	case "BRATISLAVA":
		return CapitalSK
	case "VADUZ":
		return CapitalLI
	case "MASERU":
		return CapitalLS
	case "LUXEMBOURG":
		return CapitalLU
	case "TRIPOLI":
		return CapitalLY
	case "ANTANANARIVO":
		return CapitalMG
	case "FORT-DE-FRANCE":
		return CapitalMQ
	case "MACAO":
		return CapitalMO
	case "CHISINAU":
		return CapitalMD
	case "MAMOUDZOU":
		return CapitalYT
	case "ULAANBAATAR":
		return CapitalMN
	case "PLYMOUTH":
		return CapitalMS
	case "LILONGWE":
		return CapitalMW
	case "PODGORICA":
		return CapitalME
	case "SKOPJE":
		return CapitalMK
	case "BAMAKO":
		return CapitalML
	case "MONACO":
		return CapitalMC
	case "RABAT":
		return CapitalMA
	case "PORTLOUIS":
		return CapitalMU
	case "NOUAKCHOTT":
		return CapitalMR
	case "VALLETTA":
		return CapitalMT
	case "MUSCAT":
		return CapitalOM
	case "MALE":
		return CapitalMV
	case "MEXICOCITY":
		return CapitalMX
	case "KUALALUMPUR":
		return CapitalMY
	case "MAPUTO":
		return CapitalMZ
	case "NOUMEA":
		return CapitalNC
	case "ALOFI":
		return CapitalNU
	case "KINGSTONNORFOLKISLAND", "KINGSTONNORFOLK", "NORFOLKKINGSTON", "NORFOLKISLANDKINGSTON":
		return CapitalNF
	case "NIAMEY":
		return CapitalNE
	case "PORTVILA":
		return CapitalVU
	case "ABUJA":
		return CapitalNG
	case "AMSTERDAM":
		return CapitalNL
	case "PHILIPSBURG":
		return CapitalSX
	case "OSLO":
		return CapitalNO
	case "KATHMANDU":
		return CapitalNP
	case "YAREN":
		return CapitalNR
	case "PARAMARIBO":
		return CapitalSR
	case "WILLEMSTADNETHERLANDSANTILLES", "WILLEMSTADNETHERLANDS", "NETHERLANDSWILLEMSTAD", "NETHERLANDSANTILLESWILLEMSTAD":
		return CapitalAN
	case "MANAGUA":
		return CapitalNI
	case "WELLINGTON":
		return CapitalNZ
	case "JUBA":
		return CapitalSS
	case "PRISTINA", "PRISHTINA", "PRIHTINA", "ПРИШТИНА", "PRISHTINË", "PRIŠTINA":
		return CapitalXK
	case "ASUNCION":
		return CapitalPY
	case "ADAMSTOWN":
		return CapitalPN
	case "LIMA":
		return CapitalPE
	case "ISLAMABAD":
		return CapitalPK
	case "WARSAW":
		return CapitalPL
	case "PANAMACITY":
		return CapitalPA
	case "LISBON":
		return CapitalPT
	case "PORTMORESBY":
		return CapitalPG
	case "MELEKEOK":
		return CapitalPW
	case "BISSAU":
		return CapitalGW
	case "DOHA":
		return CapitalQA
	case "SAINT-DENIS":
		return CapitalRE
	case "BELGRADE":
		return CapitalRS
	case "MAJURO":
		return CapitalMH
	case "MARIGOT":
		return CapitalMF
	case "BUCHAREST":
		return CapitalRO
	case "MANILA":
		return CapitalPH
	case "SANJUAN":
		return CapitalPR
	case "MOSCOW":
		return CapitalRU
	case "KIGALI":
		return CapitalRW
	case "RIYADH":
		return CapitalSA
	case "SAINT-PIERRE":
		return CapitalPM
	case "BASSETERRE":
		return CapitalKN
	case "VICTORIA":
		return CapitalSC
	case "PRETORIA":
		return CapitalZA
	case "DAKAR":
		return CapitalSN
	case "JAMESTOWN":
		return CapitalSH
	case "LJUBLJANA":
		return CapitalSI
	case "FREETOWN":
		return CapitalSL
	case "SANMARINO":
		return CapitalSM
	case "SINGAPORE":
		return CapitalSG
	case "MOGADISHU":
		return CapitalSO
	case "MADRID":
		return CapitalES
	case "CASTRIES":
		return CapitalLC
	case "KHARTOUM":
		return CapitalSD
	case "LONGYEARBYEN":
		return CapitalSJ
	case "STOCKHOLM":
		return CapitalSE
	case "GRYTVIKEN":
		return CapitalGS
	case "DAMASCUS":
		return CapitalSY
	case "BERN":
		return CapitalCH
	case "GUSTAVIA":
		return CapitalBL
	case "PORTOFSPAIN":
		return CapitalTT
	case "BANGKOK":
		return CapitalTH
	case "DUSHANBE":
		return CapitalTJ
	case "COCKBURNTOWN":
		return CapitalTC
	case "NUKUALOFA":
		return CapitalTO
	case "LOME":
		return CapitalTG
	case "SAOTOME":
		return CapitalST
	case "TUNIS":
		return CapitalTN
	case "DILI":
		return CapitalTL
	case "ANKARA":
		return CapitalTR
	case "FUNAFUTI":
		return CapitalTV
	case "TAIPEI":
		return CapitalTW
	case "ASHGABAT":
		return CapitalTM
	case "DODOMA":
		return CapitalTZ
	case "WILLEMSTADCURACAO", "WILLEMSTADCURAKAO", "WILLEMSTADKURAKAO", "CURACAOWILLEMSTAD", "CURAKAOWILLEMSTAD", "KURAKAOWILLEMSTAD":
		return CapitalCW
	case "KAMPALA":
		return CapitalUG
	case "LONDON":
		return CapitalGB
	case "KYIV":
		return CapitalUA
	case "WASHINGTON":
		return CapitalUS
	case "OUAGADOUGOU":
		return CapitalBF
	case "MONTEVIDEO":
		return CapitalUY
	case "TASHKENT":
		return CapitalUZ
	case "KINGSTOWN":
		return CapitalVC
	case "CARACAS":
		return CapitalVE
	case "ROADTOWN":
		return CapitalVG
	case "HANOI":
		return CapitalVN
	case "CHARLOTTEAMALIE":
		return CapitalVI
	case "VATICANCITY":
		return CapitalVA
	case "WINDHOEK":
		return CapitalNA
	case "EASTJERUSALEM":
		return CapitalPS
	case "MATAUTU":
		return CapitalWF
	case "EL-AAIUN":
		return CapitalEH
	case "APIA":
		return CapitalWS
	case "MBABANE":
		return CapitalSZ
	case "SANAA":
		return CapitalYE
	case "LUSAKA":
		return CapitalZM
	case "HARARE":
		return CapitalZW
	}
	return CapitalUnknown
}
