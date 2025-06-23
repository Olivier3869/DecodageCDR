package structuresEecASN1

import (
	"local/structures/defStructs"
	"local/structures/structuresEecASN1/spgw"
)

// TTepdg
var TTepdg = map[int]defStructs.TabRechercheTypeDc{
	0: {"TTepdg", 0x00ff0000, 0x00960000, 0, &Tepdg},
}

var Tepdg = map[int]defStructs.Champ{
	96: {"Tepdg", -1, 0, &EPDGRecord},
}

var EPDGRecord = map[int]defStructs.Champ{
	0:   {"RecordType", -1, 0, nil},
	3:   {"ServedIMSI", -1, 8, nil},
	4:   {"ePPDGAddress", -1, 0, &ePPDGAddress1},
	5:   {"chargingID", -1, 0, nil},
	7:   {"accessPointNameNI", -1, 2, nil},
	8:   {"pdpPDNType", -1, 0, nil},
	9:   {"servedPDPPDNAddress", -1, 0, &PDPPDNAddress},
	11:  {"dynamicAddressFlag", -1, 0, nil},
	12:  {"listOfTrafficVolumes", -1, 0, &spgw.ListOfTrafficVolumesExt},
	13:  {"recordOpeningTime", -1, 6, nil},
	14:  {"duration", -1, 7, nil},
	15:  {"causeForRecClosing", -1, 0, nil},
	16:  {"diagnostics", -1, 0, &spgw.DiagnosticExt},
	17:  {"recordSequenceNumber", -1, 0, nil},
	18:  {"nodeID", -1, 2, nil},
	19:  {"recordExtensions", -1, 0, &spgw.RecordExt},
	20:  {"localSequenceNumber", -1, 0, nil},
	21:  {"apnSelectionMode", -1, 0, nil},
	22:  {"servedMSISDN", -1, 8, nil},
	23:  {"chargingCharacteristics", -1, 0, nil},
	24:  {"chChSelectionMode", -1, 0, nil},
	25:  {"iMSsignalingContext", -1, 0, nil},
	29:  {"servedIMEISV", -1, 8, nil},
	30:  {"rATType", -1, 0, nil},
	34:  {"sGWChange", -1, 0, nil},
	36:  {"p_GWAddressUsed", -1, 0, &pGWAddressUsedExt2},
	37:  {"p_GWPLMNIdentifier", -1, 8, nil},
	38:  {"startTime", -1, 6, nil},
	39:  {"stopTime", -1, 6, nil},
	40:  {"pDNConnectionID", -1, 3, nil},
	43:  {"servedPDPPDNAddressExt", -1, 0, &servedPDPPDNAddressExt1},
	47:  {"dynamicAddressFlagExt", -1, 0, nil},
	48:  {"ePDGiPv6AddressUsed", -1, 0, &ePDGGSNAddressUsed},
	50:  {"p_GWiPv6AddressUsed", -1, 0, &p_GWGSNAddressUsed},
	51:  {"retransmission", -1, 0, nil},
	103: {"UETunnelInfo", -1, 0, &uETunnelInfo},
}

var ePPDGAddress1 = map[int]defStructs.Champ{
	0: {"siPBinV4Address", -1, 9, nil},
	1: {"siPBinV6Address", -1, 9, nil},
	2: {"siPTextV4Address", -1, 0, nil},
	3: {"siPTextV6Address", -1, 0, nil},
}

var PDPPDNAddress = map[int]defStructs.Champ{
	0: {"sPDPPDNiPAddress", -1, 0, &sPDPPDNiPAddressE},
	1: {"sPDPPDNeTSIAddress", -1, 0, nil},
}

var sPDPPDNiPAddressE = map[int]defStructs.Champ{
	0: {"sPDPPDNiPBinV4Adress", -1, 9, nil},
	1: {"sPDPPDNiPBinV6Adress", -1, 9, nil},
	2: {"sPDPPDNIPTextV4Address", -1, 0, nil},
	3: {"sPDPPDNIPTextV6Address", -1, 0, nil},
}

