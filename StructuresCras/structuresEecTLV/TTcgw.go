package structuresEecTLV

import "local/structures/defStructs"

// TTcgw
var TTcgw = map[int]defStructs.TabRechercheTypeDc{
	0: {"EnteteBloc", 0xffff0000, 0x80770000, 0, &Tbivsip},
	1: {"Tmmm", 0xffff0000, 0x80570000, 0, &Tbticsip},
	2: {"TmmmOL", 0xffff0000, 0x80580000, 0, &Topitml},
}

var Tbivsip = map[int]defStructs.Champ{
	0x77: {"T1_Tbivsip", -1, 0, &T2_Tbivsip},
}

var Tbticsip = map[int]defStructs.Champ{
	0x57: {"T1_Tbticsip", -1, 0, &T2_Tbticsip},
}

var Topitml = map[int]defStructs.Champ{
	0x58: {"T1_Topitml", -1, 0, &T2_Topitml},
}

var T2_Tbivsip = map[int]defStructs.Champ{
	0x02: {"T2_bivsip", -1, 0, &T3_Topitml},
}

var T2_Tbticsip = map[int]defStructs.Champ{
	0x02: {"Dc", -1, 0, &T3_Topitml},
}

var T2_Topitml = map[int]defStructs.Champ{
	0x02: {"Dc", -1, 0, &T3_Topitml},
}

var T3_Topitml = map[int]defStructs.Champ{
	0x04: {"CallingParty_Adress", -1, 0, &T3_CallingParty_Adress},
	0x05: {"Requested_Party_Address", -1, 1, nil},
	0x06: {"Translated_Called_Number", -1, 0, nil},
	0x0a: {"Service_Request_Timestamp", -1, 7, nil},
	0x0c: {"Call_Start_Timestamp", -1, 7, nil},
	0x0d: {"Call_End_TimeStamp", -1, 7, nil},
	0x0e: {"Call_Duration", -1, 0, nil},
	0x11: {"Role_of_the_Node", -1, 0, nil},
	0x12: {"Call_Status", -1, 0, nil},
	0x13: {"Session_Termination_Cause", -1, 0, nil},
	0x14: {"AS_Termination_Cause", -1, 0, nil},
	0x15: {"IMS_Charging_Identifier", -1, 1, nil},
	0x16: {"Originating_IOI", -1, 0, nil},
	0x17: {"Terminating_IOI", -1, 0, nil},
	0x19: {"Access_SIP_Session_Identifier", -1, 1, nil},
	0x1a: {"Network_SIP_Session_dentifier", -1, 0, nil},
	0x1b: {"Access_Network_Information", -1, 0, nil},
	0x1c: {"Cause_Code", -1, 0, nil},
	0x1d: {"Cause_for_Record_Closing", -1, 0, nil},
	0x1e: {"Answer_Indicator", -1, 0, nil},
	0x1f: {"Session_Id", -1, 1, nil},
	0x21: {"TrunkGroupInfo", -1, 0, nil},
	0x30: {"Service_Specific", -1, 0, &T3_Service_Specific},
	0x50: {"Codec_used", -1, 0, nil},
	0x52: {"Session_Media_Components", -1, 0, &T3_Session_Media_Components},
	0x70: {"Nodes_Identifier", -1, 0, &T3_Nodes_Identifier},
	0x71: {"AS_Releasing_Party", -1, 0, nil},
	0x73: {"Company_identifier", -1, 0, nil},
	0x74: {"Group", -1, 0, nil},
	0x77: {"SIP_Method", -1, 0, nil},
	0x78: {"Private_User_ID", -1, 1, nil},
}

var T3_CallingParty_Adress = map[int]defStructs.Champ{
	0x01: {"T4_CallingParty_Adr02", -1, 1, nil},
}

var T3_Service_Specific = map[int]defStructs.Champ{
	0x05: {"T4_Related_Call_ID", -1, 0, nil},
	0x06: {"T4_Related_Call_ID_Reason", -1, 0, nil},
	0x07: {"T4_Local_Call_ID", -1, 1, nil},
	0x08: {"T4_AS_Original_Called_Number", -1, 0, nil},
	0x09: {"T4_AS_Redirecting_Number", -1, 0, nil},
	0x0a: {"T4_Redirecting_Reason", -1, 0, nil},
	0x0b: {"T4_Original_Called_Reason", -1, 0, nil},
}

var T3_Session_Media_Components = map[int]defStructs.Champ{
	0x00: {"T4_Time_of_Media_change", 0, 0, nil},
	0x01: {"T4_Name_media", -1, 0, nil},
	0x03: {"T4_MediaDesc11", -1, 0, nil},
	0x04: {"T4_MediaDesc12", -1, 0, nil},
	0x06: {"T4_MediaDesc14", -1, 0, nil},
	0x07: {"T4_MediaDesc15", -1, 0, nil},
	0x09: {"T4_MediaDesc17", -1, 0, nil},
	0x0a: {"T4_MediaDesc18", -1, 0, nil},
	0x0c: {"T4_Duration_per_Media", -1, 0, nil},
}

var T3_Nodes_Identifier = map[int]defStructs.Champ{
	0x00: {"T4_S-CSCF_identifier", 0, 1, nil},
	0x01: {"T4_App_Server_id", -1, 1, nil},
}
