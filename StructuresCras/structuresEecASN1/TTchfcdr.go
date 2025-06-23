package structuresEecASN1

import "local/structures/defStructs"

// TTchfcdr
var TTchfcdr = map[int]defStructs.TabRechercheTypeDc{
	0: {"CHFRecord", 0x00ff0000, 0x00810000, 0, &Tchfcdr},
}

var Tchfcdr = map[int]defStructs.Champ{
	200: {"CHFRecord", -1, -1, &ChfRecord},
}

var ChfRecord = map[int]defStructs.Champ{
	0:  {"recordType", -1, 10, nil},
	1:  {"recordingNetworkFunctionID", -1, 1, nil},
	2:  {"subscriberIdentifier", -1, 0, &SubscriptionID},
	3:  {"nFunctionConsumerInformation", -1, 0, &NetworkFunctionInformation},
	4:  {"triggers", -1, 0, &Trigger},
	5:  {"listOfMultipleUnitUsage", -1, 0, &MultipleUnitUsage},
	6:  {"recordOpeningTime", -1, 6, nil},
	7:  {"duration", -1, 0, nil},
	8:  {"recordSequenceNumber", -1, 0, nil},
	9:  {"causeForRecClosing", -1, 0, nil},
	11: {"localRecordSequenceNumber", -1, 0, nil},
	12: {"recordExtensions", -1, 0, &ManagementEXTENSIONS},
	13: {"pDUSessionChargingInformation", -1, 0, &PDUSessionChargingInformation},
	14: {"roamingQBCInformation", -1, 0, &RoamingQBCInformation},
	15: {"sMSChargingInformation", -1, 0, &SMSChargingInformation},
	16: {"chargingSessionIdentifier", -1, 1, nil},
	17: {"serviceSpecificationInformation", -1, 0, nil},
	22: {"incompleteCDRIndication", -1, 0, &IncompleteCDRIndication},
	27: {"chargingID", -1, 0, nil},
}

var PDUSessionChargingInformation = map[int]defStructs.Champ{
	0:  {"pDUSessionChargingID", -1, 0, nil},
	1:  {"userIdentifier", -1, 0, &InvolvedParty}, // vers InvolvedParty TTtasnokia
	2:  {"userEquipmentInfo", -1, 0, &SubscriberEquipmentNumber},
	3:  {"userLocationInformation", -1, 1, nil},
	4:  {"userRoamerInOut", -1, 0, nil},
	5:  {"presenceReportingAreaInfo", -1, 0, &PresenceReportingAreaInfo},
	6:  {"pDUSessionId", -1, 0, nil},
	7:  {"networkSliceInstanceID", -1, 0, &SingleNSSAI},
	8:  {"pDUType", -1, 0, nil},
	9:  {"sSCMode", -1, 0, nil},
	10: {"sUPIPLMNIdentifier", -1, 0, nil},
	11: {"servingNetworkFunctionID", -1, 0, &ServingNetworkFunctionID},
	12: {"rATType", -1, 0, nil},
	13: {"dataNetworkNameIdentifier", -1, 1, nil},
	14: {"pDUAddress", -1, 0, &PDuaddress},
	15: {"authorizedQoSInformation", -1, 0, &AuthorizedQoSInformation},
	16: {"uETimeZone", -1, 0, nil},
	17: {"pDUSessionstartTime", -1, 6, nil},
	18: {"pDUSessionstopTime", -1, 6, nil},
	20: {"chargingCharacteristics", -1, 0, nil},
	21: {"chChSelectionMode", -1, 0, nil},
	22: {"threeGPPPSDataOffStatus", -1, 0, nil},
	23: {"rANSecondaryRATUsageReport", -1, 0, &NGRANSecondaryRATUsageReport},
	24: {"subscribedQoSInformation", -1, 0, &SubscribedQoSInformation},
	25: {"authorizedSessionAMBR", -1, 0, &SessionAMBR},
	26: {"subscribedSessionAMBR", -1, 0, &SessionAMBR},
	27: {"servingCNPLMNID", -1, 0, nil},
	28: {"sUPIunauthenticatedFlag", -1, 0, nil},
	29: {"dnnSelectionMode", -1, 0, nil},
	30: {"homeProvidedChargingID", -1, 0, nil},
	31: {"mAPDUNonThreeGPPUserLocationInfo", -1, 0, nil},
	32: {"mAPDUNonThreeGPPRATType", -1, 0, nil},
	33: {"mAPDUSessionInformation", -1, 0, &MAPDUSessionInformation},
	34: {"enhancedDiagnostics", -1, 0, &EnhancedDiagnostics},
}

