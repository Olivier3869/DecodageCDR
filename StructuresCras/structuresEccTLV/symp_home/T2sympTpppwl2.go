package symp_home

import "local/structures/defStructs"

// T2sympTpppwl2
var T2sympTpppwl2 = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3sympTpppwl2dc},
}

// T3sympTpppwl2dc
var T3sympTpppwl2dc = map[int]defStructs.Champ{
	0x01: {"sendtime", -1, -1, nil},
	0x02: {"accpoid", -1, 1, nil},
	0x03: {"accinstance", -1, 1, nil},
	0x04: {"servpoid", -1, 1, nil},
	0x05: {"servinstance", -1, 1, nil},
	0x06: {"serial", -1, 1, nil},
	0x07: {"starttime", -1, -1, nil},
	0x08: {"starttimeTZ", -1, 1, nil},
	0x09: {"starttimedst", -1, 1, nil},
	0x10: {"ppplogin", -1, 1, nil},
	0x11: {"chechpointed", -1, 1, nil},
	0x12: {"sestimevalue", -1, 1, nil},
	0x13: {"sestimeunit", -1, 1, nil},
	0x14: {"useripaddr", -1, 1, nil},
	0x15: {"nasipaddr", -1, 1, nil},
	0x16: {"nasname", -1, 1, nil},
	0x17: {"nastimezone", -1, 1, nil},
	0x18: {"nor", -1, 1, nil},
	0x19: {"fr", -1, 1, nil},
	0x20: {"fa", -1, 1, nil},
	0x21: {"mad", -1, 1, nil},
	0x22: {"ir", -1, 1, nil},
	0x23: {"ra", -1, 1, nil},
	0x24: {"mac", -1, 1, nil},
	0x25: {"ca", -1, 1, nil},
	0x26: {"roamintype", -1, 1, nil},
	0x27: {"callerid", -1, 1, nil},
	0x28: {"dnis", -1, 1, nil},
	0x29: {"disconcause", -1, 1, nil},
	0x30: {"modemspeed", -1, 1, nil},
	0x31: {"octestsin", -1, 1, nil},
	0x32: {"octestout", -1, 1, nil},
	0x33: {"packetsin", -1, 1, nil},
	0x34: {"packetsout", -1, 1, nil},
	0x35: {"unitpriceval", -1, 1, nil},
	0x36: {"unitpriceunit", -1, 1, nil},
	0x37: {"vat", -1, 1, nil},
	0x38: {"ratio", -1, 1, nil},
	0x39: {"strattimeS", -1, 1, nil},
	0x40: {"sestimevalH", -1, 1, nil},
	0x41: {"koctestsin", -1, 1, nil},
	0x42: {"koctestsout", -1, 1, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4sympAjoutPlatine},
}
