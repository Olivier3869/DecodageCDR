package structuresEecASN1

import "local/structures/defStructs"

// TTmmTelTas
var TTmmTelTas = map[int]defStructs.TabRechercheTypeDc{
	0: {"MMtelTAS", 0x00ff0000, 0x00170000, 0, &TmmTelTas},
}

// TTmmTelTas
var TmmTelTas = map[int]defStructs.Champ{
	17: {"MMtelTAS", -1, 0, &MMtel},
}

var MMtel = map[int]defStructs.Champ{
	0:   {"recordType", -1, 0, nil},
	2:   {"sIP_Method", -1, 2, nil},
	3:   {"role_of_Node", -1, 0, nil},
	4:   {"nodeAddress", -1, 0, &NodeADRESS},
	5:   {"session_Id", -1, 2, nil},
	6:   {"list_Of_Calling_Party_Address", -1, 0, &list_Of_Calling_Party_Address1},
	7:   {"called_Party_Address", -1, 0, &called_Party_Address1},
	9:   {"serviceRequestTimeStamp", -1, 6, nil},
	10:  {"serviceDeliveryStartTimeStamp", -1, 6, nil},
	11:  {"serviceDeliveryEndTimeStamp", -1, 6, nil},
	12:  {"recordOpeningTime", -1, 6, nil},
	13:  {"recordClosureTime", -1, 6, nil},
	14:  {"interOperatorIdentifiers", -1, 0, &InterOperatorIdentifierlist},
	15:  {"localRecordSequenceNumber", -1, 0, nil},
	16:  {"recordSequenceNumber", -1, 0, nil},
	17:  {"causeForRecordClosing", -1, 0, nil},
	19:  {"iMS_Charging_Identifier", -1, 2, nil},
	21:  {"list_Of_SDP_Media_Components", -1, 0, &list_Of_SDP_Media_Components1},
	23:  {"serviceReasonReturnCode", -1, 0, nil},
	25:  {"recordExtensions", -1, 0, &ManagementExtensionS},
	29:  {"accessNetworkInformation", -1, 2, nil},
	30:  {"serviceContextID", -1, 2, nil},
	31:  {"list_of_subscription_ID", -1, 0, &SubscriptionId},
	32:  {"list_Of_Early_SDP_Media_Components", -1, 0, &Early_Media_Components_List1},
	34:  {"numberPortabilityRouting", -1, 0, nil},
	37:  {"serviceRequestTimeStampFraction", -1, 0, nil},
	38:  {"serviceDeliveryStartTimeStampFraction", -1, 0, nil},
	39:  {"serviceDeliveryEndTimeStampFraction", -1, 0, nil},
	43:  {"online_charging_flag", -1, 0, nil},
	50:  {"servedPartyIPAddress", -1, 0, &ServedPartyIPAddress},
	55:  {"listOfReasonHeader", -1, 0, &listOfReasonHeader1},
	56:  {"additionalAccessNetworkInformation", -1, 2, nil},
	101: {"requested_Party_Address", -1, 0, &requested_Party_Address1},
	102: {"list_Of_Called_Asserted_Identity", -1, 0, &list_Of_Called_Asserted_Identity1},
	104: {"outgoingSessionId", -1, 2, nil},
	106: {"list_Of_AccessTransferInformation", -1, 0, &AccessTransferInformation2},
	110: {"mMTelInformation", -1, 0, &mmTelInformation},
}

var NodeADRESS = map[int]defStructs.Champ{
	0: {"iPAddress", -1, 10, &IPADRESS},
	1: {"domaineName", -1, 2, nil},
}

var IPADRESS = map[int]defStructs.Champ{
	0: {"iPBinV4Address", -1, 0, nil},
	1: {"iPBinV6Address", -1, 0, nil},
	2: {"iPTextV4Address", -1, 0, nil},
	3: {"iPTextV6Address", -1, 0, nil},
}

var list_Of_Calling_Party_Address1 = map[int]defStructs.Champ{
	0: {"LCPA_sIP_URI", -1, 2, nil},
	1: {"LCPA_tEL_URI", -1, 2, nil},
}

var LCPA = map[int]defStructs.Champ{
	0: {"LCPA_sIP_URI", -1, 2, nil},
	1: {"LCPA_tEL_URI", -1, 2, nil},
}

var called_Party_Address1 = map[int]defStructs.Champ{
	0: {"c_d_sIP_URI", -1, 2, nil},
	1: {"c_d_tEL_URI", -1, 2, nil},
}

var InterOperatorIdentifierlist = map[int]defStructs.Champ{
	16: {"InterOperatorIdList", -1, 0, &INTEROperatorIdentifiers},
}

var INTEROperatorIdentifiers = map[int]defStructs.Champ{
	0: {"originatingIOI", -1, 2, nil},
	1: {"terminatingIOI", -1, 2, nil},
}

var list_Of_SDP_Media_Components1 = map[int]defStructs.Champ{
	16: {"Media_Components_List", -1, 0, &list_Of_SDP_Media_Components2},
}