var InvolvedPARTY = map[int]defStructs.Champ{
	2: {"uRN", -1, 0, nil},
	4: {"externalId", -1, 0, nil},
}

var SubscriptionID = map[int]defStructs.Champ{
	0: {"subscriptionIDType", -1, 0, nil},
	1: {"subscriptionIDData", -1, 1, nil},
}

var NetworkFunctionInformation = map[int]defStructs.Champ{
	0: {"networkFunctionality", -1, 0, nil},
	1: {"networkFunctionName", -1, 1, nil},
	2: {"networkFunctionIPv4Address", -1, 1, &networkFunctionIPv4},
	3: {"networkFunctionPLMNIdentifier", -1, 0, nil},
	4: {"networkFunctionIPv6Address", -1, 1, nil},
	5: {"networkFunctionFQDN", -1, 0, &NodeAddress},
}

var networkFunctionIPv4 = map[int]defStructs.Champ{
	2: {"networkFunctionIPv4Adr", -1, 1, nil},
}

var NodeAddress = map[int]defStructs.Champ{
	0: {"iPAddress", -1, 0, nil},
	1: {"domainName", -1, 1, nil},
}

var MultipleUnitUsage = map[int]defStructs.Champ{
	16: {"multipleUnitUsage", -1, 0, &multipleUnitusage},
}

var multipleUnitusage = map[int]defStructs.Champ{
	0: {"ratingGroup", -1, 0, nil},
	1: {"usedUnitContainers", -1, 0, &usedUnitContainer},
	2: {"uPFID", -1, 1, nil},
	3: {"multihomedPDUAddress", -1, 0, &PDuaddress},
}

var ManagementEXTENSIONS = map[int]defStructs.Champ{
	16: {"managementExtensions", -1, 0, &ManagementExtension},
}

var ManagementExtension = map[int]defStructs.Champ{
	6: {"identifier", -1, 0, nil},
	1: {"significance", -1, 0, nil},
	2: {"information", -1, 0, &ProtocolEnhancements},
}

var ProtocolEnhancements = map[int]defStructs.Champ{
	16: {"ProtocolENhancements", -1, 0, &Protocolenhancements},
}

var Protocolenhancements = map[int]defStructs.Champ{
	0: {"ieKey", -1, 1, nil},
	1: {"ieRow", -1, 1, nil},
}

var RoamingQBCInformation = map[int]defStructs.Champ{
	0: {"multipleQFIcontainer", -1, 0, &MultipleQFIContainer},
	1: {"uPFID", -1, 0, nil},
	2: {"roamingChargingProfile", -1, 0, &RoamingChargingProfile},
}

//var MultipleQFIContainer = map[int]defStructs.Champ{
//	16: {"MultipleQFIcontainer", -1, 0, &MultipleQFiContainer},
//}

