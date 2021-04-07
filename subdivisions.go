package countries

import (
	"encoding/json"
	"fmt"
)

// SubdivisionCode - the code of a subdivision
type SubdivisionCode string

// Subdivision - all info about a subdivision
type Subdivision struct {
	Name    string          `json:"name"`
	Code    SubdivisionCode `json:"code"`
	Country CountryCode     `json:"countryCode"`
}

// Type implements Typer interface
func (s SubdivisionCode) Type() string {
	return TypeSubdivisionCode
}

// String - implements fmt.Stringer, returns an english name of the subdivision
//nolint:cyclop,funlen,gocyclo
func (s SubdivisionCode) String() string { //nolint:cyclop,gocyclo
	switch s {
	case SubdivisionAD02:
		return "Canillo"
	case SubdivisionAD03:
		return "Encamp"
	case SubdivisionAD04:
		return "La Massana"
	case SubdivisionAD05:
		return "Ordino"
	case SubdivisionAD06:
		return "Sant Julia de Loria"
	case SubdivisionAD07:
		return "Andorra la Vella"
	case SubdivisionAD08:
		return "Escaldes-Engordany"
	case SubdivisionAEAJ:
		return "'Ajman"
	case SubdivisionAEAZ:
		return "Abu Zaby"
	case SubdivisionAEDU:
		return "Dubayy"
	case SubdivisionAEFU:
		return "Al Fujayrah"
	case SubdivisionAERK:
		return "Ra's al Khaymah"
	case SubdivisionAESH:
		return "Ash Shariqah"
	case SubdivisionAEUQ:
		return "Umm al Qaywayn"
	case SubdivisionAFBAL:
		return "Balkh"
	case SubdivisionAFBAM:
		return "Bamyan"
	case SubdivisionAFBDG:
		return "Badghis"
	case SubdivisionAFBDS:
		return "Badakhshan"
	case SubdivisionAFBGL:
		return "Baghlan"
	case SubdivisionAFDAY:
		return "Daykundi"
	case SubdivisionAFFRA:
		return "Farah"
	case SubdivisionAFFYB:
		return "Faryab"
	case SubdivisionAFGHA:
		return "Ghazni"
	case SubdivisionAFGHO:
		return "Ghor"
	case SubdivisionAFHEL:
		return "Helmand"
	case SubdivisionAFHER:
		return "Herat"
	case SubdivisionAFJOW:
		return "Jowzjan"
	case SubdivisionAFKAB:
		return "Kabul"
	case SubdivisionAFKAN:
		return "Kandahar"
	case SubdivisionAFKDZ:
		return "Kunduz"
	case SubdivisionAFKHO:
		return "Khost"
	case SubdivisionAFLAG:
		return "Laghman"
	case SubdivisionAFLOG:
		return "Logar"
	case SubdivisionAFNAN:
		return "Nangarhar"
	case SubdivisionAFNIM:
		return "Nimroz"
	case SubdivisionAFPAR:
		return "Parwan"
	case SubdivisionAFPIA:
		return "Paktiya"
	case SubdivisionAFPKA:
		return "Paktika"
	case SubdivisionAFTAK:
		return "Takhar"
	case SubdivisionAFURU:
		return "Uruzgan"
	case SubdivisionAG03:
		return "Saint George"
	case SubdivisionAG04:
		return "Saint John"
	case SubdivisionAG05:
		return "Saint Mary"
	case SubdivisionAG06:
		return "Saint Paul"
	case SubdivisionAG07:
		return "Saint Peter"
	case SubdivisionAG08:
		return "Saint Philip"
	case SubdivisionAG11:
		return "Redonda"
	case SubdivisionAL01:
		return "Berat"
	case SubdivisionAL02:
		return "Durres"
	case SubdivisionAL03:
		return "Elbasan"
	case SubdivisionAL04:
		return "Fier"
	case SubdivisionAL05:
		return "Gjirokaster"
	case SubdivisionAL06:
		return "Korce"
	case SubdivisionAL07:
		return "Kukes"
	case SubdivisionAL08:
		return "Lezhe"
	case SubdivisionAL09:
		return "Diber"
	case SubdivisionAL10:
		return "Shkoder"
	case SubdivisionAL11:
		return "Tirane"
	case SubdivisionAL12:
		return "Vlore"
	case SubdivisionAMAG:
		return "Aragacotn"
	case SubdivisionAMAR:
		return "Ararat"
	case SubdivisionAMAV:
		return "Armavir"
	case SubdivisionAMER:
		return "Erevan"
	case SubdivisionAMGR:
		return "Gegark'unik'"
	case SubdivisionAMKT:
		return "Kotayk'"
	case SubdivisionAMLO:
		return "Lori"
	case SubdivisionAMSH:
		return "Sirak"
	case SubdivisionAMSU:
		return "Syunik'"
	case SubdivisionAMTV:
		return "Tavus"
	case SubdivisionAMVD:
		return "Vayoc Jor"
	case SubdivisionAOBGO:
		return "Bengo"
	case SubdivisionAOBGU:
		return "Benguela"
	case SubdivisionAOBIE:
		return "Bie"
	case SubdivisionAOCAB:
		return "Cabinda"
	case SubdivisionAOCCU:
		return "Cuando Cubango"
	case SubdivisionAOCNN:
		return "Cunene"
	case SubdivisionAOCNO:
		return "Cuanza-Norte"
	case SubdivisionAOCUS:
		return "Cuanza-Sul"
	case SubdivisionAOHUA:
		return "Huambo"
	case SubdivisionAOHUI:
		return "Huila"
	case SubdivisionAOLNO:
		return "Lunda-Norte"
	case SubdivisionAOLSU:
		return "Lunda-Sul"
	case SubdivisionAOLUA:
		return "Luanda"
	case SubdivisionAOMAL:
		return "Malange"
	case SubdivisionAOMOX:
		return "Moxico"
	case SubdivisionAONAM:
		return "Namibe"
	case SubdivisionAOUIG:
		return "Uige"
	case SubdivisionAOZAI:
		return "Zaire"
	case SubdivisionARA:
		return "Salta"
	case SubdivisionARB:
		return "Buenos Aires"
	case SubdivisionARC:
		return "Ciudad Autonoma de Buenos Aires"
	case SubdivisionARD:
		return "San Luis"
	case SubdivisionARE:
		return "Entre Rios"
	case SubdivisionARF:
		return "La Rioja"
	case SubdivisionARG:
		return "Santiago del Estero"
	case SubdivisionARH:
		return "Chaco"
	case SubdivisionARJ:
		return "San Juan"
	case SubdivisionARK:
		return "Catamarca"
	case SubdivisionARL:
		return "La Pampa"
	case SubdivisionARM:
		return "Mendoza"
	case SubdivisionARN:
		return "Misiones"
	case SubdivisionARP:
		return "Formosa"
	case SubdivisionARQ:
		return "Neuquen"
	case SubdivisionARR:
		return "Rio Negro"
	case SubdivisionARS:
		return "Santa Fe"
	case SubdivisionART:
		return "Tucuman"
	case SubdivisionARU:
		return "Chubut"
	case SubdivisionARV:
		return "Tierra del Fuego"
	case SubdivisionARW:
		return "Corrientes"
	case SubdivisionARX:
		return "Cordoba"
	case SubdivisionARY:
		return "Jujuy"
	case SubdivisionARZ:
		return "Santa Cruz"
	case SubdivisionAT1:
		return "Burgenland"
	case SubdivisionAT2:
		return "Karnten"
	case SubdivisionAT3:
		return "Niederosterreich"
	case SubdivisionAT4:
		return "Oberosterreich"
	case SubdivisionAT5:
		return "Salzburg"
	case SubdivisionAT6:
		return "Steiermark"
	case SubdivisionAT7:
		return "Tirol"
	case SubdivisionAT8:
		return "Vorarlberg"
	case SubdivisionAT9:
		return "Wien"
	case SubdivisionAUACT:
		return "Australian Capital Territory"
	case SubdivisionAUNSW:
		return "New South Wales"
	case SubdivisionAUNT:
		return "Northern Territory"
	case SubdivisionAUQLD:
		return "Queensland"
	case SubdivisionAUSA:
		return "South Australia"
	case SubdivisionAUTAS:
		return "Tasmania"
	case SubdivisionAUVIC:
		return "Victoria"
	case SubdivisionAUWA:
		return "Western Australia"
	case SubdivisionAZABS:
		return "Abseron"
	case SubdivisionAZAGC:
		return "Agcabadi"
	case SubdivisionAZAGS:
		return "Agdas"
	case SubdivisionAZAST:
		return "Astara"
	case SubdivisionAZBA:
		return "Baki"
	case SubdivisionAZBAL:
		return "Balakan"
	case SubdivisionAZBAR:
		return "Barda"
	case SubdivisionAZBEY:
		return "Beylaqan"
	case SubdivisionAZBIL:
		return "Bilasuvar"
	case SubdivisionAZCAL:
		return "Calilabad"
	case SubdivisionAZFUZ:
		return "Fuzuli"
	case SubdivisionAZGA:
		return "Ganca"
	case SubdivisionAZGAD:
		return "Gadabay"
	case SubdivisionAZGOR:
		return "Goranboy"
	case SubdivisionAZGOY:
		return "Goycay"
	case SubdivisionAZGYG:
		return "Goygol"
	case SubdivisionAZIMI:
		return "Imisli"
	case SubdivisionAZISM:
		return "Ismayilli"
	case SubdivisionAZKUR:
		return "Kurdamir"
	case SubdivisionAZLA:
		return "Lankaran"
	case SubdivisionAZMAS:
		return "Masalli"
	case SubdivisionAZMI:
		return "Mingacevir"
	case SubdivisionAZNEF:
		return "Neftcala"
	case SubdivisionAZNX:
		return "Naxcivan"
	case SubdivisionAZOGU:
		return "Oguz"
	case SubdivisionAZQAB:
		return "Qabala"
	case SubdivisionAZQAX:
		return "Qax"
	case SubdivisionAZQBA:
		return "Quba"
	case SubdivisionAZQUS:
		return "Qusar"
	case SubdivisionAZSAB:
		return "Sabirabad"
	case SubdivisionAZSAK:
		return "Saki"
	case SubdivisionAZSAL:
		return "Salyan"
	case SubdivisionAZSAT:
		return "Saatli"
	case SubdivisionAZSIY:
		return "Siyazan"
	case SubdivisionAZSKR:
		return "Samkir"
	case SubdivisionAZSM:
		return "Sumqayit"
	case SubdivisionAZSMI:
		return "Samaxi"
	case SubdivisionAZSMX:
		return "Samux"
	case SubdivisionAZSR:
		return "Sirvan"
	case SubdivisionAZTAR:
		return "Tartar"
	case SubdivisionAZXAC:
		return "Xacmaz"
	case SubdivisionAZXIZ:
		return "Xizi"
	case SubdivisionAZYAR:
		return "Yardimli"
	case SubdivisionAZYEV:
		return "Yevlax"
	case SubdivisionAZZAQ:
		return "Zaqatala"
	case SubdivisionAZZAR:
		return "Zardab"
	case SubdivisionBABIH:
		return "Federacija Bosne i Hercegovine"
	case SubdivisionBABRC:
		return "Brcko distrikt"
	case SubdivisionBASRP:
		return "Republika Srpska"
	case SubdivisionBB01:
		return "Christ Church"
	case SubdivisionBB02:
		return "Saint Andrew"
	case SubdivisionBB03:
		return "Saint George"
	case SubdivisionBB04:
		return "Saint James"
	case SubdivisionBB05:
		return "Saint John"
	case SubdivisionBB07:
		return "Saint Lucy"
	case SubdivisionBB08:
		return "Saint Michael"
	case SubdivisionBB09:
		return "Saint Peter"
	case SubdivisionBB10:
		return "Saint Philip"
	case SubdivisionBB11:
		return "Saint Thomas"
	case SubdivisionBDA:
		return "Barishal"
	case SubdivisionBDB:
		return "Chattogram"
	case SubdivisionBDC:
		return "Dhaka"
	case SubdivisionBDD:
		return "Khulna"
	case SubdivisionBDE:
		return "Rajshahi"
	case SubdivisionBDF:
		return "Rangpur"
	case SubdivisionBDG:
		return "Sylhet"
	case SubdivisionBEBRU:
		return "Brussels Hoofdstedelijk Gewest"
	case SubdivisionBEVAN:
		return "Antwerpen"
	case SubdivisionBEVBR:
		return "Vlaams-Brabant"
	case SubdivisionBEVLI:
		return "Limburg"
	case SubdivisionBEVOV:
		return "Oost-Vlaanderen"
	case SubdivisionBEVWV:
		return "West-Vlaanderen"
	case SubdivisionBEWBR:
		return "Brabant wallon"
	case SubdivisionBEWHT:
		return "Hainaut"
	case SubdivisionBEWLG:
		return "Liege"
	case SubdivisionBEWLX:
		return "Luxembourg"
	case SubdivisionBEWNA:
		return "Namur"
	case SubdivisionBFBAM:
		return "Bam"
	case SubdivisionBFBAZ:
		return "Bazega"
	case SubdivisionBFBLG:
		return "Boulgou"
	case SubdivisionBFBLK:
		return "Boulkiemde"
	case SubdivisionBFGAN:
		return "Ganzourgou"
	case SubdivisionBFGNA:
		return "Gnagna"
	case SubdivisionBFGOU:
		return "Gourma"
	case SubdivisionBFHOU:
		return "Houet"
	case SubdivisionBFKAD:
		return "Kadiogo"
	case SubdivisionBFKMD:
		return "Komondjari"
	case SubdivisionBFKMP:
		return "Kompienga"
	case SubdivisionBFKOP:
		return "Koulpelogo"
	case SubdivisionBFKOT:
		return "Kouritenga"
	case SubdivisionBFKOW:
		return "Kourweogo"
	case SubdivisionBFLER:
		return "Leraba"
	case SubdivisionBFLOR:
		return "Loroum"
	case SubdivisionBFMOU:
		return "Mouhoun"
	case SubdivisionBFNAM:
		return "Namentenga"
	case SubdivisionBFNAO:
		return "Nahouri"
	case SubdivisionBFNAY:
		return "Nayala"
	case SubdivisionBFOUB:
		return "Oubritenga"
	case SubdivisionBFOUD:
		return "Oudalan"
	case SubdivisionBFPAS:
		return "Passore"
	case SubdivisionBFPON:
		return "Poni"
	case SubdivisionBFSEN:
		return "Seno"
	case SubdivisionBFSIS:
		return "Sissili"
	case SubdivisionBFSMT:
		return "Sanmatenga"
	case SubdivisionBFSOM:
		return "Soum"
	case SubdivisionBFSOR:
		return "Sourou"
	case SubdivisionBFTAP:
		return "Tapoa"
	case SubdivisionBFTUI:
		return "Tuy"
	case SubdivisionBFYAT:
		return "Yatenga"
	case SubdivisionBFZIR:
		return "Ziro"
	case SubdivisionBFZON:
		return "Zondoma"
	case SubdivisionBFZOU:
		return "Zoundweogo"
	case SubdivisionBG01:
		return "Blagoevgrad"
	case SubdivisionBG02:
		return "Burgas"
	case SubdivisionBG03:
		return "Varna"
	case SubdivisionBG04:
		return "Veliko Tarnovo"
	case SubdivisionBG05:
		return "Vidin"
	case SubdivisionBG06:
		return "Vratsa"
	case SubdivisionBG07:
		return "Gabrovo"
	case SubdivisionBG08:
		return "Dobrich"
	case SubdivisionBG09:
		return "Kardzhali"
	case SubdivisionBG10:
		return "Kyustendil"
	case SubdivisionBG11:
		return "Lovech"
	case SubdivisionBG12:
		return "Montana"
	case SubdivisionBG13:
		return "Pazardzhik"
	case SubdivisionBG14:
		return "Pernik"
	case SubdivisionBG15:
		return "Pleven"
	case SubdivisionBG16:
		return "Plovdiv"
	case SubdivisionBG17:
		return "Razgrad"
	case SubdivisionBG18:
		return "Ruse"
	case SubdivisionBG19:
		return "Silistra"
	case SubdivisionBG20:
		return "Sliven"
	case SubdivisionBG21:
		return "Smolyan"
	case SubdivisionBG22:
		return "Sofia (stolitsa)"
	case SubdivisionBG23:
		return "Sofia"
	case SubdivisionBG24:
		return "Stara Zagora"
	case SubdivisionBG25:
		return "Targovishte"
	case SubdivisionBG26:
		return "Haskovo"
	case SubdivisionBG27:
		return "Shumen"
	case SubdivisionBG28:
		return "Yambol"
	case SubdivisionBH13:
		return "Al 'Asimah"
	case SubdivisionBH14:
		return "Al Janubiyah"
	case SubdivisionBH15:
		return "Al Muharraq"
	case SubdivisionBH17:
		return "Ash Shamaliyah"
	case SubdivisionBIBM:
		return "Bujumbura Mairie"
	case SubdivisionBICI:
		return "Cibitoke"
	case SubdivisionBIGI:
		return "Gitega"
	case SubdivisionBIKI:
		return "Kirundo"
	case SubdivisionBIMW:
		return "Mwaro"
	case SubdivisionBING:
		return "Ngozi"
	case SubdivisionBIRM:
		return "Rumonge"
	case SubdivisionBIRT:
		return "Rutana"
	case SubdivisionBIRY:
		return "Ruyigi"
	case SubdivisionBJAK:
		return "Atacora"
	case SubdivisionBJAL:
		return "Alibori"
	case SubdivisionBJAQ:
		return "Atlantique"
	case SubdivisionBJBO:
		return "Borgou"
	case SubdivisionBJCO:
		return "Collines"
	case SubdivisionBJDO:
		return "Donga"
	case SubdivisionBJLI:
		return "Littoral"
	case SubdivisionBJMO:
		return "Mono"
	case SubdivisionBJOU:
		return "Oueme"
	case SubdivisionBJPL:
		return "Plateau"
	case SubdivisionBJZO:
		return "Zou"
	case SubdivisionBNBE:
		return "Belait"
	case SubdivisionBNBM:
		return "Brunei-Muara"
	case SubdivisionBNTE:
		return "Temburong"
	case SubdivisionBNTU:
		return "Tutong"
	case SubdivisionBOB:
		return "El Beni"
	case SubdivisionBOC:
		return "Cochabamba"
	case SubdivisionBOH:
		return "Chuquisaca"
	case SubdivisionBOL:
		return "La Paz"
	case SubdivisionBON:
		return "Pando"
	case SubdivisionBOO:
		return "Oruro"
	case SubdivisionBOP:
		return "Potosi"
	case SubdivisionBOS:
		return "Santa Cruz"
	case SubdivisionBOT:
		return "Tarija"
	case SubdivisionBQBO:
		return "Bonaire"
	case SubdivisionBQSA:
		return "Saba"
	case SubdivisionBQSE:
		return "Sint Eustatius"
	case SubdivisionBRAC:
		return "Acre"
	case SubdivisionBRAL:
		return "Alagoas"
	case SubdivisionBRAM:
		return "Amazonas"
	case SubdivisionBRAP:
		return "Amapa"
	case SubdivisionBRBA:
		return "Bahia"
	case SubdivisionBRCE:
		return "Ceara"
	case SubdivisionBRDF:
		return "Distrito Federal"
	case SubdivisionBRES:
		return "Espirito Santo"
	case SubdivisionBRGO:
		return "Goias"
	case SubdivisionBRMA:
		return "Maranhao"
	case SubdivisionBRMG:
		return "Minas Gerais"
	case SubdivisionBRMS:
		return "Mato Grosso do Sul"
	case SubdivisionBRMT:
		return "Mato Grosso"
	case SubdivisionBRPA:
		return "Para"
	case SubdivisionBRPB:
		return "Paraiba"
	case SubdivisionBRPE:
		return "Pernambuco"
	case SubdivisionBRPI:
		return "Piaui"
	case SubdivisionBRPR:
		return "Parana"
	case SubdivisionBRRJ:
		return "Rio de Janeiro"
	case SubdivisionBRRN:
		return "Rio Grande do Norte"
	case SubdivisionBRRO:
		return "Rondonia"
	case SubdivisionBRRR:
		return "Roraima"
	case SubdivisionBRRS:
		return "Rio Grande do Sul"
	case SubdivisionBRSC:
		return "Santa Catarina"
	case SubdivisionBRSE:
		return "Sergipe"
	case SubdivisionBRSP:
		return "Sao Paulo"
	case SubdivisionBRTO:
		return "Tocantins"
	case SubdivisionBSCO:
		return "Central Abaco"
	case SubdivisionBSFP:
		return "City of Freeport"
	case SubdivisionBSLI:
		return "Long Island"
	case SubdivisionBSNO:
		return "North Abaco"
	case SubdivisionBSNP:
		return "New Providence"
	case SubdivisionBSNS:
		return "North Andros"
	case SubdivisionBSSE:
		return "South Eleuthera"
	case SubdivisionBT11:
		return "Paro"
	case SubdivisionBT12:
		return "Chhukha"
	case SubdivisionBT13:
		return "Haa"
	case SubdivisionBT14:
		return "Samtse"
	case SubdivisionBT15:
		return "Thimphu"
	case SubdivisionBT21:
		return "Tsirang"
	case SubdivisionBT23:
		return "Punakha"
	case SubdivisionBT24:
		return "Wangdue Phodrang"
	case SubdivisionBT32:
		return "Trongsa"
	case SubdivisionBT33:
		return "Bumthang"
	case SubdivisionBT41:
		return "Trashigang"
	case SubdivisionBT42:
		return "Monggar"
	case SubdivisionBT43:
		return "Pema Gatshel"
	case SubdivisionBT44:
		return "Lhuentse"
	case SubdivisionBT45:
		return "Samdrup Jongkhar"
	case SubdivisionBTGA:
		return "Gasa"
	case SubdivisionBWCE:
		return "Central"
	case SubdivisionBWCH:
		return "Chobe"
	case SubdivisionBWKG:
		return "Kgalagadi"
	case SubdivisionBWKL:
		return "Kgatleng"
	case SubdivisionBWKW:
		return "Kweneng"
	case SubdivisionBWNE:
		return "North East"
	case SubdivisionBWNW:
		return "North West"
	case SubdivisionBWSE:
		return "South East"
	case SubdivisionBWSO:
		return "Southern"
	case SubdivisionBYBR:
		return "Brestskaya voblasts'"
	case SubdivisionBYHM:
		return "Horad Minsk"
	case SubdivisionBYHO:
		return "Homyel'skaya voblasts'"
	case SubdivisionBYHR:
		return "Hrodzyenskaya voblasts'"
	case SubdivisionBYMA:
		return "Mahilyowskaya voblasts'"
	case SubdivisionBYMI:
		return "Minskaya voblasts'"
	case SubdivisionBYVI:
		return "Vitsyebskaya voblasts'"
	case SubdivisionBZBZ:
		return "Belize"
	case SubdivisionBZCY:
		return "Cayo"
	case SubdivisionBZCZL:
		return "Corozal"
	case SubdivisionBZOW:
		return "Orange Walk"
	case SubdivisionBZSC:
		return "Stann Creek"
	case SubdivisionBZTOL:
		return "Toledo"
	case SubdivisionCAAB:
		return "Alberta"
	case SubdivisionCABC:
		return "British Columbia"
	case SubdivisionCAMB:
		return "Manitoba"
	case SubdivisionCANB:
		return "New Brunswick"
	case SubdivisionCANL:
		return "Newfoundland and Labrador"
	case SubdivisionCANS:
		return "Nova Scotia"
	case SubdivisionCANT:
		return "Northwest Territories"
	case SubdivisionCANU:
		return "Nunavut"
	case SubdivisionCAON:
		return "Ontario"
	case SubdivisionCAPE:
		return "Prince Edward Island"
	case SubdivisionCAQC:
		return "Quebec"
	case SubdivisionCASK:
		return "Saskatchewan"
	case SubdivisionCAYT:
		return "Yukon"
	case SubdivisionCDEQ:
		return "Equateur"
	case SubdivisionCDHK:
		return "Haut-Katanga"
	case SubdivisionCDIT:
		return "Ituri"
	case SubdivisionCDKC:
		return "Kasai Central"
	case SubdivisionCDKE:
		return "Kasai Oriental"
	case SubdivisionCDKL:
		return "Kwilu"
	case SubdivisionCDKN:
		return "Kinshasa"
	case SubdivisionCDLU:
		return "Lualaba"
	case SubdivisionCDNK:
		return "Nord-Kivu"
	case SubdivisionCDSA:
		return "Sankuru"
	case SubdivisionCDSK:
		return "Sud-Kivu"
	case SubdivisionCDTA:
		return "Tanganyika"
	case SubdivisionCDTO:
		return "Tshopo"
	case SubdivisionCFAC:
		return "Ouham"
	case SubdivisionCFBGF:
		return "Bangui"
	case SubdivisionCFNM:
		return "Nana-Mambere"
	case SubdivisionCFOP:
		return "Ouham-Pende"
	case SubdivisionCG13:
		return "Sangha"
	case SubdivisionCG16:
		return "Pointe-Noire"
	case SubdivisionCGBZV:
		return "Brazzaville"
	case SubdivisionCHAG:
		return "Aargau"
	case SubdivisionCHAI:
		return "Appenzell Innerrhoden"
	case SubdivisionCHAR:
		return "Appenzell Ausserrhoden"
	case SubdivisionCHBE:
		return "Bern"
	case SubdivisionCHBL:
		return "Basel-Landschaft"
	case SubdivisionCHBS:
		return "Basel-Stadt"
	case SubdivisionCHFR:
		return "Fribourg"
	case SubdivisionCHGE:
		return "Geneve"
	case SubdivisionCHGL:
		return "Glarus"
	case SubdivisionCHGR:
		return "Graubunden"
	case SubdivisionCHJU:
		return "Jura"
	case SubdivisionCHLU:
		return "Luzern"
	case SubdivisionCHNE:
		return "Neuchatel"
	case SubdivisionCHNW:
		return "Nidwalden"
	case SubdivisionCHOW:
		return "Obwalden"
	case SubdivisionCHSG:
		return "Sankt Gallen"
	case SubdivisionCHSH:
		return "Schaffhausen"
	case SubdivisionCHSO:
		return "Solothurn"
	case SubdivisionCHSZ:
		return "Schwyz"
	case SubdivisionCHTG:
		return "Thurgau"
	case SubdivisionCHTI:
		return "Ticino"
	case SubdivisionCHUR:
		return "Uri"
	case SubdivisionCHVD:
		return "Vaud"
	case SubdivisionCHVS:
		return "Valais"
	case SubdivisionCHZG:
		return "Zug"
	case SubdivisionCHZH:
		return "Zurich"
	case SubdivisionCIAB:
		return "Abidjan"
	case SubdivisionCIBS:
		return "Bas-Sassandra"
	case SubdivisionCICM:
		return "Comoe"
	case SubdivisionCIDN:
		return "Denguele"
	case SubdivisionCIGD:
		return "Goh-Djiboua"
	case SubdivisionCILC:
		return "Lacs"
	case SubdivisionCILG:
		return "Lagunes"
	case SubdivisionCIMG:
		return "Montagnes"
	case SubdivisionCISM:
		return "Sassandra-Marahoue"
	case SubdivisionCISV:
		return "Savanes"
	case SubdivisionCIVB:
		return "Vallee du Bandama"
	case SubdivisionCIWR:
		return "Woroba"
	case SubdivisionCIYM:
		return "Yamoussoukro"
	case SubdivisionCIZZ:
		return "Zanzan"
	case SubdivisionCLAI:
		return "Aisen del General Carlos Ibanez del Campo"
	case SubdivisionCLAN:
		return "Antofagasta"
	case SubdivisionCLAP:
		return "Arica y Parinacota"
	case SubdivisionCLAR:
		return "La Araucania"
	case SubdivisionCLAT:
		return "Atacama"
	case SubdivisionCLBI:
		return "Biobio"
	case SubdivisionCLCO:
		return "Coquimbo"
	case SubdivisionCLLI:
		return "Libertador General Bernardo O'Higgins"
	case SubdivisionCLLL:
		return "Los Lagos"
	case SubdivisionCLLR:
		return "Los Rios"
	case SubdivisionCLMA:
		return "Magallanes"
	case SubdivisionCLML:
		return "Maule"
	case SubdivisionCLNB:
		return "Nuble"
	case SubdivisionCLRM:
		return "Region Metropolitana de Santiago"
	case SubdivisionCLTA:
		return "Tarapaca"
	case SubdivisionCLVS:
		return "Valparaiso"
	case SubdivisionCMAD:
		return "Adamaoua"
	case SubdivisionCMCE:
		return "Centre"
	case SubdivisionCMEN:
		return "Extreme-Nord"
	case SubdivisionCMES:
		return "Est"
	case SubdivisionCMLT:
		return "Littoral"
	case SubdivisionCMNO:
		return "Nord"
	case SubdivisionCMNW:
		return "Nord-Ouest"
	case SubdivisionCMOU:
		return "Ouest"
	case SubdivisionCMSU:
		return "Sud"
	case SubdivisionCMSW:
		return "Sud-Ouest"
	case SubdivisionCNAH:
		return "Anhui"
	case SubdivisionCNBJ:
		return "Beijing"
	case SubdivisionCNCQ:
		return "Chongqing"
	case SubdivisionCNFJ:
		return "Fujian"
	case SubdivisionCNGD:
		return "Guangdong"
	case SubdivisionCNGS:
		return "Gansu"
	case SubdivisionCNGX:
		return "Guangxi"
	case SubdivisionCNGZ:
		return "Guizhou"
	case SubdivisionCNHA:
		return "Henan"
	case SubdivisionCNHB:
		return "Hubei"
	case SubdivisionCNHE:
		return "Hebei"
	case SubdivisionCNHI:
		return "Hainan"
	case SubdivisionCNHL:
		return "Heilongjiang"
	case SubdivisionCNHN:
		return "Hunan"
	case SubdivisionCNJL:
		return "Jilin"
	case SubdivisionCNJS:
		return "Jiangsu"
	case SubdivisionCNJX:
		return "Jiangxi"
	case SubdivisionCNLN:
		return "Liaoning"
	case SubdivisionCNNM:
		return "Nei Mongol"
	case SubdivisionCNNX:
		return "Ningxia"
	case SubdivisionCNQH:
		return "Qinghai"
	case SubdivisionCNSC:
		return "Sichuan"
	case SubdivisionCNSD:
		return "Shandong"
	case SubdivisionCNSH:
		return "Shanghai"
	case SubdivisionCNSN:
		return "Shaanxi"
	case SubdivisionCNSX:
		return "Shanxi"
	case SubdivisionCNTJ:
		return "Tianjin"
	case SubdivisionCNXJ:
		return "Xinjiang"
	case SubdivisionCNXZ:
		return "Xizang"
	case SubdivisionCNYN:
		return "Yunnan"
	case SubdivisionCNZJ:
		return "Zhejiang"
	case SubdivisionCOAMA:
		return "Amazonas"
	case SubdivisionCOANT:
		return "Antioquia"
	case SubdivisionCOARA:
		return "Arauca"
	case SubdivisionCOATL:
		return "Atlantico"
	case SubdivisionCOBOL:
		return "Bolivar"
	case SubdivisionCOBOY:
		return "Boyaca"
	case SubdivisionCOCAL:
		return "Caldas"
	case SubdivisionCOCAQ:
		return "Caqueta"
	case SubdivisionCOCAS:
		return "Casanare"
	case SubdivisionCOCAU:
		return "Cauca"
	case SubdivisionCOCES:
		return "Cesar"
	case SubdivisionCOCHO:
		return "Choco"
	case SubdivisionCOCOR:
		return "Cordoba"
	case SubdivisionCOCUN:
		return "Cundinamarca"
	case SubdivisionCODC:
		return "Distrito Capital de Bogota"
	case SubdivisionCOGUA:
		return "Guainia"
	case SubdivisionCOGUV:
		return "Guaviare"
	case SubdivisionCOHUI:
		return "Huila"
	case SubdivisionCOLAG:
		return "La Guajira"
	case SubdivisionCOMAG:
		return "Magdalena"
	case SubdivisionCOMET:
		return "Meta"
	case SubdivisionCONAR:
		return "Narino"
	case SubdivisionCONSA:
		return "Norte de Santander"
	case SubdivisionCOPUT:
		return "Putumayo"
	case SubdivisionCOQUI:
		return "Quindio"
	case SubdivisionCORIS:
		return "Risaralda"
	case SubdivisionCOSAN:
		return "Santander"
	case SubdivisionCOSAP:
		return "San Andres, Providencia y Santa Catalina"
	case SubdivisionCOSUC:
		return "Sucre"
	case SubdivisionCOTOL:
		return "Tolima"
	case SubdivisionCOVAC:
		return "Valle del Cauca"
	case SubdivisionCOVID:
		return "Vichada"
	case SubdivisionCRA:
		return "Alajuela"
	case SubdivisionCRC:
		return "Cartago"
	case SubdivisionCRG:
		return "Guanacaste"
	case SubdivisionCRH:
		return "Heredia"
	case SubdivisionCRL:
		return "Limon"
	case SubdivisionCRP:
		return "Puntarenas"
	case SubdivisionCRSJ:
		return "San Jose"
	case SubdivisionCU01:
		return "Pinar del Rio"
	case SubdivisionCU03:
		return "La Habana"
	case SubdivisionCU04:
		return "Matanzas"
	case SubdivisionCU05:
		return "Villa Clara"
	case SubdivisionCU06:
		return "Cienfuegos"
	case SubdivisionCU07:
		return "Sancti Spiritus"
	case SubdivisionCU08:
		return "Ciego de Avila"
	case SubdivisionCU09:
		return "Camaguey"
	case SubdivisionCU10:
		return "Las Tunas"
	case SubdivisionCU11:
		return "Holguin"
	case SubdivisionCU12:
		return "Granma"
	case SubdivisionCU13:
		return "Santiago de Cuba"
	case SubdivisionCU14:
		return "Guantanamo"
	case SubdivisionCU15:
		return "Artemisa"
	case SubdivisionCU16:
		return "Mayabeque"
	case SubdivisionCVBR:
		return "Brava"
	case SubdivisionCVBV:
		return "Boa Vista"
	case SubdivisionCVCR:
		return "Santa Cruz"
	case SubdivisionCVMA:
		return "Maio"
	case SubdivisionCVPN:
		return "Porto Novo"
	case SubdivisionCVPR:
		return "Praia"
	case SubdivisionCVRG:
		return "Ribeira Grande"
	case SubdivisionCVRS:
		return "Ribeira Grande de Santiago"
	case SubdivisionCVSF:
		return "Sao Filipe"
	case SubdivisionCVSL:
		return "Sal"
	case SubdivisionCVSV:
		return "Sao Vicente"
	case SubdivisionCVTA:
		return "Tarrafal"
	case SubdivisionCY01:
		return "Lefkosia"
	case SubdivisionCY02:
		return "Lemesos"
	case SubdivisionCY03:
		return "Larnaka"
	case SubdivisionCY04:
		return "Ammochostos"
	case SubdivisionCY05:
		return "Pafos"
	case SubdivisionCY06:
		return "Keryneia"
	case SubdivisionCZ10:
		return "Praha, Hlavni mesto"
	case SubdivisionCZ20:
		return "Stredocesky kraj"
	case SubdivisionCZ31:
		return "Jihocesky kraj"
	case SubdivisionCZ32:
		return "Plzensky kraj"
	case SubdivisionCZ41:
		return "Karlovarsky kraj"
	case SubdivisionCZ42:
		return "Ustecky kraj"
	case SubdivisionCZ51:
		return "Liberecky kraj"
	case SubdivisionCZ52:
		return "Kralovehradecky kraj"
	case SubdivisionCZ53:
		return "Pardubicky kraj"
	case SubdivisionCZ63:
		return "Kraj Vysocina"
	case SubdivisionCZ64:
		return "Jihomoravsky kraj"
	case SubdivisionCZ71:
		return "Olomoucky kraj"
	case SubdivisionCZ72:
		return "Zlinsky kraj"
	case SubdivisionCZ80:
		return "Moravskoslezsky kraj"
	case SubdivisionDEBB:
		return "Brandenburg"
	case SubdivisionDEBE:
		return "Berlin"
	case SubdivisionDEBW:
		return "Baden-Wurttemberg"
	case SubdivisionDEBY:
		return "Bayern"
	case SubdivisionDEHB:
		return "Bremen"
	case SubdivisionDEHE:
		return "Hessen"
	case SubdivisionDEHH:
		return "Hamburg"
	case SubdivisionDEMV:
		return "Mecklenburg-Vorpommern"
	case SubdivisionDENI:
		return "Niedersachsen"
	case SubdivisionDENW:
		return "Nordrhein-Westfalen"
	case SubdivisionDERP:
		return "Rheinland-Pfalz"
	case SubdivisionDESH:
		return "Schleswig-Holstein"
	case SubdivisionDESL:
		return "Saarland"
	case SubdivisionDESN:
		return "Sachsen"
	case SubdivisionDEST:
		return "Sachsen-Anhalt"
	case SubdivisionDETH:
		return "Thuringen"
	case SubdivisionDJDJ:
		return "Djibouti"
	case SubdivisionDK81:
		return "Nordjylland"
	case SubdivisionDK82:
		return "Midtjylland"
	case SubdivisionDK83:
		return "Syddanmark"
	case SubdivisionDK84:
		return "Hovedstaden"
	case SubdivisionDK85:
		return "Sjaelland"
	case SubdivisionDM02:
		return "Saint Andrew"
	case SubdivisionDM04:
		return "Saint George"
	case SubdivisionDM05:
		return "Saint John"
	case SubdivisionDM09:
		return "Saint Patrick"
	case SubdivisionDM10:
		return "Saint Paul"
	case SubdivisionDO01:
		return "Distrito Nacional (Santo Domingo)"
	case SubdivisionDO02:
		return "Azua"
	case SubdivisionDO03:
		return "Baoruco"
	case SubdivisionDO04:
		return "Barahona"
	case SubdivisionDO05:
		return "Dajabon"
	case SubdivisionDO06:
		return "Duarte"
	case SubdivisionDO07:
		return "Elias Pina"
	case SubdivisionDO08:
		return "El Seibo"
	case SubdivisionDO09:
		return "Espaillat"
	case SubdivisionDO10:
		return "Independencia"
	case SubdivisionDO11:
		return "La Altagracia"
	case SubdivisionDO12:
		return "La Romana"
	case SubdivisionDO13:
		return "La Vega"
	case SubdivisionDO14:
		return "Maria Trinidad Sanchez"
	case SubdivisionDO15:
		return "Monte Cristi"
	case SubdivisionDO17:
		return "Peravia"
	case SubdivisionDO18:
		return "Puerto Plata"
	case SubdivisionDO19:
		return "Hermanas Mirabal"
	case SubdivisionDO20:
		return "Samana"
	case SubdivisionDO21:
		return "San Cristobal"
	case SubdivisionDO22:
		return "San Juan"
	case SubdivisionDO23:
		return "San Pedro de Macoris"
	case SubdivisionDO24:
		return "Sanchez Ramirez"
	case SubdivisionDO25:
		return "Santiago"
	case SubdivisionDO26:
		return "Santiago Rodriguez"
	case SubdivisionDO27:
		return "Valverde"
	case SubdivisionDO28:
		return "Monsenor Nouel"
	case SubdivisionDO29:
		return "Monte Plata"
	case SubdivisionDO30:
		return "Hato Mayor"
	case SubdivisionDO31:
		return "San Jose de Ocoa"
	case SubdivisionDZ01:
		return "Adrar"
	case SubdivisionDZ02:
		return "Chlef"
	case SubdivisionDZ03:
		return "Laghouat"
	case SubdivisionDZ04:
		return "Oum el Bouaghi"
	case SubdivisionDZ05:
		return "Batna"
	case SubdivisionDZ06:
		return "Bejaia"
	case SubdivisionDZ07:
		return "Biskra"
	case SubdivisionDZ08:
		return "Bechar"
	case SubdivisionDZ09:
		return "Blida"
	case SubdivisionDZ10:
		return "Bouira"
	case SubdivisionDZ11:
		return "Tamanrasset"
	case SubdivisionDZ12:
		return "Tebessa"
	case SubdivisionDZ13:
		return "Tlemcen"
	case SubdivisionDZ14:
		return "Tiaret"
	case SubdivisionDZ15:
		return "Tizi Ouzou"
	case SubdivisionDZ16:
		return "Alger"
	case SubdivisionDZ17:
		return "Djelfa"
	case SubdivisionDZ18:
		return "Jijel"
	case SubdivisionDZ19:
		return "Setif"
	case SubdivisionDZ20:
		return "Saida"
	case SubdivisionDZ21:
		return "Skikda"
	case SubdivisionDZ22:
		return "Sidi Bel Abbes"
	case SubdivisionDZ23:
		return "Annaba"
	case SubdivisionDZ24:
		return "Guelma"
	case SubdivisionDZ25:
		return "Constantine"
	case SubdivisionDZ26:
		return "Medea"
	case SubdivisionDZ27:
		return "Mostaganem"
	case SubdivisionDZ28:
		return "M'sila"
	case SubdivisionDZ29:
		return "Mascara"
	case SubdivisionDZ30:
		return "Ouargla"
	case SubdivisionDZ31:
		return "Oran"
	case SubdivisionDZ32:
		return "El Bayadh"
	case SubdivisionDZ33:
		return "Illizi"
	case SubdivisionDZ34:
		return "Bordj Bou Arreridj"
	case SubdivisionDZ35:
		return "Boumerdes"
	case SubdivisionDZ36:
		return "El Tarf"
	case SubdivisionDZ37:
		return "Tindouf"
	case SubdivisionDZ38:
		return "Tissemsilt"
	case SubdivisionDZ39:
		return "El Oued"
	case SubdivisionDZ40:
		return "Khenchela"
	case SubdivisionDZ41:
		return "Souk Ahras"
	case SubdivisionDZ42:
		return "Tipaza"
	case SubdivisionDZ43:
		return "Mila"
	case SubdivisionDZ44:
		return "Ain Defla"
	case SubdivisionDZ45:
		return "Naama"
	case SubdivisionDZ46:
		return "Ain Temouchent"
	case SubdivisionDZ47:
		return "Ghardaia"
	case SubdivisionDZ48:
		return "Relizane"
	case SubdivisionECA:
		return "Azuay"
	case SubdivisionECB:
		return "Bolivar"
	case SubdivisionECC:
		return "Carchi"
	case SubdivisionECD:
		return "Orellana"
	case SubdivisionECE:
		return "Esmeraldas"
	case SubdivisionECF:
		return "Canar"
	case SubdivisionECG:
		return "Guayas"
	case SubdivisionECH:
		return "Chimborazo"
	case SubdivisionECI:
		return "Imbabura"
	case SubdivisionECL:
		return "Loja"
	case SubdivisionECM:
		return "Manabi"
	case SubdivisionECN:
		return "Napo"
	case SubdivisionECO:
		return "El Oro"
	case SubdivisionECP:
		return "Pichincha"
	case SubdivisionECR:
		return "Los Rios"
	case SubdivisionECS:
		return "Morona Santiago"
	case SubdivisionECSD:
		return "Santo Domingo de los Tsachilas"
	case SubdivisionECSE:
		return "Santa Elena"
	case SubdivisionECT:
		return "Tungurahua"
	case SubdivisionECU:
		return "Sucumbios"
	case SubdivisionECW:
		return "Galapagos"
	case SubdivisionECX:
		return "Cotopaxi"
	case SubdivisionECY:
		return "Pastaza"
	case SubdivisionECZ:
		return "Zamora Chinchipe"
	case SubdivisionEE37:
		return "Harjumaa"
	case SubdivisionEE39:
		return "Hiiumaa"
	case SubdivisionEE44:
		return "Ida-Virumaa"
	case SubdivisionEE49:
		return "Jogevamaa"
	case SubdivisionEE51:
		return "Jarvamaa"
	case SubdivisionEE57:
		return "Laanemaa"
	case SubdivisionEE59:
		return "Laane-Virumaa"
	case SubdivisionEE65:
		return "Polvamaa"
	case SubdivisionEE67:
		return "Parnumaa"
	case SubdivisionEE70:
		return "Raplamaa"
	case SubdivisionEE74:
		return "Saaremaa"
	case SubdivisionEE78:
		return "Tartumaa"
	case SubdivisionEE82:
		return "Valgamaa"
	case SubdivisionEE84:
		return "Viljandimaa"
	case SubdivisionEE86:
		return "Vorumaa"
	case SubdivisionEGALX:
		return "Al Iskandariyah"
	case SubdivisionEGASN:
		return "Aswan"
	case SubdivisionEGAST:
		return "Asyut"
	case SubdivisionEGBA:
		return "Al Bahr al Ahmar"
	case SubdivisionEGBH:
		return "Al Buhayrah"
	case SubdivisionEGBNS:
		return "Bani Suwayf"
	case SubdivisionEGC:
		return "Al Qahirah"
	case SubdivisionEGDK:
		return "Ad Daqahliyah"
	case SubdivisionEGDT:
		return "Dumyat"
	case SubdivisionEGFYM:
		return "Al Fayyum"
	case SubdivisionEGGH:
		return "Al Gharbiyah"
	case SubdivisionEGGZ:
		return "Al Jizah"
	case SubdivisionEGIS:
		return "Al Isma'iliyah"
	case SubdivisionEGJS:
		return "Janub Sina'"
	case SubdivisionEGKB:
		return "Al Qalyubiyah"
	case SubdivisionEGKFS:
		return "Kafr ash Shaykh"
	case SubdivisionEGKN:
		return "Qina"
	case SubdivisionEGLX:
		return "Al Uqsur"
	case SubdivisionEGMN:
		return "Al Minya"
	case SubdivisionEGMNF:
		return "Al Minufiyah"
	case SubdivisionEGMT:
		return "Matruh"
	case SubdivisionEGPTS:
		return "Bur Sa'id"
	case SubdivisionEGSHG:
		return "Suhaj"
	case SubdivisionEGSHR:
		return "Ash Sharqiyah"
	case SubdivisionEGSIN:
		return "Shamal Sina'"
	case SubdivisionEGSUZ:
		return "As Suways"
	case SubdivisionEGWAD:
		return "Al Wadi al Jadid"
	case SubdivisionERMA:
		return "Al Awsat"
	case SubdivisionESAN:
		return "Andalucia"
	case SubdivisionESAR:
		return "Aragon"
	case SubdivisionESAS:
		return "Asturias, Principado de"
	case SubdivisionESCB:
		return "Cantabria"
	case SubdivisionESCE:
		return "Ceuta"
	case SubdivisionESCL:
		return "Castilla y Leon"
	case SubdivisionESCM:
		return "Castilla-La Mancha"
	case SubdivisionESCN:
		return "Canarias"
	case SubdivisionESCT:
		return "Catalunya"
	case SubdivisionESEX:
		return "Extremadura"
	case SubdivisionESGA:
		return "Galicia"
	case SubdivisionESIB:
		return "Illes Balears"
	case SubdivisionESMC:
		return "Murcia, Region de"
	case SubdivisionESMD:
		return "Madrid, Comunidad de"
	case SubdivisionESML:
		return "Melilla"
	case SubdivisionESNC:
		return "Navarra, Comunidad Foral de"
	case SubdivisionESPV:
		return "Pais Vasco"
	case SubdivisionESRI:
		return "La Rioja"
	case SubdivisionESVC:
		return "Valenciana, Comunidad"
	case SubdivisionETAA:
		return "Adis Abeba"
	case SubdivisionETAF:
		return "Afar"
	case SubdivisionETAM:
		return "Amara"
	case SubdivisionETBE:
		return "Binshangul Gumuz"
	case SubdivisionETDD:
		return "Dire Dawa"
	case SubdivisionETHA:
		return "Hareri Hizb"
	case SubdivisionETOR:
		return "Oromiya"
	case SubdivisionETSN:
		return "YeDebub Biheroch Bihereseboch na Hizboch"
	case SubdivisionETSO:
		return "Sumale"
	case SubdivisionETTI:
		return "Tigray"
	case SubdivisionFI02:
		return "Etela-Karjala"
	case SubdivisionFI03:
		return "Etela-Pohjanmaa"
	case SubdivisionFI04:
		return "Etela-Savo"
	case SubdivisionFI05:
		return "Kainuu"
	case SubdivisionFI06:
		return "Kanta-Hame"
	case SubdivisionFI07:
		return "Keski-Pohjanmaa"
	case SubdivisionFI08:
		return "Keski-Suomi"
	case SubdivisionFI09:
		return "Kymenlaakso"
	case SubdivisionFI10:
		return "Lappi"
	case SubdivisionFI11:
		return "Pirkanmaa"
	case SubdivisionFI12:
		return "Pohjanmaa"
	case SubdivisionFI13:
		return "Pohjois-Karjala"
	case SubdivisionFI14:
		return "Pohjois-Pohjanmaa"
	case SubdivisionFI15:
		return "Pohjois-Savo"
	case SubdivisionFI16:
		return "Paijat-Hame"
	case SubdivisionFI17:
		return "Satakunta"
	case SubdivisionFI18:
		return "Uusimaa"
	case SubdivisionFI19:
		return "Varsinais-Suomi"
	case SubdivisionFJC:
		return "Central"
	case SubdivisionFJE:
		return "Eastern"
	case SubdivisionFJN:
		return "Northern"
	case SubdivisionFJR:
		return "Rotuma"
	case SubdivisionFJW:
		return "Western"
	case SubdivisionFMKSA:
		return "Kosrae"
	case SubdivisionFMPNI:
		return "Pohnpei"
	case SubdivisionFMTRK:
		return "Chuuk"
	case SubdivisionFMYAP:
		return "Yap"
	case SubdivisionFRARA:
		return "Auvergne-Rhone-Alpes"
	case SubdivisionFRBFC:
		return "Bourgogne-Franche-Comte"
	case SubdivisionFRBRE:
		return "Bretagne"
	case SubdivisionFRCOR:
		return "Corse"
	case SubdivisionFRCVL:
		return "Centre-Val de Loire"
	case SubdivisionFRGES:
		return "Grand-Est"
	case SubdivisionFRHDF:
		return "Hauts-de-France"
	case SubdivisionFRIDF:
		return "Ile-de-France"
	case SubdivisionFRNAQ:
		return "Nouvelle-Aquitaine"
	case SubdivisionFRNOR:
		return "Normandie"
	case SubdivisionFROCC:
		return "Occitanie"
	case SubdivisionFRPAC:
		return "Provence-Alpes-Cote-d'Azur"
	case SubdivisionFRPDL:
		return "Pays-de-la-Loire"
	case SubdivisionGA1:
		return "Estuaire"
	case SubdivisionGA2:
		return "Haut-Ogooue"
	case SubdivisionGA4:
		return "Ngounie"
	case SubdivisionGA8:
		return "Ogooue-Maritime"
	case SubdivisionGA9:
		return "Woleu-Ntem"
	case SubdivisionGBENG:
		return "England"
	case SubdivisionGBNIR:
		return "Northern Ireland"
	case SubdivisionGBSCT:
		return "Scotland"
	case SubdivisionGBWLS:
		return "Wales"
	case SubdivisionGD01:
		return "Saint Andrew"
	case SubdivisionGD02:
		return "Saint David"
	case SubdivisionGD03:
		return "Saint George"
	case SubdivisionGD04:
		return "Saint John"
	case SubdivisionGD05:
		return "Saint Mark"
	case SubdivisionGD10:
		return "Southern Grenadine Islands"
	case SubdivisionGEAB:
		return "Abkhazia"
	case SubdivisionGEAJ:
		return "Ajaria"
	case SubdivisionGEGU:
		return "Guria"
	case SubdivisionGEIM:
		return "Imereti"
	case SubdivisionGEKA:
		return "K'akheti"
	case SubdivisionGEKK:
		return "Kvemo Kartli"
	case SubdivisionGEMM:
		return "Mtskheta-Mtianeti"
	case SubdivisionGERL:
		return "Rach'a-Lechkhumi-Kvemo Svaneti"
	case SubdivisionGESJ:
		return "Samtskhe-Javakheti"
	case SubdivisionGESK:
		return "Shida Kartli"
	case SubdivisionGESZ:
		return "Samegrelo-Zemo Svaneti"
	case SubdivisionGETB:
		return "Tbilisi"
	case SubdivisionGHAA:
		return "Greater Accra"
	case SubdivisionGHAF:
		return "Ahafo"
	case SubdivisionGHAH:
		return "Ashanti"
	case SubdivisionGHBE:
		return "Bono East"
	case SubdivisionGHBO:
		return "Bono"
	case SubdivisionGHCP:
		return "Central"
	case SubdivisionGHEP:
		return "Eastern"
	case SubdivisionGHNP:
		return "Northern"
	case SubdivisionGHTV:
		return "Volta"
	case SubdivisionGHUE:
		return "Upper East"
	case SubdivisionGHWP:
		return "Western"
	case SubdivisionGLAV:
		return "Avannaata Kommunia"
	case SubdivisionGLKU:
		return "Kommune Kujalleq"
	case SubdivisionGLQE:
		return "Qeqqata Kommunia"
	case SubdivisionGLQT:
		return "Kommune Qeqertalik"
	case SubdivisionGLSM:
		return "Kommuneqarfik Sermersooq"
	case SubdivisionGMB:
		return "Banjul"
	case SubdivisionGML:
		return "Lower River"
	case SubdivisionGMM:
		return "Central River"
	case SubdivisionGMN:
		return "North Bank"
	case SubdivisionGMU:
		return "Upper River"
	case SubdivisionGMW:
		return "Western"
	case SubdivisionGNB:
		return "Boke"
	case SubdivisionGNBF:
		return "Boffa"
	case SubdivisionGNC:
		return "Conakry"
	case SubdivisionGNCO:
		return "Coyah"
	case SubdivisionGND:
		return "Kindia"
	case SubdivisionGNDB:
		return "Dabola"
	case SubdivisionGNK:
		return "Kankan"
	case SubdivisionGNN:
		return "Nzerekore"
	case SubdivisionGNSI:
		return "Siguiri"
	case SubdivisionGQBN:
		return "Bioko Norte"
	case SubdivisionGQBS:
		return "Bioko Sur"
	case SubdivisionGQLI:
		return "Litoral"
	case SubdivisionGQWN:
		return "Wele-Nzas"
	case SubdivisionGR69:
		return "Agion Oros"
	case SubdivisionGRA:
		return "Anatoliki Makedonia kai Thraki"
	case SubdivisionGRB:
		return "Kentriki Makedonia"
	case SubdivisionGRC:
		return "Dytiki Makedonia"
	case SubdivisionGRD:
		return "Ipeiros"
	case SubdivisionGRE:
		return "Thessalia"
	case SubdivisionGRF:
		return "Ionia Nisia"
	case SubdivisionGRG:
		return "Dytiki Ellada"
	case SubdivisionGRH:
		return "Sterea Ellada"
	case SubdivisionGRI:
		return "Attiki"
	case SubdivisionGRJ:
		return "Peloponnisos"
	case SubdivisionGRK:
		return "Voreio Aigaio"
	case SubdivisionGRL:
		return "Notio Aigaio"
	case SubdivisionGRM:
		return "Kriti"
	case SubdivisionGTAV:
		return "Alta Verapaz"
	case SubdivisionGTBV:
		return "Baja Verapaz"
	case SubdivisionGTCM:
		return "Chimaltenango"
	case SubdivisionGTCQ:
		return "Chiquimula"
	case SubdivisionGTES:
		return "Escuintla"
	case SubdivisionGTGU:
		return "Guatemala"
	case SubdivisionGTHU:
		return "Huehuetenango"
	case SubdivisionGTIZ:
		return "Izabal"
	case SubdivisionGTJA:
		return "Jalapa"
	case SubdivisionGTJU:
		return "Jutiapa"
	case SubdivisionGTPE:
		return "Peten"
	case SubdivisionGTPR:
		return "El Progreso"
	case SubdivisionGTQC:
		return "Quiche"
	case SubdivisionGTQZ:
		return "Quetzaltenango"
	case SubdivisionGTRE:
		return "Retalhuleu"
	case SubdivisionGTSA:
		return "Sacatepequez"
	case SubdivisionGTSM:
		return "San Marcos"
	case SubdivisionGTSO:
		return "Solola"
	case SubdivisionGTSR:
		return "Santa Rosa"
	case SubdivisionGTSU:
		return "Suchitepequez"
	case SubdivisionGTTO:
		return "Totonicapan"
	case SubdivisionGTZA:
		return "Zacapa"
	case SubdivisionGWBS:
		return "Bissau"
	case SubdivisionGWGA:
		return "Gabu"
	case SubdivisionGYBA:
		return "Barima-Waini"
	case SubdivisionGYCU:
		return "Cuyuni-Mazaruni"
	case SubdivisionGYDE:
		return "Demerara-Mahaica"
	case SubdivisionGYEB:
		return "East Berbice-Corentyne"
	case SubdivisionGYES:
		return "Essequibo Islands-West Demerara"
	case SubdivisionGYMA:
		return "Mahaica-Berbice"
	case SubdivisionGYPT:
		return "Potaro-Siparuni"
	case SubdivisionGYUD:
		return "Upper Demerara-Berbice"
	case SubdivisionHNAT:
		return "Atlantida"
	case SubdivisionHNCH:
		return "Choluteca"
	case SubdivisionHNCL:
		return "Colon"
	case SubdivisionHNCM:
		return "Comayagua"
	case SubdivisionHNCP:
		return "Copan"
	case SubdivisionHNCR:
		return "Cortes"
	case SubdivisionHNEP:
		return "El Paraiso"
	case SubdivisionHNFM:
		return "Francisco Morazan"
	case SubdivisionHNIB:
		return "Islas de la Bahia"
	case SubdivisionHNIN:
		return "Intibuca"
	case SubdivisionHNLE:
		return "Lempira"
	case SubdivisionHNLP:
		return "La Paz"
	case SubdivisionHNOC:
		return "Ocotepeque"
	case SubdivisionHNOL:
		return "Olancho"
	case SubdivisionHNSB:
		return "Santa Barbara"
	case SubdivisionHNVA:
		return "Valle"
	case SubdivisionHNYO:
		return "Yoro"
	case SubdivisionHR01:
		return "Zagrebacka zupanija"
	case SubdivisionHR02:
		return "Krapinsko-zagorska zupanija"
	case SubdivisionHR03:
		return "Sisacko-moslavacka zupanija"
	case SubdivisionHR04:
		return "Karlovacka zupanija"
	case SubdivisionHR05:
		return "Varazdinska zupanija"
	case SubdivisionHR06:
		return "Koprivnicko-krizevacka zupanija"
	case SubdivisionHR07:
		return "Bjelovarsko-bilogorska zupanija"
	case SubdivisionHR08:
		return "Primorsko-goranska zupanija"
	case SubdivisionHR09:
		return "Licko-senjska zupanija"
	case SubdivisionHR10:
		return "Viroviticko-podravska zupanija"
	case SubdivisionHR11:
		return "Pozesko-slavonska zupanija"
	case SubdivisionHR12:
		return "Brodsko-posavska zupanija"
	case SubdivisionHR13:
		return "Zadarska zupanija"
	case SubdivisionHR14:
		return "Osjecko-baranjska zupanija"
	case SubdivisionHR15:
		return "Sibensko-kninska zupanija"
	case SubdivisionHR16:
		return "Vukovarsko-srijemska zupanija"
	case SubdivisionHR17:
		return "Splitsko-dalmatinska zupanija"
	case SubdivisionHR18:
		return "Istarska zupanija"
	case SubdivisionHR19:
		return "Dubrovacko-neretvanska zupanija"
	case SubdivisionHR20:
		return "Medimurska zupanija"
	case SubdivisionHR21:
		return "Grad Zagreb"
	case SubdivisionHTAR:
		return "Artibonite"
	case SubdivisionHTCE:
		return "Centre"
	case SubdivisionHTND:
		return "Nord"
	case SubdivisionHTOU:
		return "Ouest"
	case SubdivisionHTSD:
		return "Sud"
	case SubdivisionHTSE:
		return "Sud-Est"
	case SubdivisionHUBA:
		return "Baranya"
	case SubdivisionHUBE:
		return "Bekes"
	case SubdivisionHUBK:
		return "Bacs-Kiskun"
	case SubdivisionHUBU:
		return "Budapest"
	case SubdivisionHUBZ:
		return "Borsod-Abauj-Zemplen"
	case SubdivisionHUCS:
		return "Csongrad"
	case SubdivisionHUFE:
		return "Fejer"
	case SubdivisionHUGS:
		return "Gyor-Moson-Sopron"
	case SubdivisionHUHB:
		return "Hajdu-Bihar"
	case SubdivisionHUHE:
		return "Heves"
	case SubdivisionHUJN:
		return "Jasz-Nagykun-Szolnok"
	case SubdivisionHUKE:
		return "Komarom-Esztergom"
	case SubdivisionHUNO:
		return "Nograd"
	case SubdivisionHUPE:
		return "Pest"
	case SubdivisionHUSO:
		return "Somogy"
	case SubdivisionHUSZ:
		return "Szabolcs-Szatmar-Bereg"
	case SubdivisionHUTO:
		return "Tolna"
	case SubdivisionHUVA:
		return "Vas"
	case SubdivisionHUVE:
		return "Veszprem"
	case SubdivisionHUZA:
		return "Zala"
	case SubdivisionIDAC:
		return "Aceh"
	case SubdivisionIDBA:
		return "Bali"
	case SubdivisionIDBB:
		return "Kepulauan Bangka Belitung"
	case SubdivisionIDBE:
		return "Bengkulu"
	case SubdivisionIDBT:
		return "Banten"
	case SubdivisionIDGO:
		return "Gorontalo"
	case SubdivisionIDJA:
		return "Jambi"
	case SubdivisionIDJB:
		return "Jawa Barat"
	case SubdivisionIDJI:
		return "Jawa Timur"
	case SubdivisionIDJK:
		return "Jakarta Raya"
	case SubdivisionIDJT:
		return "Jawa Tengah"
	case SubdivisionIDKB:
		return "Kalimantan Barat"
	case SubdivisionIDKI:
		return "Kalimantan Timur"
	case SubdivisionIDKR:
		return "Kepulauan Riau"
	case SubdivisionIDKS:
		return "Kalimantan Selatan"
	case SubdivisionIDKT:
		return "Kalimantan Tengah"
	case SubdivisionIDKU:
		return "Kalimantan Utara"
	case SubdivisionIDLA:
		return "Lampung"
	case SubdivisionIDML:
		return "Maluku"
	case SubdivisionIDMU:
		return "Maluku Utara"
	case SubdivisionIDNB:
		return "Nusa Tenggara Barat"
	case SubdivisionIDNT:
		return "Nusa Tenggara Timur"
	case SubdivisionIDPB:
		return "Papua Barat"
	case SubdivisionIDPP:
		return "Papua"
	case SubdivisionIDRI:
		return "Riau"
	case SubdivisionIDSA:
		return "Sulawesi Utara"
	case SubdivisionIDSB:
		return "Sumatera Barat"
	case SubdivisionIDSG:
		return "Sulawesi Tenggara"
	case SubdivisionIDSN:
		return "Sulawesi Selatan"
	case SubdivisionIDSR:
		return "Sulawesi Barat"
	case SubdivisionIDSS:
		return "Sumatera Selatan"
	case SubdivisionIDST:
		return "Sulawesi Tengah"
	case SubdivisionIDSU:
		return "Sumatera Utara"
	case SubdivisionIDYO:
		return "Yogyakarta"
	case SubdivisionIECE:
		return "Clare"
	case SubdivisionIECN:
		return "Cavan"
	case SubdivisionIECO:
		return "Cork"
	case SubdivisionIECW:
		return "Carlow"
	case SubdivisionIED:
		return "Dublin"
	case SubdivisionIEDL:
		return "Donegal"
	case SubdivisionIEG:
		return "Galway"
	case SubdivisionIEKE:
		return "Kildare"
	case SubdivisionIEKK:
		return "Kilkenny"
	case SubdivisionIEKY:
		return "Kerry"
	case SubdivisionIELD:
		return "Longford"
	case SubdivisionIELH:
		return "Louth"
	case SubdivisionIELK:
		return "Limerick"
	case SubdivisionIELM:
		return "Leitrim"
	case SubdivisionIELS:
		return "Laois"
	case SubdivisionIEMH:
		return "Meath"
	case SubdivisionIEMN:
		return "Monaghan"
	case SubdivisionIEMO:
		return "Mayo"
	case SubdivisionIEOY:
		return "Offaly"
	case SubdivisionIERN:
		return "Roscommon"
	case SubdivisionIESO:
		return "Sligo"
	case SubdivisionIETA:
		return "Tipperary"
	case SubdivisionIEWD:
		return "Waterford"
	case SubdivisionIEWH:
		return "Westmeath"
	case SubdivisionIEWW:
		return "Wicklow"
	case SubdivisionIEWX:
		return "Wexford"
	case SubdivisionILD:
		return "HaDarom"
	case SubdivisionILHA:
		return "Hefa"
	case SubdivisionILJM:
		return "Yerushalayim"
	case SubdivisionILM:
		return "HaMerkaz"
	case SubdivisionILTA:
		return "Tel Aviv"
	case SubdivisionILZ:
		return "HaTsafon"
	case SubdivisionINAN:
		return "Andaman and Nicobar Islands"
	case SubdivisionINAP:
		return "Andhra Pradesh"
	case SubdivisionINAR:
		return "Arunachal Pradesh"
	case SubdivisionINAS:
		return "Assam"
	case SubdivisionINBR:
		return "Bihar"
	case SubdivisionINCH:
		return "Chandigarh"
	case SubdivisionINCT:
		return "Chhattisgarh"
	case SubdivisionINDH:
		return "Dadra and Nagar Haveli and Daman and Diu"
	case SubdivisionINDL:
		return "Delhi"
	case SubdivisionINDN:
		return "Dadra and Nagar Haveli"
	case SubdivisionINGA:
		return "Goa"
	case SubdivisionINGJ:
		return "Gujarat"
	case SubdivisionINHP:
		return "Himachal Pradesh"
	case SubdivisionINHR:
		return "Haryana"
	case SubdivisionINJH:
		return "Jharkhand"
	case SubdivisionINJK:
		return "Jammu and Kashmir"
	case SubdivisionINKA:
		return "Karnataka"
	case SubdivisionINKL:
		return "Kerala"
	case SubdivisionINLD:
		return "Lakshadweep"
	case SubdivisionINMH:
		return "Maharashtra"
	case SubdivisionINML:
		return "Meghalaya"
	case SubdivisionINMN:
		return "Manipur"
	case SubdivisionINMP:
		return "Madhya Pradesh"
	case SubdivisionINMZ:
		return "Mizoram"
	case SubdivisionINNL:
		return "Nagaland"
	case SubdivisionINOR:
		return "Odisha"
	case SubdivisionINPB:
		return "Punjab"
	case SubdivisionINPY:
		return "Puducherry"
	case SubdivisionINRJ:
		return "Rajasthan"
	case SubdivisionINSK:
		return "Sikkim"
	case SubdivisionINTG:
		return "Telangana"
	case SubdivisionINTN:
		return "Tamil Nadu"
	case SubdivisionINTR:
		return "Tripura"
	case SubdivisionINUP:
		return "Uttar Pradesh"
	case SubdivisionINUT:
		return "Uttarakhand"
	case SubdivisionINWB:
		return "West Bengal"
	case SubdivisionIQAN:
		return "Al Anbar"
	case SubdivisionIQAR:
		return "Arbil"
	case SubdivisionIQBA:
		return "Al Basrah"
	case SubdivisionIQBB:
		return "Babil"
	case SubdivisionIQBG:
		return "Baghdad"
	case SubdivisionIQDA:
		return "Dahuk"
	case SubdivisionIQDI:
		return "Diyala"
	case SubdivisionIQDQ:
		return "Dhi Qar"
	case SubdivisionIQKA:
		return "Karbala'"
	case SubdivisionIQKI:
		return "Kirkuk"
	case SubdivisionIQMA:
		return "Maysan"
	case SubdivisionIQMU:
		return "Al Muthanna"
	case SubdivisionIQNA:
		return "An Najaf"
	case SubdivisionIQNI:
		return "Ninawa"
	case SubdivisionIQQA:
		return "Al Qadisiyah"
	case SubdivisionIQSD:
		return "Salah ad Din"
	case SubdivisionIQSU:
		return "As Sulaymaniyah"
	case SubdivisionIQWA:
		return "Wasit"
	case SubdivisionIR01:
		return "Azarbayjan-e Sharqi"
	case SubdivisionIR02:
		return "Azarbayjan-e Gharbi"
	case SubdivisionIR03:
		return "Ardabil"
	case SubdivisionIR04:
		return "Esfahan"
	case SubdivisionIR05:
		return "Ilam"
	case SubdivisionIR06:
		return "Bushehr"
	case SubdivisionIR07:
		return "Tehran"
	case SubdivisionIR08:
		return "Chahar Mahal va Bakhtiari"
	case SubdivisionIR10:
		return "Khuzestan"
	case SubdivisionIR11:
		return "Zanjan"
	case SubdivisionIR12:
		return "Semnan"
	case SubdivisionIR13:
		return "Sistan va Baluchestan"
	case SubdivisionIR14:
		return "Fars"
	case SubdivisionIR15:
		return "Kerman"
	case SubdivisionIR16:
		return "Kordestan"
	case SubdivisionIR17:
		return "Kermanshah"
	case SubdivisionIR18:
		return "Kohgiluyeh va Bowyer Ahmad"
	case SubdivisionIR19:
		return "Gilan"
	case SubdivisionIR20:
		return "Lorestan"
	case SubdivisionIR21:
		return "Mazandaran"
	case SubdivisionIR22:
		return "Markazi"
	case SubdivisionIR23:
		return "Hormozgan"
	case SubdivisionIR24:
		return "Hamadan"
	case SubdivisionIR25:
		return "Yazd"
	case SubdivisionIR26:
		return "Qom"
	case SubdivisionIR27:
		return "Golestan"
	case SubdivisionIR28:
		return "Qazvin"
	case SubdivisionIR29:
		return "Khorasan-e Jonubi"
	case SubdivisionIR30:
		return "Khorasan-e Razavi"
	case SubdivisionIR31:
		return "Khorasan-e Shomali"
	case SubdivisionIR32:
		return "Alborz"
	case SubdivisionIS1:
		return "Hofudborgarsvaedi"
	case SubdivisionIS2:
		return "Sudurnes"
	case SubdivisionIS3:
		return "Vesturland"
	case SubdivisionIS4:
		return "Vestfirdir"
	case SubdivisionIS5:
		return "Nordurland vestra"
	case SubdivisionIS6:
		return "Nordurland eystra"
	case SubdivisionIS7:
		return "Austurland"
	case SubdivisionIS8:
		return "Sudurland"
	case SubdivisionIT21:
		return "Piemonte"
	case SubdivisionIT23:
		return "Valle d'Aosta"
	case SubdivisionIT25:
		return "Lombardia"
	case SubdivisionIT32:
		return "Trentino-Alto Adige"
	case SubdivisionIT34:
		return "Veneto"
	case SubdivisionIT36:
		return "Friuli-Venezia Giulia"
	case SubdivisionIT42:
		return "Liguria"
	case SubdivisionIT45:
		return "Emilia-Romagna"
	case SubdivisionIT52:
		return "Toscana"
	case SubdivisionIT55:
		return "Umbria"
	case SubdivisionIT57:
		return "Marche"
	case SubdivisionIT62:
		return "Lazio"
	case SubdivisionIT65:
		return "Abruzzo"
	case SubdivisionIT67:
		return "Molise"
	case SubdivisionIT72:
		return "Campania"
	case SubdivisionIT75:
		return "Puglia"
	case SubdivisionIT77:
		return "Basilicata"
	case SubdivisionIT78:
		return "Calabria"
	case SubdivisionIT82:
		return "Sicilia"
	case SubdivisionIT88:
		return "Sardegna"
	case SubdivisionJM01:
		return "Kingston"
	case SubdivisionJM02:
		return "Saint Andrew"
	case SubdivisionJM03:
		return "Saint Thomas"
	case SubdivisionJM04:
		return "Portland"
	case SubdivisionJM05:
		return "Saint Mary"
	case SubdivisionJM06:
		return "Saint Ann"
	case SubdivisionJM07:
		return "Trelawny"
	case SubdivisionJM08:
		return "Saint James"
	case SubdivisionJM09:
		return "Hanover"
	case SubdivisionJM10:
		return "Westmoreland"
	case SubdivisionJM11:
		return "Saint Elizabeth"
	case SubdivisionJM12:
		return "Manchester"
	case SubdivisionJM13:
		return "Clarendon"
	case SubdivisionJM14:
		return "Saint Catherine"
	case SubdivisionJOAJ:
		return "'Ajlun"
	case SubdivisionJOAM:
		return "Al 'Asimah"
	case SubdivisionJOAQ:
		return "Al 'Aqabah"
	case SubdivisionJOAZ:
		return "Az Zarqa'"
	case SubdivisionJOBA:
		return "Al Balqa'"
	case SubdivisionJOIR:
		return "Irbid"
	case SubdivisionJOJA:
		return "Jarash"
	case SubdivisionJOKA:
		return "Al Karak"
	case SubdivisionJOMA:
		return "Al Mafraq"
	case SubdivisionJOMD:
		return "Madaba"
	case SubdivisionJOMN:
		return "Ma'an"
	case SubdivisionJP01:
		return "Hokkaido"
	case SubdivisionJP02:
		return "Aomori"
	case SubdivisionJP03:
		return "Iwate"
	case SubdivisionJP04:
		return "Miyagi"
	case SubdivisionJP05:
		return "Akita"
	case SubdivisionJP06:
		return "Yamagata"
	case SubdivisionJP07:
		return "Fukushima"
	case SubdivisionJP08:
		return "Ibaraki"
	case SubdivisionJP09:
		return "Tochigi"
	case SubdivisionJP10:
		return "Gunma"
	case SubdivisionJP11:
		return "Saitama"
	case SubdivisionJP12:
		return "Chiba"
	case SubdivisionJP13:
		return "Tokyo"
	case SubdivisionJP14:
		return "Kanagawa"
	case SubdivisionJP15:
		return "Niigata"
	case SubdivisionJP16:
		return "Toyama"
	case SubdivisionJP17:
		return "Ishikawa"
	case SubdivisionJP18:
		return "Fukui"
	case SubdivisionJP19:
		return "Yamanashi"
	case SubdivisionJP20:
		return "Nagano"
	case SubdivisionJP21:
		return "Gifu"
	case SubdivisionJP22:
		return "Shizuoka"
	case SubdivisionJP23:
		return "Aichi"
	case SubdivisionJP24:
		return "Mie"
	case SubdivisionJP25:
		return "Shiga"
	case SubdivisionJP26:
		return "Kyoto"
	case SubdivisionJP27:
		return "Osaka"
	case SubdivisionJP28:
		return "Hyogo"
	case SubdivisionJP29:
		return "Nara"
	case SubdivisionJP30:
		return "Wakayama"
	case SubdivisionJP31:
		return "Tottori"
	case SubdivisionJP32:
		return "Shimane"
	case SubdivisionJP33:
		return "Okayama"
	case SubdivisionJP34:
		return "Hiroshima"
	case SubdivisionJP35:
		return "Yamaguchi"
	case SubdivisionJP36:
		return "Tokushima"
	case SubdivisionJP37:
		return "Kagawa"
	case SubdivisionJP38:
		return "Ehime"
	case SubdivisionJP39:
		return "Kochi"
	case SubdivisionJP40:
		return "Fukuoka"
	case SubdivisionJP41:
		return "Saga"
	case SubdivisionJP42:
		return "Nagasaki"
	case SubdivisionJP43:
		return "Kumamoto"
	case SubdivisionJP44:
		return "Oita"
	case SubdivisionJP45:
		return "Miyazaki"
	case SubdivisionJP46:
		return "Kagoshima"
	case SubdivisionJP47:
		return "Okinawa"
	case SubdivisionKE01:
		return "Baringo"
	case SubdivisionKE02:
		return "Bomet"
	case SubdivisionKE03:
		return "Bungoma"
	case SubdivisionKE04:
		return "Busia"
	case SubdivisionKE05:
		return "Elgeyo/Marakwet"
	case SubdivisionKE06:
		return "Embu"
	case SubdivisionKE07:
		return "Garissa"
	case SubdivisionKE08:
		return "Homa Bay"
	case SubdivisionKE09:
		return "Isiolo"
	case SubdivisionKE10:
		return "Kajiado"
	case SubdivisionKE11:
		return "Kakamega"
	case SubdivisionKE12:
		return "Kericho"
	case SubdivisionKE13:
		return "Kiambu"
	case SubdivisionKE14:
		return "Kilifi"
	case SubdivisionKE15:
		return "Kirinyaga"
	case SubdivisionKE16:
		return "Kisii"
	case SubdivisionKE17:
		return "Kisumu"
	case SubdivisionKE18:
		return "Kitui"
	case SubdivisionKE19:
		return "Kwale"
	case SubdivisionKE20:
		return "Laikipia"
	case SubdivisionKE21:
		return "Lamu"
	case SubdivisionKE22:
		return "Machakos"
	case SubdivisionKE23:
		return "Makueni"
	case SubdivisionKE24:
		return "Mandera"
	case SubdivisionKE25:
		return "Marsabit"
	case SubdivisionKE26:
		return "Meru"
	case SubdivisionKE27:
		return "Migori"
	case SubdivisionKE28:
		return "Mombasa"
	case SubdivisionKE29:
		return "Murang'a"
	case SubdivisionKE30:
		return "Nairobi City"
	case SubdivisionKE31:
		return "Nakuru"
	case SubdivisionKE32:
		return "Nandi"
	case SubdivisionKE33:
		return "Narok"
	case SubdivisionKE34:
		return "Nyamira"
	case SubdivisionKE35:
		return "Nyandarua"
	case SubdivisionKE36:
		return "Nyeri"
	case SubdivisionKE38:
		return "Siaya"
	case SubdivisionKE39:
		return "Taita/Taveta"
	case SubdivisionKE41:
		return "Tharaka-Nithi"
	case SubdivisionKE42:
		return "Trans Nzoia"
	case SubdivisionKE43:
		return "Turkana"
	case SubdivisionKE44:
		return "Uasin Gishu"
	case SubdivisionKE46:
		return "Wajir"
	case SubdivisionKGB:
		return "Batken"
	case SubdivisionKGC:
		return "Chuy"
	case SubdivisionKGGB:
		return "Bishkek Shaary"
	case SubdivisionKGGO:
		return "Osh Shaary"
	case SubdivisionKGJ:
		return "Jalal-Abad"
	case SubdivisionKGN:
		return "Naryn"
	case SubdivisionKGT:
		return "Talas"
	case SubdivisionKGY:
		return "Ysyk-Kol"
	case SubdivisionKH10:
		return "Kracheh"
	case SubdivisionKH11:
		return "Mondol Kiri"
	case SubdivisionKH12:
		return "Phnom Penh"
	case SubdivisionKH14:
		return "Prey Veaeng"
	case SubdivisionKH15:
		return "Pousaat"
	case SubdivisionKH17:
		return "Siem Reab"
	case SubdivisionKH18:
		return "Preah Sihanouk"
	case SubdivisionKH19:
		return "Stueng Traeng"
	case SubdivisionKH1:
		return "Banteay Mean Choay"
	case SubdivisionKH20:
		return "Svaay Rieng"
	case SubdivisionKH21:
		return "Taakaev"
	case SubdivisionKH23:
		return "Krong Kaeb"
	case SubdivisionKH24:
		return "Krong Pailin"
	case SubdivisionKH2:
		return "Baat Dambang"
	case SubdivisionKH3:
		return "Kampong Chaam"
	case SubdivisionKH4:
		return "Kampong Chhnang"
	case SubdivisionKH5:
		return "Kampong Spueu"
	case SubdivisionKH6:
		return "Kampong Thum"
	case SubdivisionKH7:
		return "Kampot"
	case SubdivisionKH8:
		return "Kandaal"
	case SubdivisionKH9:
		return "Kaoh Kong"
	case SubdivisionKIG:
		return "Gilbert Islands"
	case SubdivisionKMG:
		return "Grande Comore"
	case SubdivisionKN02:
		return "Saint Anne Sandy Point"
	case SubdivisionKN03:
		return "Saint George Basseterre"
	case SubdivisionKN05:
		return "Saint James Windward"
	case SubdivisionKN06:
		return "Saint John Capisterre"
	case SubdivisionKN07:
		return "Saint John Figtree"
	case SubdivisionKN08:
		return "Saint Mary Cayon"
	case SubdivisionKN09:
		return "Saint Paul Capisterre"
	case SubdivisionKN10:
		return "Saint Paul Charlestown"
	case SubdivisionKN11:
		return "Saint Peter Basseterre"
	case SubdivisionKN12:
		return "Saint Thomas Lowland"
	case SubdivisionKN13:
		return "Saint Thomas Middle Island"
	case SubdivisionKP01:
		return "P'yongyang"
	case SubdivisionKR11:
		return "Seoul-teukbyeolsi"
	case SubdivisionKR26:
		return "Busan-gwangyeoksi"
	case SubdivisionKR27:
		return "Daegu-gwangyeoksi"
	case SubdivisionKR28:
		return "Incheon-gwangyeoksi"
	case SubdivisionKR29:
		return "Gwangju-gwangyeoksi"
	case SubdivisionKR30:
		return "Daejeon-gwangyeoksi"
	case SubdivisionKR31:
		return "Ulsan-gwangyeoksi"
	case SubdivisionKR41:
		return "Gyeonggi-do"
	case SubdivisionKR42:
		return "Gangwon-do"
	case SubdivisionKR43:
		return "Chungcheongbuk-do"
	case SubdivisionKR44:
		return "Chungcheongnam-do"
	case SubdivisionKR45:
		return "Jeollabuk-do"
	case SubdivisionKR46:
		return "Jeollanam-do"
	case SubdivisionKR47:
		return "Gyeongsangbuk-do"
	case SubdivisionKR48:
		return "Gyeongsangnam-do"
	case SubdivisionKR49:
		return "Jeju-teukbyeoljachido"
	case SubdivisionKWAH:
		return "Al Ahmadi"
	case SubdivisionKWFA:
		return "Al Farwaniyah"
	case SubdivisionKWHA:
		return "Hawalli"
	case SubdivisionKWJA:
		return "Al Jahra'"
	case SubdivisionKWKU:
		return "Al 'Asimah"
	case SubdivisionKWMU:
		return "Mubarak al Kabir"
	case SubdivisionKZAKM:
		return "Aqmola oblysy"
	case SubdivisionKZAKT:
		return "Aqtobe oblysy"
	case SubdivisionKZALA:
		return "Almaty"
	case SubdivisionKZALM:
		return "Almaty oblysy"
	case SubdivisionKZAST:
		return "Nur-Sultan"
	case SubdivisionKZATY:
		return "Atyrau oblysy"
	case SubdivisionKZKAR:
		return "Qaraghandy oblysy"
	case SubdivisionKZKUS:
		return "Qostanay oblysy"
	case SubdivisionKZKZY:
		return "Qyzylorda oblysy"
	case SubdivisionKZMAN:
		return "Mangghystau oblysy"
	case SubdivisionKZPAV:
		return "Pavlodar oblysy"
	case SubdivisionKZSEV:
		return "Soltustik Qazaqstan oblysy"
	case SubdivisionKZSHY:
		return "Shymkent"
	case SubdivisionKZVOS:
		return "Shyghys Qazaqstan oblysy"
	case SubdivisionKZYUZ:
		return "Ongtustik Qazaqstan oblysy"
	case SubdivisionKZZAP:
		return "Batys Qazaqstan oblysy"
	case SubdivisionKZZHA:
		return "Zhambyl oblysy"
	case SubdivisionLABL:
		return "Bolikhamxai"
	case SubdivisionLACH:
		return "Champasak"
	case SubdivisionLAKH:
		return "Khammouan"
	case SubdivisionLALP:
		return "Louangphabang"
	case SubdivisionLAOU:
		return "Oudomxai"
	case SubdivisionLAPH:
		return "Phongsali"
	case SubdivisionLASV:
		return "Savannakhet"
	case SubdivisionLAVI:
		return "Viangchan"
	case SubdivisionLAXA:
		return "Xaignabouli"
	case SubdivisionLAXE:
		return "Xekong"
	case SubdivisionLAXI:
		return "Xiangkhouang"
	case SubdivisionLBAK:
		return "Aakkar"
	case SubdivisionLBAS:
		return "Liban-Nord"
	case SubdivisionLBBA:
		return "Beyrouth"
	case SubdivisionLBBH:
		return "Baalbek-Hermel"
	case SubdivisionLBBI:
		return "Beqaa"
	case SubdivisionLBJA:
		return "Liban-Sud"
	case SubdivisionLBJL:
		return "Mont-Liban"
	case SubdivisionLBNA:
		return "Nabatiye"
	case SubdivisionLC01:
		return "Anse la Raye"
	case SubdivisionLC02:
		return "Castries"
	case SubdivisionLC05:
		return "Dennery"
	case SubdivisionLC06:
		return "Gros Islet"
	case SubdivisionLC07:
		return "Laborie"
	case SubdivisionLC10:
		return "Soufriere"
	case SubdivisionLC11:
		return "Vieux Fort"
	case SubdivisionLI01:
		return "Balzers"
	case SubdivisionLI02:
		return "Eschen"
	case SubdivisionLI03:
		return "Gamprin"
	case SubdivisionLI04:
		return "Mauren"
	case SubdivisionLI05:
		return "Planken"
	case SubdivisionLI06:
		return "Ruggell"
	case SubdivisionLI07:
		return "Schaan"
	case SubdivisionLI09:
		return "Triesen"
	case SubdivisionLI10:
		return "Triesenberg"
	case SubdivisionLI11:
		return "Vaduz"
	case SubdivisionLK1:
		return "Western Province"
	case SubdivisionLK2:
		return "Central Province"
	case SubdivisionLK3:
		return "Southern Province"
	case SubdivisionLK4:
		return "Northern Province"
	case SubdivisionLK5:
		return "Eastern Province"
	case SubdivisionLK6:
		return "North Western Province"
	case SubdivisionLK7:
		return "North Central Province"
	case SubdivisionLK8:
		return "Uva Province"
	case SubdivisionLK9:
		return "Sabaragamuwa Province"
	case SubdivisionLRBM:
		return "Bomi"
	case SubdivisionLRGB:
		return "Grand Bassa"
	case SubdivisionLRGG:
		return "Grand Gedeh"
	case SubdivisionLRMG:
		return "Margibi"
	case SubdivisionLRMO:
		return "Montserrado"
	case SubdivisionLRNI:
		return "Nimba"
	case SubdivisionLRSI:
		return "Sinoe"
	case SubdivisionLSA:
		return "Maseru"
	case SubdivisionLSC:
		return "Leribe"
	case SubdivisionLSE:
		return "Mafeteng"
	case SubdivisionLSG:
		return "Quthing"
	case SubdivisionLSJ:
		return "Mokhotlong"
	case SubdivisionLSK:
		return "Thaba-Tseka"
	case SubdivisionLTAL:
		return "Alytaus apskritis"
	case SubdivisionLTKL:
		return "Klaipedos apskritis"
	case SubdivisionLTKU:
		return "Kauno apskritis"
	case SubdivisionLTMR:
		return "Marijampoles apskritis"
	case SubdivisionLTPN:
		return "Panevezio apskritis"
	case SubdivisionLTSA:
		return "Siauliu apskritis"
	case SubdivisionLTTA:
		return "Taurages apskritis"
	case SubdivisionLTTE:
		return "Telsiu apskritis"
	case SubdivisionLTUT:
		return "Utenos apskritis"
	case SubdivisionLTVL:
		return "Vilniaus apskritis"
	case SubdivisionLUCA:
		return "Capellen"
	case SubdivisionLUCL:
		return "Clervaux"
	case SubdivisionLUDI:
		return "Diekirch"
	case SubdivisionLUEC:
		return "Echternach"
	case SubdivisionLUES:
		return "Esch-sur-Alzette"
	case SubdivisionLUGR:
		return "Grevenmacher"
	case SubdivisionLULU:
		return "Luxembourg"
	case SubdivisionLUME:
		return "Mersch"
	case SubdivisionLURD:
		return "Redange"
	case SubdivisionLURM:
		return "Remich"
	case SubdivisionLUVD:
		return "Vianden"
	case SubdivisionLUWI:
		return "Wiltz"
	case SubdivisionLV001:
		return "Aglonas novads"
	case SubdivisionLV002:
		return "Aizkraukles novads"
	case SubdivisionLV003:
		return "Aizputes novads"
	case SubdivisionLV005:
		return "Alojas novads"
	case SubdivisionLV007:
		return "Aluksnes novads"
	case SubdivisionLV010:
		return "Auces novads"
	case SubdivisionLV011:
		return "Adazu novads"
	case SubdivisionLV012:
		return "Babites novads"
	case SubdivisionLV015:
		return "Balvu novads"
	case SubdivisionLV016:
		return "Bauskas novads"
	case SubdivisionLV017:
		return "Beverinas novads"
	case SubdivisionLV018:
		return "Brocenu novads"
	case SubdivisionLV020:
		return "Carnikavas novads"
	case SubdivisionLV021:
		return "Cesvaines novads"
	case SubdivisionLV022:
		return "Cesu novads"
	case SubdivisionLV024:
		return "Dagdas novads"
	case SubdivisionLV025:
		return "Daugavpils novads"
	case SubdivisionLV026:
		return "Dobeles novads"
	case SubdivisionLV027:
		return "Dundagas novads"
	case SubdivisionLV029:
		return "Engures novads"
	case SubdivisionLV030:
		return "Erglu novads"
	case SubdivisionLV033:
		return "Gulbenes novads"
	case SubdivisionLV034:
		return "Iecavas novads"
	case SubdivisionLV035:
		return "Ikskiles novads"
	case SubdivisionLV037:
		return "Incukalna novads"
	case SubdivisionLV038:
		return "Jaunjelgavas novads"
	case SubdivisionLV039:
		return "Jaunpiebalgas novads"
	case SubdivisionLV040:
		return "Jaunpils novads"
	case SubdivisionLV041:
		return "Jelgavas novads"
	case SubdivisionLV042:
		return "Jekabpils novads"
	case SubdivisionLV043:
		return "Kandavas novads"
	case SubdivisionLV046:
		return "Kokneses novads"
	case SubdivisionLV047:
		return "Kraslavas novads"
	case SubdivisionLV050:
		return "Kuldigas novads"
	case SubdivisionLV052:
		return "Kekavas novads"
	case SubdivisionLV053:
		return "Lielvardes novads"
	case SubdivisionLV054:
		return "Limbazu novads"
	case SubdivisionLV056:
		return "Livanu novads"
	case SubdivisionLV057:
		return "Lubanas novads"
	case SubdivisionLV058:
		return "Ludzas novads"
	case SubdivisionLV059:
		return "Madonas novads"
	case SubdivisionLV061:
		return "Malpils novads"
	case SubdivisionLV064:
		return "Nauksenu novads"
	case SubdivisionLV067:
		return "Ogres novads"
	case SubdivisionLV068:
		return "Olaines novads"
	case SubdivisionLV069:
		return "Ozolnieku novads"
	case SubdivisionLV073:
		return "Preilu novads"
	case SubdivisionLV075:
		return "Priekules novads"
	case SubdivisionLV077:
		return "Rezeknes novads"
	case SubdivisionLV078:
		return "Riebinu novads"
	case SubdivisionLV079:
		return "Rojas novads"
	case SubdivisionLV083:
		return "Rundales novads"
	case SubdivisionLV086:
		return "Salacgrivas novads"
	case SubdivisionLV087:
		return "Salaspils novads"
	case SubdivisionLV088:
		return "Saldus novads"
	case SubdivisionLV089:
		return "Saulkrastu novads"
	case SubdivisionLV090:
		return "Sejas novads"
	case SubdivisionLV091:
		return "Siguldas novads"
	case SubdivisionLV093:
		return "Skrundas novads"
	case SubdivisionLV094:
		return "Smiltenes novads"
	case SubdivisionLV095:
		return "Stopinu novads"
	case SubdivisionLV097:
		return "Talsu novads"
	case SubdivisionLV099:
		return "Tukuma novads"
	case SubdivisionLV100:
		return "Vainodes novads"
	case SubdivisionLV101:
		return "Valkas novads"
	case SubdivisionLV103:
		return "Varkavas novads"
	case SubdivisionLV105:
		return "Vecumnieku novads"
	case SubdivisionLV106:
		return "Ventspils novads"
	case SubdivisionLV110:
		return "Zilupes novads"
	case SubdivisionLVJEL:
		return "Jelgava"
	case SubdivisionLVJUR:
		return "Jurmala"
	case SubdivisionLVLPX:
		return "Liepaja"
	case SubdivisionLVRIX:
		return "Riga"
	case SubdivisionLVVMR:
		return "Valmiera"
	case SubdivisionLYBA:
		return "Banghazi"
	case SubdivisionLYBU:
		return "Al Butnan"
	case SubdivisionLYDR:
		return "Darnah"
	case SubdivisionLYJA:
		return "Al Jabal al Akhdar"
	case SubdivisionLYJG:
		return "Al Jabal al Gharbi"
	case SubdivisionLYJI:
		return "Al Jafarah"
	case SubdivisionLYJU:
		return "Al Jufrah"
	case SubdivisionLYMB:
		return "Al Marqab"
	case SubdivisionLYMI:
		return "Misratah"
	case SubdivisionLYMJ:
		return "Al Marj"
	case SubdivisionLYMQ:
		return "Murzuq"
	case SubdivisionLYNL:
		return "Nalut"
	case SubdivisionLYNQ:
		return "An Nuqat al Khams"
	case SubdivisionLYSB:
		return "Sabha"
	case SubdivisionLYSR:
		return "Surt"
	case SubdivisionLYTB:
		return "Tarabulus"
	case SubdivisionLYWA:
		return "Al Wahat"
	case SubdivisionLYZA:
		return "Az Zawiyah"
	case SubdivisionMA01:
		return "Tanger-Tetouan-Al Hoceima"
	case SubdivisionMA02:
		return "L'Oriental"
	case SubdivisionMA03:
		return "Fes- Meknes"
	case SubdivisionMA04:
		return "Rabat-Sale-Kenitra"
	case SubdivisionMA05:
		return "Beni-Mellal-Khenifra"
	case SubdivisionMA06:
		return "Casablanca-Settat"
	case SubdivisionMA07:
		return "Marrakech-Safi"
	case SubdivisionMA08:
		return "Draa-Tafilalet"
	case SubdivisionMA09:
		return "Souss-Massa"
	case SubdivisionMA10:
		return "Guelmim-Oued Noun (EH-partial)"
	case SubdivisionMA11:
		return "Laayoune-Sakia El Hamra (EH-partial)"
	case SubdivisionMCCO:
		return "La Condamine"
	case SubdivisionMCFO:
		return "Fontvieille"
	case SubdivisionMCMC:
		return "Monte-Carlo"
	case SubdivisionMCMO:
		return "Monaco-Ville"
	case SubdivisionMCSR:
		return "Saint-Roman"
	case SubdivisionMDAN:
		return "Anenii Noi"
	case SubdivisionMDBA:
		return "Balti"
	case SubdivisionMDBD:
		return "Bender"
	case SubdivisionMDBR:
		return "Briceni"
	case SubdivisionMDBS:
		return "Basarabeasca"
	case SubdivisionMDCA:
		return "Cahul"
	case SubdivisionMDCL:
		return "Calarasi"
	case SubdivisionMDCM:
		return "Cimislia"
	case SubdivisionMDCR:
		return "Criuleni"
	case SubdivisionMDCS:
		return "Causeni"
	case SubdivisionMDCT:
		return "Cantemir"
	case SubdivisionMDCU:
		return "Chisinau"
	case SubdivisionMDDO:
		return "Donduseni"
	case SubdivisionMDDR:
		return "Drochia"
	case SubdivisionMDDU:
		return "Dubasari"
	case SubdivisionMDED:
		return "Edinet"
	case SubdivisionMDFA:
		return "Falesti"
	case SubdivisionMDFL:
		return "Floresti"
	case SubdivisionMDGA:
		return "Gagauzia, Unitatea teritoriala autonoma"
	case SubdivisionMDGL:
		return "Glodeni"
	case SubdivisionMDHI:
		return "Hincesti"
	case SubdivisionMDIA:
		return "Ialoveni"
	case SubdivisionMDLE:
		return "Leova"
	case SubdivisionMDNI:
		return "Nisporeni"
	case SubdivisionMDOC:
		return "Ocnita"
	case SubdivisionMDOR:
		return "Orhei"
	case SubdivisionMDRE:
		return "Rezina"
	case SubdivisionMDRI:
		return "Riscani"
	case SubdivisionMDSD:
		return "Soldanesti"
	case SubdivisionMDSI:
		return "Singerei"
	case SubdivisionMDSN:
		return "Stinga Nistrului, unitatea teritoriala din"
	case SubdivisionMDSO:
		return "Soroca"
	case SubdivisionMDST:
		return "Straseni"
	case SubdivisionMDSV:
		return "Stefan Voda"
	case SubdivisionMDTA:
		return "Taraclia"
	case SubdivisionMDTE:
		return "Telenesti"
	case SubdivisionMDUN:
		return "Ungheni"
	case SubdivisionME02:
		return "Bar"
	case SubdivisionME03:
		return "Berane"
	case SubdivisionME04:
		return "Bijelo Polje"
	case SubdivisionME05:
		return "Budva"
	case SubdivisionME06:
		return "Cetinje"
	case SubdivisionME07:
		return "Danilovgrad"
	case SubdivisionME08:
		return "Herceg-Novi"
	case SubdivisionME09:
		return "Kolasin"
	case SubdivisionME10:
		return "Kotor"
	case SubdivisionME12:
		return "Niksic"
	case SubdivisionME13:
		return "Plav"
	case SubdivisionME14:
		return "Pljevlja"
	case SubdivisionME15:
		return "Pluzine"
	case SubdivisionME16:
		return "Podgorica"
	case SubdivisionME17:
		return "Rozaje"
	case SubdivisionME19:
		return "Tivat"
	case SubdivisionME20:
		return "Ulcinj"
	case SubdivisionME21:
		return "Zabljak"
	case SubdivisionME24:
		return "Tuzi"
	case SubdivisionMGA:
		return "Toamasina"
	case SubdivisionMGD:
		return "Antsiranana"
	case SubdivisionMGF:
		return "Fianarantsoa"
	case SubdivisionMGM:
		return "Mahajanga"
	case SubdivisionMGT:
		return "Antananarivo"
	case SubdivisionMGU:
		return "Toliara"
	case SubdivisionMHKWA:
		return "Kwajalein"
	case SubdivisionMHMAJ:
		return "Majuro"
	case SubdivisionMK101:
		return "Veles"
	case SubdivisionMK104:
		return "Kavadarci"
	case SubdivisionMK106:
		return "Negotino"
	case SubdivisionMK107:
		return "Rosoman"
	case SubdivisionMK108:
		return "Sveti Nikole"
	case SubdivisionMK109:
		return "Caska"
	case SubdivisionMK201:
		return "Berovo"
	case SubdivisionMK202:
		return "Vinica"
	case SubdivisionMK203:
		return "Delcevo"
	case SubdivisionMK205:
		return "Karbinci"
	case SubdivisionMK206:
		return "Kocani"
	case SubdivisionMK207:
		return "Makedonska Kamenica"
	case SubdivisionMK208:
		return "Pehcevo"
	case SubdivisionMK209:
		return "Probistip"
	case SubdivisionMK210:
		return "Cesinovo-Oblesevo"
	case SubdivisionMK211:
		return "Stip"
	case SubdivisionMK303:
		return "Debar"
	case SubdivisionMK307:
		return "Kicevo"
	case SubdivisionMK310:
		return "Ohrid"
	case SubdivisionMK311:
		return "Plasnica"
	case SubdivisionMK312:
		return "Struga"
	case SubdivisionMK313:
		return "Centar Zupa"
	case SubdivisionMK401:
		return "Bogdanci"
	case SubdivisionMK402:
		return "Bosilovo"
	case SubdivisionMK403:
		return "Valandovo"
	case SubdivisionMK404:
		return "Vasilevo"
	case SubdivisionMK405:
		return "Gevgelija"
	case SubdivisionMK406:
		return "Dojran"
	case SubdivisionMK408:
		return "Novo Selo"
	case SubdivisionMK409:
		return "Radovis"
	case SubdivisionMK410:
		return "Strumica"
	case SubdivisionMK501:
		return "Bitola"
	case SubdivisionMK503:
		return "Dolneni"
	case SubdivisionMK504:
		return "Krivogastani"
	case SubdivisionMK505:
		return "Krusevo"
	case SubdivisionMK506:
		return "Mogila"
	case SubdivisionMK507:
		return "Novaci"
	case SubdivisionMK508:
		return "Prilep"
	case SubdivisionMK509:
		return "Resen"
	case SubdivisionMK601:
		return "Bogovinje"
	case SubdivisionMK602:
		return "Brvenica"
	case SubdivisionMK603:
		return "Vrapciste"
	case SubdivisionMK604:
		return "Gostivar"
	case SubdivisionMK605:
		return "Zelino"
	case SubdivisionMK606:
		return "Jegunovce"
	case SubdivisionMK607:
		return "Mavrovo i Rostusa"
	case SubdivisionMK608:
		return "Tearce"
	case SubdivisionMK609:
		return "Tetovo"
	case SubdivisionMK701:
		return "Kratovo"
	case SubdivisionMK702:
		return "Kriva Palanka"
	case SubdivisionMK703:
		return "Kumanovo"
	case SubdivisionMK704:
		return "Lipkovo"
	case SubdivisionMK705:
		return "Rankovce"
	case SubdivisionMK802:
		return "Aracinovo"
	case SubdivisionMK803:
		return "Butel"
	case SubdivisionMK804:
		return "Gazi Baba"
	case SubdivisionMK806:
		return "Zelenikovo"
	case SubdivisionMK807:
		return "Ilinden"
	case SubdivisionMK809:
		return "Kisela Voda"
	case SubdivisionMK810:
		return "Petrovec"
	case SubdivisionMK811:
		return "Saraj"
	case SubdivisionMK812:
		return "Sopiste"
	case SubdivisionMK813:
		return "Studenicani"
	case SubdivisionMK814:
		return "Centar"
	case SubdivisionMK817:
		return "Suto Orizari"
	case SubdivisionML1:
		return "Kayes"
	case SubdivisionML2:
		return "Koulikoro"
	case SubdivisionML3:
		return "Sikasso"
	case SubdivisionML4:
		return "Segou"
	case SubdivisionML5:
		return "Mopti"
	case SubdivisionML6:
		return "Tombouctou"
	case SubdivisionML7:
		return "Gao"
	case SubdivisionML8:
		return "Kidal"
	case SubdivisionMLBKO:
		return "Bamako"
	case SubdivisionMM01:
		return "Sagaing"
	case SubdivisionMM02:
		return "Bago"
	case SubdivisionMM03:
		return "Magway"
	case SubdivisionMM04:
		return "Mandalay"
	case SubdivisionMM05:
		return "Tanintharyi"
	case SubdivisionMM06:
		return "Yangon"
	case SubdivisionMM07:
		return "Ayeyarwady"
	case SubdivisionMM11:
		return "Kachin"
	case SubdivisionMM12:
		return "Kayah"
	case SubdivisionMM13:
		return "Kayin"
	case SubdivisionMM15:
		return "Mon"
	case SubdivisionMM16:
		return "Rakhine"
	case SubdivisionMM17:
		return "Shan"
	case SubdivisionMM18:
		return "Nay Pyi Taw"
	case SubdivisionMN035:
		return "Orhon"
	case SubdivisionMN037:
		return "Darhan uul"
	case SubdivisionMN047:
		return "Tov"
	case SubdivisionMN049:
		return "Selenge"
	case SubdivisionMN055:
		return "Ovorhangay"
	case SubdivisionMN061:
		return "Dornod"
	case SubdivisionMN065:
		return "Govi-Altay"
	case SubdivisionMN071:
		return "Bayan-Olgiy"
	case SubdivisionMN1:
		return "Ulaanbaatar"
	case SubdivisionMR04:
		return "Gorgol"
	case SubdivisionMR06:
		return "Trarza"
	case SubdivisionMR08:
		return "Dakhlet Nouadhibou"
	case SubdivisionMR11:
		return "Tiris Zemmour"
	case SubdivisionMR12:
		return "Inchiri"
	case SubdivisionMR13:
		return "Nouakchott Ouest"
	case SubdivisionMT01:
		return "Attard"
	case SubdivisionMT02:
		return "Balzan"
	case SubdivisionMT03:
		return "Birgu"
	case SubdivisionMT04:
		return "Birkirkara"
	case SubdivisionMT05:
		return "Birzebbuga"
	case SubdivisionMT06:
		return "Bormla"
	case SubdivisionMT07:
		return "Dingli"
	case SubdivisionMT08:
		return "Fgura"
	case SubdivisionMT09:
		return "Floriana"
	case SubdivisionMT10:
		return "Fontana"
	case SubdivisionMT11:
		return "Gudja"
	case SubdivisionMT12:
		return "Gzira"
	case SubdivisionMT14:
		return "Gharb"
	case SubdivisionMT15:
		return "Gharghur"
	case SubdivisionMT16:
		return "Ghasri"
	case SubdivisionMT17:
		return "Ghaxaq"
	case SubdivisionMT18:
		return "Hamrun"
	case SubdivisionMT19:
		return "Iklin"
	case SubdivisionMT20:
		return "Isla"
	case SubdivisionMT21:
		return "Kalkara"
	case SubdivisionMT22:
		return "Kercem"
	case SubdivisionMT23:
		return "Kirkop"
	case SubdivisionMT24:
		return "Lija"
	case SubdivisionMT25:
		return "Luqa"
	case SubdivisionMT26:
		return "Marsa"
	case SubdivisionMT27:
		return "Marsaskala"
	case SubdivisionMT28:
		return "Marsaxlokk"
	case SubdivisionMT29:
		return "Mdina"
	case SubdivisionMT30:
		return "Mellieha"
	case SubdivisionMT31:
		return "Mgarr"
	case SubdivisionMT32:
		return "Mosta"
	case SubdivisionMT33:
		return "Mqabba"
	case SubdivisionMT34:
		return "Msida"
	case SubdivisionMT35:
		return "Mtarfa"
	case SubdivisionMT36:
		return "Munxar"
	case SubdivisionMT37:
		return "Nadur"
	case SubdivisionMT38:
		return "Naxxar"
	case SubdivisionMT39:
		return "Paola"
	case SubdivisionMT40:
		return "Pembroke"
	case SubdivisionMT41:
		return "Pieta"
	case SubdivisionMT42:
		return "Qala"
	case SubdivisionMT43:
		return "Qormi"
	case SubdivisionMT45:
		return "Rabat Gozo"
	case SubdivisionMT46:
		return "Rabat Malta"
	case SubdivisionMT48:
		return "Saint Julian's"
	case SubdivisionMT49:
		return "Saint John"
	case SubdivisionMT51:
		return "Saint Paul's Bay"
	case SubdivisionMT52:
		return "Sannat"
	case SubdivisionMT53:
		return "Saint Lucia's"
	case SubdivisionMT54:
		return "Santa Venera"
	case SubdivisionMT55:
		return "Siggiewi"
	case SubdivisionMT56:
		return "Sliema"
	case SubdivisionMT57:
		return "Swieqi"
	case SubdivisionMT58:
		return "Ta' Xbiex"
	case SubdivisionMT59:
		return "Tarxien"
	case SubdivisionMT60:
		return "Valletta"
	case SubdivisionMT61:
		return "Xaghra"
	case SubdivisionMT62:
		return "Xewkija"
	case SubdivisionMT63:
		return "Xghajra"
	case SubdivisionMT64:
		return "Zabbar"
	case SubdivisionMT65:
		return "Zebbug Gozo"
	case SubdivisionMT67:
		return "Zejtun"
	case SubdivisionMT68:
		return "Zurrieq"
	case SubdivisionMUBL:
		return "Black River"
	case SubdivisionMUFL:
		return "Flacq"
	case SubdivisionMUGP:
		return "Grand Port"
	case SubdivisionMUMO:
		return "Moka"
	case SubdivisionMUPA:
		return "Pamplemousses"
	case SubdivisionMUPL:
		return "Port Louis"
	case SubdivisionMUPW:
		return "Plaines Wilhems"
	case SubdivisionMURO:
		return "Rodrigues Islands"
	case SubdivisionMURR:
		return "Riviere du Rempart"
	case SubdivisionMUSA:
		return "Savanne"
	case SubdivisionMV00:
		return "South Ari Atoll"
	case SubdivisionMV01:
		return "Addu City"
	case SubdivisionMV03:
		return "Faadhippolhu"
	case SubdivisionMV04:
		return "Felidhu Atoll"
	case SubdivisionMV05:
		return "Hahdhunmathi"
	case SubdivisionMV07:
		return "North Thiladhunmathi"
	case SubdivisionMV12:
		return "Mulaku Atoll"
	case SubdivisionMV13:
		return "North Maalhosmadulu"
	case SubdivisionMV17:
		return "South Nilandhe Atoll"
	case SubdivisionMV20:
		return "South Maalhosmadulu"
	case SubdivisionMV25:
		return "South Miladhunmadulu"
	case SubdivisionMV28:
		return "South Huvadhu Atoll"
	case SubdivisionMVMLE:
		return "Male"
	case SubdivisionMWBA:
		return "Balaka"
	case SubdivisionMWBL:
		return "Blantyre"
	case SubdivisionMWCR:
		return "Chiradzulu"
	case SubdivisionMWDE:
		return "Dedza"
	case SubdivisionMWDO:
		return "Dowa"
	case SubdivisionMWKR:
		return "Karonga"
	case SubdivisionMWLI:
		return "Lilongwe"
	case SubdivisionMWMG:
		return "Mangochi"
	case SubdivisionMWMH:
		return "Machinga"
	case SubdivisionMWMZ:
		return "Mzimba"
	case SubdivisionMWNK:
		return "Nkhotakota"
	case SubdivisionMWSA:
		return "Salima"
	case SubdivisionMWZO:
		return "Zomba"
	case SubdivisionMXAGU:
		return "Aguascalientes"
	case SubdivisionMXBCN:
		return "Baja California"
	case SubdivisionMXBCS:
		return "Baja California Sur"
	case SubdivisionMXCAM:
		return "Campeche"
	case SubdivisionMXCHH:
		return "Chihuahua"
	case SubdivisionMXCHP:
		return "Chiapas"
	case SubdivisionMXCMX:
		return "Ciudad de Mexico"
	case SubdivisionMXCOA:
		return "Coahuila de Zaragoza"
	case SubdivisionMXCOL:
		return "Colima"
	case SubdivisionMXDUR:
		return "Durango"
	case SubdivisionMXGRO:
		return "Guerrero"
	case SubdivisionMXGUA:
		return "Guanajuato"
	case SubdivisionMXHID:
		return "Hidalgo"
	case SubdivisionMXJAL:
		return "Jalisco"
	case SubdivisionMXMEX:
		return "Mexico"
	case SubdivisionMXMIC:
		return "Michoacan de Ocampo"
	case SubdivisionMXMOR:
		return "Morelos"
	case SubdivisionMXNAY:
		return "Nayarit"
	case SubdivisionMXNLE:
		return "Nuevo Leon"
	case SubdivisionMXOAX:
		return "Oaxaca"
	case SubdivisionMXPUE:
		return "Puebla"
	case SubdivisionMXQUE:
		return "Queretaro"
	case SubdivisionMXROO:
		return "Quintana Roo"
	case SubdivisionMXSIN:
		return "Sinaloa"
	case SubdivisionMXSLP:
		return "San Luis Potosi"
	case SubdivisionMXSON:
		return "Sonora"
	case SubdivisionMXTAB:
		return "Tabasco"
	case SubdivisionMXTAM:
		return "Tamaulipas"
	case SubdivisionMXTLA:
		return "Tlaxcala"
	case SubdivisionMXVER:
		return "Veracruz de Ignacio de la Llave"
	case SubdivisionMXYUC:
		return "Yucatan"
	case SubdivisionMXZAC:
		return "Zacatecas"
	case SubdivisionMY01:
		return "Johor"
	case SubdivisionMY02:
		return "Kedah"
	case SubdivisionMY03:
		return "Kelantan"
	case SubdivisionMY04:
		return "Melaka"
	case SubdivisionMY05:
		return "Negeri Sembilan"
	case SubdivisionMY06:
		return "Pahang"
	case SubdivisionMY07:
		return "Pulau Pinang"
	case SubdivisionMY08:
		return "Perak"
	case SubdivisionMY09:
		return "Perlis"
	case SubdivisionMY10:
		return "Selangor"
	case SubdivisionMY11:
		return "Terengganu"
	case SubdivisionMY12:
		return "Sabah"
	case SubdivisionMY13:
		return "Sarawak"
	case SubdivisionMY14:
		return "Wilayah Persekutuan Kuala Lumpur"
	case SubdivisionMY15:
		return "Wilayah Persekutuan Labuan"
	case SubdivisionMY16:
		return "Wilayah Persekutuan Putrajaya"
	case SubdivisionMZA:
		return "Niassa"
	case SubdivisionMZB:
		return "Manica"
	case SubdivisionMZG:
		return "Gaza"
	case SubdivisionMZI:
		return "Inhambane"
	case SubdivisionMZL:
		return "Maputo"
	case SubdivisionMZN:
		return "Nampula"
	case SubdivisionMZP:
		return "Cabo Delgado"
	case SubdivisionMZQ:
		return "Zambezia"
	case SubdivisionMZS:
		return "Sofala"
	case SubdivisionMZT:
		return "Tete"
	case SubdivisionNACA:
		return "Zambezi"
	case SubdivisionNAER:
		return "Erongo"
	case SubdivisionNAHA:
		return "Hardap"
	case SubdivisionNAKA:
		return "Karas"
	case SubdivisionNAKE:
		return "Kavango East"
	case SubdivisionNAKH:
		return "Khomas"
	case SubdivisionNAKU:
		return "Kunene"
	case SubdivisionNAOD:
		return "Otjozondjupa"
	case SubdivisionNAOH:
		return "Omaheke"
	case SubdivisionNAON:
		return "Oshana"
	case SubdivisionNAOS:
		return "Omusati"
	case SubdivisionNAOT:
		return "Oshikoto"
	case SubdivisionNAOW:
		return "Ohangwena"
	case SubdivisionNE1:
		return "Agadez"
	case SubdivisionNE3:
		return "Dosso"
	case SubdivisionNE5:
		return "Tahoua"
	case SubdivisionNE7:
		return "Zinder"
	case SubdivisionNE8:
		return "Niamey"
	case SubdivisionNGAB:
		return "Abia"
	case SubdivisionNGAD:
		return "Adamawa"
	case SubdivisionNGAK:
		return "Akwa Ibom"
	case SubdivisionNGAN:
		return "Anambra"
	case SubdivisionNGBA:
		return "Bauchi"
	case SubdivisionNGBE:
		return "Benue"
	case SubdivisionNGBO:
		return "Borno"
	case SubdivisionNGBY:
		return "Bayelsa"
	case SubdivisionNGCR:
		return "Cross River"
	case SubdivisionNGDE:
		return "Delta"
	case SubdivisionNGED:
		return "Edo"
	case SubdivisionNGEK:
		return "Ekiti"
	case SubdivisionNGEN:
		return "Enugu"
	case SubdivisionNGFC:
		return "Abuja Federal Capital Territory"
	case SubdivisionNGGO:
		return "Gombe"
	case SubdivisionNGIM:
		return "Imo"
	case SubdivisionNGJI:
		return "Jigawa"
	case SubdivisionNGKD:
		return "Kaduna"
	case SubdivisionNGKE:
		return "Kebbi"
	case SubdivisionNGKN:
		return "Kano"
	case SubdivisionNGKO:
		return "Kogi"
	case SubdivisionNGKT:
		return "Katsina"
	case SubdivisionNGKW:
		return "Kwara"
	case SubdivisionNGLA:
		return "Lagos"
	case SubdivisionNGNA:
		return "Nasarawa"
	case SubdivisionNGNI:
		return "Niger"
	case SubdivisionNGOG:
		return "Ogun"
	case SubdivisionNGON:
		return "Ondo"
	case SubdivisionNGOS:
		return "Osun"
	case SubdivisionNGOY:
		return "Oyo"
	case SubdivisionNGPL:
		return "Plateau"
	case SubdivisionNGRI:
		return "Rivers"
	case SubdivisionNGSO:
		return "Sokoto"
	case SubdivisionNGTA:
		return "Taraba"
	case SubdivisionNGYO:
		return "Yobe"
	case SubdivisionNGZA:
		return "Zamfara"
	case SubdivisionNIAN:
		return "Costa Caribe Norte"
	case SubdivisionNIAS:
		return "Costa Caribe Sur"
	case SubdivisionNIBO:
		return "Boaco"
	case SubdivisionNICA:
		return "Carazo"
	case SubdivisionNICI:
		return "Chinandega"
	case SubdivisionNICO:
		return "Chontales"
	case SubdivisionNIES:
		return "Esteli"
	case SubdivisionNIGR:
		return "Granada"
	case SubdivisionNIJI:
		return "Jinotega"
	case SubdivisionNILE:
		return "Leon"
	case SubdivisionNIMD:
		return "Madriz"
	case SubdivisionNIMN:
		return "Managua"
	case SubdivisionNIMS:
		return "Masaya"
	case SubdivisionNIMT:
		return "Matagalpa"
	case SubdivisionNINS:
		return "Nueva Segovia"
	case SubdivisionNIRI:
		return "Rivas"
	case SubdivisionNISJ:
		return "Rio San Juan"
	case SubdivisionNLDR:
		return "Drenthe"
	case SubdivisionNLFL:
		return "Flevoland"
	case SubdivisionNLFR:
		return "Fryslan"
	case SubdivisionNLGE:
		return "Gelderland"
	case SubdivisionNLGR:
		return "Groningen"
	case SubdivisionNLLI:
		return "Limburg"
	case SubdivisionNLNB:
		return "Noord-Brabant"
	case SubdivisionNLNH:
		return "Noord-Holland"
	case SubdivisionNLOV:
		return "Overijssel"
	case SubdivisionNLUT:
		return "Utrecht"
	case SubdivisionNLZE:
		return "Zeeland"
	case SubdivisionNLZH:
		return "Zuid-Holland"
	case SubdivisionNO03:
		return "Oslo"
	case SubdivisionNO11:
		return "Rogaland"
	case SubdivisionNO15:
		return "More og Romsdal"
	case SubdivisionNO18:
		return "Nordland"
	case SubdivisionNO30:
		return "Viken"
	case SubdivisionNO34:
		return "Innlandet"
	case SubdivisionNO38:
		return "Vestfold og Telemark"
	case SubdivisionNO42:
		return "Agder"
	case SubdivisionNO46:
		return "Vestland"
	case SubdivisionNO50:
		return "Trondelag"
	case SubdivisionNO54:
		return "Troms og Finnmark"
	case SubdivisionNPBA:
		return "Bagmati"
	case SubdivisionNPBH:
		return "Bheri"
	case SubdivisionNPDH:
		return "Dhawalagiri"
	case SubdivisionNPGA:
		return "Gandaki"
	case SubdivisionNPJA:
		return "Janakpur"
	case SubdivisionNPKA:
		return "Karnali"
	case SubdivisionNPKO:
		return "Kosi"
	case SubdivisionNPLU:
		return "Lumbini"
	case SubdivisionNPMA:
		return "Mahakali"
	case SubdivisionNPME:
		return "Mechi"
	case SubdivisionNPNA:
		return "Narayani"
	case SubdivisionNPRA:
		return "Rapti"
	case SubdivisionNPSA:
		return "Sagarmatha"
	case SubdivisionNPSE:
		return "Seti"
	case SubdivisionNR14:
		return "Yaren"
	case SubdivisionNZAUK:
		return "Auckland"
	case SubdivisionNZBOP:
		return "Bay of Plenty"
	case SubdivisionNZCAN:
		return "Canterbury"
	case SubdivisionNZCIT:
		return "Chatham Islands Territory"
	case SubdivisionNZGIS:
		return "Gisborne"
	case SubdivisionNZHKB:
		return "Hawke's Bay"
	case SubdivisionNZMBH:
		return "Marlborough"
	case SubdivisionNZMWT:
		return "Manawatu-Wanganui"
	case SubdivisionNZNSN:
		return "Nelson"
	case SubdivisionNZNTL:
		return "Northland"
	case SubdivisionNZOTA:
		return "Otago"
	case SubdivisionNZSTL:
		return "Southland"
	case SubdivisionNZTAS:
		return "Tasman"
	case SubdivisionNZTKI:
		return "Taranaki"
	case SubdivisionNZWGN:
		return "Wellington"
	case SubdivisionNZWKO:
		return "Waikato"
	case SubdivisionNZWTC:
		return "West Coast"
	case SubdivisionOMBJ:
		return "Janub al Batinah"
	case SubdivisionOMBS:
		return "Shamal al Batinah"
	case SubdivisionOMBU:
		return "Al Buraymi"
	case SubdivisionOMDA:
		return "Ad Dakhiliyah"
	case SubdivisionOMMA:
		return "Masqat"
	case SubdivisionOMMU:
		return "Musandam"
	case SubdivisionOMSJ:
		return "Janub ash Sharqiyah"
	case SubdivisionOMSS:
		return "Shamal ash Sharqiyah"
	case SubdivisionOMWU:
		return "Al Wusta"
	case SubdivisionOMZA:
		return "Az Zahirah"
	case SubdivisionOMZU:
		return "Zufar"
	case SubdivisionPA1:
		return "Bocas del Toro"
	case SubdivisionPA2:
		return "Cocle"
	case SubdivisionPA3:
		return "Colon"
	case SubdivisionPA4:
		return "Chiriqui"
	case SubdivisionPA5:
		return "Darien"
	case SubdivisionPA6:
		return "Herrera"
	case SubdivisionPA7:
		return "Los Santos"
	case SubdivisionPA8:
		return "Panama"
	case SubdivisionPA9:
		return "Veraguas"
	case SubdivisionPANB:
		return "Ngobe-Bugle"
	case SubdivisionPEAMA:
		return "Amazonas"
	case SubdivisionPEANC:
		return "Ancash"
	case SubdivisionPEAPU:
		return "Apurimac"
	case SubdivisionPEARE:
		return "Arequipa"
	case SubdivisionPEAYA:
		return "Ayacucho"
	case SubdivisionPECAJ:
		return "Cajamarca"
	case SubdivisionPECAL:
		return "El Callao"
	case SubdivisionPECUS:
		return "Cusco"
	case SubdivisionPEHUC:
		return "Huanuco"
	case SubdivisionPEHUV:
		return "Huancavelica"
	case SubdivisionPEICA:
		return "Ica"
	case SubdivisionPEJUN:
		return "Junin"
	case SubdivisionPELAL:
		return "La Libertad"
	case SubdivisionPELAM:
		return "Lambayeque"
	case SubdivisionPELIM:
		return "Lima"
	case SubdivisionPELOR:
		return "Loreto"
	case SubdivisionPEMDD:
		return "Madre de Dios"
	case SubdivisionPEMOQ:
		return "Moquegua"
	case SubdivisionPEPAS:
		return "Pasco"
	case SubdivisionPEPIU:
		return "Piura"
	case SubdivisionPEPUN:
		return "Puno"
	case SubdivisionPESAM:
		return "San Martin"
	case SubdivisionPETAC:
		return "Tacna"
	case SubdivisionPETUM:
		return "Tumbes"
	case SubdivisionPEUCA:
		return "Ucayali"
	case SubdivisionPGCPM:
		return "Central"
	case SubdivisionPGEBR:
		return "East New Britain"
	case SubdivisionPGEHG:
		return "Eastern Highlands"
	case SubdivisionPGESW:
		return "East Sepik"
	case SubdivisionPGMBA:
		return "Milne Bay"
	case SubdivisionPGMPL:
		return "Morobe"
	case SubdivisionPGMPM:
		return "Madang"
	case SubdivisionPGMRL:
		return "Manus"
	case SubdivisionPGNCD:
		return "National Capital District (Port Moresby)"
	case SubdivisionPGNIK:
		return "New Ireland"
	case SubdivisionPGNSB:
		return "Bougainville"
	case SubdivisionPGSAN:
		return "West Sepik"
	case SubdivisionPGSHM:
		return "Southern Highlands"
	case SubdivisionPGWBK:
		return "West New Britain"
	case SubdivisionPGWHM:
		return "Western Highlands"
	case SubdivisionPGWPD:
		return "Western"
	case SubdivisionPH00:
		return "National Capital Region"
	case SubdivisionPHABR:
		return "Abra"
	case SubdivisionPHAGN:
		return "Agusan del Norte"
	case SubdivisionPHAGS:
		return "Agusan del Sur"
	case SubdivisionPHAKL:
		return "Aklan"
	case SubdivisionPHALB:
		return "Albay"
	case SubdivisionPHANT:
		return "Antique"
	case SubdivisionPHAPA:
		return "Apayao"
	case SubdivisionPHAUR:
		return "Aurora"
	case SubdivisionPHBAN:
		return "Bataan"
	case SubdivisionPHBAS:
		return "Basilan"
	case SubdivisionPHBEN:
		return "Benguet"
	case SubdivisionPHBIL:
		return "Biliran"
	case SubdivisionPHBOH:
		return "Bohol"
	case SubdivisionPHBTG:
		return "Batangas"
	case SubdivisionPHBTN:
		return "Batanes"
	case SubdivisionPHBUK:
		return "Bukidnon"
	case SubdivisionPHBUL:
		return "Bulacan"
	case SubdivisionPHCAG:
		return "Cagayan"
	case SubdivisionPHCAM:
		return "Camiguin"
	case SubdivisionPHCAN:
		return "Camarines Norte"
	case SubdivisionPHCAP:
		return "Capiz"
	case SubdivisionPHCAS:
		return "Camarines Sur"
	case SubdivisionPHCAT:
		return "Catanduanes"
	case SubdivisionPHCAV:
		return "Cavite"
	case SubdivisionPHCEB:
		return "Cebu"
	case SubdivisionPHCOM:
		return "Davao de Oro"
	case SubdivisionPHDAO:
		return "Davao Oriental"
	case SubdivisionPHDAS:
		return "Davao del Sur"
	case SubdivisionPHDAV:
		return "Davao del Norte"
	case SubdivisionPHDIN:
		return "Dinagat Islands"
	case SubdivisionPHEAS:
		return "Eastern Samar"
	case SubdivisionPHGUI:
		return "Guimaras"
	case SubdivisionPHIFU:
		return "Ifugao"
	case SubdivisionPHILI:
		return "Iloilo"
	case SubdivisionPHILN:
		return "Ilocos Norte"
	case SubdivisionPHILS:
		return "Ilocos Sur"
	case SubdivisionPHISA:
		return "Isabela"
	case SubdivisionPHKAL:
		return "Kalinga"
	case SubdivisionPHLAG:
		return "Laguna"
	case SubdivisionPHLAN:
		return "Lanao del Norte"
	case SubdivisionPHLAS:
		return "Lanao del Sur"
	case SubdivisionPHLEY:
		return "Leyte"
	case SubdivisionPHLUN:
		return "La Union"
	case SubdivisionPHMAD:
		return "Marinduque"
	case SubdivisionPHMAG:
		return "Maguindanao"
	case SubdivisionPHMAS:
		return "Masbate"
	case SubdivisionPHMDC:
		return "Mindoro Occidental"
	case SubdivisionPHMDR:
		return "Mindoro Oriental"
	case SubdivisionPHMOU:
		return "Mountain Province"
	case SubdivisionPHMSC:
		return "Misamis Occidental"
	case SubdivisionPHMSR:
		return "Misamis Oriental"
	case SubdivisionPHNCO:
		return "Cotabato"
	case SubdivisionPHNEC:
		return "Negros Occidental"
	case SubdivisionPHNER:
		return "Negros Oriental"
	case SubdivisionPHNSA:
		return "Northern Samar"
	case SubdivisionPHNUE:
		return "Nueva Ecija"
	case SubdivisionPHNUV:
		return "Nueva Vizcaya"
	case SubdivisionPHPAM:
		return "Pampanga"
	case SubdivisionPHPAN:
		return "Pangasinan"
	case SubdivisionPHPLW:
		return "Palawan"
	case SubdivisionPHQUE:
		return "Quezon"
	case SubdivisionPHQUI:
		return "Quirino"
	case SubdivisionPHRIZ:
		return "Rizal"
	case SubdivisionPHROM:
		return "Romblon"
	case SubdivisionPHSAR:
		return "Sarangani"
	case SubdivisionPHSCO:
		return "South Cotabato"
	case SubdivisionPHSIG:
		return "Siquijor"
	case SubdivisionPHSLE:
		return "Southern Leyte"
	case SubdivisionPHSLU:
		return "Sulu"
	case SubdivisionPHSOR:
		return "Sorsogon"
	case SubdivisionPHSUK:
		return "Sultan Kudarat"
	case SubdivisionPHSUN:
		return "Surigao del Norte"
	case SubdivisionPHSUR:
		return "Surigao del Sur"
	case SubdivisionPHTAR:
		return "Tarlac"
	case SubdivisionPHTAW:
		return "Tawi-Tawi"
	case SubdivisionPHWSA:
		return "Samar"
	case SubdivisionPHZAN:
		return "Zamboanga del Norte"
	case SubdivisionPHZAS:
		return "Zamboanga del Sur"
	case SubdivisionPHZMB:
		return "Zambales"
	case SubdivisionPHZSI:
		return "Zamboanga Sibugay"
	case SubdivisionPKBA:
		return "Balochistan"
	case SubdivisionPKGB:
		return "Gilgit-Baltistan"
	case SubdivisionPKIS:
		return "Islamabad"
	case SubdivisionPKJK:
		return "Azad Jammu and Kashmir"
	case SubdivisionPKKP:
		return "Khyber Pakhtunkhwa"
	case SubdivisionPKPB:
		return "Punjab"
	case SubdivisionPKSD:
		return "Sindh"
	case SubdivisionPL02:
		return "Dolnoslaskie"
	case SubdivisionPL04:
		return "Kujawsko-pomorskie"
	case SubdivisionPL06:
		return "Lubelskie"
	case SubdivisionPL08:
		return "Lubuskie"
	case SubdivisionPL10:
		return "Lodzkie"
	case SubdivisionPL12:
		return "Malopolskie"
	case SubdivisionPL14:
		return "Mazowieckie"
	case SubdivisionPL16:
		return "Opolskie"
	case SubdivisionPL18:
		return "Podkarpackie"
	case SubdivisionPL20:
		return "Podlaskie"
	case SubdivisionPL22:
		return "Pomorskie"
	case SubdivisionPL24:
		return "Slaskie"
	case SubdivisionPL26:
		return "Swietokrzyskie"
	case SubdivisionPL28:
		return "Warminsko-mazurskie"
	case SubdivisionPL30:
		return "Wielkopolskie"
	case SubdivisionPL32:
		return "Zachodniopomorskie"
	case SubdivisionPSBTH:
		return "Bethlehem"
	case SubdivisionPSDEB:
		return "Deir El Balah"
	case SubdivisionPSGZA:
		return "Gaza"
	case SubdivisionPSHBN:
		return "Hebron"
	case SubdivisionPSJEM:
		return "Jerusalem"
	case SubdivisionPSJEN:
		return "Jenin"
	case SubdivisionPSJRH:
		return "Jericho and Al Aghwar"
	case SubdivisionPSKYS:
		return "Khan Yunis"
	case SubdivisionPSNBS:
		return "Nablus"
	case SubdivisionPSQQA:
		return "Qalqilya"
	case SubdivisionPSRBH:
		return "Ramallah"
	case SubdivisionPSRFH:
		return "Rafah"
	case SubdivisionPSSLT:
		return "Salfit"
	case SubdivisionPSTBS:
		return "Tubas"
	case SubdivisionPSTKM:
		return "Tulkarm"
	case SubdivisionPT01:
		return "Aveiro"
	case SubdivisionPT02:
		return "Beja"
	case SubdivisionPT03:
		return "Braga"
	case SubdivisionPT04:
		return "Braganca"
	case SubdivisionPT05:
		return "Castelo Branco"
	case SubdivisionPT06:
		return "Coimbra"
	case SubdivisionPT07:
		return "Evora"
	case SubdivisionPT08:
		return "Faro"
	case SubdivisionPT09:
		return "Guarda"
	case SubdivisionPT10:
		return "Leiria"
	case SubdivisionPT11:
		return "Lisboa"
	case SubdivisionPT12:
		return "Portalegre"
	case SubdivisionPT13:
		return "Porto"
	case SubdivisionPT14:
		return "Santarem"
	case SubdivisionPT15:
		return "Setubal"
	case SubdivisionPT16:
		return "Viana do Castelo"
	case SubdivisionPT17:
		return "Vila Real"
	case SubdivisionPT18:
		return "Viseu"
	case SubdivisionPT20:
		return "Regiao Autonoma dos Acores"
	case SubdivisionPT30:
		return "Regiao Autonoma da Madeira"
	case SubdivisionPW004:
		return "Airai"
	case SubdivisionPW100:
		return "Kayangel"
	case SubdivisionPW150:
		return "Koror"
	case SubdivisionPW212:
		return "Melekeok"
	case SubdivisionPW214:
		return "Ngaraard"
	case SubdivisionPW222:
		return "Ngardmau"
	case SubdivisionPY10:
		return "Alto Parana"
	case SubdivisionPY11:
		return "Central"
	case SubdivisionPY12:
		return "Neembucu"
	case SubdivisionPY13:
		return "Amambay"
	case SubdivisionPY14:
		return "Canindeyu"
	case SubdivisionPY15:
		return "Presidente Hayes"
	case SubdivisionPY16:
		return "Alto Paraguay"
	case SubdivisionPY19:
		return "Boqueron"
	case SubdivisionPY1:
		return "Concepcion"
	case SubdivisionPY2:
		return "San Pedro"
	case SubdivisionPY3:
		return "Cordillera"
	case SubdivisionPY4:
		return "Guaira"
	case SubdivisionPY5:
		return "Caaguazu"
	case SubdivisionPY6:
		return "Caazapa"
	case SubdivisionPY7:
		return "Itapua"
	case SubdivisionPY8:
		return "Misiones"
	case SubdivisionPY9:
		return "Paraguari"
	case SubdivisionPYASU:
		return "Asuncion"
	case SubdivisionQADA:
		return "Ad Dawhah"
	case SubdivisionQAKH:
		return "Al Khawr wa adh Dhakhirah"
	case SubdivisionQAMS:
		return "Ash Shamal"
	case SubdivisionQARA:
		return "Ar Rayyan"
	case SubdivisionQAUS:
		return "Umm Salal"
	case SubdivisionQAWA:
		return "Al Wakrah"
	case SubdivisionQAZA:
		return "Az Za'ayin"
	case SubdivisionROAB:
		return "Alba"
	case SubdivisionROAG:
		return "Arges"
	case SubdivisionROAR:
		return "Arad"
	case SubdivisionROB:
		return "Bucuresti"
	case SubdivisionROBC:
		return "Bacau"
	case SubdivisionROBH:
		return "Bihor"
	case SubdivisionROBN:
		return "Bistrita-Nasaud"
	case SubdivisionROBR:
		return "Braila"
	case SubdivisionROBT:
		return "Botosani"
	case SubdivisionROBV:
		return "Brasov"
	case SubdivisionROBZ:
		return "Buzau"
	case SubdivisionROCJ:
		return "Cluj"
	case SubdivisionROCL:
		return "Calarasi"
	case SubdivisionROCS:
		return "Caras-Severin"
	case SubdivisionROCT:
		return "Constanta"
	case SubdivisionROCV:
		return "Covasna"
	case SubdivisionRODB:
		return "Dambovita"
	case SubdivisionRODJ:
		return "Dolj"
	case SubdivisionROGJ:
		return "Gorj"
	case SubdivisionROGL:
		return "Galati"
	case SubdivisionROGR:
		return "Giurgiu"
	case SubdivisionROHD:
		return "Hunedoara"
	case SubdivisionROHR:
		return "Harghita"
	case SubdivisionROIF:
		return "Ilfov"
	case SubdivisionROIL:
		return "Ialomita"
	case SubdivisionROIS:
		return "Iasi"
	case SubdivisionROMH:
		return "Mehedinti"
	case SubdivisionROMM:
		return "Maramures"
	case SubdivisionROMS:
		return "Mures"
	case SubdivisionRONT:
		return "Neamt"
	case SubdivisionROOT:
		return "Olt"
	case SubdivisionROPH:
		return "Prahova"
	case SubdivisionROSB:
		return "Sibiu"
	case SubdivisionROSJ:
		return "Salaj"
	case SubdivisionROSM:
		return "Satu Mare"
	case SubdivisionROSV:
		return "Suceava"
	case SubdivisionROTL:
		return "Tulcea"
	case SubdivisionROTM:
		return "Timis"
	case SubdivisionROTR:
		return "Teleorman"
	case SubdivisionROVL:
		return "Valcea"
	case SubdivisionROVN:
		return "Vrancea"
	case SubdivisionROVS:
		return "Vaslui"
	case SubdivisionRS00:
		return "Beograd"
	case SubdivisionRS01:
		return "Severnobacki okrug"
	case SubdivisionRS02:
		return "Srednjebanatski okrug"
	case SubdivisionRS03:
		return "Severnobanatski okrug"
	case SubdivisionRS04:
		return "Juznobanatski okrug"
	case SubdivisionRS05:
		return "Zapadnobacki okrug"
	case SubdivisionRS06:
		return "Juznobacki okrug"
	case SubdivisionRS07:
		return "Sremski okrug"
	case SubdivisionRS08:
		return "Macvanski okrug"
	case SubdivisionRS09:
		return "Kolubarski okrug"
	case SubdivisionRS10:
		return "Podunavski okrug"
	case SubdivisionRS11:
		return "Branicevski okrug"
	case SubdivisionRS12:
		return "Sumadijski okrug"
	case SubdivisionRS13:
		return "Pomoravski okrug"
	case SubdivisionRS14:
		return "Borski okrug"
	case SubdivisionRS15:
		return "Zajecarski okrug"
	case SubdivisionRS16:
		return "Zlatiborski okrug"
	case SubdivisionRS17:
		return "Moravicki okrug"
	case SubdivisionRS18:
		return "Raski okrug"
	case SubdivisionRS19:
		return "Rasinski okrug"
	case SubdivisionRS20:
		return "Nisavski okrug"
	case SubdivisionRS21:
		return "Toplicki okrug"
	case SubdivisionRS22:
		return "Pirotski okrug"
	case SubdivisionRS23:
		return "Jablanicki okrug"
	case SubdivisionRS24:
		return "Pcinjski okrug"
	case SubdivisionRS26:
		return "Pecki okrug"
	case SubdivisionRS27:
		return "Prizrenski okrug"
	case SubdivisionRS28:
		return "Kosovsko-Mitrovacki okrug"
	case SubdivisionRUAD:
		return "Adygeya, Respublika"
	case SubdivisionRUAL:
		return "Altay, Respublika"
	case SubdivisionRUALT:
		return "Altayskiy kray"
	case SubdivisionRUAMU:
		return "Amurskaya oblast'"
	case SubdivisionRUARK:
		return "Arkhangel'skaya oblast'"
	case SubdivisionRUAST:
		return "Astrakhanskaya oblast'"
	case SubdivisionRUBA:
		return "Bashkortostan, Respublika"
	case SubdivisionRUBEL:
		return "Belgorodskaya oblast'"
	case SubdivisionRUBRY:
		return "Bryanskaya oblast'"
	case SubdivisionRUBU:
		return "Buryatiya, Respublika"
	case SubdivisionRUCE:
		return "Chechenskaya Respublika"
	case SubdivisionRUCHE:
		return "Chelyabinskaya oblast'"
	case SubdivisionRUCHU:
		return "Chukotskiy avtonomnyy okrug"
	case SubdivisionRUCU:
		return "Chuvashskaya Respublika"
	case SubdivisionRUDA:
		return "Dagestan, Respublika"
	case SubdivisionRUIN:
		return "Ingushetiya, Respublika"
	case SubdivisionRUIRK:
		return "Irkutskaya oblast'"
	case SubdivisionRUIVA:
		return "Ivanovskaya oblast'"
	case SubdivisionRUKAM:
		return "Kamchatskiy kray"
	case SubdivisionRUKB:
		return "Kabardino-Balkarskaya Respublika"
	case SubdivisionRUKC:
		return "Karachayevo-Cherkesskaya Respublika"
	case SubdivisionRUKDA:
		return "Krasnodarskiy kray"
	case SubdivisionRUKEM:
		return "Kemerovskaya oblast'"
	case SubdivisionRUKGD:
		return "Kaliningradskaya oblast'"
	case SubdivisionRUKGN:
		return "Kurganskaya oblast'"
	case SubdivisionRUKHA:
		return "Khabarovskiy kray"
	case SubdivisionRUKHM:
		return "Khanty-Mansiyskiy avtonomnyy okrug"
	case SubdivisionRUKIR:
		return "Kirovskaya oblast'"
	case SubdivisionRUKK:
		return "Khakasiya, Respublika"
	case SubdivisionRUKL:
		return "Kalmykiya, Respublika"
	case SubdivisionRUKLU:
		return "Kaluzhskaya oblast'"
	case SubdivisionRUKO:
		return "Komi, Respublika"
	case SubdivisionRUKOS:
		return "Kostromskaya oblast'"
	case SubdivisionRUKR:
		return "Kareliya, Respublika"
	case SubdivisionRUKRS:
		return "Kurskaya oblast'"
	case SubdivisionRUKYA:
		return "Krasnoyarskiy kray"
	case SubdivisionRULEN:
		return "Leningradskaya oblast'"
	case SubdivisionRULIP:
		return "Lipetskaya oblast'"
	case SubdivisionRUMAG:
		return "Magadanskaya oblast'"
	case SubdivisionRUME:
		return "Mariy El, Respublika"
	case SubdivisionRUMO:
		return "Mordoviya, Respublika"
	case SubdivisionRUMOS:
		return "Moskovskaya oblast'"
	case SubdivisionRUMOW:
		return "Moskva"
	case SubdivisionRUMUR:
		return "Murmanskaya oblast'"
	case SubdivisionRUNEN:
		return "Nenetskiy avtonomnyy okrug"
	case SubdivisionRUNGR:
		return "Novgorodskaya oblast'"
	case SubdivisionRUNIZ:
		return "Nizhegorodskaya oblast'"
	case SubdivisionRUNVS:
		return "Novosibirskaya oblast'"
	case SubdivisionRUOMS:
		return "Omskaya oblast'"
	case SubdivisionRUORE:
		return "Orenburgskaya oblast'"
	case SubdivisionRUORL:
		return "Orlovskaya oblast'"
	case SubdivisionRUPER:
		return "Permskiy kray"
	case SubdivisionRUPNZ:
		return "Penzenskaya oblast'"
	case SubdivisionRUPRI:
		return "Primorskiy kray"
	case SubdivisionRUPSK:
		return "Pskovskaya oblast'"
	case SubdivisionRUROS:
		return "Rostovskaya oblast'"
	case SubdivisionRURYA:
		return "Ryazanskaya oblast'"
	case SubdivisionRUSA:
		return "Saha, Respublika"
	case SubdivisionRUSAK:
		return "Sakhalinskaya oblast'"
	case SubdivisionRUSAM:
		return "Samarskaya oblast'"
	case SubdivisionRUSAR:
		return "Saratovskaya oblast'"
	case SubdivisionRUSE:
		return "Severnaya Osetiya, Respublika"
	case SubdivisionRUSMO:
		return "Smolenskaya oblast'"
	case SubdivisionRUSPE:
		return "Sankt-Peterburg"
	case SubdivisionRUSTA:
		return "Stavropol'skiy kray"
	case SubdivisionRUSVE:
		return "Sverdlovskaya oblast'"
	case SubdivisionRUTA:
		return "Tatarstan, Respublika"
	case SubdivisionRUTAM:
		return "Tambovskaya oblast'"
	case SubdivisionRUTOM:
		return "Tomskaya oblast'"
	case SubdivisionRUTUL:
		return "Tul'skaya oblast'"
	case SubdivisionRUTVE:
		return "Tverskaya oblast'"
	case SubdivisionRUTY:
		return "Tyva, Respublika"
	case SubdivisionRUTYU:
		return "Tyumenskaya oblast'"
	case SubdivisionRUUD:
		return "Udmurtskaya Respublika"
	case SubdivisionRUULY:
		return "Ul'yanovskaya oblast'"
	case SubdivisionRUVGG:
		return "Volgogradskaya oblast'"
	case SubdivisionRUVLA:
		return "Vladimirskaya oblast'"
	case SubdivisionRUVLG:
		return "Vologodskaya oblast'"
	case SubdivisionRUVOR:
		return "Voronezhskaya oblast'"
	case SubdivisionRUYAN:
		return "Yamalo-Nenetskiy avtonomnyy okrug"
	case SubdivisionRUYAR:
		return "Yaroslavskaya oblast'"
	case SubdivisionRUYEV:
		return "Yevreyskaya avtonomnaya oblast'"
	case SubdivisionRUZAB:
		return "Zabaykal'skiy kray"
	case SubdivisionRW01:
		return "Ville de Kigali"
	case SubdivisionRW02:
		return "Est"
	case SubdivisionRW03:
		return "Nord"
	case SubdivisionRW04:
		return "Ouest"
	case SubdivisionRW05:
		return "Sud"
	case SubdivisionSA01:
		return "Ar Riyad"
	case SubdivisionSA02:
		return "Makkah al Mukarramah"
	case SubdivisionSA03:
		return "Al Madinah al Munawwarah"
	case SubdivisionSA04:
		return "Ash Sharqiyah"
	case SubdivisionSA05:
		return "Al Qasim"
	case SubdivisionSA06:
		return "Ha'il"
	case SubdivisionSA07:
		return "Tabuk"
	case SubdivisionSA08:
		return "Al Hudud ash Shamaliyah"
	case SubdivisionSA09:
		return "Jazan"
	case SubdivisionSA10:
		return "Najran"
	case SubdivisionSA11:
		return "Al Bahah"
	case SubdivisionSA12:
		return "Al Jawf"
	case SubdivisionSA14:
		return "'Asir"
	case SubdivisionSBCH:
		return "Choiseul"
	case SubdivisionSBGU:
		return "Guadalcanal"
	case SubdivisionSBWE:
		return "Western"
	case SubdivisionSC01:
		return "Anse aux Pins"
	case SubdivisionSC02:
		return "Anse Boileau"
	case SubdivisionSC06:
		return "Baie Lazare"
	case SubdivisionSC07:
		return "Baie Sainte Anne"
	case SubdivisionSC08:
		return "Beau Vallon"
	case SubdivisionSC10:
		return "Bel Ombre"
	case SubdivisionSC11:
		return "Cascade"
	case SubdivisionSC13:
		return "Grand Anse Mahe"
	case SubdivisionSC14:
		return "Grand Anse Praslin"
	case SubdivisionSC15:
		return "La Digue"
	case SubdivisionSC16:
		return "English River"
	case SubdivisionSC20:
		return "Pointe Larue"
	case SubdivisionSC23:
		return "Takamaka"
	case SubdivisionSDDC:
		return "Central Darfur"
	case SubdivisionSDDN:
		return "North Darfur"
	case SubdivisionSDDS:
		return "South Darfur"
	case SubdivisionSDDW:
		return "West Darfur"
	case SubdivisionSDGD:
		return "Gedaref"
	case SubdivisionSDGK:
		return "West Kordofan"
	case SubdivisionSDGZ:
		return "Gezira"
	case SubdivisionSDKA:
		return "Kassala"
	case SubdivisionSDKH:
		return "Khartoum"
	case SubdivisionSDKN:
		return "North Kordofan"
	case SubdivisionSDKS:
		return "South Kordofan"
	case SubdivisionSDNB:
		return "Blue Nile"
	case SubdivisionSDNO:
		return "Northern"
	case SubdivisionSDNR:
		return "River Nile"
	case SubdivisionSDNW:
		return "White Nile"
	case SubdivisionSDRS:
		return "Red Sea"
	case SubdivisionSDSI:
		return "Sennar"
	case SubdivisionSEAB:
		return "Stockholms lan"
	case SubdivisionSEAC:
		return "Vasterbottens lan"
	case SubdivisionSEBD:
		return "Norrbottens lan"
	case SubdivisionSEC:
		return "Uppsala lan"
	case SubdivisionSED:
		return "Sodermanlands lan"
	case SubdivisionSEE:
		return "Ostergotlands lan"
	case SubdivisionSEF:
		return "Jonkopings lan"
	case SubdivisionSEG:
		return "Kronobergs lan"
	case SubdivisionSEH:
		return "Kalmar lan"
	case SubdivisionSEI:
		return "Gotlands lan"
	case SubdivisionSEK:
		return "Blekinge lan"
	case SubdivisionSEM:
		return "Skane lan"
	case SubdivisionSEN:
		return "Hallands lan"
	case SubdivisionSEO:
		return "Vastra Gotalands lan"
	case SubdivisionSES:
		return "Varmlands lan"
	case SubdivisionSET:
		return "Orebro lan"
	case SubdivisionSEU:
		return "Vastmanlands lan"
	case SubdivisionSEW:
		return "Dalarnas lan"
	case SubdivisionSEX:
		return "Gavleborgs lan"
	case SubdivisionSEY:
		return "Vasternorrlands lan"
	case SubdivisionSEZ:
		return "Jamtlands lan"
	case SubdivisionSG01:
		return "Central Singapore"
	case SubdivisionSG02:
		return "North East"
	case SubdivisionSG03:
		return "North West"
	case SubdivisionSG04:
		return "South East"
	case SubdivisionSG05:
		return "South West"
	case SubdivisionSHHL:
		return "Saint Helena"
	case SubdivisionSI001:
		return "Ajdovscina"
	case SubdivisionSI002:
		return "Beltinci"
	case SubdivisionSI003:
		return "Bled"
	case SubdivisionSI004:
		return "Bohinj"
	case SubdivisionSI005:
		return "Borovnica"
	case SubdivisionSI006:
		return "Bovec"
	case SubdivisionSI007:
		return "Brda"
	case SubdivisionSI008:
		return "Brezovica"
	case SubdivisionSI009:
		return "Brezice"
	case SubdivisionSI010:
		return "Tisina"
	case SubdivisionSI011:
		return "Celje"
	case SubdivisionSI012:
		return "Cerklje na Gorenjskem"
	case SubdivisionSI013:
		return "Cerknica"
	case SubdivisionSI014:
		return "Cerkno"
	case SubdivisionSI015:
		return "Crensovci"
	case SubdivisionSI017:
		return "Crnomelj"
	case SubdivisionSI018:
		return "Destrnik"
	case SubdivisionSI019:
		return "Divaca"
	case SubdivisionSI020:
		return "Dobrepolje"
	case SubdivisionSI021:
		return "Dobrova-Polhov Gradec"
	case SubdivisionSI023:
		return "Domzale"
	case SubdivisionSI024:
		return "Dornava"
	case SubdivisionSI025:
		return "Dravograd"
	case SubdivisionSI026:
		return "Duplek"
	case SubdivisionSI029:
		return "Gornja Radgona"
	case SubdivisionSI031:
		return "Gornji Petrovci"
	case SubdivisionSI032:
		return "Grosuplje"
	case SubdivisionSI033:
		return "Salovci"
	case SubdivisionSI034:
		return "Hrastnik"
	case SubdivisionSI035:
		return "Hrpelje-Kozina"
	case SubdivisionSI036:
		return "Idrija"
	case SubdivisionSI037:
		return "Ig"
	case SubdivisionSI038:
		return "Ilirska Bistrica"
	case SubdivisionSI039:
		return "Ivancna Gorica"
	case SubdivisionSI040:
		return "Izola"
	case SubdivisionSI041:
		return "Jesenice"
	case SubdivisionSI042:
		return "Jursinci"
	case SubdivisionSI043:
		return "Kamnik"
	case SubdivisionSI044:
		return "Kanal"
	case SubdivisionSI045:
		return "Kidricevo"
	case SubdivisionSI046:
		return "Kobarid"
	case SubdivisionSI047:
		return "Kobilje"
	case SubdivisionSI048:
		return "Kocevje"
	case SubdivisionSI049:
		return "Komen"
	case SubdivisionSI050:
		return "Koper"
	case SubdivisionSI052:
		return "Kranj"
	case SubdivisionSI053:
		return "Kranjska Gora"
	case SubdivisionSI054:
		return "Krsko"
	case SubdivisionSI055:
		return "Kungota"
	case SubdivisionSI056:
		return "Kuzma"
	case SubdivisionSI057:
		return "Lasko"
	case SubdivisionSI058:
		return "Lenart"
	case SubdivisionSI059:
		return "Lendava"
	case SubdivisionSI060:
		return "Litija"
	case SubdivisionSI061:
		return "Ljubljana"
	case SubdivisionSI063:
		return "Ljutomer"
	case SubdivisionSI064:
		return "Logatec"
	case SubdivisionSI065:
		return "Loska dolina"
	case SubdivisionSI066:
		return "Loski Potok"
	case SubdivisionSI067:
		return "Luce"
	case SubdivisionSI068:
		return "Lukovica"
	case SubdivisionSI069:
		return "Majsperk"
	case SubdivisionSI070:
		return "Maribor"
	case SubdivisionSI071:
		return "Medvode"
	case SubdivisionSI072:
		return "Menges"
	case SubdivisionSI073:
		return "Metlika"
	case SubdivisionSI074:
		return "Mezica"
	case SubdivisionSI075:
		return "Miren-Kostanjevica"
	case SubdivisionSI076:
		return "Mislinja"
	case SubdivisionSI077:
		return "Moravce"
	case SubdivisionSI079:
		return "Mozirje"
	case SubdivisionSI080:
		return "Murska Sobota"
	case SubdivisionSI081:
		return "Muta"
	case SubdivisionSI082:
		return "Naklo"
	case SubdivisionSI083:
		return "Nazarje"
	case SubdivisionSI084:
		return "Nova Gorica"
	case SubdivisionSI085:
		return "Novo Mesto"
	case SubdivisionSI086:
		return "Odranci"
	case SubdivisionSI087:
		return "Ormoz"
	case SubdivisionSI090:
		return "Piran"
	case SubdivisionSI091:
		return "Pivka"
	case SubdivisionSI092:
		return "Podcetrtek"
	case SubdivisionSI094:
		return "Postojna"
	case SubdivisionSI095:
		return "Preddvor"
	case SubdivisionSI096:
		return "Ptuj"
	case SubdivisionSI097:
		return "Puconci"
	case SubdivisionSI098:
		return "Race-Fram"
	case SubdivisionSI099:
		return "Radece"
	case SubdivisionSI100:
		return "Radenci"
	case SubdivisionSI101:
		return "Radlje ob Dravi"
	case SubdivisionSI102:
		return "Radovljica"
	case SubdivisionSI103:
		return "Ravne na Koroskem"
	case SubdivisionSI104:
		return "Ribnica"
	case SubdivisionSI105:
		return "Rogasovci"
	case SubdivisionSI106:
		return "Rogaska Slatina"
	case SubdivisionSI108:
		return "Ruse"
	case SubdivisionSI109:
		return "Semic"
	case SubdivisionSI110:
		return "Sevnica"
	case SubdivisionSI111:
		return "Sezana"
	case SubdivisionSI112:
		return "Slovenj Gradec"
	case SubdivisionSI113:
		return "Slovenska Bistrica"
	case SubdivisionSI114:
		return "Slovenske Konjice"
	case SubdivisionSI115:
		return "Starse"
	case SubdivisionSI116:
		return "Sveti Jurij ob Scavnici"
	case SubdivisionSI117:
		return "Sencur"
	case SubdivisionSI118:
		return "Sentilj"
	case SubdivisionSI119:
		return "Sentjernej"
	case SubdivisionSI120:
		return "Sentjur"
	case SubdivisionSI121:
		return "Skocjan"
	case SubdivisionSI122:
		return "Skofja Loka"
	case SubdivisionSI123:
		return "Skofljica"
	case SubdivisionSI124:
		return "Smarje pri Jelsah"
	case SubdivisionSI125:
		return "Smartno ob Paki"
	case SubdivisionSI126:
		return "Sostanj"
	case SubdivisionSI127:
		return "Store"
	case SubdivisionSI128:
		return "Tolmin"
	case SubdivisionSI129:
		return "Trbovlje"
	case SubdivisionSI130:
		return "Trebnje"
	case SubdivisionSI131:
		return "Trzic"
	case SubdivisionSI132:
		return "Turnisce"
	case SubdivisionSI133:
		return "Velenje"
	case SubdivisionSI134:
		return "Velike Lasce"
	case SubdivisionSI135:
		return "Videm"
	case SubdivisionSI136:
		return "Vipava"
	case SubdivisionSI137:
		return "Vitanje"
	case SubdivisionSI138:
		return "Vodice"
	case SubdivisionSI139:
		return "Vojnik"
	case SubdivisionSI140:
		return "Vrhnika"
	case SubdivisionSI141:
		return "Vuzenica"
	case SubdivisionSI142:
		return "Zagorje ob Savi"
	case SubdivisionSI143:
		return "Zavrc"
	case SubdivisionSI144:
		return "Zrece"
	case SubdivisionSI146:
		return "Zelezniki"
	case SubdivisionSI147:
		return "Ziri"
	case SubdivisionSI148:
		return "Benedikt"
	case SubdivisionSI149:
		return "Bistrica ob Sotli"
	case SubdivisionSI150:
		return "Bloke"
	case SubdivisionSI151:
		return "Braslovce"
	case SubdivisionSI152:
		return "Cankova"
	case SubdivisionSI154:
		return "Dobje"
	case SubdivisionSI155:
		return "Dobrna"
	case SubdivisionSI156:
		return "Dobrovnik"
	case SubdivisionSI158:
		return "Grad"
	case SubdivisionSI159:
		return "Hajdina"
	case SubdivisionSI160:
		return "Hoce-Slivnica"
	case SubdivisionSI161:
		return "Hodos"
	case SubdivisionSI162:
		return "Horjul"
	case SubdivisionSI164:
		return "Komenda"
	case SubdivisionSI165:
		return "Kostel"
	case SubdivisionSI166:
		return "Krizevci"
	case SubdivisionSI167:
		return "Lovrenc na Pohorju"
	case SubdivisionSI168:
		return "Markovci"
	case SubdivisionSI169:
		return "Miklavz na Dravskem polju"
	case SubdivisionSI170:
		return "Mirna Pec"
	case SubdivisionSI171:
		return "Oplotnica"
	case SubdivisionSI172:
		return "Podlehnik"
	case SubdivisionSI173:
		return "Polzela"
	case SubdivisionSI174:
		return "Prebold"
	case SubdivisionSI175:
		return "Prevalje"
	case SubdivisionSI176:
		return "Razkrizje"
	case SubdivisionSI179:
		return "Sodrazica"
	case SubdivisionSI180:
		return "Solcava"
	case SubdivisionSI182:
		return "Sveti Andraz v Slovenskih Goricah"
	case SubdivisionSI183:
		return "Sempeter-Vrtojba"
	case SubdivisionSI184:
		return "Tabor"
	case SubdivisionSI185:
		return "Trnovska Vas"
	case SubdivisionSI186:
		return "Trzin"
	case SubdivisionSI187:
		return "Velika Polana"
	case SubdivisionSI188:
		return "Verzej"
	case SubdivisionSI189:
		return "Vransko"
	case SubdivisionSI190:
		return "Zalec"
	case SubdivisionSI191:
		return "Zetale"
	case SubdivisionSI193:
		return "Zuzemberk"
	case SubdivisionSI194:
		return "Smartno pri Litiji"
	case SubdivisionSI195:
		return "Apace"
	case SubdivisionSI196:
		return "Cirkulane"
	case SubdivisionSI197:
		return "Kosanjevica na Krki"
	case SubdivisionSI198:
		return "Makole"
	case SubdivisionSI199:
		return "Mokronog-Trebelno"
	case SubdivisionSI200:
		return "Poljcane"
	case SubdivisionSI201:
		return "Rence-Vogrsko"
	case SubdivisionSI203:
		return "Straza"
	case SubdivisionSI204:
		return "Sveta Trojica v Slovenskih goricah"
	case SubdivisionSI205:
		return "Sveti Tomaz"
	case SubdivisionSI206:
		return "Smarjeske Toplice"
	case SubdivisionSI207:
		return "Gorje"
	case SubdivisionSI208:
		return "Log-Dragomer"
	case SubdivisionSI209:
		return "Recica ob Savinji"
	case SubdivisionSI210:
		return "Sveti Jurij v Slovenskih goricah"
	case SubdivisionSI211:
		return "Sentrupert"
	case SubdivisionSI212:
		return "Mirna"
	case SubdivisionSI213:
		return "Ankaran"
	case SubdivisionSKBC:
		return "Banskobystricky kraj"
	case SubdivisionSKBL:
		return "Bratislavsky kraj"
	case SubdivisionSKKI:
		return "Kosicky kraj"
	case SubdivisionSKNI:
		return "Nitriansky kraj"
	case SubdivisionSKPV:
		return "Presovsky kraj"
	case SubdivisionSKTA:
		return "Trnavsky kraj"
	case SubdivisionSKTC:
		return "Trenciansky kraj"
	case SubdivisionSKZI:
		return "Zilinsky kraj"
	case SubdivisionSLE:
		return "Eastern"
	case SubdivisionSLN:
		return "Northern"
	case SubdivisionSLS:
		return "Southern"
	case SubdivisionSLW:
		return "Western Area"
	case SubdivisionSM03:
		return "Domagnano"
	case SubdivisionSM07:
		return "Citta di San Marino"
	case SubdivisionSM09:
		return "Serravalle"
	case SubdivisionSNDB:
		return "Diourbel"
	case SubdivisionSNDK:
		return "Dakar"
	case SubdivisionSNFK:
		return "Fatick"
	case SubdivisionSNKA:
		return "Kaffrine"
	case SubdivisionSNKD:
		return "Kolda"
	case SubdivisionSNKE:
		return "Kedougou"
	case SubdivisionSNKL:
		return "Kaolack"
	case SubdivisionSNLG:
		return "Louga"
	case SubdivisionSNMT:
		return "Matam"
	case SubdivisionSNSL:
		return "Saint-Louis"
	case SubdivisionSNTC:
		return "Tambacounda"
	case SubdivisionSNTH:
		return "Thies"
	case SubdivisionSNZG:
		return "Ziguinchor"
	case SubdivisionSOAW:
		return "Awdal"
	case SubdivisionSOBN:
		return "Banaadir"
	case SubdivisionSOMU:
		return "Mudug"
	case SubdivisionSONU:
		return "Nugaal"
	case SubdivisionSOSH:
		return "Shabeellaha Hoose"
	case SubdivisionSOSO:
		return "Sool"
	case SubdivisionSOTO:
		return "Togdheer"
	case SubdivisionSOWO:
		return "Woqooyi Galbeed"
	case SubdivisionSRCM:
		return "Commewijne"
	case SubdivisionSRNI:
		return "Nickerie"
	case SubdivisionSRPM:
		return "Paramaribo"
	case SubdivisionSRPR:
		return "Para"
	case SubdivisionSRSI:
		return "Sipaliwini"
	case SubdivisionSRWA:
		return "Wanica"
	case SubdivisionSSBN:
		return "Northern Bahr el Ghazal"
	case SubdivisionSSEC:
		return "Central Equatoria"
	case SubdivisionSSEE:
		return "Eastern Equatoria"
	case SubdivisionSSEW:
		return "Western Equatoria"
	case SubdivisionSSNU:
		return "Upper Nile"
	case SubdivisionST01:
		return "Agua Grande"
	case SubdivisionSVAH:
		return "Ahuachapan"
	case SubdivisionSVCA:
		return "Cabanas"
	case SubdivisionSVCH:
		return "Chalatenango"
	case SubdivisionSVCU:
		return "Cuscatlan"
	case SubdivisionSVLI:
		return "La Libertad"
	case SubdivisionSVMO:
		return "Morazan"
	case SubdivisionSVPA:
		return "La Paz"
	case SubdivisionSVSA:
		return "Santa Ana"
	case SubdivisionSVSM:
		return "San Miguel"
	case SubdivisionSVSO:
		return "Sonsonate"
	case SubdivisionSVSS:
		return "San Salvador"
	case SubdivisionSVSV:
		return "San Vicente"
	case SubdivisionSVUN:
		return "La Union"
	case SubdivisionSVUS:
		return "Usulutan"
	case SubdivisionSYDI:
		return "Dimashq"
	case SubdivisionSYDR:
		return "Dar'a"
	case SubdivisionSYDY:
		return "Dayr az Zawr"
	case SubdivisionSYHA:
		return "Al Hasakah"
	case SubdivisionSYHI:
		return "Hims"
	case SubdivisionSYHL:
		return "Halab"
	case SubdivisionSYHM:
		return "Hamah"
	case SubdivisionSYID:
		return "Idlib"
	case SubdivisionSYLA:
		return "Al Ladhiqiyah"
	case SubdivisionSYQU:
		return "Al Qunaytirah"
	case SubdivisionSYRA:
		return "Ar Raqqah"
	case SubdivisionSYRD:
		return "Rif Dimashq"
	case SubdivisionSYSU:
		return "As Suwayda'"
	case SubdivisionSYTA:
		return "Tartus"
	case SubdivisionSZHH:
		return "Hhohho"
	case SubdivisionSZLU:
		return "Lubombo"
	case SubdivisionSZMA:
		return "Manzini"
	case SubdivisionTDGR:
		return "Guera"
	case SubdivisionTDLO:
		return "Logone-Occidental"
	case SubdivisionTDME:
		return "Mayo-Kebbi-Est"
	case SubdivisionTDND:
		return "Ville de Ndjamena"
	case SubdivisionTDOD:
		return "Ouaddai"
	case SubdivisionTGC:
		return "Centrale"
	case SubdivisionTGK:
		return "Kara"
	case SubdivisionTGM:
		return "Maritime"
	case SubdivisionTGP:
		return "Plateaux"
	case SubdivisionTH10:
		return "Krung Thep Maha Nakhon"
	case SubdivisionTH11:
		return "Samut Prakan"
	case SubdivisionTH12:
		return "Nonthaburi"
	case SubdivisionTH13:
		return "Pathum Thani"
	case SubdivisionTH14:
		return "Phra Nakhon Si Ayutthaya"
	case SubdivisionTH15:
		return "Ang Thong"
	case SubdivisionTH16:
		return "Lop Buri"
	case SubdivisionTH17:
		return "Sing Buri"
	case SubdivisionTH18:
		return "Chai Nat"
	case SubdivisionTH19:
		return "Saraburi"
	case SubdivisionTH20:
		return "Chon Buri"
	case SubdivisionTH21:
		return "Rayong"
	case SubdivisionTH22:
		return "Chanthaburi"
	case SubdivisionTH23:
		return "Trat"
	case SubdivisionTH24:
		return "Chachoengsao"
	case SubdivisionTH25:
		return "Prachin Buri"
	case SubdivisionTH26:
		return "Nakhon Nayok"
	case SubdivisionTH27:
		return "Sa Kaeo"
	case SubdivisionTH30:
		return "Nakhon Ratchasima"
	case SubdivisionTH31:
		return "Buri Ram"
	case SubdivisionTH32:
		return "Surin"
	case SubdivisionTH33:
		return "Si Sa Ket"
	case SubdivisionTH34:
		return "Ubon Ratchathani"
	case SubdivisionTH35:
		return "Yasothon"
	case SubdivisionTH36:
		return "Chaiyaphum"
	case SubdivisionTH37:
		return "Amnat Charoen"
	case SubdivisionTH38:
		return "Bueng Kan"
	case SubdivisionTH39:
		return "Nong Bua Lam Phu"
	case SubdivisionTH40:
		return "Khon Kaen"
	case SubdivisionTH41:
		return "Udon Thani"
	case SubdivisionTH42:
		return "Loei"
	case SubdivisionTH43:
		return "Nong Khai"
	case SubdivisionTH44:
		return "Maha Sarakham"
	case SubdivisionTH45:
		return "Roi Et"
	case SubdivisionTH46:
		return "Kalasin"
	case SubdivisionTH47:
		return "Sakon Nakhon"
	case SubdivisionTH48:
		return "Nakhon Phanom"
	case SubdivisionTH49:
		return "Mukdahan"
	case SubdivisionTH50:
		return "Chiang Mai"
	case SubdivisionTH51:
		return "Lamphun"
	case SubdivisionTH52:
		return "Lampang"
	case SubdivisionTH53:
		return "Uttaradit"
	case SubdivisionTH54:
		return "Phrae"
	case SubdivisionTH55:
		return "Nan"
	case SubdivisionTH56:
		return "Phayao"
	case SubdivisionTH57:
		return "Chiang Rai"
	case SubdivisionTH58:
		return "Mae Hong Son"
	case SubdivisionTH60:
		return "Nakhon Sawan"
	case SubdivisionTH61:
		return "Uthai Thani"
	case SubdivisionTH62:
		return "Kamphaeng Phet"
	case SubdivisionTH63:
		return "Tak"
	case SubdivisionTH64:
		return "Sukhothai"
	case SubdivisionTH65:
		return "Phitsanulok"
	case SubdivisionTH66:
		return "Phichit"
	case SubdivisionTH67:
		return "Phetchabun"
	case SubdivisionTH70:
		return "Ratchaburi"
	case SubdivisionTH71:
		return "Kanchanaburi"
	case SubdivisionTH72:
		return "Suphan Buri"
	case SubdivisionTH73:
		return "Nakhon Pathom"
	case SubdivisionTH74:
		return "Samut Sakhon"
	case SubdivisionTH75:
		return "Samut Songkhram"
	case SubdivisionTH76:
		return "Phetchaburi"
	case SubdivisionTH77:
		return "Prachuap Khiri Khan"
	case SubdivisionTH80:
		return "Nakhon Si Thammarat"
	case SubdivisionTH81:
		return "Krabi"
	case SubdivisionTH82:
		return "Phangnga"
	case SubdivisionTH83:
		return "Phuket"
	case SubdivisionTH84:
		return "Surat Thani"
	case SubdivisionTH85:
		return "Ranong"
	case SubdivisionTH86:
		return "Chumphon"
	case SubdivisionTH90:
		return "Songkhla"
	case SubdivisionTH91:
		return "Satun"
	case SubdivisionTH92:
		return "Trang"
	case SubdivisionTH93:
		return "Phatthalung"
	case SubdivisionTH94:
		return "Pattani"
	case SubdivisionTH95:
		return "Yala"
	case SubdivisionTH96:
		return "Narathiwat"
	case SubdivisionTJDU:
		return "Dushanbe"
	case SubdivisionTJKT:
		return "Khatlon"
	case SubdivisionTJRA:
		return "Nohiyahoi Tobei Jumhuri"
	case SubdivisionTJSU:
		return "Sughd"
	case SubdivisionTLAN:
		return "Ainaro"
	case SubdivisionTLCO:
		return "Cova Lima"
	case SubdivisionTLDI:
		return "Dili"
	case SubdivisionTMA:
		return "Ahal"
	case SubdivisionTMB:
		return "Balkan"
	case SubdivisionTMD:
		return "Dasoguz"
	case SubdivisionTML:
		return "Lebap"
	case SubdivisionTMM:
		return "Mary"
	case SubdivisionTN11:
		return "Tunis"
	case SubdivisionTN12:
		return "L'Ariana"
	case SubdivisionTN13:
		return "Ben Arous"
	case SubdivisionTN14:
		return "La Manouba"
	case SubdivisionTN21:
		return "Nabeul"
	case SubdivisionTN22:
		return "Zaghouan"
	case SubdivisionTN23:
		return "Bizerte"
	case SubdivisionTN31:
		return "Beja"
	case SubdivisionTN32:
		return "Jendouba"
	case SubdivisionTN33:
		return "Le Kef"
	case SubdivisionTN34:
		return "Siliana"
	case SubdivisionTN41:
		return "Kairouan"
	case SubdivisionTN42:
		return "Kasserine"
	case SubdivisionTN43:
		return "Sidi Bouzid"
	case SubdivisionTN51:
		return "Sousse"
	case SubdivisionTN52:
		return "Monastir"
	case SubdivisionTN53:
		return "Mahdia"
	case SubdivisionTN61:
		return "Sfax"
	case SubdivisionTN71:
		return "Gafsa"
	case SubdivisionTN72:
		return "Tozeur"
	case SubdivisionTN73:
		return "Kebili"
	case SubdivisionTN81:
		return "Gabes"
	case SubdivisionTN82:
		return "Medenine"
	case SubdivisionTN83:
		return "Tataouine"
	case SubdivisionTO03:
		return "Niuas"
	case SubdivisionTO04:
		return "Tongatapu"
	case SubdivisionTR01:
		return "Adana"
	case SubdivisionTR02:
		return "Adiyaman"
	case SubdivisionTR03:
		return "Afyonkarahisar"
	case SubdivisionTR04:
		return "Agri"
	case SubdivisionTR05:
		return "Amasya"
	case SubdivisionTR06:
		return "Ankara"
	case SubdivisionTR07:
		return "Antalya"
	case SubdivisionTR08:
		return "Artvin"
	case SubdivisionTR09:
		return "Aydin"
	case SubdivisionTR10:
		return "Balikesir"
	case SubdivisionTR11:
		return "Bilecik"
	case SubdivisionTR12:
		return "Bingol"
	case SubdivisionTR13:
		return "Bitlis"
	case SubdivisionTR14:
		return "Bolu"
	case SubdivisionTR15:
		return "Burdur"
	case SubdivisionTR16:
		return "Bursa"
	case SubdivisionTR17:
		return "Canakkale"
	case SubdivisionTR18:
		return "Cankiri"
	case SubdivisionTR19:
		return "Corum"
	case SubdivisionTR20:
		return "Denizli"
	case SubdivisionTR21:
		return "Diyarbakir"
	case SubdivisionTR22:
		return "Edirne"
	case SubdivisionTR23:
		return "Elazig"
	case SubdivisionTR24:
		return "Erzincan"
	case SubdivisionTR25:
		return "Erzurum"
	case SubdivisionTR26:
		return "Eskisehir"
	case SubdivisionTR27:
		return "Gaziantep"
	case SubdivisionTR28:
		return "Giresun"
	case SubdivisionTR29:
		return "Gumushane"
	case SubdivisionTR30:
		return "Hakkari"
	case SubdivisionTR31:
		return "Hatay"
	case SubdivisionTR32:
		return "Isparta"
	case SubdivisionTR33:
		return "Mersin"
	case SubdivisionTR34:
		return "Istanbul"
	case SubdivisionTR35:
		return "Izmir"
	case SubdivisionTR36:
		return "Kars"
	case SubdivisionTR37:
		return "Kastamonu"
	case SubdivisionTR38:
		return "Kayseri"
	case SubdivisionTR39:
		return "Kirklareli"
	case SubdivisionTR40:
		return "Kirsehir"
	case SubdivisionTR41:
		return "Kocaeli"
	case SubdivisionTR42:
		return "Konya"
	case SubdivisionTR43:
		return "Kutahya"
	case SubdivisionTR44:
		return "Malatya"
	case SubdivisionTR45:
		return "Manisa"
	case SubdivisionTR46:
		return "Kahramanmaras"
	case SubdivisionTR47:
		return "Mardin"
	case SubdivisionTR48:
		return "Mugla"
	case SubdivisionTR49:
		return "Mus"
	case SubdivisionTR50:
		return "Nevsehir"
	case SubdivisionTR51:
		return "Nigde"
	case SubdivisionTR52:
		return "Ordu"
	case SubdivisionTR53:
		return "Rize"
	case SubdivisionTR54:
		return "Sakarya"
	case SubdivisionTR55:
		return "Samsun"
	case SubdivisionTR56:
		return "Siirt"
	case SubdivisionTR57:
		return "Sinop"
	case SubdivisionTR58:
		return "Sivas"
	case SubdivisionTR59:
		return "Tekirdag"
	case SubdivisionTR60:
		return "Tokat"
	case SubdivisionTR61:
		return "Trabzon"
	case SubdivisionTR62:
		return "Tunceli"
	case SubdivisionTR63:
		return "Sanliurfa"
	case SubdivisionTR64:
		return "Usak"
	case SubdivisionTR65:
		return "Van"
	case SubdivisionTR66:
		return "Yozgat"
	case SubdivisionTR67:
		return "Zonguldak"
	case SubdivisionTR68:
		return "Aksaray"
	case SubdivisionTR69:
		return "Bayburt"
	case SubdivisionTR70:
		return "Karaman"
	case SubdivisionTR71:
		return "Kirikkale"
	case SubdivisionTR72:
		return "Batman"
	case SubdivisionTR73:
		return "Sirnak"
	case SubdivisionTR74:
		return "Bartin"
	case SubdivisionTR75:
		return "Ardahan"
	case SubdivisionTR76:
		return "Igdir"
	case SubdivisionTR77:
		return "Yalova"
	case SubdivisionTR78:
		return "Karabuk"
	case SubdivisionTR79:
		return "Kilis"
	case SubdivisionTR80:
		return "Osmaniye"
	case SubdivisionTR81:
		return "Duzce"
	case SubdivisionTTARI:
		return "Arima"
	case SubdivisionTTCHA:
		return "Chaguanas"
	case SubdivisionTTCTT:
		return "Couva-Tabaquite-Talparo"
	case SubdivisionTTDMN:
		return "Diego Martin"
	case SubdivisionTTMRC:
		return "Mayaro-Rio Claro"
	case SubdivisionTTPED:
		return "Penal-Debe"
	case SubdivisionTTPOS:
		return "Port of Spain"
	case SubdivisionTTPRT:
		return "Princes Town"
	case SubdivisionTTPTF:
		return "Point Fortin"
	case SubdivisionTTSFO:
		return "San Fernando"
	case SubdivisionTTSGE:
		return "Sangre Grande"
	case SubdivisionTTSIP:
		return "Siparia"
	case SubdivisionTTSJL:
		return "San Juan-Laventille"
	case SubdivisionTTTOB:
		return "Tobago"
	case SubdivisionTTTUP:
		return "Tunapuna-Piarco"
	case SubdivisionTVFUN:
		return "Funafuti"
	case SubdivisionTWCHA:
		return "Changhua"
	case SubdivisionTWCYQ:
		return "Chiayi"
	case SubdivisionTWHSQ:
		return "Hsinchu"
	case SubdivisionTWHUA:
		return "Hualien"
	case SubdivisionTWILA:
		return "Yilan"
	case SubdivisionTWKEE:
		return "Keelung"
	case SubdivisionTWKHH:
		return "Kaohsiung"
	case SubdivisionTWKIN:
		return "Kinmen"
	case SubdivisionTWLIE:
		return "Lienchiang"
	case SubdivisionTWMIA:
		return "Miaoli"
	case SubdivisionTWNAN:
		return "Nantou"
	case SubdivisionTWNWT:
		return "New Taipei"
	case SubdivisionTWPEN:
		return "Penghu"
	case SubdivisionTWPIF:
		return "Pingtung"
	case SubdivisionTWTAO:
		return "Taoyuan"
	case SubdivisionTWTNN:
		return "Tainan"
	case SubdivisionTWTPE:
		return "Taipei"
	case SubdivisionTWTTT:
		return "Taitung"
	case SubdivisionTWTXG:
		return "Taichung"
	case SubdivisionTWYUN:
		return "Yunlin"
	case SubdivisionTZ01:
		return "Arusha"
	case SubdivisionTZ02:
		return "Dar es Salaam"
	case SubdivisionTZ03:
		return "Dodoma"
	case SubdivisionTZ04:
		return "Iringa"
	case SubdivisionTZ05:
		return "Kagera"
	case SubdivisionTZ07:
		return "Kaskazini Unguja"
	case SubdivisionTZ08:
		return "Kigoma"
	case SubdivisionTZ09:
		return "Kilimanjaro"
	case SubdivisionTZ11:
		return "Kusini Unguja"
	case SubdivisionTZ12:
		return "Lindi"
	case SubdivisionTZ13:
		return "Mara"
	case SubdivisionTZ14:
		return "Mbeya"
	case SubdivisionTZ15:
		return "Mjini Magharibi"
	case SubdivisionTZ16:
		return "Morogoro"
	case SubdivisionTZ17:
		return "Mtwara"
	case SubdivisionTZ18:
		return "Mwanza"
	case SubdivisionTZ19:
		return "Pwani"
	case SubdivisionTZ20:
		return "Rukwa"
	case SubdivisionTZ21:
		return "Ruvuma"
	case SubdivisionTZ22:
		return "Shinyanga"
	case SubdivisionTZ23:
		return "Singida"
	case SubdivisionTZ24:
		return "Tabora"
	case SubdivisionTZ25:
		return "Tanga"
	case SubdivisionTZ26:
		return "Manyara"
	case SubdivisionTZ27:
		return "Geita"
	case SubdivisionTZ28:
		return "Katavi"
	case SubdivisionTZ29:
		return "Njombe"
	case SubdivisionTZ30:
		return "Simiyu"
	case SubdivisionTZ31:
		return "Songwe"
	case SubdivisionUA05:
		return "Vinnytska oblast"
	case SubdivisionUA07:
		return "Volynska oblast"
	case SubdivisionUA09:
		return "Luhanska oblast"
	case SubdivisionUA12:
		return "Dnipropetrovska oblast"
	case SubdivisionUA14:
		return "Donetska oblast"
	case SubdivisionUA18:
		return "Zhytomyrska oblast"
	case SubdivisionUA21:
		return "Zakarpatska oblast"
	case SubdivisionUA23:
		return "Zaporizka oblast"
	case SubdivisionUA26:
		return "Ivano-Frankivska oblast"
	case SubdivisionUA30:
		return "Kyiv"
	case SubdivisionUA32:
		return "Kyivska oblast"
	case SubdivisionUA35:
		return "Kirovohradska oblast"
	case SubdivisionUA40:
		return "Sevastopol"
	case SubdivisionUA43:
		return "Avtonomna Respublika Krym"
	case SubdivisionUA46:
		return "Lvivska oblast"
	case SubdivisionUA48:
		return "Mykolaivska oblast"
	case SubdivisionUA51:
		return "Odeska oblast"
	case SubdivisionUA53:
		return "Poltavska oblast"
	case SubdivisionUA56:
		return "Rivnenska oblast"
	case SubdivisionUA59:
		return "Sumska oblast"
	case SubdivisionUA61:
		return "Ternopilska oblast"
	case SubdivisionUA63:
		return "Kharkivska oblast"
	case SubdivisionUA65:
		return "Khersonska oblast"
	case SubdivisionUA68:
		return "Khmelnytska oblast"
	case SubdivisionUA71:
		return "Cherkaska oblast"
	case SubdivisionUA74:
		return "Chernihivska oblast"
	case SubdivisionUA77:
		return "Chernivetska oblast"
	case SubdivisionUG101:
		return "Kalangala"
	case SubdivisionUG102:
		return "Kampala"
	case SubdivisionUG103:
		return "Kiboga"
	case SubdivisionUG104:
		return "Luwero"
	case SubdivisionUG105:
		return "Masaka"
	case SubdivisionUG106:
		return "Mpigi"
	case SubdivisionUG107:
		return "Mubende"
	case SubdivisionUG108:
		return "Mukono"
	case SubdivisionUG109:
		return "Nakasongola"
	case SubdivisionUG112:
		return "Kayunga"
	case SubdivisionUG113:
		return "Wakiso"
	case SubdivisionUG115:
		return "Mityana"
	case SubdivisionUG116:
		return "Nakaseke"
	case SubdivisionUG117:
		return "Buikwe"
	case SubdivisionUG120:
		return "Buvuma"
	case SubdivisionUG122:
		return "Kalungu"
	case SubdivisionUG201:
		return "Bugiri"
	case SubdivisionUG203:
		return "Iganga"
	case SubdivisionUG204:
		return "Jinja"
	case SubdivisionUG205:
		return "Kamuli"
	case SubdivisionUG206:
		return "Kapchorwa"
	case SubdivisionUG209:
		return "Mbale"
	case SubdivisionUG210:
		return "Pallisa"
	case SubdivisionUG211:
		return "Soroti"
	case SubdivisionUG214:
		return "Mayuge"
	case SubdivisionUG215:
		return "Sironko"
	case SubdivisionUG219:
		return "Bukedea"
	case SubdivisionUG222:
		return "Kaliro"
	case SubdivisionUG303:
		return "Arua"
	case SubdivisionUG304:
		return "Gulu"
	case SubdivisionUG305:
		return "Kitgum"
	case SubdivisionUG307:
		return "Lira"
	case SubdivisionUG316:
		return "Amuru"
	case SubdivisionUG321:
		return "Oyam"
	case SubdivisionUG403:
		return "Hoima"
	case SubdivisionUG404:
		return "Kabale"
	case SubdivisionUG405:
		return "Kabarole"
	case SubdivisionUG406:
		return "Kasese"
	case SubdivisionUG407:
		return "Kibaale"
	case SubdivisionUG408:
		return "Kisoro"
	case SubdivisionUG409:
		return "Masindi"
	case SubdivisionUG410:
		return "Mbarara"
	case SubdivisionUG411:
		return "Ntungamo"
	case SubdivisionUG412:
		return "Rukungiri"
	case SubdivisionUG413:
		return "Kamwenge"
	case SubdivisionUG415:
		return "Kyenjojo"
	case SubdivisionUG419:
		return "Kiruhura"
	case SubdivisionUM95:
		return "Palmyra Atoll"
	case SubdivisionUSAK:
		return "Alaska"
	case SubdivisionUSAL:
		return "Alabama"
	case SubdivisionUSAR:
		return "Arkansas"
	case SubdivisionUSAZ:
		return "Arizona"
	case SubdivisionUSCA:
		return "California"
	case SubdivisionUSCO:
		return "Colorado"
	case SubdivisionUSCT:
		return "Connecticut"
	case SubdivisionUSDC:
		return "District of Columbia"
	case SubdivisionUSDE:
		return "Delaware"
	case SubdivisionUSFL:
		return "Florida"
	case SubdivisionUSGA:
		return "Georgia"
	case SubdivisionUSHI:
		return "Hawaii"
	case SubdivisionUSIA:
		return "Iowa"
	case SubdivisionUSID:
		return "Idaho"
	case SubdivisionUSIL:
		return "Illinois"
	case SubdivisionUSIN:
		return "Indiana"
	case SubdivisionUSKS:
		return "Kansas"
	case SubdivisionUSKY:
		return "Kentucky"
	case SubdivisionUSLA:
		return "Louisiana"
	case SubdivisionUSMA:
		return "Massachusetts"
	case SubdivisionUSMD:
		return "Maryland"
	case SubdivisionUSME:
		return "Maine"
	case SubdivisionUSMI:
		return "Michigan"
	case SubdivisionUSMN:
		return "Minnesota"
	case SubdivisionUSMO:
		return "Missouri"
	case SubdivisionUSMS:
		return "Mississippi"
	case SubdivisionUSMT:
		return "Montana"
	case SubdivisionUSNC:
		return "North Carolina"
	case SubdivisionUSND:
		return "North Dakota"
	case SubdivisionUSNE:
		return "Nebraska"
	case SubdivisionUSNH:
		return "New Hampshire"
	case SubdivisionUSNJ:
		return "New Jersey"
	case SubdivisionUSNM:
		return "New Mexico"
	case SubdivisionUSNV:
		return "Nevada"
	case SubdivisionUSNY:
		return "New York"
	case SubdivisionUSOH:
		return "Ohio"
	case SubdivisionUSOK:
		return "Oklahoma"
	case SubdivisionUSOR:
		return "Oregon"
	case SubdivisionUSPA:
		return "Pennsylvania"
	case SubdivisionUSRI:
		return "Rhode Island"
	case SubdivisionUSSC:
		return "South Carolina"
	case SubdivisionUSSD:
		return "South Dakota"
	case SubdivisionUSTN:
		return "Tennessee"
	case SubdivisionUSTX:
		return "Texas"
	case SubdivisionUSUT:
		return "Utah"
	case SubdivisionUSVA:
		return "Virginia"
	case SubdivisionUSVT:
		return "Vermont"
	case SubdivisionUSWA:
		return "Washington"
	case SubdivisionUSWI:
		return "Wisconsin"
	case SubdivisionUSWV:
		return "West Virginia"
	case SubdivisionUSWY:
		return "Wyoming"
	case SubdivisionUYAR:
		return "Artigas"
	case SubdivisionUYCA:
		return "Canelones"
	case SubdivisionUYCL:
		return "Cerro Largo"
	case SubdivisionUYCO:
		return "Colonia"
	case SubdivisionUYDU:
		return "Durazno"
	case SubdivisionUYFD:
		return "Florida"
	case SubdivisionUYFS:
		return "Flores"
	case SubdivisionUYLA:
		return "Lavalleja"
	case SubdivisionUYMA:
		return "Maldonado"
	case SubdivisionUYMO:
		return "Montevideo"
	case SubdivisionUYPA:
		return "Paysandu"
	case SubdivisionUYRN:
		return "Rio Negro"
	case SubdivisionUYRO:
		return "Rocha"
	case SubdivisionUYRV:
		return "Rivera"
	case SubdivisionUYSA:
		return "Salto"
	case SubdivisionUYSJ:
		return "San Jose"
	case SubdivisionUYSO:
		return "Soriano"
	case SubdivisionUYTA:
		return "Tacuarembo"
	case SubdivisionUYTT:
		return "Treinta y Tres"
	case SubdivisionUZAN:
		return "Andijon"
	case SubdivisionUZBU:
		return "Buxoro"
	case SubdivisionUZFA:
		return "Farg'ona"
	case SubdivisionUZJI:
		return "Jizzax"
	case SubdivisionUZNG:
		return "Namangan"
	case SubdivisionUZNW:
		return "Navoiy"
	case SubdivisionUZQA:
		return "Qashqadaryo"
	case SubdivisionUZQR:
		return "Qoraqalpog'iston Respublikasi"
	case SubdivisionUZSA:
		return "Samarqand"
	case SubdivisionUZSI:
		return "Sirdaryo"
	case SubdivisionUZSU:
		return "Surxondaryo"
	case SubdivisionUZTK:
		return "Toshkent"
	case SubdivisionUZXO:
		return "Xorazm"
	case SubdivisionVC01:
		return "Charlotte"
	case SubdivisionVC04:
		return "Saint George"
	case SubdivisionVC05:
		return "Saint Patrick"
	case SubdivisionVC06:
		return "Grenadines"
	case SubdivisionVEA:
		return "Distrito Capital"
	case SubdivisionVEB:
		return "Anzoategui"
	case SubdivisionVEC:
		return "Apure"
	case SubdivisionVED:
		return "Aragua"
	case SubdivisionVEE:
		return "Barinas"
	case SubdivisionVEF:
		return "Bolivar"
	case SubdivisionVEG:
		return "Carabobo"
	case SubdivisionVEH:
		return "Cojedes"
	case SubdivisionVEI:
		return "Falcon"
	case SubdivisionVEJ:
		return "Guarico"
	case SubdivisionVEK:
		return "Lara"
	case SubdivisionVEL:
		return "Merida"
	case SubdivisionVEM:
		return "Miranda"
	case SubdivisionVEN:
		return "Monagas"
	case SubdivisionVEO:
		return "Nueva Esparta"
	case SubdivisionVEP:
		return "Portuguesa"
	case SubdivisionVER:
		return "Sucre"
	case SubdivisionVES:
		return "Tachira"
	case SubdivisionVET:
		return "Trujillo"
	case SubdivisionVEU:
		return "Yaracuy"
	case SubdivisionVEV:
		return "Zulia"
	case SubdivisionVEX:
		return "La Guaira"
	case SubdivisionVEY:
		return "Delta Amacuro"
	case SubdivisionVEZ:
		return "Amazonas"
	case SubdivisionVN01:
		return "Lai Chau"
	case SubdivisionVN02:
		return "Lao Cai"
	case SubdivisionVN03:
		return "Ha Giang"
	case SubdivisionVN04:
		return "Cao Bang"
	case SubdivisionVN05:
		return "Son La"
	case SubdivisionVN06:
		return "Yen Bai"
	case SubdivisionVN07:
		return "Tuyen Quang"
	case SubdivisionVN09:
		return "Lang Son"
	case SubdivisionVN13:
		return "Quang Ninh"
	case SubdivisionVN14:
		return "Hoa Binh"
	case SubdivisionVN18:
		return "Ninh Binh"
	case SubdivisionVN20:
		return "Thai Binh"
	case SubdivisionVN21:
		return "Thanh Hoa"
	case SubdivisionVN22:
		return "Nghe An"
	case SubdivisionVN23:
		return "Ha Tinh"
	case SubdivisionVN24:
		return "Quang Binh"
	case SubdivisionVN25:
		return "Quang Tri"
	case SubdivisionVN26:
		return "Thua Thien-Hue"
	case SubdivisionVN27:
		return "Quang Nam"
	case SubdivisionVN28:
		return "Kon Tum"
	case SubdivisionVN29:
		return "Quang Ngai"
	case SubdivisionVN30:
		return "Gia Lai"
	case SubdivisionVN31:
		return "Binh Dinh"
	case SubdivisionVN32:
		return "Phu Yen"
	case SubdivisionVN33:
		return "Dak Lak"
	case SubdivisionVN34:
		return "Khanh Hoa"
	case SubdivisionVN35:
		return "Lam Dong"
	case SubdivisionVN36:
		return "Ninh Thuan"
	case SubdivisionVN37:
		return "Tay Ninh"
	case SubdivisionVN39:
		return "Dong Nai"
	case SubdivisionVN40:
		return "Binh Thuan"
	case SubdivisionVN41:
		return "Long An"
	case SubdivisionVN43:
		return "Ba Ria - Vung Tau"
	case SubdivisionVN44:
		return "An Giang"
	case SubdivisionVN45:
		return "Dong Thap"
	case SubdivisionVN46:
		return "Tien Giang"
	case SubdivisionVN47:
		return "Kien Giang"
	case SubdivisionVN49:
		return "Vinh Long"
	case SubdivisionVN50:
		return "Ben Tre"
	case SubdivisionVN51:
		return "Tra Vinh"
	case SubdivisionVN52:
		return "Soc Trang"
	case SubdivisionVN53:
		return "Bac Kan"
	case SubdivisionVN54:
		return "Bac Giang"
	case SubdivisionVN55:
		return "Bac Lieu"
	case SubdivisionVN56:
		return "Bac Ninh"
	case SubdivisionVN57:
		return "Binh Duong"
	case SubdivisionVN58:
		return "Binh Phuoc"
	case SubdivisionVN59:
		return "Ca Mau"
	case SubdivisionVN61:
		return "Hai Duong"
	case SubdivisionVN63:
		return "Ha Nam"
	case SubdivisionVN66:
		return "Hung Yen"
	case SubdivisionVN67:
		return "Nam Dinh"
	case SubdivisionVN68:
		return "Phu Tho"
	case SubdivisionVN69:
		return "Thai Nguyen"
	case SubdivisionVN70:
		return "Vinh Phuc"
	case SubdivisionVN71:
		return "Dien Bien"
	case SubdivisionVN72:
		return "Dak Nong"
	case SubdivisionVNCT:
		return "Can Tho"
	case SubdivisionVNDN:
		return "Da Nang"
	case SubdivisionVNHN:
		return "Ha Noi"
	case SubdivisionVNHP:
		return "Hai Phong"
	case SubdivisionVNSG:
		return "Ho Chi Minh"
	case SubdivisionVUSEE:
		return "Shefa"
	case SubdivisionVUTAE:
		return "Tafea"
	case SubdivisionVUTOB:
		return "Torba"
	case SubdivisionWFSG:
		return "Sigave"
	case SubdivisionWFUV:
		return "Uvea"
	case SubdivisionWSAT:
		return "Atua"
	case SubdivisionWSFA:
		return "Fa'asaleleaga"
	case SubdivisionWSTU:
		return "Tuamasaga"
	case SubdivisionYEAB:
		return "Abyan"
	case SubdivisionYEAD:
		return "'Adan"
	case SubdivisionYEAM:
		return "'Amran"
	case SubdivisionYEBA:
		return "Al Bayda'"
	case SubdivisionYEDA:
		return "Ad Dali'"
	case SubdivisionYEDH:
		return "Dhamar"
	case SubdivisionYEHD:
		return "Hadramawt"
	case SubdivisionYEHJ:
		return "Hajjah"
	case SubdivisionYEHU:
		return "Al Hudaydah"
	case SubdivisionYEIB:
		return "Ibb"
	case SubdivisionYELA:
		return "Lahij"
	case SubdivisionYEMA:
		return "Ma'rib"
	case SubdivisionYEMR:
		return "Al Mahrah"
	case SubdivisionYESA:
		return "Amanat al 'Asimah"
	case SubdivisionYESD:
		return "Sa'dah"
	case SubdivisionYESH:
		return "Shabwah"
	case SubdivisionYESN:
		return "San'a'"
	case SubdivisionYETA:
		return "Ta'izz"
	case SubdivisionZAEC:
		return "Eastern Cape"
	case SubdivisionZAFS:
		return "Free State"
	case SubdivisionZAGP:
		return "Gauteng"
	case SubdivisionZAKZN:
		return "Kwazulu-Natal"
	case SubdivisionZALP:
		return "Limpopo"
	case SubdivisionZAMP:
		return "Mpumalanga"
	case SubdivisionZANC:
		return "Northern Cape"
	case SubdivisionZANW:
		return "North-West"
	case SubdivisionZAWC:
		return "Western Cape"
	case SubdivisionZM01:
		return "Western"
	case SubdivisionZM02:
		return "Central"
	case SubdivisionZM03:
		return "Eastern"
	case SubdivisionZM04:
		return "Luapula"
	case SubdivisionZM05:
		return "Northern"
	case SubdivisionZM06:
		return "North-Western"
	case SubdivisionZM07:
		return "Southern"
	case SubdivisionZM08:
		return "Copperbelt"
	case SubdivisionZM09:
		return "Lusaka"
	case SubdivisionZM10:
		return "Muchinga"
	case SubdivisionZWBU:
		return "Bulawayo"
	case SubdivisionZWHA:
		return "Harare"
	case SubdivisionZWMA:
		return "Manicaland"
	case SubdivisionZWMC:
		return "Mashonaland Central"
	case SubdivisionZWME:
		return "Mashonaland East"
	case SubdivisionZWMI:
		return "Midlands"
	case SubdivisionZWMN:
		return "Matabeleland North"
	case SubdivisionZWMS:
		return "Matabeleland South"
	case SubdivisionZWMV:
		return "Masvingo"
	case SubdivisionZWMW:
		return "Mashonaland West"
	}
	return UnknownMsg
}

// Country - returns a country of the subdivision
//nolint:cyclop,funlen,gocyclo
func (s SubdivisionCode) Country() CountryCode {
	switch s {
	case SubdivisionAD02:
		return AD
	case SubdivisionAD03:
		return AD
	case SubdivisionAD04:
		return AD
	case SubdivisionAD05:
		return AD
	case SubdivisionAD06:
		return AD
	case SubdivisionAD07:
		return AD
	case SubdivisionAD08:
		return AD
	case SubdivisionAEAJ:
		return AE
	case SubdivisionAEAZ:
		return AE
	case SubdivisionAEDU:
		return AE
	case SubdivisionAEFU:
		return AE
	case SubdivisionAERK:
		return AE
	case SubdivisionAESH:
		return AE
	case SubdivisionAEUQ:
		return AE
	case SubdivisionAFBAL:
		return AF
	case SubdivisionAFBAM:
		return AF
	case SubdivisionAFBDG:
		return AF
	case SubdivisionAFBDS:
		return AF
	case SubdivisionAFBGL:
		return AF
	case SubdivisionAFDAY:
		return AF
	case SubdivisionAFFRA:
		return AF
	case SubdivisionAFFYB:
		return AF
	case SubdivisionAFGHA:
		return AF
	case SubdivisionAFGHO:
		return AF
	case SubdivisionAFHEL:
		return AF
	case SubdivisionAFHER:
		return AF
	case SubdivisionAFJOW:
		return AF
	case SubdivisionAFKAB:
		return AF
	case SubdivisionAFKAN:
		return AF
	case SubdivisionAFKDZ:
		return AF
	case SubdivisionAFKHO:
		return AF
	case SubdivisionAFLAG:
		return AF
	case SubdivisionAFLOG:
		return AF
	case SubdivisionAFNAN:
		return AF
	case SubdivisionAFNIM:
		return AF
	case SubdivisionAFPAR:
		return AF
	case SubdivisionAFPIA:
		return AF
	case SubdivisionAFPKA:
		return AF
	case SubdivisionAFTAK:
		return AF
	case SubdivisionAFURU:
		return AF
	case SubdivisionAG03:
		return AG
	case SubdivisionAG04:
		return AG
	case SubdivisionAG05:
		return AG
	case SubdivisionAG06:
		return AG
	case SubdivisionAG07:
		return AG
	case SubdivisionAG08:
		return AG
	case SubdivisionAG11:
		return AG
	case SubdivisionAL01:
		return AL
	case SubdivisionAL02:
		return AL
	case SubdivisionAL03:
		return AL
	case SubdivisionAL04:
		return AL
	case SubdivisionAL05:
		return AL
	case SubdivisionAL06:
		return AL
	case SubdivisionAL07:
		return AL
	case SubdivisionAL08:
		return AL
	case SubdivisionAL09:
		return AL
	case SubdivisionAL10:
		return AL
	case SubdivisionAL11:
		return AL
	case SubdivisionAL12:
		return AL
	case SubdivisionAMAG:
		return AM
	case SubdivisionAMAR:
		return AM
	case SubdivisionAMAV:
		return AM
	case SubdivisionAMER:
		return AM
	case SubdivisionAMGR:
		return AM
	case SubdivisionAMKT:
		return AM
	case SubdivisionAMLO:
		return AM
	case SubdivisionAMSH:
		return AM
	case SubdivisionAMSU:
		return AM
	case SubdivisionAMTV:
		return AM
	case SubdivisionAMVD:
		return AM
	case SubdivisionAOBGO:
		return AO
	case SubdivisionAOBGU:
		return AO
	case SubdivisionAOBIE:
		return AO
	case SubdivisionAOCAB:
		return AO
	case SubdivisionAOCCU:
		return AO
	case SubdivisionAOCNN:
		return AO
	case SubdivisionAOCNO:
		return AO
	case SubdivisionAOCUS:
		return AO
	case SubdivisionAOHUA:
		return AO
	case SubdivisionAOHUI:
		return AO
	case SubdivisionAOLNO:
		return AO
	case SubdivisionAOLSU:
		return AO
	case SubdivisionAOLUA:
		return AO
	case SubdivisionAOMAL:
		return AO
	case SubdivisionAOMOX:
		return AO
	case SubdivisionAONAM:
		return AO
	case SubdivisionAOUIG:
		return AO
	case SubdivisionAOZAI:
		return AO
	case SubdivisionARA:
		return AR
	case SubdivisionARB:
		return AR
	case SubdivisionARC:
		return AR
	case SubdivisionARD:
		return AR
	case SubdivisionARE:
		return AR
	case SubdivisionARF:
		return AR
	case SubdivisionARG:
		return AR
	case SubdivisionARH:
		return AR
	case SubdivisionARJ:
		return AR
	case SubdivisionARK:
		return AR
	case SubdivisionARL:
		return AR
	case SubdivisionARM:
		return AR
	case SubdivisionARN:
		return AR
	case SubdivisionARP:
		return AR
	case SubdivisionARQ:
		return AR
	case SubdivisionARR:
		return AR
	case SubdivisionARS:
		return AR
	case SubdivisionART:
		return AR
	case SubdivisionARU:
		return AR
	case SubdivisionARV:
		return AR
	case SubdivisionARW:
		return AR
	case SubdivisionARX:
		return AR
	case SubdivisionARY:
		return AR
	case SubdivisionARZ:
		return AR
	case SubdivisionAT1:
		return AT
	case SubdivisionAT2:
		return AT
	case SubdivisionAT3:
		return AT
	case SubdivisionAT4:
		return AT
	case SubdivisionAT5:
		return AT
	case SubdivisionAT6:
		return AT
	case SubdivisionAT7:
		return AT
	case SubdivisionAT8:
		return AT
	case SubdivisionAT9:
		return AT
	case SubdivisionAUACT:
		return AU
	case SubdivisionAUNSW:
		return AU
	case SubdivisionAUNT:
		return AU
	case SubdivisionAUQLD:
		return AU
	case SubdivisionAUSA:
		return AU
	case SubdivisionAUTAS:
		return AU
	case SubdivisionAUVIC:
		return AU
	case SubdivisionAUWA:
		return AU
	case SubdivisionAZABS:
		return AZ
	case SubdivisionAZAGC:
		return AZ
	case SubdivisionAZAGS:
		return AZ
	case SubdivisionAZAST:
		return AZ
	case SubdivisionAZBA:
		return AZ
	case SubdivisionAZBAL:
		return AZ
	case SubdivisionAZBAR:
		return AZ
	case SubdivisionAZBEY:
		return AZ
	case SubdivisionAZBIL:
		return AZ
	case SubdivisionAZCAL:
		return AZ
	case SubdivisionAZFUZ:
		return AZ
	case SubdivisionAZGA:
		return AZ
	case SubdivisionAZGAD:
		return AZ
	case SubdivisionAZGOR:
		return AZ
	case SubdivisionAZGOY:
		return AZ
	case SubdivisionAZGYG:
		return AZ
	case SubdivisionAZIMI:
		return AZ
	case SubdivisionAZISM:
		return AZ
	case SubdivisionAZKUR:
		return AZ
	case SubdivisionAZLA:
		return AZ
	case SubdivisionAZMAS:
		return AZ
	case SubdivisionAZMI:
		return AZ
	case SubdivisionAZNEF:
		return AZ
	case SubdivisionAZNX:
		return AZ
	case SubdivisionAZOGU:
		return AZ
	case SubdivisionAZQAB:
		return AZ
	case SubdivisionAZQAX:
		return AZ
	case SubdivisionAZQBA:
		return AZ
	case SubdivisionAZQUS:
		return AZ
	case SubdivisionAZSAB:
		return AZ
	case SubdivisionAZSAK:
		return AZ
	case SubdivisionAZSAL:
		return AZ
	case SubdivisionAZSAT:
		return AZ
	case SubdivisionAZSIY:
		return AZ
	case SubdivisionAZSKR:
		return AZ
	case SubdivisionAZSM:
		return AZ
	case SubdivisionAZSMI:
		return AZ
	case SubdivisionAZSMX:
		return AZ
	case SubdivisionAZSR:
		return AZ
	case SubdivisionAZTAR:
		return AZ
	case SubdivisionAZXAC:
		return AZ
	case SubdivisionAZXIZ:
		return AZ
	case SubdivisionAZYAR:
		return AZ
	case SubdivisionAZYEV:
		return AZ
	case SubdivisionAZZAQ:
		return AZ
	case SubdivisionAZZAR:
		return AZ
	case SubdivisionBABIH:
		return BA
	case SubdivisionBABRC:
		return BA
	case SubdivisionBASRP:
		return BA
	case SubdivisionBB01:
		return BB
	case SubdivisionBB02:
		return BB
	case SubdivisionBB03:
		return BB
	case SubdivisionBB04:
		return BB
	case SubdivisionBB05:
		return BB
	case SubdivisionBB07:
		return BB
	case SubdivisionBB08:
		return BB
	case SubdivisionBB09:
		return BB
	case SubdivisionBB10:
		return BB
	case SubdivisionBB11:
		return BB
	case SubdivisionBDA:
		return BD
	case SubdivisionBDB:
		return BD
	case SubdivisionBDC:
		return BD
	case SubdivisionBDD:
		return BD
	case SubdivisionBDE:
		return BD
	case SubdivisionBDF:
		return BD
	case SubdivisionBDG:
		return BD
	case SubdivisionBEBRU:
		return BE
	case SubdivisionBEVAN:
		return BE
	case SubdivisionBEVBR:
		return BE
	case SubdivisionBEVLI:
		return BE
	case SubdivisionBEVOV:
		return BE
	case SubdivisionBEVWV:
		return BE
	case SubdivisionBEWBR:
		return BE
	case SubdivisionBEWHT:
		return BE
	case SubdivisionBEWLG:
		return BE
	case SubdivisionBEWLX:
		return BE
	case SubdivisionBEWNA:
		return BE
	case SubdivisionBFBAM:
		return BF
	case SubdivisionBFBAZ:
		return BF
	case SubdivisionBFBLG:
		return BF
	case SubdivisionBFBLK:
		return BF
	case SubdivisionBFGAN:
		return BF
	case SubdivisionBFGNA:
		return BF
	case SubdivisionBFGOU:
		return BF
	case SubdivisionBFHOU:
		return BF
	case SubdivisionBFKAD:
		return BF
	case SubdivisionBFKMD:
		return BF
	case SubdivisionBFKMP:
		return BF
	case SubdivisionBFKOP:
		return BF
	case SubdivisionBFKOT:
		return BF
	case SubdivisionBFKOW:
		return BF
	case SubdivisionBFLER:
		return BF
	case SubdivisionBFLOR:
		return BF
	case SubdivisionBFMOU:
		return BF
	case SubdivisionBFNAM:
		return BF
	case SubdivisionBFNAO:
		return BF
	case SubdivisionBFNAY:
		return BF
	case SubdivisionBFOUB:
		return BF
	case SubdivisionBFOUD:
		return BF
	case SubdivisionBFPAS:
		return BF
	case SubdivisionBFPON:
		return BF
	case SubdivisionBFSEN:
		return BF
	case SubdivisionBFSIS:
		return BF
	case SubdivisionBFSMT:
		return BF
	case SubdivisionBFSOM:
		return BF
	case SubdivisionBFSOR:
		return BF
	case SubdivisionBFTAP:
		return BF
	case SubdivisionBFTUI:
		return BF
	case SubdivisionBFYAT:
		return BF
	case SubdivisionBFZIR:
		return BF
	case SubdivisionBFZON:
		return BF
	case SubdivisionBFZOU:
		return BF
	case SubdivisionBG01:
		return BG
	case SubdivisionBG02:
		return BG
	case SubdivisionBG03:
		return BG
	case SubdivisionBG04:
		return BG
	case SubdivisionBG05:
		return BG
	case SubdivisionBG06:
		return BG
	case SubdivisionBG07:
		return BG
	case SubdivisionBG08:
		return BG
	case SubdivisionBG09:
		return BG
	case SubdivisionBG10:
		return BG
	case SubdivisionBG11:
		return BG
	case SubdivisionBG12:
		return BG
	case SubdivisionBG13:
		return BG
	case SubdivisionBG14:
		return BG
	case SubdivisionBG15:
		return BG
	case SubdivisionBG16:
		return BG
	case SubdivisionBG17:
		return BG
	case SubdivisionBG18:
		return BG
	case SubdivisionBG19:
		return BG
	case SubdivisionBG20:
		return BG
	case SubdivisionBG21:
		return BG
	case SubdivisionBG22:
		return BG
	case SubdivisionBG23:
		return BG
	case SubdivisionBG24:
		return BG
	case SubdivisionBG25:
		return BG
	case SubdivisionBG26:
		return BG
	case SubdivisionBG27:
		return BG
	case SubdivisionBG28:
		return BG
	case SubdivisionBH13:
		return BH
	case SubdivisionBH14:
		return BH
	case SubdivisionBH15:
		return BH
	case SubdivisionBH17:
		return BH
	case SubdivisionBIBM:
		return BI
	case SubdivisionBICI:
		return BI
	case SubdivisionBIGI:
		return BI
	case SubdivisionBIKI:
		return BI
	case SubdivisionBIMW:
		return BI
	case SubdivisionBING:
		return BI
	case SubdivisionBIRM:
		return BI
	case SubdivisionBIRT:
		return BI
	case SubdivisionBIRY:
		return BI
	case SubdivisionBJAK:
		return BJ
	case SubdivisionBJAL:
		return BJ
	case SubdivisionBJAQ:
		return BJ
	case SubdivisionBJBO:
		return BJ
	case SubdivisionBJCO:
		return BJ
	case SubdivisionBJDO:
		return BJ
	case SubdivisionBJLI:
		return BJ
	case SubdivisionBJMO:
		return BJ
	case SubdivisionBJOU:
		return BJ
	case SubdivisionBJPL:
		return BJ
	case SubdivisionBJZO:
		return BJ
	case SubdivisionBNBE:
		return BN
	case SubdivisionBNBM:
		return BN
	case SubdivisionBNTE:
		return BN
	case SubdivisionBNTU:
		return BN
	case SubdivisionBOB:
		return BO
	case SubdivisionBOC:
		return BO
	case SubdivisionBOH:
		return BO
	case SubdivisionBOL:
		return BO
	case SubdivisionBON:
		return BO
	case SubdivisionBOO:
		return BO
	case SubdivisionBOP:
		return BO
	case SubdivisionBOS:
		return BO
	case SubdivisionBOT:
		return BO
	case SubdivisionBQBO:
		return BQ
	case SubdivisionBQSA:
		return BQ
	case SubdivisionBQSE:
		return BQ
	case SubdivisionBRAC:
		return BR
	case SubdivisionBRAL:
		return BR
	case SubdivisionBRAM:
		return BR
	case SubdivisionBRAP:
		return BR
	case SubdivisionBRBA:
		return BR
	case SubdivisionBRCE:
		return BR
	case SubdivisionBRDF:
		return BR
	case SubdivisionBRES:
		return BR
	case SubdivisionBRGO:
		return BR
	case SubdivisionBRMA:
		return BR
	case SubdivisionBRMG:
		return BR
	case SubdivisionBRMS:
		return BR
	case SubdivisionBRMT:
		return BR
	case SubdivisionBRPA:
		return BR
	case SubdivisionBRPB:
		return BR
	case SubdivisionBRPE:
		return BR
	case SubdivisionBRPI:
		return BR
	case SubdivisionBRPR:
		return BR
	case SubdivisionBRRJ:
		return BR
	case SubdivisionBRRN:
		return BR
	case SubdivisionBRRO:
		return BR
	case SubdivisionBRRR:
		return BR
	case SubdivisionBRRS:
		return BR
	case SubdivisionBRSC:
		return BR
	case SubdivisionBRSE:
		return BR
	case SubdivisionBRSP:
		return BR
	case SubdivisionBRTO:
		return BR
	case SubdivisionBSCO:
		return BS
	case SubdivisionBSFP:
		return BS
	case SubdivisionBSLI:
		return BS
	case SubdivisionBSNO:
		return BS
	case SubdivisionBSNP:
		return BS
	case SubdivisionBSNS:
		return BS
	case SubdivisionBSSE:
		return BS
	case SubdivisionBT11:
		return BT
	case SubdivisionBT12:
		return BT
	case SubdivisionBT13:
		return BT
	case SubdivisionBT14:
		return BT
	case SubdivisionBT15:
		return BT
	case SubdivisionBT21:
		return BT
	case SubdivisionBT23:
		return BT
	case SubdivisionBT24:
		return BT
	case SubdivisionBT32:
		return BT
	case SubdivisionBT33:
		return BT
	case SubdivisionBT41:
		return BT
	case SubdivisionBT42:
		return BT
	case SubdivisionBT43:
		return BT
	case SubdivisionBT44:
		return BT
	case SubdivisionBT45:
		return BT
	case SubdivisionBTGA:
		return BT
	case SubdivisionBWCE:
		return BW
	case SubdivisionBWCH:
		return BW
	case SubdivisionBWKG:
		return BW
	case SubdivisionBWKL:
		return BW
	case SubdivisionBWKW:
		return BW
	case SubdivisionBWNE:
		return BW
	case SubdivisionBWNW:
		return BW
	case SubdivisionBWSE:
		return BW
	case SubdivisionBWSO:
		return BW
	case SubdivisionBYBR:
		return BY
	case SubdivisionBYHM:
		return BY
	case SubdivisionBYHO:
		return BY
	case SubdivisionBYHR:
		return BY
	case SubdivisionBYMA:
		return BY
	case SubdivisionBYMI:
		return BY
	case SubdivisionBYVI:
		return BY
	case SubdivisionBZBZ:
		return BZ
	case SubdivisionBZCY:
		return BZ
	case SubdivisionBZCZL:
		return BZ
	case SubdivisionBZOW:
		return BZ
	case SubdivisionBZSC:
		return BZ
	case SubdivisionBZTOL:
		return BZ
	case SubdivisionCAAB:
		return CA
	case SubdivisionCABC:
		return CA
	case SubdivisionCAMB:
		return CA
	case SubdivisionCANB:
		return CA
	case SubdivisionCANL:
		return CA
	case SubdivisionCANS:
		return CA
	case SubdivisionCANT:
		return CA
	case SubdivisionCANU:
		return CA
	case SubdivisionCAON:
		return CA
	case SubdivisionCAPE:
		return CA
	case SubdivisionCAQC:
		return CA
	case SubdivisionCASK:
		return CA
	case SubdivisionCAYT:
		return CA
	case SubdivisionCDEQ:
		return CD
	case SubdivisionCDHK:
		return CD
	case SubdivisionCDIT:
		return CD
	case SubdivisionCDKC:
		return CD
	case SubdivisionCDKE:
		return CD
	case SubdivisionCDKL:
		return CD
	case SubdivisionCDKN:
		return CD
	case SubdivisionCDLU:
		return CD
	case SubdivisionCDNK:
		return CD
	case SubdivisionCDSA:
		return CD
	case SubdivisionCDSK:
		return CD
	case SubdivisionCDTA:
		return CD
	case SubdivisionCDTO:
		return CD
	case SubdivisionCFAC:
		return CF
	case SubdivisionCFBGF:
		return CF
	case SubdivisionCFNM:
		return CF
	case SubdivisionCFOP:
		return CF
	case SubdivisionCG13:
		return CG
	case SubdivisionCG16:
		return CG
	case SubdivisionCGBZV:
		return CG
	case SubdivisionCHAG:
		return CH
	case SubdivisionCHAI:
		return CH
	case SubdivisionCHAR:
		return CH
	case SubdivisionCHBE:
		return CH
	case SubdivisionCHBL:
		return CH
	case SubdivisionCHBS:
		return CH
	case SubdivisionCHFR:
		return CH
	case SubdivisionCHGE:
		return CH
	case SubdivisionCHGL:
		return CH
	case SubdivisionCHGR:
		return CH
	case SubdivisionCHJU:
		return CH
	case SubdivisionCHLU:
		return CH
	case SubdivisionCHNE:
		return CH
	case SubdivisionCHNW:
		return CH
	case SubdivisionCHOW:
		return CH
	case SubdivisionCHSG:
		return CH
	case SubdivisionCHSH:
		return CH
	case SubdivisionCHSO:
		return CH
	case SubdivisionCHSZ:
		return CH
	case SubdivisionCHTG:
		return CH
	case SubdivisionCHTI:
		return CH
	case SubdivisionCHUR:
		return CH
	case SubdivisionCHVD:
		return CH
	case SubdivisionCHVS:
		return CH
	case SubdivisionCHZG:
		return CH
	case SubdivisionCHZH:
		return CH
	case SubdivisionCIAB:
		return CI
	case SubdivisionCIBS:
		return CI
	case SubdivisionCICM:
		return CI
	case SubdivisionCIDN:
		return CI
	case SubdivisionCIGD:
		return CI
	case SubdivisionCILC:
		return CI
	case SubdivisionCILG:
		return CI
	case SubdivisionCIMG:
		return CI
	case SubdivisionCISM:
		return CI
	case SubdivisionCISV:
		return CI
	case SubdivisionCIVB:
		return CI
	case SubdivisionCIWR:
		return CI
	case SubdivisionCIYM:
		return CI
	case SubdivisionCIZZ:
		return CI
	case SubdivisionCLAI:
		return CL
	case SubdivisionCLAN:
		return CL
	case SubdivisionCLAP:
		return CL
	case SubdivisionCLAR:
		return CL
	case SubdivisionCLAT:
		return CL
	case SubdivisionCLBI:
		return CL
	case SubdivisionCLCO:
		return CL
	case SubdivisionCLLI:
		return CL
	case SubdivisionCLLL:
		return CL
	case SubdivisionCLLR:
		return CL
	case SubdivisionCLMA:
		return CL
	case SubdivisionCLML:
		return CL
	case SubdivisionCLNB:
		return CL
	case SubdivisionCLRM:
		return CL
	case SubdivisionCLTA:
		return CL
	case SubdivisionCLVS:
		return CL
	case SubdivisionCMAD:
		return CM
	case SubdivisionCMCE:
		return CM
	case SubdivisionCMEN:
		return CM
	case SubdivisionCMES:
		return CM
	case SubdivisionCMLT:
		return CM
	case SubdivisionCMNO:
		return CM
	case SubdivisionCMNW:
		return CM
	case SubdivisionCMOU:
		return CM
	case SubdivisionCMSU:
		return CM
	case SubdivisionCMSW:
		return CM
	case SubdivisionCNAH:
		return CN
	case SubdivisionCNBJ:
		return CN
	case SubdivisionCNCQ:
		return CN
	case SubdivisionCNFJ:
		return CN
	case SubdivisionCNGD:
		return CN
	case SubdivisionCNGS:
		return CN
	case SubdivisionCNGX:
		return CN
	case SubdivisionCNGZ:
		return CN
	case SubdivisionCNHA:
		return CN
	case SubdivisionCNHB:
		return CN
	case SubdivisionCNHE:
		return CN
	case SubdivisionCNHI:
		return CN
	case SubdivisionCNHL:
		return CN
	case SubdivisionCNHN:
		return CN
	case SubdivisionCNJL:
		return CN
	case SubdivisionCNJS:
		return CN
	case SubdivisionCNJX:
		return CN
	case SubdivisionCNLN:
		return CN
	case SubdivisionCNNM:
		return CN
	case SubdivisionCNNX:
		return CN
	case SubdivisionCNQH:
		return CN
	case SubdivisionCNSC:
		return CN
	case SubdivisionCNSD:
		return CN
	case SubdivisionCNSH:
		return CN
	case SubdivisionCNSN:
		return CN
	case SubdivisionCNSX:
		return CN
	case SubdivisionCNTJ:
		return CN
	case SubdivisionCNXJ:
		return CN
	case SubdivisionCNXZ:
		return CN
	case SubdivisionCNYN:
		return CN
	case SubdivisionCNZJ:
		return CN
	case SubdivisionCOAMA:
		return CO
	case SubdivisionCOANT:
		return CO
	case SubdivisionCOARA:
		return CO
	case SubdivisionCOATL:
		return CO
	case SubdivisionCOBOL:
		return CO
	case SubdivisionCOBOY:
		return CO
	case SubdivisionCOCAL:
		return CO
	case SubdivisionCOCAQ:
		return CO
	case SubdivisionCOCAS:
		return CO
	case SubdivisionCOCAU:
		return CO
	case SubdivisionCOCES:
		return CO
	case SubdivisionCOCHO:
		return CO
	case SubdivisionCOCOR:
		return CO
	case SubdivisionCOCUN:
		return CO
	case SubdivisionCODC:
		return CO
	case SubdivisionCOGUA:
		return CO
	case SubdivisionCOGUV:
		return CO
	case SubdivisionCOHUI:
		return CO
	case SubdivisionCOLAG:
		return CO
	case SubdivisionCOMAG:
		return CO
	case SubdivisionCOMET:
		return CO
	case SubdivisionCONAR:
		return CO
	case SubdivisionCONSA:
		return CO
	case SubdivisionCOPUT:
		return CO
	case SubdivisionCOQUI:
		return CO
	case SubdivisionCORIS:
		return CO
	case SubdivisionCOSAN:
		return CO
	case SubdivisionCOSAP:
		return CO
	case SubdivisionCOSUC:
		return CO
	case SubdivisionCOTOL:
		return CO
	case SubdivisionCOVAC:
		return CO
	case SubdivisionCOVID:
		return CO
	case SubdivisionCRA:
		return CR
	case SubdivisionCRC:
		return CR
	case SubdivisionCRG:
		return CR
	case SubdivisionCRH:
		return CR
	case SubdivisionCRL:
		return CR
	case SubdivisionCRP:
		return CR
	case SubdivisionCRSJ:
		return CR
	case SubdivisionCU01:
		return CU
	case SubdivisionCU03:
		return CU
	case SubdivisionCU04:
		return CU
	case SubdivisionCU05:
		return CU
	case SubdivisionCU06:
		return CU
	case SubdivisionCU07:
		return CU
	case SubdivisionCU08:
		return CU
	case SubdivisionCU09:
		return CU
	case SubdivisionCU10:
		return CU
	case SubdivisionCU11:
		return CU
	case SubdivisionCU12:
		return CU
	case SubdivisionCU13:
		return CU
	case SubdivisionCU14:
		return CU
	case SubdivisionCU15:
		return CU
	case SubdivisionCU16:
		return CU
	case SubdivisionCVBR:
		return CV
	case SubdivisionCVBV:
		return CV
	case SubdivisionCVCR:
		return CV
	case SubdivisionCVMA:
		return CV
	case SubdivisionCVPN:
		return CV
	case SubdivisionCVPR:
		return CV
	case SubdivisionCVRG:
		return CV
	case SubdivisionCVRS:
		return CV
	case SubdivisionCVSF:
		return CV
	case SubdivisionCVSL:
		return CV
	case SubdivisionCVSV:
		return CV
	case SubdivisionCVTA:
		return CV
	case SubdivisionCY01:
		return CY
	case SubdivisionCY02:
		return CY
	case SubdivisionCY03:
		return CY
	case SubdivisionCY04:
		return CY
	case SubdivisionCY05:
		return CY
	case SubdivisionCY06:
		return CY
	case SubdivisionCZ10:
		return CZ
	case SubdivisionCZ20:
		return CZ
	case SubdivisionCZ31:
		return CZ
	case SubdivisionCZ32:
		return CZ
	case SubdivisionCZ41:
		return CZ
	case SubdivisionCZ42:
		return CZ
	case SubdivisionCZ51:
		return CZ
	case SubdivisionCZ52:
		return CZ
	case SubdivisionCZ53:
		return CZ
	case SubdivisionCZ63:
		return CZ
	case SubdivisionCZ64:
		return CZ
	case SubdivisionCZ71:
		return CZ
	case SubdivisionCZ72:
		return CZ
	case SubdivisionCZ80:
		return CZ
	case SubdivisionDEBB:
		return DE
	case SubdivisionDEBE:
		return DE
	case SubdivisionDEBW:
		return DE
	case SubdivisionDEBY:
		return DE
	case SubdivisionDEHB:
		return DE
	case SubdivisionDEHE:
		return DE
	case SubdivisionDEHH:
		return DE
	case SubdivisionDEMV:
		return DE
	case SubdivisionDENI:
		return DE
	case SubdivisionDENW:
		return DE
	case SubdivisionDERP:
		return DE
	case SubdivisionDESH:
		return DE
	case SubdivisionDESL:
		return DE
	case SubdivisionDESN:
		return DE
	case SubdivisionDEST:
		return DE
	case SubdivisionDETH:
		return DE
	case SubdivisionDJDJ:
		return DJ
	case SubdivisionDK81:
		return DK
	case SubdivisionDK82:
		return DK
	case SubdivisionDK83:
		return DK
	case SubdivisionDK84:
		return DK
	case SubdivisionDK85:
		return DK
	case SubdivisionDM02:
		return DM
	case SubdivisionDM04:
		return DM
	case SubdivisionDM05:
		return DM
	case SubdivisionDM09:
		return DM
	case SubdivisionDM10:
		return DM
	case SubdivisionDO01:
		return DO
	case SubdivisionDO02:
		return DO
	case SubdivisionDO03:
		return DO
	case SubdivisionDO04:
		return DO
	case SubdivisionDO05:
		return DO
	case SubdivisionDO06:
		return DO
	case SubdivisionDO07:
		return DO
	case SubdivisionDO08:
		return DO
	case SubdivisionDO09:
		return DO
	case SubdivisionDO10:
		return DO
	case SubdivisionDO11:
		return DO
	case SubdivisionDO12:
		return DO
	case SubdivisionDO13:
		return DO
	case SubdivisionDO14:
		return DO
	case SubdivisionDO15:
		return DO
	case SubdivisionDO17:
		return DO
	case SubdivisionDO18:
		return DO
	case SubdivisionDO19:
		return DO
	case SubdivisionDO20:
		return DO
	case SubdivisionDO21:
		return DO
	case SubdivisionDO22:
		return DO
	case SubdivisionDO23:
		return DO
	case SubdivisionDO24:
		return DO
	case SubdivisionDO25:
		return DO
	case SubdivisionDO26:
		return DO
	case SubdivisionDO27:
		return DO
	case SubdivisionDO28:
		return DO
	case SubdivisionDO29:
		return DO
	case SubdivisionDO30:
		return DO
	case SubdivisionDO31:
		return DO
	case SubdivisionDZ01:
		return DZ
	case SubdivisionDZ02:
		return DZ
	case SubdivisionDZ03:
		return DZ
	case SubdivisionDZ04:
		return DZ
	case SubdivisionDZ05:
		return DZ
	case SubdivisionDZ06:
		return DZ
	case SubdivisionDZ07:
		return DZ
	case SubdivisionDZ08:
		return DZ
	case SubdivisionDZ09:
		return DZ
	case SubdivisionDZ10:
		return DZ
	case SubdivisionDZ11:
		return DZ
	case SubdivisionDZ12:
		return DZ
	case SubdivisionDZ13:
		return DZ
	case SubdivisionDZ14:
		return DZ
	case SubdivisionDZ15:
		return DZ
	case SubdivisionDZ16:
		return DZ
	case SubdivisionDZ17:
		return DZ
	case SubdivisionDZ18:
		return DZ
	case SubdivisionDZ19:
		return DZ
	case SubdivisionDZ20:
		return DZ
	case SubdivisionDZ21:
		return DZ
	case SubdivisionDZ22:
		return DZ
	case SubdivisionDZ23:
		return DZ
	case SubdivisionDZ24:
		return DZ
	case SubdivisionDZ25:
		return DZ
	case SubdivisionDZ26:
		return DZ
	case SubdivisionDZ27:
		return DZ
	case SubdivisionDZ28:
		return DZ
	case SubdivisionDZ29:
		return DZ
	case SubdivisionDZ30:
		return DZ
	case SubdivisionDZ31:
		return DZ
	case SubdivisionDZ32:
		return DZ
	case SubdivisionDZ33:
		return DZ
	case SubdivisionDZ34:
		return DZ
	case SubdivisionDZ35:
		return DZ
	case SubdivisionDZ36:
		return DZ
	case SubdivisionDZ37:
		return DZ
	case SubdivisionDZ38:
		return DZ
	case SubdivisionDZ39:
		return DZ
	case SubdivisionDZ40:
		return DZ
	case SubdivisionDZ41:
		return DZ
	case SubdivisionDZ42:
		return DZ
	case SubdivisionDZ43:
		return DZ
	case SubdivisionDZ44:
		return DZ
	case SubdivisionDZ45:
		return DZ
	case SubdivisionDZ46:
		return DZ
	case SubdivisionDZ47:
		return DZ
	case SubdivisionDZ48:
		return DZ
	case SubdivisionECA:
		return EC
	case SubdivisionECB:
		return EC
	case SubdivisionECC:
		return EC
	case SubdivisionECD:
		return EC
	case SubdivisionECE:
		return EC
	case SubdivisionECF:
		return EC
	case SubdivisionECG:
		return EC
	case SubdivisionECH:
		return EC
	case SubdivisionECI:
		return EC
	case SubdivisionECL:
		return EC
	case SubdivisionECM:
		return EC
	case SubdivisionECN:
		return EC
	case SubdivisionECO:
		return EC
	case SubdivisionECP:
		return EC
	case SubdivisionECR:
		return EC
	case SubdivisionECS:
		return EC
	case SubdivisionECSD:
		return EC
	case SubdivisionECSE:
		return EC
	case SubdivisionECT:
		return EC
	case SubdivisionECU:
		return EC
	case SubdivisionECW:
		return EC
	case SubdivisionECX:
		return EC
	case SubdivisionECY:
		return EC
	case SubdivisionECZ:
		return EC
	case SubdivisionEE37:
		return EE
	case SubdivisionEE39:
		return EE
	case SubdivisionEE44:
		return EE
	case SubdivisionEE49:
		return EE
	case SubdivisionEE51:
		return EE
	case SubdivisionEE57:
		return EE
	case SubdivisionEE59:
		return EE
	case SubdivisionEE65:
		return EE
	case SubdivisionEE67:
		return EE
	case SubdivisionEE70:
		return EE
	case SubdivisionEE74:
		return EE
	case SubdivisionEE78:
		return EE
	case SubdivisionEE82:
		return EE
	case SubdivisionEE84:
		return EE
	case SubdivisionEE86:
		return EE
	case SubdivisionEGALX:
		return EG
	case SubdivisionEGASN:
		return EG
	case SubdivisionEGAST:
		return EG
	case SubdivisionEGBA:
		return EG
	case SubdivisionEGBH:
		return EG
	case SubdivisionEGBNS:
		return EG
	case SubdivisionEGC:
		return EG
	case SubdivisionEGDK:
		return EG
	case SubdivisionEGDT:
		return EG
	case SubdivisionEGFYM:
		return EG
	case SubdivisionEGGH:
		return EG
	case SubdivisionEGGZ:
		return EG
	case SubdivisionEGIS:
		return EG
	case SubdivisionEGJS:
		return EG
	case SubdivisionEGKB:
		return EG
	case SubdivisionEGKFS:
		return EG
	case SubdivisionEGKN:
		return EG
	case SubdivisionEGLX:
		return EG
	case SubdivisionEGMN:
		return EG
	case SubdivisionEGMNF:
		return EG
	case SubdivisionEGMT:
		return EG
	case SubdivisionEGPTS:
		return EG
	case SubdivisionEGSHG:
		return EG
	case SubdivisionEGSHR:
		return EG
	case SubdivisionEGSIN:
		return EG
	case SubdivisionEGSUZ:
		return EG
	case SubdivisionEGWAD:
		return EG
	case SubdivisionERMA:
		return ER
	case SubdivisionESAN:
		return ES
	case SubdivisionESAR:
		return ES
	case SubdivisionESAS:
		return ES
	case SubdivisionESCB:
		return ES
	case SubdivisionESCE:
		return ES
	case SubdivisionESCL:
		return ES
	case SubdivisionESCM:
		return ES
	case SubdivisionESCN:
		return ES
	case SubdivisionESCT:
		return ES
	case SubdivisionESEX:
		return ES
	case SubdivisionESGA:
		return ES
	case SubdivisionESIB:
		return ES
	case SubdivisionESMC:
		return ES
	case SubdivisionESMD:
		return ES
	case SubdivisionESML:
		return ES
	case SubdivisionESNC:
		return ES
	case SubdivisionESPV:
		return ES
	case SubdivisionESRI:
		return ES
	case SubdivisionESVC:
		return ES
	case SubdivisionETAA:
		return ET
	case SubdivisionETAF:
		return ET
	case SubdivisionETAM:
		return ET
	case SubdivisionETBE:
		return ET
	case SubdivisionETDD:
		return ET
	case SubdivisionETHA:
		return ET
	case SubdivisionETOR:
		return ET
	case SubdivisionETSN:
		return ET
	case SubdivisionETSO:
		return ET
	case SubdivisionETTI:
		return ET
	case SubdivisionFI02:
		return FI
	case SubdivisionFI03:
		return FI
	case SubdivisionFI04:
		return FI
	case SubdivisionFI05:
		return FI
	case SubdivisionFI06:
		return FI
	case SubdivisionFI07:
		return FI
	case SubdivisionFI08:
		return FI
	case SubdivisionFI09:
		return FI
	case SubdivisionFI10:
		return FI
	case SubdivisionFI11:
		return FI
	case SubdivisionFI12:
		return FI
	case SubdivisionFI13:
		return FI
	case SubdivisionFI14:
		return FI
	case SubdivisionFI15:
		return FI
	case SubdivisionFI16:
		return FI
	case SubdivisionFI17:
		return FI
	case SubdivisionFI18:
		return FI
	case SubdivisionFI19:
		return FI
	case SubdivisionFJC:
		return FJ
	case SubdivisionFJE:
		return FJ
	case SubdivisionFJN:
		return FJ
	case SubdivisionFJR:
		return FJ
	case SubdivisionFJW:
		return FJ
	case SubdivisionFMKSA:
		return FM
	case SubdivisionFMPNI:
		return FM
	case SubdivisionFMTRK:
		return FM
	case SubdivisionFMYAP:
		return FM
	case SubdivisionFRARA:
		return FR
	case SubdivisionFRBFC:
		return FR
	case SubdivisionFRBRE:
		return FR
	case SubdivisionFRCOR:
		return FR
	case SubdivisionFRCVL:
		return FR
	case SubdivisionFRGES:
		return FR
	case SubdivisionFRHDF:
		return FR
	case SubdivisionFRIDF:
		return FR
	case SubdivisionFRNAQ:
		return FR
	case SubdivisionFRNOR:
		return FR
	case SubdivisionFROCC:
		return FR
	case SubdivisionFRPAC:
		return FR
	case SubdivisionFRPDL:
		return FR
	case SubdivisionGA1:
		return GA
	case SubdivisionGA2:
		return GA
	case SubdivisionGA4:
		return GA
	case SubdivisionGA8:
		return GA
	case SubdivisionGA9:
		return GA
	case SubdivisionGBENG:
		return GB
	case SubdivisionGBNIR:
		return GB
	case SubdivisionGBSCT:
		return GB
	case SubdivisionGBWLS:
		return GB
	case SubdivisionGD01:
		return GD
	case SubdivisionGD02:
		return GD
	case SubdivisionGD03:
		return GD
	case SubdivisionGD04:
		return GD
	case SubdivisionGD05:
		return GD
	case SubdivisionGD10:
		return GD
	case SubdivisionGEAB:
		return GE
	case SubdivisionGEAJ:
		return GE
	case SubdivisionGEGU:
		return GE
	case SubdivisionGEIM:
		return GE
	case SubdivisionGEKA:
		return GE
	case SubdivisionGEKK:
		return GE
	case SubdivisionGEMM:
		return GE
	case SubdivisionGERL:
		return GE
	case SubdivisionGESJ:
		return GE
	case SubdivisionGESK:
		return GE
	case SubdivisionGESZ:
		return GE
	case SubdivisionGETB:
		return GE
	case SubdivisionGHAA:
		return GH
	case SubdivisionGHAF:
		return GH
	case SubdivisionGHAH:
		return GH
	case SubdivisionGHBE:
		return GH
	case SubdivisionGHBO:
		return GH
	case SubdivisionGHCP:
		return GH
	case SubdivisionGHEP:
		return GH
	case SubdivisionGHNP:
		return GH
	case SubdivisionGHTV:
		return GH
	case SubdivisionGHUE:
		return GH
	case SubdivisionGHWP:
		return GH
	case SubdivisionGLAV:
		return GL
	case SubdivisionGLKU:
		return GL
	case SubdivisionGLQE:
		return GL
	case SubdivisionGLQT:
		return GL
	case SubdivisionGLSM:
		return GL
	case SubdivisionGMB:
		return GM
	case SubdivisionGML:
		return GM
	case SubdivisionGMM:
		return GM
	case SubdivisionGMN:
		return GM
	case SubdivisionGMU:
		return GM
	case SubdivisionGMW:
		return GM
	case SubdivisionGNB:
		return GN
	case SubdivisionGNBF:
		return GN
	case SubdivisionGNC:
		return GN
	case SubdivisionGNCO:
		return GN
	case SubdivisionGND:
		return GN
	case SubdivisionGNDB:
		return GN
	case SubdivisionGNK:
		return GN
	case SubdivisionGNN:
		return GN
	case SubdivisionGNSI:
		return GN
	case SubdivisionGQBN:
		return GQ
	case SubdivisionGQBS:
		return GQ
	case SubdivisionGQLI:
		return GQ
	case SubdivisionGQWN:
		return GQ
	case SubdivisionGR69:
		return GR
	case SubdivisionGRA:
		return GR
	case SubdivisionGRB:
		return GR
	case SubdivisionGRC:
		return GR
	case SubdivisionGRD:
		return GR
	case SubdivisionGRE:
		return GR
	case SubdivisionGRF:
		return GR
	case SubdivisionGRG:
		return GR
	case SubdivisionGRH:
		return GR
	case SubdivisionGRI:
		return GR
	case SubdivisionGRJ:
		return GR
	case SubdivisionGRK:
		return GR
	case SubdivisionGRL:
		return GR
	case SubdivisionGRM:
		return GR
	case SubdivisionGTAV:
		return GT
	case SubdivisionGTBV:
		return GT
	case SubdivisionGTCM:
		return GT
	case SubdivisionGTCQ:
		return GT
	case SubdivisionGTES:
		return GT
	case SubdivisionGTGU:
		return GT
	case SubdivisionGTHU:
		return GT
	case SubdivisionGTIZ:
		return GT
	case SubdivisionGTJA:
		return GT
	case SubdivisionGTJU:
		return GT
	case SubdivisionGTPE:
		return GT
	case SubdivisionGTPR:
		return GT
	case SubdivisionGTQC:
		return GT
	case SubdivisionGTQZ:
		return GT
	case SubdivisionGTRE:
		return GT
	case SubdivisionGTSA:
		return GT
	case SubdivisionGTSM:
		return GT
	case SubdivisionGTSO:
		return GT
	case SubdivisionGTSR:
		return GT
	case SubdivisionGTSU:
		return GT
	case SubdivisionGTTO:
		return GT
	case SubdivisionGTZA:
		return GT
	case SubdivisionGWBS:
		return GW
	case SubdivisionGWGA:
		return GW
	case SubdivisionGYBA:
		return GY
	case SubdivisionGYCU:
		return GY
	case SubdivisionGYDE:
		return GY
	case SubdivisionGYEB:
		return GY
	case SubdivisionGYES:
		return GY
	case SubdivisionGYMA:
		return GY
	case SubdivisionGYPT:
		return GY
	case SubdivisionGYUD:
		return GY
	case SubdivisionHNAT:
		return HN
	case SubdivisionHNCH:
		return HN
	case SubdivisionHNCL:
		return HN
	case SubdivisionHNCM:
		return HN
	case SubdivisionHNCP:
		return HN
	case SubdivisionHNCR:
		return HN
	case SubdivisionHNEP:
		return HN
	case SubdivisionHNFM:
		return HN
	case SubdivisionHNIB:
		return HN
	case SubdivisionHNIN:
		return HN
	case SubdivisionHNLE:
		return HN
	case SubdivisionHNLP:
		return HN
	case SubdivisionHNOC:
		return HN
	case SubdivisionHNOL:
		return HN
	case SubdivisionHNSB:
		return HN
	case SubdivisionHNVA:
		return HN
	case SubdivisionHNYO:
		return HN
	case SubdivisionHR01:
		return HR
	case SubdivisionHR02:
		return HR
	case SubdivisionHR03:
		return HR
	case SubdivisionHR04:
		return HR
	case SubdivisionHR05:
		return HR
	case SubdivisionHR06:
		return HR
	case SubdivisionHR07:
		return HR
	case SubdivisionHR08:
		return HR
	case SubdivisionHR09:
		return HR
	case SubdivisionHR10:
		return HR
	case SubdivisionHR11:
		return HR
	case SubdivisionHR12:
		return HR
	case SubdivisionHR13:
		return HR
	case SubdivisionHR14:
		return HR
	case SubdivisionHR15:
		return HR
	case SubdivisionHR16:
		return HR
	case SubdivisionHR17:
		return HR
	case SubdivisionHR18:
		return HR
	case SubdivisionHR19:
		return HR
	case SubdivisionHR20:
		return HR
	case SubdivisionHR21:
		return HR
	case SubdivisionHTAR:
		return HT
	case SubdivisionHTCE:
		return HT
	case SubdivisionHTND:
		return HT
	case SubdivisionHTOU:
		return HT
	case SubdivisionHTSD:
		return HT
	case SubdivisionHTSE:
		return HT
	case SubdivisionHUBA:
		return HU
	case SubdivisionHUBE:
		return HU
	case SubdivisionHUBK:
		return HU
	case SubdivisionHUBU:
		return HU
	case SubdivisionHUBZ:
		return HU
	case SubdivisionHUCS:
		return HU
	case SubdivisionHUFE:
		return HU
	case SubdivisionHUGS:
		return HU
	case SubdivisionHUHB:
		return HU
	case SubdivisionHUHE:
		return HU
	case SubdivisionHUJN:
		return HU
	case SubdivisionHUKE:
		return HU
	case SubdivisionHUNO:
		return HU
	case SubdivisionHUPE:
		return HU
	case SubdivisionHUSO:
		return HU
	case SubdivisionHUSZ:
		return HU
	case SubdivisionHUTO:
		return HU
	case SubdivisionHUVA:
		return HU
	case SubdivisionHUVE:
		return HU
	case SubdivisionHUZA:
		return HU
	case SubdivisionIDAC:
		return ID
	case SubdivisionIDBA:
		return ID
	case SubdivisionIDBB:
		return ID
	case SubdivisionIDBE:
		return ID
	case SubdivisionIDBT:
		return ID
	case SubdivisionIDGO:
		return ID
	case SubdivisionIDJA:
		return ID
	case SubdivisionIDJB:
		return ID
	case SubdivisionIDJI:
		return ID
	case SubdivisionIDJK:
		return ID
	case SubdivisionIDJT:
		return ID
	case SubdivisionIDKB:
		return ID
	case SubdivisionIDKI:
		return ID
	case SubdivisionIDKR:
		return ID
	case SubdivisionIDKS:
		return ID
	case SubdivisionIDKT:
		return ID
	case SubdivisionIDKU:
		return ID
	case SubdivisionIDLA:
		return ID
	case SubdivisionIDML:
		return ID
	case SubdivisionIDMU:
		return ID
	case SubdivisionIDNB:
		return ID
	case SubdivisionIDNT:
		return ID
	case SubdivisionIDPB:
		return ID
	case SubdivisionIDPP:
		return ID
	case SubdivisionIDRI:
		return ID
	case SubdivisionIDSA:
		return ID
	case SubdivisionIDSB:
		return ID
	case SubdivisionIDSG:
		return ID
	case SubdivisionIDSN:
		return ID
	case SubdivisionIDSR:
		return ID
	case SubdivisionIDSS:
		return ID
	case SubdivisionIDST:
		return ID
	case SubdivisionIDSU:
		return ID
	case SubdivisionIDYO:
		return ID
	case SubdivisionIECE:
		return IE
	case SubdivisionIECN:
		return IE
	case SubdivisionIECO:
		return IE
	case SubdivisionIECW:
		return IE
	case SubdivisionIED:
		return IE
	case SubdivisionIEDL:
		return IE
	case SubdivisionIEG:
		return IE
	case SubdivisionIEKE:
		return IE
	case SubdivisionIEKK:
		return IE
	case SubdivisionIEKY:
		return IE
	case SubdivisionIELD:
		return IE
	case SubdivisionIELH:
		return IE
	case SubdivisionIELK:
		return IE
	case SubdivisionIELM:
		return IE
	case SubdivisionIELS:
		return IE
	case SubdivisionIEMH:
		return IE
	case SubdivisionIEMN:
		return IE
	case SubdivisionIEMO:
		return IE
	case SubdivisionIEOY:
		return IE
	case SubdivisionIERN:
		return IE
	case SubdivisionIESO:
		return IE
	case SubdivisionIETA:
		return IE
	case SubdivisionIEWD:
		return IE
	case SubdivisionIEWH:
		return IE
	case SubdivisionIEWW:
		return IE
	case SubdivisionIEWX:
		return IE
	case SubdivisionILD:
		return IL
	case SubdivisionILHA:
		return IL
	case SubdivisionILJM:
		return IL
	case SubdivisionILM:
		return IL
	case SubdivisionILTA:
		return IL
	case SubdivisionILZ:
		return IL
	case SubdivisionINAN:
		return IN
	case SubdivisionINAP:
		return IN
	case SubdivisionINAR:
		return IN
	case SubdivisionINAS:
		return IN
	case SubdivisionINBR:
		return IN
	case SubdivisionINCH:
		return IN
	case SubdivisionINCT:
		return IN
	case SubdivisionINDH:
		return IN
	case SubdivisionINDL:
		return IN
	case SubdivisionINDN:
		return IN
	case SubdivisionINGA:
		return IN
	case SubdivisionINGJ:
		return IN
	case SubdivisionINHP:
		return IN
	case SubdivisionINHR:
		return IN
	case SubdivisionINJH:
		return IN
	case SubdivisionINJK:
		return IN
	case SubdivisionINKA:
		return IN
	case SubdivisionINKL:
		return IN
	case SubdivisionINLD:
		return IN
	case SubdivisionINMH:
		return IN
	case SubdivisionINML:
		return IN
	case SubdivisionINMN:
		return IN
	case SubdivisionINMP:
		return IN
	case SubdivisionINMZ:
		return IN
	case SubdivisionINNL:
		return IN
	case SubdivisionINOR:
		return IN
	case SubdivisionINPB:
		return IN
	case SubdivisionINPY:
		return IN
	case SubdivisionINRJ:
		return IN
	case SubdivisionINSK:
		return IN
	case SubdivisionINTG:
		return IN
	case SubdivisionINTN:
		return IN
	case SubdivisionINTR:
		return IN
	case SubdivisionINUP:
		return IN
	case SubdivisionINUT:
		return IN
	case SubdivisionINWB:
		return IN
	case SubdivisionIQAN:
		return IQ
	case SubdivisionIQAR:
		return IQ
	case SubdivisionIQBA:
		return IQ
	case SubdivisionIQBB:
		return IQ
	case SubdivisionIQBG:
		return IQ
	case SubdivisionIQDA:
		return IQ
	case SubdivisionIQDI:
		return IQ
	case SubdivisionIQDQ:
		return IQ
	case SubdivisionIQKA:
		return IQ
	case SubdivisionIQKI:
		return IQ
	case SubdivisionIQMA:
		return IQ
	case SubdivisionIQMU:
		return IQ
	case SubdivisionIQNA:
		return IQ
	case SubdivisionIQNI:
		return IQ
	case SubdivisionIQQA:
		return IQ
	case SubdivisionIQSD:
		return IQ
	case SubdivisionIQSU:
		return IQ
	case SubdivisionIQWA:
		return IQ
	case SubdivisionIR01:
		return IR
	case SubdivisionIR02:
		return IR
	case SubdivisionIR03:
		return IR
	case SubdivisionIR04:
		return IR
	case SubdivisionIR05:
		return IR
	case SubdivisionIR06:
		return IR
	case SubdivisionIR07:
		return IR
	case SubdivisionIR08:
		return IR
	case SubdivisionIR10:
		return IR
	case SubdivisionIR11:
		return IR
	case SubdivisionIR12:
		return IR
	case SubdivisionIR13:
		return IR
	case SubdivisionIR14:
		return IR
	case SubdivisionIR15:
		return IR
	case SubdivisionIR16:
		return IR
	case SubdivisionIR17:
		return IR
	case SubdivisionIR18:
		return IR
	case SubdivisionIR19:
		return IR
	case SubdivisionIR20:
		return IR
	case SubdivisionIR21:
		return IR
	case SubdivisionIR22:
		return IR
	case SubdivisionIR23:
		return IR
	case SubdivisionIR24:
		return IR
	case SubdivisionIR25:
		return IR
	case SubdivisionIR26:
		return IR
	case SubdivisionIR27:
		return IR
	case SubdivisionIR28:
		return IR
	case SubdivisionIR29:
		return IR
	case SubdivisionIR30:
		return IR
	case SubdivisionIR31:
		return IR
	case SubdivisionIR32:
		return IR
	case SubdivisionIS1:
		return IS
	case SubdivisionIS2:
		return IS
	case SubdivisionIS3:
		return IS
	case SubdivisionIS4:
		return IS
	case SubdivisionIS5:
		return IS
	case SubdivisionIS6:
		return IS
	case SubdivisionIS7:
		return IS
	case SubdivisionIS8:
		return IS
	case SubdivisionIT21:
		return IT
	case SubdivisionIT23:
		return IT
	case SubdivisionIT25:
		return IT
	case SubdivisionIT32:
		return IT
	case SubdivisionIT34:
		return IT
	case SubdivisionIT36:
		return IT
	case SubdivisionIT42:
		return IT
	case SubdivisionIT45:
		return IT
	case SubdivisionIT52:
		return IT
	case SubdivisionIT55:
		return IT
	case SubdivisionIT57:
		return IT
	case SubdivisionIT62:
		return IT
	case SubdivisionIT65:
		return IT
	case SubdivisionIT67:
		return IT
	case SubdivisionIT72:
		return IT
	case SubdivisionIT75:
		return IT
	case SubdivisionIT77:
		return IT
	case SubdivisionIT78:
		return IT
	case SubdivisionIT82:
		return IT
	case SubdivisionIT88:
		return IT
	case SubdivisionJM01:
		return JM
	case SubdivisionJM02:
		return JM
	case SubdivisionJM03:
		return JM
	case SubdivisionJM04:
		return JM
	case SubdivisionJM05:
		return JM
	case SubdivisionJM06:
		return JM
	case SubdivisionJM07:
		return JM
	case SubdivisionJM08:
		return JM
	case SubdivisionJM09:
		return JM
	case SubdivisionJM10:
		return JM
	case SubdivisionJM11:
		return JM
	case SubdivisionJM12:
		return JM
	case SubdivisionJM13:
		return JM
	case SubdivisionJM14:
		return JM
	case SubdivisionJOAJ:
		return JO
	case SubdivisionJOAM:
		return JO
	case SubdivisionJOAQ:
		return JO
	case SubdivisionJOAZ:
		return JO
	case SubdivisionJOBA:
		return JO
	case SubdivisionJOIR:
		return JO
	case SubdivisionJOJA:
		return JO
	case SubdivisionJOKA:
		return JO
	case SubdivisionJOMA:
		return JO
	case SubdivisionJOMD:
		return JO
	case SubdivisionJOMN:
		return JO
	case SubdivisionJP01:
		return JP
	case SubdivisionJP02:
		return JP
	case SubdivisionJP03:
		return JP
	case SubdivisionJP04:
		return JP
	case SubdivisionJP05:
		return JP
	case SubdivisionJP06:
		return JP
	case SubdivisionJP07:
		return JP
	case SubdivisionJP08:
		return JP
	case SubdivisionJP09:
		return JP
	case SubdivisionJP10:
		return JP
	case SubdivisionJP11:
		return JP
	case SubdivisionJP12:
		return JP
	case SubdivisionJP13:
		return JP
	case SubdivisionJP14:
		return JP
	case SubdivisionJP15:
		return JP
	case SubdivisionJP16:
		return JP
	case SubdivisionJP17:
		return JP
	case SubdivisionJP18:
		return JP
	case SubdivisionJP19:
		return JP
	case SubdivisionJP20:
		return JP
	case SubdivisionJP21:
		return JP
	case SubdivisionJP22:
		return JP
	case SubdivisionJP23:
		return JP
	case SubdivisionJP24:
		return JP
	case SubdivisionJP25:
		return JP
	case SubdivisionJP26:
		return JP
	case SubdivisionJP27:
		return JP
	case SubdivisionJP28:
		return JP
	case SubdivisionJP29:
		return JP
	case SubdivisionJP30:
		return JP
	case SubdivisionJP31:
		return JP
	case SubdivisionJP32:
		return JP
	case SubdivisionJP33:
		return JP
	case SubdivisionJP34:
		return JP
	case SubdivisionJP35:
		return JP
	case SubdivisionJP36:
		return JP
	case SubdivisionJP37:
		return JP
	case SubdivisionJP38:
		return JP
	case SubdivisionJP39:
		return JP
	case SubdivisionJP40:
		return JP
	case SubdivisionJP41:
		return JP
	case SubdivisionJP42:
		return JP
	case SubdivisionJP43:
		return JP
	case SubdivisionJP44:
		return JP
	case SubdivisionJP45:
		return JP
	case SubdivisionJP46:
		return JP
	case SubdivisionJP47:
		return JP
	case SubdivisionKE01:
		return KE
	case SubdivisionKE02:
		return KE
	case SubdivisionKE03:
		return KE
	case SubdivisionKE04:
		return KE
	case SubdivisionKE05:
		return KE
	case SubdivisionKE06:
		return KE
	case SubdivisionKE07:
		return KE
	case SubdivisionKE08:
		return KE
	case SubdivisionKE09:
		return KE
	case SubdivisionKE10:
		return KE
	case SubdivisionKE11:
		return KE
	case SubdivisionKE12:
		return KE
	case SubdivisionKE13:
		return KE
	case SubdivisionKE14:
		return KE
	case SubdivisionKE15:
		return KE
	case SubdivisionKE16:
		return KE
	case SubdivisionKE17:
		return KE
	case SubdivisionKE18:
		return KE
	case SubdivisionKE19:
		return KE
	case SubdivisionKE20:
		return KE
	case SubdivisionKE21:
		return KE
	case SubdivisionKE22:
		return KE
	case SubdivisionKE23:
		return KE
	case SubdivisionKE24:
		return KE
	case SubdivisionKE25:
		return KE
	case SubdivisionKE26:
		return KE
	case SubdivisionKE27:
		return KE
	case SubdivisionKE28:
		return KE
	case SubdivisionKE29:
		return KE
	case SubdivisionKE30:
		return KE
	case SubdivisionKE31:
		return KE
	case SubdivisionKE32:
		return KE
	case SubdivisionKE33:
		return KE
	case SubdivisionKE34:
		return KE
	case SubdivisionKE35:
		return KE
	case SubdivisionKE36:
		return KE
	case SubdivisionKE38:
		return KE
	case SubdivisionKE39:
		return KE
	case SubdivisionKE41:
		return KE
	case SubdivisionKE42:
		return KE
	case SubdivisionKE43:
		return KE
	case SubdivisionKE44:
		return KE
	case SubdivisionKE46:
		return KE
	case SubdivisionKGB:
		return KG
	case SubdivisionKGC:
		return KG
	case SubdivisionKGGB:
		return KG
	case SubdivisionKGGO:
		return KG
	case SubdivisionKGJ:
		return KG
	case SubdivisionKGN:
		return KG
	case SubdivisionKGT:
		return KG
	case SubdivisionKGY:
		return KG
	case SubdivisionKH10:
		return KH
	case SubdivisionKH11:
		return KH
	case SubdivisionKH12:
		return KH
	case SubdivisionKH14:
		return KH
	case SubdivisionKH15:
		return KH
	case SubdivisionKH17:
		return KH
	case SubdivisionKH18:
		return KH
	case SubdivisionKH19:
		return KH
	case SubdivisionKH1:
		return KH
	case SubdivisionKH20:
		return KH
	case SubdivisionKH21:
		return KH
	case SubdivisionKH23:
		return KH
	case SubdivisionKH24:
		return KH
	case SubdivisionKH2:
		return KH
	case SubdivisionKH3:
		return KH
	case SubdivisionKH4:
		return KH
	case SubdivisionKH5:
		return KH
	case SubdivisionKH6:
		return KH
	case SubdivisionKH7:
		return KH
	case SubdivisionKH8:
		return KH
	case SubdivisionKH9:
		return KH
	case SubdivisionKIG:
		return KI
	case SubdivisionKMG:
		return KM
	case SubdivisionKN02:
		return KN
	case SubdivisionKN03:
		return KN
	case SubdivisionKN05:
		return KN
	case SubdivisionKN06:
		return KN
	case SubdivisionKN07:
		return KN
	case SubdivisionKN08:
		return KN
	case SubdivisionKN09:
		return KN
	case SubdivisionKN10:
		return KN
	case SubdivisionKN11:
		return KN
	case SubdivisionKN12:
		return KN
	case SubdivisionKN13:
		return KN
	case SubdivisionKP01:
		return KP
	case SubdivisionKR11:
		return KR
	case SubdivisionKR26:
		return KR
	case SubdivisionKR27:
		return KR
	case SubdivisionKR28:
		return KR
	case SubdivisionKR29:
		return KR
	case SubdivisionKR30:
		return KR
	case SubdivisionKR31:
		return KR
	case SubdivisionKR41:
		return KR
	case SubdivisionKR42:
		return KR
	case SubdivisionKR43:
		return KR
	case SubdivisionKR44:
		return KR
	case SubdivisionKR45:
		return KR
	case SubdivisionKR46:
		return KR
	case SubdivisionKR47:
		return KR
	case SubdivisionKR48:
		return KR
	case SubdivisionKR49:
		return KR
	case SubdivisionKWAH:
		return KW
	case SubdivisionKWFA:
		return KW
	case SubdivisionKWHA:
		return KW
	case SubdivisionKWJA:
		return KW
	case SubdivisionKWKU:
		return KW
	case SubdivisionKWMU:
		return KW
	case SubdivisionKZAKM:
		return KZ
	case SubdivisionKZAKT:
		return KZ
	case SubdivisionKZALA:
		return KZ
	case SubdivisionKZALM:
		return KZ
	case SubdivisionKZAST:
		return KZ
	case SubdivisionKZATY:
		return KZ
	case SubdivisionKZKAR:
		return KZ
	case SubdivisionKZKUS:
		return KZ
	case SubdivisionKZKZY:
		return KZ
	case SubdivisionKZMAN:
		return KZ
	case SubdivisionKZPAV:
		return KZ
	case SubdivisionKZSEV:
		return KZ
	case SubdivisionKZSHY:
		return KZ
	case SubdivisionKZVOS:
		return KZ
	case SubdivisionKZYUZ:
		return KZ
	case SubdivisionKZZAP:
		return KZ
	case SubdivisionKZZHA:
		return KZ
	case SubdivisionLABL:
		return LA
	case SubdivisionLACH:
		return LA
	case SubdivisionLAKH:
		return LA
	case SubdivisionLALP:
		return LA
	case SubdivisionLAOU:
		return LA
	case SubdivisionLAPH:
		return LA
	case SubdivisionLASV:
		return LA
	case SubdivisionLAVI:
		return LA
	case SubdivisionLAXA:
		return LA
	case SubdivisionLAXE:
		return LA
	case SubdivisionLAXI:
		return LA
	case SubdivisionLBAK:
		return LB
	case SubdivisionLBAS:
		return LB
	case SubdivisionLBBA:
		return LB
	case SubdivisionLBBH:
		return LB
	case SubdivisionLBBI:
		return LB
	case SubdivisionLBJA:
		return LB
	case SubdivisionLBJL:
		return LB
	case SubdivisionLBNA:
		return LB
	case SubdivisionLC01:
		return LC
	case SubdivisionLC02:
		return LC
	case SubdivisionLC05:
		return LC
	case SubdivisionLC06:
		return LC
	case SubdivisionLC07:
		return LC
	case SubdivisionLC10:
		return LC
	case SubdivisionLC11:
		return LC
	case SubdivisionLI01:
		return LI
	case SubdivisionLI02:
		return LI
	case SubdivisionLI03:
		return LI
	case SubdivisionLI04:
		return LI
	case SubdivisionLI05:
		return LI
	case SubdivisionLI06:
		return LI
	case SubdivisionLI07:
		return LI
	case SubdivisionLI09:
		return LI
	case SubdivisionLI10:
		return LI
	case SubdivisionLI11:
		return LI
	case SubdivisionLK1:
		return LK
	case SubdivisionLK2:
		return LK
	case SubdivisionLK3:
		return LK
	case SubdivisionLK4:
		return LK
	case SubdivisionLK5:
		return LK
	case SubdivisionLK6:
		return LK
	case SubdivisionLK7:
		return LK
	case SubdivisionLK8:
		return LK
	case SubdivisionLK9:
		return LK
	case SubdivisionLRBM:
		return LR
	case SubdivisionLRGB:
		return LR
	case SubdivisionLRGG:
		return LR
	case SubdivisionLRMG:
		return LR
	case SubdivisionLRMO:
		return LR
	case SubdivisionLRNI:
		return LR
	case SubdivisionLRSI:
		return LR
	case SubdivisionLSA:
		return LS
	case SubdivisionLSC:
		return LS
	case SubdivisionLSE:
		return LS
	case SubdivisionLSG:
		return LS
	case SubdivisionLSJ:
		return LS
	case SubdivisionLSK:
		return LS
	case SubdivisionLTAL:
		return LT
	case SubdivisionLTKL:
		return LT
	case SubdivisionLTKU:
		return LT
	case SubdivisionLTMR:
		return LT
	case SubdivisionLTPN:
		return LT
	case SubdivisionLTSA:
		return LT
	case SubdivisionLTTA:
		return LT
	case SubdivisionLTTE:
		return LT
	case SubdivisionLTUT:
		return LT
	case SubdivisionLTVL:
		return LT
	case SubdivisionLUCA:
		return LU
	case SubdivisionLUCL:
		return LU
	case SubdivisionLUDI:
		return LU
	case SubdivisionLUEC:
		return LU
	case SubdivisionLUES:
		return LU
	case SubdivisionLUGR:
		return LU
	case SubdivisionLULU:
		return LU
	case SubdivisionLUME:
		return LU
	case SubdivisionLURD:
		return LU
	case SubdivisionLURM:
		return LU
	case SubdivisionLUVD:
		return LU
	case SubdivisionLUWI:
		return LU
	case SubdivisionLV001:
		return LV
	case SubdivisionLV002:
		return LV
	case SubdivisionLV003:
		return LV
	case SubdivisionLV005:
		return LV
	case SubdivisionLV007:
		return LV
	case SubdivisionLV010:
		return LV
	case SubdivisionLV011:
		return LV
	case SubdivisionLV012:
		return LV
	case SubdivisionLV015:
		return LV
	case SubdivisionLV016:
		return LV
	case SubdivisionLV017:
		return LV
	case SubdivisionLV018:
		return LV
	case SubdivisionLV020:
		return LV
	case SubdivisionLV021:
		return LV
	case SubdivisionLV022:
		return LV
	case SubdivisionLV024:
		return LV
	case SubdivisionLV025:
		return LV
	case SubdivisionLV026:
		return LV
	case SubdivisionLV027:
		return LV
	case SubdivisionLV029:
		return LV
	case SubdivisionLV030:
		return LV
	case SubdivisionLV033:
		return LV
	case SubdivisionLV034:
		return LV
	case SubdivisionLV035:
		return LV
	case SubdivisionLV037:
		return LV
	case SubdivisionLV038:
		return LV
	case SubdivisionLV039:
		return LV
	case SubdivisionLV040:
		return LV
	case SubdivisionLV041:
		return LV
	case SubdivisionLV042:
		return LV
	case SubdivisionLV043:
		return LV
	case SubdivisionLV046:
		return LV
	case SubdivisionLV047:
		return LV
	case SubdivisionLV050:
		return LV
	case SubdivisionLV052:
		return LV
	case SubdivisionLV053:
		return LV
	case SubdivisionLV054:
		return LV
	case SubdivisionLV056:
		return LV
	case SubdivisionLV057:
		return LV
	case SubdivisionLV058:
		return LV
	case SubdivisionLV059:
		return LV
	case SubdivisionLV061:
		return LV
	case SubdivisionLV064:
		return LV
	case SubdivisionLV067:
		return LV
	case SubdivisionLV068:
		return LV
	case SubdivisionLV069:
		return LV
	case SubdivisionLV073:
		return LV
	case SubdivisionLV075:
		return LV
	case SubdivisionLV077:
		return LV
	case SubdivisionLV078:
		return LV
	case SubdivisionLV079:
		return LV
	case SubdivisionLV083:
		return LV
	case SubdivisionLV086:
		return LV
	case SubdivisionLV087:
		return LV
	case SubdivisionLV088:
		return LV
	case SubdivisionLV089:
		return LV
	case SubdivisionLV090:
		return LV
	case SubdivisionLV091:
		return LV
	case SubdivisionLV093:
		return LV
	case SubdivisionLV094:
		return LV
	case SubdivisionLV095:
		return LV
	case SubdivisionLV097:
		return LV
	case SubdivisionLV099:
		return LV
	case SubdivisionLV100:
		return LV
	case SubdivisionLV101:
		return LV
	case SubdivisionLV103:
		return LV
	case SubdivisionLV105:
		return LV
	case SubdivisionLV106:
		return LV
	case SubdivisionLV110:
		return LV
	case SubdivisionLVJEL:
		return LV
	case SubdivisionLVJUR:
		return LV
	case SubdivisionLVLPX:
		return LV
	case SubdivisionLVRIX:
		return LV
	case SubdivisionLVVMR:
		return LV
	case SubdivisionLYBA:
		return LY
	case SubdivisionLYBU:
		return LY
	case SubdivisionLYDR:
		return LY
	case SubdivisionLYJA:
		return LY
	case SubdivisionLYJG:
		return LY
	case SubdivisionLYJI:
		return LY
	case SubdivisionLYJU:
		return LY
	case SubdivisionLYMB:
		return LY
	case SubdivisionLYMI:
		return LY
	case SubdivisionLYMJ:
		return LY
	case SubdivisionLYMQ:
		return LY
	case SubdivisionLYNL:
		return LY
	case SubdivisionLYNQ:
		return LY
	case SubdivisionLYSB:
		return LY
	case SubdivisionLYSR:
		return LY
	case SubdivisionLYTB:
		return LY
	case SubdivisionLYWA:
		return LY
	case SubdivisionLYZA:
		return LY
	case SubdivisionMA01:
		return MA
	case SubdivisionMA02:
		return MA
	case SubdivisionMA03:
		return MA
	case SubdivisionMA04:
		return MA
	case SubdivisionMA05:
		return MA
	case SubdivisionMA06:
		return MA
	case SubdivisionMA07:
		return MA
	case SubdivisionMA08:
		return MA
	case SubdivisionMA09:
		return MA
	case SubdivisionMA10:
		return MA
	case SubdivisionMA11:
		return MA
	case SubdivisionMCCO:
		return MC
	case SubdivisionMCFO:
		return MC
	case SubdivisionMCMC:
		return MC
	case SubdivisionMCMO:
		return MC
	case SubdivisionMCSR:
		return MC
	case SubdivisionMDAN:
		return MD
	case SubdivisionMDBA:
		return MD
	case SubdivisionMDBD:
		return MD
	case SubdivisionMDBR:
		return MD
	case SubdivisionMDBS:
		return MD
	case SubdivisionMDCA:
		return MD
	case SubdivisionMDCL:
		return MD
	case SubdivisionMDCM:
		return MD
	case SubdivisionMDCR:
		return MD
	case SubdivisionMDCS:
		return MD
	case SubdivisionMDCT:
		return MD
	case SubdivisionMDCU:
		return MD
	case SubdivisionMDDO:
		return MD
	case SubdivisionMDDR:
		return MD
	case SubdivisionMDDU:
		return MD
	case SubdivisionMDED:
		return MD
	case SubdivisionMDFA:
		return MD
	case SubdivisionMDFL:
		return MD
	case SubdivisionMDGA:
		return MD
	case SubdivisionMDGL:
		return MD
	case SubdivisionMDHI:
		return MD
	case SubdivisionMDIA:
		return MD
	case SubdivisionMDLE:
		return MD
	case SubdivisionMDNI:
		return MD
	case SubdivisionMDOC:
		return MD
	case SubdivisionMDOR:
		return MD
	case SubdivisionMDRE:
		return MD
	case SubdivisionMDRI:
		return MD
	case SubdivisionMDSD:
		return MD
	case SubdivisionMDSI:
		return MD
	case SubdivisionMDSN:
		return MD
	case SubdivisionMDSO:
		return MD
	case SubdivisionMDST:
		return MD
	case SubdivisionMDSV:
		return MD
	case SubdivisionMDTA:
		return MD
	case SubdivisionMDTE:
		return MD
	case SubdivisionMDUN:
		return MD
	case SubdivisionME02:
		return ME
	case SubdivisionME03:
		return ME
	case SubdivisionME04:
		return ME
	case SubdivisionME05:
		return ME
	case SubdivisionME06:
		return ME
	case SubdivisionME07:
		return ME
	case SubdivisionME08:
		return ME
	case SubdivisionME09:
		return ME
	case SubdivisionME10:
		return ME
	case SubdivisionME12:
		return ME
	case SubdivisionME13:
		return ME
	case SubdivisionME14:
		return ME
	case SubdivisionME15:
		return ME
	case SubdivisionME16:
		return ME
	case SubdivisionME17:
		return ME
	case SubdivisionME19:
		return ME
	case SubdivisionME20:
		return ME
	case SubdivisionME21:
		return ME
	case SubdivisionME24:
		return ME
	case SubdivisionMGA:
		return MG
	case SubdivisionMGD:
		return MG
	case SubdivisionMGF:
		return MG
	case SubdivisionMGM:
		return MG
	case SubdivisionMGT:
		return MG
	case SubdivisionMGU:
		return MG
	case SubdivisionMHKWA:
		return MH
	case SubdivisionMHMAJ:
		return MH
	case SubdivisionMK101:
		return MK
	case SubdivisionMK104:
		return MK
	case SubdivisionMK106:
		return MK
	case SubdivisionMK107:
		return MK
	case SubdivisionMK108:
		return MK
	case SubdivisionMK109:
		return MK
	case SubdivisionMK201:
		return MK
	case SubdivisionMK202:
		return MK
	case SubdivisionMK203:
		return MK
	case SubdivisionMK205:
		return MK
	case SubdivisionMK206:
		return MK
	case SubdivisionMK207:
		return MK
	case SubdivisionMK208:
		return MK
	case SubdivisionMK209:
		return MK
	case SubdivisionMK210:
		return MK
	case SubdivisionMK211:
		return MK
	case SubdivisionMK303:
		return MK
	case SubdivisionMK307:
		return MK
	case SubdivisionMK310:
		return MK
	case SubdivisionMK311:
		return MK
	case SubdivisionMK312:
		return MK
	case SubdivisionMK313:
		return MK
	case SubdivisionMK401:
		return MK
	case SubdivisionMK402:
		return MK
	case SubdivisionMK403:
		return MK
	case SubdivisionMK404:
		return MK
	case SubdivisionMK405:
		return MK
	case SubdivisionMK406:
		return MK
	case SubdivisionMK408:
		return MK
	case SubdivisionMK409:
		return MK
	case SubdivisionMK410:
		return MK
	case SubdivisionMK501:
		return MK
	case SubdivisionMK503:
		return MK
	case SubdivisionMK504:
		return MK
	case SubdivisionMK505:
		return MK
	case SubdivisionMK506:
		return MK
	case SubdivisionMK507:
		return MK
	case SubdivisionMK508:
		return MK
	case SubdivisionMK509:
		return MK
	case SubdivisionMK601:
		return MK
	case SubdivisionMK602:
		return MK
	case SubdivisionMK603:
		return MK
	case SubdivisionMK604:
		return MK
	case SubdivisionMK605:
		return MK
	case SubdivisionMK606:
		return MK
	case SubdivisionMK607:
		return MK
	case SubdivisionMK608:
		return MK
	case SubdivisionMK609:
		return MK
	case SubdivisionMK701:
		return MK
	case SubdivisionMK702:
		return MK
	case SubdivisionMK703:
		return MK
	case SubdivisionMK704:
		return MK
	case SubdivisionMK705:
		return MK
	case SubdivisionMK802:
		return MK
	case SubdivisionMK803:
		return MK
	case SubdivisionMK804:
		return MK
	case SubdivisionMK806:
		return MK
	case SubdivisionMK807:
		return MK
	case SubdivisionMK809:
		return MK
	case SubdivisionMK810:
		return MK
	case SubdivisionMK811:
		return MK
	case SubdivisionMK812:
		return MK
	case SubdivisionMK813:
		return MK
	case SubdivisionMK814:
		return MK
	case SubdivisionMK817:
		return MK
	case SubdivisionML1:
		return ML
	case SubdivisionML2:
		return ML
	case SubdivisionML3:
		return ML
	case SubdivisionML4:
		return ML
	case SubdivisionML5:
		return ML
	case SubdivisionML6:
		return ML
	case SubdivisionML7:
		return ML
	case SubdivisionML8:
		return ML
	case SubdivisionMLBKO:
		return ML
	case SubdivisionMM01:
		return MM
	case SubdivisionMM02:
		return MM
	case SubdivisionMM03:
		return MM
	case SubdivisionMM04:
		return MM
	case SubdivisionMM05:
		return MM
	case SubdivisionMM06:
		return MM
	case SubdivisionMM07:
		return MM
	case SubdivisionMM11:
		return MM
	case SubdivisionMM12:
		return MM
	case SubdivisionMM13:
		return MM
	case SubdivisionMM15:
		return MM
	case SubdivisionMM16:
		return MM
	case SubdivisionMM17:
		return MM
	case SubdivisionMM18:
		return MM
	case SubdivisionMN035:
		return MN
	case SubdivisionMN037:
		return MN
	case SubdivisionMN047:
		return MN
	case SubdivisionMN049:
		return MN
	case SubdivisionMN055:
		return MN
	case SubdivisionMN061:
		return MN
	case SubdivisionMN065:
		return MN
	case SubdivisionMN071:
		return MN
	case SubdivisionMN1:
		return MN
	case SubdivisionMR04:
		return MR
	case SubdivisionMR06:
		return MR
	case SubdivisionMR08:
		return MR
	case SubdivisionMR11:
		return MR
	case SubdivisionMR12:
		return MR
	case SubdivisionMR13:
		return MR
	case SubdivisionMT01:
		return MT
	case SubdivisionMT02:
		return MT
	case SubdivisionMT03:
		return MT
	case SubdivisionMT04:
		return MT
	case SubdivisionMT05:
		return MT
	case SubdivisionMT06:
		return MT
	case SubdivisionMT07:
		return MT
	case SubdivisionMT08:
		return MT
	case SubdivisionMT09:
		return MT
	case SubdivisionMT10:
		return MT
	case SubdivisionMT11:
		return MT
	case SubdivisionMT12:
		return MT
	case SubdivisionMT14:
		return MT
	case SubdivisionMT15:
		return MT
	case SubdivisionMT16:
		return MT
	case SubdivisionMT17:
		return MT
	case SubdivisionMT18:
		return MT
	case SubdivisionMT19:
		return MT
	case SubdivisionMT20:
		return MT
	case SubdivisionMT21:
		return MT
	case SubdivisionMT22:
		return MT
	case SubdivisionMT23:
		return MT
	case SubdivisionMT24:
		return MT
	case SubdivisionMT25:
		return MT
	case SubdivisionMT26:
		return MT
	case SubdivisionMT27:
		return MT
	case SubdivisionMT28:
		return MT
	case SubdivisionMT29:
		return MT
	case SubdivisionMT30:
		return MT
	case SubdivisionMT31:
		return MT
	case SubdivisionMT32:
		return MT
	case SubdivisionMT33:
		return MT
	case SubdivisionMT34:
		return MT
	case SubdivisionMT35:
		return MT
	case SubdivisionMT36:
		return MT
	case SubdivisionMT37:
		return MT
	case SubdivisionMT38:
		return MT
	case SubdivisionMT39:
		return MT
	case SubdivisionMT40:
		return MT
	case SubdivisionMT41:
		return MT
	case SubdivisionMT42:
		return MT
	case SubdivisionMT43:
		return MT
	case SubdivisionMT45:
		return MT
	case SubdivisionMT46:
		return MT
	case SubdivisionMT48:
		return MT
	case SubdivisionMT49:
		return MT
	case SubdivisionMT51:
		return MT
	case SubdivisionMT52:
		return MT
	case SubdivisionMT53:
		return MT
	case SubdivisionMT54:
		return MT
	case SubdivisionMT55:
		return MT
	case SubdivisionMT56:
		return MT
	case SubdivisionMT57:
		return MT
	case SubdivisionMT58:
		return MT
	case SubdivisionMT59:
		return MT
	case SubdivisionMT60:
		return MT
	case SubdivisionMT61:
		return MT
	case SubdivisionMT62:
		return MT
	case SubdivisionMT63:
		return MT
	case SubdivisionMT64:
		return MT
	case SubdivisionMT65:
		return MT
	case SubdivisionMT67:
		return MT
	case SubdivisionMT68:
		return MT
	case SubdivisionMUBL:
		return MU
	case SubdivisionMUFL:
		return MU
	case SubdivisionMUGP:
		return MU
	case SubdivisionMUMO:
		return MU
	case SubdivisionMUPA:
		return MU
	case SubdivisionMUPL:
		return MU
	case SubdivisionMUPW:
		return MU
	case SubdivisionMURO:
		return MU
	case SubdivisionMURR:
		return MU
	case SubdivisionMUSA:
		return MU
	case SubdivisionMV00:
		return MV
	case SubdivisionMV01:
		return MV
	case SubdivisionMV03:
		return MV
	case SubdivisionMV04:
		return MV
	case SubdivisionMV05:
		return MV
	case SubdivisionMV07:
		return MV
	case SubdivisionMV12:
		return MV
	case SubdivisionMV13:
		return MV
	case SubdivisionMV17:
		return MV
	case SubdivisionMV20:
		return MV
	case SubdivisionMV25:
		return MV
	case SubdivisionMV28:
		return MV
	case SubdivisionMVMLE:
		return MV
	case SubdivisionMWBA:
		return MW
	case SubdivisionMWBL:
		return MW
	case SubdivisionMWCR:
		return MW
	case SubdivisionMWDE:
		return MW
	case SubdivisionMWDO:
		return MW
	case SubdivisionMWKR:
		return MW
	case SubdivisionMWLI:
		return MW
	case SubdivisionMWMG:
		return MW
	case SubdivisionMWMH:
		return MW
	case SubdivisionMWMZ:
		return MW
	case SubdivisionMWNK:
		return MW
	case SubdivisionMWSA:
		return MW
	case SubdivisionMWZO:
		return MW
	case SubdivisionMXAGU:
		return MX
	case SubdivisionMXBCN:
		return MX
	case SubdivisionMXBCS:
		return MX
	case SubdivisionMXCAM:
		return MX
	case SubdivisionMXCHH:
		return MX
	case SubdivisionMXCHP:
		return MX
	case SubdivisionMXCMX:
		return MX
	case SubdivisionMXCOA:
		return MX
	case SubdivisionMXCOL:
		return MX
	case SubdivisionMXDUR:
		return MX
	case SubdivisionMXGRO:
		return MX
	case SubdivisionMXGUA:
		return MX
	case SubdivisionMXHID:
		return MX
	case SubdivisionMXJAL:
		return MX
	case SubdivisionMXMEX:
		return MX
	case SubdivisionMXMIC:
		return MX
	case SubdivisionMXMOR:
		return MX
	case SubdivisionMXNAY:
		return MX
	case SubdivisionMXNLE:
		return MX
	case SubdivisionMXOAX:
		return MX
	case SubdivisionMXPUE:
		return MX
	case SubdivisionMXQUE:
		return MX
	case SubdivisionMXROO:
		return MX
	case SubdivisionMXSIN:
		return MX
	case SubdivisionMXSLP:
		return MX
	case SubdivisionMXSON:
		return MX
	case SubdivisionMXTAB:
		return MX
	case SubdivisionMXTAM:
		return MX
	case SubdivisionMXTLA:
		return MX
	case SubdivisionMXVER:
		return MX
	case SubdivisionMXYUC:
		return MX
	case SubdivisionMXZAC:
		return MX
	case SubdivisionMY01:
		return MY
	case SubdivisionMY02:
		return MY
	case SubdivisionMY03:
		return MY
	case SubdivisionMY04:
		return MY
	case SubdivisionMY05:
		return MY
	case SubdivisionMY06:
		return MY
	case SubdivisionMY07:
		return MY
	case SubdivisionMY08:
		return MY
	case SubdivisionMY09:
		return MY
	case SubdivisionMY10:
		return MY
	case SubdivisionMY11:
		return MY
	case SubdivisionMY12:
		return MY
	case SubdivisionMY13:
		return MY
	case SubdivisionMY14:
		return MY
	case SubdivisionMY15:
		return MY
	case SubdivisionMY16:
		return MY
	case SubdivisionMZA:
		return MZ
	case SubdivisionMZB:
		return MZ
	case SubdivisionMZG:
		return MZ
	case SubdivisionMZI:
		return MZ
	case SubdivisionMZL:
		return MZ
	case SubdivisionMZN:
		return MZ
	case SubdivisionMZP:
		return MZ
	case SubdivisionMZQ:
		return MZ
	case SubdivisionMZS:
		return MZ
	case SubdivisionMZT:
		return MZ
	case SubdivisionNACA:
		return NA
	case SubdivisionNAER:
		return NA
	case SubdivisionNAHA:
		return NA
	case SubdivisionNAKA:
		return NA
	case SubdivisionNAKE:
		return NA
	case SubdivisionNAKH:
		return NA
	case SubdivisionNAKU:
		return NA
	case SubdivisionNAOD:
		return NA
	case SubdivisionNAOH:
		return NA
	case SubdivisionNAON:
		return NA
	case SubdivisionNAOS:
		return NA
	case SubdivisionNAOT:
		return NA
	case SubdivisionNAOW:
		return NA
	case SubdivisionNE1:
		return NE
	case SubdivisionNE3:
		return NE
	case SubdivisionNE5:
		return NE
	case SubdivisionNE7:
		return NE
	case SubdivisionNE8:
		return NE
	case SubdivisionNGAB:
		return NG
	case SubdivisionNGAD:
		return NG
	case SubdivisionNGAK:
		return NG
	case SubdivisionNGAN:
		return NG
	case SubdivisionNGBA:
		return NG
	case SubdivisionNGBE:
		return NG
	case SubdivisionNGBO:
		return NG
	case SubdivisionNGBY:
		return NG
	case SubdivisionNGCR:
		return NG
	case SubdivisionNGDE:
		return NG
	case SubdivisionNGED:
		return NG
	case SubdivisionNGEK:
		return NG
	case SubdivisionNGEN:
		return NG
	case SubdivisionNGFC:
		return NG
	case SubdivisionNGGO:
		return NG
	case SubdivisionNGIM:
		return NG
	case SubdivisionNGJI:
		return NG
	case SubdivisionNGKD:
		return NG
	case SubdivisionNGKE:
		return NG
	case SubdivisionNGKN:
		return NG
	case SubdivisionNGKO:
		return NG
	case SubdivisionNGKT:
		return NG
	case SubdivisionNGKW:
		return NG
	case SubdivisionNGLA:
		return NG
	case SubdivisionNGNA:
		return NG
	case SubdivisionNGNI:
		return NG
	case SubdivisionNGOG:
		return NG
	case SubdivisionNGON:
		return NG
	case SubdivisionNGOS:
		return NG
	case SubdivisionNGOY:
		return NG
	case SubdivisionNGPL:
		return NG
	case SubdivisionNGRI:
		return NG
	case SubdivisionNGSO:
		return NG
	case SubdivisionNGTA:
		return NG
	case SubdivisionNGYO:
		return NG
	case SubdivisionNGZA:
		return NG
	case SubdivisionNIAN:
		return NI
	case SubdivisionNIAS:
		return NI
	case SubdivisionNIBO:
		return NI
	case SubdivisionNICA:
		return NI
	case SubdivisionNICI:
		return NI
	case SubdivisionNICO:
		return NI
	case SubdivisionNIES:
		return NI
	case SubdivisionNIGR:
		return NI
	case SubdivisionNIJI:
		return NI
	case SubdivisionNILE:
		return NI
	case SubdivisionNIMD:
		return NI
	case SubdivisionNIMN:
		return NI
	case SubdivisionNIMS:
		return NI
	case SubdivisionNIMT:
		return NI
	case SubdivisionNINS:
		return NI
	case SubdivisionNIRI:
		return NI
	case SubdivisionNISJ:
		return NI
	case SubdivisionNLDR:
		return NL
	case SubdivisionNLFL:
		return NL
	case SubdivisionNLFR:
		return NL
	case SubdivisionNLGE:
		return NL
	case SubdivisionNLGR:
		return NL
	case SubdivisionNLLI:
		return NL
	case SubdivisionNLNB:
		return NL
	case SubdivisionNLNH:
		return NL
	case SubdivisionNLOV:
		return NL
	case SubdivisionNLUT:
		return NL
	case SubdivisionNLZE:
		return NL
	case SubdivisionNLZH:
		return NL
	case SubdivisionNO03:
		return NO
	case SubdivisionNO11:
		return NO
	case SubdivisionNO15:
		return NO
	case SubdivisionNO18:
		return NO
	case SubdivisionNO30:
		return NO
	case SubdivisionNO34:
		return NO
	case SubdivisionNO38:
		return NO
	case SubdivisionNO42:
		return NO
	case SubdivisionNO46:
		return NO
	case SubdivisionNO50:
		return NO
	case SubdivisionNO54:
		return NO
	case SubdivisionNPBA:
		return NP
	case SubdivisionNPBH:
		return NP
	case SubdivisionNPDH:
		return NP
	case SubdivisionNPGA:
		return NP
	case SubdivisionNPJA:
		return NP
	case SubdivisionNPKA:
		return NP
	case SubdivisionNPKO:
		return NP
	case SubdivisionNPLU:
		return NP
	case SubdivisionNPMA:
		return NP
	case SubdivisionNPME:
		return NP
	case SubdivisionNPNA:
		return NP
	case SubdivisionNPRA:
		return NP
	case SubdivisionNPSA:
		return NP
	case SubdivisionNPSE:
		return NP
	case SubdivisionNR14:
		return NR
	case SubdivisionNZAUK:
		return NZ
	case SubdivisionNZBOP:
		return NZ
	case SubdivisionNZCAN:
		return NZ
	case SubdivisionNZCIT:
		return NZ
	case SubdivisionNZGIS:
		return NZ
	case SubdivisionNZHKB:
		return NZ
	case SubdivisionNZMBH:
		return NZ
	case SubdivisionNZMWT:
		return NZ
	case SubdivisionNZNSN:
		return NZ
	case SubdivisionNZNTL:
		return NZ
	case SubdivisionNZOTA:
		return NZ
	case SubdivisionNZSTL:
		return NZ
	case SubdivisionNZTAS:
		return NZ
	case SubdivisionNZTKI:
		return NZ
	case SubdivisionNZWGN:
		return NZ
	case SubdivisionNZWKO:
		return NZ
	case SubdivisionNZWTC:
		return NZ
	case SubdivisionOMBJ:
		return OM
	case SubdivisionOMBS:
		return OM
	case SubdivisionOMBU:
		return OM
	case SubdivisionOMDA:
		return OM
	case SubdivisionOMMA:
		return OM
	case SubdivisionOMMU:
		return OM
	case SubdivisionOMSJ:
		return OM
	case SubdivisionOMSS:
		return OM
	case SubdivisionOMWU:
		return OM
	case SubdivisionOMZA:
		return OM
	case SubdivisionOMZU:
		return OM
	case SubdivisionPA1:
		return PA
	case SubdivisionPA2:
		return PA
	case SubdivisionPA3:
		return PA
	case SubdivisionPA4:
		return PA
	case SubdivisionPA5:
		return PA
	case SubdivisionPA6:
		return PA
	case SubdivisionPA7:
		return PA
	case SubdivisionPA8:
		return PA
	case SubdivisionPA9:
		return PA
	case SubdivisionPANB:
		return PA
	case SubdivisionPEAMA:
		return PE
	case SubdivisionPEANC:
		return PE
	case SubdivisionPEAPU:
		return PE
	case SubdivisionPEARE:
		return PE
	case SubdivisionPEAYA:
		return PE
	case SubdivisionPECAJ:
		return PE
	case SubdivisionPECAL:
		return PE
	case SubdivisionPECUS:
		return PE
	case SubdivisionPEHUC:
		return PE
	case SubdivisionPEHUV:
		return PE
	case SubdivisionPEICA:
		return PE
	case SubdivisionPEJUN:
		return PE
	case SubdivisionPELAL:
		return PE
	case SubdivisionPELAM:
		return PE
	case SubdivisionPELIM:
		return PE
	case SubdivisionPELOR:
		return PE
	case SubdivisionPEMDD:
		return PE
	case SubdivisionPEMOQ:
		return PE
	case SubdivisionPEPAS:
		return PE
	case SubdivisionPEPIU:
		return PE
	case SubdivisionPEPUN:
		return PE
	case SubdivisionPESAM:
		return PE
	case SubdivisionPETAC:
		return PE
	case SubdivisionPETUM:
		return PE
	case SubdivisionPEUCA:
		return PE
	case SubdivisionPGCPM:
		return PG
	case SubdivisionPGEBR:
		return PG
	case SubdivisionPGEHG:
		return PG
	case SubdivisionPGESW:
		return PG
	case SubdivisionPGMBA:
		return PG
	case SubdivisionPGMPL:
		return PG
	case SubdivisionPGMPM:
		return PG
	case SubdivisionPGMRL:
		return PG
	case SubdivisionPGNCD:
		return PG
	case SubdivisionPGNIK:
		return PG
	case SubdivisionPGNSB:
		return PG
	case SubdivisionPGSAN:
		return PG
	case SubdivisionPGSHM:
		return PG
	case SubdivisionPGWBK:
		return PG
	case SubdivisionPGWHM:
		return PG
	case SubdivisionPGWPD:
		return PG
	case SubdivisionPH00:
		return PH
	case SubdivisionPHABR:
		return PH
	case SubdivisionPHAGN:
		return PH
	case SubdivisionPHAGS:
		return PH
	case SubdivisionPHAKL:
		return PH
	case SubdivisionPHALB:
		return PH
	case SubdivisionPHANT:
		return PH
	case SubdivisionPHAPA:
		return PH
	case SubdivisionPHAUR:
		return PH
	case SubdivisionPHBAN:
		return PH
	case SubdivisionPHBAS:
		return PH
	case SubdivisionPHBEN:
		return PH
	case SubdivisionPHBIL:
		return PH
	case SubdivisionPHBOH:
		return PH
	case SubdivisionPHBTG:
		return PH
	case SubdivisionPHBTN:
		return PH
	case SubdivisionPHBUK:
		return PH
	case SubdivisionPHBUL:
		return PH
	case SubdivisionPHCAG:
		return PH
	case SubdivisionPHCAM:
		return PH
	case SubdivisionPHCAN:
		return PH
	case SubdivisionPHCAP:
		return PH
	case SubdivisionPHCAS:
		return PH
	case SubdivisionPHCAT:
		return PH
	case SubdivisionPHCAV:
		return PH
	case SubdivisionPHCEB:
		return PH
	case SubdivisionPHCOM:
		return PH
	case SubdivisionPHDAO:
		return PH
	case SubdivisionPHDAS:
		return PH
	case SubdivisionPHDAV:
		return PH
	case SubdivisionPHDIN:
		return PH
	case SubdivisionPHEAS:
		return PH
	case SubdivisionPHGUI:
		return PH
	case SubdivisionPHIFU:
		return PH
	case SubdivisionPHILI:
		return PH
	case SubdivisionPHILN:
		return PH
	case SubdivisionPHILS:
		return PH
	case SubdivisionPHISA:
		return PH
	case SubdivisionPHKAL:
		return PH
	case SubdivisionPHLAG:
		return PH
	case SubdivisionPHLAN:
		return PH
	case SubdivisionPHLAS:
		return PH
	case SubdivisionPHLEY:
		return PH
	case SubdivisionPHLUN:
		return PH
	case SubdivisionPHMAD:
		return PH
	case SubdivisionPHMAG:
		return PH
	case SubdivisionPHMAS:
		return PH
	case SubdivisionPHMDC:
		return PH
	case SubdivisionPHMDR:
		return PH
	case SubdivisionPHMOU:
		return PH
	case SubdivisionPHMSC:
		return PH
	case SubdivisionPHMSR:
		return PH
	case SubdivisionPHNCO:
		return PH
	case SubdivisionPHNEC:
		return PH
	case SubdivisionPHNER:
		return PH
	case SubdivisionPHNSA:
		return PH
	case SubdivisionPHNUE:
		return PH
	case SubdivisionPHNUV:
		return PH
	case SubdivisionPHPAM:
		return PH
	case SubdivisionPHPAN:
		return PH
	case SubdivisionPHPLW:
		return PH
	case SubdivisionPHQUE:
		return PH
	case SubdivisionPHQUI:
		return PH
	case SubdivisionPHRIZ:
		return PH
	case SubdivisionPHROM:
		return PH
	case SubdivisionPHSAR:
		return PH
	case SubdivisionPHSCO:
		return PH
	case SubdivisionPHSIG:
		return PH
	case SubdivisionPHSLE:
		return PH
	case SubdivisionPHSLU:
		return PH
	case SubdivisionPHSOR:
		return PH
	case SubdivisionPHSUK:
		return PH
	case SubdivisionPHSUN:
		return PH
	case SubdivisionPHSUR:
		return PH
	case SubdivisionPHTAR:
		return PH
	case SubdivisionPHTAW:
		return PH
	case SubdivisionPHWSA:
		return PH
	case SubdivisionPHZAN:
		return PH
	case SubdivisionPHZAS:
		return PH
	case SubdivisionPHZMB:
		return PH
	case SubdivisionPHZSI:
		return PH
	case SubdivisionPKBA:
		return PK
	case SubdivisionPKGB:
		return PK
	case SubdivisionPKIS:
		return PK
	case SubdivisionPKJK:
		return PK
	case SubdivisionPKKP:
		return PK
	case SubdivisionPKPB:
		return PK
	case SubdivisionPKSD:
		return PK
	case SubdivisionPL02:
		return PL
	case SubdivisionPL04:
		return PL
	case SubdivisionPL06:
		return PL
	case SubdivisionPL08:
		return PL
	case SubdivisionPL10:
		return PL
	case SubdivisionPL12:
		return PL
	case SubdivisionPL14:
		return PL
	case SubdivisionPL16:
		return PL
	case SubdivisionPL18:
		return PL
	case SubdivisionPL20:
		return PL
	case SubdivisionPL22:
		return PL
	case SubdivisionPL24:
		return PL
	case SubdivisionPL26:
		return PL
	case SubdivisionPL28:
		return PL
	case SubdivisionPL30:
		return PL
	case SubdivisionPL32:
		return PL
	case SubdivisionPSBTH:
		return PS
	case SubdivisionPSDEB:
		return PS
	case SubdivisionPSGZA:
		return PS
	case SubdivisionPSHBN:
		return PS
	case SubdivisionPSJEM:
		return PS
	case SubdivisionPSJEN:
		return PS
	case SubdivisionPSJRH:
		return PS
	case SubdivisionPSKYS:
		return PS
	case SubdivisionPSNBS:
		return PS
	case SubdivisionPSQQA:
		return PS
	case SubdivisionPSRBH:
		return PS
	case SubdivisionPSRFH:
		return PS
	case SubdivisionPSSLT:
		return PS
	case SubdivisionPSTBS:
		return PS
	case SubdivisionPSTKM:
		return PS
	case SubdivisionPT01:
		return PT
	case SubdivisionPT02:
		return PT
	case SubdivisionPT03:
		return PT
	case SubdivisionPT04:
		return PT
	case SubdivisionPT05:
		return PT
	case SubdivisionPT06:
		return PT
	case SubdivisionPT07:
		return PT
	case SubdivisionPT08:
		return PT
	case SubdivisionPT09:
		return PT
	case SubdivisionPT10:
		return PT
	case SubdivisionPT11:
		return PT
	case SubdivisionPT12:
		return PT
	case SubdivisionPT13:
		return PT
	case SubdivisionPT14:
		return PT
	case SubdivisionPT15:
		return PT
	case SubdivisionPT16:
		return PT
	case SubdivisionPT17:
		return PT
	case SubdivisionPT18:
		return PT
	case SubdivisionPT20:
		return PT
	case SubdivisionPT30:
		return PT
	case SubdivisionPW004:
		return PW
	case SubdivisionPW100:
		return PW
	case SubdivisionPW150:
		return PW
	case SubdivisionPW212:
		return PW
	case SubdivisionPW214:
		return PW
	case SubdivisionPW222:
		return PW
	case SubdivisionPY10:
		return PY
	case SubdivisionPY11:
		return PY
	case SubdivisionPY12:
		return PY
	case SubdivisionPY13:
		return PY
	case SubdivisionPY14:
		return PY
	case SubdivisionPY15:
		return PY
	case SubdivisionPY16:
		return PY
	case SubdivisionPY19:
		return PY
	case SubdivisionPY1:
		return PY
	case SubdivisionPY2:
		return PY
	case SubdivisionPY3:
		return PY
	case SubdivisionPY4:
		return PY
	case SubdivisionPY5:
		return PY
	case SubdivisionPY6:
		return PY
	case SubdivisionPY7:
		return PY
	case SubdivisionPY8:
		return PY
	case SubdivisionPY9:
		return PY
	case SubdivisionPYASU:
		return PY
	case SubdivisionQADA:
		return QA
	case SubdivisionQAKH:
		return QA
	case SubdivisionQAMS:
		return QA
	case SubdivisionQARA:
		return QA
	case SubdivisionQAUS:
		return QA
	case SubdivisionQAWA:
		return QA
	case SubdivisionQAZA:
		return QA
	case SubdivisionROAB:
		return RO
	case SubdivisionROAG:
		return RO
	case SubdivisionROAR:
		return RO
	case SubdivisionROB:
		return RO
	case SubdivisionROBC:
		return RO
	case SubdivisionROBH:
		return RO
	case SubdivisionROBN:
		return RO
	case SubdivisionROBR:
		return RO
	case SubdivisionROBT:
		return RO
	case SubdivisionROBV:
		return RO
	case SubdivisionROBZ:
		return RO
	case SubdivisionROCJ:
		return RO
	case SubdivisionROCL:
		return RO
	case SubdivisionROCS:
		return RO
	case SubdivisionROCT:
		return RO
	case SubdivisionROCV:
		return RO
	case SubdivisionRODB:
		return RO
	case SubdivisionRODJ:
		return RO
	case SubdivisionROGJ:
		return RO
	case SubdivisionROGL:
		return RO
	case SubdivisionROGR:
		return RO
	case SubdivisionROHD:
		return RO
	case SubdivisionROHR:
		return RO
	case SubdivisionROIF:
		return RO
	case SubdivisionROIL:
		return RO
	case SubdivisionROIS:
		return RO
	case SubdivisionROMH:
		return RO
	case SubdivisionROMM:
		return RO
	case SubdivisionROMS:
		return RO
	case SubdivisionRONT:
		return RO
	case SubdivisionROOT:
		return RO
	case SubdivisionROPH:
		return RO
	case SubdivisionROSB:
		return RO
	case SubdivisionROSJ:
		return RO
	case SubdivisionROSM:
		return RO
	case SubdivisionROSV:
		return RO
	case SubdivisionROTL:
		return RO
	case SubdivisionROTM:
		return RO
	case SubdivisionROTR:
		return RO
	case SubdivisionROVL:
		return RO
	case SubdivisionROVN:
		return RO
	case SubdivisionROVS:
		return RO
	case SubdivisionRS00:
		return RS
	case SubdivisionRS01:
		return RS
	case SubdivisionRS02:
		return RS
	case SubdivisionRS03:
		return RS
	case SubdivisionRS04:
		return RS
	case SubdivisionRS05:
		return RS
	case SubdivisionRS06:
		return RS
	case SubdivisionRS07:
		return RS
	case SubdivisionRS08:
		return RS
	case SubdivisionRS09:
		return RS
	case SubdivisionRS10:
		return RS
	case SubdivisionRS11:
		return RS
	case SubdivisionRS12:
		return RS
	case SubdivisionRS13:
		return RS
	case SubdivisionRS14:
		return RS
	case SubdivisionRS15:
		return RS
	case SubdivisionRS16:
		return RS
	case SubdivisionRS17:
		return RS
	case SubdivisionRS18:
		return RS
	case SubdivisionRS19:
		return RS
	case SubdivisionRS20:
		return RS
	case SubdivisionRS21:
		return RS
	case SubdivisionRS22:
		return RS
	case SubdivisionRS23:
		return RS
	case SubdivisionRS24:
		return RS
	case SubdivisionRS26:
		return RS
	case SubdivisionRS27:
		return RS
	case SubdivisionRS28:
		return RS
	case SubdivisionRUAD:
		return RU
	case SubdivisionRUAL:
		return RU
	case SubdivisionRUALT:
		return RU
	case SubdivisionRUAMU:
		return RU
	case SubdivisionRUARK:
		return RU
	case SubdivisionRUAST:
		return RU
	case SubdivisionRUBA:
		return RU
	case SubdivisionRUBEL:
		return RU
	case SubdivisionRUBRY:
		return RU
	case SubdivisionRUBU:
		return RU
	case SubdivisionRUCE:
		return RU
	case SubdivisionRUCHE:
		return RU
	case SubdivisionRUCHU:
		return RU
	case SubdivisionRUCU:
		return RU
	case SubdivisionRUDA:
		return RU
	case SubdivisionRUIN:
		return RU
	case SubdivisionRUIRK:
		return RU
	case SubdivisionRUIVA:
		return RU
	case SubdivisionRUKAM:
		return RU
	case SubdivisionRUKB:
		return RU
	case SubdivisionRUKC:
		return RU
	case SubdivisionRUKDA:
		return RU
	case SubdivisionRUKEM:
		return RU
	case SubdivisionRUKGD:
		return RU
	case SubdivisionRUKGN:
		return RU
	case SubdivisionRUKHA:
		return RU
	case SubdivisionRUKHM:
		return RU
	case SubdivisionRUKIR:
		return RU
	case SubdivisionRUKK:
		return RU
	case SubdivisionRUKL:
		return RU
	case SubdivisionRUKLU:
		return RU
	case SubdivisionRUKO:
		return RU
	case SubdivisionRUKOS:
		return RU
	case SubdivisionRUKR:
		return RU
	case SubdivisionRUKRS:
		return RU
	case SubdivisionRUKYA:
		return RU
	case SubdivisionRULEN:
		return RU
	case SubdivisionRULIP:
		return RU
	case SubdivisionRUMAG:
		return RU
	case SubdivisionRUME:
		return RU
	case SubdivisionRUMO:
		return RU
	case SubdivisionRUMOS:
		return RU
	case SubdivisionRUMOW:
		return RU
	case SubdivisionRUMUR:
		return RU
	case SubdivisionRUNEN:
		return RU
	case SubdivisionRUNGR:
		return RU
	case SubdivisionRUNIZ:
		return RU
	case SubdivisionRUNVS:
		return RU
	case SubdivisionRUOMS:
		return RU
	case SubdivisionRUORE:
		return RU
	case SubdivisionRUORL:
		return RU
	case SubdivisionRUPER:
		return RU
	case SubdivisionRUPNZ:
		return RU
	case SubdivisionRUPRI:
		return RU
	case SubdivisionRUPSK:
		return RU
	case SubdivisionRUROS:
		return RU
	case SubdivisionRURYA:
		return RU
	case SubdivisionRUSA:
		return RU
	case SubdivisionRUSAK:
		return RU
	case SubdivisionRUSAM:
		return RU
	case SubdivisionRUSAR:
		return RU
	case SubdivisionRUSE:
		return RU
	case SubdivisionRUSMO:
		return RU
	case SubdivisionRUSPE:
		return RU
	case SubdivisionRUSTA:
		return RU
	case SubdivisionRUSVE:
		return RU
	case SubdivisionRUTA:
		return RU
	case SubdivisionRUTAM:
		return RU
	case SubdivisionRUTOM:
		return RU
	case SubdivisionRUTUL:
		return RU
	case SubdivisionRUTVE:
		return RU
	case SubdivisionRUTY:
		return RU
	case SubdivisionRUTYU:
		return RU
	case SubdivisionRUUD:
		return RU
	case SubdivisionRUULY:
		return RU
	case SubdivisionRUVGG:
		return RU
	case SubdivisionRUVLA:
		return RU
	case SubdivisionRUVLG:
		return RU
	case SubdivisionRUVOR:
		return RU
	case SubdivisionRUYAN:
		return RU
	case SubdivisionRUYAR:
		return RU
	case SubdivisionRUYEV:
		return RU
	case SubdivisionRUZAB:
		return RU
	case SubdivisionRW01:
		return RW
	case SubdivisionRW02:
		return RW
	case SubdivisionRW03:
		return RW
	case SubdivisionRW04:
		return RW
	case SubdivisionRW05:
		return RW
	case SubdivisionSA01:
		return SA
	case SubdivisionSA02:
		return SA
	case SubdivisionSA03:
		return SA
	case SubdivisionSA04:
		return SA
	case SubdivisionSA05:
		return SA
	case SubdivisionSA06:
		return SA
	case SubdivisionSA07:
		return SA
	case SubdivisionSA08:
		return SA
	case SubdivisionSA09:
		return SA
	case SubdivisionSA10:
		return SA
	case SubdivisionSA11:
		return SA
	case SubdivisionSA12:
		return SA
	case SubdivisionSA14:
		return SA
	case SubdivisionSBCH:
		return SB
	case SubdivisionSBGU:
		return SB
	case SubdivisionSBWE:
		return SB
	case SubdivisionSC01:
		return SC
	case SubdivisionSC02:
		return SC
	case SubdivisionSC06:
		return SC
	case SubdivisionSC07:
		return SC
	case SubdivisionSC08:
		return SC
	case SubdivisionSC10:
		return SC
	case SubdivisionSC11:
		return SC
	case SubdivisionSC13:
		return SC
	case SubdivisionSC14:
		return SC
	case SubdivisionSC15:
		return SC
	case SubdivisionSC16:
		return SC
	case SubdivisionSC20:
		return SC
	case SubdivisionSC23:
		return SC
	case SubdivisionSDDC:
		return SD
	case SubdivisionSDDN:
		return SD
	case SubdivisionSDDS:
		return SD
	case SubdivisionSDDW:
		return SD
	case SubdivisionSDGD:
		return SD
	case SubdivisionSDGK:
		return SD
	case SubdivisionSDGZ:
		return SD
	case SubdivisionSDKA:
		return SD
	case SubdivisionSDKH:
		return SD
	case SubdivisionSDKN:
		return SD
	case SubdivisionSDKS:
		return SD
	case SubdivisionSDNB:
		return SD
	case SubdivisionSDNO:
		return SD
	case SubdivisionSDNR:
		return SD
	case SubdivisionSDNW:
		return SD
	case SubdivisionSDRS:
		return SD
	case SubdivisionSDSI:
		return SD
	case SubdivisionSEAB:
		return SE
	case SubdivisionSEAC:
		return SE
	case SubdivisionSEBD:
		return SE
	case SubdivisionSEC:
		return SE
	case SubdivisionSED:
		return SE
	case SubdivisionSEE:
		return SE
	case SubdivisionSEF:
		return SE
	case SubdivisionSEG:
		return SE
	case SubdivisionSEH:
		return SE
	case SubdivisionSEI:
		return SE
	case SubdivisionSEK:
		return SE
	case SubdivisionSEM:
		return SE
	case SubdivisionSEN:
		return SE
	case SubdivisionSEO:
		return SE
	case SubdivisionSES:
		return SE
	case SubdivisionSET:
		return SE
	case SubdivisionSEU:
		return SE
	case SubdivisionSEW:
		return SE
	case SubdivisionSEX:
		return SE
	case SubdivisionSEY:
		return SE
	case SubdivisionSEZ:
		return SE
	case SubdivisionSG01:
		return SG
	case SubdivisionSG02:
		return SG
	case SubdivisionSG03:
		return SG
	case SubdivisionSG04:
		return SG
	case SubdivisionSG05:
		return SG
	case SubdivisionSHHL:
		return SH
	case SubdivisionSI001:
		return SI
	case SubdivisionSI002:
		return SI
	case SubdivisionSI003:
		return SI
	case SubdivisionSI004:
		return SI
	case SubdivisionSI005:
		return SI
	case SubdivisionSI006:
		return SI
	case SubdivisionSI007:
		return SI
	case SubdivisionSI008:
		return SI
	case SubdivisionSI009:
		return SI
	case SubdivisionSI010:
		return SI
	case SubdivisionSI011:
		return SI
	case SubdivisionSI012:
		return SI
	case SubdivisionSI013:
		return SI
	case SubdivisionSI014:
		return SI
	case SubdivisionSI015:
		return SI
	case SubdivisionSI017:
		return SI
	case SubdivisionSI018:
		return SI
	case SubdivisionSI019:
		return SI
	case SubdivisionSI020:
		return SI
	case SubdivisionSI021:
		return SI
	case SubdivisionSI023:
		return SI
	case SubdivisionSI024:
		return SI
	case SubdivisionSI025:
		return SI
	case SubdivisionSI026:
		return SI
	case SubdivisionSI029:
		return SI
	case SubdivisionSI031:
		return SI
	case SubdivisionSI032:
		return SI
	case SubdivisionSI033:
		return SI
	case SubdivisionSI034:
		return SI
	case SubdivisionSI035:
		return SI
	case SubdivisionSI036:
		return SI
	case SubdivisionSI037:
		return SI
	case SubdivisionSI038:
		return SI
	case SubdivisionSI039:
		return SI
	case SubdivisionSI040:
		return SI
	case SubdivisionSI041:
		return SI
	case SubdivisionSI042:
		return SI
	case SubdivisionSI043:
		return SI
	case SubdivisionSI044:
		return SI
	case SubdivisionSI045:
		return SI
	case SubdivisionSI046:
		return SI
	case SubdivisionSI047:
		return SI
	case SubdivisionSI048:
		return SI
	case SubdivisionSI049:
		return SI
	case SubdivisionSI050:
		return SI
	case SubdivisionSI052:
		return SI
	case SubdivisionSI053:
		return SI
	case SubdivisionSI054:
		return SI
	case SubdivisionSI055:
		return SI
	case SubdivisionSI056:
		return SI
	case SubdivisionSI057:
		return SI
	case SubdivisionSI058:
		return SI
	case SubdivisionSI059:
		return SI
	case SubdivisionSI060:
		return SI
	case SubdivisionSI061:
		return SI
	case SubdivisionSI063:
		return SI
	case SubdivisionSI064:
		return SI
	case SubdivisionSI065:
		return SI
	case SubdivisionSI066:
		return SI
	case SubdivisionSI067:
		return SI
	case SubdivisionSI068:
		return SI
	case SubdivisionSI069:
		return SI
	case SubdivisionSI070:
		return SI
	case SubdivisionSI071:
		return SI
	case SubdivisionSI072:
		return SI
	case SubdivisionSI073:
		return SI
	case SubdivisionSI074:
		return SI
	case SubdivisionSI075:
		return SI
	case SubdivisionSI076:
		return SI
	case SubdivisionSI077:
		return SI
	case SubdivisionSI079:
		return SI
	case SubdivisionSI080:
		return SI
	case SubdivisionSI081:
		return SI
	case SubdivisionSI082:
		return SI
	case SubdivisionSI083:
		return SI
	case SubdivisionSI084:
		return SI
	case SubdivisionSI085:
		return SI
	case SubdivisionSI086:
		return SI
	case SubdivisionSI087:
		return SI
	case SubdivisionSI090:
		return SI
	case SubdivisionSI091:
		return SI
	case SubdivisionSI092:
		return SI
	case SubdivisionSI094:
		return SI
	case SubdivisionSI095:
		return SI
	case SubdivisionSI096:
		return SI
	case SubdivisionSI097:
		return SI
	case SubdivisionSI098:
		return SI
	case SubdivisionSI099:
		return SI
	case SubdivisionSI100:
		return SI
	case SubdivisionSI101:
		return SI
	case SubdivisionSI102:
		return SI
	case SubdivisionSI103:
		return SI
	case SubdivisionSI104:
		return SI
	case SubdivisionSI105:
		return SI
	case SubdivisionSI106:
		return SI
	case SubdivisionSI108:
		return SI
	case SubdivisionSI109:
		return SI
	case SubdivisionSI110:
		return SI
	case SubdivisionSI111:
		return SI
	case SubdivisionSI112:
		return SI
	case SubdivisionSI113:
		return SI
	case SubdivisionSI114:
		return SI
	case SubdivisionSI115:
		return SI
	case SubdivisionSI116:
		return SI
	case SubdivisionSI117:
		return SI
	case SubdivisionSI118:
		return SI
	case SubdivisionSI119:
		return SI
	case SubdivisionSI120:
		return SI
	case SubdivisionSI121:
		return SI
	case SubdivisionSI122:
		return SI
	case SubdivisionSI123:
		return SI
	case SubdivisionSI124:
		return SI
	case SubdivisionSI125:
		return SI
	case SubdivisionSI126:
		return SI
	case SubdivisionSI127:
		return SI
	case SubdivisionSI128:
		return SI
	case SubdivisionSI129:
		return SI
	case SubdivisionSI130:
		return SI
	case SubdivisionSI131:
		return SI
	case SubdivisionSI132:
		return SI
	case SubdivisionSI133:
		return SI
	case SubdivisionSI134:
		return SI
	case SubdivisionSI135:
		return SI
	case SubdivisionSI136:
		return SI
	case SubdivisionSI137:
		return SI
	case SubdivisionSI138:
		return SI
	case SubdivisionSI139:
		return SI
	case SubdivisionSI140:
		return SI
	case SubdivisionSI141:
		return SI
	case SubdivisionSI142:
		return SI
	case SubdivisionSI143:
		return SI
	case SubdivisionSI144:
		return SI
	case SubdivisionSI146:
		return SI
	case SubdivisionSI147:
		return SI
	case SubdivisionSI148:
		return SI
	case SubdivisionSI149:
		return SI
	case SubdivisionSI150:
		return SI
	case SubdivisionSI151:
		return SI
	case SubdivisionSI152:
		return SI
	case SubdivisionSI154:
		return SI
	case SubdivisionSI155:
		return SI
	case SubdivisionSI156:
		return SI
	case SubdivisionSI158:
		return SI
	case SubdivisionSI159:
		return SI
	case SubdivisionSI160:
		return SI
	case SubdivisionSI161:
		return SI
	case SubdivisionSI162:
		return SI
	case SubdivisionSI164:
		return SI
	case SubdivisionSI165:
		return SI
	case SubdivisionSI166:
		return SI
	case SubdivisionSI167:
		return SI
	case SubdivisionSI168:
		return SI
	case SubdivisionSI169:
		return SI
	case SubdivisionSI170:
		return SI
	case SubdivisionSI171:
		return SI
	case SubdivisionSI172:
		return SI
	case SubdivisionSI173:
		return SI
	case SubdivisionSI174:
		return SI
	case SubdivisionSI175:
		return SI
	case SubdivisionSI176:
		return SI
	case SubdivisionSI179:
		return SI
	case SubdivisionSI180:
		return SI
	case SubdivisionSI182:
		return SI
	case SubdivisionSI183:
		return SI
	case SubdivisionSI184:
		return SI
	case SubdivisionSI185:
		return SI
	case SubdivisionSI186:
		return SI
	case SubdivisionSI187:
		return SI
	case SubdivisionSI188:
		return SI
	case SubdivisionSI189:
		return SI
	case SubdivisionSI190:
		return SI
	case SubdivisionSI191:
		return SI
	case SubdivisionSI193:
		return SI
	case SubdivisionSI194:
		return SI
	case SubdivisionSI195:
		return SI
	case SubdivisionSI196:
		return SI
	case SubdivisionSI197:
		return SI
	case SubdivisionSI198:
		return SI
	case SubdivisionSI199:
		return SI
	case SubdivisionSI200:
		return SI
	case SubdivisionSI201:
		return SI
	case SubdivisionSI203:
		return SI
	case SubdivisionSI204:
		return SI
	case SubdivisionSI205:
		return SI
	case SubdivisionSI206:
		return SI
	case SubdivisionSI207:
		return SI
	case SubdivisionSI208:
		return SI
	case SubdivisionSI209:
		return SI
	case SubdivisionSI210:
		return SI
	case SubdivisionSI211:
		return SI
	case SubdivisionSI212:
		return SI
	case SubdivisionSI213:
		return SI
	case SubdivisionSKBC:
		return SK
	case SubdivisionSKBL:
		return SK
	case SubdivisionSKKI:
		return SK
	case SubdivisionSKNI:
		return SK
	case SubdivisionSKPV:
		return SK
	case SubdivisionSKTA:
		return SK
	case SubdivisionSKTC:
		return SK
	case SubdivisionSKZI:
		return SK
	case SubdivisionSLE:
		return SL
	case SubdivisionSLN:
		return SL
	case SubdivisionSLS:
		return SL
	case SubdivisionSLW:
		return SL
	case SubdivisionSM03:
		return SM
	case SubdivisionSM07:
		return SM
	case SubdivisionSM09:
		return SM
	case SubdivisionSNDB:
		return SN
	case SubdivisionSNDK:
		return SN
	case SubdivisionSNFK:
		return SN
	case SubdivisionSNKA:
		return SN
	case SubdivisionSNKD:
		return SN
	case SubdivisionSNKE:
		return SN
	case SubdivisionSNKL:
		return SN
	case SubdivisionSNLG:
		return SN
	case SubdivisionSNMT:
		return SN
	case SubdivisionSNSL:
		return SN
	case SubdivisionSNTC:
		return SN
	case SubdivisionSNTH:
		return SN
	case SubdivisionSNZG:
		return SN
	case SubdivisionSOAW:
		return SO
	case SubdivisionSOBN:
		return SO
	case SubdivisionSOMU:
		return SO
	case SubdivisionSONU:
		return SO
	case SubdivisionSOSH:
		return SO
	case SubdivisionSOSO:
		return SO
	case SubdivisionSOTO:
		return SO
	case SubdivisionSOWO:
		return SO
	case SubdivisionSRCM:
		return SR
	case SubdivisionSRNI:
		return SR
	case SubdivisionSRPM:
		return SR
	case SubdivisionSRPR:
		return SR
	case SubdivisionSRSI:
		return SR
	case SubdivisionSRWA:
		return SR
	case SubdivisionSSBN:
		return SS
	case SubdivisionSSEC:
		return SS
	case SubdivisionSSEE:
		return SS
	case SubdivisionSSEW:
		return SS
	case SubdivisionSSNU:
		return SS
	case SubdivisionST01:
		return ST
	case SubdivisionSVAH:
		return SV
	case SubdivisionSVCA:
		return SV
	case SubdivisionSVCH:
		return SV
	case SubdivisionSVCU:
		return SV
	case SubdivisionSVLI:
		return SV
	case SubdivisionSVMO:
		return SV
	case SubdivisionSVPA:
		return SV
	case SubdivisionSVSA:
		return SV
	case SubdivisionSVSM:
		return SV
	case SubdivisionSVSO:
		return SV
	case SubdivisionSVSS:
		return SV
	case SubdivisionSVSV:
		return SV
	case SubdivisionSVUN:
		return SV
	case SubdivisionSVUS:
		return SV
	case SubdivisionSYDI:
		return SY
	case SubdivisionSYDR:
		return SY
	case SubdivisionSYDY:
		return SY
	case SubdivisionSYHA:
		return SY
	case SubdivisionSYHI:
		return SY
	case SubdivisionSYHL:
		return SY
	case SubdivisionSYHM:
		return SY
	case SubdivisionSYID:
		return SY
	case SubdivisionSYLA:
		return SY
	case SubdivisionSYQU:
		return SY
	case SubdivisionSYRA:
		return SY
	case SubdivisionSYRD:
		return SY
	case SubdivisionSYSU:
		return SY
	case SubdivisionSYTA:
		return SY
	case SubdivisionSZHH:
		return SZ
	case SubdivisionSZLU:
		return SZ
	case SubdivisionSZMA:
		return SZ
	case SubdivisionTDGR:
		return TD
	case SubdivisionTDLO:
		return TD
	case SubdivisionTDME:
		return TD
	case SubdivisionTDND:
		return TD
	case SubdivisionTDOD:
		return TD
	case SubdivisionTGC:
		return TG
	case SubdivisionTGK:
		return TG
	case SubdivisionTGM:
		return TG
	case SubdivisionTGP:
		return TG
	case SubdivisionTH10:
		return TH
	case SubdivisionTH11:
		return TH
	case SubdivisionTH12:
		return TH
	case SubdivisionTH13:
		return TH
	case SubdivisionTH14:
		return TH
	case SubdivisionTH15:
		return TH
	case SubdivisionTH16:
		return TH
	case SubdivisionTH17:
		return TH
	case SubdivisionTH18:
		return TH
	case SubdivisionTH19:
		return TH
	case SubdivisionTH20:
		return TH
	case SubdivisionTH21:
		return TH
	case SubdivisionTH22:
		return TH
	case SubdivisionTH23:
		return TH
	case SubdivisionTH24:
		return TH
	case SubdivisionTH25:
		return TH
	case SubdivisionTH26:
		return TH
	case SubdivisionTH27:
		return TH
	case SubdivisionTH30:
		return TH
	case SubdivisionTH31:
		return TH
	case SubdivisionTH32:
		return TH
	case SubdivisionTH33:
		return TH
	case SubdivisionTH34:
		return TH
	case SubdivisionTH35:
		return TH
	case SubdivisionTH36:
		return TH
	case SubdivisionTH37:
		return TH
	case SubdivisionTH38:
		return TH
	case SubdivisionTH39:
		return TH
	case SubdivisionTH40:
		return TH
	case SubdivisionTH41:
		return TH
	case SubdivisionTH42:
		return TH
	case SubdivisionTH43:
		return TH
	case SubdivisionTH44:
		return TH
	case SubdivisionTH45:
		return TH
	case SubdivisionTH46:
		return TH
	case SubdivisionTH47:
		return TH
	case SubdivisionTH48:
		return TH
	case SubdivisionTH49:
		return TH
	case SubdivisionTH50:
		return TH
	case SubdivisionTH51:
		return TH
	case SubdivisionTH52:
		return TH
	case SubdivisionTH53:
		return TH
	case SubdivisionTH54:
		return TH
	case SubdivisionTH55:
		return TH
	case SubdivisionTH56:
		return TH
	case SubdivisionTH57:
		return TH
	case SubdivisionTH58:
		return TH
	case SubdivisionTH60:
		return TH
	case SubdivisionTH61:
		return TH
	case SubdivisionTH62:
		return TH
	case SubdivisionTH63:
		return TH
	case SubdivisionTH64:
		return TH
	case SubdivisionTH65:
		return TH
	case SubdivisionTH66:
		return TH
	case SubdivisionTH67:
		return TH
	case SubdivisionTH70:
		return TH
	case SubdivisionTH71:
		return TH
	case SubdivisionTH72:
		return TH
	case SubdivisionTH73:
		return TH
	case SubdivisionTH74:
		return TH
	case SubdivisionTH75:
		return TH
	case SubdivisionTH76:
		return TH
	case SubdivisionTH77:
		return TH
	case SubdivisionTH80:
		return TH
	case SubdivisionTH81:
		return TH
	case SubdivisionTH82:
		return TH
	case SubdivisionTH83:
		return TH
	case SubdivisionTH84:
		return TH
	case SubdivisionTH85:
		return TH
	case SubdivisionTH86:
		return TH
	case SubdivisionTH90:
		return TH
	case SubdivisionTH91:
		return TH
	case SubdivisionTH92:
		return TH
	case SubdivisionTH93:
		return TH
	case SubdivisionTH94:
		return TH
	case SubdivisionTH95:
		return TH
	case SubdivisionTH96:
		return TH
	case SubdivisionTJDU:
		return TJ
	case SubdivisionTJKT:
		return TJ
	case SubdivisionTJRA:
		return TJ
	case SubdivisionTJSU:
		return TJ
	case SubdivisionTLAN:
		return TL
	case SubdivisionTLCO:
		return TL
	case SubdivisionTLDI:
		return TL
	case SubdivisionTMA:
		return TM
	case SubdivisionTMB:
		return TM
	case SubdivisionTMD:
		return TM
	case SubdivisionTML:
		return TM
	case SubdivisionTMM:
		return TM
	case SubdivisionTN11:
		return TN
	case SubdivisionTN12:
		return TN
	case SubdivisionTN13:
		return TN
	case SubdivisionTN14:
		return TN
	case SubdivisionTN21:
		return TN
	case SubdivisionTN22:
		return TN
	case SubdivisionTN23:
		return TN
	case SubdivisionTN31:
		return TN
	case SubdivisionTN32:
		return TN
	case SubdivisionTN33:
		return TN
	case SubdivisionTN34:
		return TN
	case SubdivisionTN41:
		return TN
	case SubdivisionTN42:
		return TN
	case SubdivisionTN43:
		return TN
	case SubdivisionTN51:
		return TN
	case SubdivisionTN52:
		return TN
	case SubdivisionTN53:
		return TN
	case SubdivisionTN61:
		return TN
	case SubdivisionTN71:
		return TN
	case SubdivisionTN72:
		return TN
	case SubdivisionTN73:
		return TN
	case SubdivisionTN81:
		return TN
	case SubdivisionTN82:
		return TN
	case SubdivisionTN83:
		return TN
	case SubdivisionTO03:
		return TO
	case SubdivisionTO04:
		return TO
	case SubdivisionTR01:
		return TR
	case SubdivisionTR02:
		return TR
	case SubdivisionTR03:
		return TR
	case SubdivisionTR04:
		return TR
	case SubdivisionTR05:
		return TR
	case SubdivisionTR06:
		return TR
	case SubdivisionTR07:
		return TR
	case SubdivisionTR08:
		return TR
	case SubdivisionTR09:
		return TR
	case SubdivisionTR10:
		return TR
	case SubdivisionTR11:
		return TR
	case SubdivisionTR12:
		return TR
	case SubdivisionTR13:
		return TR
	case SubdivisionTR14:
		return TR
	case SubdivisionTR15:
		return TR
	case SubdivisionTR16:
		return TR
	case SubdivisionTR17:
		return TR
	case SubdivisionTR18:
		return TR
	case SubdivisionTR19:
		return TR
	case SubdivisionTR20:
		return TR
	case SubdivisionTR21:
		return TR
	case SubdivisionTR22:
		return TR
	case SubdivisionTR23:
		return TR
	case SubdivisionTR24:
		return TR
	case SubdivisionTR25:
		return TR
	case SubdivisionTR26:
		return TR
	case SubdivisionTR27:
		return TR
	case SubdivisionTR28:
		return TR
	case SubdivisionTR29:
		return TR
	case SubdivisionTR30:
		return TR
	case SubdivisionTR31:
		return TR
	case SubdivisionTR32:
		return TR
	case SubdivisionTR33:
		return TR
	case SubdivisionTR34:
		return TR
	case SubdivisionTR35:
		return TR
	case SubdivisionTR36:
		return TR
	case SubdivisionTR37:
		return TR
	case SubdivisionTR38:
		return TR
	case SubdivisionTR39:
		return TR
	case SubdivisionTR40:
		return TR
	case SubdivisionTR41:
		return TR
	case SubdivisionTR42:
		return TR
	case SubdivisionTR43:
		return TR
	case SubdivisionTR44:
		return TR
	case SubdivisionTR45:
		return TR
	case SubdivisionTR46:
		return TR
	case SubdivisionTR47:
		return TR
	case SubdivisionTR48:
		return TR
	case SubdivisionTR49:
		return TR
	case SubdivisionTR50:
		return TR
	case SubdivisionTR51:
		return TR
	case SubdivisionTR52:
		return TR
	case SubdivisionTR53:
		return TR
	case SubdivisionTR54:
		return TR
	case SubdivisionTR55:
		return TR
	case SubdivisionTR56:
		return TR
	case SubdivisionTR57:
		return TR
	case SubdivisionTR58:
		return TR
	case SubdivisionTR59:
		return TR
	case SubdivisionTR60:
		return TR
	case SubdivisionTR61:
		return TR
	case SubdivisionTR62:
		return TR
	case SubdivisionTR63:
		return TR
	case SubdivisionTR64:
		return TR
	case SubdivisionTR65:
		return TR
	case SubdivisionTR66:
		return TR
	case SubdivisionTR67:
		return TR
	case SubdivisionTR68:
		return TR
	case SubdivisionTR69:
		return TR
	case SubdivisionTR70:
		return TR
	case SubdivisionTR71:
		return TR
	case SubdivisionTR72:
		return TR
	case SubdivisionTR73:
		return TR
	case SubdivisionTR74:
		return TR
	case SubdivisionTR75:
		return TR
	case SubdivisionTR76:
		return TR
	case SubdivisionTR77:
		return TR
	case SubdivisionTR78:
		return TR
	case SubdivisionTR79:
		return TR
	case SubdivisionTR80:
		return TR
	case SubdivisionTR81:
		return TR
	case SubdivisionTTARI:
		return TT
	case SubdivisionTTCHA:
		return TT
	case SubdivisionTTCTT:
		return TT
	case SubdivisionTTDMN:
		return TT
	case SubdivisionTTMRC:
		return TT
	case SubdivisionTTPED:
		return TT
	case SubdivisionTTPOS:
		return TT
	case SubdivisionTTPRT:
		return TT
	case SubdivisionTTPTF:
		return TT
	case SubdivisionTTSFO:
		return TT
	case SubdivisionTTSGE:
		return TT
	case SubdivisionTTSIP:
		return TT
	case SubdivisionTTSJL:
		return TT
	case SubdivisionTTTOB:
		return TT
	case SubdivisionTTTUP:
		return TT
	case SubdivisionTVFUN:
		return TV
	case SubdivisionTWCHA:
		return TW
	case SubdivisionTWCYQ:
		return TW
	case SubdivisionTWHSQ:
		return TW
	case SubdivisionTWHUA:
		return TW
	case SubdivisionTWILA:
		return TW
	case SubdivisionTWKEE:
		return TW
	case SubdivisionTWKHH:
		return TW
	case SubdivisionTWKIN:
		return TW
	case SubdivisionTWLIE:
		return TW
	case SubdivisionTWMIA:
		return TW
	case SubdivisionTWNAN:
		return TW
	case SubdivisionTWNWT:
		return TW
	case SubdivisionTWPEN:
		return TW
	case SubdivisionTWPIF:
		return TW
	case SubdivisionTWTAO:
		return TW
	case SubdivisionTWTNN:
		return TW
	case SubdivisionTWTPE:
		return TW
	case SubdivisionTWTTT:
		return TW
	case SubdivisionTWTXG:
		return TW
	case SubdivisionTWYUN:
		return TW
	case SubdivisionTZ01:
		return TZ
	case SubdivisionTZ02:
		return TZ
	case SubdivisionTZ03:
		return TZ
	case SubdivisionTZ04:
		return TZ
	case SubdivisionTZ05:
		return TZ
	case SubdivisionTZ07:
		return TZ
	case SubdivisionTZ08:
		return TZ
	case SubdivisionTZ09:
		return TZ
	case SubdivisionTZ11:
		return TZ
	case SubdivisionTZ12:
		return TZ
	case SubdivisionTZ13:
		return TZ
	case SubdivisionTZ14:
		return TZ
	case SubdivisionTZ15:
		return TZ
	case SubdivisionTZ16:
		return TZ
	case SubdivisionTZ17:
		return TZ
	case SubdivisionTZ18:
		return TZ
	case SubdivisionTZ19:
		return TZ
	case SubdivisionTZ20:
		return TZ
	case SubdivisionTZ21:
		return TZ
	case SubdivisionTZ22:
		return TZ
	case SubdivisionTZ23:
		return TZ
	case SubdivisionTZ24:
		return TZ
	case SubdivisionTZ25:
		return TZ
	case SubdivisionTZ26:
		return TZ
	case SubdivisionTZ27:
		return TZ
	case SubdivisionTZ28:
		return TZ
	case SubdivisionTZ29:
		return TZ
	case SubdivisionTZ30:
		return TZ
	case SubdivisionTZ31:
		return TZ
	case SubdivisionUA05:
		return UA
	case SubdivisionUA07:
		return UA
	case SubdivisionUA09:
		return UA
	case SubdivisionUA12:
		return UA
	case SubdivisionUA14:
		return UA
	case SubdivisionUA18:
		return UA
	case SubdivisionUA21:
		return UA
	case SubdivisionUA23:
		return UA
	case SubdivisionUA26:
		return UA
	case SubdivisionUA30:
		return UA
	case SubdivisionUA32:
		return UA
	case SubdivisionUA35:
		return UA
	case SubdivisionUA40:
		return UA
	case SubdivisionUA43:
		return UA
	case SubdivisionUA46:
		return UA
	case SubdivisionUA48:
		return UA
	case SubdivisionUA51:
		return UA
	case SubdivisionUA53:
		return UA
	case SubdivisionUA56:
		return UA
	case SubdivisionUA59:
		return UA
	case SubdivisionUA61:
		return UA
	case SubdivisionUA63:
		return UA
	case SubdivisionUA65:
		return UA
	case SubdivisionUA68:
		return UA
	case SubdivisionUA71:
		return UA
	case SubdivisionUA74:
		return UA
	case SubdivisionUA77:
		return UA
	case SubdivisionUG101:
		return UG
	case SubdivisionUG102:
		return UG
	case SubdivisionUG103:
		return UG
	case SubdivisionUG104:
		return UG
	case SubdivisionUG105:
		return UG
	case SubdivisionUG106:
		return UG
	case SubdivisionUG107:
		return UG
	case SubdivisionUG108:
		return UG
	case SubdivisionUG109:
		return UG
	case SubdivisionUG112:
		return UG
	case SubdivisionUG113:
		return UG
	case SubdivisionUG115:
		return UG
	case SubdivisionUG116:
		return UG
	case SubdivisionUG117:
		return UG
	case SubdivisionUG120:
		return UG
	case SubdivisionUG122:
		return UG
	case SubdivisionUG201:
		return UG
	case SubdivisionUG203:
		return UG
	case SubdivisionUG204:
		return UG
	case SubdivisionUG205:
		return UG
	case SubdivisionUG206:
		return UG
	case SubdivisionUG209:
		return UG
	case SubdivisionUG210:
		return UG
	case SubdivisionUG211:
		return UG
	case SubdivisionUG214:
		return UG
	case SubdivisionUG215:
		return UG
	case SubdivisionUG219:
		return UG
	case SubdivisionUG222:
		return UG
	case SubdivisionUG303:
		return UG
	case SubdivisionUG304:
		return UG
	case SubdivisionUG305:
		return UG
	case SubdivisionUG307:
		return UG
	case SubdivisionUG316:
		return UG
	case SubdivisionUG321:
		return UG
	case SubdivisionUG403:
		return UG
	case SubdivisionUG404:
		return UG
	case SubdivisionUG405:
		return UG
	case SubdivisionUG406:
		return UG
	case SubdivisionUG407:
		return UG
	case SubdivisionUG408:
		return UG
	case SubdivisionUG409:
		return UG
	case SubdivisionUG410:
		return UG
	case SubdivisionUG411:
		return UG
	case SubdivisionUG412:
		return UG
	case SubdivisionUG413:
		return UG
	case SubdivisionUG415:
		return UG
	case SubdivisionUG419:
		return UG
	case SubdivisionUM95:
		return UM
	case SubdivisionUSAK:
		return US
	case SubdivisionUSAL:
		return US
	case SubdivisionUSAR:
		return US
	case SubdivisionUSAZ:
		return US
	case SubdivisionUSCA:
		return US
	case SubdivisionUSCO:
		return US
	case SubdivisionUSCT:
		return US
	case SubdivisionUSDC:
		return US
	case SubdivisionUSDE:
		return US
	case SubdivisionUSFL:
		return US
	case SubdivisionUSGA:
		return US
	case SubdivisionUSHI:
		return US
	case SubdivisionUSIA:
		return US
	case SubdivisionUSID:
		return US
	case SubdivisionUSIL:
		return US
	case SubdivisionUSIN:
		return US
	case SubdivisionUSKS:
		return US
	case SubdivisionUSKY:
		return US
	case SubdivisionUSLA:
		return US
	case SubdivisionUSMA:
		return US
	case SubdivisionUSMD:
		return US
	case SubdivisionUSME:
		return US
	case SubdivisionUSMI:
		return US
	case SubdivisionUSMN:
		return US
	case SubdivisionUSMO:
		return US
	case SubdivisionUSMS:
		return US
	case SubdivisionUSMT:
		return US
	case SubdivisionUSNC:
		return US
	case SubdivisionUSND:
		return US
	case SubdivisionUSNE:
		return US
	case SubdivisionUSNH:
		return US
	case SubdivisionUSNJ:
		return US
	case SubdivisionUSNM:
		return US
	case SubdivisionUSNV:
		return US
	case SubdivisionUSNY:
		return US
	case SubdivisionUSOH:
		return US
	case SubdivisionUSOK:
		return US
	case SubdivisionUSOR:
		return US
	case SubdivisionUSPA:
		return US
	case SubdivisionUSRI:
		return US
	case SubdivisionUSSC:
		return US
	case SubdivisionUSSD:
		return US
	case SubdivisionUSTN:
		return US
	case SubdivisionUSTX:
		return US
	case SubdivisionUSUT:
		return US
	case SubdivisionUSVA:
		return US
	case SubdivisionUSVT:
		return US
	case SubdivisionUSWA:
		return US
	case SubdivisionUSWI:
		return US
	case SubdivisionUSWV:
		return US
	case SubdivisionUSWY:
		return US
	case SubdivisionUYAR:
		return UY
	case SubdivisionUYCA:
		return UY
	case SubdivisionUYCL:
		return UY
	case SubdivisionUYCO:
		return UY
	case SubdivisionUYDU:
		return UY
	case SubdivisionUYFD:
		return UY
	case SubdivisionUYFS:
		return UY
	case SubdivisionUYLA:
		return UY
	case SubdivisionUYMA:
		return UY
	case SubdivisionUYMO:
		return UY
	case SubdivisionUYPA:
		return UY
	case SubdivisionUYRN:
		return UY
	case SubdivisionUYRO:
		return UY
	case SubdivisionUYRV:
		return UY
	case SubdivisionUYSA:
		return UY
	case SubdivisionUYSJ:
		return UY
	case SubdivisionUYSO:
		return UY
	case SubdivisionUYTA:
		return UY
	case SubdivisionUYTT:
		return UY
	case SubdivisionUZAN:
		return UZ
	case SubdivisionUZBU:
		return UZ
	case SubdivisionUZFA:
		return UZ
	case SubdivisionUZJI:
		return UZ
	case SubdivisionUZNG:
		return UZ
	case SubdivisionUZNW:
		return UZ
	case SubdivisionUZQA:
		return UZ
	case SubdivisionUZQR:
		return UZ
	case SubdivisionUZSA:
		return UZ
	case SubdivisionUZSI:
		return UZ
	case SubdivisionUZSU:
		return UZ
	case SubdivisionUZTK:
		return UZ
	case SubdivisionUZXO:
		return UZ
	case SubdivisionVC01:
		return VC
	case SubdivisionVC04:
		return VC
	case SubdivisionVC05:
		return VC
	case SubdivisionVC06:
		return VC
	case SubdivisionVEA:
		return VE
	case SubdivisionVEB:
		return VE
	case SubdivisionVEC:
		return VE
	case SubdivisionVED:
		return VE
	case SubdivisionVEE:
		return VE
	case SubdivisionVEF:
		return VE
	case SubdivisionVEG:
		return VE
	case SubdivisionVEH:
		return VE
	case SubdivisionVEI:
		return VE
	case SubdivisionVEJ:
		return VE
	case SubdivisionVEK:
		return VE
	case SubdivisionVEL:
		return VE
	case SubdivisionVEM:
		return VE
	case SubdivisionVEN:
		return VE
	case SubdivisionVEO:
		return VE
	case SubdivisionVEP:
		return VE
	case SubdivisionVER:
		return VE
	case SubdivisionVES:
		return VE
	case SubdivisionVET:
		return VE
	case SubdivisionVEU:
		return VE
	case SubdivisionVEV:
		return VE
	case SubdivisionVEX:
		return VE
	case SubdivisionVEY:
		return VE
	case SubdivisionVEZ:
		return VE
	case SubdivisionVN01:
		return VN
	case SubdivisionVN02:
		return VN
	case SubdivisionVN03:
		return VN
	case SubdivisionVN04:
		return VN
	case SubdivisionVN05:
		return VN
	case SubdivisionVN06:
		return VN
	case SubdivisionVN07:
		return VN
	case SubdivisionVN09:
		return VN
	case SubdivisionVN13:
		return VN
	case SubdivisionVN14:
		return VN
	case SubdivisionVN18:
		return VN
	case SubdivisionVN20:
		return VN
	case SubdivisionVN21:
		return VN
	case SubdivisionVN22:
		return VN
	case SubdivisionVN23:
		return VN
	case SubdivisionVN24:
		return VN
	case SubdivisionVN25:
		return VN
	case SubdivisionVN26:
		return VN
	case SubdivisionVN27:
		return VN
	case SubdivisionVN28:
		return VN
	case SubdivisionVN29:
		return VN
	case SubdivisionVN30:
		return VN
	case SubdivisionVN31:
		return VN
	case SubdivisionVN32:
		return VN
	case SubdivisionVN33:
		return VN
	case SubdivisionVN34:
		return VN
	case SubdivisionVN35:
		return VN
	case SubdivisionVN36:
		return VN
	case SubdivisionVN37:
		return VN
	case SubdivisionVN39:
		return VN
	case SubdivisionVN40:
		return VN
	case SubdivisionVN41:
		return VN
	case SubdivisionVN43:
		return VN
	case SubdivisionVN44:
		return VN
	case SubdivisionVN45:
		return VN
	case SubdivisionVN46:
		return VN
	case SubdivisionVN47:
		return VN
	case SubdivisionVN49:
		return VN
	case SubdivisionVN50:
		return VN
	case SubdivisionVN51:
		return VN
	case SubdivisionVN52:
		return VN
	case SubdivisionVN53:
		return VN
	case SubdivisionVN54:
		return VN
	case SubdivisionVN55:
		return VN
	case SubdivisionVN56:
		return VN
	case SubdivisionVN57:
		return VN
	case SubdivisionVN58:
		return VN
	case SubdivisionVN59:
		return VN
	case SubdivisionVN61:
		return VN
	case SubdivisionVN63:
		return VN
	case SubdivisionVN66:
		return VN
	case SubdivisionVN67:
		return VN
	case SubdivisionVN68:
		return VN
	case SubdivisionVN69:
		return VN
	case SubdivisionVN70:
		return VN
	case SubdivisionVN71:
		return VN
	case SubdivisionVN72:
		return VN
	case SubdivisionVNCT:
		return VN
	case SubdivisionVNDN:
		return VN
	case SubdivisionVNHN:
		return VN
	case SubdivisionVNHP:
		return VN
	case SubdivisionVNSG:
		return VN
	case SubdivisionVUSEE:
		return VU
	case SubdivisionVUTAE:
		return VU
	case SubdivisionVUTOB:
		return VU
	case SubdivisionWFSG:
		return WF
	case SubdivisionWFUV:
		return WF
	case SubdivisionWSAT:
		return WS
	case SubdivisionWSFA:
		return WS
	case SubdivisionWSTU:
		return WS
	case SubdivisionYEAB:
		return YE
	case SubdivisionYEAD:
		return YE
	case SubdivisionYEAM:
		return YE
	case SubdivisionYEBA:
		return YE
	case SubdivisionYEDA:
		return YE
	case SubdivisionYEDH:
		return YE
	case SubdivisionYEHD:
		return YE
	case SubdivisionYEHJ:
		return YE
	case SubdivisionYEHU:
		return YE
	case SubdivisionYEIB:
		return YE
	case SubdivisionYELA:
		return YE
	case SubdivisionYEMA:
		return YE
	case SubdivisionYEMR:
		return YE
	case SubdivisionYESA:
		return YE
	case SubdivisionYESD:
		return YE
	case SubdivisionYESH:
		return YE
	case SubdivisionYESN:
		return YE
	case SubdivisionYETA:
		return YE
	case SubdivisionZAEC:
		return ZA
	case SubdivisionZAFS:
		return ZA
	case SubdivisionZAGP:
		return ZA
	case SubdivisionZAKZN:
		return ZA
	case SubdivisionZALP:
		return ZA
	case SubdivisionZAMP:
		return ZA
	case SubdivisionZANC:
		return ZA
	case SubdivisionZANW:
		return ZA
	case SubdivisionZAWC:
		return ZA
	case SubdivisionZM01:
		return ZM
	case SubdivisionZM02:
		return ZM
	case SubdivisionZM03:
		return ZM
	case SubdivisionZM04:
		return ZM
	case SubdivisionZM05:
		return ZM
	case SubdivisionZM06:
		return ZM
	case SubdivisionZM07:
		return ZM
	case SubdivisionZM08:
		return ZM
	case SubdivisionZM09:
		return ZM
	case SubdivisionZM10:
		return ZM
	case SubdivisionZWBU:
		return ZW
	case SubdivisionZWHA:
		return ZW
	case SubdivisionZWMA:
		return ZW
	case SubdivisionZWMC:
		return ZW
	case SubdivisionZWME:
		return ZW
	case SubdivisionZWMI:
		return ZW
	case SubdivisionZWMN:
		return ZW
	case SubdivisionZWMS:
		return ZW
	case SubdivisionZWMV:
		return ZW
	case SubdivisionZWMW:
		return ZW
	}

	return Unknown
}

// IsValid - returns true, if code is correct
func (s SubdivisionCode) IsValid() bool {
	return s.String() != UnknownMsg
}

// Info - return CapitalCode as Capital info
func (s SubdivisionCode) Info() *Subdivision {
	return &Subdivision{
		Name:    s.String(),
		Code:    s,
		Country: s.Country(),
	}
}

// Type implements Typer interface
func (s Subdivision) Type() string {
	return TypeSubdivision
}

// Value implements database/sql/driver.Valuer
func (s Subdivision) Value() (Value, error) {
	return json.Marshal(s)
}

// Scan implements database/sql.Scanner
func (s *Subdivision) Scan(src interface{}) error {
	if s == nil {
		return fmt.Errorf("countries::Scan: Subdivision scan err: subdivision == nil")
	}
	switch src := src.(type) {
	case *Subdivision:
		*s = *src
	case Subdivision:
		*s = src
	default:
		return fmt.Errorf("countries::Scan: Subdivision scan err: unexpected value of type %T for %T", src, *s)
	}
	return nil
}

// AllSubdivisions - return all subdivision codes
//nolint:funlen
func AllSubdivisions() []SubdivisionCode {
	return []SubdivisionCode{
		SubdivisionUnknown,
		SubdivisionAD02,
		SubdivisionAD03,
		SubdivisionAD04,
		SubdivisionAD05,
		SubdivisionAD06,
		SubdivisionAD07,
		SubdivisionAD08,
		SubdivisionAEAJ,
		SubdivisionAEAZ,
		SubdivisionAEDU,
		SubdivisionAEFU,
		SubdivisionAERK,
		SubdivisionAESH,
		SubdivisionAEUQ,
		SubdivisionAFBAL,
		SubdivisionAFBAM,
		SubdivisionAFBDG,
		SubdivisionAFBDS,
		SubdivisionAFBGL,
		SubdivisionAFDAY,
		SubdivisionAFFRA,
		SubdivisionAFFYB,
		SubdivisionAFGHA,
		SubdivisionAFGHO,
		SubdivisionAFHEL,
		SubdivisionAFHER,
		SubdivisionAFJOW,
		SubdivisionAFKAB,
		SubdivisionAFKAN,
		SubdivisionAFKDZ,
		SubdivisionAFKHO,
		SubdivisionAFLAG,
		SubdivisionAFLOG,
		SubdivisionAFNAN,
		SubdivisionAFNIM,
		SubdivisionAFPAR,
		SubdivisionAFPIA,
		SubdivisionAFPKA,
		SubdivisionAFTAK,
		SubdivisionAFURU,
		SubdivisionAG03,
		SubdivisionAG04,
		SubdivisionAG05,
		SubdivisionAG06,
		SubdivisionAG07,
		SubdivisionAG08,
		SubdivisionAG11,
		SubdivisionAL01,
		SubdivisionAL02,
		SubdivisionAL03,
		SubdivisionAL04,
		SubdivisionAL05,
		SubdivisionAL06,
		SubdivisionAL07,
		SubdivisionAL08,
		SubdivisionAL09,
		SubdivisionAL10,
		SubdivisionAL11,
		SubdivisionAL12,
		SubdivisionAMAG,
		SubdivisionAMAR,
		SubdivisionAMAV,
		SubdivisionAMER,
		SubdivisionAMGR,
		SubdivisionAMKT,
		SubdivisionAMLO,
		SubdivisionAMSH,
		SubdivisionAMSU,
		SubdivisionAMTV,
		SubdivisionAMVD,
		SubdivisionAOBGO,
		SubdivisionAOBGU,
		SubdivisionAOBIE,
		SubdivisionAOCAB,
		SubdivisionAOCCU,
		SubdivisionAOCNN,
		SubdivisionAOCNO,
		SubdivisionAOCUS,
		SubdivisionAOHUA,
		SubdivisionAOHUI,
		SubdivisionAOLNO,
		SubdivisionAOLSU,
		SubdivisionAOLUA,
		SubdivisionAOMAL,
		SubdivisionAOMOX,
		SubdivisionAONAM,
		SubdivisionAOUIG,
		SubdivisionAOZAI,
		SubdivisionARA,
		SubdivisionARB,
		SubdivisionARC,
		SubdivisionARD,
		SubdivisionARE,
		SubdivisionARF,
		SubdivisionARG,
		SubdivisionARH,
		SubdivisionARJ,
		SubdivisionARK,
		SubdivisionARL,
		SubdivisionARM,
		SubdivisionARN,
		SubdivisionARP,
		SubdivisionARQ,
		SubdivisionARR,
		SubdivisionARS,
		SubdivisionART,
		SubdivisionARU,
		SubdivisionARV,
		SubdivisionARW,
		SubdivisionARX,
		SubdivisionARY,
		SubdivisionARZ,
		SubdivisionAT1,
		SubdivisionAT2,
		SubdivisionAT3,
		SubdivisionAT4,
		SubdivisionAT5,
		SubdivisionAT6,
		SubdivisionAT7,
		SubdivisionAT8,
		SubdivisionAT9,
		SubdivisionAUACT,
		SubdivisionAUNSW,
		SubdivisionAUNT,
		SubdivisionAUQLD,
		SubdivisionAUSA,
		SubdivisionAUTAS,
		SubdivisionAUVIC,
		SubdivisionAUWA,
		SubdivisionAZABS,
		SubdivisionAZAGC,
		SubdivisionAZAGS,
		SubdivisionAZAST,
		SubdivisionAZBA,
		SubdivisionAZBAL,
		SubdivisionAZBAR,
		SubdivisionAZBEY,
		SubdivisionAZBIL,
		SubdivisionAZCAL,
		SubdivisionAZFUZ,
		SubdivisionAZGA,
		SubdivisionAZGAD,
		SubdivisionAZGOR,
		SubdivisionAZGOY,
		SubdivisionAZGYG,
		SubdivisionAZIMI,
		SubdivisionAZISM,
		SubdivisionAZKUR,
		SubdivisionAZLA,
		SubdivisionAZMAS,
		SubdivisionAZMI,
		SubdivisionAZNEF,
		SubdivisionAZNX,
		SubdivisionAZOGU,
		SubdivisionAZQAB,
		SubdivisionAZQAX,
		SubdivisionAZQBA,
		SubdivisionAZQUS,
		SubdivisionAZSAB,
		SubdivisionAZSAK,
		SubdivisionAZSAL,
		SubdivisionAZSAT,
		SubdivisionAZSIY,
		SubdivisionAZSKR,
		SubdivisionAZSM,
		SubdivisionAZSMI,
		SubdivisionAZSMX,
		SubdivisionAZSR,
		SubdivisionAZTAR,
		SubdivisionAZXAC,
		SubdivisionAZXIZ,
		SubdivisionAZYAR,
		SubdivisionAZYEV,
		SubdivisionAZZAQ,
		SubdivisionAZZAR,
		SubdivisionBABIH,
		SubdivisionBABRC,
		SubdivisionBASRP,
		SubdivisionBB01,
		SubdivisionBB02,
		SubdivisionBB03,
		SubdivisionBB04,
		SubdivisionBB05,
		SubdivisionBB07,
		SubdivisionBB08,
		SubdivisionBB09,
		SubdivisionBB10,
		SubdivisionBB11,
		SubdivisionBDA,
		SubdivisionBDB,
		SubdivisionBDC,
		SubdivisionBDD,
		SubdivisionBDE,
		SubdivisionBDF,
		SubdivisionBDG,
		SubdivisionBEBRU,
		SubdivisionBEVAN,
		SubdivisionBEVBR,
		SubdivisionBEVLI,
		SubdivisionBEVOV,
		SubdivisionBEVWV,
		SubdivisionBEWBR,
		SubdivisionBEWHT,
		SubdivisionBEWLG,
		SubdivisionBEWLX,
		SubdivisionBEWNA,
		SubdivisionBFBAM,
		SubdivisionBFBAZ,
		SubdivisionBFBLG,
		SubdivisionBFBLK,
		SubdivisionBFGAN,
		SubdivisionBFGNA,
		SubdivisionBFGOU,
		SubdivisionBFHOU,
		SubdivisionBFKAD,
		SubdivisionBFKMD,
		SubdivisionBFKMP,
		SubdivisionBFKOP,
		SubdivisionBFKOT,
		SubdivisionBFKOW,
		SubdivisionBFLER,
		SubdivisionBFLOR,
		SubdivisionBFMOU,
		SubdivisionBFNAM,
		SubdivisionBFNAO,
		SubdivisionBFNAY,
		SubdivisionBFOUB,
		SubdivisionBFOUD,
		SubdivisionBFPAS,
		SubdivisionBFPON,
		SubdivisionBFSEN,
		SubdivisionBFSIS,
		SubdivisionBFSMT,
		SubdivisionBFSOM,
		SubdivisionBFSOR,
		SubdivisionBFTAP,
		SubdivisionBFTUI,
		SubdivisionBFYAT,
		SubdivisionBFZIR,
		SubdivisionBFZON,
		SubdivisionBFZOU,
		SubdivisionBG01,
		SubdivisionBG02,
		SubdivisionBG03,
		SubdivisionBG04,
		SubdivisionBG05,
		SubdivisionBG06,
		SubdivisionBG07,
		SubdivisionBG08,
		SubdivisionBG09,
		SubdivisionBG10,
		SubdivisionBG11,
		SubdivisionBG12,
		SubdivisionBG13,
		SubdivisionBG14,
		SubdivisionBG15,
		SubdivisionBG16,
		SubdivisionBG17,
		SubdivisionBG18,
		SubdivisionBG19,
		SubdivisionBG20,
		SubdivisionBG21,
		SubdivisionBG22,
		SubdivisionBG23,
		SubdivisionBG24,
		SubdivisionBG25,
		SubdivisionBG26,
		SubdivisionBG27,
		SubdivisionBG28,
		SubdivisionBH13,
		SubdivisionBH14,
		SubdivisionBH15,
		SubdivisionBH17,
		SubdivisionBIBM,
		SubdivisionBICI,
		SubdivisionBIGI,
		SubdivisionBIKI,
		SubdivisionBIMW,
		SubdivisionBING,
		SubdivisionBIRM,
		SubdivisionBIRT,
		SubdivisionBIRY,
		SubdivisionBJAK,
		SubdivisionBJAL,
		SubdivisionBJAQ,
		SubdivisionBJBO,
		SubdivisionBJCO,
		SubdivisionBJDO,
		SubdivisionBJLI,
		SubdivisionBJMO,
		SubdivisionBJOU,
		SubdivisionBJPL,
		SubdivisionBJZO,
		SubdivisionBNBE,
		SubdivisionBNBM,
		SubdivisionBNTE,
		SubdivisionBNTU,
		SubdivisionBOB,
		SubdivisionBOC,
		SubdivisionBOH,
		SubdivisionBOL,
		SubdivisionBON,
		SubdivisionBOO,
		SubdivisionBOP,
		SubdivisionBOS,
		SubdivisionBOT,
		SubdivisionBQBO,
		SubdivisionBQSA,
		SubdivisionBQSE,
		SubdivisionBRAC,
		SubdivisionBRAL,
		SubdivisionBRAM,
		SubdivisionBRAP,
		SubdivisionBRBA,
		SubdivisionBRCE,
		SubdivisionBRDF,
		SubdivisionBRES,
		SubdivisionBRGO,
		SubdivisionBRMA,
		SubdivisionBRMG,
		SubdivisionBRMS,
		SubdivisionBRMT,
		SubdivisionBRPA,
		SubdivisionBRPB,
		SubdivisionBRPE,
		SubdivisionBRPI,
		SubdivisionBRPR,
		SubdivisionBRRJ,
		SubdivisionBRRN,
		SubdivisionBRRO,
		SubdivisionBRRR,
		SubdivisionBRRS,
		SubdivisionBRSC,
		SubdivisionBRSE,
		SubdivisionBRSP,
		SubdivisionBRTO,
		SubdivisionBSCO,
		SubdivisionBSFP,
		SubdivisionBSLI,
		SubdivisionBSNO,
		SubdivisionBSNP,
		SubdivisionBSNS,
		SubdivisionBSSE,
		SubdivisionBT11,
		SubdivisionBT12,
		SubdivisionBT13,
		SubdivisionBT14,
		SubdivisionBT15,
		SubdivisionBT21,
		SubdivisionBT23,
		SubdivisionBT24,
		SubdivisionBT32,
		SubdivisionBT33,
		SubdivisionBT41,
		SubdivisionBT42,
		SubdivisionBT43,
		SubdivisionBT44,
		SubdivisionBT45,
		SubdivisionBTGA,
		SubdivisionBWCE,
		SubdivisionBWCH,
		SubdivisionBWKG,
		SubdivisionBWKL,
		SubdivisionBWKW,
		SubdivisionBWNE,
		SubdivisionBWNW,
		SubdivisionBWSE,
		SubdivisionBWSO,
		SubdivisionBYBR,
		SubdivisionBYHM,
		SubdivisionBYHO,
		SubdivisionBYHR,
		SubdivisionBYMA,
		SubdivisionBYMI,
		SubdivisionBYVI,
		SubdivisionBZBZ,
		SubdivisionBZCY,
		SubdivisionBZCZL,
		SubdivisionBZOW,
		SubdivisionBZSC,
		SubdivisionBZTOL,
		SubdivisionCAAB,
		SubdivisionCABC,
		SubdivisionCAMB,
		SubdivisionCANB,
		SubdivisionCANL,
		SubdivisionCANS,
		SubdivisionCANT,
		SubdivisionCANU,
		SubdivisionCAON,
		SubdivisionCAPE,
		SubdivisionCAQC,
		SubdivisionCASK,
		SubdivisionCAYT,
		SubdivisionCDEQ,
		SubdivisionCDHK,
		SubdivisionCDIT,
		SubdivisionCDKC,
		SubdivisionCDKE,
		SubdivisionCDKL,
		SubdivisionCDKN,
		SubdivisionCDLU,
		SubdivisionCDNK,
		SubdivisionCDSA,
		SubdivisionCDSK,
		SubdivisionCDTA,
		SubdivisionCDTO,
		SubdivisionCFAC,
		SubdivisionCFBGF,
		SubdivisionCFNM,
		SubdivisionCFOP,
		SubdivisionCG13,
		SubdivisionCG16,
		SubdivisionCGBZV,
		SubdivisionCHAG,
		SubdivisionCHAI,
		SubdivisionCHAR,
		SubdivisionCHBE,
		SubdivisionCHBL,
		SubdivisionCHBS,
		SubdivisionCHFR,
		SubdivisionCHGE,
		SubdivisionCHGL,
		SubdivisionCHGR,
		SubdivisionCHJU,
		SubdivisionCHLU,
		SubdivisionCHNE,
		SubdivisionCHNW,
		SubdivisionCHOW,
		SubdivisionCHSG,
		SubdivisionCHSH,
		SubdivisionCHSO,
		SubdivisionCHSZ,
		SubdivisionCHTG,
		SubdivisionCHTI,
		SubdivisionCHUR,
		SubdivisionCHVD,
		SubdivisionCHVS,
		SubdivisionCHZG,
		SubdivisionCHZH,
		SubdivisionCIAB,
		SubdivisionCIBS,
		SubdivisionCICM,
		SubdivisionCIDN,
		SubdivisionCIGD,
		SubdivisionCILC,
		SubdivisionCILG,
		SubdivisionCIMG,
		SubdivisionCISM,
		SubdivisionCISV,
		SubdivisionCIVB,
		SubdivisionCIWR,
		SubdivisionCIYM,
		SubdivisionCIZZ,
		SubdivisionCLAI,
		SubdivisionCLAN,
		SubdivisionCLAP,
		SubdivisionCLAR,
		SubdivisionCLAT,
		SubdivisionCLBI,
		SubdivisionCLCO,
		SubdivisionCLLI,
		SubdivisionCLLL,
		SubdivisionCLLR,
		SubdivisionCLMA,
		SubdivisionCLML,
		SubdivisionCLNB,
		SubdivisionCLRM,
		SubdivisionCLTA,
		SubdivisionCLVS,
		SubdivisionCMAD,
		SubdivisionCMCE,
		SubdivisionCMEN,
		SubdivisionCMES,
		SubdivisionCMLT,
		SubdivisionCMNO,
		SubdivisionCMNW,
		SubdivisionCMOU,
		SubdivisionCMSU,
		SubdivisionCMSW,
		SubdivisionCNAH,
		SubdivisionCNBJ,
		SubdivisionCNCQ,
		SubdivisionCNFJ,
		SubdivisionCNGD,
		SubdivisionCNGS,
		SubdivisionCNGX,
		SubdivisionCNGZ,
		SubdivisionCNHA,
		SubdivisionCNHB,
		SubdivisionCNHE,
		SubdivisionCNHI,
		SubdivisionCNHL,
		SubdivisionCNHN,
		SubdivisionCNJL,
		SubdivisionCNJS,
		SubdivisionCNJX,
		SubdivisionCNLN,
		SubdivisionCNNM,
		SubdivisionCNNX,
		SubdivisionCNQH,
		SubdivisionCNSC,
		SubdivisionCNSD,
		SubdivisionCNSH,
		SubdivisionCNSN,
		SubdivisionCNSX,
		SubdivisionCNTJ,
		SubdivisionCNXJ,
		SubdivisionCNXZ,
		SubdivisionCNYN,
		SubdivisionCNZJ,
		SubdivisionCOAMA,
		SubdivisionCOANT,
		SubdivisionCOARA,
		SubdivisionCOATL,
		SubdivisionCOBOL,
		SubdivisionCOBOY,
		SubdivisionCOCAL,
		SubdivisionCOCAQ,
		SubdivisionCOCAS,
		SubdivisionCOCAU,
		SubdivisionCOCES,
		SubdivisionCOCHO,
		SubdivisionCOCOR,
		SubdivisionCOCUN,
		SubdivisionCODC,
		SubdivisionCOGUA,
		SubdivisionCOGUV,
		SubdivisionCOHUI,
		SubdivisionCOLAG,
		SubdivisionCOMAG,
		SubdivisionCOMET,
		SubdivisionCONAR,
		SubdivisionCONSA,
		SubdivisionCOPUT,
		SubdivisionCOQUI,
		SubdivisionCORIS,
		SubdivisionCOSAN,
		SubdivisionCOSAP,
		SubdivisionCOSUC,
		SubdivisionCOTOL,
		SubdivisionCOVAC,
		SubdivisionCOVID,
		SubdivisionCRA,
		SubdivisionCRC,
		SubdivisionCRG,
		SubdivisionCRH,
		SubdivisionCRL,
		SubdivisionCRP,
		SubdivisionCRSJ,
		SubdivisionCU01,
		SubdivisionCU03,
		SubdivisionCU04,
		SubdivisionCU05,
		SubdivisionCU06,
		SubdivisionCU07,
		SubdivisionCU08,
		SubdivisionCU09,
		SubdivisionCU10,
		SubdivisionCU11,
		SubdivisionCU12,
		SubdivisionCU13,
		SubdivisionCU14,
		SubdivisionCU15,
		SubdivisionCU16,
		SubdivisionCVBR,
		SubdivisionCVBV,
		SubdivisionCVCR,
		SubdivisionCVMA,
		SubdivisionCVPN,
		SubdivisionCVPR,
		SubdivisionCVRG,
		SubdivisionCVRS,
		SubdivisionCVSF,
		SubdivisionCVSL,
		SubdivisionCVSV,
		SubdivisionCVTA,
		SubdivisionCY01,
		SubdivisionCY02,
		SubdivisionCY03,
		SubdivisionCY04,
		SubdivisionCY05,
		SubdivisionCY06,
		SubdivisionCZ10,
		SubdivisionCZ20,
		SubdivisionCZ31,
		SubdivisionCZ32,
		SubdivisionCZ41,
		SubdivisionCZ42,
		SubdivisionCZ51,
		SubdivisionCZ52,
		SubdivisionCZ53,
		SubdivisionCZ63,
		SubdivisionCZ64,
		SubdivisionCZ71,
		SubdivisionCZ72,
		SubdivisionCZ80,
		SubdivisionDEBB,
		SubdivisionDEBE,
		SubdivisionDEBW,
		SubdivisionDEBY,
		SubdivisionDEHB,
		SubdivisionDEHE,
		SubdivisionDEHH,
		SubdivisionDEMV,
		SubdivisionDENI,
		SubdivisionDENW,
		SubdivisionDERP,
		SubdivisionDESH,
		SubdivisionDESL,
		SubdivisionDESN,
		SubdivisionDEST,
		SubdivisionDETH,
		SubdivisionDJDJ,
		SubdivisionDK81,
		SubdivisionDK82,
		SubdivisionDK83,
		SubdivisionDK84,
		SubdivisionDK85,
		SubdivisionDM02,
		SubdivisionDM04,
		SubdivisionDM05,
		SubdivisionDM09,
		SubdivisionDM10,
		SubdivisionDO01,
		SubdivisionDO02,
		SubdivisionDO03,
		SubdivisionDO04,
		SubdivisionDO05,
		SubdivisionDO06,
		SubdivisionDO07,
		SubdivisionDO08,
		SubdivisionDO09,
		SubdivisionDO10,
		SubdivisionDO11,
		SubdivisionDO12,
		SubdivisionDO13,
		SubdivisionDO14,
		SubdivisionDO15,
		SubdivisionDO17,
		SubdivisionDO18,
		SubdivisionDO19,
		SubdivisionDO20,
		SubdivisionDO21,
		SubdivisionDO22,
		SubdivisionDO23,
		SubdivisionDO24,
		SubdivisionDO25,
		SubdivisionDO26,
		SubdivisionDO27,
		SubdivisionDO28,
		SubdivisionDO29,
		SubdivisionDO30,
		SubdivisionDO31,
		SubdivisionDZ01,
		SubdivisionDZ02,
		SubdivisionDZ03,
		SubdivisionDZ04,
		SubdivisionDZ05,
		SubdivisionDZ06,
		SubdivisionDZ07,
		SubdivisionDZ08,
		SubdivisionDZ09,
		SubdivisionDZ10,
		SubdivisionDZ11,
		SubdivisionDZ12,
		SubdivisionDZ13,
		SubdivisionDZ14,
		SubdivisionDZ15,
		SubdivisionDZ16,
		SubdivisionDZ17,
		SubdivisionDZ18,
		SubdivisionDZ19,
		SubdivisionDZ20,
		SubdivisionDZ21,
		SubdivisionDZ22,
		SubdivisionDZ23,
		SubdivisionDZ24,
		SubdivisionDZ25,
		SubdivisionDZ26,
		SubdivisionDZ27,
		SubdivisionDZ28,
		SubdivisionDZ29,
		SubdivisionDZ30,
		SubdivisionDZ31,
		SubdivisionDZ32,
		SubdivisionDZ33,
		SubdivisionDZ34,
		SubdivisionDZ35,
		SubdivisionDZ36,
		SubdivisionDZ37,
		SubdivisionDZ38,
		SubdivisionDZ39,
		SubdivisionDZ40,
		SubdivisionDZ41,
		SubdivisionDZ42,
		SubdivisionDZ43,
		SubdivisionDZ44,
		SubdivisionDZ45,
		SubdivisionDZ46,
		SubdivisionDZ47,
		SubdivisionDZ48,
		SubdivisionECA,
		SubdivisionECB,
		SubdivisionECC,
		SubdivisionECD,
		SubdivisionECE,
		SubdivisionECF,
		SubdivisionECG,
		SubdivisionECH,
		SubdivisionECI,
		SubdivisionECL,
		SubdivisionECM,
		SubdivisionECN,
		SubdivisionECO,
		SubdivisionECP,
		SubdivisionECR,
		SubdivisionECS,
		SubdivisionECSD,
		SubdivisionECSE,
		SubdivisionECT,
		SubdivisionECU,
		SubdivisionECW,
		SubdivisionECX,
		SubdivisionECY,
		SubdivisionECZ,
		SubdivisionEE37,
		SubdivisionEE39,
		SubdivisionEE44,
		SubdivisionEE49,
		SubdivisionEE51,
		SubdivisionEE57,
		SubdivisionEE59,
		SubdivisionEE65,
		SubdivisionEE67,
		SubdivisionEE70,
		SubdivisionEE74,
		SubdivisionEE78,
		SubdivisionEE82,
		SubdivisionEE84,
		SubdivisionEE86,
		SubdivisionEGALX,
		SubdivisionEGASN,
		SubdivisionEGAST,
		SubdivisionEGBA,
		SubdivisionEGBH,
		SubdivisionEGBNS,
		SubdivisionEGC,
		SubdivisionEGDK,
		SubdivisionEGDT,
		SubdivisionEGFYM,
		SubdivisionEGGH,
		SubdivisionEGGZ,
		SubdivisionEGIS,
		SubdivisionEGJS,
		SubdivisionEGKB,
		SubdivisionEGKFS,
		SubdivisionEGKN,
		SubdivisionEGLX,
		SubdivisionEGMN,
		SubdivisionEGMNF,
		SubdivisionEGMT,
		SubdivisionEGPTS,
		SubdivisionEGSHG,
		SubdivisionEGSHR,
		SubdivisionEGSIN,
		SubdivisionEGSUZ,
		SubdivisionEGWAD,
		SubdivisionERMA,
		SubdivisionESAN,
		SubdivisionESAR,
		SubdivisionESAS,
		SubdivisionESCB,
		SubdivisionESCE,
		SubdivisionESCL,
		SubdivisionESCM,
		SubdivisionESCN,
		SubdivisionESCT,
		SubdivisionESEX,
		SubdivisionESGA,
		SubdivisionESIB,
		SubdivisionESMC,
		SubdivisionESMD,
		SubdivisionESML,
		SubdivisionESNC,
		SubdivisionESPV,
		SubdivisionESRI,
		SubdivisionESVC,
		SubdivisionETAA,
		SubdivisionETAF,
		SubdivisionETAM,
		SubdivisionETBE,
		SubdivisionETDD,
		SubdivisionETHA,
		SubdivisionETOR,
		SubdivisionETSN,
		SubdivisionETSO,
		SubdivisionETTI,
		SubdivisionFI02,
		SubdivisionFI03,
		SubdivisionFI04,
		SubdivisionFI05,
		SubdivisionFI06,
		SubdivisionFI07,
		SubdivisionFI08,
		SubdivisionFI09,
		SubdivisionFI10,
		SubdivisionFI11,
		SubdivisionFI12,
		SubdivisionFI13,
		SubdivisionFI14,
		SubdivisionFI15,
		SubdivisionFI16,
		SubdivisionFI17,
		SubdivisionFI18,
		SubdivisionFI19,
		SubdivisionFJC,
		SubdivisionFJE,
		SubdivisionFJN,
		SubdivisionFJR,
		SubdivisionFJW,
		SubdivisionFMKSA,
		SubdivisionFMPNI,
		SubdivisionFMTRK,
		SubdivisionFMYAP,
		SubdivisionFRARA,
		SubdivisionFRBFC,
		SubdivisionFRBRE,
		SubdivisionFRCOR,
		SubdivisionFRCVL,
		SubdivisionFRGES,
		SubdivisionFRHDF,
		SubdivisionFRIDF,
		SubdivisionFRNAQ,
		SubdivisionFRNOR,
		SubdivisionFROCC,
		SubdivisionFRPAC,
		SubdivisionFRPDL,
		SubdivisionGA1,
		SubdivisionGA2,
		SubdivisionGA4,
		SubdivisionGA8,
		SubdivisionGA9,
		SubdivisionGBENG,
		SubdivisionGBNIR,
		SubdivisionGBSCT,
		SubdivisionGBWLS,
		SubdivisionGD01,
		SubdivisionGD02,
		SubdivisionGD03,
		SubdivisionGD04,
		SubdivisionGD05,
		SubdivisionGD10,
		SubdivisionGEAB,
		SubdivisionGEAJ,
		SubdivisionGEGU,
		SubdivisionGEIM,
		SubdivisionGEKA,
		SubdivisionGEKK,
		SubdivisionGEMM,
		SubdivisionGERL,
		SubdivisionGESJ,
		SubdivisionGESK,
		SubdivisionGESZ,
		SubdivisionGETB,
		SubdivisionGHAA,
		SubdivisionGHAF,
		SubdivisionGHAH,
		SubdivisionGHBE,
		SubdivisionGHBO,
		SubdivisionGHCP,
		SubdivisionGHEP,
		SubdivisionGHNP,
		SubdivisionGHTV,
		SubdivisionGHUE,
		SubdivisionGHWP,
		SubdivisionGLAV,
		SubdivisionGLKU,
		SubdivisionGLQE,
		SubdivisionGLQT,
		SubdivisionGLSM,
		SubdivisionGMB,
		SubdivisionGML,
		SubdivisionGMM,
		SubdivisionGMN,
		SubdivisionGMU,
		SubdivisionGMW,
		SubdivisionGNB,
		SubdivisionGNBF,
		SubdivisionGNC,
		SubdivisionGNCO,
		SubdivisionGND,
		SubdivisionGNDB,
		SubdivisionGNK,
		SubdivisionGNN,
		SubdivisionGNSI,
		SubdivisionGQBN,
		SubdivisionGQBS,
		SubdivisionGQLI,
		SubdivisionGQWN,
		SubdivisionGR69,
		SubdivisionGRA,
		SubdivisionGRB,
		SubdivisionGRC,
		SubdivisionGRD,
		SubdivisionGRE,
		SubdivisionGRF,
		SubdivisionGRG,
		SubdivisionGRH,
		SubdivisionGRI,
		SubdivisionGRJ,
		SubdivisionGRK,
		SubdivisionGRL,
		SubdivisionGRM,
		SubdivisionGTAV,
		SubdivisionGTBV,
		SubdivisionGTCM,
		SubdivisionGTCQ,
		SubdivisionGTES,
		SubdivisionGTGU,
		SubdivisionGTHU,
		SubdivisionGTIZ,
		SubdivisionGTJA,
		SubdivisionGTJU,
		SubdivisionGTPE,
		SubdivisionGTPR,
		SubdivisionGTQC,
		SubdivisionGTQZ,
		SubdivisionGTRE,
		SubdivisionGTSA,
		SubdivisionGTSM,
		SubdivisionGTSO,
		SubdivisionGTSR,
		SubdivisionGTSU,
		SubdivisionGTTO,
		SubdivisionGTZA,
		SubdivisionGWBS,
		SubdivisionGWGA,
		SubdivisionGYBA,
		SubdivisionGYCU,
		SubdivisionGYDE,
		SubdivisionGYEB,
		SubdivisionGYES,
		SubdivisionGYMA,
		SubdivisionGYPT,
		SubdivisionGYUD,
		SubdivisionHNAT,
		SubdivisionHNCH,
		SubdivisionHNCL,
		SubdivisionHNCM,
		SubdivisionHNCP,
		SubdivisionHNCR,
		SubdivisionHNEP,
		SubdivisionHNFM,
		SubdivisionHNIB,
		SubdivisionHNIN,
		SubdivisionHNLE,
		SubdivisionHNLP,
		SubdivisionHNOC,
		SubdivisionHNOL,
		SubdivisionHNSB,
		SubdivisionHNVA,
		SubdivisionHNYO,
		SubdivisionHR01,
		SubdivisionHR02,
		SubdivisionHR03,
		SubdivisionHR04,
		SubdivisionHR05,
		SubdivisionHR06,
		SubdivisionHR07,
		SubdivisionHR08,
		SubdivisionHR09,
		SubdivisionHR10,
		SubdivisionHR11,
		SubdivisionHR12,
		SubdivisionHR13,
		SubdivisionHR14,
		SubdivisionHR15,
		SubdivisionHR16,
		SubdivisionHR17,
		SubdivisionHR18,
		SubdivisionHR19,
		SubdivisionHR20,
		SubdivisionHR21,
		SubdivisionHTAR,
		SubdivisionHTCE,
		SubdivisionHTND,
		SubdivisionHTOU,
		SubdivisionHTSD,
		SubdivisionHTSE,
		SubdivisionHUBA,
		SubdivisionHUBE,
		SubdivisionHUBK,
		SubdivisionHUBU,
		SubdivisionHUBZ,
		SubdivisionHUCS,
		SubdivisionHUFE,
		SubdivisionHUGS,
		SubdivisionHUHB,
		SubdivisionHUHE,
		SubdivisionHUJN,
		SubdivisionHUKE,
		SubdivisionHUNO,
		SubdivisionHUPE,
		SubdivisionHUSO,
		SubdivisionHUSZ,
		SubdivisionHUTO,
		SubdivisionHUVA,
		SubdivisionHUVE,
		SubdivisionHUZA,
		SubdivisionIDAC,
		SubdivisionIDBA,
		SubdivisionIDBB,
		SubdivisionIDBE,
		SubdivisionIDBT,
		SubdivisionIDGO,
		SubdivisionIDJA,
		SubdivisionIDJB,
		SubdivisionIDJI,
		SubdivisionIDJK,
		SubdivisionIDJT,
		SubdivisionIDKB,
		SubdivisionIDKI,
		SubdivisionIDKR,
		SubdivisionIDKS,
		SubdivisionIDKT,
		SubdivisionIDKU,
		SubdivisionIDLA,
		SubdivisionIDML,
		SubdivisionIDMU,
		SubdivisionIDNB,
		SubdivisionIDNT,
		SubdivisionIDPB,
		SubdivisionIDPP,
		SubdivisionIDRI,
		SubdivisionIDSA,
		SubdivisionIDSB,
		SubdivisionIDSG,
		SubdivisionIDSN,
		SubdivisionIDSR,
		SubdivisionIDSS,
		SubdivisionIDST,
		SubdivisionIDSU,
		SubdivisionIDYO,
		SubdivisionIECE,
		SubdivisionIECN,
		SubdivisionIECO,
		SubdivisionIECW,
		SubdivisionIED,
		SubdivisionIEDL,
		SubdivisionIEG,
		SubdivisionIEKE,
		SubdivisionIEKK,
		SubdivisionIEKY,
		SubdivisionIELD,
		SubdivisionIELH,
		SubdivisionIELK,
		SubdivisionIELM,
		SubdivisionIELS,
		SubdivisionIEMH,
		SubdivisionIEMN,
		SubdivisionIEMO,
		SubdivisionIEOY,
		SubdivisionIERN,
		SubdivisionIESO,
		SubdivisionIETA,
		SubdivisionIEWD,
		SubdivisionIEWH,
		SubdivisionIEWW,
		SubdivisionIEWX,
		SubdivisionILD,
		SubdivisionILHA,
		SubdivisionILJM,
		SubdivisionILM,
		SubdivisionILTA,
		SubdivisionILZ,
		SubdivisionINAN,
		SubdivisionINAP,
		SubdivisionINAR,
		SubdivisionINAS,
		SubdivisionINBR,
		SubdivisionINCH,
		SubdivisionINCT,
		SubdivisionINDH,
		SubdivisionINDL,
		SubdivisionINDN,
		SubdivisionINGA,
		SubdivisionINGJ,
		SubdivisionINHP,
		SubdivisionINHR,
		SubdivisionINJH,
		SubdivisionINJK,
		SubdivisionINKA,
		SubdivisionINKL,
		SubdivisionINLD,
		SubdivisionINMH,
		SubdivisionINML,
		SubdivisionINMN,
		SubdivisionINMP,
		SubdivisionINMZ,
		SubdivisionINNL,
		SubdivisionINOR,
		SubdivisionINPB,
		SubdivisionINPY,
		SubdivisionINRJ,
		SubdivisionINSK,
		SubdivisionINTG,
		SubdivisionINTN,
		SubdivisionINTR,
		SubdivisionINUP,
		SubdivisionINUT,
		SubdivisionINWB,
		SubdivisionIQAN,
		SubdivisionIQAR,
		SubdivisionIQBA,
		SubdivisionIQBB,
		SubdivisionIQBG,
		SubdivisionIQDA,
		SubdivisionIQDI,
		SubdivisionIQDQ,
		SubdivisionIQKA,
		SubdivisionIQKI,
		SubdivisionIQMA,
		SubdivisionIQMU,
		SubdivisionIQNA,
		SubdivisionIQNI,
		SubdivisionIQQA,
		SubdivisionIQSD,
		SubdivisionIQSU,
		SubdivisionIQWA,
		SubdivisionIR01,
		SubdivisionIR02,
		SubdivisionIR03,
		SubdivisionIR04,
		SubdivisionIR05,
		SubdivisionIR06,
		SubdivisionIR07,
		SubdivisionIR08,
		SubdivisionIR10,
		SubdivisionIR11,
		SubdivisionIR12,
		SubdivisionIR13,
		SubdivisionIR14,
		SubdivisionIR15,
		SubdivisionIR16,
		SubdivisionIR17,
		SubdivisionIR18,
		SubdivisionIR19,
		SubdivisionIR20,
		SubdivisionIR21,
		SubdivisionIR22,
		SubdivisionIR23,
		SubdivisionIR24,
		SubdivisionIR25,
		SubdivisionIR26,
		SubdivisionIR27,
		SubdivisionIR28,
		SubdivisionIR29,
		SubdivisionIR30,
		SubdivisionIR31,
		SubdivisionIR32,
		SubdivisionIS1,
		SubdivisionIS2,
		SubdivisionIS3,
		SubdivisionIS4,
		SubdivisionIS5,
		SubdivisionIS6,
		SubdivisionIS7,
		SubdivisionIS8,
		SubdivisionIT21,
		SubdivisionIT23,
		SubdivisionIT25,
		SubdivisionIT32,
		SubdivisionIT34,
		SubdivisionIT36,
		SubdivisionIT42,
		SubdivisionIT45,
		SubdivisionIT52,
		SubdivisionIT55,
		SubdivisionIT57,
		SubdivisionIT62,
		SubdivisionIT65,
		SubdivisionIT67,
		SubdivisionIT72,
		SubdivisionIT75,
		SubdivisionIT77,
		SubdivisionIT78,
		SubdivisionIT82,
		SubdivisionIT88,
		SubdivisionJM01,
		SubdivisionJM02,
		SubdivisionJM03,
		SubdivisionJM04,
		SubdivisionJM05,
		SubdivisionJM06,
		SubdivisionJM07,
		SubdivisionJM08,
		SubdivisionJM09,
		SubdivisionJM10,
		SubdivisionJM11,
		SubdivisionJM12,
		SubdivisionJM13,
		SubdivisionJM14,
		SubdivisionJOAJ,
		SubdivisionJOAM,
		SubdivisionJOAQ,
		SubdivisionJOAZ,
		SubdivisionJOBA,
		SubdivisionJOIR,
		SubdivisionJOJA,
		SubdivisionJOKA,
		SubdivisionJOMA,
		SubdivisionJOMD,
		SubdivisionJOMN,
		SubdivisionJP01,
		SubdivisionJP02,
		SubdivisionJP03,
		SubdivisionJP04,
		SubdivisionJP05,
		SubdivisionJP06,
		SubdivisionJP07,
		SubdivisionJP08,
		SubdivisionJP09,
		SubdivisionJP10,
		SubdivisionJP11,
		SubdivisionJP12,
		SubdivisionJP13,
		SubdivisionJP14,
		SubdivisionJP15,
		SubdivisionJP16,
		SubdivisionJP17,
		SubdivisionJP18,
		SubdivisionJP19,
		SubdivisionJP20,
		SubdivisionJP21,
		SubdivisionJP22,
		SubdivisionJP23,
		SubdivisionJP24,
		SubdivisionJP25,
		SubdivisionJP26,
		SubdivisionJP27,
		SubdivisionJP28,
		SubdivisionJP29,
		SubdivisionJP30,
		SubdivisionJP31,
		SubdivisionJP32,
		SubdivisionJP33,
		SubdivisionJP34,
		SubdivisionJP35,
		SubdivisionJP36,
		SubdivisionJP37,
		SubdivisionJP38,
		SubdivisionJP39,
		SubdivisionJP40,
		SubdivisionJP41,
		SubdivisionJP42,
		SubdivisionJP43,
		SubdivisionJP44,
		SubdivisionJP45,
		SubdivisionJP46,
		SubdivisionJP47,
		SubdivisionKE01,
		SubdivisionKE02,
		SubdivisionKE03,
		SubdivisionKE04,
		SubdivisionKE05,
		SubdivisionKE06,
		SubdivisionKE07,
		SubdivisionKE08,
		SubdivisionKE09,
		SubdivisionKE10,
		SubdivisionKE11,
		SubdivisionKE12,
		SubdivisionKE13,
		SubdivisionKE14,
		SubdivisionKE15,
		SubdivisionKE16,
		SubdivisionKE17,
		SubdivisionKE18,
		SubdivisionKE19,
		SubdivisionKE20,
		SubdivisionKE21,
		SubdivisionKE22,
		SubdivisionKE23,
		SubdivisionKE24,
		SubdivisionKE25,
		SubdivisionKE26,
		SubdivisionKE27,
		SubdivisionKE28,
		SubdivisionKE29,
		SubdivisionKE30,
		SubdivisionKE31,
		SubdivisionKE32,
		SubdivisionKE33,
		SubdivisionKE34,
		SubdivisionKE35,
		SubdivisionKE36,
		SubdivisionKE38,
		SubdivisionKE39,
		SubdivisionKE41,
		SubdivisionKE42,
		SubdivisionKE43,
		SubdivisionKE44,
		SubdivisionKE46,
		SubdivisionKGB,
		SubdivisionKGC,
		SubdivisionKGGB,
		SubdivisionKGGO,
		SubdivisionKGJ,
		SubdivisionKGN,
		SubdivisionKGT,
		SubdivisionKGY,
		SubdivisionKH1,
		SubdivisionKH10,
		SubdivisionKH11,
		SubdivisionKH12,
		SubdivisionKH14,
		SubdivisionKH15,
		SubdivisionKH17,
		SubdivisionKH18,
		SubdivisionKH19,
		SubdivisionKH2,
		SubdivisionKH20,
		SubdivisionKH21,
		SubdivisionKH23,
		SubdivisionKH24,
		SubdivisionKH3,
		SubdivisionKH4,
		SubdivisionKH5,
		SubdivisionKH6,
		SubdivisionKH7,
		SubdivisionKH8,
		SubdivisionKH9,
		SubdivisionKIG,
		SubdivisionKMG,
		SubdivisionKN02,
		SubdivisionKN03,
		SubdivisionKN05,
		SubdivisionKN06,
		SubdivisionKN07,
		SubdivisionKN08,
		SubdivisionKN09,
		SubdivisionKN10,
		SubdivisionKN11,
		SubdivisionKN12,
		SubdivisionKN13,
		SubdivisionKP01,
		SubdivisionKR11,
		SubdivisionKR26,
		SubdivisionKR27,
		SubdivisionKR28,
		SubdivisionKR29,
		SubdivisionKR30,
		SubdivisionKR31,
		SubdivisionKR41,
		SubdivisionKR42,
		SubdivisionKR43,
		SubdivisionKR44,
		SubdivisionKR45,
		SubdivisionKR46,
		SubdivisionKR47,
		SubdivisionKR48,
		SubdivisionKR49,
		SubdivisionKWAH,
		SubdivisionKWFA,
		SubdivisionKWHA,
		SubdivisionKWJA,
		SubdivisionKWKU,
		SubdivisionKWMU,
		SubdivisionKZAKM,
		SubdivisionKZAKT,
		SubdivisionKZALA,
		SubdivisionKZALM,
		SubdivisionKZAST,
		SubdivisionKZATY,
		SubdivisionKZKAR,
		SubdivisionKZKUS,
		SubdivisionKZKZY,
		SubdivisionKZMAN,
		SubdivisionKZPAV,
		SubdivisionKZSEV,
		SubdivisionKZSHY,
		SubdivisionKZVOS,
		SubdivisionKZYUZ,
		SubdivisionKZZAP,
		SubdivisionKZZHA,
		SubdivisionLABL,
		SubdivisionLACH,
		SubdivisionLAKH,
		SubdivisionLALP,
		SubdivisionLAOU,
		SubdivisionLAPH,
		SubdivisionLASV,
		SubdivisionLAVI,
		SubdivisionLAXA,
		SubdivisionLAXE,
		SubdivisionLAXI,
		SubdivisionLBAK,
		SubdivisionLBAS,
		SubdivisionLBBA,
		SubdivisionLBBH,
		SubdivisionLBBI,
		SubdivisionLBJA,
		SubdivisionLBJL,
		SubdivisionLBNA,
		SubdivisionLC01,
		SubdivisionLC02,
		SubdivisionLC05,
		SubdivisionLC06,
		SubdivisionLC07,
		SubdivisionLC10,
		SubdivisionLC11,
		SubdivisionLI01,
		SubdivisionLI02,
		SubdivisionLI03,
		SubdivisionLI04,
		SubdivisionLI05,
		SubdivisionLI06,
		SubdivisionLI07,
		SubdivisionLI09,
		SubdivisionLI10,
		SubdivisionLI11,
		SubdivisionLK1,
		SubdivisionLK2,
		SubdivisionLK3,
		SubdivisionLK4,
		SubdivisionLK5,
		SubdivisionLK6,
		SubdivisionLK7,
		SubdivisionLK8,
		SubdivisionLK9,
		SubdivisionLRBM,
		SubdivisionLRGB,
		SubdivisionLRGG,
		SubdivisionLRMG,
		SubdivisionLRMO,
		SubdivisionLRNI,
		SubdivisionLRSI,
		SubdivisionLSA,
		SubdivisionLSC,
		SubdivisionLSE,
		SubdivisionLSG,
		SubdivisionLSJ,
		SubdivisionLSK,
		SubdivisionLTAL,
		SubdivisionLTKL,
		SubdivisionLTKU,
		SubdivisionLTMR,
		SubdivisionLTPN,
		SubdivisionLTSA,
		SubdivisionLTTA,
		SubdivisionLTTE,
		SubdivisionLTUT,
		SubdivisionLTVL,
		SubdivisionLUCA,
		SubdivisionLUCL,
		SubdivisionLUDI,
		SubdivisionLUEC,
		SubdivisionLUES,
		SubdivisionLUGR,
		SubdivisionLULU,
		SubdivisionLUME,
		SubdivisionLURD,
		SubdivisionLURM,
		SubdivisionLUVD,
		SubdivisionLUWI,
		SubdivisionLV001,
		SubdivisionLV002,
		SubdivisionLV003,
		SubdivisionLV005,
		SubdivisionLV007,
		SubdivisionLV010,
		SubdivisionLV011,
		SubdivisionLV012,
		SubdivisionLV015,
		SubdivisionLV016,
		SubdivisionLV017,
		SubdivisionLV018,
		SubdivisionLV020,
		SubdivisionLV021,
		SubdivisionLV022,
		SubdivisionLV024,
		SubdivisionLV025,
		SubdivisionLV026,
		SubdivisionLV027,
		SubdivisionLV029,
		SubdivisionLV030,
		SubdivisionLV033,
		SubdivisionLV034,
		SubdivisionLV035,
		SubdivisionLV037,
		SubdivisionLV038,
		SubdivisionLV039,
		SubdivisionLV040,
		SubdivisionLV041,
		SubdivisionLV042,
		SubdivisionLV043,
		SubdivisionLV046,
		SubdivisionLV047,
		SubdivisionLV050,
		SubdivisionLV052,
		SubdivisionLV053,
		SubdivisionLV054,
		SubdivisionLV056,
		SubdivisionLV057,
		SubdivisionLV058,
		SubdivisionLV059,
		SubdivisionLV061,
		SubdivisionLV064,
		SubdivisionLV067,
		SubdivisionLV068,
		SubdivisionLV069,
		SubdivisionLV073,
		SubdivisionLV075,
		SubdivisionLV077,
		SubdivisionLV078,
		SubdivisionLV079,
		SubdivisionLV083,
		SubdivisionLV086,
		SubdivisionLV087,
		SubdivisionLV088,
		SubdivisionLV089,
		SubdivisionLV090,
		SubdivisionLV091,
		SubdivisionLV093,
		SubdivisionLV094,
		SubdivisionLV095,
		SubdivisionLV097,
		SubdivisionLV099,
		SubdivisionLV100,
		SubdivisionLV101,
		SubdivisionLV103,
		SubdivisionLV105,
		SubdivisionLV106,
		SubdivisionLV110,
		SubdivisionLVJEL,
		SubdivisionLVJUR,
		SubdivisionLVLPX,
		SubdivisionLVRIX,
		SubdivisionLVVMR,
		SubdivisionLYBA,
		SubdivisionLYBU,
		SubdivisionLYDR,
		SubdivisionLYJA,
		SubdivisionLYJG,
		SubdivisionLYJI,
		SubdivisionLYJU,
		SubdivisionLYMB,
		SubdivisionLYMI,
		SubdivisionLYMJ,
		SubdivisionLYMQ,
		SubdivisionLYNL,
		SubdivisionLYNQ,
		SubdivisionLYSB,
		SubdivisionLYSR,
		SubdivisionLYTB,
		SubdivisionLYWA,
		SubdivisionLYZA,
		SubdivisionMA01,
		SubdivisionMA02,
		SubdivisionMA03,
		SubdivisionMA04,
		SubdivisionMA05,
		SubdivisionMA06,
		SubdivisionMA07,
		SubdivisionMA08,
		SubdivisionMA09,
		SubdivisionMA10,
		SubdivisionMA11,
		SubdivisionMCCO,
		SubdivisionMCFO,
		SubdivisionMCMC,
		SubdivisionMCMO,
		SubdivisionMCSR,
		SubdivisionMDAN,
		SubdivisionMDBA,
		SubdivisionMDBD,
		SubdivisionMDBR,
		SubdivisionMDBS,
		SubdivisionMDCA,
		SubdivisionMDCL,
		SubdivisionMDCM,
		SubdivisionMDCR,
		SubdivisionMDCS,
		SubdivisionMDCT,
		SubdivisionMDCU,
		SubdivisionMDDO,
		SubdivisionMDDR,
		SubdivisionMDDU,
		SubdivisionMDED,
		SubdivisionMDFA,
		SubdivisionMDFL,
		SubdivisionMDGA,
		SubdivisionMDGL,
		SubdivisionMDHI,
		SubdivisionMDIA,
		SubdivisionMDLE,
		SubdivisionMDNI,
		SubdivisionMDOC,
		SubdivisionMDOR,
		SubdivisionMDRE,
		SubdivisionMDRI,
		SubdivisionMDSD,
		SubdivisionMDSI,
		SubdivisionMDSN,
		SubdivisionMDSO,
		SubdivisionMDST,
		SubdivisionMDSV,
		SubdivisionMDTA,
		SubdivisionMDTE,
		SubdivisionMDUN,
		SubdivisionME02,
		SubdivisionME03,
		SubdivisionME04,
		SubdivisionME05,
		SubdivisionME06,
		SubdivisionME07,
		SubdivisionME08,
		SubdivisionME09,
		SubdivisionME10,
		SubdivisionME12,
		SubdivisionME13,
		SubdivisionME14,
		SubdivisionME15,
		SubdivisionME16,
		SubdivisionME17,
		SubdivisionME19,
		SubdivisionME20,
		SubdivisionME21,
		SubdivisionME24,
		SubdivisionMGA,
		SubdivisionMGD,
		SubdivisionMGF,
		SubdivisionMGM,
		SubdivisionMGT,
		SubdivisionMGU,
		SubdivisionMHKWA,
		SubdivisionMHMAJ,
		SubdivisionMK101,
		SubdivisionMK104,
		SubdivisionMK106,
		SubdivisionMK107,
		SubdivisionMK108,
		SubdivisionMK109,
		SubdivisionMK201,
		SubdivisionMK202,
		SubdivisionMK203,
		SubdivisionMK205,
		SubdivisionMK206,
		SubdivisionMK207,
		SubdivisionMK208,
		SubdivisionMK209,
		SubdivisionMK210,
		SubdivisionMK211,
		SubdivisionMK303,
		SubdivisionMK307,
		SubdivisionMK310,
		SubdivisionMK311,
		SubdivisionMK312,
		SubdivisionMK313,
		SubdivisionMK401,
		SubdivisionMK402,
		SubdivisionMK403,
		SubdivisionMK404,
		SubdivisionMK405,
		SubdivisionMK406,
		SubdivisionMK408,
		SubdivisionMK409,
		SubdivisionMK410,
		SubdivisionMK501,
		SubdivisionMK503,
		SubdivisionMK504,
		SubdivisionMK505,
		SubdivisionMK506,
		SubdivisionMK507,
		SubdivisionMK508,
		SubdivisionMK509,
		SubdivisionMK601,
		SubdivisionMK602,
		SubdivisionMK603,
		SubdivisionMK604,
		SubdivisionMK605,
		SubdivisionMK606,
		SubdivisionMK607,
		SubdivisionMK608,
		SubdivisionMK609,
		SubdivisionMK701,
		SubdivisionMK702,
		SubdivisionMK703,
		SubdivisionMK704,
		SubdivisionMK705,
		SubdivisionMK802,
		SubdivisionMK803,
		SubdivisionMK804,
		SubdivisionMK806,
		SubdivisionMK807,
		SubdivisionMK809,
		SubdivisionMK810,
		SubdivisionMK811,
		SubdivisionMK812,
		SubdivisionMK813,
		SubdivisionMK814,
		SubdivisionMK817,
		SubdivisionML1,
		SubdivisionML2,
		SubdivisionML3,
		SubdivisionML4,
		SubdivisionML5,
		SubdivisionML6,
		SubdivisionML7,
		SubdivisionML8,
		SubdivisionMLBKO,
		SubdivisionMM01,
		SubdivisionMM02,
		SubdivisionMM03,
		SubdivisionMM04,
		SubdivisionMM05,
		SubdivisionMM06,
		SubdivisionMM07,
		SubdivisionMM11,
		SubdivisionMM12,
		SubdivisionMM13,
		SubdivisionMM15,
		SubdivisionMM16,
		SubdivisionMM17,
		SubdivisionMM18,
		SubdivisionMN035,
		SubdivisionMN037,
		SubdivisionMN047,
		SubdivisionMN049,
		SubdivisionMN055,
		SubdivisionMN061,
		SubdivisionMN065,
		SubdivisionMN071,
		SubdivisionMN1,
		SubdivisionMR04,
		SubdivisionMR06,
		SubdivisionMR08,
		SubdivisionMR11,
		SubdivisionMR12,
		SubdivisionMR13,
		SubdivisionMT01,
		SubdivisionMT02,
		SubdivisionMT03,
		SubdivisionMT04,
		SubdivisionMT05,
		SubdivisionMT06,
		SubdivisionMT07,
		SubdivisionMT08,
		SubdivisionMT09,
		SubdivisionMT10,
		SubdivisionMT11,
		SubdivisionMT12,
		SubdivisionMT14,
		SubdivisionMT15,
		SubdivisionMT16,
		SubdivisionMT17,
		SubdivisionMT18,
		SubdivisionMT19,
		SubdivisionMT20,
		SubdivisionMT21,
		SubdivisionMT22,
		SubdivisionMT23,
		SubdivisionMT24,
		SubdivisionMT25,
		SubdivisionMT26,
		SubdivisionMT27,
		SubdivisionMT28,
		SubdivisionMT29,
		SubdivisionMT30,
		SubdivisionMT31,
		SubdivisionMT32,
		SubdivisionMT33,
		SubdivisionMT34,
		SubdivisionMT35,
		SubdivisionMT36,
		SubdivisionMT37,
		SubdivisionMT38,
		SubdivisionMT39,
		SubdivisionMT40,
		SubdivisionMT41,
		SubdivisionMT42,
		SubdivisionMT43,
		SubdivisionMT45,
		SubdivisionMT46,
		SubdivisionMT48,
		SubdivisionMT49,
		SubdivisionMT51,
		SubdivisionMT52,
		SubdivisionMT53,
		SubdivisionMT54,
		SubdivisionMT55,
		SubdivisionMT56,
		SubdivisionMT57,
		SubdivisionMT58,
		SubdivisionMT59,
		SubdivisionMT60,
		SubdivisionMT61,
		SubdivisionMT62,
		SubdivisionMT63,
		SubdivisionMT64,
		SubdivisionMT65,
		SubdivisionMT67,
		SubdivisionMT68,
		SubdivisionMUBL,
		SubdivisionMUFL,
		SubdivisionMUGP,
		SubdivisionMUMO,
		SubdivisionMUPA,
		SubdivisionMUPL,
		SubdivisionMUPW,
		SubdivisionMURO,
		SubdivisionMURR,
		SubdivisionMUSA,
		SubdivisionMV00,
		SubdivisionMV01,
		SubdivisionMV03,
		SubdivisionMV04,
		SubdivisionMV05,
		SubdivisionMV07,
		SubdivisionMV12,
		SubdivisionMV13,
		SubdivisionMV17,
		SubdivisionMV20,
		SubdivisionMV25,
		SubdivisionMV28,
		SubdivisionMVMLE,
		SubdivisionMWBA,
		SubdivisionMWBL,
		SubdivisionMWCR,
		SubdivisionMWDE,
		SubdivisionMWDO,
		SubdivisionMWKR,
		SubdivisionMWLI,
		SubdivisionMWMG,
		SubdivisionMWMH,
		SubdivisionMWMZ,
		SubdivisionMWNK,
		SubdivisionMWSA,
		SubdivisionMWZO,
		SubdivisionMXAGU,
		SubdivisionMXBCN,
		SubdivisionMXBCS,
		SubdivisionMXCAM,
		SubdivisionMXCHH,
		SubdivisionMXCHP,
		SubdivisionMXCMX,
		SubdivisionMXCOA,
		SubdivisionMXCOL,
		SubdivisionMXDUR,
		SubdivisionMXGRO,
		SubdivisionMXGUA,
		SubdivisionMXHID,
		SubdivisionMXJAL,
		SubdivisionMXMEX,
		SubdivisionMXMIC,
		SubdivisionMXMOR,
		SubdivisionMXNAY,
		SubdivisionMXNLE,
		SubdivisionMXOAX,
		SubdivisionMXPUE,
		SubdivisionMXQUE,
		SubdivisionMXROO,
		SubdivisionMXSIN,
		SubdivisionMXSLP,
		SubdivisionMXSON,
		SubdivisionMXTAB,
		SubdivisionMXTAM,
		SubdivisionMXTLA,
		SubdivisionMXVER,
		SubdivisionMXYUC,
		SubdivisionMXZAC,
		SubdivisionMY01,
		SubdivisionMY02,
		SubdivisionMY03,
		SubdivisionMY04,
		SubdivisionMY05,
		SubdivisionMY06,
		SubdivisionMY07,
		SubdivisionMY08,
		SubdivisionMY09,
		SubdivisionMY10,
		SubdivisionMY11,
		SubdivisionMY12,
		SubdivisionMY13,
		SubdivisionMY14,
		SubdivisionMY15,
		SubdivisionMY16,
		SubdivisionMZA,
		SubdivisionMZB,
		SubdivisionMZG,
		SubdivisionMZI,
		SubdivisionMZL,
		SubdivisionMZN,
		SubdivisionMZP,
		SubdivisionMZQ,
		SubdivisionMZS,
		SubdivisionMZT,
		SubdivisionNACA,
		SubdivisionNAER,
		SubdivisionNAHA,
		SubdivisionNAKA,
		SubdivisionNAKE,
		SubdivisionNAKH,
		SubdivisionNAKU,
		SubdivisionNAOD,
		SubdivisionNAOH,
		SubdivisionNAON,
		SubdivisionNAOS,
		SubdivisionNAOT,
		SubdivisionNAOW,
		SubdivisionNE1,
		SubdivisionNE3,
		SubdivisionNE5,
		SubdivisionNE7,
		SubdivisionNE8,
		SubdivisionNGAB,
		SubdivisionNGAD,
		SubdivisionNGAK,
		SubdivisionNGAN,
		SubdivisionNGBA,
		SubdivisionNGBE,
		SubdivisionNGBO,
		SubdivisionNGBY,
		SubdivisionNGCR,
		SubdivisionNGDE,
		SubdivisionNGED,
		SubdivisionNGEK,
		SubdivisionNGEN,
		SubdivisionNGFC,
		SubdivisionNGGO,
		SubdivisionNGIM,
		SubdivisionNGJI,
		SubdivisionNGKD,
		SubdivisionNGKE,
		SubdivisionNGKN,
		SubdivisionNGKO,
		SubdivisionNGKT,
		SubdivisionNGKW,
		SubdivisionNGLA,
		SubdivisionNGNA,
		SubdivisionNGNI,
		SubdivisionNGOG,
		SubdivisionNGON,
		SubdivisionNGOS,
		SubdivisionNGOY,
		SubdivisionNGPL,
		SubdivisionNGRI,
		SubdivisionNGSO,
		SubdivisionNGTA,
		SubdivisionNGYO,
		SubdivisionNGZA,
		SubdivisionNIAN,
		SubdivisionNIAS,
		SubdivisionNIBO,
		SubdivisionNICA,
		SubdivisionNICI,
		SubdivisionNICO,
		SubdivisionNIES,
		SubdivisionNIGR,
		SubdivisionNIJI,
		SubdivisionNILE,
		SubdivisionNIMD,
		SubdivisionNIMN,
		SubdivisionNIMS,
		SubdivisionNIMT,
		SubdivisionNINS,
		SubdivisionNIRI,
		SubdivisionNISJ,
		SubdivisionNLDR,
		SubdivisionNLFL,
		SubdivisionNLFR,
		SubdivisionNLGE,
		SubdivisionNLGR,
		SubdivisionNLLI,
		SubdivisionNLNB,
		SubdivisionNLNH,
		SubdivisionNLOV,
		SubdivisionNLUT,
		SubdivisionNLZE,
		SubdivisionNLZH,
		SubdivisionNO03,
		SubdivisionNO11,
		SubdivisionNO15,
		SubdivisionNO18,
		SubdivisionNO30,
		SubdivisionNO34,
		SubdivisionNO38,
		SubdivisionNO42,
		SubdivisionNO46,
		SubdivisionNO50,
		SubdivisionNO54,
		SubdivisionNPBA,
		SubdivisionNPBH,
		SubdivisionNPDH,
		SubdivisionNPGA,
		SubdivisionNPJA,
		SubdivisionNPKA,
		SubdivisionNPKO,
		SubdivisionNPLU,
		SubdivisionNPMA,
		SubdivisionNPME,
		SubdivisionNPNA,
		SubdivisionNPRA,
		SubdivisionNPSA,
		SubdivisionNPSE,
		SubdivisionNR14,
		SubdivisionNZAUK,
		SubdivisionNZBOP,
		SubdivisionNZCAN,
		SubdivisionNZCIT,
		SubdivisionNZGIS,
		SubdivisionNZHKB,
		SubdivisionNZMBH,
		SubdivisionNZMWT,
		SubdivisionNZNSN,
		SubdivisionNZNTL,
		SubdivisionNZOTA,
		SubdivisionNZSTL,
		SubdivisionNZTAS,
		SubdivisionNZTKI,
		SubdivisionNZWGN,
		SubdivisionNZWKO,
		SubdivisionNZWTC,
		SubdivisionOMBJ,
		SubdivisionOMBS,
		SubdivisionOMBU,
		SubdivisionOMDA,
		SubdivisionOMMA,
		SubdivisionOMMU,
		SubdivisionOMSJ,
		SubdivisionOMSS,
		SubdivisionOMWU,
		SubdivisionOMZA,
		SubdivisionOMZU,
		SubdivisionPA1,
		SubdivisionPA2,
		SubdivisionPA3,
		SubdivisionPA4,
		SubdivisionPA5,
		SubdivisionPA6,
		SubdivisionPA7,
		SubdivisionPA8,
		SubdivisionPA9,
		SubdivisionPANB,
		SubdivisionPEAMA,
		SubdivisionPEANC,
		SubdivisionPEAPU,
		SubdivisionPEARE,
		SubdivisionPEAYA,
		SubdivisionPECAJ,
		SubdivisionPECAL,
		SubdivisionPECUS,
		SubdivisionPEHUC,
		SubdivisionPEHUV,
		SubdivisionPEICA,
		SubdivisionPEJUN,
		SubdivisionPELAL,
		SubdivisionPELAM,
		SubdivisionPELIM,
		SubdivisionPELOR,
		SubdivisionPEMDD,
		SubdivisionPEMOQ,
		SubdivisionPEPAS,
		SubdivisionPEPIU,
		SubdivisionPEPUN,
		SubdivisionPESAM,
		SubdivisionPETAC,
		SubdivisionPETUM,
		SubdivisionPEUCA,
		SubdivisionPGCPM,
		SubdivisionPGEBR,
		SubdivisionPGEHG,
		SubdivisionPGESW,
		SubdivisionPGMBA,
		SubdivisionPGMPL,
		SubdivisionPGMPM,
		SubdivisionPGMRL,
		SubdivisionPGNCD,
		SubdivisionPGNIK,
		SubdivisionPGNSB,
		SubdivisionPGSAN,
		SubdivisionPGSHM,
		SubdivisionPGWBK,
		SubdivisionPGWHM,
		SubdivisionPGWPD,
		SubdivisionPH00,
		SubdivisionPHABR,
		SubdivisionPHAGN,
		SubdivisionPHAGS,
		SubdivisionPHAKL,
		SubdivisionPHALB,
		SubdivisionPHANT,
		SubdivisionPHAPA,
		SubdivisionPHAUR,
		SubdivisionPHBAN,
		SubdivisionPHBAS,
		SubdivisionPHBEN,
		SubdivisionPHBIL,
		SubdivisionPHBOH,
		SubdivisionPHBTG,
		SubdivisionPHBTN,
		SubdivisionPHBUK,
		SubdivisionPHBUL,
		SubdivisionPHCAG,
		SubdivisionPHCAM,
		SubdivisionPHCAN,
		SubdivisionPHCAP,
		SubdivisionPHCAS,
		SubdivisionPHCAT,
		SubdivisionPHCAV,
		SubdivisionPHCEB,
		SubdivisionPHCOM,
		SubdivisionPHDAO,
		SubdivisionPHDAS,
		SubdivisionPHDAV,
		SubdivisionPHDIN,
		SubdivisionPHEAS,
		SubdivisionPHGUI,
		SubdivisionPHIFU,
		SubdivisionPHILI,
		SubdivisionPHILN,
		SubdivisionPHILS,
		SubdivisionPHISA,
		SubdivisionPHKAL,
		SubdivisionPHLAG,
		SubdivisionPHLAN,
		SubdivisionPHLAS,
		SubdivisionPHLEY,
		SubdivisionPHLUN,
		SubdivisionPHMAD,
		SubdivisionPHMAG,
		SubdivisionPHMAS,
		SubdivisionPHMDC,
		SubdivisionPHMDR,
		SubdivisionPHMOU,
		SubdivisionPHMSC,
		SubdivisionPHMSR,
		SubdivisionPHNCO,
		SubdivisionPHNEC,
		SubdivisionPHNER,
		SubdivisionPHNSA,
		SubdivisionPHNUE,
		SubdivisionPHNUV,
		SubdivisionPHPAM,
		SubdivisionPHPAN,
		SubdivisionPHPLW,
		SubdivisionPHQUE,
		SubdivisionPHQUI,
		SubdivisionPHRIZ,
		SubdivisionPHROM,
		SubdivisionPHSAR,
		SubdivisionPHSCO,
		SubdivisionPHSIG,
		SubdivisionPHSLE,
		SubdivisionPHSLU,
		SubdivisionPHSOR,
		SubdivisionPHSUK,
		SubdivisionPHSUN,
		SubdivisionPHSUR,
		SubdivisionPHTAR,
		SubdivisionPHTAW,
		SubdivisionPHWSA,
		SubdivisionPHZAN,
		SubdivisionPHZAS,
		SubdivisionPHZMB,
		SubdivisionPHZSI,
		SubdivisionPKBA,
		SubdivisionPKGB,
		SubdivisionPKIS,
		SubdivisionPKJK,
		SubdivisionPKKP,
		SubdivisionPKPB,
		SubdivisionPKSD,
		SubdivisionPL02,
		SubdivisionPL04,
		SubdivisionPL06,
		SubdivisionPL08,
		SubdivisionPL10,
		SubdivisionPL12,
		SubdivisionPL14,
		SubdivisionPL16,
		SubdivisionPL18,
		SubdivisionPL20,
		SubdivisionPL22,
		SubdivisionPL24,
		SubdivisionPL26,
		SubdivisionPL28,
		SubdivisionPL30,
		SubdivisionPL32,
		SubdivisionPSBTH,
		SubdivisionPSDEB,
		SubdivisionPSGZA,
		SubdivisionPSHBN,
		SubdivisionPSJEM,
		SubdivisionPSJEN,
		SubdivisionPSJRH,
		SubdivisionPSKYS,
		SubdivisionPSNBS,
		SubdivisionPSQQA,
		SubdivisionPSRBH,
		SubdivisionPSRFH,
		SubdivisionPSSLT,
		SubdivisionPSTBS,
		SubdivisionPSTKM,
		SubdivisionPT01,
		SubdivisionPT02,
		SubdivisionPT03,
		SubdivisionPT04,
		SubdivisionPT05,
		SubdivisionPT06,
		SubdivisionPT07,
		SubdivisionPT08,
		SubdivisionPT09,
		SubdivisionPT10,
		SubdivisionPT11,
		SubdivisionPT12,
		SubdivisionPT13,
		SubdivisionPT14,
		SubdivisionPT15,
		SubdivisionPT16,
		SubdivisionPT17,
		SubdivisionPT18,
		SubdivisionPT20,
		SubdivisionPT30,
		SubdivisionPW004,
		SubdivisionPW100,
		SubdivisionPW150,
		SubdivisionPW212,
		SubdivisionPW214,
		SubdivisionPW222,
		SubdivisionPY1,
		SubdivisionPY10,
		SubdivisionPY11,
		SubdivisionPY12,
		SubdivisionPY13,
		SubdivisionPY14,
		SubdivisionPY15,
		SubdivisionPY16,
		SubdivisionPY19,
		SubdivisionPY2,
		SubdivisionPY3,
		SubdivisionPY4,
		SubdivisionPY5,
		SubdivisionPY6,
		SubdivisionPY7,
		SubdivisionPY8,
		SubdivisionPY9,
		SubdivisionPYASU,
		SubdivisionQADA,
		SubdivisionQAKH,
		SubdivisionQAMS,
		SubdivisionQARA,
		SubdivisionQAUS,
		SubdivisionQAWA,
		SubdivisionQAZA,
		SubdivisionROAB,
		SubdivisionROAG,
		SubdivisionROAR,
		SubdivisionROB,
		SubdivisionROBC,
		SubdivisionROBH,
		SubdivisionROBN,
		SubdivisionROBR,
		SubdivisionROBT,
		SubdivisionROBV,
		SubdivisionROBZ,
		SubdivisionROCJ,
		SubdivisionROCL,
		SubdivisionROCS,
		SubdivisionROCT,
		SubdivisionROCV,
		SubdivisionRODB,
		SubdivisionRODJ,
		SubdivisionROGJ,
		SubdivisionROGL,
		SubdivisionROGR,
		SubdivisionROHD,
		SubdivisionROHR,
		SubdivisionROIF,
		SubdivisionROIL,
		SubdivisionROIS,
		SubdivisionROMH,
		SubdivisionROMM,
		SubdivisionROMS,
		SubdivisionRONT,
		SubdivisionROOT,
		SubdivisionROPH,
		SubdivisionROSB,
		SubdivisionROSJ,
		SubdivisionROSM,
		SubdivisionROSV,
		SubdivisionROTL,
		SubdivisionROTM,
		SubdivisionROTR,
		SubdivisionROVL,
		SubdivisionROVN,
		SubdivisionROVS,
		SubdivisionRS00,
		SubdivisionRS01,
		SubdivisionRS02,
		SubdivisionRS03,
		SubdivisionRS04,
		SubdivisionRS05,
		SubdivisionRS06,
		SubdivisionRS07,
		SubdivisionRS08,
		SubdivisionRS09,
		SubdivisionRS10,
		SubdivisionRS11,
		SubdivisionRS12,
		SubdivisionRS13,
		SubdivisionRS14,
		SubdivisionRS15,
		SubdivisionRS16,
		SubdivisionRS17,
		SubdivisionRS18,
		SubdivisionRS19,
		SubdivisionRS20,
		SubdivisionRS21,
		SubdivisionRS22,
		SubdivisionRS23,
		SubdivisionRS24,
		SubdivisionRS26,
		SubdivisionRS27,
		SubdivisionRS28,
		SubdivisionRUAD,
		SubdivisionRUAL,
		SubdivisionRUALT,
		SubdivisionRUAMU,
		SubdivisionRUARK,
		SubdivisionRUAST,
		SubdivisionRUBA,
		SubdivisionRUBEL,
		SubdivisionRUBRY,
		SubdivisionRUBU,
		SubdivisionRUCE,
		SubdivisionRUCHE,
		SubdivisionRUCHU,
		SubdivisionRUCU,
		SubdivisionRUDA,
		SubdivisionRUIN,
		SubdivisionRUIRK,
		SubdivisionRUIVA,
		SubdivisionRUKAM,
		SubdivisionRUKB,
		SubdivisionRUKC,
		SubdivisionRUKDA,
		SubdivisionRUKEM,
		SubdivisionRUKGD,
		SubdivisionRUKGN,
		SubdivisionRUKHA,
		SubdivisionRUKHM,
		SubdivisionRUKIR,
		SubdivisionRUKK,
		SubdivisionRUKL,
		SubdivisionRUKLU,
		SubdivisionRUKO,
		SubdivisionRUKOS,
		SubdivisionRUKR,
		SubdivisionRUKRS,
		SubdivisionRUKYA,
		SubdivisionRULEN,
		SubdivisionRULIP,
		SubdivisionRUMAG,
		SubdivisionRUME,
		SubdivisionRUMO,
		SubdivisionRUMOS,
		SubdivisionRUMOW,
		SubdivisionRUMUR,
		SubdivisionRUNEN,
		SubdivisionRUNGR,
		SubdivisionRUNIZ,
		SubdivisionRUNVS,
		SubdivisionRUOMS,
		SubdivisionRUORE,
		SubdivisionRUORL,
		SubdivisionRUPER,
		SubdivisionRUPNZ,
		SubdivisionRUPRI,
		SubdivisionRUPSK,
		SubdivisionRUROS,
		SubdivisionRURYA,
		SubdivisionRUSA,
		SubdivisionRUSAK,
		SubdivisionRUSAM,
		SubdivisionRUSAR,
		SubdivisionRUSE,
		SubdivisionRUSMO,
		SubdivisionRUSPE,
		SubdivisionRUSTA,
		SubdivisionRUSVE,
		SubdivisionRUTA,
		SubdivisionRUTAM,
		SubdivisionRUTOM,
		SubdivisionRUTUL,
		SubdivisionRUTVE,
		SubdivisionRUTY,
		SubdivisionRUTYU,
		SubdivisionRUUD,
		SubdivisionRUULY,
		SubdivisionRUVGG,
		SubdivisionRUVLA,
		SubdivisionRUVLG,
		SubdivisionRUVOR,
		SubdivisionRUYAN,
		SubdivisionRUYAR,
		SubdivisionRUYEV,
		SubdivisionRUZAB,
		SubdivisionRW01,
		SubdivisionRW02,
		SubdivisionRW03,
		SubdivisionRW04,
		SubdivisionRW05,
		SubdivisionSA01,
		SubdivisionSA02,
		SubdivisionSA03,
		SubdivisionSA04,
		SubdivisionSA05,
		SubdivisionSA06,
		SubdivisionSA07,
		SubdivisionSA08,
		SubdivisionSA09,
		SubdivisionSA10,
		SubdivisionSA11,
		SubdivisionSA12,
		SubdivisionSA14,
		SubdivisionSBCH,
		SubdivisionSBGU,
		SubdivisionSBWE,
		SubdivisionSC01,
		SubdivisionSC02,
		SubdivisionSC06,
		SubdivisionSC07,
		SubdivisionSC08,
		SubdivisionSC10,
		SubdivisionSC11,
		SubdivisionSC13,
		SubdivisionSC14,
		SubdivisionSC15,
		SubdivisionSC16,
		SubdivisionSC20,
		SubdivisionSC23,
		SubdivisionSDDC,
		SubdivisionSDDN,
		SubdivisionSDDS,
		SubdivisionSDDW,
		SubdivisionSDGD,
		SubdivisionSDGK,
		SubdivisionSDGZ,
		SubdivisionSDKA,
		SubdivisionSDKH,
		SubdivisionSDKN,
		SubdivisionSDKS,
		SubdivisionSDNB,
		SubdivisionSDNO,
		SubdivisionSDNR,
		SubdivisionSDNW,
		SubdivisionSDRS,
		SubdivisionSDSI,
		SubdivisionSEAB,
		SubdivisionSEAC,
		SubdivisionSEBD,
		SubdivisionSEC,
		SubdivisionSED,
		SubdivisionSEE,
		SubdivisionSEF,
		SubdivisionSEG,
		SubdivisionSEH,
		SubdivisionSEI,
		SubdivisionSEK,
		SubdivisionSEM,
		SubdivisionSEN,
		SubdivisionSEO,
		SubdivisionSES,
		SubdivisionSET,
		SubdivisionSEU,
		SubdivisionSEW,
		SubdivisionSEX,
		SubdivisionSEY,
		SubdivisionSEZ,
		SubdivisionSG01,
		SubdivisionSG02,
		SubdivisionSG03,
		SubdivisionSG04,
		SubdivisionSG05,
		SubdivisionSHHL,
		SubdivisionSI001,
		SubdivisionSI002,
		SubdivisionSI003,
		SubdivisionSI004,
		SubdivisionSI005,
		SubdivisionSI006,
		SubdivisionSI007,
		SubdivisionSI008,
		SubdivisionSI009,
		SubdivisionSI010,
		SubdivisionSI011,
		SubdivisionSI012,
		SubdivisionSI013,
		SubdivisionSI014,
		SubdivisionSI015,
		SubdivisionSI017,
		SubdivisionSI018,
		SubdivisionSI019,
		SubdivisionSI020,
		SubdivisionSI021,
		SubdivisionSI023,
		SubdivisionSI024,
		SubdivisionSI025,
		SubdivisionSI026,
		SubdivisionSI029,
		SubdivisionSI031,
		SubdivisionSI032,
		SubdivisionSI033,
		SubdivisionSI034,
		SubdivisionSI035,
		SubdivisionSI036,
		SubdivisionSI037,
		SubdivisionSI038,
		SubdivisionSI039,
		SubdivisionSI040,
		SubdivisionSI041,
		SubdivisionSI042,
		SubdivisionSI043,
		SubdivisionSI044,
		SubdivisionSI045,
		SubdivisionSI046,
		SubdivisionSI047,
		SubdivisionSI048,
		SubdivisionSI049,
		SubdivisionSI050,
		SubdivisionSI052,
		SubdivisionSI053,
		SubdivisionSI054,
		SubdivisionSI055,
		SubdivisionSI056,
		SubdivisionSI057,
		SubdivisionSI058,
		SubdivisionSI059,
		SubdivisionSI060,
		SubdivisionSI061,
		SubdivisionSI063,
		SubdivisionSI064,
		SubdivisionSI065,
		SubdivisionSI066,
		SubdivisionSI067,
		SubdivisionSI068,
		SubdivisionSI069,
		SubdivisionSI070,
		SubdivisionSI071,
		SubdivisionSI072,
		SubdivisionSI073,
		SubdivisionSI074,
		SubdivisionSI075,
		SubdivisionSI076,
		SubdivisionSI077,
		SubdivisionSI079,
		SubdivisionSI080,
		SubdivisionSI081,
		SubdivisionSI082,
		SubdivisionSI083,
		SubdivisionSI084,
		SubdivisionSI085,
		SubdivisionSI086,
		SubdivisionSI087,
		SubdivisionSI090,
		SubdivisionSI091,
		SubdivisionSI092,
		SubdivisionSI094,
		SubdivisionSI095,
		SubdivisionSI096,
		SubdivisionSI097,
		SubdivisionSI098,
		SubdivisionSI099,
		SubdivisionSI100,
		SubdivisionSI101,
		SubdivisionSI102,
		SubdivisionSI103,
		SubdivisionSI104,
		SubdivisionSI105,
		SubdivisionSI106,
		SubdivisionSI108,
		SubdivisionSI109,
		SubdivisionSI110,
		SubdivisionSI111,
		SubdivisionSI112,
		SubdivisionSI113,
		SubdivisionSI114,
		SubdivisionSI115,
		SubdivisionSI116,
		SubdivisionSI117,
		SubdivisionSI118,
		SubdivisionSI119,
		SubdivisionSI120,
		SubdivisionSI121,
		SubdivisionSI122,
		SubdivisionSI123,
		SubdivisionSI124,
		SubdivisionSI125,
		SubdivisionSI126,
		SubdivisionSI127,
		SubdivisionSI128,
		SubdivisionSI129,
		SubdivisionSI130,
		SubdivisionSI131,
		SubdivisionSI132,
		SubdivisionSI133,
		SubdivisionSI134,
		SubdivisionSI135,
		SubdivisionSI136,
		SubdivisionSI137,
		SubdivisionSI138,
		SubdivisionSI139,
		SubdivisionSI140,
		SubdivisionSI141,
		SubdivisionSI142,
		SubdivisionSI143,
		SubdivisionSI144,
		SubdivisionSI146,
		SubdivisionSI147,
		SubdivisionSI148,
		SubdivisionSI149,
		SubdivisionSI150,
		SubdivisionSI151,
		SubdivisionSI152,
		SubdivisionSI154,
		SubdivisionSI155,
		SubdivisionSI156,
		SubdivisionSI158,
		SubdivisionSI159,
		SubdivisionSI160,
		SubdivisionSI161,
		SubdivisionSI162,
		SubdivisionSI164,
		SubdivisionSI165,
		SubdivisionSI166,
		SubdivisionSI167,
		SubdivisionSI168,
		SubdivisionSI169,
		SubdivisionSI170,
		SubdivisionSI171,
		SubdivisionSI172,
		SubdivisionSI173,
		SubdivisionSI174,
		SubdivisionSI175,
		SubdivisionSI176,
		SubdivisionSI179,
		SubdivisionSI180,
		SubdivisionSI182,
		SubdivisionSI183,
		SubdivisionSI184,
		SubdivisionSI185,
		SubdivisionSI186,
		SubdivisionSI187,
		SubdivisionSI188,
		SubdivisionSI189,
		SubdivisionSI190,
		SubdivisionSI191,
		SubdivisionSI193,
		SubdivisionSI194,
		SubdivisionSI195,
		SubdivisionSI196,
		SubdivisionSI197,
		SubdivisionSI198,
		SubdivisionSI199,
		SubdivisionSI200,
		SubdivisionSI201,
		SubdivisionSI203,
		SubdivisionSI204,
		SubdivisionSI205,
		SubdivisionSI206,
		SubdivisionSI207,
		SubdivisionSI208,
		SubdivisionSI209,
		SubdivisionSI210,
		SubdivisionSI211,
		SubdivisionSI212,
		SubdivisionSI213,
		SubdivisionSKBC,
		SubdivisionSKBL,
		SubdivisionSKKI,
		SubdivisionSKNI,
		SubdivisionSKPV,
		SubdivisionSKTA,
		SubdivisionSKTC,
		SubdivisionSKZI,
		SubdivisionSLE,
		SubdivisionSLN,
		SubdivisionSLS,
		SubdivisionSLW,
		SubdivisionSM03,
		SubdivisionSM07,
		SubdivisionSM09,
		SubdivisionSNDB,
		SubdivisionSNDK,
		SubdivisionSNFK,
		SubdivisionSNKA,
		SubdivisionSNKD,
		SubdivisionSNKE,
		SubdivisionSNKL,
		SubdivisionSNLG,
		SubdivisionSNMT,
		SubdivisionSNSL,
		SubdivisionSNTC,
		SubdivisionSNTH,
		SubdivisionSNZG,
		SubdivisionSOAW,
		SubdivisionSOBN,
		SubdivisionSOMU,
		SubdivisionSONU,
		SubdivisionSOSH,
		SubdivisionSOSO,
		SubdivisionSOTO,
		SubdivisionSOWO,
		SubdivisionSRCM,
		SubdivisionSRNI,
		SubdivisionSRPM,
		SubdivisionSRPR,
		SubdivisionSRSI,
		SubdivisionSRWA,
		SubdivisionSSBN,
		SubdivisionSSEC,
		SubdivisionSSEE,
		SubdivisionSSEW,
		SubdivisionSSNU,
		SubdivisionST01,
		SubdivisionSVAH,
		SubdivisionSVCA,
		SubdivisionSVCH,
		SubdivisionSVCU,
		SubdivisionSVLI,
		SubdivisionSVMO,
		SubdivisionSVPA,
		SubdivisionSVSA,
		SubdivisionSVSM,
		SubdivisionSVSO,
		SubdivisionSVSS,
		SubdivisionSVSV,
		SubdivisionSVUN,
		SubdivisionSVUS,
		SubdivisionSYDI,
		SubdivisionSYDR,
		SubdivisionSYDY,
		SubdivisionSYHA,
		SubdivisionSYHI,
		SubdivisionSYHL,
		SubdivisionSYHM,
		SubdivisionSYID,
		SubdivisionSYLA,
		SubdivisionSYQU,
		SubdivisionSYRA,
		SubdivisionSYRD,
		SubdivisionSYSU,
		SubdivisionSYTA,
		SubdivisionSZHH,
		SubdivisionSZLU,
		SubdivisionSZMA,
		SubdivisionTDGR,
		SubdivisionTDLO,
		SubdivisionTDME,
		SubdivisionTDND,
		SubdivisionTDOD,
		SubdivisionTGC,
		SubdivisionTGK,
		SubdivisionTGM,
		SubdivisionTGP,
		SubdivisionTH10,
		SubdivisionTH11,
		SubdivisionTH12,
		SubdivisionTH13,
		SubdivisionTH14,
		SubdivisionTH15,
		SubdivisionTH16,
		SubdivisionTH17,
		SubdivisionTH18,
		SubdivisionTH19,
		SubdivisionTH20,
		SubdivisionTH21,
		SubdivisionTH22,
		SubdivisionTH23,
		SubdivisionTH24,
		SubdivisionTH25,
		SubdivisionTH26,
		SubdivisionTH27,
		SubdivisionTH30,
		SubdivisionTH31,
		SubdivisionTH32,
		SubdivisionTH33,
		SubdivisionTH34,
		SubdivisionTH35,
		SubdivisionTH36,
		SubdivisionTH37,
		SubdivisionTH38,
		SubdivisionTH39,
		SubdivisionTH40,
		SubdivisionTH41,
		SubdivisionTH42,
		SubdivisionTH43,
		SubdivisionTH44,
		SubdivisionTH45,
		SubdivisionTH46,
		SubdivisionTH47,
		SubdivisionTH48,
		SubdivisionTH49,
		SubdivisionTH50,
		SubdivisionTH51,
		SubdivisionTH52,
		SubdivisionTH53,
		SubdivisionTH54,
		SubdivisionTH55,
		SubdivisionTH56,
		SubdivisionTH57,
		SubdivisionTH58,
		SubdivisionTH60,
		SubdivisionTH61,
		SubdivisionTH62,
		SubdivisionTH63,
		SubdivisionTH64,
		SubdivisionTH65,
		SubdivisionTH66,
		SubdivisionTH67,
		SubdivisionTH70,
		SubdivisionTH71,
		SubdivisionTH72,
		SubdivisionTH73,
		SubdivisionTH74,
		SubdivisionTH75,
		SubdivisionTH76,
		SubdivisionTH77,
		SubdivisionTH80,
		SubdivisionTH81,
		SubdivisionTH82,
		SubdivisionTH83,
		SubdivisionTH84,
		SubdivisionTH85,
		SubdivisionTH86,
		SubdivisionTH90,
		SubdivisionTH91,
		SubdivisionTH92,
		SubdivisionTH93,
		SubdivisionTH94,
		SubdivisionTH95,
		SubdivisionTH96,
		SubdivisionTJDU,
		SubdivisionTJKT,
		SubdivisionTJRA,
		SubdivisionTJSU,
		SubdivisionTLAN,
		SubdivisionTLCO,
		SubdivisionTLDI,
		SubdivisionTMA,
		SubdivisionTMB,
		SubdivisionTMD,
		SubdivisionTML,
		SubdivisionTMM,
		SubdivisionTN11,
		SubdivisionTN12,
		SubdivisionTN13,
		SubdivisionTN14,
		SubdivisionTN21,
		SubdivisionTN22,
		SubdivisionTN23,
		SubdivisionTN31,
		SubdivisionTN32,
		SubdivisionTN33,
		SubdivisionTN34,
		SubdivisionTN41,
		SubdivisionTN42,
		SubdivisionTN43,
		SubdivisionTN51,
		SubdivisionTN52,
		SubdivisionTN53,
		SubdivisionTN61,
		SubdivisionTN71,
		SubdivisionTN72,
		SubdivisionTN73,
		SubdivisionTN81,
		SubdivisionTN82,
		SubdivisionTN83,
		SubdivisionTO03,
		SubdivisionTO04,
		SubdivisionTR01,
		SubdivisionTR02,
		SubdivisionTR03,
		SubdivisionTR04,
		SubdivisionTR05,
		SubdivisionTR06,
		SubdivisionTR07,
		SubdivisionTR08,
		SubdivisionTR09,
		SubdivisionTR10,
		SubdivisionTR11,
		SubdivisionTR12,
		SubdivisionTR13,
		SubdivisionTR14,
		SubdivisionTR15,
		SubdivisionTR16,
		SubdivisionTR17,
		SubdivisionTR18,
		SubdivisionTR19,
		SubdivisionTR20,
		SubdivisionTR21,
		SubdivisionTR22,
		SubdivisionTR23,
		SubdivisionTR24,
		SubdivisionTR25,
		SubdivisionTR26,
		SubdivisionTR27,
		SubdivisionTR28,
		SubdivisionTR29,
		SubdivisionTR30,
		SubdivisionTR31,
		SubdivisionTR32,
		SubdivisionTR33,
		SubdivisionTR34,
		SubdivisionTR35,
		SubdivisionTR36,
		SubdivisionTR37,
		SubdivisionTR38,
		SubdivisionTR39,
		SubdivisionTR40,
		SubdivisionTR41,
		SubdivisionTR42,
		SubdivisionTR43,
		SubdivisionTR44,
		SubdivisionTR45,
		SubdivisionTR46,
		SubdivisionTR47,
		SubdivisionTR48,
		SubdivisionTR49,
		SubdivisionTR50,
		SubdivisionTR51,
		SubdivisionTR52,
		SubdivisionTR53,
		SubdivisionTR54,
		SubdivisionTR55,
		SubdivisionTR56,
		SubdivisionTR57,
		SubdivisionTR58,
		SubdivisionTR59,
		SubdivisionTR60,
		SubdivisionTR61,
		SubdivisionTR62,
		SubdivisionTR63,
		SubdivisionTR64,
		SubdivisionTR65,
		SubdivisionTR66,
		SubdivisionTR67,
		SubdivisionTR68,
		SubdivisionTR69,
		SubdivisionTR70,
		SubdivisionTR71,
		SubdivisionTR72,
		SubdivisionTR73,
		SubdivisionTR74,
		SubdivisionTR75,
		SubdivisionTR76,
		SubdivisionTR77,
		SubdivisionTR78,
		SubdivisionTR79,
		SubdivisionTR80,
		SubdivisionTR81,
		SubdivisionTTARI,
		SubdivisionTTCHA,
		SubdivisionTTCTT,
		SubdivisionTTDMN,
		SubdivisionTTMRC,
		SubdivisionTTPED,
		SubdivisionTTPOS,
		SubdivisionTTPRT,
		SubdivisionTTPTF,
		SubdivisionTTSFO,
		SubdivisionTTSGE,
		SubdivisionTTSIP,
		SubdivisionTTSJL,
		SubdivisionTTTOB,
		SubdivisionTTTUP,
		SubdivisionTVFUN,
		SubdivisionTWCHA,
		SubdivisionTWCYQ,
		SubdivisionTWHSQ,
		SubdivisionTWHUA,
		SubdivisionTWILA,
		SubdivisionTWKEE,
		SubdivisionTWKHH,
		SubdivisionTWKIN,
		SubdivisionTWLIE,
		SubdivisionTWMIA,
		SubdivisionTWNAN,
		SubdivisionTWNWT,
		SubdivisionTWPEN,
		SubdivisionTWPIF,
		SubdivisionTWTAO,
		SubdivisionTWTNN,
		SubdivisionTWTPE,
		SubdivisionTWTTT,
		SubdivisionTWTXG,
		SubdivisionTWYUN,
		SubdivisionTZ01,
		SubdivisionTZ02,
		SubdivisionTZ03,
		SubdivisionTZ04,
		SubdivisionTZ05,
		SubdivisionTZ07,
		SubdivisionTZ08,
		SubdivisionTZ09,
		SubdivisionTZ11,
		SubdivisionTZ12,
		SubdivisionTZ13,
		SubdivisionTZ14,
		SubdivisionTZ15,
		SubdivisionTZ16,
		SubdivisionTZ17,
		SubdivisionTZ18,
		SubdivisionTZ19,
		SubdivisionTZ20,
		SubdivisionTZ21,
		SubdivisionTZ22,
		SubdivisionTZ23,
		SubdivisionTZ24,
		SubdivisionTZ25,
		SubdivisionTZ26,
		SubdivisionTZ27,
		SubdivisionTZ28,
		SubdivisionTZ29,
		SubdivisionTZ30,
		SubdivisionTZ31,
		SubdivisionUA05,
		SubdivisionUA07,
		SubdivisionUA09,
		SubdivisionUA12,
		SubdivisionUA14,
		SubdivisionUA18,
		SubdivisionUA21,
		SubdivisionUA23,
		SubdivisionUA26,
		SubdivisionUA30,
		SubdivisionUA32,
		SubdivisionUA35,
		SubdivisionUA40,
		SubdivisionUA43,
		SubdivisionUA46,
		SubdivisionUA48,
		SubdivisionUA51,
		SubdivisionUA53,
		SubdivisionUA56,
		SubdivisionUA59,
		SubdivisionUA61,
		SubdivisionUA63,
		SubdivisionUA65,
		SubdivisionUA68,
		SubdivisionUA71,
		SubdivisionUA74,
		SubdivisionUA77,
		SubdivisionUG101,
		SubdivisionUG102,
		SubdivisionUG103,
		SubdivisionUG104,
		SubdivisionUG105,
		SubdivisionUG106,
		SubdivisionUG107,
		SubdivisionUG108,
		SubdivisionUG109,
		SubdivisionUG112,
		SubdivisionUG113,
		SubdivisionUG115,
		SubdivisionUG116,
		SubdivisionUG117,
		SubdivisionUG120,
		SubdivisionUG122,
		SubdivisionUG201,
		SubdivisionUG203,
		SubdivisionUG204,
		SubdivisionUG205,
		SubdivisionUG206,
		SubdivisionUG209,
		SubdivisionUG210,
		SubdivisionUG211,
		SubdivisionUG214,
		SubdivisionUG215,
		SubdivisionUG219,
		SubdivisionUG222,
		SubdivisionUG303,
		SubdivisionUG304,
		SubdivisionUG305,
		SubdivisionUG307,
		SubdivisionUG316,
		SubdivisionUG321,
		SubdivisionUG403,
		SubdivisionUG404,
		SubdivisionUG405,
		SubdivisionUG406,
		SubdivisionUG407,
		SubdivisionUG408,
		SubdivisionUG409,
		SubdivisionUG410,
		SubdivisionUG411,
		SubdivisionUG412,
		SubdivisionUG413,
		SubdivisionUG415,
		SubdivisionUG419,
		SubdivisionUM95,
		SubdivisionUSAK,
		SubdivisionUSAL,
		SubdivisionUSAR,
		SubdivisionUSAZ,
		SubdivisionUSCA,
		SubdivisionUSCO,
		SubdivisionUSCT,
		SubdivisionUSDC,
		SubdivisionUSDE,
		SubdivisionUSFL,
		SubdivisionUSGA,
		SubdivisionUSHI,
		SubdivisionUSIA,
		SubdivisionUSID,
		SubdivisionUSIL,
		SubdivisionUSIN,
		SubdivisionUSKS,
		SubdivisionUSKY,
		SubdivisionUSLA,
		SubdivisionUSMA,
		SubdivisionUSMD,
		SubdivisionUSME,
		SubdivisionUSMI,
		SubdivisionUSMN,
		SubdivisionUSMO,
		SubdivisionUSMS,
		SubdivisionUSMT,
		SubdivisionUSNC,
		SubdivisionUSND,
		SubdivisionUSNE,
		SubdivisionUSNH,
		SubdivisionUSNJ,
		SubdivisionUSNM,
		SubdivisionUSNV,
		SubdivisionUSNY,
		SubdivisionUSOH,
		SubdivisionUSOK,
		SubdivisionUSOR,
		SubdivisionUSPA,
		SubdivisionUSRI,
		SubdivisionUSSC,
		SubdivisionUSSD,
		SubdivisionUSTN,
		SubdivisionUSTX,
		SubdivisionUSUT,
		SubdivisionUSVA,
		SubdivisionUSVT,
		SubdivisionUSWA,
		SubdivisionUSWI,
		SubdivisionUSWV,
		SubdivisionUSWY,
		SubdivisionUYAR,
		SubdivisionUYCA,
		SubdivisionUYCL,
		SubdivisionUYCO,
		SubdivisionUYDU,
		SubdivisionUYFD,
		SubdivisionUYFS,
		SubdivisionUYLA,
		SubdivisionUYMA,
		SubdivisionUYMO,
		SubdivisionUYPA,
		SubdivisionUYRN,
		SubdivisionUYRO,
		SubdivisionUYRV,
		SubdivisionUYSA,
		SubdivisionUYSJ,
		SubdivisionUYSO,
		SubdivisionUYTA,
		SubdivisionUYTT,
		SubdivisionUZAN,
		SubdivisionUZBU,
		SubdivisionUZFA,
		SubdivisionUZJI,
		SubdivisionUZNG,
		SubdivisionUZNW,
		SubdivisionUZQA,
		SubdivisionUZQR,
		SubdivisionUZSA,
		SubdivisionUZSI,
		SubdivisionUZSU,
		SubdivisionUZTK,
		SubdivisionUZXO,
		SubdivisionVC01,
		SubdivisionVC04,
		SubdivisionVC05,
		SubdivisionVC06,
		SubdivisionVEA,
		SubdivisionVEB,
		SubdivisionVEC,
		SubdivisionVED,
		SubdivisionVEE,
		SubdivisionVEF,
		SubdivisionVEG,
		SubdivisionVEH,
		SubdivisionVEI,
		SubdivisionVEJ,
		SubdivisionVEK,
		SubdivisionVEL,
		SubdivisionVEM,
		SubdivisionVEN,
		SubdivisionVEO,
		SubdivisionVEP,
		SubdivisionVER,
		SubdivisionVES,
		SubdivisionVET,
		SubdivisionVEU,
		SubdivisionVEV,
		SubdivisionVEX,
		SubdivisionVEY,
		SubdivisionVEZ,
		SubdivisionVN01,
		SubdivisionVN02,
		SubdivisionVN03,
		SubdivisionVN04,
		SubdivisionVN05,
		SubdivisionVN06,
		SubdivisionVN07,
		SubdivisionVN09,
		SubdivisionVN13,
		SubdivisionVN14,
		SubdivisionVN18,
		SubdivisionVN20,
		SubdivisionVN21,
		SubdivisionVN22,
		SubdivisionVN23,
		SubdivisionVN24,
		SubdivisionVN25,
		SubdivisionVN26,
		SubdivisionVN27,
		SubdivisionVN28,
		SubdivisionVN29,
		SubdivisionVN30,
		SubdivisionVN31,
		SubdivisionVN32,
		SubdivisionVN33,
		SubdivisionVN34,
		SubdivisionVN35,
		SubdivisionVN36,
		SubdivisionVN37,
		SubdivisionVN39,
		SubdivisionVN40,
		SubdivisionVN41,
		SubdivisionVN43,
		SubdivisionVN44,
		SubdivisionVN45,
		SubdivisionVN46,
		SubdivisionVN47,
		SubdivisionVN49,
		SubdivisionVN50,
		SubdivisionVN51,
		SubdivisionVN52,
		SubdivisionVN53,
		SubdivisionVN54,
		SubdivisionVN55,
		SubdivisionVN56,
		SubdivisionVN57,
		SubdivisionVN58,
		SubdivisionVN59,
		SubdivisionVN61,
		SubdivisionVN63,
		SubdivisionVN66,
		SubdivisionVN67,
		SubdivisionVN68,
		SubdivisionVN69,
		SubdivisionVN70,
		SubdivisionVN71,
		SubdivisionVN72,
		SubdivisionVNCT,
		SubdivisionVNDN,
		SubdivisionVNHN,
		SubdivisionVNHP,
		SubdivisionVNSG,
		SubdivisionVUSEE,
		SubdivisionVUTAE,
		SubdivisionVUTOB,
		SubdivisionWFSG,
		SubdivisionWFUV,
		SubdivisionWSAT,
		SubdivisionWSFA,
		SubdivisionWSTU,
		SubdivisionYEAB,
		SubdivisionYEAD,
		SubdivisionYEAM,
		SubdivisionYEBA,
		SubdivisionYEDA,
		SubdivisionYEDH,
		SubdivisionYEHD,
		SubdivisionYEHJ,
		SubdivisionYEHU,
		SubdivisionYEIB,
		SubdivisionYELA,
		SubdivisionYEMA,
		SubdivisionYEMR,
		SubdivisionYESA,
		SubdivisionYESD,
		SubdivisionYESH,
		SubdivisionYESN,
		SubdivisionYETA,
		SubdivisionZAEC,
		SubdivisionZAFS,
		SubdivisionZAGP,
		SubdivisionZAKZN,
		SubdivisionZALP,
		SubdivisionZAMP,
		SubdivisionZANC,
		SubdivisionZANW,
		SubdivisionZAWC,
		SubdivisionZM01,
		SubdivisionZM02,
		SubdivisionZM03,
		SubdivisionZM04,
		SubdivisionZM05,
		SubdivisionZM06,
		SubdivisionZM07,
		SubdivisionZM08,
		SubdivisionZM09,
		SubdivisionZM10,
		SubdivisionZWBU,
		SubdivisionZWHA,
		SubdivisionZWMA,
		SubdivisionZWMC,
		SubdivisionZWME,
		SubdivisionZWMI,
		SubdivisionZWMN,
		SubdivisionZWMS,
		SubdivisionZWMV,
		SubdivisionZWMW,
	}
}

// AllSubdivisionsInfo - return all subdivision codes as []Subdivision
func AllSubdivisionsInfo() []*Subdivision {
	all := AllSubdivisions()
	subdivisions := make([]*Subdivision, 0, len(all))
	for _, v := range all {
		subdivisions = append(subdivisions, v.Info())
	}
	return subdivisions
}

// AllSubdivisionsByCountryCode - returns all the subdivisions, mapped to their country code
func AllSubdivisionsByCountryCode() map[CountryCode][]SubdivisionCode {
	resp := map[CountryCode][]SubdivisionCode{}

	for _, s := range AllSubdivisions() {
		c := s.Country()
		if _, ok := resp[c]; !ok {
			resp[c] = []SubdivisionCode{}
		}
		resp[c] = append(resp[c], s)
	}

	return resp
}

// SubdivisionsByCountryCode - returns all subdivisions for a particular country code
func SubdivisionsByCountryCode(c CountryCode) []SubdivisionCode {
	return AllSubdivisionsByCountryCode()[c]
}

// TotalSubdivisions - returns number of subdivisions in the package
func TotalSubdivisions() int {
	return 3332
}
