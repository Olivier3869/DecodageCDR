package structuresEecCSV

import "local/structures/defStructs"

// TTsvaipTAG
var TTsvaipTAG = map[int]defStructs.TabRechercheTypeDc{
	0: {"Entete", 0xffffffff, 0x474c4132, 0x7c, nil},
	1: {"TTsvaipTAG", 0x00000000, 0x00000000, 0x7c, &TsvaipTAG},
}
var TsvaipTAG = map[int]defStructs.Champ{
	0:  {"GLA2ICAR", -1, 0, nil},
	1:  {"GLA2HFAPP", -1, 0, nil},
	2:  {"GLA2NATACH", -1, 0, nil},
	3:  {"DEDEiniNUM", -1, 0, nil},
	4:  {"DETR.NUM", -1, 0, nil},
	5:  {"IDUS.INF", -1, 0, nil},
	6:  {"IDUS.NUM", -1, 0, nil},
	7:  {"ICCM.INF", -1, 0, nil},
	8:  {"ICCM.NUM", -1, 0, nil},
	9:  {"TLA.CDEUR", -1, 0, nil},
	10: {"TLA_HDCONV", -1, 0, nil},
	11: {"TLA.DUR", -1, 0, nil},
	12: {"TLA.DURSN", -1, 0, nil},
	13: {"ENAP.REF", -1, 0, nil},
	14: {"ENAP.CALLID", -1, 0, nil},
	15: {"ENAP.SEQ", -1, 0, nil},
	16: {"IPER.INF", -1, 0, nil},
	17: {"IPER_NUM", -1, 0, nil},
	18: {"FINSER.CAU", -1, 0, nil},
}