var MultipleQFIContainer = map[int]defStructs.Champ{
	0:  {"qosFlowId", -1, 0, nil},
	1:  {"triggers", -1, 0, &Trigger},
	2:  {"triggerTimeStamp", -1, -1, nil},
	3:  {"dataTotalVolume", -1, 0, nil},
	4:  {"dataVolumeUplink", -1, 0, nil},
	5:  {"dataVolumeDownlink", -1, 0, nil},
	6:  {"localSequenceNumber", -1, 0, nil},
	8:  {"timeOfFirstUsage", -1, 0, nil},
	9:  {"timeOfLastUsage", -1, 0, nil},
	10: {"qoSInformation", -1, 0, &FiveGQoSInformation},
	11: {"userLocationInformation", -1, 1, nil},
	12: {"uETimeZone", -1, 0, nil},
	13: {"presenceReportingAreaInfo", -1, 0, &PresenceReportingAreaInfo},
	14: {"rATType", -1, 0, nil},
	15: {"reportTime", -1, 0, nil},
	16: {"servingNetworkFunctionID", -1, 0, &ServingNetworkFunctionID},
	17: {"threeGPPPSDataOffStatus", -1, 0, nil},
	18: {"threeGPPChargingID", -1, 0, nil},
	20: {"extensionDiagnostics", -1, 0, &EnhancedDiagnostics5G},
	21: {"qoSCharacteristics", -1, 0, nil},
	22: {"time", -1, 0, nil},
}

var RoamingChargingProfile = map[int]defStructs.Champ{
	16: {"RoamingChargingPRofile", -1, 0, &RoamingChargingprofile},
}

var RoamingChargingprofile = map[int]defStructs.Champ{
	0: {"roamingTriggers", -1, 0, &RoamingTriggers},
	1: {"partialRecordMethod", -1, 0, nil},
}

var RoamingTriggers = map[int]defStructs.Champ{
	16: {"Roamingtriggers", -1, 0, &RoamingTRiggers},
}

var RoamingTRiggers = map[int]defStructs.Champ{
	0: {"trigger", -1, 0, nil},
	1: {"triggerCategory", -1, 0, nil},
	2: {"timeLimit", -1, 0, nil},
	3: {"volumeLimit", -1, 0, nil},
	4: {"maxNbChargingConditions", -1, 0, nil},
}

var SMSChargingInformation = map[int]defStructs.Champ{
	1:  {"originatorInfo", -1, 0, &OriginatorInfo},
	2:  {"recipientInfos", -1, 0, &Recipientinfo},
	3:  {"userEquipmentInfo", -1, 0, &SubscriberEquipmentNumber},
	4:  {"userLocationInformation", -1, 1, nil},
	5:  {"uETimeZone", -1, 0, nil},
	6:  {"rATType", -1, 0, nil},
	7:  {"sMSCAddress", -1, 5, nil},
	8:  {"eventtimestamp", -1, 0, nil},
	20: {"sMDataCodingScheme", -1, 0, nil},
	21: {"sMMessageType", -1, 0, nil},
	22: {"sMReplyPathRequested", -1, 0, nil},
	23: {"sMUserDataHeader", -1, 0, nil},
	24: {"sMSStatus", -1, 0, nil},
	25: {"sMDischargeTime", -1, 0, nil},
	26: {"sMTotalNumber", -1, 0, nil},
	27: {"sMServiceType", -1, 0, nil},
	28: {"sMSequenceNumber", -1, 0, nil},
	30: {"submissionTime", -1, 6, nil},
	31: {"sMPriority", -1, 0, nil},
	32: {"messageReference", -1, 0, nil},
	33: {"messageSize", -1, 0, nil},
	34: {"messageClass", -1, 0, nil},
	35: {"sMdeliveryReportRequested", -1, 0, nil},
}

//var inCompleteCdrIndication = map[int]defStructs.Champ{
//	16: {"InCompleteCDRIndication", -1, 0, &inCompleteCdrIndication},
//}

var IncompleteCDRIndication = map[int]defStructs.Champ{
	0: {"initialLost", -1, 0, nil},
	1: {"updateLost", -1, 0, nil},
	2: {"terminationLost", -1, 0, nil},
}

var Recipientinfo = map[int]defStructs.Champ{
	16: {"RecipientinFO", -1, 0, &RecipientInfo},
}

