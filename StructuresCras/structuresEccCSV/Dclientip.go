package structuresEccCSV

import "local/structures/defStructs"

// Tclientip
var Tclientip = map[int]defStructs.Champ{
	0:  {"eventType", -1, -1, nil},
	1:  {"sessionType", -1, -1, nil},
	2:  {"userLogin", -1, -1, nil},
	3:  {"sessionId", -1, -1, nil},
	4:  {"circuitId", -1, -1, nil},
	5:  {"userLineId", -1, -1, nil},
	6:  {"userIpAddress", -1, -1, nil},
	7:  {"startDate", -1, -1, nil},
	8:  {"sessionTime", -1, -1, nil},
	9:  {"eventDate", -1, -1, nil},
	10: {"accessProvider", -1, -1, nil},
}

// Dclientip
var Dclientip = map[int]defStructs.TabRechercheTypeDc{
	0: {"Tclientip", 0x00000000, 0x00000000, 0x7c, &Tclientip},
}
