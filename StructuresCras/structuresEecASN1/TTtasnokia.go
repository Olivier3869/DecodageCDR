package structuresEecASN1

import "local/structures/defStructs"

// Niveau 4
var SDP_Session_Description1 = map[int]defStructs.Champ{
	25: {"SDP_Session_Description", -1, -1, nil},
}

var ServiceTypes = map[int]defStructs.Champ{
	0: {"serviceType", -1, -1, nil},
	1: {"serviceMode", -1, -1, nil},
	2: {"numberOfDiversions", -1, -1, nil},
	3: {"associated_Party_Adress", -1, -1, &InvolvedParty},
	4: {"serviceId", -1, -1, nil},
}

var Information2 = map[int]defStructs.Champ{
	4:  {"Initial_accessNetworkInformation", -1, 1, nil},
	13: {"subscriber_Type_Id", -1, 1, nil},
}

var ASrecordExtension = map[int]defStructs.Champ{
	0: {"tAS_Roles", -1, -1, nil},
	1: {"sServiceInformation", -1, -1, nil},
	2: {"UTF8String", -1, -1, nil},
	3: {"offline_Charging_Faillure", -1, -1, nil},
	4: {"initAccessNetInformation", -1, -1, nil},
}

// Niveau 3
var InterOperatorIdent = map[int]defStructs.Champ{
	0: {"originatingIOI", -1, 1, nil},
	1: {"terminatingIOI", -1, 1, nil},
}

var Media_Components_List2 = map[int]defStructs.Champ{
	0: {"sIP_Request_Timestamp", -1, -1, nil},
	1: {"sIP_Response_Timestamp", -1, -1, nil},
	2: {"sDP_Media_Components", -1, -1, nil},
	3: {"mediaInitiatorFlag", -1, -1, nil},
	4: {"sDP_Session_Description", -1, -1, &SDP_Session_Description1},
	6: {"IP_Request_Timestamp_Fraction", -1, -1, nil},
	7: {"sIP_Response_Timestamp_Fraction", -1, -1, nil},
	8: {"sDP_Type", -1, -1, nil},
}

var ManagementExtension3 = map[int]defStructs.Champ{
	1: {"significance", -1, -1, nil},
	2: {"information", -1, -1, &Information},
	6: {"identifier", -1, -1, nil},
}

var SubscriptionID2 = map[int]defStructs.Champ{
	0: {"subscriptionIDType", -1, -1, nil},
	1: {"subscriptionIDData", -1, 1, nil},
}

var InvolvedParty1 = map[int]defStructs.Champ{
	0: {"sIP_URI", -1, 1, nil},
	1: {"teL_URI", -1, 1, nil},
	2: {"uRN", -1, -1, nil},
	3: {"isDN_E164", -1, -1, nil},
}

var AccessTransferInformation = map[int]defStructs.Champ{
	0: {"accessTransferType", -1, -1, nil},
	1: {"inter_UE_Transfer", -1, -1, nil},
}

var SupplService = map[int]defStructs.Champ{
	17: {"supplService", -1, -1, &ServiceTypes},
}

var Information = map[int]defStructs.Champ{
	17: {"Information2", -1, -1, &Information2},
}

// Niveau 2
var ListOfInvolvedParties1 = map[int]defStructs.Champ{
	0: {"sIP_URI", -1, 1, nil},
	1: {"teL_URI", -1, 1, nil},
	2: {"uRN", -1, -1, nil},
	3: {"isDN_E164", -1, -1, nil},
}

var InvolvedParty = map[int]defStructs.Champ{
	0: {"sIP_URI", -1, 1, nil},
	1: {"teL_URI", -1, 1, nil},
	2: {"uRN", -1, -1, nil},
	3: {"isDN_E164", -1, -1, nil},
}

var InterOperatorIdentifiers = map[int]defStructs.Champ{
	16: {"InterOperatorId", -1, -1, &InterOperatorIdent},
}

var Incomplete_CDR_Indication = map[int]defStructs.Champ{
	0: {"aCRStartLost", -1, -1, nil},
	1: {"aCRInterimLost", -1, -1, nil},
	2: {"aCRStopLost", -1, -1, nil},
}

var Media_Components_List = map[int]defStructs.Champ{
	16: {"Media_Components_List1", -1, -1, &Media_Components_List2},
}

var MessageBody1 = map[int]defStructs.Champ{
	0: {"content_Type", -1, -1, nil},
	2: {"content_Length", -1, -1, nil},
}

var ManagementExtensions = map[int]defStructs.Champ{
	16: {"ManagementExtension2", -1, -1, &ManagementExtension3},
}

var SubscriptionID1 = map[int]defStructs.Champ{
	17: {"Subscription", -1, -1, &SubscriptionID2},
}

var ListOfInvolvedParties2 = map[int]defStructs.Champ{
	16: {"ListOfInvdParties", -1, -1, &InvolvedParty1},
}

