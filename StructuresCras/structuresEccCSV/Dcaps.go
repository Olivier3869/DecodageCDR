package structuresEccCSV

import "local/structures/defStructs"

// Tcaps
var Tcaps = map[int]defStructs.Champ{
	0: {"DOMAINE", -1, -1, nil},
	1: {"IDUS_NUM", -1, -1, nil},
	2: {"DEDE_NUM", -1, -1, nil},
	3: {"ICAR", -1, -1, nil},
	4: {"DUR", -1, -1, nil},
	5: {"SGTQS", -1, -1, nil},
	6: {"TYPE_EDIT", -1, -1, nil},
	7: {"CODE_TARIF", -1, -1, nil},
	8: {"codAppFact", -1, -1, nil},
}

// Tcaps
var Dcaps = map[int]defStructs.TabRechercheTypeDc{
	0: {"Tcaps", 0x00000000, 0x00000000, 0x3b, &Tcaps},
}
