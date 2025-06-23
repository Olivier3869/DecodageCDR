package structuresEecCSV

import "local/structures/defStructs"

// Tines
var Tines = map[int]defStructs.Champ{
	0:  {"serial", -1, -1, nil},
	1:  {"send_date", -1, -1, nil},
	2:  {"send_date_tz", -1, -1, nil},
	3:  {"send_date_dst", -1, -1, nil},
	4:  {"event_type", -1, -1, nil},
	5:  {"event_date", -1, -1, nil},
	6:  {"event_cause", -1, -1, nil},
	7:  {"account_poid", -1, -1, nil},
	8:  {"user_login", -1, -1, nil},
	9:  {"user_line_id", -1, -1, nil},
	10: {"start_date", -1, -1, nil},
	11: {"session_id", -1, -1, nil},
	12: {"multi_session_id", -1, -1, nil},
	13: {"session_type", -1, -1, nil},
	14: {"nas_name", -1, -1, nil},
	15: {"nas_tz", -1, -1, nil},
	16: {"nas_ipv4_address", -1, -1, nil},
	17: {"user_ipv4_address", -1, -1, nil},
	18: {"user_ipv6_address", -1, -1, nil},
	19: {"theoretical_lease_end_date", -1, -1, nil},
	20: {"access_provider", -1, -1, nil},
	21: {"address_provider", -1, -1, nil},
	22: {"allocation_mode", -1, -1, nil},
	23: {"access_mode", -1, -1, nil},
	24: {"roaming_type", -1, -1, nil},
	25: {"synchro_rate", -1, -1, nil},
	26: {"caller_id", -1, -1, nil},
	27: {"circuit_id", -1, -1, nil},
	28: {"dnis", -1, -1, nil},
	29: {"MAC", -1, -1, nil},
	30: {"session_time", -1, -1, nil},
	31: {"session_time_unit", -1, -1, nil},
	32: {"octets_in", -1, -1, nil},
	33: {"octets_in_v4", -1, -1, nil},
	34: {"octets_in_v6", -1, -1, nil},
	35: {"octets_out", -1, -1, nil},
	36: {"octets_out_v4", -1, -1, nil},
	37: {"octets_out_v6", -1, -1, nil},
	38: {"disconnect_author", -1, -1, nil},
	39: {"authentication_error", -1, -1, nil},
	40: {"context_end_date", -1, -1, nil},
}

// TTines
var TTines = map[int]defStructs.TabRechercheTypeDc{
	0: {"Tclientip", 0x00000000, 0x00000000, 0x7c, &Tines},
}
