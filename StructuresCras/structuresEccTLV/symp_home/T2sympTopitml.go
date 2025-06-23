package symp_home

import "local/structures/defStructs"

// T2sympTopitml
var T2sympTopitml = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3sympTopitml},
}

// T3sympTopitml
var T3sympTopitml = map[int]defStructs.Champ{
	0x04: {"CallingPartyAddress", -1, -1, &T4sympTopitmlCallingPartyAddress},
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
	0x30: {"ServiceSpecificInformation", -1, -1, &T4sympTopitmlServiceSpecificInformation},
	0x50: {"CodecUsed", -1, 1, nil},
	0x52: {"SessionMediaComponents", -1, -1, &T4sympTopitmlSessionMediaComponents},
	0x70: {"NodesID", -1, -1, &T4sympTopitmlNodesID},
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
	0x7e: {"ajoutPlatine", -1, -1, &T4sympAjoutPlatine},
}

// T4sympTopitmlCallingPartyAddress
var T4sympTopitmlCallingPartyAddress = map[int]defStructs.Champ{
	0x01: {"CallingPartyAddr02", -1, 1, nil},
}

// T4sympTopitmlServiceSpecificInformation
var T4sympTopitmlServiceSpecificInformation = map[int]defStructs.Champ{
	0x05: {"RelatedCallID", -1, 1, nil},
	0x06: {"RelatedCallIDreason", -1, 1, nil},
	0x07: {"LocalCallID", -1, 1, nil},
	0x08: {"ASoriginalCalledNumber", -1, 1, nil},
	0x09: {"ASredirectingNumber", -1, 1, nil},
	0x0a: {"RedirectingReason", -1, -1, nil},
	0x0b: {"OriginalCalledReason", -1, -1, nil},
}

// T4sympTopitmlSessionMediaComponents
var T4sympTopitmlSessionMediaComponents = map[int]defStructs.Champ{
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

// T4sympTopitmlNodesID
var T4sympTopitmlNodesID = map[int]defStructs.Champ{
	0x00: {"S-CSCF_ID", -1, 1, nil},
	0x01: {"ApplicationServerID", -1, 1, nil},
}
