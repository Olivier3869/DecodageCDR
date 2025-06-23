package symp_home

import "local/structures/defStructs"

// T2symphTetelephonie
var T2symphTetelephonie = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3symphTetelephonie},
}

// T3symphTetelephonie
var T3symphTetelephonie = map[int]defStructs.Champ{
	0x01: {"AuthorizeTime", -1, 7, nil},
	0x02: {"StartTime", -1, 7, nil},
	0x03: {"Duree", -1, -1, nil},
	0x10: {"SelectorIN", -1, -1, nil},
	0x11: {"SelectorVP", -1, -1, nil},
	0x12: {"SelectorSI", -1, -1, nil},
	0x13: {"SelectorSIID", -1, -1, nil},
	0x14: {"SelectorTZ", -1, -1, nil},
	0x15: {"SelectorInterco", -1, -1, nil},
	0x20: {"VtrunkIN", -1, -1, nil},
	0x21: {"VtrunkVP", -1, -1, nil},
	0x22: {"VtrunkSI", -1, -1, nil},
	0x23: {"VtrunkSIID", -1, 1, nil},
	0x24: {"VtrunkTZ", -1, 1, nil},
	0x30: {"CgidentityPU", -1, -1, nil},
	0x31: {"CgidentityPR", -1, -1, nil},
	0x33: {"CgidendityBI", -1, 1, nil},
	0x32: {"CdidentityPU", -1, -1, nil},
	0x34: {"CdidentityPR", -1, -1, nil},
	0x35: {"CdidentityBI", -1, 1, nil},
	0x40: {"OriginCaller", -1, -1, nil},
	0x41: {"TerminCallee", -1, -1, nil},
	0x42: {"Caller", -1, -1, nil},
	0x43: {"Callee", -1, -1, nil},
	0x44: {"TerminCause", -1, -1, nil},
	0x45: {"ConferenceId", -1, -1, nil},
	0x46: {"SimulationCall", -1, -1, nil},
	0x47: {"LegNumber", -1, 1, nil},
	0x48: {"SourceIP", -1, 1, nil},
	0x49: {"DestinationIP", -1, 1, nil},
	0x50: {"OriginCallee", -1, -1, nil},
	0x51: {"Service", -1, 1, nil},
	0x52: {"TerminCaller", -1, 1, nil},
	0x60: {"SuIdentifier", -1, 1, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4sympAjoutPlatine},
}