var list_Of_SDP_Media_Components2 = map[int]defStructs.Champ{
	0: {"sIP_Request_Timestamp", -1, 6, nil},
	1: {"sIP_Response_Timestamp", -1, 6, nil},
	2: {"sDP_Media_Components", -1, 0, &SDP_Media_Comp},
	3: {"mediaInitiatorFlag", -1, 0, nil},
	4: {"sDP_Session_Description", -1, 0, &SDP_Session_Description},
	6: {"sIP_Request_Timestamp_Fraction", -1, 2, nil},
	7: {"sIP_Response_Timestamp_Fraction", -1, 0, nil},
	8: {"sDP_Type", -1, 0, nil},
}

var SDP_Session_Description = map[int]defStructs.Champ{
	25: {"Chaine_sDP_Session_Description", -1, 2, nil},
}

var ManagementExtensionS = map[int]defStructs.Champ{
	53:  {"ringBackTone", -1, 0, nil},
	200: {"node_id", -1, 0, nil},
	211: {"disconnectingParty", -1, 0, nil},
	271: {"announcementPlayed", -1, 0, nil},
	274: {"servedIMEISV", -1, 8, nil},
	276: {"servedTASAddress", -1, 0, &Address_Type},
	280: {"privacyHdr", -1, 2, nil},
	282: {"cgiType", -1, 0, nil},
	300: {"callDuration", -1, 0, nil},
	301: {"chargingFunctionAddress", -1, 0, &Charging_Function_Address},
	303: {"chargeBandNumber", -1, 0, nil},
	304: {"transactionType", -1, 0, nil},
	305: {"lastKnownPANI", -1, 2, nil},
	306: {"internalCauseCode", -1, 0, nil},
	307: {"cgiInformation", -1, 0, nil},
	308: {"vplmnId", -1, 0, nil},
	309: {"visitedNetworkIdentifier", -1, 2, nil},
	500: {"iN_AMA_Extension", -1, 0, &IN_AMA_EXT_Type},
	501: {"multiSIMIndicator", -1, 0, nil},
	502: {"onlineChargingServiceReturnCode", -1, 0, nil},
	503: {"locationRoutNum", -1, 0, nil},
	504: {"lrnSoInd", -1, 0, nil},
	505: {"lrnQuryStatus", -1, 0, nil},
}

var SubscriptioniD = map[int]defStructs.Champ{
	17: {"SubscriptionID", -1, 0, &Subscription},
}

var Subscription = map[int]defStructs.Champ{
	0: {"subscriptionIDType", -1, 0, nil},
	1: {"subscriptionIDData", -1, 2, nil},
}

var Address_Type = map[int]defStructs.Champ{
	0: {"nai", -1, 0, nil},
	1: {"npi", -1, 0, nil},
	2: {"addressDigit", -1, 2, nil},
}

var Charging_Function_Address = map[int]defStructs.Champ{
	0: {"ccf", -1, 0, &ccf_NodeAddress},
	1: {"ecf", -1, 0, &ecf_NodeAddress},
}

var ccf_NodeAddress = map[int]defStructs.Champ{
	0: {"ccf_iPAddress", -1, 10, &ccf_IPAddress},
	1: {"ccf_domainName", -1, 0, nil},
}

var ccf_IPAddress = map[int]defStructs.Champ{
	0: {"ccf_iPBinV4Address", -1, 10, nil},
	1: {"ccf_iPBinV6Address", -1, 0, nil},
	2: {"ccf_iPTextV4Address", -1, 2, nil},
	3: {"ccf_iPTextV6Address", -1, 0, nil},
}

var ecf_NodeAddress = map[int]defStructs.Champ{
	0: {"ecf_iPAddress", -1, 10, &ecf_IPAddress},
	1: {"ecf_domainName", -1, 0, nil},
}

var ecf_IPAddress = map[int]defStructs.Champ{
	0: {"ecf_iPBinV4Address", -1, 0, nil},
	1: {"ecf_iPBinV6Address", -1, 0, nil},
	2: {"ecf_iPTextV4Address", -1, 2, nil},
	3: {"ecf_iPTextV6Address", -1, 0, nil},
}

var IN_AMA_EXT_Type = map[int]defStructs.Champ{
	0:  {"recordType", -1, 0, nil},
	1:  {"callTransactionType", -1, 0, nil},
	2:  {"chargedPartyId", -1, 0, nil},
	3:  {"chargingTimeData", -1, 0, &ChargingTimeDataType},
	4:  {"timeQuality", -1, 0, nil},
	5:  {"mSCID", -1, 0, nil},
	6:  {"otherParty", -1, 0, nil},
	7:  {"basicService", -1, 0, nil},
	8:  {"callingPartyNumber", -1, 0, nil},
	9:  {"cAMELDestinationNumber", -1, 0, nil},
	10: {"correlationId", -1, 0, nil},
	11: {"chargeInfos", -1, 0, nil},
	12: {"serviceKey", -1, 0, nil},
}