var RecipientInfo = map[int]defStructs.Champ{
	0: {"recipientIMSI", -1, 0, nil},
	1: {"recipientMSISDN", -1, 0, nil},
	2: {"recipientOtherAddress", -1, 0, &SMAddressInfo},
	3: {"recipientSCCPAddress", -1, 0, nil},
	4: {"recipientReceivedAddress", -1, 0, &SMAddressInfo},
	5: {"sMDestinationInterface", -1, 0, &SMInterface},
	6: {"sMRecipientProtocolID", -1, 0, nil},
}

//var SMAddressInfo = map[int]defStructs.Champ{
//	16: {"smaddressInfo", -1, 0, &SmaddressInFo},
//}

var SMAddressInfo = map[int]defStructs.Champ{
	0: {"sMAddressType", -1, 0, nil},
	1: {"sMAddressData", -1, 0, nil},
	2: {"sMAddressDomain", -1, 0, &SMAddressDomain},
}

var SMAddressDomain = map[int]defStructs.Champ{
	16: {"smAddressDomain", -1, 0, &SMAddressDOMain},
}

var SMAddressDOMain = map[int]defStructs.Champ{
	0: {"sMDomainName", -1, 0, nil},
}

//var OriginatorInfo = map[int]defStructs.Champ{
//	16: {"originatorINFO", -1, 0, &OriginatorInfO},
//}

var OriginatorInfo = map[int]defStructs.Champ{
	0: {"originatorIMSI", -1, 0, nil},
	1: {"originatorMSISDN", -1, 0, nil},
	2: {"originatorOtherAddress", -1, 0, &SMAdressInfo},
	3: {"originatorSCCPAddress", -1, 0, nil},
	4: {"originatorReceivedAddress", -1, 0, &SMAdressInfo},
	5: {"sMOriginatorInterface", -1, 0, &SMInterface},
	6: {"sMOriginatorProtocolID", -1, 0, nil},
}

//var SMAdressInfo = map[int]defStructs.Champ{
//	16: {"smadAdressInfo", -1, 0, &SmadAdressInfo},
//}

var SMAdressInfo = map[int]defStructs.Champ{
	0: {"sMAddressType", -1, 0, nil},
	1: {"sMAddressData", -1, 0, nil},
	2: {"sMAddressDomain", -1, 0, &SMAddressDomain},
}

//var SMAddressDomain = map[int]defStructs.Champ{
//	16: {"smaAddressDomaine", -1, 0, &SmaAddressDomaine},
//}
//
//var SmaAddressDomaine = map[int]defStructs.Champ{
//	0: {"sMDomainName", -1, 0, nil},
//}

//var SMInterface = map[int]defStructs.Champ{
//	16: {"sminterFace", -1, 0, &SmInterFace},
//}

var SMInterface = map[int]defStructs.Champ{
	0: {"interfaceId", -1, 0, nil},
	1: {"interfaceText", -1, 0, nil},
	2: {"interfacePort", -1, 0, nil},
	3: {"interfaceType", -1, 0, nil},
}

var SubscriberEquipmentNumber = map[int]defStructs.Champ{
	0: {"subscriberEquipmentNumberType", -1, 0, nil},
	1: {"subscriberEquipmentNumberData", -1, 1, nil},
}

//var PresenceReportingAreaInfo = map[int]defStructs.Champ{
//	16: {"PresenceReportingAreaINFO", -1, 0, &PresenceReportingAreainfo},
//}

var PresenceReportingAreaInfo = map[int]defStructs.Champ{
	0: {"presenceReportingAreaIdentifier", -1, 0, nil},
	1: {"presenceReportingAreaStatus", -1, 0, nil},
	2: {"presenceReportingAreaElementsList", -1, 0, nil},
	3: {"presenceReportingAreaNode", -1, 0, nil},
}

