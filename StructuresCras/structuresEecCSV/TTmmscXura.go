package structuresEecCSV

import "local/structures/defStructs"

// TmmscXura
var TmmscXura = map[int]defStructs.Champ{
	0:  {"MSG_ID", -1, -1, nil},
	1:  {"NODE_ID", -1, -1, nil},
	2:  {"RECIPIENT_NUMBER", -1, -1, nil},
	3:  {"NUMBER_OF_RECIPIENTS", -1, -1, nil},
	4:  {"SDR_OPERATION_TIME_LOCAL", -1, -1, nil},
	5:  {"SUBMIT_TIME_LOCAL", -1, -1, nil},
	6:  {"PDU_SIZE", -1, -1, nil},
	7:  {"A_PREPAID_IND", -1, -1, nil},
	8:  {"EVENT_TYPE", -1, -1, nil},
	9:  {"SDR_TYPE", -1, -1, nil},
	10: {"CAUSE_CODE", -1, -1, nil},
	11: {"A_MSISDN", -1, -1, nil},
	12: {"A_TYPE", -1, -1, nil},
	13: {"A_ALIAS", -1, -1, nil},
	14: {"A_DOMAIN", -1, -1, nil},
	15: {"A_IMSI", -1, -1, nil},
	16: {"A_MSC_VLR_ID", -1, -1, nil},
	17: {"A_SGSN_IP", -1, -1, nil},
	18: {"A_SGSN_OPERATOR_ID", -1, -1, nil},
	19: {"VAS_MSISDN", -1, -1, nil},
	20: {"B_MSISDN", -1, -1, nil},
	21: {"B_TYPE", -1, -1, nil},
	22: {"B_ALIAS", -1, -1, nil},
	23: {"B_DOMAIN", -1, -1, nil},
	24: {"B_IMSI", -1, -1, nil},
	25: {"B_MSC_VLR_ID", -1, -1, nil},
	26: {"B_SGSN_IP", -1, -1, nil},
	27: {"B_SGSN_OPERATOR_ID", -1, -1, nil},
	28: {"CONTENT_TYPE", -1, -1, nil},
	29: {"TRANSCODED", -1, -1, nil},
}

// TTines
var TTmmscXura = map[int]defStructs.TabRechercheTypeDc{
	0: {"TmmscXura", 0x00000000, 0x00000000, 0x2c, &TmmscXura},
}
