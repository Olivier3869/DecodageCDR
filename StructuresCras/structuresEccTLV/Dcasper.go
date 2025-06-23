package structuresEccTLV

import "local/structures/defStructs"

// T4casperOpitmlCallingPartyAddress
var T4casperOpitmlCallingPartyAddress = map[int]defStructs.Champ{
	0x01: {"CallingPartyAddr02", -1, 1, nil},
}

// T4casperOpitmlServiceSpecificInformation
var T4casperOpitmlServiceSpecificInformation = map[int]defStructs.Champ{
	0x05: {"RelatedCallID", -1, 1, nil},
	0x06: {"RelatedCallIDreason", -1, 1, nil},
	0x07: {"LocalCallID", -1, 1, nil},
	0x08: {"ASoriginalCalledNumber", -1, 1, nil},
	0x09: {"ASredirectingNumber", -1, 1, nil},
	0x0a: {"RedirectingReason", -1, -1, nil},
	0x0b: {"OriginalCalledReason", -1, -1, nil},
}

// T4casperOpitmlSessionMediaComponents
var T4casperOpitmlSessionMediaComponents = map[int]defStructs.Champ{
	0x00: {"TimeMediaChange", -1, 3, nil},
	0x01: {"NameMedia", -1, 1, nil},
	0x02: {"DescriptionOfMedia", -1, 1, nil},
	0x03: {"DescriptionOfMedia", -1, 1, nil},
	0x04: {"DescriptionOfMedia", -1, 1, nil},
	0x05: {"DescriptionOfMedia", -1, 1, nil},
	0x06: {"DescriptionOfMedia", -1, 1, nil},
	0x07: {"DescriptionOfMedia", -1, 1, nil},
	0x08: {"DescriptionOfMedia", -1, 1, nil},
	0x09: {"DescriptionOfMedia", -1, 1, nil},
	0x0a: {"DescriptionOfMedia", -1, 1, nil},
	0x0b: {"DescriptionOfMedia", -1, 1, nil},
	0x0c: {"DurationPerMedia", -1, -1, nil},
}

// T4casperOpitmlNodesID
var T4casperOpitmlNodesID = map[int]defStructs.Champ{
	0x00: {"S-CSCF_ID", -1, 1, nil},
	0x01: {"ApplicationServerID", -1, 1, nil},
}

// T4casperV2oipASaSservice
var T4casperV2oipASaSservice = map[int]defStructs.Champ{
	0x01: {"aSserviceName", -1, -1, nil},
	0x02: {"aSfacResult", -1, -1, nil},
}

// T4casperAjoutPlatine
var T4casperAjoutPlatine = map[int]defStructs.Champ{
	0x00: {"CCD", -1, -1, nil},
	0x50: {"SGTQS", -1, 1, nil},
	0x51: {"CODCIR", -1, 1, nil},
	0x52: {"CODCIDR", -1, 1, nil},
	0x21: {"NSUC", -1, 1, nil},
}

