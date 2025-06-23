package symp_home

import "local/structures/defStructs"

// T2symphTsrnbi
var T2symphTsrnbi = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3symphTsrnbidc},
}

// T3symphTsrnbidc
var T3symphTsrnbidc = map[int]defStructs.Champ{
	0x40: {"CallinNb", -1, 1, nil},
	0x43: {"CalledNb", -1, 1, nil},
	0x01: {"Date", -1, 8, nil},
	0x08: {"sessionEndTime", -1, 8, nil},
	0x13: {"Duree", -1, -1, nil},
	0x09: {"RoleOfNode", -1, -1, nil},
	0x20: {"Efficiency", -1, -1, nil},
	0x0b: {"ChargingId", -1, -1, nil},
	0x0e: {"SBCnodeId", -1, 1, nil},
	0x0f: {"TermIPaddr", -1, -1, nil},
	0x1f: {"Q850cause", -1, -1, nil},
	0x16: {"AccessCallId", -1, -1, nil},
	0x50: {"IngCallId", -1, 1, nil},
	0x51: {"EgrCallId", -1, 1, nil},
	0x52: {"CallingStationId", -1, 1, nil},
	0x53: {"CalledStationId", -1, 1, nil},
	0x54: {"H323setupTime", -1, 1, nil},
	0x55: {"H323ConnectTime", -1, 1, nil},
	0x56: {"H323DisconnectTime", -1, 1, nil},
	0x57: {"IngRealm", -1, 1, nil},
	0x58: {"IngRemoteAddr", -1, 1, nil},
	0x59: {"EgrRemoteAddr", -1, 1, nil},
	0x5a: {"SIPstatus", -1, 1, nil},
	0x5b: {"EgrRealm", -1, 1, nil},
	0x5c: {"ChargingFctAddr", -1, -1, nil},
	0x5d: {"PostDialDelay", -1, -1, nil},
	0x5e: {"SipDiversion", -1, -1, nil},
	0x5f: {"SessionDispo", -1, -1, nil},
	0x60: {"DisconnectInit", -1, -1, nil},
	0x61: {"FinalRoutingNb", -1, 1, nil},
	0x62: {"VSA200", -1, 1, nil},
	0x63: {"VSA201", -1, 1, nil},
	0x64: {"VSA202", -1, 1, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4sympAjoutPlatine},
}