var ListOfInvolvedParties3 = map[int]defStructs.Champ{
	16: {"ListOfInvdParties", -1, -1, &InvolvedParty1},
}

var AccessTransferInformation1 = map[int]defStructs.Champ{
	16: {"AccessTransferInfo2", -1, -1, &AccessTransferInformation},
}

var MMTelInformation = map[int]defStructs.Champ{
	0: {"listOfSupplServices", -1, -1, &SupplService},
}

// Niveau 1
var ASRecord = map[int]defStructs.Champ{
	0:   {"recordType", -1, -1, nil},
	1:   {"retransmission", -1, -1, nil},
	2:   {"sIP_Method", -1, -1, nil},
	5:   {"session_Id", -1, 1, nil},
	6:   {"list_Of_Calling_Party_Address", -1, -1, &ListOfInvolvedParties1},
	7:   {"called_Party_Address", -1, -1, &InvolvedParty},
	14:  {"interOperatorIdentifiers", -1, -1, &InterOperatorIdent}, //&InterOperatorIdentifiers},
	18:  {"incomplete_CDR_Indication", -1, -1, &Incomplete_CDR_Indication},
	19:  {"iMS_Charging_Identifier", -1, 1, nil},
	21:  {"list_Of_SDP_Media_Components", -1, -1, &Media_Components_List},
	22:  {"gGSNaddress", -1, -1, nil},
	23:  {"serviceReasonReturnCode", -1, -1, nil},
	24:  {"list_Of_Message_Bodies", -1, -1, &MessageBody1},
	25:  {"recordExtensions", -1, -1, &ManagementExtensions},
	29:  {"accessNetworkInformation", -1, 1, nil},
	31:  {"list_of_subscription_ID", -1, -1, &SubscriptionID1},
	33:  {"iMSCommunicationServiceIdentifier", -1, -1, nil},
	36:  {"sessionPriority", -1, -1, nil},
	41:  {"list_of_Requested_Party_Address", -1, -1, &InvolvedParty1}, //&ListOfInvolvedParties2},
	43:  {"online_Charging_Flag", -1, -1, nil},
	47:  {"userLocationInformation", -1, -1, nil},
	48:  {"mSTimeZone", -1, -1, nil},
	51:  {"fromAddress", -1, -1, nil},
	53:  {"transit_IOI_Lists", -1, -1, nil},
	54:  {"iMSVisitedNetworkIdentifier", -1, 1, nil},
	57:  {"instanceId", -1, 1, nil},
	64:  {"cellularNetworkInformation", -1, -1, nil},
	102: {"list_Of_Called_Asserted_Identity", -1, -1, &InvolvedParty1}, //&ListOfInvolvedParties3},
	104: {"outgoingSessionId", -1, 1, nil},
	106: {"list_Of_AccessTransferInformation", -1, -1, &AccessTransferInformation1},
	107: {"relatedICID", -1, -1, nil},
	109: {"tADS_Identifier", -1, -1, nil},
}

var MMTelRecord = map[int]defStructs.Champ{
	0:   {"recordType", -1, -1, nil},
	110: {"mMTelInformation", -1, -1, &MMTelInformation},
}

var SOC = map[int]defStructs.Champ{
	0:  {"answer_time", 7, 12, nil},
	1:  {"call_reference_time", 7, 12, nil},
	2:  {"called_number", 12, 5, nil},
	3:  {"calling_imeisv", 8, 0, nil},
	4:  {"calling_imsi", 8, 0, nil},
	5:  {"calling_number", 12, 5, nil},
	6:  {"calling_subs_first_ci", 4, 0, nil},
	7:  {"calling_subs_first_lac", 2, 0, nil},
	8:  {"calling_subs_first_location_type", 1, 0, nil},
	9:  {"calling_subs_first_mcc", 2, 0, nil},
	10: {"calling_subs_first_mnc", 2, 0, nil},
	11: {"calling_subs_last_ci", 4, 0, nil},
	12: {"calling_subs_last_lac", 2, 0, nil},
	13: {"calling_subs_last_location_type", 1, 0, nil},
	14: {"calling_subs_last_mcc", 2, 0, nil},
	15: {"calling_subs_last_mnc", 2, 0, nil},
	16: {"cause_for_termination", 4, 0, nil},
	17: {"charging_end_time", 7, 12, nil},
	18: {"charging_start_time", 7, 12, nil},
	19: {"facility_usage", 5, 0, nil},
	20: {"intermediate_charging_ind", 1, 0, nil},
	21: {"intermediate_chrg_cause", 5, 0, nil},
	22: {"orig_mcz_chrg_type", 1, 0, nil},
	23: {"orig_mcz_duration", 3, 0, nil},
	24: {"radio_network_type", 1, 0, nil},
	25: {"virtual_msc_id", 16, 0, nil},
}

