package structuresEecASN1

import "local/structures/defStructs"

// TTims
var TTims = map[int]defStructs.TabRechercheTypeDc{
	0: {"Tims", 0x00ff0000, 0x003f0000, 0, &Tims},
}

var Tims = map[int]defStructs.Champ{
	63: {"sCSCFRecord", -1, 7, &SCSCFRecord},
}

var SCSCFRecord = map[int]defStructs.Champ{
	0:  {"recordType", -1, 10, nil},
	1:  {"retransmission", -1, 0, nil},
	2:  {"sIP-Method", -1, 1, nil},
	3:  {"role-of-Node", -1, 0, nil},
	4:  {"nodeAddress", -1, 0, &NodeAddreSS},
	5:  {"session-Id", -1, 0, nil},
	6:  {"list-Of-Calling-Party-Address", -1, 0, &InvolveDParty},
	7:  {"called-Party-Address", -1, 0, &InvolveDParty},
	8:  {"privateUserID", -1, 0, nil},
	9:  {"serviceRequestTimeStamp", -1, 6, nil},
	10: {"serviceDeliveryStartTimeStamp", -1, 6, nil},
	11: {"serviceDeliveryEndTimeStamp", -1, 6, nil},
	12: {"recordOpeningTime", -1, 6, nil},
	13: {"recordClosureTime", -1, 6, nil},
	14: {"interOperatorIdentifiers", -1, 0, &interOperatorIdentifierlist},
	15: {"localRecordSequenceNumber", -1, 0, nil},
	16: {"recordSequenceNumber", -1, 0, nil},
	17: {"causeForRecordClosing", -1, 0, nil},
	18: {"incomplete-CDR-Indication", -1, 0, &incompleteCDRIndication},
	19: {"iMS-Charging-Identifier", -1, 1, nil},
	21: {"list-Of-SDP-Media-Components", -1, 0, &MediaComponentsList},
	22: {"gGSNaddress", -1, 0, &NodeAddreSS},
	23: {"serviceReasonReturnCode", -1, 0, nil},
	24: {"list-Of-Message-Bodies", -1, 0, &MessageBody},
	25: {"recordExtensions", -1, 0, &ManagemenTExtensionS},
	26: {"expiresInformation", -1, 0, nil},
	27: {"list-Of-Associated-URI", -1, 0, &InvolveDParty},
	28: {"event", -1, 2, nil},
	29: {"accessNetworkInformation", -1, 1, nil},
	34: {"numberPortabilityRouting", -1, 0, nil},
	35: {"carrierSelectRouting", -1, 0, nil},
	36: {"sessionPriority", -1, 0, nil},
	37: {"serviceRequestTimeStampFraction", -1, 0, nil},
	38: {"serviceDeliveryStartTimeStampFraction", -1, 0, nil},
	39: {"serviceDeliveryEndTimeStampFraction", -1, 0, nil},
	40: {"applicationServersInformation", -1, 0, &ApplicationServersInformation},
	41: {"requested-Party-Address", -1, 0, &InvolveDParty},
	42: {"list-Of-Called-Asserted-Identity", -1, 0, &InvolveDParty},
	43: {"online-charging-flag", -1, 0, nil},
	47: {"userLocationInformation", -1, 0, nil},
	48: {"mSTimeZone", -1, 0, nil},
	52: {"iMSEmergencyIndicator", -1, 0, nil},
	54: {"iMSVisitedNetworkIdentifier", -1, 0, nil},
	56: {"additionalAccessNetworkInformation", -1, 0, nil},
	58: {"subscriberEquipmentNumber", -1, 0, &subscriberEquipmentNumber},
	64: {"cellularNetworkInformation", -1, 0, nil},
}

var ApplicationServersInformation = map[int]defStructs.Champ{
	16: {"sequence", -1, 0, &SequenceOfApplicationServersInformation},
}

var SequenceOfApplicationServersInformation = map[int]defStructs.Champ{
	0: {"applicationServersInvolved", -1, 0, &NodeAddreSS},
	1: {"applicationProvidedCalledParties", -1, 0, &InvolvedParty},
}

var interOperatorIdentifierlist = map[int]defStructs.Champ{
	16: {"sequence", -1, 0, &sequenceOfInterOperatorIdentifierlist},
}

var sequenceOfInterOperatorIdentifierlist = map[int]defStructs.Champ{
	0: {"originatingIOI", -1, 1, nil},
	1: {"terminatingIOI", -1, 1, nil},
}

var incompleteCDRIndication = map[int]defStructs.Champ{
	0: {"aCRStartLost", -1, 1, nil},
	1: {"aCRInterimLost", -1, 1, nil},
	2: {"aCRStopLost", -1, 1, nil},
}

var MediaComponentsList = map[int]defStructs.Champ{
	16: {"sequence", -1, 0, &SequenceOfMediaComponentsList},
}

var SequenceOfMediaComponentsList = map[int]defStructs.Champ{
	0:  {"sIP-Request-Timestamp", -1, 6, nil},
	1:  {"sIP-Response-Timestamp", -1, 6, nil},
	2:  {"sDP-Media-Components", -1, 0, &SDPMediaComponent},
	3:  {"mediaInitiatorFlag", -1, 1, nil},
	4:  {"sDP-Session-Description", -1, 1, nil},
	6:  {"sIP-Request-Timestamp-Fraction", -1, 0, nil},
	7:  {"sIP-Response-Timestamp-Fraction", -1, 0, nil},
	9:  {"localGWInsertedIndication", -1, 1, nil},
	10: {"iPRealmDefaultIndication", -1, 1, nil},
}

