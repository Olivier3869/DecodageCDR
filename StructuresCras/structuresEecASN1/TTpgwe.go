package structuresEecASN1

import (
	"local/structures/defStructs"

	"local/structures/structuresEecASN1/spgw"
)

// TTpgwe
var TTpgwe = map[int]defStructs.TabRechercheTypeDc{
	0: {"Ttpgwe", 0x00ff0000, 0x004f0000, 0, &Ttpgwe},
}

var Ttpgwe = map[int]defStructs.Champ{
	0x4f: {"TTpgwe", -1, -1, &spgw.CallER},
}

//var PgweCallER = map[int]defStructs.Champ{
//	0:  {"RecordType", -1, 0, 1, nil},
//	3:  {"ServedIMSI", -1, 8, 1, nil},
//	4:  {"PGWAddress", -1, 0, 1, &pGbinAddress},
//	5:  {"chargingID", -1, 7, 1, nil},
//	6:  {"servingNodeAddress", -1, 0, 1, &serviIP},
//	7:  {"accessPointNameNI", -1, 2, 1, nil},
//	8:  {"pdpPDNType", -1, 0, 1, nil},
//	9:  {"servedPDPPDNAddress", -1, 0, 1, &servedIP},
//	11: {"dynamicAddressFlag", -1, 0, 1, nil},
//	13: {"recordOpeningTime", -1, 6, 1, nil},
//	14: {"duration", -1, 7, 1, nil},
//	15: {"causeForRecClosing", -1, 0, 1, nil},
//	16: {"diagnostics", -1, 0, 1, &cause},
//	17: {"recordSequenceNumber", -1, 7, 1, nil},
//	18: {"nodeID", -1, 2, 1, nil},
//	19: {"recordExtensions", -1, 0, 1, &ManagementExt},
//	20: {"localSequenceNumber", -1, 0, 1, nil},
//	21: {"apnSelectionMode", -1, 0, 1, nil},
//	22: {"servedMSISDN", -1, 8, 1, nil},
//	23: {"chargingCharacteristics", -1, 0, 1, nil},
//	24: {"chChSelectionMode", -1, 0, 1, nil},
//	27: {"servingNodePLMNIdentifier", -1, 8, 1, nil},
//	29: {"servedIMEISV", -1, 8, 1, nil},
//	30: {"rATType", -1, 7, 1, nil},
//	31: {"mSTimeZone", -1, 8, 1, nil},
//	32: {"userLocationInformation", -1, 0, 1, nil},
//	34: {"listOfServiceData", -1, 0, 1, &spgw.ListService},
//	35: {"servingNodeType", -1, 0, 1, &spgw.Serving},
//	36: {"servedMNNAI", -1, 0, 1, &spgw.SubscriptionID},
//	37: {"PGWPLMNIdentifier", -1, 8, 1, nil},
//	38: {"startTime", -1, 6, 1, nil},
//	39: {"stopTime", -1, 6, 1, nil},
//	41: {"pDNConnectionID", -1, 0, 1, nil},
//}

//var pGbinAddress = map[int]defStructs.Champ{
//	0: {"PGWV4address", -1, 9, 2, nil},
//	1: {"PGWV6address", -1, 9, 2, nil},
//}

//var servedIP = map[int]defStructs.Champ{
//	0: {"iPAdress", -1, 0, 2, &serviIP},
//}

//var serviIP = map[int]defStructs.Champ{
//	0: {"servediPBinV4Address", -1, 9, 2, nil},
//	1: {"servediPBinV6Address", -1, 9, 2, nil},
//}

//var ManagementExt = map[int]defStructs.Champ{
//	16: {"ManagementExten", -1, 0, 2, &Manage},
//}

//var Manage = map[int]defStructs.Champ{
//	6: {"Identifier", -1, 0, 3, nil},
//	2: {"serviceLevelCDR", -1, 0, 3, &spgw.ServiceLevel},
//}

//var ServiceLevel = map[int]defStructs.Champ{
//	1: {"sessionID", -1, 2, 4, nil},
//	2: {"serviceID", -1, 0, 4, nil},
//	3: {"serviceIPVolumeUsage", -1, 0, 4, &serviceIPVolume},
//	4: {"quotaServerFlags", -1, 0, 4, nil},
//	5: {"qualifiedUsage", -1, 0, 4, &qualisage},
//	6: {"timeOfFirstUsage", -1, 6, 4, nil},
//}

//var serviceIPVolume = map[int]defStructs.Champ{
//	1: {"cumIPBytesUploaded", -1, 0, 4, nil},
//	2: {"cumIPBytesdwnloaded", -1, 0, 4, nil},
//	3: {"deltaIPBytesUploaded", -1, 0, 4, nil},
//	4: {"deltaIPBytesdwnloaded", -1, 0, 4, nil},
//	6: {"pcsRealm", -1, 0, 4, nil},
//	7: {"policyControlSessionId", -1, 0, 4, nil},
//}
//
//var qualisage = map[int]defStructs.Champ{
//	1: {"quadrons", -1, 0, 4, nil},
//	2: {"type", -1, 0, 4, nil},
//}

//var cause = map[int]defStructs.Champ{
//	0: {"gsm0408Cause", -1, 0, 2, nil},
//}

//var listService = map[int]defStructs.Champ{
//	16: {"ChangeOfServiceCondition", -1, 0, 2, &ChgService},
//}

//var ChgService = map[int]defStructs.Champ{
//	1:  {"ratingGroup", -1, 7, 3, nil},
//	2:  {"chargingRuleBaseName", -1, 2, 3, nil},
//	3:  {"resultCode", -1, 0, 3, nil},
//	4:  {"localSequenceNumber", -1, 0, 3, nil},
//	5:  {"timeOfFirstUsage", -1, 6, 3, nil},
//	6:  {"timeOfLastUsage", -1, 6, 3, nil},
//	7:  {"timeUsage", -1, 7, 3, nil},
//	8:  {"serviceConditionChange", -1, 0, 3, nil},
//	9:  {"qoSInformationNeg", -1, 0, 3, &EPCQoSInformation},
//	10: {"servingNodeAddress", -1, 0, 3, &IPbinAdress2},
//	12: {"datavolumeFBCUplink", -1, 0, 3, nil},
//	13: {"datavolumeFBCDownlink", -1, 0, 2, nil},
//	14: {"timeOfReport", -1, 6, 3, nil},
//	16: {"failureHandlingContinue", -1, 0, 3, nil},
//	17: {"serviceIdentifier", -1, 0, 3, nil},
//	20: {"userLocationInformation", -1, 0, 3, nil},
//}

//var serving = map[int]defStructs.Champ{
//	10: {"servingNodeType", -1, 0, 2, nil},
//}

//var SubscriptionID = map[int]defStructs.Champ{
//0: {"subscriptionIDType", -1, 0, 2, nil},
//1: {"subscriptionIDData", -1, 2, 2, nil},
//}

//var EPCQoSInformation = map[int]defStructs.Champ{
//	1: {"qci", -1, 7, 4, nil},
//	2: {"maxRequestedBandwithUL", -1, 7, 4, nil},
//	3: {"maxRequestedBandwithDL", -1, 7, 4, nil},
//	4: {"guaranteedBitrateUL", -1, 7, 4, nil},
//	5: {"guaranteedBitrateDL", -1, 7, 4, nil},
//	6: {"aRP", -1, 7, 4, nil},
//}
//
//var IPbinAdress2 = map[int]defStructs.Champ{
//	0: {"servingiPBinV4address", -1, 9, 4, nil},
//	1: {"servingiPBinV6address", -1, 9, 4, nil}}
