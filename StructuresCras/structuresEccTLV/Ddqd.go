package structuresEccTLV

import "local/structures/defStructs"

// T4dqdIPdcASservice
var T4dqdIPdcASservice = map[int]defStructs.Champ{
	0x01: {"ASserviceName", -1, -1, nil},
	0x02: {"ASfacResult", -1, -1, nil},
}

// T4dqdRIdcGLA2
var T4dqdRIdcGLA2 = map[int]defStructs.Champ{
	0: {"IDSV", 1, -1, nil},
	1: {"ICAR", 6, 5, nil},
	2: {"HFAPP", 3, 6, nil},
	3: {"CPCS", 2, -1, nil},
	4: {"CCAS", 2, -1, nil},
	5: {"NATACH", 1, -1, nil},
}

// T4dqdRIdcDEDE
var T4dqdRIdcDEDE = map[int]defStructs.Champ{
	0: {"INF", 2, -1, nil},
	1: {"NUM", -1, -1, nil},
}

// T4dqdRIdcDETR
var T4dqdRIdcDETR = map[int]defStructs.Champ{
	0: {"INT", 2, -1, nil},
	1: {"INF", 1, -1, nil},
	2: {"NUM", -1, -1, nil},
}

// T4dqdRIdcDE
var T4dqdRIdcDE = map[int]defStructs.Champ{
	0: {"INF", 1, -1, nil},
	1: {"NUM", -1, -1, nil},
}

// T4dqdRIdcMV
var T4dqdRIdcMV = map[int]defStructs.Champ{
	0: {"MV", 2, -1, nil},
	1: {"DURPAC", 2, -1, nil},
	2: {"DURACC", 2, -1, nil},
	3: {"DURaMV", -1, -1, nil},
}

