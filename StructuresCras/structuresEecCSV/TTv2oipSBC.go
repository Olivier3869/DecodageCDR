package structuresEecCSV

import "local/structures/defStructs"

// TTv2oipSBC
var TTv2oipSBC = map[int]defStructs.TabRechercheTypeDc{

	0: {"TTv2oipSBC", 0x00000000, 0x00000000, 0x2c, &Tv2oipsbc},
}

var Tv2oipsbc = map[int]defStructs.Champ{
	0:  {"Acct-Status-Type", -1, 0, nil},
	1:  {"NAS-IP-Address", -1, 0, nil},
	2:  {"NAS-Port", -1, 0, nil},
	3:  {"Acct-Session-Id", -1, 0, nil},
	4:  {"Acme-Session-Ingress-CallId", -1, 0, nil},
	5:  {"Acme-Session-Egress-CallId", -1, 0, nil},
	6:  {"Acme-Session-Protocol-Type", -1, 0, nil},
	7:  {"Acme-Session-Generic-Id", -1, 0, nil},
	8:  {"Calling-Station-Id", -1, 0, nil},
	9:  {"Called-Station-Id", -1, 0, nil},
	10: {"Acct-Terminate-Cause", -1, 0, nil},
	11: {"Acct-Session-Time", -1, 0, nil},
	12: {"h323-setup-time", -1, 22, nil},
	13: {"h323-connect-time", -1, 22, nil},
	14: {"h323-disconnect-time", -1, 22, nil},
	15: {"h323-disconnect-cause", -1, 0, nil},
	16: {"Acme-Session-Egress-Realm", -1, 0, nil},
	17: {"Acme-Session-Ingress-Realm", -1, 0, nil},
	18: {"Acme-FlowId_FS1_F", -1, 0, nil},
	19: {"Acme-FlowType_FS1_F", -1, 0, nil},
	20: {"Nu", -1, 0, nil},
	21: {"Nu", -1, 0, nil},
	22: {"Nu", -1, 0, nil},
	23: {"Nu", -1, 0, nil},
	24: {"Nu", -1, 0, nil},
	25: {"Nu", -1, 0, nil},
	26: {"Nu", -1, 0, nil},
	27: {"Nu", -1, 0, nil},
	28: {"Nu", -1, 0, nil},
	29: {"Nu", -1, 0, nil},
	30: {"Acme-calling-RTCP-Packets-Lost_FS1", -1, 0, nil},
	31: {"Acme-calling-RTCP-Avg-Jitter_FS1", -1, 0, nil},
	32: {"Acme-calling-RTCP-Avg-Lattency_FS1", -1, 0, nil},
	33: {"Acme-calling-RTCP-MaxJitter_FS1", -1, 0, nil},
	34: {"Acme-calling-RTCP-MaxLattency_FS1", -1, 0, nil},
	35: {"Acme-calling-RTP-Packets-Lost_FS1", -1, 0, nil},
	36: {"Acme-calling-RTP-Avg-Jitter_FS1", -1, 0, nil},
	37: {"Acme-calling-RTP-MaxJitter_FS1", -1, 0, nil},
	38: {"Acme-calling-Octets_FS1", -1, 0, nil},
	39: {"Acme-calling-Packets_FS1", -1, 0, nil},
	40: {"Acme-FlowID_FS1_R", -1, 0, nil},
	41: {"Acme-FlowType_FS1_R", -1, 0, nil},
	42: {"Nu", -1, 0, nil},
	43: {"Nu", -1, 0, nil},
	44: {"Nu", -1, 0, nil},
	45: {"Nu", -1, 0, nil},
	46: {"Nu", -1, 0, nil},
	47: {"Nu", -1, 0, nil},
	48: {"Nu", -1, 0, nil},
	49: {"Nu", -1, 0, nil},
	50: {"Nu", -1, 0, nil},
	51: {"Nu", -1, 0, nil},
	52: {"Acme-called-RTCP-Packets-Lost_FS1", -1, 0, nil},
	53: {"Acme-called-RTCP-Avg-Jitter_FS1", -1, 0, nil},
	54: {"Acme-called-RTCP-Avg-Lattency_FS1", -1, 0, nil},
	55: {"Acme-called-RTCP-MaxJitter_FS1", -1, 0, nil},
	56: {"Acme-called-RTCP-MaxLattency_FS1", -1, 0, nil},
	57: {"Acme-called-RTP-Packets-Lost_FS1", -1, 0, nil},
	58: {"Acme-called-RTP-Avg-Jitter_FS1", -1, 0, nil},
	59: {"Acme-called-RTP-MaxJitter_FS1", -1, 0, nil},
	60: {"Acme-called-Octets_FS1", -1, 0, nil},
	61: {"Acme-called-Packets_FS1", -1, 0, nil},
	62: {"Acme-Session-Charging-Vector", -1, 0, nil},
	63: {"Acme-Session-Charging-Function-Address", -1, 0, nil},
	64: {"Acme-Firmware-Version", -1, 0, nil},
	65: {"Nu", -1, 0, nil},
	66: {"Nu", -1, 0, nil},
	67: {"Nu", -1, 0, nil},
	68: {"Nu", -1, 0, nil},
	69: {"Nu", -1, 0, nil},
	70: {"Nu", -1, 0, nil},
	71: {"Nu", -1, 0, nil},
	72: {"Acme-P-Asserted-ID", -1, 0, nil},
	73: {"Acme-Ingress-Local-Addr", -1, 0, nil},
	74: {"Acme-Ingress-Remote-Addr", -1, 0, nil},
	75: {"Acme-Egress-Local-Addr", -1, 0, nil},
	76: {"Acme-Egress-Remote-Addr", -1, 0, nil},
	77: {"Nu", -1, 0, nil},
	78: {"Nu", -1, 0, nil},
	79: {"Nu", -1, 0, nil},
	80: {"Acme-Disconnect-Cause", -1, 0, nil},
	81: {"Acme-SIP-Status", -1, 0, nil},
	82: {"Nu", -1, 0, nil},
	83: {"Nu", -1, 0, nil},
}
