package spgw

import "local/structures/defStructs"

var ListOfTrafficVolumesExt = map[int]defStructs.Champ{
	16: {"ChangeOfCharCondition", -1, 0, &ChgServiceExt},
}

var ChgServiceExt = map[int]defStructs.Champ{
	1:  {"qosRequested", -1, 0, nil},
	2:  {"qosNegotiated", -1, 0, nil},
	3:  {"dataVolumeGPRSUplink", -1, 0, nil},
	4:  {"dataVolumeGPRSDownlink", -1, 0, nil},
	5:  {"changeCondition", -1, 0, nil},
	6:  {"changeTime", -1, 6, nil},
	8:  {"userLocationInformation", -1, 0, nil},
	9:  {"ePCQoSInformation", -1, 0, &ePCQoSInforExt},
	10: {"chargingID", -1, 0, nil},
	11: {"presenceReportingAreaStatus", -1, 0, nil},
	12: {"userCSGInformation", -1, 0, nil},
}

var ePCQoSInforExt = map[int]defStructs.Champ{
	1: {"qCI", -1, 0, nil},
	2: {"maxRequestedBandwithUL", -1, 0, nil},
	3: {"maxRequestedBandwithDL", -1, 0, nil},
	4: {"guaranteedBitrateUL", -1, 0, nil},
	5: {"guaranteedBitrateDL", -1, 0, nil},
	6: {"aRP", -1, 0, nil},
}

var DiagnosticExt = map[int]defStructs.Champ{
	0: {"gsm0408Cause", -1, 0, nil},
	1: {"gsm0902MapErrorValue", -1, 0, nil},
	2: {"itutQ767Cause", -1, 0, nil},
	3: {"networkSpecificCause", -1, 0, &networkExtensions},
	4: {"manufacturerSpecificCause", -1, 0, &manufactureExtensions},
	5: {"positionMethodFailureCause", -1, 0, nil},
	6: {"unauthorizedLCSClientCause", -1, 0, nil},
	7: {"diameterResultCodeAndExpResult", -1, 0, nil},
}

var networkExtensions = map[int]defStructs.Champ{
	1: {"significance", -1, 0, nil},
}

var manufactureExtensions = map[int]defStructs.Champ{
	1: {"manuSignificance", -1, 0, nil},
}

var RecordExt = map[int]defStructs.Champ{
	17: {"ManagementExtensions", -1, 0, &ManagementEXT},
}

var ManagementEXT = map[int]defStructs.Champ{
	16: {"ManagementExtension", -1, 0, &MsigniExt},
}

var MsigniExt = map[int]defStructs.Champ{
	1: {"Msignificance", -1, 0, nil},
}
