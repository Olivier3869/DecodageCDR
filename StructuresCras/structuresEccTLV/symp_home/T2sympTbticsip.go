package symp_home

import "local/structures/defStructs"

// T2symphTbticsip
var T2symphTbticsip = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3symphTbticsipdc},
}

// T3symphTbticsipdc
var T3symphTbticsipdc = map[int]defStructs.Champ{
	0x03: {"ASCallType", -1, -1, nil},
	0x04: {"CallingPartyAdress", -1, -1, &T4symphTbticsipdcCallingPartyAdress},
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
	0x1f: {"SessionId", -1, 1, nil},
	0x30: {"ServiceSpecificInformation", -1, -1, &T4symphTbticsipdcServiceSpecificInformation},
	0x50: {"CodecUsed", -1, -1, nil},
	0x52: {"SessionMediaComponents", -1, -1, &T4symphTbticsipdcSessionMediaComponents},
	0x70: {"NodesIdentifier", -1, -1, &T4symphTbticsipdcNodesIdentifier},
	0x71: {"ASreleasingParty", -1, -1, nil},
	0x73: {"CompanyIdentifier", -1, 1, nil},
	0x74: {"Group", -1, -1, nil},
	0x77: {"SIPmethod", -1, -1, nil},
	0x78: {"PrivateUserID", -1, -1, nil},
	0xa0: {"ServiceComputedInfo01", -1, 1, nil},
	0xa1: {"ServiceComputedInfo02", -1, 1, nil},
	0xa2: {"ServiceComputedInfo03", -1, 1, nil},
	0xa3: {"ServiceComputedInfo04", -1, 1, nil},
	0xa4: {"ServiceComputedInfo05", -1, 1, nil},
	0xa5: {"ServiceComputedInfo06", -1, 1, nil},
	0xa6: {"ServiceComputedInfo07", -1, 1, nil},
	0xa7: {"ServiceComputedInfo08", -1, 1, nil},
	0xa8: {"ServiceComputedInfo09", -1, 1, nil},
	0xa9: {"ServiceComputedInfo10", -1, 1, nil},
	0xb0: {"ServiceComputedInfo11", -1, 1, nil},
	0xb1: {"ServiceComputedInfo12", -1, 1, nil},
	0xb2: {"ServiceComputedInfo13", -1, 1, nil},
	0xb3: {"ServiceComputedInfo14", -1, 1, nil},
	0xb4: {"ServiceComputedInfo15", -1, 1, nil},
	0xb5: {"ServiceComputedInfo16", -1, 1, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4sympAjoutPlatine},
}

// T4symphTbticsipdcCallingPartyAdress
var T4symphTbticsipdcCallingPartyAdress = map[int]defStructs.Champ{
	0x01: {"CallingPartyAddr02", -1, 1, nil},
}

// T4symphTbticsipdcServiceSpecificInformation
var T4symphTbticsipdcServiceSpecificInformation = map[int]defStructs.Champ{
	0x00: {"servicedata01", -1, 1, nil},
	0x02: {"servicename", -1, 1, nil},
	0x03: {"ResultCode", -1, 1, nil},
	0x04: {"InvocationTime", -1, -1, nil},
	0x05: {"RelatedCallID", -1, 1, nil},
	0x06: {"RelatedCallIDreason", -1, -1, nil},
	0x07: {"LocalCallID", -1, 1, nil},
	0x08: {"ASoriginalCalledNumber", -1, 1, nil},
	0x09: {"ASredirectingNumber", -1, 1, nil},
	0x0a: {"RedirectingReason", -1, -1, nil},
	0x0b: {"OriginalCalleReason", -1, -1, nil},
}

// T4symphTbticsipdcSessionMediaComponents
var T4symphTbticsipdcSessionMediaComponents = map[int]defStructs.Champ{
	0x00: {"TimeOfMediaChange", -1, -1, nil},
	0x01: {"NameOfMedia", -1, 1, nil},
	0x02: {"DescriptionOfMedia01", -1, 1, nil},
	0x03: {"DescriptionOfMedia02", -1, 1, nil},
	0x04: {"DescriptionOfMedia03", -1, 1, nil},
	0x05: {"DescriptionOfMedia04", -1, 1, nil},
	0x06: {"DescriptionOfMedia05", -1, 1, nil},
	0x07: {"DescriptionOfMedia06", -1, 1, nil},
	0x08: {"DescriptionOfMedia07", -1, 1, nil},
	0x09: {"DescriptionOfMedia08", -1, 1, nil},
	0x0a: {"DescriptionOfMedia09", -1, 1, nil},
	0x0b: {"DescriptionOfMedia10", -1, 1, nil},
	0x0c: {"DurationPerMedia", -1, -1, nil},
}

// T4symphTbticsipdcNodesIdentifier
var T4symphTbticsipdcNodesIdentifier = map[int]defStructs.Champ{
	0x00: {"S-CSCFidentifier", -1, 1, nil},
	0x01: {"ASidentifier", -1, 1, nil},
}
