package structuresEecCSV

import "local/structures/defStructs"

// Tines
var Tlmisf = map[int]defStructs.Champ{
	0:  {"Node_Type", -1, -1, nil},
	1:  {"Sess_Type", -1, -1, nil},
	2:  {"Sess_Dir", -1, -1, nil},
	3:  {"Sess_Id", -1, -1, nil},
	4:  {"Sess_Start_time", -1, -1, nil},
	5:  {"Sess_Dur_Vol", -1, -1, nil},
	6:  {"Calling_Id", -1, -1, nil},
	7:  {"Called_Id", -1, -1, nil},
	8:  {"CallForward_service", -1, -1, nil},
	9:  {"IMEI", -1, -1, nil},
	10: {"IMSI", -1, -1, nil},
	11: {"Location", -1, -1, nil},
	12: {"SMSC_Addr", -1, -1, nil},
	13: {"H_Network_id", -1, -1, nil},
	14: {"V_Network_id", -1, -1, nil},
	15: {"RAT_Type", -1, -1, nil},
}

// TTines
var TTlmisf = map[int]defStructs.TabRechercheTypeDc{
	0: {"TTlmisf", 0x00000000, 0x00000000, 0x2c, &Tlmisf},
}
