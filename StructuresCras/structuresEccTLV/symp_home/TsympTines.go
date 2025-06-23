package symp_home

import (
	"local/structures/defStructs"
)

// TsympTinesT2
var TsympTinesT2 = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &TsympTinesT3},
}

// TsympTinesT3
var TsympTinesT3 = map[int]defStructs.Champ{
	0x01: {"SendDate", -1, 8, nil},
	0x02: {"accountPoid", -1, 1, nil},
	0x05: {"multiSessionId", -1, 1, nil},
	0x06: {"Serial", -1, 1, nil},
	0x07: {"startDate", -1, 8, nil},
	0x08: {"sendDatetz", -1, 1, nil},
	0x09: {"sendDatedst", -1, -1, nil},
	0x0a: {"eventType", -1, -1, nil},
	0x0b: {"EventDate", -1, 8, nil},
	0x0c: {"userLineId", -1, 1, nil},
	0x0d: {"SessionId", -1, 1, nil},
	0x0e: {"sessionType", -1, -1, nil},
	0x0f: {"userIpv6Addr", -1, 1, nil},
	0x10: {"userLogin", -1, 1, nil},
	0x11: {"DisconnectAuth", -1, -1, nil},
	0x12: {"sessionTime", -1, 1, nil},
	0x13: {"sessionTimeUnit", -1, 1, nil},
	0x14: {"userIpv4Addr", -1, -1, nil},
	0x15: {"nasIPv4Addr", -1, 1, nil},
	0x16: {"nasName", -1, 1, nil},
	0x17: {"nasTz", -1, 1, nil},
	0x19: {"accessProvider", -1, -1, nil},
	0x1a: {"theoLeasEndDate", -1, -1, nil},
	0x1b: {"circuitId", -1, 1, nil},
	0x1c: {"MACAddr", -1, -1, nil},
	0x20: {"addressProvider", -1, -1, nil},
	0x21: {"allocationMode", -1, -1, nil},
	0x24: {"accessMode", -1, -1, nil},
	0x26: {"roamingType", -1, -1, nil},
	0x27: {"callerId", -1, 1, nil},
	0x28: {"dnis", -1, -1, nil},
	0x29: {"eventCause", -1, -1, nil},
	0x2a: {"AuthentError", -1, -1, nil},
	0x2b: {"contextEndDate", -1, -1, nil},
	0x2c: {"contexEndDateIn", -1, -1, nil},
	0x30: {"synchroRate", -1, -1, nil},
	0x31: {"octetsIn", -1, 1, nil},
	0x32: {"octetsOut", -1, 1, nil},
	0x33: {"octetsInV4", -1, 1, nil},
	0x34: {"octetsInV6", -1, 1, nil},
	0x35: {"octetsOutV4", -1, 1, nil},
	0x36: {"octetsOutV6", -1, 1, nil},
	0x39: {"startDateIn", -1, 1, nil},
	0x3a: {"EventDateIn", -1, 1, nil},
	0x3b: {"theoLeEndDateIn", -1, -1, nil},
	0x3d: {"UNKNOWN", -1, -1, nil},
	0x40: {"sessionTimeH", -1, -1, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4TsympTines},
}

// T4TsympTines
var T4TsympTines = map[int]defStructs.Champ{
	0x50: {"SGTQS", 0, 1, nil},
	0x51: {"userIpAddr", -1, 1, nil},
}