// T3casperOpitml
var T3casperOpitml = map[int]defStructs.Champ{
	0x04: {"CallingPartyAddress", -1, -1, &T4casperOpitmlCallingPartyAddress},
	0x05: {"RequestedPartyAddress", -1, 1, nil},
	0x06: {"TranslatedCalledNumber", -1, 1, nil},
	0x0a: {"ServiceRequestTimestamp", -1, 3, nil},
	0x0c: {"CallStartTimestamp", -1, 3, nil},
	0x0d: {"CallEndTimeStamp", -1, 3, nil},
	0x0e: {"CallDuration", -1, -1, nil},
	0x11: {"RoleOfNode", -1, -1, nil},
	0x12: {"CallStatus", -1, -1, nil},
	0x13: {"SessionTerminationCause", -1, -1, nil},
	0x14: {"ASterminationCause", -1, -1, nil},
	0x15: {"IMSchargingIdentifier", -1, 1, nil},
	0x16: {"OriginatingIOI", -1, 1, nil},
	0x17: {"TerminatingIOI", -1, 1, nil},
	0x19: {"AccessSIPsessionIdentifier", -1, 1, nil},
	0x1a: {"NetworkSIPsessionIdentifier", -1, 1, nil},
	0x1b: {"AccessNetworkInformation", -1, 1, nil},
	0x1c: {"CauseCode", -1, -1, nil},
	0x1d: {"CauseForRecordClosing", -1, -1, nil},
	0x1e: {"AnswerIndicator", -1, -1, nil},
	0x1f: {"SessionID", -1, 1, nil},
	0x21: {"TrunkGroupInfo", -1, 1, nil},
	0x30: {"ServiceSpecificInformation", -1, -1, &T4casperOpitmlServiceSpecificInformation},
	0x50: {"CodecUsed", -1, 1, nil},
	0x52: {"SessionMediaComponents", -1, -1, &T4casperOpitmlSessionMediaComponents},
	0x70: {"NodesID", -1, -1, &T4casperOpitmlNodesID},
	0x71: {"ASreasingParty", -1, 1, nil},
	0x73: {"CompanyID", -1, 1, nil},
	0x74: {"group", -1, 1, nil},
	0x77: {"SIPmethod", -1, 1, nil},
	0x78: {"PrivateuserID", -1, 1, nil},
	0xa0: {"ServiceComputedInfo01", -1, 1, nil},
	0xa1: {"ServiceComputedInfo02", -1, 1, nil},
	0xa2: {"ServiceComputedInfo03", -1, 1, nil},
	0xa3: {"ServiceComputedInfo04", -1, 1, nil},
	0xa4: {"ServiceComputedInfo05", -1, 1, nil},
	0xa5: {"ServiceComputedInfo06", -1, 1, nil},
	0xa6: {"ServiceComputedInfo07", -1, 1, nil},
	0xa7: {"ServiceComputedInfo08", -1, 1, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4casperAjoutPlatine},
}

// T3casperSrnbi
var T3casperSrnbi = map[int]defStructs.Champ{
	0x03: {"codeEEC", -1, 1, nil},
	0x02: {"typeENR", -1, 1, nil},
	0x50: {"callId", -1, -1, nil},
	0x01: {"ICAR", -1, 7, nil},
	0x13: {"duree", -1, -1, nil},
	0x05: {"indRenv", -1, -1, nil},
	0x23: {"callType", -1, -1, nil},
	0x07: {"idusNAI", -1, -1, nil},
	0x08: {"idusNbPlan", -1, -1, nil},
	0x09: {"idusIndRedir", -1, -1, nil},
	0x40: {"idusNum", -1, -1, nil},
	0x0b: {"idusInf", -1, -1, nil},
	0x54: {"locIdusNAI", -1, -1, nil},
	0x55: {"locIdusNbPlan", -1, -1, nil},
	0x56: {"locIdusNum", -1, -1, nil},
	0x0c: {"dedeNAI", -1, -1, nil},
	0x0d: {"dedeNbPlan", -1, -1, nil},
	0x12: {"dedeNum", -1, -1, nil},
	0x0f: {"dedeInf", -1, -1, nil},
	0x10: {"destNAI", -1, -1, nil},
	0x11: {"destNbPlan", -1, -1, nil},
	0x43: {"DETRnum", -1, -1, nil},
	0x44: {"DETRinf", -1, -1, nil},
	0x14: {"callingPartyCateg", -1, -1, nil},
	0x3c: {"SUP", -1, -1, nil},
	0x16: {"reroutInf", -1, -1, nil},
	0x17: {"incomingPort", -1, -1, nil},
	0x18: {"outgoingPort", -1, -1, nil},
	0x1a: {"outgoingPOP", -1, -1, nil},
	0x1b: {"ougoingModule", -1, -1, nil},
	0x1c: {"CIC", -1, -1, nil},
	0x52: {"sourceIP", -1, 9, nil},
	0x53: {"destIP", -1, 9, nil},
	0x1f: {"GRFSEND", -1, -1, nil},
	0x1e: {"GRTSSORT", -1, -1, nil},
	0x21: {"anoCode", -1, -1, nil},
	0x22: {"anoInd", -1, -1, nil},
	0x25: {"locationEchec", -1, -1, nil},
	0x20: {"INEF", -1, -1, nil},
	0x60: {"status", -1, -1, nil},
	0x26: {"BidTime", -1, 1, nil},
	0x27: {"TCSdelay", -1, 1, nil},
	0x2a: {"DCSdelay", -1, -1, nil},
	0x28: {"CCdelay", -1, -1, nil},
	0x29: {"ICCdelay", -1, -1, nil},
	0x30: {"transportDelay", -1, -1, nil},
	0x31: {"interJitter", -1, -1, nil},
	0x5e: {"OUTPAC", -1, -1, nil},
	0x5c: {"OUTOCT", -1, -1, nil},
	0x5d: {"INPAC", -1, -1, nil},
	0x5b: {"INOCT", -1, -1, nil},
	0x36: {"lostPacket", -1, -1, nil},
	0x37: {"lostPacketOut", -1, -1, nil},
	0x38: {"packingPeriod", -1, -1, nil},
	0x39: {"coding", -1, -1, nil},
	0x3a: {"silenceSup", -1, -1, nil},
	0x3b: {"durInterPh", -1, -1, nil},
	0x3d: {"FSENT", -1, 1, nil},
	0x3e: {"FSSORT", -1, 1, nil},
	0x3f: {"delCliIdusNAI", -1, -1, nil},
	0x41: {"delCliIdusNbPlan", -1, -1, nil},
	0x42: {"delCliIdusPubPriv", -1, -1, nil},
	0x32: {"delCliIdusNum", -1, -1, nil},
	0x45: {"OCNDEDENAI", -1, -1, nil},
	0x46: {"OCNDEDEnbPlan", -1, -1, nil},
	0x33: {"OCNDEDEnum", -1, -1, nil},
	0x48: {"RDNinfRedirNAI", -1, -1, nil},
	0x49: {"RDNinfRedirNbPlan", -1, -1, nil},
	0x34: {"RDNinfRedirNum", -1, -1, nil},
	0x35: {"callEchoDevice", -1, -1, nil},
	0x4c: {"FSENT2", -1, -1, nil},
	0x4d: {"FSSORT2", -1, -1, nil},
	0x4e: {"netwkCallId", -1, -1, nil},
	0x4f: {"nbPulses", -1, -1, nil},
	0x75: {"chargingFctAddr", -1, 1, nil},
	0x76: {"SIPstatus", -1, -1, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4casperAjoutPlatine},
}

