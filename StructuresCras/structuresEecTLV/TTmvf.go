package structuresEecTLV

import "local/structures/defStructs"

// TTmvf
var TTmvf = map[int]defStructs.TabRechercheTypeDc{
	0: {"Header", 0xffff0000, 0x80030000, 0, &Entete},
	1: {"Tmvf", 0xffff0000, 0x806f0000, 0, &Tmvf},
}

// Entete
var Entete = map[int]defStructs.Champ{
	0x03: {"Header", -1, 0, &EnteteBLOC},
}

// EnteteBLOC
var EnteteBLOC = map[int]defStructs.Champ{
	0: {"Bloc", 3, 0, nil},
	1: {"Date", 6, 9, nil},
	2: {"Caa", 6, 1, nil},
	3: {"Emi", 2, 0, nil},
	4: {"Codage", 1, 0, nil},
	5: {"Mine", 1, 0, nil},
}

// Tmvf
var Tmvf = map[int]defStructs.Champ{
	0x6f: {"T1", -1, 0, &t2},
}

// t2
var t2 = map[int]defStructs.Champ{
	0x02: {"T2", -1, 0, &t3},
}

// t3
var t3 = map[int]defStructs.Champ{
	0x01: {"APPLI", -1, 0, nil},
	0x02: {"ICARDEB", -1, 13, nil},
	0x03: {"ICARFIN", -1, 13, nil},
	0x04: {"INEF", -1, 0, nil},
	0x05: {"IDUS", -1, 0, &T3_IDUS},
	0x06: {"Mail", -1, 2, nil},
	0x07: {"IDUSGENE", -1, 0, nil},
	0x08: {"IDUSPRIV", -1, 0, nil},
	0x09: {"APP", -1, 0, &T3_APP},
	0x0a: {"DEDE", -1, 0, &T3_DEDE},
	0x0b: {"DEDEMAIL", -1, 0, nil},
	0x0c: {"DEDEGENE", -1, 0, nil},
	0x0d: {"DEDEPRIV", -1, 0, nil},
	0x0e: {"VOLDATA", -1, 0, nil},
	0x0f: {"Origine", -1, 0, nil},
	0x10: {"Destination", -1, 0, nil},
	0x11: {"Domaine", -1, 0, nil},
	0x12: {"Dur", -1, 0, nil},
	0x13: {"Fact", -1, 0, nil},
	0x14: {"Etap", -1, 0, nil},
}

// T3_IDUS
var T3_IDUS = map[int]defStructs.Champ{
	0: {"INF", 1, 0, nil},
	1: {"NUM", -1, 0, nil},
}

// T3_APP
var T3_APP = map[int]defStructs.Champ{
	0: {"APPinf", 1, 0, nil},
	1: {"APPnum", -1, 0, nil},
}

// T3_DEDE
var T3_DEDE = map[int]defStructs.Champ{
	0: {"INF", 1, 0, nil},
	1: {"NUM", -1, 0, nil},
}
