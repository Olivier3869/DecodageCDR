package structuresEecTLV

import "local/structures/defStructs"

// T3DATE_LOCALET3
var T3DATE_LOCALET3 = map[int]defStructs.Champ{
	0: {"DATE_LOCALE", -1, 12, nil},
}

// T3DATE_METROPOLET3
var T3DATE_METROPOLET3 = map[int]defStructs.Champ{
	0: {"DATE_METROPOLE", -1, 12, nil},
}

// T3NATACHT3
var T3NATACHT3 = map[int]defStructs.Champ{
	0: {"NATACH", -1, 0, nil},
}

var T3DEDE_NUMT3 = map[int]defStructs.Champ{
	0: {"DEDE_NUM", -1, 0, nil},
}

var T3IDUS_INFT3 = map[int]defStructs.Champ{
	0: {"IDUS_INF", -1, 0, nil},
}

var T3IDUS_NUMT3 = map[int]defStructs.Champ{
	0: {"IDUS_NUM", -1, 0, nil},
}

var T3ICCM_INFT3 = map[int]defStructs.Champ{
	0: {"ICCM_INF", -1, 0, nil},
}

var T3ICCM_NUMT3 = map[int]defStructs.Champ{
	0: {"ICCM_NUM", -1, 0, nil},
}

var T3DURT3 = map[int]defStructs.Champ{
	0: {"DUR", -1, 0, nil},
}

var T3REFT3 = map[int]defStructs.Champ{
	0: {"REF", -1, 0, nil},
}

var T3SEQT3 = map[int]defStructs.Champ{
	0: {"SEQ", -1, 0, nil},
}

var T3IPER_INFT3 = map[int]defStructs.Champ{
	0: {"IPER_INF", -1, 0, nil},
}

var T3IPER_NUMT3 = map[int]defStructs.Champ{
	0: {"IPER_NUM", -1, 0, nil},
}

var T3COD_OPERT3 = map[int]defStructs.Champ{
	0: {"COD_OPER", -1, 2, nil},
}

var T3COD_OFFRET3 = map[int]defStructs.Champ{
	0: {"COD_OFFRE", -1, 2, nil},
}

var T3TAGT3 = map[int]defStructs.Champ{
	0: {"TAG", -1, 0, nil},
}

// Tsvaip
var TsvaipT3 = map[int]defStructs.Champ{
	0x01: {"T3DATE_LOCALE", -1, 0, &T3DATE_LOCALET3},
	0x02: {"T3DATE_METROPOLE", -1, 0, &T3DATE_METROPOLET3},
	0x03: {"T3NATACH", -1, 0, &T3NATACHT3},
	0x04: {"T3DEDE_NUM", -1, 0, &T3DEDE_NUMT3},
	0x05: {"T3IDUS_INF", -1, 0, &T3IDUS_INFT3},
	0x06: {"T3IDUS_NUM", -1, 0, &T3IDUS_NUMT3},
	0x07: {"T3ICCM_INF", -1, 0, &T3ICCM_INFT3},
	0x08: {"T3ICCM_NUM", -1, 0, &T3ICCM_NUMT3},
	0x09: {"T3DUR", -1, 0, &T3DURT3},
	0x0a: {"T3REF", -1, 0, &T3REFT3},
	0x0b: {"T3SEQ", -1, 0, &T3SEQT3},
	0x0c: {"T3IPER_INF", -1, 0, &T3IPER_INFT3},
	0x0d: {"T3IPER_NUM", -1, 0, &T3IPER_NUMT3},
	0x0e: {"T3COD_OPER", -1, 0, &T3COD_OPERT3},
	0x0f: {"T3COD_OFFRE", -1, 0, &T3COD_OFFRET3},
	0x10: {"T3TAG", -1, 0, &T3TAGT3},
}

// Tsvaip
var TsvaipT2 = map[int]defStructs.Champ{
	0x02: {"TsvaipT3", -1, -1, &TsvaipT3},
}

// Tsvaip
var Tsvaip = map[int]defStructs.Champ{
	0x10: {"Tsvaip", -1, -1, &TsvaipT2},
}

// TTsvaip
var TTsvaip = map[int]defStructs.TabRechercheTypeDc{
	0: {"Tsvaip", 0xffff0000, 0x80100000, 0, &Tsvaip},
}