// T3casperDevil
var T3casperDevil = map[int]defStructs.Champ{
	0x03: {"CodeEEC", -1, 1, nil},
	0x02: {"TYPENR", -1, 1, nil},
	0x50: {"CallId", -1, -1, nil},
	0x01: {"ICAR", -1, 7, nil},
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
	0x52: {"SourceIP", -1, 9, nil},
	0x53: {"DestIP", -1, 9, nil},
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
	0x7e: {"ajoutPlatine", -1, -1, &T4casperAjoutPlatine},
}

// T3casperV2oipAS
var T3casperV2oipAS = map[int]defStructs.Champ{
	0x03: {"networkInfo", -1, 1, nil},
	0x40: {"callingNb", -1, 1, nil},
	0x43: {"calledNb", -1, 1, nil},
	0x12: {"translatedNb", -1, 1, nil},
	0x04: {"origCalledNb", -1, 1, nil},
	0x05: {"redirectingNb", -1, 1, nil},
	0x06: {"redirectingReason", -1, 1, nil},
	0x07: {"aScodec", -1, -1, nil},
	0x01: {"date", -1, 8, nil},
	0x08: {"sessionEndTime", -1, 8, nil},
	0x13: {"duree", -1, -1, nil},
	0x09: {"roleOfNode", -1, -1, nil},
	0x20: {"efficiency", -1, -1, nil},
	0x0a: {"termation", -1, -1, nil},
	0x0b: {"chargingId", -1, 1, nil},
	0x0c: {"aSreleasingParty", -1, -1, nil},
	0x11: {"cSaSId", -1, 1, nil},
	0x0d: {"chargeIndicator", -1, 1, nil},
	0x1a: {"localCallId", -1, 1, nil},
	0x1b: {"relatedCallId", -1, 1, nil},
	0x1c: {"relatedCallIdReason", -1, 1, nil},
	0x1d: {"answerIndicator", -1, 1, nil},
	0x1f: {"Q850cause", -1, -1, nil},
	0x30: {"aSservice1", -1, -1, &T4casperV2oipASaSservice},
	0x31: {"aSservice2", -1, -1, &T4casperV2oipASaSservice},
	0x32: {"aSservice3", -1, -1, &T4casperV2oipASaSservice},
	0x33: {"aSservice4", -1, -1, &T4casperV2oipASaSservice},
	0x34: {"aSservice5", -1, -1, &T4casperV2oipASaSservice},
	0x16: {"accessCallId", -1, -1, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4casperAjoutPlatine},
}