var SingleNSSAI = map[int]defStructs.Champ{
	0: {"sST", -1, 0, nil},
	1: {"sD", -1, 0, nil},
	// 16: {"SingleNssai", -1, 0, &SingleNSSai},
}

//var SingleNSSai = map[int]defStructs.Champ{
//	0: {"sST", -1, 0, nil},
//	1: {"sD", -1, 0, nil},
//}

//var PDUaddress = map[int]defStructs.Champ{
//	0: {"Pduaddress", -1, 0, &PDUaddress},
//}

var PDuaddress = map[int]defStructs.Champ{
	0: {"pDUIPv4Address", -1, 0, &iPBinV4Address},
	1: {"pDUIPv6AddresswithPrefix", -1, 0, &PdUIPv6AddresswithPrefix},
	2: {"iPV4dynamicAddressFlag", -1, 0, nil},
	3: {"iPV6dynamicPrefixFlag", -1, 0, nil},
}

var iPBinV4Address = map[int]defStructs.Champ{
	0: {"iPBinV4Address", -1, 1, nil},
	2: {"iPTextV4Address", -1, 0, nil},
	3: {"iPTextV6Address", -1, 0, nil},
}

var PdUIPv6AddresswithPrefix = map[int]defStructs.Champ{
	3: {"PdUIPv6AddresswithPrefix", -1, 1, nil},
}

//var AuthorizedQoSInformation = map[int]defStructs.Champ{
//	16: {"AUthorizedQoSInformation", -1, 0, &AUthorizedQOSInformation},
//}

var AuthorizedQoSInformation = map[int]defStructs.Champ{
	1: {"fiveQi", -1, 0, nil},
	2: {"aRP", -1, 0, &AllocationRetentionPriority},
	3: {"priorityLevel", -1, 0, nil},
	4: {"averWindow", -1, 0, nil},
	5: {"maxDataBurstVol", -1, 0, nil},
}

//var AllocationRetentionPriority = map[int]defStructs.Champ{
//	16: {"AllocationRetentionPRIORITY", -1, 0, &AllocationRetentionPRIORITY},
//}

var AllocationRetentionPriority = map[int]defStructs.Champ{
	1: {"priorityLevel", -1, 0, nil},
	2: {"preemptionCapability", -1, 0, nil},
	3: {"preemptionVulnerability", -1, 0, nil},
}

var NGRANSecondaryRATUsageReport = map[int]defStructs.Champ{
	16: {"NgranSecondaryRATUsageReport", -1, 0, &NgranSecondaryRatuSageReport},
}

var NgranSecondaryRatuSageReport = map[int]defStructs.Champ{
	0: {"nGRANSecondaryRATType", -1, 0, nil},
	1: {"qosFlowsUsageReports", -1, 0, &QosFlowsUsageReport},
}

//var QosFlowsUsageReport = map[int]defStructs.Champ{
//	16: {"QosFlowsUsageREport", -1, 0, &QOSFlowsUsageReport},
//}

var QosFlowsUsageReport = map[int]defStructs.Champ{
	1: {"qosFlowId", -1, 0, nil},
	2: {"startTime", -1, 6, nil},
	3: {"endTime", -1, 6, nil},
	4: {"dataVolumeDownlink", -1, 0, nil},
	5: {"dataVolumeUplink", -1, 0, nil},
}

//var SubscribedQoSInformation = map[int]defStructs.Champ{
//	16: {"SubscribedQosinformation", -1, 0, &SubscribedQosInformation},
//}

var SubscribedQoSInformation = map[int]defStructs.Champ{
	1: {"fiveQi", -1, 0, nil},
	2: {"aRP", -1, 0, &AllocationRetentionPriority},
	3: {"priorityLevel", -1, 0, nil},
}

//var SessionAMBR = map[int]defStructs.Champ{
//	16: {"sessionAmbr", -1, 0, &SessionAmbr},
//}

var SessionAMBR = map[int]defStructs.Champ{
	1: {"ambrUL", -1, 1, nil},
	2: {"ambrDL", -1, 1, nil},
}

