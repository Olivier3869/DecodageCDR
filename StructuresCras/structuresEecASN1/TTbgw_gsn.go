package structuresEecASN1

import "local/structures/defStructs"

// TTbgw_gsn
var TTbgw_gsn = map[int]defStructs.TabRechercheTypeDc{
	0: {"CallEventRecord", 0x00ff0000, 0x00800000, 0, &Tbgw_gsn},
	//1: {"CallEventRecordType", 0x00ff0000, 0x00800000, 0, &Tbgw_gsn},
}

// TTbgw_gsn
var Tbgw_gsn = map[int]defStructs.Champ{
	20: {"CallEventRecord", -1, 0, &Caller},
	21: {"CallEventRecordType", -1, 0, &CallerType},
}

var Caller = map[int]defStructs.Champ{
	0:   {"RecordType", -1, 7, nil},
	3:   {"ServedIMSI", -1, 5, nil},
	4:   {"ServedIMEI", -1, 5, nil},
	5:   {"SgsnAdress", -1, 0, &IPAdr},
	6:   {"MsNetworkCapability", -1, 0, nil},
	7:   {"RoutingArea", -1, 0, nil},
	8:   {"LocationAreaCode", -1, 0, nil},
	9:   {"CellIdentifier", -1, 0, nil},
	10:  {"ChargingID", -1, 7, nil},
	11:  {"GgsnAddressUsed", -1, 0, &IPAdr2},
	12:  {"AccessPointNameNI", -1, 1, nil},
	13:  {"PdpType", -1, 0, nil},
	14:  {"ServedPDaddress", -1, 0, &IP},
	15:  {"ListOfTrafficVolumes", -1, 0, &IP},
	16:  {"RecordOpeningTime", -1, 6, nil},
	17:  {"Duration", -1, 7, nil},
	18:  {"SgsnChange", -1, 0, nil},
	19:  {"CauseForRecClosing", -1, 0, nil},
	20:  {"Diagnostics", -1, 0, &Gsm},
	21:  {"RecordSequenceNumber", -1, 0, nil},
	22:  {"NodeID", -1, 1, nil},
	24:  {"LocalSequenceNumber", -1, 0, nil},
	25:  {"ApnSelectionMode", -1, 0, nil},
	26:  {"AccessPointNameOI", -1, 1, nil},
	27:  {"ServedMSISDN", -1, 16, nil},
	28:  {"ChargingCharacteristics", -1, 0, nil},
	29:  {"SystemType", -1, 0, nil},
	30:  {"CAMELinformationPDP", -1, 0, &CAMEL},
	32:  {"ChChSelectionMode", -1, 0, nil},
	33:  {"DynamicAddress", -1, 0, nil},
	101: {"PLMNidentifier", -1, 8, nil},
}

var IP = map[int]defStructs.Champ{
	0:  {"IpAddress", -1, 0, &IPAdr2},
	16: {"ChargeOfCharCondition", -1, 0, &Requested},
}

var IPAdr = map[int]defStructs.Champ{
	0: {"IPbinV4Address", -1, 3, nil},
	3: {"IPbinV6Address", -1, 3, nil},
	4: {"IPtextV4Address", -1, 3, nil},
	5: {"IPtextV6Address", -1, 3, nil},
}

var IPAdr2 = map[int]defStructs.Champ{
	0: {"IPbinV4Address", -1, 3, nil},
	1: {"IPbinV6Address", -1, 3, nil},
	2: {"IPtextV4Address", -1, 3, nil},
	3: {"IPtextV6Address", -1, 3, nil},
}

var Requested = map[int]defStructs.Champ{
	1: {"QoSrequested", -1, 0, nil},
	2: {"QoSnegotiated", -1, 0, nil},
	3: {"DataVolumeGPRSuplink", -1, 0, nil},
	4: {"DataVolumeGPRSdownlink", -1, 0, nil},
	5: {"ChangeCondition", -1, 0, nil},
	6: {"ChangeTime", -1, 6, nil},
}

var Gsm = map[int]defStructs.Champ{
	0: {"Gsm0408Cause", -1, 0, nil},
	1: {"Gsm0902MapErrorValue", -1, 0, nil},
}

var CAMEL = map[int]defStructs.Champ{
	1: {"SCFaddress", -1, 0, nil},
	2: {"ServiceKey", -1, 0, nil},
	3: {"DefaultTransactionHandling", -1, 0, nil},
	4: {"CAMELaccessPointNameNI", -1, 2, nil},
	5: {"CAMELaccessPointNameOI", -1, 2, nil},
	6: {"NumberofDPEencountered", -1, 0, nil},
	7: {"LevelofCamelService", -1, 0, nil},
	8: {"FreeFormatData", -1, 0, nil},
	9: {"FFDappendIndicator", -1, 0, nil},
}

