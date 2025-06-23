package structuresEccFIXE

import "local/structures/defStructs"

// TcentaureDCtrealer
var TcentaureDCtrealer = map[int]defStructs.Champ{
	0: {"CODARTFIN", 6, 1, nil},
	1: {"NBART", 8, 1, nil},
	2: {"CODENT", 3, 1, nil},
	3: {"FILLER", 333, -1, nil},
	4: {"FIN", 1, -1, nil},
}

// TcentaureDCheader
var TcentaureDCheader = map[int]defStructs.Champ{
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

// TcentaureDC
var TcentaureDC = map[int]defStructs.Champ{
	0:  {"T1", 3, 1, nil},
	1:  {"T2", 3, 1, nil},
	2:  {"SCT", 8, 1, nil},
	3:  {"CAX", 3, 1, nil},
	4:  {"TAP", 8, 1, nil},
	5:  {"STY", 8, 1, nil},
	6:  {"INDEX", 4, 1, nil},
	7:  {"DATE", 14, 1, nil},
	8:  {"ALINK", 8, 1, nil},
	9:  {"NLINK", 5, 1, nil},
	10: {"ILOCUS", 3, 1, nil},
	11: {"NLOCUS", 15, 1, nil},
	12: {"IIDUS", 3, 1, nil},
	13: {"NIDUS", 15, 1, nil},
	14: {"IIPRV", 3, 1, nil},
	15: {"NIPRV", 15, 1, nil},
	16: {"IIDRV", 3, 1, nil},
	17: {"NIDRV", 15, 1, nil},
	18: {"IDEDE", 3, 1, nil},
	19: {"NDEDE", 32, 1, nil},
	20: {"IIDAV", 3, 1, nil},
	21: {"NIDAV", 32, 1, nil},
	22: {"INOS", 3, 1, nil},
	23: {"NNOS", 10, 1, nil},
	24: {"IICOM", 3, 1, nil},
	25: {"NICOM", 10, 1, nil},
	26: {"REF", 3, 1, nil},
	27: {"NBRV", 3, 1, nil},
	28: {"DUREE", 6, 1, nil},
	29: {"TAX", 8, 1, nil},
	30: {"DUSEL", 5, 1, nil},
	31: {"DUSON", 5, 1, nil},
	32: {"TTX", 1, 1, nil},
	33: {"SIDR", 1, 1, nil},
	34: {"MES1", 3, 1, nil},
	35: {"MES2", 3, 1, nil},
	36: {"INEF", 3, 1, nil},
	37: {"CI", 6, 1, nil},
	38: {"CLE1", 3, 1, nil},
	39: {"CLE2", 3, 1, nil},
	40: {"CLE3", 3, 1, nil},
	41: {"CLE4", 3, 1, nil},
	42: {"COCA", 4, 1, nil},
	43: {"COCAV", 4, 1, nil},
	44: {"NUMLOT", 6, 1, nil},
	45: {"CDSGTQS", 5, 1, nil},
	46: {"SUP", 1, 1, nil},
	47: {"INVOCATION", 3, 1, nil},
	48: {"MERE", 3, 1, nil},
	49: {"ECHO", 8, 1, nil},
	50: {"MENV", 3, 1, nil},
	51: {"TYPE_EDIT", 1, 1, nil},
	52: {"CODE_TARIF", 4, 1, nil},
	53: {"BUFFEUR", 17, -1, nil},
	54: {"FIN", 1, -1, nil},
}

// DcentaureDC
var DcentaureDC = map[int]defStructs.TabRechercheTypeDc{
	0: {"TcentaureDCheader", 0xffff0000, 0x30300000, 351, &TcentaureDCheader},
	1: {"TcentaureDC", 0xffff0000, 0x30310000, 351, &TcentaureDC},
	2: {"TcentaureDCtrealer", 0xffff0000, 0x39390000, 351, &TcentaureDCtrealer},
}