var ChargingTimeDataType = map[int]defStructs.Champ{
	0: {"inStartOfChargingDate", -1, 0, nil},
	1: {"inStartOfChargingTime", -1, 0, &InStartOfChargingTimeType},
	2: {"inCallDuration", -1, 0, nil},
}

var InStartOfChargingTimeType = map[int]defStructs.Champ{
	0: {"inTimeStamp", -1, 0, nil},
	1: {"inTimeStamp2", -1, 0, nil},
}

var SubscriptionId = map[int]defStructs.Champ{
	0: {"subscriptionIDType", -1, 2, nil},
	1: {"subscriptionIDData", -1, 2, nil},
}

var Early_Media_Components_List1 = map[int]defStructs.Champ{
	16: {"Early_Media_Components_List", -1, 0, &Early_Media_Components_List2},
}

var Early_Media_Components_List2 = map[int]defStructs.Champ{
	0: {"Early_sDP_Offer_Timestamp", -1, 6, nil},
	1: {"Early_sDP_Answer_Timestamp", -1, 6, nil},
	2: {"Early_sDP_Media_Components", -1, 0, &Early_SDP_Media_Component1},
	3: {"Early_mediaInitiatorFlag", -1, 0, nil},
	4: {"Early_sDP_Session_Description", -1, 0, &Early_sDP_Session_Description1},
	5: {"Early_sDP_Type", -1, 0, nil},
}

var SDP_Media_Comp = map[int]defStructs.Champ{
	16: {"SDP_Media_Component", -1, 0, &SDP_Media_Comp2},
}

var SDP_Media_Comp2 = map[int]defStructs.Champ{
	0: {"sDP_Media_Name", -1, 2, nil},
	1: {"sDP_Media_Descriptions", -1, 0, &sDP_Media_Descriptions1},
}

var sDP_Media_Descriptions1 = map[int]defStructs.Champ{
	25: {"Chaine_sDP_Media_Descrip", -1, 2, nil},
}

var sDP_Session_Description1 = map[int]defStructs.Champ{
	25: {"Chaine_sDP_Session_Description", -1, 2, nil},
}

var Early_SDP_Media_Component1 = map[int]defStructs.Champ{
	16: {"Early_SDP_Media_Components", -1, 0, &Early_sdp_Media_Desciptions},
}

var Early_sdp_Media_Desciptions = map[int]defStructs.Champ{
	0: {"Early_sDP_Media_Name", -1, 2, nil},
	1: {"Early_sDP_Media_Description", -1, 2, &Early_SDP_Media_Description},
}

var Early_SDP_Media_Description = map[int]defStructs.Champ{
	25: {"Early_SDP_Media_Desc", -1, 2, nil},
}

var Early_sDP_Session_Description1 = map[int]defStructs.Champ{
	25: {"Early_SDP_Session_Description", -1, 2, nil},
}

var ServedPartyIPAddress = map[int]defStructs.Champ{
	0: {"iPBinV4Address", -1, 0, nil},
	1: {"iPBinV6Address", -1, 0, nil},
	2: {"iPTextV4Address", -1, 0, nil},
	3: {"iPTextV6Address", -1, 0, nil},
}

var listOfReasonHeader1 = map[int]defStructs.Champ{
	25: {"ReasonHeaderInformation", -1, 2, nil},
}

var requested_Party_Address1 = map[int]defStructs.Champ{
	0: {"sIP_URI", -1, 2, nil},
	1: {"tEL_URI", -1, 2, nil},
}

var list_Of_Called_Asserted_Identity1 = map[int]defStructs.Champ{
	0: {"list_sIP_URI", -1, 2, nil},
	1: {"list_tEL_URI", -1, 2, nil},
}

var AccessTransferInformation2 = map[int]defStructs.Champ{
	16: {"AccessTransferInfo", -1, 0, &AccessTransferInfo1},
}

var AccessTransferInfo1 = map[int]defStructs.Champ{
	0: {"accessTransferType", -1, 0, nil},
	1: {"accessNetworkInformation", -1, 2, nil},
	2: {"additionalAccessNetworkInformation", -1, 2, nil},
}

var mmTelInformation = map[int]defStructs.Champ{
	0: {"listofsupplservices", -1, 0, &listofsupplservices1},
}

var listofsupplservices1 = map[int]defStructs.Champ{
	17: {"SupplService", -1, 0, &SupplService1},
}

var SupplService1 = map[int]defStructs.Champ{
	0: {"serviceType", -1, 0, nil},
	1: {"serviceMode", -1, 0, nil},
	2: {"numberOfDiversions", -1, 0, nil},
	3: {"associated_Party_Address", -1, 0, &InvolvedPartyB},
	4: {"serviceId", -1, 0, nil},
	5: {"changeTime", -1, 0, nil},
	6: {"numberOfParticipants", -1, 0, nil},
	7: {"participantActionType", -1, 0, nil},
}

var InvolvedPartyB = map[int]defStructs.Champ{
	0: {"Supp_sIP_URI", -1, 2, nil},
	1: {"Supp_tEL_URI", -1, 2, nil},
}
