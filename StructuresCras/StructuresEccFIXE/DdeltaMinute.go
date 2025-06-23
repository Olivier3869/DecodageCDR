package structuresEccFIXE

import "local/structures/defStructs"

// TdeltaMinuteDc_CAX
var TdeltaMinuteDc_CAX = map[int]defStructs.Champ{
	0xf0: {"non_significatif", -1, 2, nil},
	0xf1: {"analogique/numeris", -1, 2, nil},
	0xf2: {"publiphone", -1, 2, nil},
	0xf3: {"operateur_tiers", -1, 2, nil},
	0xf4: {"VTG", -1, 2, nil},
	0xf6: {"offre_transgroupe3", -1, 2, nil},
	0xf7: {"offre_8FT", -1, 2, nil},
}

// TdeltaMinuteDc_SUP
var TdeltaMinuteDc_SUP = map[int]defStructs.Champ{
	0xf0: {"non_significatif", -1, 2, nil},
	0xf1: {"analogique", -1, 2, nil},
	0xf2: {"3.1Khz_ou_numeris", -1, 2, nil},
	0xf3: {"numeris_64bits", -1, 2, nil},
}

// TdeltaMinuteDc_DOMAINE
var TdeltaMinuteDc_DOMAINE = map[int]defStructs.Champ{
	0xf0f5: {"Telephone_ou_Numeris", -1, 2, nil},
	0xf3f3: {"Renseignement", -1, 2, nil},
	0xf5f1: {"RI_LAA_ou_SVAIP_TAG", -1, 2, nil},
	0xf9f8: {"(bandes)", -1, 2, nil},
	0xf9f7: {"VTG-CAX4", -1, 2, nil},
	0xf9f3: {"VTG-convergent", -1, 2, nil},
	0xf9f4: {"VTG-RI", -1, 2, nil},
	0xf9f5: {"RPV", -1, 2, nil},
	0xf9f6: {"FCT", -1, 2, nil},
	0xf6f4: {"CE-Internet", -1, 2, nil},
	0xf9f1: {"VTG-Reroutage-fixe/mobile", -1, 2, nil},
}

// TdeltaMinuteDc_SERVICE
var TdeltaMinuteDc_SERVICE = map[int]defStructs.Champ{
	0xf0f0: {"Telephone", -1, 2, nil},
	0xf0f2: {"Numeris", -1, 2, nil},
	0xf2f1: {"RI_LAA_ou_SVAIP_TAG", -1, 2, nil},
	0xf2f5: {"Centrex", -1, 2, nil},
	0xf2f7: {"Top_message", -1, 2, nil},
	0xf2f8: {"TNN", -1, 2, nil},
	0xf2f9: {"ACRES-3000", -1, 2, nil},
	0xf3f5: {"Internet", -1, 2, nil},
	0xf3f6: {"MVT", -1, 2, nil},
	0xf4f0: {"@llo", -1, 2, nil},
	0xf4f1: {"ATP", -1, 2, nil},
	0xf4f4: {"Service-payant_a_l_acte", -1, 2, nil},
	0xf4f5: {"Transphon", -1, 2, nil},
	0xf4f6: {"Reroutage_fixe/mobile_VTG/RI2G", -1, 2, nil},
	0xf5f4: {"Mes_Contacts", -1, 2, nil},
	0xf5f6: {"DC_TOP_INFO", -1, 2, nil},
	0xf5f7: {"DC_FCP", -1, 2, nil},
	0xf6f8: {"DC_MBMM", -1, 2, nil},
}

