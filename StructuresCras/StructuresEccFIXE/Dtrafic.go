package structuresEccFIXE

import "local/structures/defStructs"

// TtraficDc
var TtraficDc = map[int]defStructs.Champ{
	0:  {"DOMAINE", 2, 2, nil},
	1:  {"SERVICE", 2, 2, nil},
	2:  {"IDUS_INF", 3, 2, nil},
	3:  {"filler", 1, 2, nil},
	4:  {"IDUS_NUM", 9, 2, nil},
	5:  {"DEDE_INF", 3, 2, nil},
	6:  {"DEDE_NUM", 20, 2, nil},
	7:  {"ICAR", 14, 2, nil},
	8:  {"DUR", 6, 2, nil},
	9:  {"CAX", 1, 2, nil},
	10: {"TAX", 8, 2, nil},
	11: {"NDS_NUM", 10, 2, nil},
	12: {"SUP", 1, 2, nil},
	13: {"REF", 3, 2, nil},
	14: {"MES1", 3, 2, nil},
	15: {"MES2", 3, 2, nil},
	16: {"nb_csv", 1, 2, nil},
	17: {"c_csv1", 2, 2, nil},
	18: {"op_csv1", 1, 2, nil},
	19: {"c_csv2", 2, 2, nil},
	20: {"op_csv2", 1, 2, nil},
	21: {"c_csv3", 2, 2, nil},
	22: {"op_csv3", 1, 2, nil},
	23: {"c_csv4", 2, 2, nil},
	24: {"op_csv4", 1, 2, nil},
	25: {"c_csv5", 2, 2, nil},
	26: {"op_csv5", 1, 2, nil},
	27: {"STY", 2, 2, nil},
	28: {"COCA", 4, 2, nil},
	29: {"COCAV", 4, 2, nil},
	30: {"IDAV_INF", 3, 2, nil},
	31: {"IDAV_NUM", 20, 2, nil},
	32: {"SGTQS", 5, 2, nil},
	33: {"MNEMO", 10, 2, nil},
	34: {"CLASSDEST", 3, 2, nil},
	35: {"OFFRE", 6, 2, nil},
	36: {"OPERATEUR", 6, 2, nil},
	37: {"TAG", 1, 2, nil},
	38: {"TYPE_EDIT", 1, 2, nil},
	39: {"CODE_TARIF", 4, 2, nil},
	40: {"IPAS", 10, 2, nil},
	41: {"filler", 4, -1, nil},
}

// TtraficHeader
var TtraficHeader = map[int]defStructs.Champ{
	0: {"CODART", 2, 2, nil},
	1: {"APPLI", 2, 2, nil},
	2: {"VER", 6, 2, nil},
	3: {"DATTRAIT", 14, 2, nil},
	4: {"TYPE_LOT", 1, 2, nil},
	5: {"SGTQS", 5, 2, nil},
	6: {"n_ser", 3, 2, nil},
	7: {"n_seq", 3, 2, nil},
	8: {"filler", 152, -1, nil},
}

// TtraficTrailer
var TtraficTrailer = map[int]defStructs.Champ{
	0: {"CODART", 2, 2, nil},
	1: {"APPLI", 2, 2, nil},
	2: {"VER", 6, 2, nil},
	3: {"DATTRAIT", 14, 2, nil},
	4: {"TYPE_LOT", 1, 2, nil},
	5: {"SGTQS", 5, 2, nil},
	6: {"NBART", 8, 2, nil},
	7: {"filler", 150, -1, nil},
}

// Dtrafic
var Dtrafic = map[int]defStructs.TabRechercheTypeDc{
	0: {"TtraficHeader", 0xffff0000, 0xf0f00000, 188, &TtraficHeader},
	1: {"TtraficTrailer", 0xffff0000, 0xf9f90000, 188, &TtraficTrailer},
	2: {"TtraficRtc", 0xffff0000, 0xf0f50000, 188, &TtraficDc},
	3: {"TtraficVoip", 0xffff0000, 0xf1f00000, 188, &TtraficDc},
	4: {"TtraficRiOsa", 0xffff0000, 0xf2f80000, 188, &TtraficDc},
	5: {"TtraficPcs118", 0xffff0000, 0xf3f30000, 188, &TtraficDc},
	6: {"TtraficPcsConv", 0xffff0000, 0xf5f10000, 188, &TtraficDc},
	7: {"TtraficPAS", 0xffff0000, 0xf6f40000, 188, &TtraficDc},
}