// T2casperOpitml
var T2casperOpitml = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3casperOpitml},
}

// T2casperSrnbi
var T2casperSrnbi = map[int]defStructs.Champ{
	0x02: {"devil", -1, -1, &T3casperDevil},
	0x03: {"sRNBIfacturation", -1, -1, &T3casperSrnbi},
}

// T2casperV2oipAS
var T2casperV2oipAS = map[int]defStructs.Champ{
	0x02: {"facturation", -1, -1, &T3casperV2oipAS},
}

// Tcasper
var Tcasper = map[int]defStructs.Champ{
	0x58: {"Topitml", -1, -1, &T2casperOpitml},
	0x48: {"CallServer", -1, -1, &T2casperSrnbi},
	0x33: {"V2oipAS", -1, -1, &T2casperV2oipAS},
}

// TcasperTta6G3
var TcasperTta6G3 = map[int]defStructs.Champ{
	0: {"sty", 1, -1, nil},
	1: {"sct", 1, -1, nil},
	2: {"icar", 5, 7, nil},
	3: {"inef", -1, -1, nil},
}

// TcasperTta6G4
var TcasperTta6G4 = map[int]defStructs.Champ{
	0: {"sty", 1, -1, nil},
	1: {"sct", 1, -1, nil},
	2: {"icar", 5, 7, nil},
	3: {"inef", 1, -1, nil},
	4: {"ref", -1, -1, nil},
}

// TcasperTta6IDUS
var TcasperTta6IDUS = map[int]defStructs.Champ{
	0: {"inf", 1, -1, nil},
	1: {"num", -1, -1, nil},
}

// TcasperTta6DEDE
var TcasperTta6DEDE = map[int]defStructs.Champ{
	0: {"inf", 1, -1, nil},
	1: {"num", -1, -1, nil},
}

// TcasperTta6IDAV
var TcasperTta6IDAV = map[int]defStructs.Champ{
	0: {"inf", 1, -1, nil},
	1: {"num", -1, -1, nil},
}

// TcasperTta6NOS
var TcasperTta6NOS = map[int]defStructs.Champ{
	0: {"inf", 1, -1, nil},
	1: {"num", -1, -1, nil},
}

// TcasperTta6CAP1
var TcasperTta6CAP1 = map[int]defStructs.Champ{
	0: {"cax", 1, -1, nil},
	1: {"tap", 1, -1, nil},
	2: {"dur", 2, -1, nil},
	3: {"tax", -1, -1, nil},
}

// TcasperTta6LOCUS
var TcasperTta6LOCUS = map[int]defStructs.Champ{
	0: {"inf", 1, -1, nil},
	1: {"num", -1, -1, nil},
}

// TcasperTta6IPRV
var TcasperTta6IPRV = map[int]defStructs.Champ{
	0: {"inf", 1, -1, nil},
	1: {"num", -1, -1, nil},
}

// TcasperTta6CAP2
var TcasperTta6CAP2 = map[int]defStructs.Champ{
	0: {"cax", 1, -1, nil},
	1: {"tap", 1, -1, nil},
	2: {"dur", 2, -1, nil},
	3: {"tax", 2, -1, nil},
	4: {"sup", -1, -1, nil},
}

// TcasperTta6ICOM
var TcasperTta6ICOM = map[int]defStructs.Champ{
	0: {"inf", 1, -1, nil},
	1: {"num", -1, -1, nil},
}

// TcasperTta6IDRV
var TcasperTta6IDRV = map[int]defStructs.Champ{
	0: {"inf", 1, -1, nil},
	1: {"num", -1, -1, nil},
}

// TcasperTta6ICX
var TcasperTta6ICX = map[int]defStructs.Champ{
	0x01: {"GRPX", -1, -1, nil},
	0x02: {"NETX", -1, -1, nil},
	0x03: {"DSTX", -1, -1, nil},
	0x04: {"NUPX", -1, -1, nil},
}

