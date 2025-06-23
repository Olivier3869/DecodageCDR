package symp_home

import "local/structures/defStructs"

// T2symphTv2oipAS
var T2symphTv2oipAS = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3symphTv2oipAS},
}

// T3symphTv2oipAS
var T3symphTv2oipAS = map[int]defStructs.Champ{
	0x03: {"NetworkInfo", -1, 1, nil},
	0x40: {"CallingNumber", -1, 1, nil},
	0x43: {"CalledNumber", -1, 1, nil},
	0x12: {"TranslatedNumber", -1, 1, nil},
	0x04: {"OrigCalledNumber", -1, 1, nil},
	0x05: {"RedirectingNumber", -1, 1, nil},
	0x06: {"RedirectingReason", -1, 1, nil},
	0x07: {"AScodec", -1, 1, nil},
	0x01: {"Date", -1, 8, nil},
	0x08: {"SessionEndTime", -1, 8, nil},
	0x13: {"Duree", -1, -1, nil},
	0x09: {"RoleOfNode", -1, -1, nil},
	0x20: {"Efficiency", -1, -1, nil},
	0x0a: {"Termination", -1, -1, nil},
	0x0b: {"ChargingId", -1, 1, nil},
	0x0c: {"ASreleasingParty", -1, -1, nil},
	0x11: {"CAasId", -1, 1, nil},
	0x0d: {"ChargeIndicator", -1, 1, nil},
	0x1a: {"LocalCallId", -1, 1, nil},
	0x1b: {"RelatedCallId", -1, 1, nil},
	0x1c: {"RealatedCallIdReason", -1, 1, nil},
	0x1d: {"AnswerIndicator", -1, 1, nil},
	0x1f: {"Q850cause", -1, -1, nil},
	0x30: {"ASservice1", -1, -1, &T4symphTv2oipASservice},
	0x31: {"ASservice2", -1, -1, &T4symphTv2oipASservice},
	0x32: {"ASservice3", -1, -1, &T4symphTv2oipASservice},
	0x33: {"ASservice4", -1, -1, &T4symphTv2oipASservice},
	0x34: {"ASservice5", -1, -1, &T4symphTv2oipASservice},
	0x16: {"AccessCallId", -1, 1, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4sympAjoutPlatine},
}

// T4symphTv2oipASservice
var T4symphTv2oipASservice = map[int]defStructs.Champ{
	0x01: {"ASserviceName", -1, -1, nil},
	0x02: {"ASfacResult", -1, -1, nil},
}
