package structuresEecASN1

import "local/structures/defStructs"

// TTSms_mcohr
var TTSms_mcohr = map[int]defStructs.TabRechercheTypeDc{
	0: {"TTsms_mcohr", 0xff000000, 0x30000000, 0, &TSms_mcohr},
}

// TTSms_mcohr
var TSms_mcohr = map[int]defStructs.Champ{
	16: {"Tsms_mcohr", -1, 0, &CALLDetailRecord},
}

var CALLDetailRecord = map[int]defStructs.Champ{
	1:   {"ORIGADDRESSGSM", -1, 5, nil},
	3:   {"RECIPADDRESSGSM", -1, 5, nil},
	4:   {"SUBMITDATE", -1, 14, nil},
	5:   {"SUBMITTIME", -1, 15, nil},
	6:   {"STATUS", -1, 0, nil},
	7:   {"TERMINDATE", -1, 14, nil},
	8:   {"TERMINTIME", -1, 15, nil},
	9:   {"lengthOfMessage", -1, 0, nil},
	10:  {"prioIndicator", -1, 0, nil},
	14:  {"notifIndicator", -1, 0, nil},
	20:  {"dgtiAddressGSM", -1, 5, nil},
	25:  {"orglSubmitDate", -1, 14, nil},
	26:  {"orglSubmitTime", -1, 15, nil},
	29:  {"recipIntlMobileSubId", -1, 5, nil},
	97:  {"notifIndicator", -1, 0, nil},
	110: {"origSCAddress", -1, 0, &origSCAddressExt},
}

var origSCAddressExt = map[int]defStructs.Champ{
	0: {"TON", -1, 0, nil},
	1: {"NPI", -1, 0, nil},
	2: {"PID", -1, 0, nil},
	3: {"MSISDN", -1, 1, nil},
	4: {"MSISDNUTF8", -1, 1, nil},
}