// T3casperTta6dc
var T3casperTta6dc = map[int]defStructs.Champ{
	0x03: {"G3", -1, -1, &TcasperTta6G3},
	0x1a: {"LINK", -1, -1, nil},
	0x1d: {"CALC", -1, -1, nil},
	0x40: {"IDUS", -1, -1, &TcasperTta6IDUS},
	0x43: {"DEDE", -1, -1, &TcasperTta6DEDE},
	0x45: {"IDAV", -1, -1, &TcasperTta6IDAV},
	0x44: {"COCA", -1, -1, nil},
	0x46: {"COCAV", -1, -1, nil},
	0x42: {"NOS", -1, -1, &TcasperTta6NOS},
	0x38: {"INDEX", -1, -1, nil},
	0x08: {"CAP1", -1, -1, &TcasperTta6CAP1},
	0x06: {"CAB", -1, -1, nil},
	0x3c: {"SIDR", -1, -1, nil},
	0x3b: {"CODE", -1, -1, nil},
	0x3f: {"INV", -1, -1, nil},
	0x3d: {"CLE", -1, -1, nil},
	0x18: {"CLER", -1, -1, nil},
	0x1b: {"SEL", -1, -1, nil},
	0x1c: {"SON", -1, -1, nil},
	0x20: {"reserve20", -1, 2, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4casperAjoutPlatine},
}

// T3casperTta6dea
var T3casperTta6dea = map[int]defStructs.Champ{
	0x04: {"G4", -1, -1, &TcasperTta6G4},
	0x1a: {"LINK", -1, -1, nil},
	0x1d: {"CALC", -1, -1, nil},
	0x40: {"IDUS", -1, -1, &TcasperTta6IDUS},
	0x42: {"NOS", -1, -1, &TcasperTta6NOS},
	0x47: {"LOCUS", -1, -1, &TcasperTta6LOCUS},
	0x43: {"DEDE", -1, -1, &TcasperTta6DEDE},
	0x45: {"IDAV", -1, -1, &TcasperTta6IDAV},
	0x44: {"COCA", -1, -1, nil},
	0x46: {"COCAV", -1, -1, nil},
	0x30: {"IPRV", -1, -1, &TcasperTta6IPRV},
	0x31: {"IDRV", -1, -1, nil},
	0x19: {"NBRV", -1, -1, nil},
	0x32: {"DUR", -1, -1, nil},
	0x07: {"UT", -1, -1, nil},
	0x3d: {"CLE", -1, -1, nil},
	0x18: {"CLER", -1, -1, nil},
	0x1b: {"SEL", -1, -1, nil},
	0x1c: {"SON", -1, -1, nil},
	0x1f: {"reserve1f", -1, 2, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4casperAjoutPlatine},
}

// T3casperTns6dc
var T3casperTns6dc = map[int]defStructs.Champ{
	0x04: {"G4", -1, -1, &TcasperTta6G4},
	0x1a: {"LINK", -1, -1, nil},
	0x1d: {"CALC", -1, -1, nil},
	0x40: {"IDUS", -1, -1, &TcasperTta6IDUS},
	0x43: {"DEDE", -1, -1, &TcasperTta6DEDE},
	0x45: {"IDAV", -1, -1, &TcasperTta6IDAV},
	0x44: {"COCA", -1, -1, nil},
	0x46: {"COCAV", -1, -1, nil},
	0x42: {"NOS", -1, -1, &TcasperTta6NOS},
	0x38: {"INDEX", -1, -1, nil},
	0x0a: {"CAP2", -1, -1, &TcasperTta6CAP1},
	0x06: {"CAB", -1, -1, nil},
	0x3c: {"SIDR", -1, -1, nil},
	0x3b: {"CODE", -1, -1, nil},
	0x3f: {"INV", -1, -1, nil},
	0x3d: {"CLE", -1, -1, nil},
	0x18: {"CLER", -1, -1, nil},
	0x3e: {"TTX", -1, -1, nil},
	0x0c: {"CHA", -1, -1, nil},
	0x0d: {"CHP", -1, -1, nil},
	0x1b: {"SEL", -1, -1, nil},
	0x1c: {"SON", -1, -1, nil},
	0x20: {"reserve20", -1, 2, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4casperAjoutPlatine},
}

