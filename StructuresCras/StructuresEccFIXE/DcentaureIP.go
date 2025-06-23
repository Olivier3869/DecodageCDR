package structuresEccFIXE

import "local/structures/defStructs"

// TcentaureIPtrealer
var TcentaureIPtrealer = map[int]defStructs.Champ{
	0: {"CODARTFIN", 6, 1, nil},
	1: {"NBART", 8, 1, nil},
	2: {"CODENT", 3, 1, nil},
	3: {"FILLER", 333, -1, nil},
	4: {"FIN", 1, -1, nil},
}

// TcentaureIPheader
var TcentaureIPheader = map[int]defStructs.Champ{
	0: {"CODARTETE", 6, 1, nil},
	1: {"CODENT", 3, 1, nil},
	2: {"FLUX", 2, 1, nil},
	3: {"DTTRT", 14, 1, nil},
	4: {"INDICE", 6, 1, nil},
	5: {"TYPELOT", 1, 1, nil},
	6: {"VERSION", 6, 1, nil},
	7: {"FILLER", 312, -1, nil},
	8: {"FIN", 1, -1, nil},
}

// TcentaureIP
var TcentaureIP = map[int]defStructs.Champ{
	0:  {"TYPENR", 3, 1, nil},
	1:  {"DATE", 14, 1, nil},
	2:  {"DUREE", 6, 1, nil},
	3:  {"IndRenv", 2, 1, nil},
	4:  {"CallType", 2, 1, nil},
	5:  {"IDUS_NAI", 2, 1, nil},
	6:  {"IdusNbPlan", 2, 1, nil},
	7:  {"IdusIndRedir", 2, 1, nil},
	8:  {"IDUS_NUM", 16, 1, nil},
	9:  {"IDUS_INF", 2, 1, nil},
	10: {"DEDE_NAI", 2, 1, nil},
	11: {"DedeNbPlan", 2, 1, nil},
	12: {"DEDE_NUM", 32, 1, nil},
	13: {"DEDE_INF", 2, 1, nil},
	14: {"DEST_NAI", 2, 1, nil},
	15: {"DestNbPlan", 2, 1, nil},
	16: {"DETR_NUM", 32, 1, nil},
	17: {"DETR_INF", 2, 1, nil},
	18: {"CallingPartyCat", 2, 1, nil},
	19: {"SUP", 2, 1, nil},
	20: {"ReroutInf", 2, 1, nil},
	21: {"IncomingPort", 4, 1, nil},
	22: {"OutgoingPort", 4, 1, nil},
	23: {"IncomingPOP", 2, 1, nil},
	24: {"OutgoingPOP", 2, 1, nil},
	25: {"Outgoingmod", 2, 1, nil},
	26: {"CIC", 2, 1, nil},
	27: {"SourceIP", 8, 1, nil},
	28: {"DestIP", 8, 1, nil},
	29: {"DurInterPh", 6, 1, nil},
	30: {"FSENT", 8, 1, nil},
	31: {"FSSORT", 8, 1, nil},
	32: {"DelCLIIDUSNAI", 2, 1, nil},
	33: {"DelCLIIDUSPlan", 2, 1, nil},
	34: {"DelCLIIDUSPubpr", 2, 1, nil},
	35: {"DelCLIIDUSNUM", 16, 1, nil},
	36: {"OCNDEDENAI", 2, 1, nil},
	37: {"OCNDEDENbPlan", 2, 1, nil},
	38: {"OCNDEDENUM", 16, 1, nil},
	39: {"RDNInfRedirNAI", 2, 1, nil},
	40: {"RDNInfRedirNbPl", 2, 1, nil},
	41: {"RDNInfRedirNUM", 16, 1, nil},
	42: {"CallEchoDevice", 2, 1, nil},
	43: {"NetwkCallID", 32, 1, nil},
	44: {"NbPulses", 8, 1, nil},
	45: {"SGTQS", 6, 1, nil},
	46: {"LOCAL_NAI", 2, 1, nil},
	47: {"LocalNbPlan", 2, 1, nil},
	48: {"LocIDUS_NUM", 16, 1, nil},
	49: {"TYPE_EDIT", 1, 1, nil},
	50: {"CODE_TARIF", 4, 1, nil},
	51: {"FILLER", 28, -1, nil},
	52: {"FIN", 1, -1, nil},
}

// DcentaureIP
var DcentaureIP = map[int]defStructs.TabRechercheTypeDc{
	0: {"TcentaureDCheader", 0xffff0000, 0x30300000, 351, &TcentaureIPheader},
	1: {"TcentaureDC", 0xffffff00, 0x4e424900, 351, &TcentaureIP},
	2: {"TcentaureDCtrealer", 0xffffff00, 0x39393900, 351, &TcentaureIPtrealer},
}
