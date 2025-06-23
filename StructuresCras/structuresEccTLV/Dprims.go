package structuresEccTLV

import "local/structures/defStructs"

// T4VOLTEmobileID
var T4VOLTEmobileID = map[int]defStructs.Champ{
	0x00: {"IMSI", -1, 1, nil},
	0x01: {"MSISDN", -1, 1, nil},
	0x02: {"ggsnAddr", -1, 1, nil},
	0x03: {"sgsnAddr", -1, 1, nil},
}

// T4VOLTEnodesIdentifier
var T4VOLTEnodesIdentifier = map[int]defStructs.Champ{
	0x00: {"SCSCFidentifier", -1, 1, nil},
	0x01: {"ASidentifier", -1, 1, nil},
}

// T4VOLTEsDPmediaComponent1
var T4VOLTEsDPmediaComponent1 = map[int]defStructs.Champ{
	0x01: {"SDPmediaName", -1, 1, nil},
	0x02: {"SDPmediaDesc10", -1, 1, nil},
}

// T4VOLTEcallingPartyAdress
var T4VOLTEcallingPartyAdress = map[int]defStructs.Champ{
	0x01: {"callingPartyAddr02", -1, 1, nil},
}

// T3VOLTE
var T3VOLTE = map[int]defStructs.Champ{
	0x02: {"recordSequenceNumber", -1, -1, nil},
	0x04: {"callingPartyAdress", -1, -1, &T4VOLTEcallingPartyAdress},
	0x05: {"requestedPartyAddress", -1, 1, nil},
	0x06: {"translatedCalledNumber", -1, 1, nil},
	0x0a: {"serviceRequestTimestamp", -1, 4, nil},
	0x0c: {"callStartTimestamp", -1, 4, nil},
	0x0d: {"callEndTimeStamp", -1, 4, nil},
	0x0e: {"callDuration", -1, -1, nil},
	0x11: {"roleOfTheNode", -1, -1, nil},
	0x12: {"callStatus", -1, -1, nil},
	0x13: {"sessionTerminationCause", -1, -1, nil},
	0x14: {"ASterminationCause", -1, -1, nil},
	0x15: {"IMSchargingIdentifier", -1, 1, nil},
	0x16: {"originatingIOI", -1, -1, nil},
	0x17: {"terminatingIOI", -1, -1, nil},
	0x1b: {"accessNetworkInformation", -1, 1, nil},
	0x1d: {"causeForRecordClosing", -1, -1, nil},
	0x1f: {"sessionId", -1, 1, nil},
	0x52: {"SDPmediaComponent1", -1, 1, &T4VOLTEsDPmediaComponent1},
	0x70: {"nodesIdentifier", -1, -1, &T4VOLTEnodesIdentifier},
	0xa6: {"serviceComputedInfo07", -1, 1, nil},
	0xa7: {"serviceComputedInfo08", -1, 1, nil},
	0x77: {"SIPmethod", -1, -1, nil},
	0x78: {"privateUserID", -1, 1, nil},
	0x7e: {"sigVolum", -1, -1, nil},
	0x7f: {"totalVolum", -1, -1, nil},
	0x81: {"extBearVolum", -1, -1, nil},
	0x82: {"mobileID", -1, -1, &T4VOLTEmobileID},
	0x83: {"subRecordType", -1, 1, nil},
}

// T4BTLUnodesIdentifier
var T4BTLUnodesIdentifier = map[int]defStructs.Champ{
	0x00: {"SCSCFidentifier", -1, 1, nil},
	0x01: {"ASidentifier", -1, 1, nil},
}

// T4BTLUserviceSpecificInformation
var T4BTLUserviceSpecificInformation = map[int]defStructs.Champ{
	0x05: {"relatedCallID", -1, 1, nil},
	0x06: {"relatedCalllIDreason", -1, -1, nil},
	0x07: {"localCallID", -1, 1, nil},
	0x08: {"ASoriginalCalledNumber", -1, 1, nil},
	0x09: {"ASredirectingNumber", -1, 1, nil},
	0x0a: {"redirectingReason", -1, -1, nil},
	0x0b: {"originalCalledReason", -1, -1, nil},
}

// T4BTLUmobileID
var T4BTLUmobileID = map[int]defStructs.Champ{
	0x00: {"IMSI", -1, 1, nil},
	0x01: {"MSISDN", -1, 1, nil},
	0x02: {"ggsnAddr", -1, 1, nil},
	0x03: {"sgsnAddr", -1, 1, nil},
}

// T4BTLUcallingPartyAdress
var T4BTLUcallingPartyAdress = map[int]defStructs.Champ{
	0x01: {"callingPartyAddr02", -1, 1, nil},
}

