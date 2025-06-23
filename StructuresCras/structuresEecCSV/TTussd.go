package structuresEecCSV

import "local/structures/defStructs"

// Tussd
var Tussd = map[int]defStructs.Champ{
	0: {"MSISDN", -1, -1, nil},
	1: {"IMSI", -1, -1, nil},
	2: {"DTA", -1, -1, nil},
	3: {"Duree", -1, -1, nil},
	4: {"NUM", -1, -1, nil},
}

// TTussd
var TTussd = map[int]defStructs.TabRechercheTypeDc{
	0: {"TTussd", 0x00000000, 0x00000000, 0x2c, &Tussd},
}
