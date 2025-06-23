package structuresEccTLV

import "local/structures/defStructs"

// T4bivSIPajoutPlatine
var T4bivSIPajoutPlatine = map[int]defStructs.Champ{
	0x50: {"SGTQS", -1, 1, nil},
}

// T4bivSIPcallingPartyAdress
var T4bivSIPcallingPartyAdress = map[int]defStructs.Champ{
	0x01: {"callingPartyAddr02", -1, 1, nil},
}

// T4bivSIPserviceSpecificInformation
var T4bivSIPserviceSpecificInformation = map[int]defStructs.Champ{
	0x05: {"relatedCallID", -1, 1, nil},
	0x06: {"relatedCallIDreason", -1, -1, nil},
	0x07: {"localCallID", -1, 1, nil},
	0x08: {"ASoriginalCalledNumber", -1, 1, nil},
	0x09: {"ASredirectingNumber", -1, 1, nil},
	0x0a: {"redirectingReason", -1, -1, nil},
	0x0b: {"originalCalledReason", -1, -1, nil},
}

// T4bivSIPsessionMediaComponents
var T4bivSIPsessionMediaComponents = map[int]defStructs.Champ{
	0x00: {"timeOfMediaChange", -1, -1, nil},
	0x01: {"nameOfTheMedia", -1, 1, nil},
	0x02: {"descriptionOfTheMedia", -1, 1, nil},
	0x03: {"descriptionOfTheMedia", -1, 1, nil},
	0x04: {"descriptionOfTheMedia", -1, 1, nil},
	0x05: {"descriptionOfTheMedia", -1, 1, nil},
	0x06: {"descriptionOfTheMedia", -1, 1, nil},
	0x07: {"descriptionOfTheMedia", -1, 1, nil},
	0x08: {"descriptionOfTheMedia", -1, 1, nil},
	0x09: {"descriptionOfTheMedia", -1, 1, nil},
	0x0a: {"descriptionOfTheMedia", -1, 1, nil},
	0x0b: {"descriptionOfTheMedia", -1, 1, nil},
	0x0c: {"durationPerMedia", -1, -1, nil},
}

// T4bivSIPnodesIdentifier
var T4bivSIPnodesIdentifier = map[int]defStructs.Champ{
	0x00: {"SCSCFidentifier", -1, 1, nil},
	0x01: {"ASidentifier", -1, 1, nil},
}

// T4etelephonieAFDBajoutPlatine
var T4etelephonieAFDBajoutPlatine = map[int]defStructs.Champ{
	0x50: {"SGTQS", -1, 1, nil},
	0x53: {"localeDate", -1, -1, nil},
	0x61: {"calltype", -1, 1, nil},
}

// T4bticSIPAFDBcallingPartyAdress
var T4bticSIPAFDBcallingPartyAdress = map[int]defStructs.Champ{
	0x01: {"CallingPartyAddr02", -1, 1, nil},
}

// T4bticSIPAFDBserviceSpecificInformation
var T4bticSIPAFDBserviceSpecificInformation = map[int]defStructs.Champ{
	0x00: {"servicedata01", -1, 1, nil},
	0x02: {"serviceName", -1, 1, nil},
	0x03: {"ResultCode", -1, 1, nil},
	0x04: {"InvocationTime", -1, -1, nil},
	0x05: {"RelatedCallID", -1, 1, nil},
	0x06: {"RelatedCallIDreason", -1, -1, nil},
	0x07: {"LocalCallID", -1, 1, nil},
	0x08: {"ASoriginalCalledNumber", -1, 1, nil},
	0x09: {"ASredirectingrNumber", -1, 1, nil},
	0x0a: {"RedirectingReason", -1, -1, nil},
	0x0b: {"OriginalCalledReason", -1, -1, nil},
}

// T4bticSIPAFDBsessionMediaComponents
var T4bticSIPAFDBsessionMediaComponents = map[int]defStructs.Champ{
	0x00: {"TimeOfMediachange", -1, -1, nil},
	0x01: {"NameOfMedia", -1, -1, nil},
	0x02: {"MediaDescription01", -1, -1, nil},
	0x03: {"MediaDescription02", -1, -1, nil},
	0x04: {"MediaDescription03", -1, -1, nil},
	0x05: {"MediaDescription04", -1, -1, nil},
	0x06: {"MediaDescription05", -1, -1, nil},
	0x07: {"MediaDescription06", -1, -1, nil},
	0x08: {"MediaDescription07", -1, -1, nil},
	0x09: {"MediaDescription08", -1, -1, nil},
	0x0a: {"MediaDescription09", -1, -1, nil},
	0x0b: {"MediaDescription10", -1, -1, nil},
	0x0c: {"DurationPerMedia", -1, -1, nil},
}

// T4bticSIPAFDBnodesIdentifier
var T4bticSIPAFDBnodesIdentifier = map[int]defStructs.Champ{
	0x00: {"S-CSCFidentifier", -1, 1, nil},
	0x01: {"ApplicationServeridentifier", -1, 1, nil},
}