// T3BTLU
var T3BTLU = map[int]defStructs.Champ{
	0x02: {"recordSequenceNumber", -1, -1, nil},
	0x04: {"callingPartyAdress", -1, -1, &T4BTLUcallingPartyAdress},
	0x05: {"requestedPartyAddress", -1, 1, nil},
	0x06: {"translatedCalledNumber", -1, 1, nil},
	0x0a: {"serviceRequestTimestamp", -1, 4, nil},
	0x0c: {"callStartTimestamp", -1, 4, nil},
	0x0d: {"callEndTimeStamp", -1, 4, nil},
	0x0e: {"callDuration", -1, -1, nil},
	0x11: {"roleOfTheNode", -1, -1, nil},
	0x12: {"callStatus", -1, -1, nil},
	0x13: {"sessionTerminationCause", -1, -1, nil},
	0x14: {"ASterminationCause", -1, -1, nil},
	0x15: {"IMSchargingIdentifier", -1, 1, nil},
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
	0x30: {"serviceSpecificInformation", -1, -1, &T4BTLUserviceSpecificInformation},
	0x50: {"codecUsed", -1, 1, nil},
	0x70: {"nodesIdentifier", -1, -1, &T4BTLUnodesIdentifier},
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
	0x7e: {"sigVolum", -1, -1, nil},
	0x7f: {"totalVolum", -1, -1, nil},
	0x81: {"extBearVolum", -1, -1, nil},
	0x82: {"mobileID", -1, -1, &T4BTLUmobileID},
	0x83: {"subRecordType", -1, 1, nil},
}

// T4BTICnodesIdentifier
var T4BTICnodesIdentifier = map[int]defStructs.Champ{
	0x00: {"SCSCFidentifier", -1, 1, nil},
	0x01: {"ASidentifier", -1, 1, nil},
}

