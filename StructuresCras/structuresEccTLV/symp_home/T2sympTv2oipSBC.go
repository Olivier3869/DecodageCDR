package symp_home

import "local/structures/defStructs"

// T2symphTv2oipSBC
var T2symphTv2oipSBC = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3symphTv2oipSBC},
}

// T3symphTv2oipSBC
var T3symphTv2oipSBC = map[int]defStructs.Champ{
	0x40: {"CallinNb", -1, 1, nil},
	0x43: {"CalledNb", -1, 1, nil},
	0x01: {"Date", -1, 8, nil},
	0x08: {"sessionEndTime", -1, 8, nil},
	0x13: {"Duree", -1, -1, nil},
	0x09: {"RoleOfNode", -1, -1, nil},
	0x20: {"Efficiency", -1, -1, nil},
	0x0b: {"ChargingId", -1, 1, nil},
	0x0e: {"SBCnodeId", -1, 1, nil},
	0x0f: {"TermIPaddr", -1, 1, nil},
	0x1f: {"Q850cause", -1, -1, nil},
	0x44: {"CallingFlow", -1, -1, &T4symphTv2oipSBCCallingFlow},
	0x45: {"CalledFlow", -1, 1, &T4symphTv2oipSBCCalledFlow},
	0x16: {"AccessCallId", -1, 1, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4sympAjoutPlatine},
}

// T4symphTv2oipSBCCallingFlow
var T4symphTv2oipSBCCallingFlow = map[int]defStructs.Champ{
	0x01: {"Octets", -1, -1, nil},
	0x02: {"Packets", -1, -1, nil},
	0x03: {"RTCPavgJitter", -1, -1, nil},
	0x04: {"RTCPavgLatency", -1, -1, nil},
	0x05: {"RTCPmaxJitter", -1, -1, nil},
	0x06: {"RTCPmaxLatency", -1, -1, nil},
	0x07: {"RTCPpacketsLost", -1, -1, nil},
	0x08: {"RTPavgJitter", -1, -1, nil},
	0x09: {"RTPmaxJitter", -1, -1, nil},
	0x0a: {"RTPpacketsLost", -1, -1, nil},
	0x0b: {"Flowtype", -1, 1, nil},
}

// T4symphTv2oipSBCCalledFlow
var T4symphTv2oipSBCCalledFlow = map[int]defStructs.Champ{
	0x01: {"Octets", -1, -1, nil},
	0x02: {"Packets", -1, -1, nil},
	0x03: {"RTCPavgJitter", -1, -1, nil},
	0x04: {"RTCPavgLatency", -1, -1, nil},
	0x05: {"RTCPmaxJitter", -1, -1, nil},
	0x06: {"RTCPmaxLatency", -1, -1, nil},
	0x07: {"RTCPpacketsLost", -1, -1, nil},
	0x08: {"RTPavgJitter", -1, -1, nil},
	0x09: {"RTPmaxJitter", -1, -1, nil},
	0x0a: {"RTPpacketsLost", -1, -1, nil},
	0x0b: {"Flowtype", -1, 1, nil},
}
