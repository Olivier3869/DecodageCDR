package structuresEccFIXE

import "local/structures/defStructs"

// TcentaureVGtrealer
var TcentaureVGtrealer = map[int]defStructs.Champ{
	0: {"CODARTFIN", 6, 1, nil},
	1: {"NBART", 8, 1, nil},
	2: {"CODENT", 3, 1, nil},
	3: {"FILLER", 171, -1, nil},
	4: {"FIN", 1, -1, nil},
}

// TcentaureVGheader
var TcentaureVGheader = map[int]defStructs.Champ{
	0: {"CODARTETE", 6, 1, nil},
	1: {"CODENT", 3, 1, nil},
	2: {"FLUX", 2, 1, nil},
	3: {"DTTRT", 14, 1, nil},
	4: {"INDICE", 6, 1, nil},
	5: {"TYPELOT", 1, 1, nil},
	6: {"VERSION", 6, 1, nil},
	7: {"FILLER", 150, -1, nil},
	8: {"FIN", 1, -1, nil},
}

// TcentaureVG
var TcentaureVG = map[int]defStructs.Champ{
	0:  {"DOMAINE", 2, 1, nil},
	1:  {"SERVICE", 2, 1, nil},
	2:  {"IDUS_INF", 3, 1, nil},
	3:  {"IDUS_NUM", 9, 1, nil},
	4:  {"DEDE_INF", 3, 1, nil},
	5:  {"DEDE_NUM", 20, 1, nil},
	6:  {"ICAR", 14, 1, nil},
	7:  {"DUR", 6, 1, nil},
	8:  {"CAX", 1, 1, nil},
	9:  {"TAX", 8, 1, nil},
	10: {"NDS_NUM", 10, 1, nil},
	11: {"SUP", 1, 1, nil},
	12: {"REF", 3, 1, nil},
	13: {"MES1", 3, 1, nil},
	14: {"MES2", 3, 1, nil},
	15: {"nb_csv", 1, 1, nil},
	16: {"C_CSV1", 2, 1, nil},
	17: {"OP_CSV1", 1, 1, nil},
	18: {"C_CSV2", 2, 1, nil},
	19: {"OP_CSV2", 1, 1, nil},
	20: {"C_CSV3", 2, 1, nil},
	21: {"OP_CSV3", 1, 1, nil},
	22: {"C_CSV4", 2, 1, nil},
	23: {"OP_CSV4", 1, 1, nil},
	24: {"C_CSV5", 2, 1, nil},
	25: {"OP_CSV5", 1, 1, nil},
	26: {"STY", 2, 1, nil},
	27: {"COCA", 4, 1, nil},
	28: {"COCAV", 4, 1, nil},
	29: {"IDAV_INF", 3, 1, nil},
	30: {"IDAV_NUM", 20, 1, nil},
	31: {"SGTQS", 5, 1, nil},
	32: {"MNEMO", 10, 1, nil},
	33: {"CLASSDEST", 3, 1, nil},
	34: {"OFFRE", 6, 1, nil},
	35: {"OPERATEUR", 6, -1, nil},
	36: {"TAG", 1, 1, nil},
	37: {"TYPE_EDIT", 1, 1, nil},
	38: {"CODE_TARIF", 4, 1, nil},
	39: {"FILLER", 15, -1, nil},
	40: {"FIN", 1, -1, nil},
}

// DcentaureVG
var DcentaureVG = map[int]defStructs.TabRechercheTypeDc{
	0: {"TcentaureVGheader", 0xffff0000, 0x30300000, 189, &TcentaureVGheader},
	1: {"TcentaureVG", 0xffff0000, 0x30350000, 189, &TcentaureVG},
	2: {"TcentaureVGtrealer", 0xffff0000, 0x39390000, 189, &TcentaureVGtrealer},
	3: {"TcentaureVG", 0xffff0000, 0x35310000, 189, &TcentaureVG},
	4: {"TcentaureVG", 0xffff0000, 0x31300000, 189, &TcentaureVG},
}
