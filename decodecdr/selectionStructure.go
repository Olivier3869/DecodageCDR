package main

import (
	"fmt"
	"local/structures/StructuresEecFIXE"
	"local/structures/defStructs"
	"local/structures/structuresEccCSV"
	"local/structures/structuresEccFIXE"
	"local/structures/structuresEccTLV"
	"local/structures/structuresEecASN1"
	"local/structures/structuresEecCSV"
	"local/structures/structuresEecTLV"
)

// selectionStructure sélectionne une structure basée sur le tag fourni.
//
// Paramètres :
// - tag (string) : une chaîne représentant le tag pour lequel la structure doit être sélectionnée.
//
// Retourne :
// - map[int]defStructs.TabRechercheTypeDc : une carte contenant la structure correspondante.
// - string : le format de la structure (TLV, FIXE, CSV, ASN1, TASNP).
// - string : une chaîne supplémentaire, qui peut être vide ou contenir des informations spécifiques (ex. "RabbitMQ", "Kafka").

func selectionStructure(tag string) (map[int]defStructs.TabRechercheTypeDc, string, string) {
	switch {
	// ECC format TLV
	case tag == "Dafdb":
		return structuresEccTLV.Dafdb, "TLV", ""
	case tag == "Dcasper":
		return structuresEccTLV.Dcasper, "TLV", ""
	case tag == "Ddpvgsm":
		return structuresEccTLV.Ddpvgsm, "TLV", ""
	case tag == "Ddpvgsm28":
		return structuresEccTLV.Ddpvgsm28, "TLV", ""
	case tag == "Ddpvgsm32":
		return structuresEccTLV.Ddpvgsm32, "TLV", ""
	case tag == "DdpvgsmFree":
		return structuresEccTLV.DdpvgsmFree, "TLV", ""
	case tag == "Ddqd":
		return structuresEccTLV.Ddqd, "TLV", ""
	case tag == "Dmeteor":
		return structuresEccTLV.Dmeteor, "TLV", ""
	case tag == "Dspot":
		return structuresEccTLV.Dspot, "TLV", ""
	case tag == "Dsymp_home":
		return structuresEccTLV.Dsymp_home, "TLV", ""
	case tag == "Dsymphonie":
		return structuresEccTLV.Dsymphonie, "TLV", ""

	// ECC format FIXE
	case tag == "D16xconv":
		return structuresEccFIXE.D16xconv, "FIXE", ""
	case tag == "DallegroNew":
		return structuresEccFIXE.DallegroNew, "FIXE", ""
	case tag == "Dandromeri":
		return structuresEccFIXE.Dandromeri, "FIXE", ""
	case tag == "DcentaureBIV":
		return structuresEccFIXE.DcentaureBIV, "FIXE", ""
	case tag == "DcentaureDC":
		return structuresEccFIXE.DcentaureDC, "FIXE", ""
	case tag == "DcentaureHF":
		return structuresEccFIXE.DcentaureHF, "FIXE", ""
	case tag == "DcentaureIP":
		return structuresEccFIXE.DcentaureIP, "FIXE", ""
	case tag == "DcentaureVG":
		return structuresEccFIXE.DcentaureVG, "FIXE", ""
	case tag == "DdeltaMinute":
		return structuresEccFIXE.DdeltaMinute, "FIXE", ""
	case tag == "Dentropi":
		return structuresEccFIXE.Dentropi, "FIXE", ""
	case tag == "Dfe":
		return structuresEccFIXE.Dfe, "FIXE", ""
	case tag == "Dferpv":
		return structuresEccFIXE.Dferpv, "FIXE", ""
	case tag == "Dtapin":
		return structuresEccFIXE.Dtapin, "FIXE", ""
	case tag == "Dtrafic":
		return structuresEccFIXE.Dtrafic, "FIXE", ""

	// ECC format CSV
	case tag == "Dcaps":
		return structuresEccCSV.Dcaps, "CSV", ""
	case tag == "Dclientip":
		return structuresEccCSV.Dclientip, "CSV", ""
	case tag == "Dentropi_sonus":
		return structuresEccCSV.Dentropi_sonus, "CSV", ""
	case tag == "DkenobiWha":
		return structuresEccCSV.DkenobiWha, "CSV", ""
	case tag == "Dmmdm":
		return structuresEccCSV.Dmmdm, "CSV", ""
	case tag == "DmmdmWha":
		return structuresEccCSV.DmmdmWha, "CSV", ""

	// EEC format CSV
	case tag == "TT02F_HP":
		return StructuresEecFIXE.TT02F_HP, "FIXE", ""

	// EEC format CSV
	case tag == "TTasbc":
		return structuresEecCSV.TTasbc, "CSV", ""
	case tag == "TTeaming":
		return structuresEecCSV.TTeaming, "CSV", ""
	case tag == "TTines":
		return structuresEecCSV.TTines, "CSV", "RabbitMQ"
	case tag == "TTlmisf":
		return structuresEecCSV.TTlmisf, "CSV", ""
	case tag == "TTmmscXura":
		return structuresEecCSV.TTmmscXura, "CSV", ""
	case tag == "TTNBI":
		return structuresEecCSV.TTNBI, "CSV", ""
	case tag == "TTorca":
		return structuresEecCSV.TTorca, "CSV", ""
	case tag == "TTsmsoIP":
		return structuresEecCSV.TTsmsoIP, "CSV", ""
	case tag == "TTsonus":
		return structuresEecCSV.TTsonus, "CSV", ""
	case tag == "TTsvaipTAG":
		return structuresEecCSV.TTsvaipTAG, "CSV", ""
	case tag == "TTv2oipAS":
		return structuresEecCSV.TTv2oipAS, "CSV", ""
	case tag == "TTv2oipSBC":
		return structuresEecCSV.TTv2oipSBC, "CSV", ""
	case tag == "TTvoipvpn2":
		return structuresEecCSV.TTvoipvpn2, "CSV", ""
	case tag == "TTwha":
		return structuresEecCSV.TTwha, "CSV", ""
	case tag == "TTwhaInternet":
		return structuresEecCSV.TTwhaInternet, "CSV", ""
	case tag == "TTwhaSelect":
		return structuresEecCSV.TTwhaSelect, "CSV", ""

	// EEC format ASN1

	case tag == "TTbgw_gsn":
		return structuresEecASN1.TTbgw_gsn, "ASN1", ""
	case tag == "TTchfcdr":
		return structuresEecASN1.TTchfcdr, "ASN1", "Kafka"
	case tag == "TTcrae":
		return structuresEecASN1.TTtap310ZB, "ASN1", "" // La structure tap310 plus complete que CRAE dans le cdf
	case tag == "TTepdg":
		return structuresEecASN1.TTepdg, "ASN1", ""
	case tag == "TTims":
		return structuresEecASN1.TTims, "ASN1", ""
	case tag == "TTmmTelTas":
		return structuresEecASN1.TTmmTelTas, "ASN1", ""
	case tag == "TTnss9":
		return structuresEecASN1.TTnss9, "ASN1", ""
	case tag == "TTpgwcdr":
		return structuresEecASN1.TTpgwcdr, "ASN1", ""
	case tag == "TTpgwe":
		return structuresEecASN1.TTpgwe, "ASN1", ""
	case tag == "TTsgwcdr":
		return structuresEecASN1.TTsgwcdr, "ASN1", ""
	case tag == "TTSms_cmg":
		return structuresEecASN1.TTSms_cmg, "ASN1", ""
	case tag == "TTSms_mcohr":
		return structuresEecASN1.TTSms_mcohr, "ASN1", ""
	case tag == "TTsssas":
		return structuresEecASN1.TTsssas, "ASN1", ""
	case tag == "TTtap310ZB":
		return structuresEecASN1.TTtap310ZB, "ASN1", ""

	// EEC format TASNP
	case tag == "TTtasnokia":
		return structuresEecASN1.TTtasnokia, "TASNP", ""

	// EEC format TLV
	case tag == "trait1":
		return structuresEecTLV.Trait1, "TLV", ""
	case tag == "TTaliasNb":
		return structuresEecTLV.TTaliasNb, "TLV", ""
	case tag == "TTcgw":
		return structuresEecTLV.TTcgw, "TLV", ""
	case tag == "TTmmm":
		return structuresEecTLV.TTmmm, "TLV", ""
	case tag == "TTmvf":
		return structuresEecTLV.TTmvf, "TLV", ""
	case tag == "TTnodes":
		return structuresEecTLV.TTnodes, "TLV", ""
	case tag == "TTsvaip":
		return structuresEecTLV.TTsvaip, "TLV", ""

	// CAS si structure non définie
	default:
		fmt.Printf("Structure %v non prise en compte.\n", tag)
		return nil, "INCONNU", ""
	}

}
