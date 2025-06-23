package structuresEccFIXE

import "local/structures/defStructs"

// TferpvDc
var TferpvDc = map[int]defStructs.Champ{
	0:  {"SERVICE", 1, 2, nil},
	1:  {"SOURCE", 2, 2, nil},
	2:  {"TYPCOM", 1, 2, nil},
	3:  {"NUMDDR", 10, 2, nil},
	4:  {"NUMDDE", 20, 2, nil},
	5:  {"DEBUT", 14, 2, nil},
	6:  {"DUREE", 6, 2, nil},
	7:  {"NBUT", 5, 2, nil},
	8:  {"MTHSTD", 12, 2, nil},
	9:  {"CANAL", 3, 2, nil},
	10: {"NBCOD", 1, 2, nil},
	11: {"C_CSV1", 2, 2, nil},
	12: {"OP_CSV1", 1, 2, nil},
	13: {"C_CSV2", 2, 2, nil},
	14: {"OP_CSV2", 1, 2, nil},
	15: {"C_CSV3", 2, 2, nil},
	16: {"OP_CSV3", 1, 2, nil},
	17: {"C_CSV4", 2, 2, nil},
	18: {"OP_CSV4", 1, 2, nil},
	19: {"C_CSV5", 2, 2, nil},
	20: {"OP_CSV5", 1, 2, nil},
	21: {"TYPACCRI", 1, 2, nil},
	22: {"SUPPORT", 1, 2, nil},
	23: {"NOCOMP", 20, 2, nil},
	24: {"NOORIG", 10, 2, nil},
	25: {"NATAPPRI", 1, 2, nil},
	26: {"CCD", 4, 2, nil},
	27: {"DDR_RMTA", 10, 2, nil},
	28: {"PREF_PORT", 5, 2, nil},
	29: {"TAG", 1, 2, nil},
	30: {"INFDDE", 3, 2, nil},
	31: {"TYPE_EDIT", 1, 2, nil},
	32: {"CODE_TARIF", 4, 2, nil},
	33: {"FILLER", 9, -1, nil},
}

// TferpvHeader2
var TferpvHeader2 = map[int]defStructs.Champ{
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

// TferpvHeader1
var TferpvHeader1 = map[int]defStructs.Champ{
	0: {"CODART", 2, 2, nil},
	1: {"FLUX", 8, 2, nil},
	2: {"N_SQ", 8, 2, nil},
	3: {"DATTRAIT", 14, 2, nil},
	4: {"LONG", 3, 2, nil},
	5: {"FILLER", 125, -1, nil},
}

// Dferpv
var Dferpv = map[int]defStructs.TabRechercheTypeDc{
	0: {"TferpvHeader1", 0xffff0000, 0xc1c10000, 160, &TferpvHeader1},
	1: {"TferpvHeader2", 0xffff0000, 0xc1c20000, 160, &TferpvHeader2},
	2: {"TferpvDc", 0x00000000, 0x00000000, 160, &TferpvDc},
}
