package structuresEecCSV

import "local/structures/defStructs"

// TsmsoIP
var TsmsoIP = map[int]defStructs.Champ{
	0:  {"LOG_TIME", -1, -1, nil},
	1:  {"SM_TEXT_LENGTH", -1, -1, nil},
	2:  {"SM_SEGMENT_MAX", -1, -1, nil},
	3:  {"SM_SEGMENT_CURRENT", -1, -1, nil},
	4:  {"MSG_PARENT_ID", -1, -1, nil},
	5:  {"MSG_STATUS", -1, -1, nil},
	6:  {"SUBMIT_TIME", -1, -1, nil},
	7:  {"A_ADDR", -1, -1, nil},
	8:  {"A_IMSI", -1, -1, nil},
	9:  {"A_IMEI", -1, -1, nil},
	10: {"A_NETWORK_TYPE", -1, -1, nil},
	11: {"A_AIM_NAME", -1, -1, nil},
	12: {"A_SC", -1, -1, nil},
	13: {"OFR_A_UPLI", -1, -1, nil},
	14: {"OFR_A_NPLI", -1, -1, nil},
	15: {"DELIVERY_TIME", -1, -1, nil},
	16: {"B_ADDR", -1, -1, nil},
	17: {"B_IMSI", -1, -1, nil},
	18: {"B_IMEI", -1, -1, nil},
	19: {"B_NETWORK_TYPE", -1, -1, nil},
	20: {"B_AIM_NAME", -1, -1, nil},
	21: {"B_SC", -1, -1, nil},
	22: {"OFR_ASID", -1, -1, nil},
	23: {"OFR_B_UPLI", -1, -1, nil},
	24: {"OFR_B_NPLI", -1, -1, nil},
	25: {"OFR_PVNI", -1, -1, nil},
	26: {"MSG_ID", -1, -1, nil},
}

// TTsmsoIP
var TTsmsoIP = map[int]defStructs.TabRechercheTypeDc{
	0: {"TsmsoIP", 0x00000000, 0x00000000, 0x09, &TsmsoIP},
}
