package structuresEecTLV

import "local/structures/defStructs"

// TTnodes
var TTnodes = map[int]defStructs.TabRechercheTypeDc{
	0: {"Txxbloc3", 0xffff0000, 0x03150000, 0, &Txxbloc3},
	1: {"Tmvt7dc", 0xff00ff00, 0x24000200, 0, &Tmvt7dc},
}

var Txxbloc3 = map[int]defStructs.Champ{
	0x03: {"Entete", -1, 0, &Entetebloc},
}

var Entetebloc = map[int]defStructs.Champ{
	0: {"Bloc", 3, 0, nil},
	1: {"Date", 6, 0, nil},
	2: {"Caa", 6, 0, nil},
	3: {"Emi", 2, 0, nil},
	4: {"Art", 1, 0, nil},
	5: {"Cod", 1, 0, nil},
	6: {"Mine", 1, 0, nil},
	7: {"Appli", -1, 0, nil},
}

var Tmvt7dc = map[int]defStructs.Champ{
	0x24: {"T1", -1, 0, &Tmvt7dcT2},
}

var Tmvt7dcT2 = map[int]defStructs.Champ{
	0x02: {"T2", -1, 0, &Tmvt7dcT3},
}

var Tmvt7dcT3 = map[int]defStructs.Champ{
	0x03: {"G3", -1, 0, &T3G3},
	0x04: {"G4", -1, 0, &T3G4},
	0x1a: {"LINK", -1, 0, &T3LINK},
	0x40: {"IDUS", -1, 0, &T3IDUS},
	0x43: {"DEDE", -1, 0, &T3DEDE},
	0x42: {"NOS", -1, 0, &T3NOS},
	0x08: {"CAP1", -1, 0, &T3CAP1},
	0x0a: {"CAP2", -1, 0, &T3CAP2},
	0x3c: {"SIDR", -1, 0, nil},
	0x78: {"ETAP", -1, 0, nil},
}

var T3G3 = map[int]defStructs.Champ{
	0: {"Sty", 1, 0, nil},
	1: {"Sct", 1, 0, nil},
	2: {"Icar", 5, 0, nil},
	3: {"Inef", -1, 0, nil},
}

var T3G4 = map[int]defStructs.Champ{
	0: {"Sty", 1, 0, nil},
	1: {"Sct", 1, 0, nil},
	2: {"Icar", 5, 0, nil},
	3: {"Inef", 1, 0, nil},
	4: {"Ref", -1, 0, nil},
}

var T3LINK = map[int]defStructs.Champ{
	0: {"appel", 3, 0, nil},
	1: {"noeud", 2, 0, nil},
}

var T3IDUS = map[int]defStructs.Champ{
	0: {"Inf", 1, 0, nil},
	1: {"Num", 5, 0, nil},
}

var T3DEDE = map[int]defStructs.Champ{
	0: {"Inf", -1, 0, nil},
}

var T3NOS = map[int]defStructs.Champ{
	0: {"inf", -1, 0, nil},
}

var T3CAP1 = map[int]defStructs.Champ{
	0: {"Cax", 1, 0, nil},
	1: {"Tap", 1, 0, nil},
	2: {"Dur", 2, 0, nil},
	3: {"Tax", -1, 0, nil},
}

var T3CAP2 = map[int]defStructs.Champ{
	0: {"Cax", 1, 0, nil},
	1: {"Tap", 1, 0, nil},
	2: {"Dur", 2, 0, nil},
	3: {"Tax", 2, 0, nil},
	4: {"sup", -1, 0, nil},
}
