package structuresEecASN1

import (
	"local/structures/defStructs"
	"local/structures/structuresEecASN1/spgw"
)

// niveau 4
//var serviceIPVolume = map[int]defStructs.Champ{
//	1: {"cumIPBytesUploaded", -1, -1, 4, nil},
//	2: {"cumIPBytesdwnloaded", -1, -1, 4, nil},
//	3: {"deltaIPBytesUploaded", -1, -1, 4, nil},
//	4: {"deltaIPBytesdwnloaded", -1, -1, 4, nil},
//}
//
//var qualisage = map[int]defStructs.Champ{
//	1: {"quadrons", -1, -1, 4, nil},
//	2: {"type", -1, -1, 4, nil},
//}

// niveau 3
//var ServiceLevel = map[int]defStructs.Champ{
//	1: {"sessionID", -1, -1, 2, nil},
//	2: {"serviceID", -1, -1, 4, nil},
//	3: {"serviceIPVolumeUsage", -1, -1, 4, &serviceIPVolume},
//	4: {"quotaServerFlags", -1, -1, 4, nil},
//	5: {"qualifiedUsage", -1, -1, 4, &qualisage},
//	6: {"timeOfFirstUsage", -1, -1, 4, nil},
//}

//var EPCQoSInformation = map[int]defStructs.Champ{
//	1: {"qci", -1, -1, 4, nil},
//	2: {"maxRequestedBandwithUL", -1, -1, 4, nil},
//	3: {"maxRequestedBandwithDL", -1, -1, 4, nil},
//	4: {"guaranteedBitrateUL", -1, -1, 4, nil},
//	5: {"guaranteedBitrateDL", -1, -1, 4, nil},
//	6: {"aRP", -1, -1, 4, nil},
//}
//
//var IPbinAdress2 = map[int]defStructs.Champ{
//	0: {"servingiPBinV4address", -1, -1, 4, nil},
//	1: {"servingiPBinV6address", -1, -1, 4, nil},
//}

//// niveau 2
//var Manage = map[int]defStructs.Champ{
//	100: {"serviceLevelCDR", -1, -1, 3, &ServiceLevel},
//}

//var ChgService = map[int]defStructs.Champ{
//	1:  {"ratingGroup", -1, -1, 3, nil},
//	2:  {"chargingRuleBaseName", -1, -1, 3, nil},
//	3:  {"resultCode", -1, -1, 2, nil},
//	4:  {"localSequenceNumber", -1, -1, 3, nil},
//	5:  {"timeOfFirstUsage", -1, -1, 3, nil},
//	6:  {"timeOfLastUsage", -1, -1, 3, nil},
//	7:  {"timeUsage", -1, -1, 3, nil},
//	8:  {"serviceConditionChange", -1, -1, 3, nil},
//	9:  {"qoSInformationNeg", -1, -1, 3, &EPCQoSInformation},
//	10: {"servingNodeAddress", -1, -1, 3, &IPbinAdress2},
//	12: {"datavolumeFBCUplink", -1, -1, 3, nil},
//	13: {"datavolumeFBCDownlink", -1, -1, 3, nil},
//	14: {"timeOfReport", -1, -1, 2, nil},
//	16: {"failureHandlingContinue", -1, -1, 3, nil},
//	17: {"serviceIdentifier", -1, -1, 2, nil},
//	20: {"userLocationInformation", -1, -1, 3, nil},
//}

//var RANSecondaryRATUsageREPORT = map[int]defStructs.Champ{
//	1: {"dataVolumeUplink", -1, -1, 3, nil},
//	2: {"dataVolumeDownlink", -1, -1, 3, nil},
//	3: {"rANStartTime", -1, -1, 3, nil},
//	4: {"rANEndTime", -1, -1, 3, nil},
//	5: {"secondaryRATType", -1, -1, 3, nil},
//}

// niveau 1
//var pGbinAddress = map[int]defStructs.Champ{
//	0: {"PGWV4address", -1, -1, 2, nil},
//	1: {"PGWV6address", -1, -1, 2, nil},
//}

//var serviIP = map[int]defStructs.Champ{
//	0: {"SgWiPBinV4Ad", -1, -1, 2, nil},
//	1: {"SgWiPBinV6Ad", -1, -1, 2, nil},
//}

