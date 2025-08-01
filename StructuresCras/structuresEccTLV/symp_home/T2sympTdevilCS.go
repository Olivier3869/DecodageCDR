package symp_home

import "local/structures/defStructs"

// T2symphTdevilCS
var T2symphTdevilCS = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3symphTdevilCS},
}

// T3symphTdevilCS
var T3symphTdevilCS = map[int]defStructs.Champ{
	0x03: {"CodeEEC", -1, 1, nil},
	0x02: {"TYPENR", -1, 1, nil},
	0x50: {"CallId", -1, -1, nil},
	0x01: {"ICAR", -1, -1, nil},
	0x13: {"DUREE", -1, -1, nil},
	0x05: {"indrenv", -1, -1, nil},
	0x23: {"Calltype", -1, -1, nil},
	0x07: {"IdusNAI", -1, -1, nil},
	0x08: {"IdusNbPlan", -1, -1, nil},
	0x09: {"IdusIndRedir", -1, -1, nil},
	0x40: {"IdusNum", -1, -1, nil},
	0x0b: {"IdusInf", -1, -1, nil},
	0x54: {"LocIdusNAI", -1, -1, nil},
	0x55: {"LocIdusNbPlan", -1, -1, nil},
	0x56: {"LocIdusNum", -1, -1, nil},
	0x0d: {"DedeNbPlan", -1, -1, nil},
	0x12: {"DedeNum", -1, -1, nil},
	0x0f: {"DedeInf", -1, -1, nil},
	0x10: {"DestNAI", -1, -1, nil},
	0x11: {"DestNbPlan", -1, -1, nil},
	0x43: {"DETRNum", -1, -1, nil},
	0x44: {"DETRInf", -1, -1, nil},
	0x14: {"CallingPartyCateg", -1, -1, nil},
	0x3c: {"SUP", -1, -1, nil},
	0x16: {"ReroutInf", -1, -1, nil},
	0x17: {"incomingPort", -1, -1, nil},
	0x18: {"OutgoingPort", -1, -1, nil},
	0x19: {"IncomingPop", -1, -1, nil},
	0x1a: {"outgoingPOP", -1, -1, nil},
	0x1b: {"Outgoingmodule", -1, -1, nil},
	0x1c: {"CIC", -1, -1, nil},
	0x52: {"SourceIP", -1, -1, nil},
	0x53: {"DestIP", -1, -1, nil},
	0x1f: {"GRFSENT", -1, -1, nil},
	0x1e: {"GRFSSORT", -1, -1, nil},
	0x21: {"Anocode", -1, -1, nil},
	0x22: {"AnoInd", -1, -1, nil},
	0x25: {"locationechec", -1, -1, nil},
	0x20: {"INEF", -1, -1, nil},
	0x60: {"Status", -1, -1, nil},
	0x26: {"BIDTIME", -1, -1, nil},
	0x27: {"TCSDelay", -1, -1, nil},
	0x2a: {"DCSDelay", -1, -1, nil},
	0x28: {"CCDelay", -1, -1, nil},
	0x29: {"ICCDelay", -1, -1, nil},
	0x30: {"TransportDelay", -1, -1, nil},
	0x31: {"InterarrivalJiller", -1, -1, nil},
	0x5e: {"OUTPAC", -1, -1, nil},
	0x5c: {"OUTOCT", -1, -1, nil},
	0x5b: {"INOCT", -1, -1, nil},
	0x36: {"LostPacket", -1, -1, nil},
	0x37: {"LostPacketOut", -1, -1, nil},
	0x38: {"PacketingPeriod", -1, -1, nil},
	0x39: {"Coding", -1, -1, nil},
	0x3a: {"SilenceSup", -1, -1, nil},
	0x3b: {"DurIntPh", -1, -1, nil},
	0x3d: {"FSENT", -1, 1, nil},
	0x3e: {"FSSORT", -1, 1, nil},
	0x3f: {"DelCLIIdusNAI", -1, -1, nil},
	0x5d: {"INPAC", -1, -1, nil},
	0x41: {"DelCLIIdusNbPlan", -1, -1, nil},
	0x42: {"DelCLIIdusPubPriv", -1, -1, nil},
	0x32: {"DelCLIIdusNum", -1, -1, nil},
	0x45: {"OCNDEDENAI", -1, -1, nil},
	0x46: {"OCNDEDENbPlan", -1, -1, nil},
	0x33: {"OCNDEDENUM", -1, -1, nil},
	0x48: {"RDNInfRedirNAI", -1, -1, nil},
	0x49: {"RDNInfRedirNbPlan", -1, -1, nil},
	0x34: {"RDNInfRedirNUM", -1, -1, nil},
	0x35: {"CallEchoDevice", -1, -1, nil},
	0x0c: {"DedeNAI", -1, -1, nil},
	0x4e: {"Tag4e?", -1, 1, nil},
	0x4f: {"Tag4f?", -1, -1, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4sympAjoutPlatine},
}
