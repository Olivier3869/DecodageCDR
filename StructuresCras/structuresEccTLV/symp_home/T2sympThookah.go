package symp_home

import "local/structures/defStructs"

// T2symphThookah
var T2symphThookah = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3symphThookahdc},
}

// T3symphThookahdc
var T3symphThookahdc = map[int]defStructs.Champ{
	0x01: {"ICAR", -1, -1, nil},
	0x02: {"SEPTID", -1, 1, nil},
	0x03: {"STATUS", -1, 1, nil},
	0x04: {"VPNID", -1, 1, nil},
	0x05: {"INDX", -1, 1, nil},
	0x06: {"Duree", -1, -1, nil},
	0x07: {"TACC", -1, 1, nil},
	0x08: {"CDE", -1, -1, nil},
	0x09: {"IDUS", -1, -1, nil},
	0x0a: {"NORIG", -1, 1, nil},
	0x0b: {"SourceIP", -1, -1, nil},
	0x0c: {"ORIGGWTYPE", -1, -1, nil},
	0x0d: {"CGSITEID", -1, 1, nil},
	0x0e: {"CGSITETYPE", -1, 1, nil},
	0x10: {"DECO", -1, 1, nil},
	0x11: {"SAP1", -1, -1, nil},
	0x12: {"NUM1", -1, -1, nil},
	0x13: {"CDSITEID1", -1, 1, nil},
	0x14: {"CDSITETYPE1", -1, 1, nil},
	0x15: {"SAP2", -1, 1, nil},
	0x16: {"NUM2", -1, -1, nil},
	0x17: {"CDSITEID2", -1, 1, nil},
	0x18: {"CDSITETYPE2", -1, 1, nil},
	0x19: {"SAP3", -1, 1, nil},
	0x1a: {"CDSITEID3", -1, 1, nil},
	0x1b: {"CDSITETYPE3", -1, 1, nil},
	0x1c: {"DETA", -1, 1, nil},
	0x1d: {"DSON", -1, 1, nil},
	0x1e: {"DESTIPADDR", -1, 1, nil},
	0x1f: {"DESTGWTYPE", -1, 1, nil},
	0x20: {"INEF", -1, 1, nil},
	0x21: {"DISC", -1, 1, nil},
	0x22: {"CallId", -1, 1, nil},
	0x23: {"CallType", -1, -1, nil},
	0x24: {"OFFNETROUTINGSITE", -1, 1, nil},
	0x25: {"IVRTERMINATED", -1, 1, nil},
	0x26: {"IVRMSGID", -1, 1, nil},
	0x27: {"IVRADDR", -1, -1, nil},
	0x28: {"IVRREPLACEDCLI", -1, -1, nil},
	0x29: {"TEXT1", -1, 1, nil},
	0x2a: {"TEXT2", -1, 1, nil},
	0x2b: {"DTOAPPLIED", -1, 1, nil},
	0x2c: {"CALLEDSITECC", -1, -1, nil},
	0x2d: {"CALLEDPREFIX", -1, 1, nil},
	0x2e: {"DETR", -1, 1, nil},
	0x2f: {"CLI", -1, 1, nil},
	0x30: {"CALLEDPUBNDI", -1, -1, nil},
	0x31: {"ORIG", -1, -1, nil},
	0x32: {"PAI", -1, 1, nil},
	0x33: {"ORIGRECEIVEDNUMBER", -1, 1, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4sympAjoutPlatine},
}
