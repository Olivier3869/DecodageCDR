package structuresEccFIXE

import "local/structures/defStructs"

// TcentaureBIVcalltype
var TcentaureBIVcalltype = map[int]defStructs.Champ{
	0:  {"callType", 1, -1, nil},
	1:  {"recordingEntityType", 1, -1, nil},
	2:  {"recordingEntityIdentity", 9, -1, nil},
	3:  {"mobileSubscriberIdentity", 8, -1, nil},
	4:  {"callStartTimeStamp", 7, -1, nil},
	5:  {"callDuration", 4, -1, nil},
	6:  {"bearerService", 1, -1, nil},
	7:  {"teleservice", 1, -1, nil},
	8:  {"bearerCapabilities", 1, -1, nil},
	9:  {"systemtype", 1, -1, nil},
	10: {"VPN", 1, -1, nil},
	11: {"callOrigine", 1, -1, nil},
	12: {"efficiencyIndicator", 1, -1, nil},
	13: {"callTerminationType", 1, -1, nil},
	14: {"chargingIndicator", 1, -1, nil},
	15: {"UTnumbers", 1, -1, nil},
	16: {"chargingCharacteristics", 4, -1, nil},
	17: {"portability", -1, -1, nil},
}

// T4centaureBIVservicesInformations
var T4centaureBIVservicesInformations = map[int]defStructs.Champ{
	0x02: {"occurenceNumber", 1, -1, nil},
}

// T4centaureBIVASdata
var T4centaureBIVASdata = map[int]defStructs.Champ{
	0x01: {"IMSstatistics", -1, -1, nil},
	0x02: {"accessSessionIdentifier", -1, 1, nil},
	0x03: {"networkSessionIdentifier", -1, 1, nil},
	0x04: {"relatedCallId", -1, 1, nil},
	0x05: {"localcallId", -1, 1, nil},
	0x06: {"codecUsed", -1, 1, nil},
	0x07: {"companyIdentifier", -1, 1, nil},
	0x08: {"receivedCallNumber", -1, 1, nil},
}

// T4centaureBIVIMScore
var T4centaureBIVIMScore = map[int]defStructs.Champ{
	0x01: {"ASstatistics", -1, -1, nil},
	0x05: {"translatedCalledNumber", -1, 1, nil},
	0x06: {"originatingIOI", -1, 1, nil},
	0x07: {"terminatingIOI", -1, 1, nil},
	0x08: {"accessNetworkInformation", -1, 1, nil},
	0x09: {"sessionId", -1, 1, nil},
	0x10: {"S-CSCFidentifier", -1, 1, nil},
	0x11: {"ASidentifier", -1, 1, nil},
	0x12: {"SIPmethod", -1, 1, nil},
	0x13: {"privateUserId", -1, 1, nil},
}

// T4centaureBIVfixedWireNetwork
var T4centaureBIVfixedWireNetwork = map[int]defStructs.Champ{
	0x03: {"serviceCode", 1, 1, nil},
}

// T4centaureBIVCDRreception
var T4centaureBIVCDRreception = map[int]defStructs.Champ{
	0x20: {"CDRreceptionDate", 1, -1, nil},
}

// T3centaureBIV
var T3centaureBIV = map[int]defStructs.Champ{
	0x11: {"callTypeCharateristics", -1, -1, &TcentaureBIVcalltype},
	0x12: {"callPartnerIdentity", -1, -1, nil},
	0x14: {"chargeableCustomer", -1, -1, nil},
	0x17: {"circuitAllocationTimestamp", -1, -1, nil},
	0x27: {"callPartnerAddress", -1, 1, nil},
	0x28: {"MsAddress", -1, 1, nil},
	0x30: {"CDRreception", -1, -1, &T4centaureBIVCDRreception},
	0x3d: {"fixedWireNetwork", -1, -1, &T4centaureBIVfixedWireNetwork},
	0x53: {"redirectingNumber", -1, -1, nil},
	0x54: {"redirectingReason", -1, -1, nil},
	0x55: {"IMScore", -1, -1, &T4centaureBIVIMScore},
	0x56: {"ASdata", -1, -1, &T4centaureBIVASdata},
	0x35: {"servicesInformations", -1, -1, &T4centaureBIVservicesInformations},
	0x4d: {"servicesList", -1, -1, nil},
	0x5b: {"VASoperatorIndicator", -1, -1, nil},
	0x5c: {"VAScodeTarif", -1, -1, nil},
}

// T2centaureBIV
var T2centaureBIV = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3centaureBIV},
}

// TcentaureBIV
var TcentaureBIV = map[int]defStructs.Champ{
	0x5d: {"TcentaureBIV", -1, -1, &T2centaureBIV},
}

// DcentaureBIV
var DcentaureBIV = map[int]defStructs.TabRechercheTypeDc{
	0: {"TcentaureBIV", 0xffff0000, 0x805d0000, 351, &TcentaureBIV},
	1: {"TcentaureBIV", 0xffff0000, 0x80020000, 351, &TcentaureBIV},
}
