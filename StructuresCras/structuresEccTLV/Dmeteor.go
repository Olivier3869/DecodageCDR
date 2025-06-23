package structuresEccTLV

import "local/structures/defStructs"

// TmeteorTLVt3
var TmeteorTLVt3 = map[int]defStructs.Champ{
	0x01: {"DATE_LOCALE", -1, 8, nil},
	0x02: {"DATE_METROPOLE", -1, 8, nil},
	0x03: {"NATACH", -1, -1, nil},
	0x04: {"DEDE_NUM", -1, -1, nil},
	0x05: {"IDUS_INF", -1, -1, nil},
	0x06: {"IDUS_NUM", -1, -1, nil},
	0x07: {"ICCM_INF", -1, -1, nil},
	0x08: {"ICCM_NUM", -1, -1, nil},
	0x09: {"DUR", -1, -1, nil},
	0x0a: {"REF", -1, -1, nil},
	0x0b: {"SEQ", -1, -1, nil},
	0x0c: {"IPER_INF", -1, -1, nil},
	0x0d: {"IPER_NUM", -1, -1, nil},
	0x0e: {"COD_OPER", -1, 1, nil},
	0x0f: {"COD_OFFRE", -1, -1, nil},
	0x10: {"TAG", -1, -1, nil},
}

// TmeteorTLVt2
var TmeteorTLVt2 = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &TmeteorTLVt3},
}

// TmeteorTLVt1
var TmeteorTLVt1 = map[int]defStructs.Champ{
	0x10: {"sVAIPTAG", -1, -1, &TmeteorTLVt2},
}

// Dmeteor
var Dmeteor = map[int]defStructs.TabRechercheTypeDc{
	0: {"TmeteorTLVdc", 0xffff0000, 0x80100000, 0, &TmeteorTLVt1},
}