// T3casperTns6dea
var T3casperTns6dea = map[int]defStructs.Champ{
	0x04: {"G4", -1, -1, &TcasperTta6G4},
	0x1a: {"LINK", -1, -1, nil},
	0x1d: {"CALC", -1, -1, nil},
	0x40: {"IDUS", -1, -1, &TcasperTta6IDUS},
	0x42: {"NOS", -1, -1, &TcasperTta6NOS},
	0x47: {"LOCUS", -1, -1, &TcasperTta6LOCUS},
	0x43: {"DEDE", -1, -1, &TcasperTta6DEDE},
	0x45: {"IDAV", -1, -1, &TcasperTta6IDAV},
	0x44: {"COCA", -1, -1, nil},
	0x46: {"COCAV", -1, -1, nil},
	0x30: {"IPRV", -1, -1, &TcasperTta6IPRV},
	0x31: {"IDRV", -1, -1, nil},
	0x19: {"NBRV", -1, -1, nil},
	0x32: {"DUR", -1, -1, nil},
	0x07: {"UT", -1, -1, nil},
	0x3d: {"CLE", -1, -1, nil},
	0x18: {"CLER", -1, -1, nil},
	0x1b: {"SEL", -1, -1, nil},
	0x1c: {"SON", -1, -1, nil},
	0x1f: {"reserve1f", -1, 2, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4casperAjoutPlatine},
}

// T3casperTrm6dc
var T3casperTrm6dc = map[int]defStructs.Champ{
	0x04: {"G4", -1, -1, &TcasperTta6G4},
	0x1a: {"LINK", -1, -1, nil},
	0x1d: {"CALC", -1, -1, nil},
	0x40: {"IDUS", -1, -1, &TcasperTta6IDUS},
	0x43: {"DEDE", -1, -1, &TcasperTta6DEDE},
	0x45: {"IDAV", -1, -1, &TcasperTta6IDAV},
	0x44: {"COCA", -1, -1, nil},
	0x46: {"COCAV", -1, -1, nil},
	0x0a: {"CAP2", -1, -1, &TcasperTta6CAP2},
	0x06: {"CAB", -1, -1, nil},
	0x3d: {"CLE", -1, -1, nil},
	0x18: {"CLER", -1, -1, nil},
	0x0c: {"CHA", -1, -1, nil},
	0x1b: {"SEL", -1, -1, nil},
	0x1c: {"SON", -1, -1, nil},
	0x47: {"LOCUS", -1, -1, &TcasperTta6LOCUS},
	0x41: {"ICOM", -1, -1, &TcasperTta6ICOM},
	0x30: {"IPRV", -1, -1, &TcasperTta6IPRV},
	0x31: {"IDRV", -1, -1, &TcasperTta6IDRV},
	0x19: {"NBRV", -1, -1, nil},
	0x48: {"IDSV", -1, -1, nil},
	0x1f: {"reserve1f", -1, 2, nil},
	0x20: {"reserve20", -1, 2, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4casperAjoutPlatine},
}

// T3casperTfe6dc
var T3casperTfe6dc = map[int]defStructs.Champ{
	0x04: {"G4", -1, -1, &TcasperTta6G4},
	0x1a: {"LINK", -1, -1, nil},
	0x1d: {"CALC", -1, -1, nil},
	0x40: {"IDUS", -1, -1, &TcasperTta6IDUS},
	0x43: {"DEDE", -1, -1, &TcasperTta6DEDE},
	0x45: {"IDAV", -1, -1, &TcasperTta6IDAV},
	0x44: {"COCA", -1, -1, nil},
	0x46: {"COCAV", -1, -1, nil},
	0x0a: {"CAP2", -1, -1, &TcasperTta6CAP2},
	0x06: {"CAB", -1, -1, nil},
	0x3d: {"CLE", -1, -1, nil},
	0x18: {"CLER", -1, -1, nil},
	0x0c: {"CHA", -1, -1, nil},
	0x1b: {"SEL", -1, -1, nil},
	0x1c: {"SON", -1, -1, nil},
	0x47: {"LOCUS", -1, -1, &TcasperTta6LOCUS},
	0x41: {"ICOM", -1, -1, &TcasperTta6ICOM},
	0x30: {"IPRV", -1, -1, &TcasperTta6IPRV},
	0x31: {"IDRV", -1, -1, &TcasperTta6IDRV},
	0x19: {"NBRV", -1, -1, nil},
	0x48: {"IDSV", -1, -1, nil},
	0x1f: {"reserve1f", -1, 2, nil},
	0x20: {"reserve20", -1, 2, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4casperAjoutPlatine},
}

