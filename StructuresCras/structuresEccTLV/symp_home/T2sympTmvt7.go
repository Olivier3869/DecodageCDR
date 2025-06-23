package symp_home

import "local/structures/defStructs"

// T2sympTmvt7
var T2sympTmvt7 = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3sympTmvt7dc},
}

// T3sympTmvt7dc
var T3sympTmvt7dc = map[int]defStructs.Champ{
	0x03: {"T3sympTmvt7G3", -1, -1, &T4sympTmvt7G3},
	0x04: {"T3sympTmvt7G4", -1, -1, &T4sympTmvt7G4},
	0x08: {"T3sympTmvt7Cap1", -1, -1, &T4sympTmvt7Cap1},
	0x0a: {"T3sympTmvt7Cap2", -1, -1, &T4sympTmvt7Cap2},
	0x1a: {"T3sympTmvt7Link", -1, -1, &T4sympTmvt7Link},
	0x3c: {"T3sympTmvt7SIDR", -1, -1, nil},
	0x40: {"T3sympTmvt7Idus", -1, -1, &T4sympTmvt7Idus},
	0x42: {"T3sympTmvt7Nos", -1, -1, &T4sympTmvt7Nos},
	0x43: {"T3sympTmvt7Dede", -1, -1, &T4sympTmvt7Dede},
	0x78: {"T3sympTmvt7ETAP", -1, -1, nil},
}

// T4sympTmvt7G3
var T4sympTmvt7G3 = map[int]defStructs.Champ{
	0: {"Sty", 1, -1, nil},
	1: {"Sct", 1, -1, nil},
	2: {"ICAR", 5, -1, nil},
	3: {"Inef", -1, -1, nil},
}

// T4sympTmvt7G4
var T4sympTmvt7G4 = map[int]defStructs.Champ{
	0: {"Sty", 1, -1, nil},
	1: {"Sct", 1, -1, nil},
	2: {"ICAR", 5, -1, nil},
	3: {"Inef", 1, -1, nil},
	4: {"Ref", -1, -1, nil},
}

// T4sympTmvt7Link
var T4sympTmvt7Link = map[int]defStructs.Champ{
	0: {"appel", 3, -1, nil},
	1: {"noeud", 2, -1, nil},
}

// T4sympTmvt7Idus
var T4sympTmvt7Idus = map[int]defStructs.Champ{
	0: {"IDUS_INF", 1, -1, nil},
	1: {"IDUS_NUM", 5, -1, nil},
}

// T4sympTmvt7Dede
var T4sympTmvt7Dede = map[int]defStructs.Champ{
	0: {"DEDE_INF", 1, -1, nil},
	1: {"DEDE_NUM", -1, -1, nil},
}

// T4sympTmvt7Nos
var T4sympTmvt7Nos = map[int]defStructs.Champ{
	0: {"NOS_INF", 1, -1, nil},
	1: {"NOS_NUM", -1, -1, nil},
}

// T4sympTmvt7Cap1
var T4sympTmvt7Cap1 = map[int]defStructs.Champ{
	0: {"Cax", 1, -1, nil},
	1: {"Tap", 1, -1, nil},
	2: {"Dur", 2, -1, nil},
	3: {"Tax", 2, -1, nil},
}

// T4sympTmvt7Cap2
var T4sympTmvt7Cap2 = map[int]defStructs.Champ{
	0: {"Cax", 1, -1, nil},
	1: {"Tap", 1, -1, nil},
	2: {"Dur", 2, -1, nil},
	3: {"Tax", 2, -1, nil},
	4: {"Sup", 1, -1, nil},
}
