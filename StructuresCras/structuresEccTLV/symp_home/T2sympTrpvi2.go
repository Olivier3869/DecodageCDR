package symp_home

import "local/structures/defStructs"

// T2sympTrpvi2
var T2sympTrpvi2 = map[int]defStructs.Champ{
	0x02: {"dc", -1, -1, &T3sympTrpvi2dc},
}

// T3sympTrpvi2dc
var T3sympTrpvi2dc = map[int]defStructs.Champ{
	0x03:  {"Date", -1, 5, nil},
	0x163: {"CallingNumber", -1, -1, &T3sympTrpvi2dcCallingNumber},
	0x164: {"CalledNumber", -1, -1, &T3sympTrpvi2dcCalledNumber},
	0x165: {"TranslatedNumber", -1, -1, &T3sympTrpvi2dcTranslatedNumber},
	0x167: {"DateDuration", -1, -1, &T3sympTrpvi2dcDateDuration},
	0x192: {"Detail", -1, -1, &T4sympTrpvi2dcDetail},
	0x7e:  {"ajoutPlatine", -1, -1, &T4sympAjoutPlatine},
}

// T3sympTrpvi2dcCallingNumber
var T3sympTrpvi2dcCallingNumber = map[int]defStructs.Champ{
	0: {"NumberingFormat", 1, -1, nil},
	1: {"CallingNumber", -1, -1, nil},
}

// T3sympTrpvi2dcCalledNumber
var T3sympTrpvi2dcCalledNumber = map[int]defStructs.Champ{
	0: {"NumberingFormat", 1, -1, nil},
	1: {"CalledNumber", -1, -1, nil},
}

// T3sympTrpvi2dcTranslatedNumber
var T3sympTrpvi2dcTranslatedNumber = map[int]defStructs.Champ{
	0: {"NumberingFormat", 1, -1, nil},
	1: {"CalledNumber", -1, -1, nil},
}

// T3sympTrpvi2dcDateDuration
var T3sympTrpvi2dcDateDuration = map[int]defStructs.Champ{
	0: {"Date", 6, 5, nil},
	1: {"Duration", -1, -1, nil},
}

// T4sympTrpvi2dcDetailCallType
var T4sympTrpvi2dcDetailCallType = map[int]defStructs.Champ{
	0: {"Calltype", 1, -1, nil},
	1: {"TeleserviceIndicator", 1, -1, nil},
	2: {"NetworkCause", 1, -1, nil},
	3: {"ServiceIndicator", -1, -1, nil},
}

// T4sympTrpvi2dcDetailOriginalCalledNumber
var T4sympTrpvi2dcDetailOriginalCalledNumber = map[int]defStructs.Champ{
	0: {"NumberingFormat", 1, -1, nil},
	1: {"OriginalCalledNumber", -1, -1, nil},
}

// T4sympTrpvi2dcDetailPrivateNetworkInformation1
var T4sympTrpvi2dcDetailPrivateNetworkInformation1 = map[int]defStructs.Champ{
	0: {"CharacteristicOfCall", 1, -1, nil},
	1: {"GroupNumber", -1, -1, nil},
}

// T4sympTrpvi2dcDetailPrivateNetworkInformation2
var T4sympTrpvi2dcDetailPrivateNetworkInformation2 = map[int]defStructs.Champ{
	0: {"LineNumber", 1, -1, nil},
	1: {"AccountCodeNumber", -1, -1, nil},
}

// T4sympTrpvi2dcDetailAddCallingNumber
var T4sympTrpvi2dcDetailAddCallingNumber = map[int]defStructs.Champ{
	0: {"NumberingFormat", 1, -1, nil},
	1: {"AddCallingNumber", -1, -1, nil},
}

// T4sympTrpvi2dcDetailCMMuserIdentity
var T4sympTrpvi2dcDetailCMMuserIdentity = map[int]defStructs.Champ{
	0: {"NumberingFormat", 1, -1, nil},
	1: {"CMMuserIdentity", -1, -1, nil},
}

// T4sympTrpvi2dcDetail
var T4sympTrpvi2dcDetail = map[int]defStructs.Champ{
	0x16: {"CallType", -1, -1, &T4sympTrpvi2dcDetailCallType},
	0x2a: {"TypeofRerouting", -1, -1, nil},
	0x30: {"PrivateCompanyInformation", -1, -1, nil},
	0x2e: {"OriginalCalledNumber", -1, -1, &T4sympTrpvi2dcDetailOriginalCalledNumber},
	0x38: {"PrivateNetworkInformation1", -1, -1, &T4sympTrpvi2dcDetailPrivateNetworkInformation1},
	0x39: {"SiteIdentifier", -1, -1, nil},
	0x3a: {"InternalNumber", -1, -1, nil},
	0x3b: {"PrivateNetworkInformation2", -1, -1, &T4sympTrpvi2dcDetailPrivateNetworkInformation2},
	0x44: {"AddCallingNumber", -1, -1, &T4sympTrpvi2dcDetailAddCallingNumber},
	0x45: {"CMMuserIdentity", -1, -1, &T4sympTrpvi2dcDetailCMMuserIdentity},
}