// T4BTICsessionMediaComponents
var T4BTICsessionMediaComponents = map[int]defStructs.Champ{
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

// T4BTICserviceSpecificInformation
var T4BTICserviceSpecificInformation = map[int]defStructs.Champ{
	0x00: {"serviceData01", -1, 1, nil},
	0x02: {"serviceName", -1, 1, nil},
	0x03: {"resultCode", -1, 1, nil},
	0x04: {"invocationTime", -1, 3, nil},
	0x05: {"relatedCallID", -1, 1, nil},
	0x06: {"relatedCalllIDreason", -1, -1, nil},
	0x07: {"localCallID", -1, 1, nil},
	0x08: {"ASoriginalCalledNumber", -1, 1, nil},
	0x09: {"ASredirectingNumber", -1, 1, nil},
	0x0a: {"redirectingReason", -1, -1, nil},
	0x0b: {"originalCalledReason", -1, -1, nil},
}

// T4BTICcallingPartyAdress
var T4BTICcallingPartyAdress = map[int]defStructs.Champ{
	0x01: {"callingPartyAddr02", -1, 1, nil},
}

// T3BTIC
var T3BTIC = map[int]defStructs.Champ{
	0x03: {"AScallType", -1, -1, nil},
	0x04: {"callingPartyeAddress", -1, -1, &T4BTICcallingPartyAdress},
	0x05: {"requestedPartyAddress", -1, -1, nil},
	0x06: {"translatedCalledAddress", -1, -1, nil},
	0x0a: {"serviceTimestamp", -1, 3, nil},
	0x0c: {"callStartTimestamp", -1, 3, nil},
	0x0d: {"callEndTimestamp", -1, 3, nil},
	0x0e: {"callDuration", -1, -1, nil},
	0x11: {"roleOfNode", -1, -1, nil},
	0x13: {"sessionTerminationCause", -1, -1, nil},
	0x14: {"ASterminationCause", -1, -1, nil},
	0x15: {"IMSchargingIdentifier", -1, 1, nil},
	0x16: {"originatingIOI", -1, -1, nil},
	0x17: {"terminatingIOI", -1, -1, nil},
	0x19: {"accessSIPsession", -1, 1, nil},
	0x1a: {"networkSIPsession", -1, 1, nil},
	0x1b: {"accessNetworkInformation", -1, -1, nil},
	0x1c: {"causeCode", -1, -1, nil},
	0x1d: {"causeForRecordClosing", -1, -1, nil},
	0x1e: {"answerIndicator", -1, -1, nil},
	0x1f: {"sessionId", -1, 1, nil},
	0x30: {"serviceSpecificInformation", -1, -1, &T4BTICserviceSpecificInformation},
	0x50: {"codecUsed", -1, 1, nil},
	0x52: {"sessionMediaComponents", -1, -1, &T4BTICsessionMediaComponents},
	0x70: {"nodesIdentifier", -1, -1, &T4BTICnodesIdentifier},
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
	0x12: {"callStatus", -1, -1, nil},
	0xa8: {"serviceComputedInfo09", -1, 1, nil},
	0xa9: {"serviceComputedInfo10", -1, 1, nil},
	0xb1: {"serviceComputedInfo11", -1, 1, nil},
	0xb2: {"serviceComputedInfo12", -1, 1, nil},
	0xb3: {"serviceComputedInfo13", -1, 1, nil},
	0xb4: {"serviceComputedInfo14", -1, 1, nil},
	0xb5: {"serviceComputedInfo15", -1, 1, nil},
	0xb6: {"serviceComputedInfo16", -1, 1, nil},
}

// T4BIVnodesIdentifier
var T4BIVnodesIdentifier = map[int]defStructs.Champ{
	0x00: {"SCSCFidentifier", -1, 1, nil},
	0x01: {"ASidentifier", -1, 1, nil},
}

// T4BIVsessionMediaComponents
var T4BIVsessionMediaComponents = map[int]defStructs.Champ{
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

// T4BIVserviceSpecificInformation
var T4BIVserviceSpecificInformation = map[int]defStructs.Champ{
	0x05: {"relatedCallID", -1, 1, nil},
	0x06: {"relatedCallIDreason", -1, -1, nil},
	0x07: {"localCallID", -1, 1, nil},
	0x08: {"ASoriginalCalledNumber", -1, 1, nil},
	0x09: {"ASredirectingNumber", -1, 1, nil},
	0x0a: {"redirectingReason", -1, -1, nil},
	0x0b: {"originalCalledReason", -1, -1, nil},
}

// T4BIVcallingPartyAdress
var T4BIVcallingPartyAdress = map[int]defStructs.Champ{
	0x01: {"callingPartyAddr02", -1, 1, nil},
}

// T3BIV
var T3BIV = map[int]defStructs.Champ{
	0x04: {"callingPartyAdress", -1, -1, &T4BIVcallingPartyAdress},
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
	0x15: {"IMSchargingIdentifier", -1, 1, nil},
	0x16: {"originatingIOI", -1, -1, nil},
	0x17: {"terminatingIOI", -1, -1, nil},
	0x19: {"accessSIPsession", -1, 1, nil},
	0x1a: {"networkSIPsession", -1, 1, nil},
	0x1b: {"accessNetworkInformation", -1, -1, nil},
	0x1c: {"causeCode", -1, -1, nil},
	0x1d: {"causeForRecordClosing", -1, -1, nil},
	0x1e: {"answerIndicator", -1, -1, nil},
	0x1f: {"sessionId", -1, 1, nil},
	0x21: {"trunkGroupInfo", -1, 1, nil},
	0x30: {"serviceSpecificInformation", -1, -1, &T4BIVserviceSpecificInformation},
	0x50: {"codecUsed", -1, 1, nil},
	0x52: {"sessionMediaComponents", -1, -1, &T4BIVsessionMediaComponents},
	0x70: {"nodesIdentifier", -1, -1, &T4BIVnodesIdentifier},
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
}

// T2VOLTE
var T2VOLTE = map[int]defStructs.Champ{
	0x02: {"T2", -1, -1, &T3VOLTE},
}

// T2BTLU
var T2BTLU = map[int]defStructs.Champ{
	0x02: {"T2", -1, -1, &T3BTLU},
}

// T2BTIC
var T2BTIC = map[int]defStructs.Champ{
	0x02: {"T2", -1, -1, &T3BTIC},
}

// T2OPIT
var T2OPIT = map[int]defStructs.Champ{
	0x02: {"T2", -1, -1, &T3BIV},
}

// T2BIV
var T2BIV = map[int]defStructs.Champ{
	0x02: {"T2", -1, -1, &T3BIV},
}

// TVOLTE
var TVOLTE = map[int]defStructs.Champ{
	0x6a: {"T1VOLTE", -1, -1, &T2VOLTE},
}

// TBTLU
var TBTLU = map[int]defStructs.Champ{
	0x78: {"T1BTLU", -1, -1, &T2BTLU},
}

// TRCSE
var TRCSE = map[int]defStructs.Champ{
	0x79: {"T1RCSE", -1, -1, nil},
}

// TBTIC
var TBTIC = map[int]defStructs.Champ{
	0x57: {"T1BTIC", -1, -1, &T2BTIC},
}

// TOPIT
var TOPIT = map[int]defStructs.Champ{
	0x58: {"T1OPIT", -1, -1, &T2OPIT},
}

// TBIV
var TBIV = map[int]defStructs.Champ{
	0x77: {"T1BIV", -1, -1, &T2BIV},
}

// Dprims
var Dprims = map[int]defStructs.TabRechercheTypeDc{
	0: {"TBIV", 0xffff0000, 0x80770000, 0, &TBIV},
	1: {"TOPIT", 0xffff0000, 0x80580000, 0, &TOPIT},
	2: {"TBTIC", 0xffff0000, 0x80570000, 0, &TBTIC},
	3: {"TRCSE", 0xffff0000, 0x80790000, 0, &TRCSE},
	4: {"TBTLU", 0xffff0000, 0x80780000, 0, &TBTLU},
	5: {"TVOLTE", 0xffff0000, 0x806a0000, 0, &TVOLTE},
}
