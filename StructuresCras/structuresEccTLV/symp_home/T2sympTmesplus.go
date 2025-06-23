package symp_home

import "local/structures/defStructs"

// T2symphTmesplus
var T2symphTmesplus = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3symphTmesplusdc},
}

// T3symphTmesplusdc
var T3symphTmesplusdc = map[int]defStructs.Champ{
	0x03: {"G3", -1, -1, &TsymphTta6G3},
	0x40: {"IDUS", -1, -1, &TsymphTta6IDUS},
	0x43: {"DEDE", -1, -1, &TsymphTta6DEDE},
	0x42: {"NOS", -1, -1, &TsymphTta6NOS},
	0x08: {"CAP1", -1, -1, &TsymphTta6CAP1},
	0x78: {"ETAP", -1, -1, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4sympAjoutPlatine},
}
