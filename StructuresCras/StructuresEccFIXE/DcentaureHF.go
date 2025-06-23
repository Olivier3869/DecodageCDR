package structuresEccFIXE

import "local/structures/defStructs"

// TcentaureHFtrealer
var TcentaureHFtrealer = map[int]defStructs.Champ{
	0: {"CODARTFIN", 6, 1, nil},
	1: {"NBART", 8, 1, nil},
	2: {"CODENT", 3, 1, nil},
	3: {"FILLER", 110, -1, nil},
	4: {"FIN", 1, -1, nil},
}

// TcentaureHFheader
var TcentaureHFheader = map[int]defStructs.Champ{
	0: {"CODARTETE", 6, 1, nil},
	1: {"CODENT", 3, 1, nil},
	2: {"FLUX", 2, 1, nil},
	3: {"DTTRT", 14, 1, nil},
	4: {"INDICE", 6, 1, nil},
	5: {"TYPELOT", 1, 1, nil},
	6: {"VERSION", 6, 1, nil},
	7: {"FILLER", 89, -1, nil},
	8: {"FIN", 1, -1, nil},
}

// TcentaureHF
var TcentaureHF = map[int]defStructs.Champ{
	0:  {"T1", 3, 1, nil},
	1:  {"T2", 3, 1, nil},
	2:  {"CNDR", 2, 1, nil},
	3:  {"DATE", 14, 1, nil},
	4:  {"IDUS_AaF", 2, 1, nil},
	5:  {"IDUS_GH", 1, 1, nil},
	6:  {"IDUS_NUM", 16, 1, nil},
	7:  {"DEDE_INF", 3, 1, nil},
	8:  {"DEDE_NUM", 20, 1, nil},
	9:  {"DUREE", 6, 1, nil},
	10: {"CTISI_CD", 1, 1, nil},
	11: {"CTISI_GH", 1, 1, nil},
	12: {"NDES_NUM", 20, 1, nil},
	13: {"NDES_Nat", 1, 1, nil},
	14: {"PREFPORT", 10, 1, nil},
	15: {"NBRE", 1, 1, nil},
	16: {"TYPAP1_ABC", 1, 1, nil},
	17: {"TYPAP1_DE", 1, 1, nil},
	18: {"TYPAP1_F", 1, 1, nil},
	19: {"TYPAP1_G", 1, 1, nil},
	20: {"SGTQS1", 4, 1, nil},
	21: {"FAISCENT1", 5, 1, nil},
	22: {"FAISSORT1", 5, 1, nil},
	23: {"TYPE_EDIT", 1, 1, nil},
	24: {"CODE_TARIF", 4, 1, nil},
	25: {"FIN", 1, -1, nil},
}

// DcentaureHF
var DcentaureHF = map[int]defStructs.TabRechercheTypeDc{
	0: {"TcentaureHFheader", 0xffff0000, 0x30300000, 128, &TcentaureHFheader},
	1: {"TcentaureHF", 0xffff0000, 0x31300000, 128, &TcentaureHF},
	2: {"TcentaureHFtrealer", 0xffff0000, 0x39390000, 128, &TcentaureHFtrealer},
}