//var MAPDUSessionInformation = map[int]defStructs.Champ{
//	16: {"MapduSessionInformation", -1, 0, &MapduSessionInfor},
//}

var MAPDUSessionInformation = map[int]defStructs.Champ{
	0: {"mAPDUSessionIndicator", -1, 0, nil},
	1: {"aTSSSCapability", -1, 0, nil}}

//var EnhancedDiagnostics5G = map[int]defStructs.Champ{
//	16: {"ENhancedDiagnostics5G", -1, 0, &EnhanceDdiagnostics5G},
//}

var EnhancedDiagnostics = map[int]defStructs.Champ{
	0: {"rANNASCause", -1, 0, &rAnnasCause},
}

var EnhancedDiagnostics5G = map[int]defStructs.Champ{
	0: {"rANNASCause", -1, 0, &rAnnasCause},
}

//var rAnnasCause = map[int]defStructs.Champ{
//	16: {"RannasCause", -1, 0, &RannasCAUSE},
//}

var rAnnasCause = map[int]defStructs.Champ{
	0: {"rannasCAUSE", -1, 0, nil},
}

//var NodeAddress = map[int]defStructs.Champ{
//	0: {"iPAddress", -1, 0, nil},
//	1: {"domainName", -1, 2, nil},
//}

var usedUnitContainer = map[int]defStructs.Champ{
	16: {"UsedUnitContaineR", -1, 0, &UseDUnitContainer},
}

var UseDUnitContainer = map[int]defStructs.Champ{
	0:  {"serviceIdentifier", -1, 0, nil},
	1:  {"time", -1, 0, nil},
	2:  {"triggers", -1, 0, &Trigger},
	3:  {"triggerTimeStamp", -1, 6, nil},
	4:  {"dataTotalVolume", -1, 0, nil},
	5:  {"dataVolumeUplink", -1, 0, nil},
	6:  {"dataVolumeDownlink", -1, 0, nil},
	7:  {"serviceSpecificUnits", -1, 0, nil},
	8:  {"eventTimeStamp", -1, 0, nil},
	9:  {"localSequenceNumber", -1, 0, nil},
	10: {"ratingIndicator", -1, 0, nil},
	11: {"pDUContainerInformation", -1, 0, &PDUContainerInformation},
	12: {"quotaManagementIndicator", -1, 0, nil},
	13: {"quotaManagementIndicatorExt", -1, 0, nil},
	15: {"UNKNOWN", -1, 0, nil},
}

var Trigger = map[int]defStructs.Champ{
	0: {"sMFTrigger", -1, 0, nil},
}

//var PDUContainerInformation = map[int]defStructs.Champ{
//	16: {"PDUContaineriNformation", -1, 0, &PDUContaineriNformation},
//}

var PDUContainerInformation = map[int]defStructs.Champ{
	0:  {"chargingRuleBaseName", -1, 1, nil},
	1:  {"aFCorrelationInformation", -1, 0, nil},
	2:  {"timeOfFirstUsage", -1, 6, nil},
	3:  {"timeOfLastUsage", -1, 6, nil},
	4:  {"qoSInformation", -1, 0, &FiveGQoSInformation},
	5:  {"userLocationInformation", -1, 1, nil},
	6:  {"presenceReportingAreaInfo", -1, 0, &PresenceReportingAreaInfo},
	7:  {"rATType", -1, 0, nil},
	8:  {"sponsorIdentity", -1, 0, nil},
	9:  {"applicationServiceProviderIdentity", -1, 0, nil},
	10: {"servingNetworkFunctionID", -1, 0, &ServingNetworkFunctionID},
	11: {"uETimeZone", -1, 0, nil},
	12: {"threeGPPPSDataOffStatus", -1, 0, nil},
	13: {"qoSCharacteristics", -1, 0, nil},
	14: {"afChargingIdentifier", -1, 0, nil},
	15: {"afChargingIdString", -1, 0, nil},
	16: {"mAPDUSteeringFunctionality", -1, 0, nil},
	17: {"mAPDUSteeringMode", -1, 0, &MAPDUSteeringMode},
}