// T4bticSIPajoutPlatine
var T4bticSIPajoutPlatine = map[int]defStructs.Champ{
	0x50: {"SGTQS", -1, 1, nil},
	0x53: {"localeDate", -1, -1, nil},
	0x61: {"calltype", -1, 1, nil},
}

// TbivSIPT3
var TbivSIPT3 = map[int]defStructs.Champ{
	0x04: {"callingPartyAdress", -1, -1, &T4bivSIPcallingPartyAdress},
	0x05: {"requestedPartyAddress", -1, 1, nil},
	0x06: {"translatedCalledNumber", -1, 1, nil},
	0x0a: {"serviceRequestTimestamp", -1, 3, nil},
	0x0c: {"callStartTimestamp", -1, 3, nil},
	0x0d: {"callEndTimeStamp", -1, 3, nil},
	0x0e: {"callDuration", -1, -1, nil},
	0x11: {"roleOfTheNode", -1, -1, nil},
	0x12: {"callStatus", -1, -1, nil},
	0x13: {"sessionTerminationCause", -1, -1, nil},
	0x14: {"ASterminationCause", -1, -1, nil},
	0x15: {"IMSchargingIdentifier", -1, -1, nil},
	0x16: {"originatingIOI", -1, -1, nil},
	0x17: {"terminatingIOI", -1, -1, nil},
	0x19: {"accessSIPsession", -1, 1, nil},
	0x1a: {"networkSIPsession", -1, 1, nil},
	0x1b: {"accessNetworkInformation", -1, 1, nil},
	0x1c: {"causeCode", -1, -1, nil},
	0x1d: {"causeForRecordClosing", -1, -1, nil},
	0x1e: {"answerIndicator", -1, -1, nil},
	0x1f: {"sessionId", -1, 1, nil},
	0x21: {"trunkGroupInfo", -1, 1, nil},
	0x30: {"serviceSpecificInformation", -1, -1, &T4bivSIPserviceSpecificInformation},
	0x50: {"codecUsed", -1, 1, nil},
	0x52: {"sessionMediaComponents", -1, -1, &T4bivSIPsessionMediaComponents},
	0x70: {"nodesIdentifier", -1, -1, &T4bivSIPnodesIdentifier},
	0x71: {"ASreleasingParty", -1, 1, nil},
	0x73: {"companyIdentifier", -1, 1, nil},
	0x74: {"group", -1, -1, nil},
	0xa0: {"serviceComputedInfo01", -1, 1, nil},
	0xa1: {"serviceComputedInfo02", -1, 1, nil},
	0xa2: {"serviceComputedInfo03", -1, 1, nil},
	0xa3: {"serviceComputedInfo04", -1, 1, nil},
	0xa4: {"serviceComputedInfo05", -1, 1, nil},
	0xa5: {"serviceComputedInfo06", -1, 1, nil},
	0xa6: {"serviceComputedInfo07", -1, 1, nil},
	0xa7: {"serviceComputedInfo08", -1, 1, nil},
	0x77: {"SIPmethod", -1, -1, nil},
	0x78: {"privateUserID", -1, 1, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4bivSIPajoutPlatine},
}

// T3etelephonieAFDB
var T3etelephonieAFDB = map[int]defStructs.Champ{
	0x01: {"authoriseTime", -1, -1, nil},
	0x02: {"startTime", -1, -1, nil},
	0x03: {"duree", -1, -1, nil},
	0x10: {"SelectorIN", -1, -1, nil},
	0x11: {"SelectorVP", -1, -1, nil},
	0x12: {"SelectorSI", -1, -1, nil},
	0x13: {"SelectorSIID", -1, -1, nil},
	0x14: {"SelectorTZ", -1, -1, nil},
	0x15: {"SelectorInterco", -1, -1, nil},
	0x20: {"VtrunkIN", -1, -1, nil},
	0x21: {"VtrunkVP", -1, -1, nil},
	0x22: {"VtrunkSI", -1, -1, nil},
	0x23: {"VtrunkSIID", -1, 1, nil},
	0x24: {"VtrunkTZ", -1, 1, nil},
	0x30: {"CgidentityPU", -1, -1, nil},
	0x31: {"CgidentityPR", -1, -1, nil},
	0x32: {"CdidentityPU", -1, 1, nil},
	0x33: {"CdidentityBI", -1, -1, nil},
	0x34: {"CdidentityPR", -1, -1, nil},
	0x35: {"CdidentityBI", -1, 1, nil},
	0x40: {"OriginCaller", -1, -1, nil},
	0x41: {"TerminCallee", -1, -1, nil},
	0x42: {"Caller", -1, -1, nil},
	0x43: {"Callee", -1, -1, nil},
	0x44: {"TerminatingCause", -1, -1, nil},
	0x45: {"ConferenceId", -1, -1, nil},
	0x46: {"SimulationCall", -1, -1, nil},
	0x47: {"LegNumber", -1, 1, nil},
	0x48: {"SourceIP", -1, 1, nil},
	0x49: {"DestinationIP", -1, 1, nil},
	0x50: {"OriginCallee", -1, -1, nil},
	0x51: {"Service", -1, 1, nil},
	0x52: {"TerminCaller", -1, -1, nil},
	0x60: {"SUIdentifier", -1, 1, nil},
	0x7e: {"ajoutPlatine", -1, 1, &T4etelephonieAFDBajoutPlatine},
}