var CallerType = map[int]defStructs.Champ{
	0:  {"RecordType", -1, 7, nil},
	3:  {"ServedIMSI", -1, 8, nil},
	4:  {"GgsnAddress", -1, 0, &IPADRR},
	5:  {"ChargingID", -1, 0, nil},
	6:  {"SgsnAddress", -1, 0, &IPADRR2},
	7:  {"AccessPointNameNI", -1, 2, nil},
	8:  {"PdpType", -1, 0, nil},
	9:  {"ServedPDaddress", -1, 0, &SPDaddress},
	11: {"DynamicAddressFlag", -1, 0, nil},
	12: {"ListOfTrafficVolumes", -1, 0, &ListOfTraffic},
	13: {"RecordOpeningTime", -1, 6, nil},
	14: {"Duration", -1, 0, nil},
	15: {"CauseForRecClosing", -1, 0, nil},
	17: {"RecordSequenceNumber", -1, 0, nil},
	18: {"NodeID", -1, 2, nil},
	19: {"RecordExtension", -1, 0, &RecExt},
	20: {"LocalSequenceNumber", -1, 0, nil},
	21: {"ApnSelectionMode", -1, 0, nil},
	22: {"ServedMSISDN", -1, 16, nil},
	23: {"ChargingCharacteristics", -1, 0, nil},
	24: {"ChChSelectionMode", -1, 0, nil},
	25: {"IMSsignalContext", -1, 0, nil},
	26: {"ExternalChargingID", -1, 0, nil},
	27: {"SgsnPLMNidentifier", -1, 0, nil},
	29: {"ServedIMEISV", -1, 0, nil},
	30: {"rATType", -1, 0, nil},
	31: {"MsTimeZone", -1, 0, nil},
	32: {"UserLocationInformation", -1, 0, &UserLocInfo},
}

var IPADRR = map[int]defStructs.Champ{
	0: {"IPbinV4Address", -1, 9, nil},
	1: {"IPbinV6Address", -1, 9, nil},
	2: {"IPtextV4Address", -1, 2, nil},
	3: {"IPtextV6Address", -1, 2, nil},
}

var IPADRR2 = map[int]defStructs.Champ{
	0: {"IPbinV4Address", -1, 9, nil},
	1: {"IPbinV6Address", -1, 9, nil},
	2: {"IPtextV4Address", -1, 2, nil},
	3: {"IPtextV6Address", -1, 2, nil},
}

var SPDaddress = map[int]defStructs.Champ{
	0: {"Ipaddress", -1, 0, &IpAddress},
	1: {"ETSIaddress", -1, 0, nil},
}

var IpAddress = map[int]defStructs.Champ{
	0: {"IPbinV4Address", -1, 9, nil},
	1: {"IPbinV6Address", -1, 9, nil},
	2: {"IPtextV4Address", -1, 2, nil},
	3: {"IPtextV6Address", -1, 2, nil},
}

var ListOfTraffic = map[int]defStructs.Champ{
	16: {"ChangeOfCharCondition", -1, 0, &ChOfChCond},
}

var ChOfChCond = map[int]defStructs.Champ{
	2: {"QoSnegotiated", -1, 0, &QoSneg},
	3: {"DataVolumeGPRSuplink", -1, 0, nil},
	4: {"DataVolumeGPRSdownlink", -1, 0, nil},
	5: {"ChangeCondition", -1, 0, nil},
	6: {"ChangeTime", -1, 6, nil},
	8: {"userLocationInformation", -1, 0, nil},
}

var QoSneg = map[int]defStructs.Champ{
	0: {"GsmQoSinformaion", -1, 0, nil},
	1: {"UmtsQoSinformation", -1, 0, &UmtsQoSInfo},
}

var UmtsQoSInfo = map[int]defStructs.Champ{
	0: {"Reability", -1, 0, nil},
	1: {"Dealy", -1, 0, nil},
	2: {"Precedence", -1, 0, nil},
	3: {"PeakThrouput", -1, 0, nil},
	4: {"MeanThrouput", -1, 0, nil},
}

var RecExt = map[int]defStructs.Champ{
	1: {"Signifiance", -1, 0, nil},
	2: {"Information", -1, 0, nil},
}

var UserLocInfo = map[int]defStructs.Champ{
	0: {"IPbinV4Address", -1, 9, nil},
	1: {"IPbinV6Address", -1, 9, nil},
	2: {"IPtextV4Address", -1, 2, nil},
	3: {"IPtextV6Address", -1, 2, &IPtextV6},
}

var IPtextV6 = map[int]defStructs.Champ{
	0: {"IPbinV4Address", -1, 9, nil},
	1: {"IPbinV6Address", -1, 9, nil},
	2: {"IPtextV4Address", -1, 2, nil},
	3: {"IPtextV6Address", -1, 2, nil},
}
