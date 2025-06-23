package symp_home

import "local/structures/defStructs"

// TsympMvfT2
var TsympMvfT2 = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &TsympMvfT3},
}

// TsympMvfT3
var TsympMvfT3 = map[int]defStructs.Champ{
	0x01: {"Appli", -1, -1, nil},
	0x02: {"IcarDeb", -1, 8, nil},
	0x04: {"Inef", -1, -1, nil},
	0x05: {"Idus", -1, -1, &TsympMvfT3Idus},
	0x06: {"IdusMail", -1, 1, nil},
	0x07: {"IdusGene", -1, -1, nil},
	0x08: {"IdusPriv", -1, -1, nil},
	0x09: {"App", -1, -1, &TsympMvfT3App},
	0x0a: {"Dede", -1, -1, &TsympMvfT3Dede},
	0x0b: {"DedeMail", -1, 1, nil},
	0x0c: {"DedeGene", -1, -1, nil},
	0x0d: {"DedePriv", -1, -1, nil},
	0x0f: {"Origine", -1, 1, nil},
	0x10: {"Destination", -1, 1, nil},
	0x11: {"Domaine", -1, 1, nil},
	0x12: {"Dur", -1, -1, nil},
	0x13: {"Fact", -1, -1, nil},
	0x14: {"Etap", -1, -1, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4sympHomeAjout},
}

// T4sympMvfAjout
var T4sympHomeAjout = map[int]defStructs.Champ{
	0x00: {"CCD", 0, -1, nil},
	0x50: {"SGTQS", -1, -1, nil},
	0x51: {"ECC_FACT", -1, -1, nil},
	0x53: {"DATE_LOCALE", -1, 7, nil},
	0x6d: {"OPE_VGA", -1, -1, nil},
	0x6e: {"OFF_VGA", -1, -1, nil},
}

// TsympMvfT3IA
var TsympMvfT3Dede = map[int]defStructs.Champ{
	0: {"Dede_Inf", 1, -1, nil},
	1: {"Dede_Num", -1, -1, nil},
}

var TsympMvfT3Idus = map[int]defStructs.Champ{
	0: {"Idus_Inf", 1, -1, nil},
	1: {"Idus_Num", -1, -1, nil},
}

var TsympMvfT3App = map[int]defStructs.Champ{
	0: {"App_Inf", 1, -1, nil},
	1: {"App_Num", -1, -1, nil},
}