// T3bticSIPAFDB
var T3bticSIPAFDB = map[int]defStructs.Champ{
	0x03: {"AScallType", -1, -1, nil},
	0x04: {"CallingPartyAdress", -1, -1, &T4bticSIPAFDBcallingPartyAdress},
	0x05: {"RequestedPartyAddress", -1, 1, nil},
	0x06: {"TranslatedCalledNumber", -1, 1, nil},
	0x0a: {"ServiceRequestTimestamp", -1, 4, nil},
	0x0c: {"CallStartTimestamp", -1, 4, nil},
	0x0d: {"CallEndTimeStamp", -1, 4, nil},
	0x0e: {"CallDuration", -1, -1, nil},
	0x11: {"RoleOfNode", -1, -1, nil},
	0x12: {"CallStatus", -1, -1, nil},
	0x13: {"SessionTerminationCause", -1, -1, nil},
	0x14: {"ASterminationCause", -1, -1, nil},
	0x15: {"IMSchargingIdentifier", -1, 1, nil},
	0x16: {"OriginatingIOI", -1, 1, nil},
	0x17: {"TerminatingIOI", -1, 1, nil},
	0x19: {"AccessSIPsession", -1, 1, nil},
	0x1a: {"NetworkSIPsession", -1, 1, nil},
	0x1b: {"AccessNetworkInformation", -1, 1, nil},
	0x1c: {"CauseCode", -1, -1, nil},
	0x1d: {"CauseForRecordClosing", -1, -1, nil},
	0x1e: {"AnswerIndicator", -1, -1, nil},
	0x1f: {"SessionId", -1, 1, nil},
	0x30: {"ServiceSpecificInformation", -1, -1, &T4bticSIPAFDBserviceSpecificInformation},
	0x50: {"CodecUsed", -1, 1, nil},
	0x52: {"SessionMediaComponents", -1, -1, &T4bticSIPAFDBsessionMediaComponents},
	0x70: {"NodesIdentifier", -1, -1, &T4bticSIPAFDBnodesIdentifier},
	0x71: {"ASreleasingParty", -1, 1, nil},
	0x73: {"Companyidentifier", -1, 1, nil},
	0x74: {"Group", -1, 1, nil},
	0x77: {"SIPmethod", -1, 1, nil},
	0x78: {"PrivateUserID", -1, 1, nil},
	0xa0: {"ServiceComputedInfo01", -1, 1, nil},
	0xa1: {"ServiceComputedInfo02", -1, 1, nil},
	0xa3: {"ServiceComputedInfo03", -1, 1, nil},
	0xa4: {"ServiceComputedInfo04", -1, 1, nil},
	0xa5: {"ServiceComputedInfo05", -1, 1, nil},
	0xa6: {"ServiceComputedInfo06", -1, 1, nil},
	0xa7: {"ServiceComputedInfo07", -1, 1, nil},
	0xa8: {"ServiceComputedInfo08", -1, 1, nil},
	0xa9: {"ServiceComputedInfo09", -1, 1, nil},
	0xb0: {"ServiceComputedInfo10", -1, 1, nil},
	0xb1: {"ServiceComputedInfo11", -1, 1, nil},
	0xb2: {"ServiceComputedInfo12", -1, 1, nil},
	0xb3: {"ServiceComputedInfo13", -1, 1, nil},
	0xb4: {"ServiceComputedInfo14", -1, 1, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4bticSIPajoutPlatine},
}

// T2bivSIPAFDB
var T2bivSIPAFDB = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &TbivSIPT3},
}

// T2eTelephonieAFDB
var T2eTelephonieAFDB = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3etelephonieAFDB},
}

// T2bticSIPAFDB
var T2bticSIPAFDB = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3bticSIPAFDB},
}

// TvoipGS
var TvoipGS = map[int]defStructs.Champ{
	0x4a: {"voipGS", -1, -1, nil},
}

// TbivSIP
var TbivSIP = map[int]defStructs.Champ{
	0x77: {"bivSIP", -1, -1, &T2bivSIPAFDB},
}

// TeTelephonie
var TeTelephonie = map[int]defStructs.Champ{
	0x43: {"eTelephonie", -1, -1, &T2eTelephonieAFDB},
}

// TbticSIP
var TbticSIP = map[int]defStructs.Champ{
	0x57: {"bticSIP", -1, -1, &T2bticSIPAFDB},
}

// Dafdb
var Dafdb = map[int]defStructs.TabRechercheTypeDc{
	0: {"TvoipGS", 0xffff0000, 0x804a0000, 0, &TvoipGS},
	1: {"TbivSIP", 0xffff0000, 0x80770000, 0, &TbivSIP},
	2: {"TeTelephonie", 0xffff0000, 0x80430000, 0, &TeTelephonie},
	3: {"TbticSIP", 0xffff0000, 0x80570000, 0, &TbticSIP},
}