// T3casperTcentrexDc
var T3casperTcentrexDc = map[int]defStructs.Champ{
	0x04: {"G4", -1, -1, &TcasperTta6G4},
	0x1a: {"LINK", -1, -1, nil},
	0x1d: {"CALC", -1, -1, nil},
	0x40: {"IDUS", -1, -1, &TcasperTta6IDUS},
	0x43: {"DEDE", -1, -1, &TcasperTta6DEDE},
	0x45: {"IDAV", -1, -1, &TcasperTta6IDAV},
	0x44: {"COCA", -1, -1, nil},
	0x46: {"COCAV", -1, -1, nil},
	0x42: {"NOS", -1, -1, &TcasperTta6NOS},
	0x38: {"INDEX", -1, -1, nil},
	0x0a: {"CAP2", -1, -1, &TcasperTta6CAP1},
	0x06: {"CAB", -1, -1, nil},
	0x3c: {"SIDR", -1, -1, nil},
	0x3b: {"CODE", -1, -1, nil},
	0x3f: {"INV", -1, -1, nil},
	0x3d: {"CLE", -1, -1, nil},
	0x18: {"CLER", -1, -1, nil},
	0x3e: {"TTX", -1, -1, nil},
	0x0c: {"CHA", -1, -1, nil},
	0x0d: {"CHP", -1, -1, nil},
	0x1b: {"SEL", -1, -1, nil},
	0x1c: {"SON", -1, -1, nil},
	0x39: {"ICX", -1, -1, &TcasperTta6ICX},
	0x20: {"reserve20", -1, 2, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4casperAjoutPlatine},
}

// T2casperTta6
var T2casperTta6 = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3casperTta6dc},
	0x18: {"dea", -1, -1, &T3casperTta6dea},
}

// T2casperTns6
var T2casperTns6 = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3casperTns6dc},
	0x18: {"dea", -1, -1, &T3casperTns6dea},
}

// T2casperTrm6
var T2casperTrm6 = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3casperTrm6dc},
}

// T2casperTfe6
var T2casperTfe6 = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3casperTfe6dc},
}

// T2casperTcentrex
var T2casperTcentrex = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3casperTcentrexDc},
}

// TcasperRTC
var TcasperRTC = map[int]defStructs.Champ{
	0x00: {"Tta6", -1, -1, &T2casperTta6},
	0x02: {"Tns6", -1, -1, &T2casperTns6},
	0x05: {"Trm6", -1, -1, &T2casperTrm6},
	0x0e: {"Tfe6", -1, -1, &T2casperTfe6},
	0x19: {"Tcentrex", -1, -1, &T2casperTcentrex},
}

// Dcasper
var Dcasper = map[int]defStructs.TabRechercheTypeDc{
	0:  {"TcasperOpitml", 0xffff0000, 0x80580000, 0, &Tcasper},
	1:  {"TcasperOpitmldc", 0xffff0000, 0x80020000, 0, &Tcasper},
	2:  {"TcasperCs", 0xffff0000, 0x80480000, 0, &Tcasper},
	3:  {"TcasperV2oip", 0xffff0000, 0x80020000, 0, &Tcasper},
	4:  {"TcasperSrnbidc", 0xffff0000, 0x80030000, 0, &Tcasper},
	5:  {"TcasperV2oipAS", 0xffff0000, 0x80330000, 0, &Tcasper},
	6:  {"TcasperV2oipAS", 0xffff0000, 0x80020000, 0, &Tcasper},
	7:  {"TcasperTta6dc", 0xff00ff00, 0x00000200, 0, &TcasperRTC},
	8:  {"TcasperTta6dea", 0xff00ff00, 0x00001800, 0, &TcasperRTC},
	9:  {"TcasperTns6dc", 0xff00ff00, 0x02000200, 0, &TcasperRTC},
	10: {"TcasperTns6dea", 0xff00ff00, 0x02001800, 0, &TcasperRTC},
	11: {"TcasperTrm6dc", 0xff00ff00, 0x05000200, 0, &TcasperRTC},
	12: {"TcasperTch6dc", 0xff00ff00, 0x0e000200, 0, &TcasperRTC},
	13: {"TcasperTcentrexdc", 0xff00ff00, 0x19000200, 0, &TcasperRTC},
}