//var listOfTrafficVolumesExt = map[int]defStructs.Champ{
//16: {"ChangeOfCharCondition", -1, 0, &ChgServiceExt},
//}
//
//var ChgServiceExt = map[int]defStructs.Champ{
//1: {"qosRequested", -1, 0, nil},
//2: {"qosNegotiated", -1, 0, nil},
//3: {"dataVolumeGPRSUplink", -1, 0, nil},
//4: {"dataVolumeGPRSDownlink", -1, 0, nil},
//5: {"changeCondition", -1, 0, nil},
//6: {"changeTime", -1, 6, nil},
//8: {"userLocationInformation", -1, 0, nil},
//9: {"ePCQoSInformation", -1, 0, &ePCQoSInforExt},
//10: {"chargingID", -1, 0, nil},
//11: {"presenceReportingAreaStatus", -1, 0, nil},
//12: {"userCSGInformation", -1, 0, nil},
//}
//
//var ePCQoSInforExt = map[int]defStructs.Champ{
//1: {"qCI", -1, 0, nil},
//2: {"maxRequestedBandwithUL", -1, 0, nil},
//3: {"maxRequestedBandwithDL", -1, 0, nil},
//4: {"guaranteedBitrateUL", -1, 0, nil},
//5: {"guaranteedBitrateDL", -1, 0, nil},
//6: {"aRP", -1, 0, nil},
//}

//var diagnosticExt = map[int]defStructs.Champ{
//0: {"gsm0408Cause", -1, 0, nil},
//1: {"gsm0902MapErrorValue", -1, 0, nil},
//2: {"itutQ767Cause", -1, 0, nil},
//3: {"networkSpecificCause", -1, 0, &networkExtensions},
//4: {"manufacturerSpecificCause", -1, 0, &manufactureExtensions},
//5: {"positionMethodFailureCause", -1, 0, nil},
//6: {"unauthorizedLCSClientCause", -1, 0, nil},
//7: {"diameterResultCodeAndExpResult", -1, 0, nil},
//}
//
//var networkExtensions = map[int]defStructs.Champ{
//1: {"significance", -1, 0, nil},
//}
//
//var manufactureExtensions = map[int]defStructs.Champ{
//1: {"manuSignificance", -1, 0, nil},
//}

//var recordExt = map[int]defStructs.Champ{
//17: {"ManagementExtensions", -1, 0, &ManagementExt},
//}
//
//var ManagementExt = map[int]defStructs.Champ{
//16: {"ManagementExtension", -1, 0, &MsigniExt},
//}
//
//var MsigniExt = map[int]defStructs.Champ{
//1: {"Msignificance", -1, 0, nil},
//}

var pGWAddressUsedExt2 = map[int]defStructs.Champ{
	0: {"piPBinV4Adress", -1, 9, nil},
	1: {"piPBinV6Adress", -1, 9, nil},
	2: {"piPTextV4Address", -1, 0, nil},
	3: {"piPTextV6Address", -1, 0, nil},
}

var servedPDPPDNAddressExt1 = map[int]defStructs.Champ{
	0: {"iPAddress", -1, 0, &IPAddress},
	1: {"eTSIAddress", -1, 0, nil},
}

var IPAddress = map[int]defStructs.Champ{
	0: {"iPBinV4Adress", -1, 9, nil},
	1: {"iPBinV6Adress", -1, 9, nil},
	2: {"iPTextV4Address", -1, 0, nil},
	3: {"iPTextV6Address", -1, 0, nil},
}

var ePDGGSNAddressUsed = map[int]defStructs.Champ{
	0: {"ePDGiPBinV4Adress", -1, 9, nil},
	1: {"ePDGiPBinV6Adress", -1, 9, nil},
	2: {"ePDGiPTextV4Address", -1, 0, nil},
	3: {"ePDGiPTextV6Address", -1, 0, nil},
}

var p_GWGSNAddressUsed = map[int]defStructs.Champ{
	0: {"p_GWiPBinV4Adress", -1, 9, nil},
	1: {"p_GWiPBinV6Adress", -1, 9, nil},
	2: {"p_GWiPTextV4Address", -1, 0, nil},
	3: {"p_GWiPTextV6Address", -1, 0, nil},
}

var uETunnelInfo = map[int]defStructs.Champ{
	0: {"UEEndpointIP", -1, 0, &uEEndpointIP},
	1: {"UEEndpointPort", -1, 0, nil},
	2: {"EPDGEndpointIP", -1, 0, &ePDGEndpointIP},
	3: {"EPDGPor", -1, 3, nil},
}

var uEEndpointIP = map[int]defStructs.Champ{
	0: {"UEIPBinV4Adress", -1, 9, nil},
	1: {"UEIPBinV6Adress", -1, 9, nil},
	2: {"UEIPTextV4Address", -1, 0, nil},
	3: {"UEIPTextV6Address", -1, 0, nil},
}

var ePDGEndpointIP = map[int]defStructs.Champ{
	0: {"EndpointIPBinV4Adress", -1, 9, nil},
	1: {"EndpointIPBinV6Adress", -1, 9, nil},
	2: {"EndpointIPTextV4Address", -1, 0, nil},
	3: {"EndpointIPTextV6Address", -1, 0, nil},
}
