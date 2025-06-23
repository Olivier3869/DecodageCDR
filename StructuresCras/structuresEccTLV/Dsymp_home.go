package structuresEccTLV

import (
	"local/structures/defStructs"
	"local/structures/structuresEccTLV/symp_home"
)

// Tsymp_h
var Tsymp_h = map[int]defStructs.Champ{
	0x6f: {"TsympMvf", -1, -1, &symp_home.TsympMvfT2},
	0x80: {"TsympTines", -1, -1, &symp_home.TsympTinesT2},
}

// Dsymp_home
var Dsymp_home = map[int]defStructs.TabRechercheTypeDc{
	0: {"Dsymp_h", 0xffff0000, 0x806f0000, 0, &Tsymp_h},
	1: {"Dsymp_hdc", 0xffff0000, 0x80020000, 0, &Tsymp_h},
	2: {"Dsymp_h", 0xffff0000, 0x81000000, 0, &Tsymp_h},
	3: {"Dsymp_hdc", 0xffff0000, 0x80020000, 0, &Tsymp_h},
}