//var FiveGQoSInformation = map[int]defStructs.Champ{
//	16: {"fiveGQosInformation", -1, 0, &FiveGQoSinformation},
//}

var FiveGQoSInformation = map[int]defStructs.Champ{
	1:  {"fiveQi", -1, 0, nil},
	2:  {"aRP", -1, 0, &AllocationRetentionPriority},
	3:  {"qoSNotificationControl", -1, 0, nil},
	4:  {"reflectiveQos", -1, 0, nil},
	5:  {"maxbitrateUL", -1, 0, nil},
	6:  {"maxbitrateDL", -1, 0, nil},
	7:  {"guaranteedbitrateUL", -1, 0, nil},
	8:  {"guaranteedbitrateDL", -1, 0, nil},
	9:  {"priorityLevel", -1, 0, nil},
	10: {"averWindow", -1, 0, nil},
	11: {"maxDataBurstVol", -1, 0, nil},
	12: {"maxPacketLossRateDL", -1, 0, nil},
	13: {"maxPacketLossRateUL", -1, 0, nil},
}

//var PresenceReportingAreaInfo = map[int]defStructs.Champ{
//	16: {"PresencEreportinGareaInfo", -1, 0, &PresenceReportinGareaINFO},
//}
//
//var PresenceReportinGareaINFO = map[int]defStructs.Champ{
//	0: {"presenceReportingAreaIdentifier", -1, 0, nil},
//	1: {"presenceReportingAreaStatus", -1, 0, nil},
//	2: {"presenceReportingAreaElementsList", -1, 0, nil},
//	3: {"presenceReportingAreaNode", -1, 0, nil},
//}

var ServingNetworkFunctionID = map[int]defStructs.Champ{
	16: {"ServingNetworkFunctionId", -1, 0, &ServingNETworkFUNCtionId},
}

var ServingNETworkFUNCtionId = map[int]defStructs.Champ{
	0: {"servingNetworkFunctionInformation", -1, 0, &NetworkFunctionInformation},
	1: {"aMFIdentifier", -1, 0, nil},
}

//var NetworkFunctionInformation = map[int]defStructs.Champ{
//	0: {"networkFunctionality", -1, 0, nil},
//	1: {"networkFunctionName", -1, 2, nil},
//	2: {"networkFunctionIPv4Address", -1, 9, nil},
//	3: {"networkFunctionPLMNIdentifier", -1, 0, nil},
//	4: {"networkFunctionIPv6Address", -1, 9, nil},
//	5: {"networkFunctionFQDN", -1, 0, &NodeAddress},
//}

//var NodeAddress = map[int]defStructs.Champ{
//	0: {"iPAddress", -1, 0, nil},
//	1: {"domainName", -1, 2, nil},
//}

//var MAPDUSteeringMode = map[int]defStructs.Champ{
//	16: {"MAPDUSteeringMODE", -1, 0, &MapdusteeringMODE},
//}

var MAPDUSteeringMode = map[int]defStructs.Champ{
	0: {"steerModeValue", -1, 0, nil},
	1: {"active", -1, 0, nil},
	2: {"standby", -1, 0, nil},
	3: {"threegLoad", -1, 0, nil},
	4: {"prioAcc", -1, 0, nil},
}

//var AllocationRetentionPriority = map[int]defStructs.Champ{
//	16: {"allocationRetentionPRiority", -1, 0, &AllocationRetentionPRiority2},
//}
//
//var AllocationRetentionPRiority2 = map[int]defStructs.Champ{
//	1: {"priorityLevel", -1, 0, nil},
//	2: {"preemptionCapability", -1, 0, nil},
//	3: {"preemptionVulnerability", -1, 0, nil},
//}