//var servedIP = map[int]defStructs.Champ{
//	0: {"iPAdressV4", -1, -1, 2, &serviIP},
//	1: {"iPAdressV6", -1, -1, 2, &serviIP},
//}

//var cause = map[int]defStructs.Champ{
//	0: {"gsm0408Cause", -1, -1, 2, nil},
//}

//var Management_Ext = map[int]defStructs.Champ{
//	16: {"ManagementExten", -1, -1, 2, &Manage},
//}

//var listService = map[int]defStructs.Champ{
//	16: {"ChangeOfServiceCondition", -1, -1, 2, &ChgService},
//}

//var serving = map[int]defStructs.Champ{
//	10: {"servingNodeType", -1, -1, 2, nil},
//}

//var SubscriptionID = map[int]defStructs.Champ{
//	0: {"subscriptionIDType", -1, -1, 2, nil},
//	1: {"subscriptionIDData", -1, 1, 2, nil},
//}

//var listOfRANSECONDARY = map[int]defStructs.Champ{
//	16: {"RANSecondaryRATUsageReport", -1, -1, 2, &RANSecondaryRATUsageREPORT},
//}

// niveau 0
//var PgwCallER = map[int]defStructs.Champ{
//	0:  {"RecordType", -1, -1, 1, nil},
//	3:  {"ServedIMSI", -1, -1, 1, nil},
//	4:  {"PGWAddress", -1, -1, 1, &pGbinAddress},
//	5:  {"chargingID", -1, -1, 1, nil},
//	6:  {"servingNodeAddress", -1, -1, 1, &serviIP},
//	7:  {"accessPointNameNI", -1, -1, 1, nil},
//	8:  {"pdpPDNType", -1, -1, 1, nil},
//	9:  {"servedPDPPDNAddress", -1, -1, 1, &servedIP},
//	11: {"dynamicAddressFlag", -1, -1, 1, nil},
//	13: {"recordOpeningTime", -1, -1, 1, nil},
//	14: {"duration", -1, -1, 1, nil},
//	15: {"causeForRecClosing", -1, -1, 1, nil},
//	16: {"diagnostics", -1, -1, 1, &cause},
//	17: {"recordSequenceNumber", -1, -1, 1, nil},
//	18: {"nodeID", -1, -1, 1, nil},
//	19: {"recordExtensions", -1, -1, 1, &Management_Ext},
//	20: {"localSequenceNumber", -1, -1, 1, nil},
//	21: {"apnSelectionMode", -1, -1, 1, nil},
//	22: {"servedMSISDN", -1, -1, 1, nil},
//	23: {"chargingCharacteristics", -1, -1, 1, nil},
//	24: {"chChSelectionMode", -1, -1, 1, nil},
//	27: {"servingNodePLMNIdentifier", -1, -1, 1, nil},
//	29: {"servedIMEISV", -1, -1, 1, nil},
//	30: {"rATType", -1, -1, 1, nil},
//	31: {"mSTimeZone", -1, -1, 1, nil},
//	32: {"userLocationInformation", -1, -1, 1, nil},
//	34: {"listOfServiceData", -1, -1, 1, &spgw.ListService},
//	35: {"servingNodeType", -1, -1, 1, &spgw.Serving},
//	36: {"servedMNNAI", -1, -1, 1, &spgw.SubscriptionID},
//	37: {"PGWPLMNIdentifier", -1, -1, 1, nil},
//	38: {"startTime", -1, -1, 1, nil},
//	39: {"stopTime", -1, -1, 1, nil},
//	41: {"pDNConnectionID", -1, -1, 1, nil},
//	73: {"listOfRANSecondary", -1, -1, 1, &listOfRANSECONDARY},
//}

var Tpgwcdr = map[int]defStructs.Champ{
	0x4f: {"TTpgwcdr", -1, -1, &spgw.CallER},
}

// TTpgwcdr
var TTpgwcdr = map[int]defStructs.TabRechercheTypeDc{
	0: {"Tpgwcdr", 0x00ff0000, 0x004f0000, 0, &Tpgwcdr},
}
