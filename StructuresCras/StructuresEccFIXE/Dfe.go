package structuresEccFIXE

import "local/structures/defStructs"

// TfeDc
var TfeDc = map[int]defStructs.Champ{
	0:  {"SERVICE", 2, 2, nil},
	1:  {"IDUS_INF", 3, 2, nil},
	2:  {"E", 1, 2, nil},
	3:  {"IDUS_NUM", 9, 2, nil},
	4:  {"DEDE_INF", 3, 2, nil},
	5:  {"DEDE_NUM", 20, 2, nil},
	6:  {"ICAR", 14, 2, nil},
	7:  {"DUR", 6, 2, nil},
	8:  {"CAX", 1, 2, nil},
	9:  {"TAX", 8, 2, nil},
	10: {"NDS_INF", 3, 2, nil},
	11: {"NDS_NUM", 10, 2, nil},
	12: {"SUP", 1, 2, nil},
	13: {"REF", 3, 2, nil},
	14: {"CLASSDEST", 3, 2, nil},
	15: {"MES1", 3, 2, nil},
	16: {"MES2", 3, 2, nil},
	17: {"NB_CSV", 1, 2, nil},
	18: {"C_CSV1", 2, 2, nil},
	19: {"OP_CSV1", 1, 2, nil},
	20: {"C_CSV2", 2, 2, nil},
	21: {"OP_CSV2", 1, 2, nil},
	22: {"C_CSV3", 2, 2, nil},
	23: {"OP_CSV3", 1, 2, nil},
	24: {"C_CSV4", 2, 2, nil},
	25: {"OP_CSV4", 1, 2, nil},
	26: {"C_CSV5", 2, 2, nil},
	27: {"OP_CSV5", 1, 2, nil},
	28: {"COCA", 4, 2, nil},
	29: {"COCAV", 4, 2, nil},
	30: {"IDAV_INF", 3, 2, nil},
	31: {"IDAV_NUM", 20, 2, nil},
	32: {"MNEMO", 10, 2, nil},
	33: {"TAG", 1, 2, nil},
	34: {"TYPE_EDIT", 1, 2, nil},
	35: {"CODE_TARIF", 4, 2, nil},
	36: {"Filler", 4, -1, nil},
}

// TfeHeader2
var TfeHeader2 = map[int]defStructs.Champ{
	0: {"CODART", 2, 2, nil},
	1: {"EMET", 6, 2, nil},
	2: {"RECEPT", 6, 2, nil},
	3: {"APPLI_E", 3, 2, nil},
	4: {"APPLI_R", 3, 2, nil},
	5: {"VER", 2, 2, nil},
	6: {"NBART", 8, 2, nil},
	7: {"TYPELOT", 1, 2, nil},
	8: {"FILLER", 129, -1, nil},
}

// TfeHeader1
var TfeHeader1 = map[int]defStructs.Champ{
	0: {"CODART", 2, 2, nil},
	1: {"FLUX", 8, 2, nil},
	2: {"N_SQ", 8, 2, nil},
	3: {"DATTRAIT", 14, 2, nil},
	4: {"LONG", 3, 2, nil},
	5: {"FILLER", 125, -1, nil},
}

// Dfe
var Dfe = map[int]defStructs.TabRechercheTypeDc{
	0: {"TfeHeader1", 0xffff0000, 0xc1c10000, 160, &TfeHeader1},
	1: {"TfeHeader2", 0xffff0000, 0xc1c20000, 160, &TfeHeader2},
	2: {"TfeDc", 0x00000000, 0x00000000, 160, &TfeDc},
}