// TdeltaMinuteDc
var TdeltaMinuteDc = map[int]defStructs.Champ{
	0:  {"DOMAINE", 2, 2, &TdeltaMinuteDc_DOMAINE},
	1:  {"SERVICE", 2, 2, &TdeltaMinuteDc_SERVICE},
	2:  {"IDUS_INF", 3, 2, nil},
	3:  {"filler", 1, 2, nil},
	4:  {"IDUS_NUM", 9, 2, nil},
	5:  {"DEDE_INF", 3, 2, nil},
	6:  {"DEDE_NUM", 20, 2, nil},
	7:  {"ICAR", 14, 2, nil},
	8:  {"DUR", 6, 2, nil},
	9:  {"CAX", 1, 2, &TdeltaMinuteDc_CAX},
	10: {"TAX", 8, 2, nil},
	11: {"NDS_NUM", 10, 2, nil},
	12: {"SUP", 1, 2, &TdeltaMinuteDc_SUP},
	13: {"REF", 3, 2, nil},
	14: {"MES1", 3, 2, nil},
	15: {"MES2", 3, 2, nil},
	16: {"nb_csv", 1, 2, nil},
	17: {"c_csv1", 2, 2, nil},
	18: {"op_csv1", 1, 2, nil},
	19: {"c_csv2", 2, 2, nil},
	20: {"op_csv2", 1, 2, nil},
	21: {"c_csv3", 2, 2, nil},
	22: {"op_csv3", 1, 2, nil},
	23: {"c_csv4", 2, 2, nil},
	24: {"op_csv4", 1, 2, nil},
	25: {"c_csv5", 2, 2, nil},
	26: {"op_csv5", 1, 2, nil},
	27: {"STY", 2, 2, nil},
	28: {"COCA", 4, 2, nil},
	29: {"COCAV", 4, 2, nil},
	30: {"IDAV_INF", 3, 2, nil},
	31: {"IDAV_NUM", 20, 2, nil},
	32: {"SGTQS", 5, 2, nil},
	33: {"MNEMO", 10, 2, nil},
	34: {"CLASSDEST", 3, 2, nil},
	35: {"filler", 32, -1, nil},
}

// TdeltaMinuteHeader
var TdeltaMinuteHeader = map[int]defStructs.Champ{
	0: {"CODART", 2, 2, nil},
	1: {"APPLI", 2, 2, nil},
	2: {"VER", 6, 2, nil},
	3: {"DATTRAIT", 14, 2, nil},
	4: {"TYPE_LOT", 1, 2, nil},
	5: {"SGTQS", 5, 2, nil},
	6: {"n_ser", 3, 2, nil},
	7: {"n_seq", 3, 2, nil},
	8: {"filler", 152, -1, nil},
}

// TdeltaMinuteTrailer
var TdeltaMinuteTrailer = map[int]defStructs.Champ{
	0: {"CODART", 2, 2, nil},
	1: {"APPLI", 2, 2, nil},
	2: {"VER", 6, 2, nil},
	3: {"DATTRAIT", 14, 2, nil},
	4: {"TYPE_LOT", 1, 2, nil},
	5: {"SGTQS", 5, 2, nil},
	6: {"NBART", 8, 2, nil},
	7: {"filler", 150, -1, nil},
}

// DdeltaMinute
var DdeltaMinute = map[int]defStructs.TabRechercheTypeDc{
	0:  {"TdeltaMinuteHeader", 0xffff0000, 0xf0f00000, 188, &TdeltaMinuteHeader},
	1:  {"TdeltaMinuteTrailer", 0xffff0000, 0xf9f90000, 188, &TdeltaMinuteTrailer},
	2:  {"TdeltaMinuteRtc", 0xffff0000, 0xf0f50000, 188, &TdeltaMinuteDc},
	3:  {"TdeltaMinute118", 0xffff0000, 0xf3f30000, 188, &TdeltaMinuteDc},
	4:  {"TdeltaMinuteBandes", 0xffff0000, 0xf9f80000, 188, &TdeltaMinuteDc},
	5:  {"TdeltaMinuteVTG", 0xffff0000, 0xf9f70000, 188, &TdeltaMinuteDc},
	6:  {"TdeltaMinuteVTGConv", 0xffff0000, 0xf9f30000, 188, &TdeltaMinuteDc},
	7:  {"TdeltaMinuteVTGri", 0xffff0000, 0xf9f40000, 188, &TdeltaMinuteDc},
	8:  {"TdeltaMinuteRpv", 0xffff0000, 0xf9f50000, 160, &TdeltaMinuteDc},
	9:  {"TdeltaMinuteFct", 0xffff0000, 0xf9f60000, 160, &TdeltaMinuteDc},
	10: {"TdeltaMinuteCEinternet", 0xffff0000, 0xf6f40000, 160, &TdeltaMinuteDc},
	11: {"TdeltaMinuteVTGreroutage", 0xffff0000, 0xf9f10000, 160, &TdeltaMinuteDc},
	12: {"TdeltaMinuteVTGreroutage", 0xffff0000, 0xf5f10000, 188, &TdeltaMinuteDc},
}