// T4dqdRIdcTLA
var T4dqdRIdcTLA = map[int]defStructs.Champ{
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

// T4dqdRIdcDECGEO
var T4dqdRIdcDECGEO = map[int]defStructs.Champ{
	0: {"ZAA", 2, -1, nil},
	1: {"ZCT", 2, -1, nil},
	2: {"ZDT", 1, -1, nil},
	3: {"ZRG", -1, -1, nil},
}

// T4dqdRIdcRDIP
var T4dqdRIdcRDIP = map[int]defStructs.Champ{
	0: {"DATE", 5, -1, nil},
	1: {"CPAYS", 2, -1, nil},
	2: {"IDPCS", 4, -1, nil},
	3: {"REFC", -1, -1, nil},
}

// T4dqdRIdcENAP
var T4dqdRIdcENAP = map[int]defStructs.Champ{
	0: {"REF", 3, -1, nil},
	1: {"SEQ", -1, -1, nil},
}

// T4dqdRIdcTRABO
var T4dqdRIdcTRABO = map[int]defStructs.Champ{
	0: {"NUTC", 2, -1, nil},
	1: {"NUAA", -1, -1, nil},
}

// T4dqdRIdcLINK
var T4dqdRIdcLINK = map[int]defStructs.Champ{
	0: {"APPEL", 3, -1, nil},
	1: {"NOEUD", -1, -1, nil},
}

// T3dqdIPdc
var T3dqdIPdc = map[int]defStructs.Champ{
	0x03: {"NetworkInfo", -1, 1, nil},
	0x40: {"CallingNumber", -1, 1, nil},
	0x43: {"CalledNumber", -1, 1, nil},
	0x12: {"TranslatedNumber", -1, 1, nil},
	0x04: {"OrigCalledNumber", -1, 1, nil},
	0x05: {"RedirectingNumber", -1, 1, nil},
	0x06: {"RedirectingReason", -1, 1, nil},
	0x07: {"AScodec", -1, 1, nil},
	0x01: {"Date", -1, 8, nil},
	0x08: {"SessionEndTime", -1, 8, nil},
	0x13: {"Duree", -1, -1, nil},
	0x09: {"RoleOfNode", -1, -1, nil},
	0x20: {"Efficiency", -1, -1, nil},
	0x0a: {"Termination", -1, -1, nil},
	0x0b: {"ChargingId", -1, 1, nil},
	0x0c: {"ASreleasingParty", -1, -1, nil},
	0x11: {"CAasId", -1, 1, nil},
	0x0d: {"ChargeIndicator", -1, 1, nil},
	0x1a: {"LocalCallId", -1, 1, nil},
	0x1b: {"RelatedCallId", -1, 1, nil},
	0x1c: {"RealatedCallIdReason", -1, 1, nil},
	0x1d: {"AnswerIndicator", -1, 1, nil},
	0x1f: {"Q850cause", -1, -1, nil},
	0x30: {"ASservice1", -1, -1, &T4dqdIPdcASservice},
	0x31: {"ASservice2", -1, -1, &T4dqdIPdcASservice},
	0x32: {"ASservice3", -1, -1, &T4dqdIPdcASservice},
	0x33: {"ASservice4", -1, -1, &T4dqdIPdcASservice},
	0x34: {"ASservice5", -1, -1, &T4dqdIPdcASservice},
	0x16: {"AccessCallId", -1, 1, nil},
}

// T3dqdRIdc
var T3dqdRIdc = map[int]defStructs.Champ{
	0x02: {"GLA2", -1, -1, &T4dqdRIdcGLA2},
	0x4b: {"DEDEini", -1, -1, nil},
	0x4c: {"DEDEini2", -1, -1, nil},
	0x43: {"DEDE", -1, -1, &T4dqdRIdcDEDE},
	0x48: {"DEDE1", -1, -1, nil},
	0x27: {"DETR", -1, -1, &T4dqdRIdcDETR},
	0x2b: {"DETR1", -1, -1, &T4dqdRIdcDETR},
	0x28: {"DETN", -1, -1, &T4dqdRIdcDE},
	0x29: {"DETD", -1, -1, &T4dqdRIdcDE},
	0x40: {"IDUS", -1, -1, &T4dqdRIdcDE},
	0x41: {"ICCM", -1, -1, &T4dqdRIdcDE},
	0x37: {"CDPO", -1, -1, nil},
	0x07: {"CPTX", -1, -1, nil},
	0x09: {"MV", -1, -1, &T4dqdRIdcMV},
	0x0b: {"TLA", -1, -1, &T4dqdRIdcTLA},
	0x05: {"DECGEO", -1, -1, &T4dqdRIdcDECGEO},
	0x31: {"NCCONF", -1, -1, nil},
	0x35: {"PLT", -1, -1, nil},
	0x14: {"RDIP", -1, -1, &T4dqdRIdcRDIP},
	0x15: {"LISTE", -1, -1, nil},
	0x08: {"MNEMO", -1, 1, nil},
	0x44: {"ENAP", -1, -1, &T4dqdRIdcENAP},
	0x0a: {"TRABO", -1, -1, &T4dqdRIdcTRABO},
	0x2a: {"TRAD", -1, -1, nil},
	0x1a: {"LINK", -1, -1, &T4dqdRIdcLINK},
	0x4e: {"IPER", -1, -1, &T4dqdRIdcDE},
	0x4f: {"OADC", -1, -1, nil},
	0x50: {"TAG50?", -1, -1, nil},
}

// T2dqdIPdc
var T2dqdIPdc = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3dqdIPdc},
}

// T2dqdRIdc
var T2dqdRIdc = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3dqdRIdc},
}

// T1dqd
var T1dqd = map[int]defStructs.Champ{
	0x09: {"RIconvergent", -1, -1, &T2dqdRIdc},
	0x33: {"VoipSIP", -1, -1, &T2dqdIPdc},
}

// Ddqd
var Ddqd = map[int]defStructs.TabRechercheTypeDc{
	0: {"TdqdIP", 0xffff0000, 0x80330000, 0, &T1dqd},
	1: {"TdqdIPdc", 0xffff0000, 0x80020000, 0, &T1dqd},
	2: {"TdqdRI", 0xffff0000, 0x80090000, 0, &T1dqd},
	3: {"TdqdRIdc", 0xffff0000, 0x80020000, 0, &T1dqd},
	//4: {"TdqdRI??", 0xffff0000, 0x800a0000, 0, "T1dqd"},
}