var STC = map[int]defStructs.Champ{
	0:  {"answer_time", 7, 12, nil},
	1:  {"call_reference_time", 7, 12, nil},
	2:  {"called_imeisv", 8, 5, nil},
	3:  {"called_imsi", 8, 5, nil},
	4:  {"called_number", 12, 5, nil},
	5:  {"called_subs_first_ci", 4, 0, nil},
	6:  {"called_subs_first_lac", 2, 0, nil},
	7:  {"called_subs_first_location_type", 1, 0, nil},
	8:  {"called_subs_first_mcc", 2, 0, nil},
	9:  {"called_subs_first_mnc", 2, 0, nil},
	10: {"called_subs_last_ci", 4, 0, nil},
	11: {"called_subs_last_lac", 2, 0, nil},
	12: {"called_subs_last_location_type", 1, 0, nil},
	13: {"called_subs_last_mcc", 2, 0, nil},
	14: {"called_subs_last_mnc", 2, 0, nil},
	15: {"calling_number", 12, 5, nil},
	16: {"cause_for_termination", 4, 0, nil},
	17: {"charging_end_time", 7, 12, nil},
	18: {"charging_start_time", 7, 12, nil},
	19: {"facility_usage", 5, 0, nil},
	20: {"intermediate_charging_ind", 1, 0, nil},
	21: {"intermediate_chrg_cause", 5, 0, nil},
	22: {"radio_network_type", 1, 0, nil},
	23: {"term_mcz_chrg_type", 1, 0, nil},
	24: {"term_mcz_duration", 3, 0, nil},
	25: {"virtual_msc_id", 16, 0, nil},
}

var FW = map[int]defStructs.Champ{
	0:  {"answer_time", 7, 12, nil},
	1:  {"call_reference_time", 7, 12, nil},
	2:  {"cause_for_forwarding", 1, 0, nil},
	3:  {"cause_for_termination", 4, 0, nil},
	4:  {"charging_end_time", 7, 12, nil},
	5:  {"charging_start_time", 7, 12, nil},
	6:  {"facility_usage", 5, 0, nil},
	7:  {"forw_mcz_chrg_type", 1, 0, nil},
	8:  {"forw_mcz_duration", 3, 0, nil},
	9:  {"forwarding_first_ci", 4, 0, nil},
	10: {"forwarding_first_lac", 2, 0, nil},
	11: {"forwarding_first_location_type", 1, 0, nil},
	12: {"forwarding_first_mcc", 2, 0, nil},
	13: {"forwarding_first_mnc", 2, 0, nil},
	14: {"forwarding_imei", 8, 0, nil},
	15: {"forwarding_imeisv", 8, 0, nil},
	16: {"forwarding_imsi", 8, 0, nil},
	17: {"forwarding_last_ci", 4, 0, nil},
	18: {"forwarding_last_lac", 2, 0, nil},
	19: {"forwarding_last_location_type", 1, 0, nil},
	20: {"forwarding_last_mcc", 2, 0, nil},
	21: {"forwarding_last_mnc", 2, 0, nil},
	22: {"forwarding_msrn", 12, 0, nil},
	23: {"forwarding_number", 12, 0, nil},
	24: {"intermediate_charging_ind", 1, 0, nil},
	25: {"intermediate_chrg_cause", 5, 0, nil},
	26: {"radio_network_type", 1, 0, nil},
	27: {"virtual_msc_id", 16, 0, nil},
}

var StaticTasNokia = map[int]defStructs.Champ{
	0x30: {"Originating", -1, -1, &SOC},
	0x31: {"Terminating", -1, -1, &STC},
	0x03: {"Forwarding", -1, -1, &FW},
}

var Ttasnokia = map[int]defStructs.Champ{
	69: {"aSRecord", -1, -1, &ASRecord},
	83: {"mMTelRecord", -1, -1, &MMTelRecord},
}

var HeaderCdrTasNokia = map[int]defStructs.Champ{
	0: {"record_length", 2, 11, nil},
	1: {"record_type", 1, 0, nil},
	2: {"record_number", 4, 0, nil},
	3: {"record_status", 1, 0, nil},
	4: {"check_sum", 2, 0, nil},
	5: {"call_reference", 5, 0, nil},
	6: {"exchange_id", 10, 0, nil},
}

var HeaderBlocTasNokia = map[int]defStructs.Champ{}

var TTtasnokia = map[int]defStructs.TabRechercheTypeDc{
	0: {"TTtasnokia", 0x00ff0000, 0x00450000, 0, &Ttasnokia},
	1: {"HeaderBloc", -1, -1, 0, &HeaderBlocTasNokia},
	2: {"HeaderCDR", -1, -1, 0, &HeaderCdrTasNokia},
	3: {"StaticCDR", -1, -1, 0, &StaticTasNokia},
}
