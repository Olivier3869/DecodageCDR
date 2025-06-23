package symp_home

import "local/structures/defStructs"

// T2sympTfe6
var T2sympTfe6 = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3sympTfe6dc},
	0x18: {"dea", -1, -1, &T3sympTfe6dea},
}

// T3sympTfe6dc
var T3sympTfe6dc = map[int]defStructs.Champ{
	0x04: {"G4", -1, -1, &TsymphTta6G4},
	0x1a: {"LINK", -1, -1, nil},
	0x1d: {"CALC", -1, -1, nil},
	0x40: {"IDUS", -1, -1, &TsymphTta6IDUS},
	0x43: {"DEDE", -1, -1, &TsymphTta6DEDE},
	0x45: {"IDAV", -1, -1, &TsymphTta6IDAV},
	0x44: {"COCA", -1, -1, nil},
	0x46: {"COCAV", -1, -1, nil},
	0x0a: {"CAP2", -1, -1, &TsymphTta6CAP2},
	0x06: {"CAB", -1, -1, nil},
	0x3d: {"CLE", -1, -1, nil},
	0x18: {"CLER", -1, -1, nil},
	0x0c: {"CHA", -1, -1, nil},
	0x1b: {"SEL", -1, -1, nil},
	0x1c: {"SON", -1, -1, nil},
	0x47: {"LOCUS", -1, -1, &TsymphTta6LOCUS},
	0x41: {"ICOM", -1, -1, &TsymphTta6ICOM},
	0x30: {"IPRV", -1, -1, &TsymphTta6IPRV},
	0x31: {"IDRV", -1, -1, &TsymphTta6IDRV},
	0x19: {"NBRV", -1, -1, nil},
	0x48: {"IDSV", -1, -1, nil},
	0x1f: {"reserve1f", -1, 2, nil},
	0x20: {"reserve20", -1, 2, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4sympAjoutPlatine},
}

// T3sympTfe6dea
var T3sympTfe6dea = map[int]defStructs.Champ{
	0x04: {"G4", -1, -1, &TsymphTta6G4},
	0x1a: {"LINK", -1, -1, nil},
	0x1d: {"CALC", -1, -1, nil},
	0x40: {"IDUS", -1, -1, &TsymphTta6IDUS},
	0x42: {"NOS", -1, -1, &TsymphTta6NOS},
	0x47: {"LOCUS", -1, -1, &TsymphTta6LOCUS},
	0x43: {"DEDE", -1, -1, &TsymphTta6DEDE},
	0x45: {"IDAV", -1, -1, &TsymphTta6IDAV},
	0x44: {"COCA", -1, -1, nil},
	0x46: {"COCAV", -1, -1, nil},
	0x30: {"IPRV", -1, -1, &TsymphTta6IPRV},
	0x31: {"IDRV", -1, -1, nil},
	0x19: {"NBRV", -1, -1, nil},
	0x32: {"DUR", -1, -1, nil},
	0x07: {"UT", -1, -1, nil},
	0x3d: {"CLE", -1, -1, nil},
	0x18: {"CLER", -1, -1, nil},
	0x1b: {"SEL", -1, -1, nil},
	0x1c: {"SON", -1, -1, nil},
	0x1f: {"reserve", -1, 2, nil},
	0x41: {"ICOM", -1, -1, &TsymphTta6ICOM},
	0x48: {"IDSV", -1, -1, nil},
	0x20: {"reserve20", -1, 2, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4sympAjoutPlatine},
}
