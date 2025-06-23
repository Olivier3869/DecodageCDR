package symp_home

import "local/structures/defStructs"

// TsymphTconvN
var TsymphTconvN = map[int]defStructs.Champ{
	0x02: {"TsymphTconvNE", -1, -1, &TsymphTconvNT3},
	0x03: {"TsymphTconvNI", -1, -1, &TsymphTconvNT3},
}

// TsymphTconvNT3
var TsymphTconvNT3 = map[int]defStructs.Champ{
	0x02: {"GLA2", -1, -1, &TsymphTconvNT3GLA2},
	0x4b: {"DEDEini", -1, -1, nil},
	0x4c: {"DEDEini2", -1, -1, nil},
	0x43: {"DEDE", -1, -1, &TsymphTconvNT3DEDE},
	0x48: {"DEDE1", -1, -1, nil},
	0x27: {"DETR", -1, -1, &TsymphTconvNT3DETR},
	0x2b: {"DETR1", -1, -1, &TsymphTconvNT3DETR},
	0x28: {"DETN", -1, -1, &TsymphTconvNT3DE},
	0x29: {"DETD", -1, -1, &TsymphTconvNT3DE},
	0x40: {"IDUS", -1, -1, &TsymphTconvNT3DE},
	0x41: {"ICCM", -1, -1, &TsymphTconvNT3DE},
	0x37: {"CDPO", -1, -1, nil},
	0x07: {"CPTX", -1, -1, nil},
	0x09: {"MV", -1, -1, &TsymphTconvNT3MV},
	0x0b: {"TLA", -1, -1, &TsymphTconvNT3TLA},
	0x05: {"DECGEO", -1, -1, &TsymphTconvNT3DECGEO},
	0x31: {"NCCONF", -1, -1, nil},
	0x35: {"PLT", -1, -1, nil},
	0x14: {"RDIP", -1, -1, &TsymphTconvNT3RDIP},
	0x15: {"LISTE", -1, -1, nil},
	0x08: {"MNEMO", -1, 1, nil},
	0x44: {"ENAP", -1, -1, &TsymphTconvNT3ENAP},
	0x0a: {"TRABO", -1, -1, &TsymphTconvNT3TRABO},
	0x2a: {"TRAD", -1, -1, nil},
	0x1a: {"LINK", -1, -1, &TsymphTconvNT3LINK},
	0x4e: {"IPER", -1, -1, &TsymphTconvNT3DE},
	0x4f: {"OADC", -1, -1, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4sympAjoutPlatine},
	0x50: {"TAG50?", -1, -1, nil},
}

// TsymphTconvNT3GLA2
var TsymphTconvNT3GLA2 = map[int]defStructs.Champ{
	0: {"IDSV", 1, -1, nil},
	1: {"ICAR", 6, 5, nil},
	2: {"HFAPP", 3, 6, nil},
	3: {"CPCS", 2, -1, nil},
	4: {"CCAS", 2, -1, nil},
	5: {"NATACH", 1, -1, nil},
}

// TsymphTconvNT3DEDE
var TsymphTconvNT3DEDE = map[int]defStructs.Champ{
	0: {"INF", 2, -1, nil},
	1: {"NUM", -1, -1, nil},
}

// TsymphTconvNT3DETR
var TsymphTconvNT3DETR = map[int]defStructs.Champ{
	0: {"INT", 2, -1, nil},
	1: {"INF", 1, -1, nil},
	2: {"NUM", -1, -1, nil},
}

// TsymphTconvNT3DE
var TsymphTconvNT3DE = map[int]defStructs.Champ{
	0: {"INF", 1, -1, nil},
	1: {"NUM", -1, -1, nil},
}

// TsymphTconvNT3MV
var TsymphTconvNT3MV = map[int]defStructs.Champ{
	0: {"MV", 2, -1, nil},
	1: {"DURPAC", 2, -1, nil},
	2: {"DURACC", 2, -1, nil},
	3: {"DURaMV", -1, -1, nil},
}

// TsymphTconvNT3TLA
var TsymphTconvNT3TLA = map[int]defStructs.Champ{
	0:  {"CDEUR", 1, -1, nil},
	1:  {"HDCONV", 3, -1, nil},
	2:  {"DUR", 3, -1, nil},
	3:  {"INEP", 1, -1, nil},
	4:  {"INEC", 1, -1, nil},
	5:  {"TAP", 2, -1, nil},
	6:  {"DURSONNTN", 2, -1, nil},
	7:  {"DURSONNTD", 2, -1, nil},
	8:  {"SUP", 1, -1, nil},
	9:  {"ITX", 1, -1, nil},
	10: {"TG", 1, -1, nil},
	11: {"NBUT", -1, -1, nil},
}

// TsymphTconvNT3DECGEO
var TsymphTconvNT3DECGEO = map[int]defStructs.Champ{
	0: {"ZAA", 2, -1, nil},
	1: {"ZCT", 2, -1, nil},
	2: {"ZDT", 1, -1, nil},
	3: {"ZRG", -1, -1, nil},
}

// TsymphTconvNT3RDIP
var TsymphTconvNT3RDIP = map[int]defStructs.Champ{
	0: {"DATE", 5, -1, nil},
	1: {"CPAYS", 2, -1, nil},
	2: {"IDPCS", 4, -1, nil},
	3: {"REFC", -1, -1, nil},
}

// TsymphTconvNT3ENAP
var TsymphTconvNT3ENAP = map[int]defStructs.Champ{
	0: {"REF", 3, -1, nil},
	1: {"SEQ", -1, -1, nil},
}

// TsymphTconvNT3TRABO
var TsymphTconvNT3TRABO = map[int]defStructs.Champ{
	0: {"NUTC", 2, -1, nil},
	1: {"NUAA", -1, -1, nil},
}

// TsymphTconvNT3LINK
var TsymphTconvNT3LINK = map[int]defStructs.Champ{
	0: {"APPEL", 3, -1, nil},
	1: {"NOEUD", -1, -1, nil},
}
