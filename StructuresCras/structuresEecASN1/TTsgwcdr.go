package structuresEecASN1

import (
	"local/structures/defStructs"
	"local/structures/structuresEecASN1/spgw"
)

// TTsgwcdr
var TTsgwcdr = map[int]defStructs.TabRechercheTypeDc{
	0: {"TTsgwcdr", 0x00ff0000, 0x004e0000, 0, &Tsgwcdr},
}

var Tsgwcdr = map[int]defStructs.Champ{
	78: {"Tsgwcdr", -1, -1, &CallER},
}

var CallER = map[int]defStructs.Champ{
	0:  {"RecordType", 0, 0, nil},
	3:  {"ServedIMSI", -1, 0, nil},
	4:  {"SGWAddress", -1, 0, &SGWAddressExt},
	5:  {"chargingID", -1, 0, nil},
	6:  {"servingNodeAddress", -1, 0, &servingNodeAddressExt},
	7:  {"accessPointNameNI", -1, 0, nil},
	8:  {"pdpPDNType", -1, 0, nil},
	9:  {"servedPDPPDNAddress", -1, 0, &servedPDPPDNAddressExt},
	11: {"dynamicAddressFlag", -1, 0, nil},
	12: {"listOfTrafficVolumes", -1, 0, &spgw.ListOfTrafficVolumesExt},
	13: {"recordOpeningTime", -1, 6, nil},
	14: {"duration", -1, 0, nil},
	15: {"causeForRecClosing", -1, 0, nil},
	16: {"diagnostics", -1, 0, &spgw.DiagnosticExt},
	17: {"recordSequenceNumber", -1, 0, nil},
	18: {"nodeID", -1, 0, nil},
	19: {"recordExtensions", -1, 0, &spgw.RecordExt},
	20: {"localSequenceNumber", -1, 0, nil},
	21: {"apnSelectionMode", -1, 0, nil},
	22: {"servedMSISDN", -1, 0, nil},
	23: {"chargingCharacteristics", -1, 0, nil},
	24: {"chChSelectionMode", -1, 0, nil},
	25: {"iMSsignalingContext", -1, 0, nil},
	27: {"servingNodePLMNIdentifier", -1, 0, nil},
	29: {"servedIMEISV", -1, 0, nil},
	30: {"rATType", -1, 0, nil},
	31: {"mSTimeZone", -1, 0, nil},
	32: {"userLocationInformation", -1, 0, nil},
	34: {"sGWChange", -1, 0, nil},
	35: {"servingNodeType", -1, 0, nil},
	36: {"pGWAddressUsed", -1, 0, &pGWAddressUsedExt},
	38: {"startTime", -1, 0, nil},
	39: {"stopTime", -1, 0, nil},
	40: {"pDNConnectionID", -1, 0, nil},
	64: {"listOfRANSecondary", -1, 0, &LISTOfRANSec},
}

var SGWAddressExt = map[int]defStructs.Champ{
	0: {"siPBinV4Address", 0, 0, nil},
	1: {"siPBinV6Address", -1, 0, nil},
	2: {"siPTextV4Address", -1, 0, nil},
	3: {"siPTextV6Address", -1, 0, nil},
}

var servingNodeAddressExt = map[int]defStructs.Champ{
	0: {"snIPBinV4Adress", 0, 0, nil},
	1: {"snIPBinV6Adress", -1, 0, nil},
	2: {"snIPTextV4Address", -1, 0, nil},
	3: {"snIPTextV6Address", -1, 0, nil},
}

var servedPDPPDNAddressExt = map[int]defStructs.Champ{
	0: {"sPDPPDNiPAddress", 0, 0, &sPDPPDNiPAddressExt},
	1: {"sPDPPDNeTSIAddress", -1, 0, nil},
}

var sPDPPDNiPAddressExt = map[int]defStructs.Champ{
	0: {"sPDPPDNiPBinV4Adress", 0, 0, nil},
	1: {"sPDPPDNiPBinV6Adress", -1, 0, nil},
	2: {"sPDPPDNIPTextV4Address", -1, 0, nil},
	3: {"sPDPPDNIPTextV6Address", -1, 0, nil},
}

//var listOfTrafficVolumesExt = map[int]defStructs.Champ{
//	16: {"ChangeOfCharCondition", -1, 0, &ChgServiceExt},
//}
//
//var ChgServiceExt = map[int]defStructs.Champ{
//	1: {"qosRequested", -1, 0, nil},
//	2: {"qosNegotiated", -1, 0, nil},
//	3: {"dataVolumeGPRSUplink", -1, 0, nil},
//	4: {"dataVolumeGPRSDownlink", -1, 0, nil},
//	5: {"changeCondition", -1, 0, nil},
//	6: {"changeTime", -1, 0, nil},
//	8: {"userLocationInformation", -1, 0, nil},
//	9: {"ePCQoSInformation", -1, 0, &ePCQoSInforExt},
//}
//
//var ePCQoSInforExt = map[int]defStructs.Champ{
//	1: {"qCI", -1, 0, nil},
//	2: {"maxRequestedBandwithUL", -1, 0, nil},
//	3: {"maxRequestedBandwithDL", -1, 0, nil},
//	4: {"guaranteedBitrateUL", -1, 0, nil},
//	5: {"guaranteedBitrateDL", -1, 0, nil},
//	6: {"aRP", -1, 0, nil},
//}

//var diagnosticExt = map[int]defStructs.Champ{
//	0: {"gsm0408Cause", 0, 0, nil},
//	1: {"gsm0902MapErrorValue", -1, 0, nil},
//	2: {"itutQ767Cause", -1, 0, nil},
//	3: {"networkSpecificCause", -1, 0, &networkExtensions},
//	4: {"manufacturerSpecificCause", -1, 0, &manufactureExtensions},
//	5: {"positionMethodFailureCause", -1, 0, nil},
//	6: {"unauthorizedLCSClientCause", -1, 0, nil},
//}
//
//var networkExtensions = map[int]defStructs.Champ{
//	1: {"significance", -1, 0, nil},
//}
//
//var manufactureExtensions = map[int]defStructs.Champ{
//	1: {"manuSignificance", -1, 0, nil},
//}

//var recordExt = map[int]defStructs.Champ{
//	17: {"ManagementExtensions", -1, 0, &ManagementExt},
//}
//
//var ManagementExt = map[int]defStructs.Champ{
//	16: {"ManagementExtension", -1, 0, &MsigniExt},
//}
//
//var MsigniExt = map[int]defStructs.Champ{
//	1: {"Msignificance", -1, 0, nil},
//}

var pGWAddressUsedExt = map[int]defStructs.Champ{
	0: {"piPBinV4Adress", 0, 0, nil},
	1: {"piPBinV6Adress", -1, 0, nil},
	2: {"piPTextV4Address", -1, 0, nil},
	3: {"piPTextV6Address", -1, 0, nil},
}

var LISTOfRANSec = map[int]defStructs.Champ{
	16: {"LISTOfRANSecondary", -1, 0, &LISTOfRANSECONDARY},
}

var LISTOfRANSECONDARY = map[int]defStructs.Champ{
	1: {"dataVolumeUplink", -1, 0, nil},
	2: {"dataVolumeDownlink", -1, 0, nil},
	3: {"rANStartTime", -1, 0, nil},
	4: {"rANEndTime", -1, 0, nil},
	5: {"secondaryRATType", -1, 0, nil},
}
