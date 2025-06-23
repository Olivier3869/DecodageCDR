package symp_home

import "local/structures/defStructs"

// T2sympThookahSVAIPtag
var T2sympThookahSVAIPtag = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3sympThookahSVAIPtagdc},
}

// T3sympThookahSVAIPtagdc
var T3sympThookahSVAIPtagdc = map[int]defStructs.Champ{
	0x01: {"ICAR", -1, 5, nil},
	0x02: {"HFAPP", -1, -1, nil},
	0x03: {"NATACH", -1, -1, nil},
	0x04: {"DEDE_NUM", -1, -1, nil},
	0x05: {"DETR_NUM", -1, -1, nil},
	0x06: {"IDUS_INF", -1, -1, nil},
	0x07: {"IDUS_NUM", -1, -1, nil},
	0x08: {"ICCM_INF", -1, -1, nil},
	0x09: {"ICCM_NUM", -1, -1, nil},
	0x0a: {"CDEUR", -1, -1, nil},
	0x0b: {"HDCONV", -1, -1, nil},
	0x0c: {"DUR", -1, -1, nil},
	0x0d: {"DURSN", -1, -1, nil},
	0x0e: {"REF", -1, -1, nil},
	0x0f: {"CALLID", -1, 1, nil},
	0x10: {"SEQ", -1, -1, nil},
	0x11: {"IPER_INF", -1, -1, nil},
	0x12: {"IPER_NUM", -1, -1, nil},
	0x13: {"FINSER", -1, -1, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4sympAjoutPlatine},
}