var SDPMediaComponent = map[int]defStructs.Champ{
	16: {"sequence", -1, 0, &SequenceOfSDPMediaComponent},
}

var SequenceOfSDPMediaComponent = map[int]defStructs.Champ{
	0: {"sDP-Media-Name", -1, 1, nil},
	1: {"sDP-Media-Descriptions", -1, 1, nil},
	2: {"accessCorrelationID@gPRS-Charging-Id", -1, 0, &AccessCorrelationID},
	3: {"authorized-QoS", -1, 2, nil},
	4: {"accessCorrelationID@accessNetworkChargingIdentifier", -1, 0, &AccessCorrelationID},
}

var AccessCorrelationID = map[int]defStructs.Champ{
	2: {"gPRS-Charging-Id", -1, 0, nil},
	4: {"accessNetworkChargingIdentifier", -1, 2, nil},
}

var MessageBody = map[int]defStructs.Champ{
	0: {"content-Type", -1, 2, nil},
	1: {"content-Disposition", -1, 2, nil},
	2: {"content-Length", -1, 0, nil},
	3: {"originator", -1, 0, &InvolveDParty},
}

var ManagemenTExtensionS = map[int]defStructs.Champ{
	0:  {"version", -1, 0, nil},
	1:  {"signallingVolume", -1, 0, nil},
	2:  {"totalVolume", -1, 0, nil},
	3:  {"extraBearerVolume", -1, 0, nil},
	4:  {"msISDN", -1, 1, nil},
	5:  {"iMSI", -1, 1, nil},
	6:  {"sgsnAddress", -1, 0, &NodeAddreSS},
	7:  {"gcidSignalling", -1, 0, nil},
	8:  {"acctSessionIdSignalling", -1, 2, nil},
	9:  {"listOfAccessInformation", -1, 1, &AccessInformation},
	10: {"extraInformation", -1, 0, &ExtraInfo},
	11: {"ccf", -1, 1, nil},
	13: {"subRecordType", -1, 1, nil},
	14: {"serviceInformation", -1, 0, &ServiceInformation},
	15: {"userAgentInformation", -1, 1, nil},
	16: {"releaseReasonText", -1, 1, nil},
	17: {"emergencyCallIndication", -1, 0, nil},
	19: {"referTo", -1, 1, nil},
	20: {"listOfAcceptContact", -1, 1, nil},
	21: {"to", -1, 1, nil},
	22: {"from", -1, 1, nil},
	23: {"listOfDiversion", -1, 1, nil},
	24: {"listOfHistoryInfo", -1, 1, nil},
	25: {"servedUserAddress", -1, 1, nil},
	26: {"lastKnownActivityTimeStamp", -1, 0, nil},
	27: {"listOfPreviousNetworkIDs", -1, 1, nil},
	28: {"listOfProprietaryParameter", -1, 1, nil},
	30: {"sIP-Ringing-Timestamp", -1, 0, nil},
	31: {"sIP-Ringing-Timestamp-Fraction", -1, 0, nil},
}

var AccessInformation = map[int]defStructs.Champ{
	16: {"sequence", -1, 0, &SequenceOfAccessInformation},
}

var SequenceOfAccessInformation = map[int]defStructs.Champ{
	0: {"pcomSiemensAccessInformation", -1, 0, nil},
	1: {"timeStamp", -1, 6, nil},
}

var ExtraInfo = map[int]defStructs.Champ{
	0: {"number-Of-Occurrence", -1, 7, nil},
	1: {"info-code", -1, 7, nil},
}

var ServiceInformation = map[int]defStructs.Champ{
	0: {"closedUserGroups", -1, 2, nil},
	1: {"callForwardingIndication", -1, 2, nil},
	2: {"seeTransparentDataList", -1, 2, nil},
	3: {"fciData", -1, 2, nil},
}

var ServedUserAddress = map[int]defStructs.Champ{
	0: {"iPAddress", -1, 2, &IPAddreSS},
	1: {"port", -1, 7, nil},
}

var NodeAddreSS = map[int]defStructs.Champ{
	0: {"iPAddress", -1, 2, &IPAddreSS},
	1: {"domainName", -1, 1, nil},
}

var IPAddreSS = map[int]defStructs.Champ{
	0: {"iPBinV4Address", -1, 1, nil},
	1: {"iPBinV6Address", -1, 1, nil},
	2: {"iPTextV4Address", -1, 1, nil},
	3: {"iPTextV6Address", -1, 1, nil},
}

var InvolveDParty = map[int]defStructs.Champ{
	0: {"sIP-URI", -1, 1, nil},
	1: {"tEL-URI", -1, 1, nil},
}

var subscriberEquipmentNumber = map[int]defStructs.Champ{
	0: {"subscriberEquipmentNumberType", -1, 1, nil},
	1: {"subscriberEquipmentNumberData", -1, 1, nil},
}
