package structuresEccTLV

import "local/structures/defStructs"

// Tdpvgsm28T3t1
var Tdpvgsm28T3t1 = map[int]defStructs.Champ{
	0:  {"type", 1, -1, nil},
	1:  {"tentite", 1, -1, nil},
	2:  {"idemet", 9, -1, nil},
	3:  {"IMSI", 8, -1, nil},
	4:  {"DATE", 7, -1, nil},
	5:  {"DUREE", 4, -1, nil},
	6:  {"svcsup", 1, -1, nil},
	7:  {"tlserv", 1, -1, nil},
	8:  {"support1", 1, -1, nil},
	9:  {"systype", 1, -1, nil},
	10: {"reserve2", 1, -1, nil},
	11: {"origine", 1, -1, nil},
	12: {"indeff", 1, -1, nil},
	13: {"tterm", 1, -1, nil},
	14: {"indtax", 1, -1, nil},
	15: {"nbUT", 4, -1, nil},
	16: {"chgcar", 2, -1, nil},
	17: {"couleur", -1, -1, nil},
}

// Tdpvgsm28T3t2
var Tdpvgsm28T3t2 = map[int]defStructs.Champ{
	0: {"idpart", -1, -1, nil},
}

// Tdpvgsm28T3t3
var Tdpvgsm28T3t3 = map[int]defStructs.Champ{
	0: {"local", -1, -1, nil},
}

// Tdpvgsm28T3t4
var Tdpvgsm28T3t4 = map[int]defStructs.Champ{
	0: {"MSISDN", -1, -1, nil},
}

// Tdpvgsm28T3t5
var Tdpvgsm28T3t5 = map[int]defStructs.Champ{
	0: {"IMEI", -1, -1, nil},
}

// Tdpvgsm28T3t6
var Tdpvgsm28T3t6 = map[int]defStructs.Champ{
	0: {"portOrigNet", -1, -1, nil},
}

// Tdpvgsm28T3t7
var Tdpvgsm28T3t7 = map[int]defStructs.Champ{
	0: {"estalloc", -1, -1, nil},
}

// Tdpvgsm28T3t8
var Tdpvgsm28T3t8 = map[int]defStructs.Champ{
	0: {"msgITX", -1, -1, nil},
}

// Tdpvgsm28T3t9
var Tdpvgsm28T3t9 = map[int]defStructs.Champ{
	0: {"estTAX", 2, -1, nil},
	1: {"msgTAX", -1, -1, nil},
}

// Tdpvgsm28T3t10
var Tdpvgsm28T3t10 = map[int]defStructs.Champ{
	0: {"estCHT1", 2, -1, nil},
	1: {"msgCHT1", 4, -1, nil},
	2: {"estCHT2", 2, -1, nil},
	3: {"msgCHT2", 4, -1, nil},
	4: {"estCHT3", 2, -1, nil},
	5: {"msgCHT3", 4, -1, nil},
	6: {"estCHT4", 2, -1, nil},
	7: {"msgCHT4", -1, -1, nil},
}

// Tdpvgsm28T3t11_1
var Tdpvgsm28T3t11_1 = map[int]defStructs.Champ{
	0: {"nbsvcIN", -1, -1, nil},
}

// Tdpvgsm28T3t11_2
var Tdpvgsm28T3t11_2 = map[int]defStructs.Champ{
	0: {"lstSVCIN", -1, -1, nil},
}

// Tdpvgsm28T3t11_3
var Tdpvgsm28T3t11_3 = map[int]defStructs.Champ{
	0: {"SCPIN", -1, -1, nil},
}

// Tdpvgsm28T3t11
var Tdpvgsm28T3t11 = map[int]defStructs.Champ{
	0x01: {"T4t11.1", -1, -1, &Tdpvgsm28T3t11_1},
	0x02: {"T4t11.2", -1, -1, &Tdpvgsm28T3t11_2},
	0x03: {"T4t11.3", -1, -1, &Tdpvgsm28T3t11_3},
}

// Tdpvgsm28T3t12_1
var Tdpvgsm28T3t12_1 = map[int]defStructs.Champ{
	0: {"nbsvcIN", -1, -1, nil},
}

// Tdpvgsm28T3t12_2
var Tdpvgsm28T3t12_2 = map[int]defStructs.Champ{
	0: {"lstSVCIN", -1, -1, nil},
}

// Tdpvgsm28T3t12_3
var Tdpvgsm28T3t12_3 = map[int]defStructs.Champ{
	0: {"SCPPC", -1, -1, nil},
}

// Tdpvgsm28T3t12
var Tdpvgsm28T3t12 = map[int]defStructs.Champ{
	0x01: {"T4t12.1", -1, -1, &Tdpvgsm28T3t12_1},
	0x02: {"T4t12.2", -1, -1, &Tdpvgsm28T3t12_2},
	0x03: {"T4t12.3", -1, -1, &Tdpvgsm28T3t12_3},
}

// Tdpvgsm28T3t13_1
var Tdpvgsm28T3t13_1 = map[int]defStructs.Champ{
	0: {"priorite", 1, -1, nil},
	1: {"classe", 1, -1, nil},
	2: {"factdest", 1, -1, nil},
	3: {"nbdemMT", 1, -1, nil},
	4: {"nbstatMT", 1, -1, nil},
	5: {"nbrempMT", 1, -1, nil},
	6: {"reserve1", -1, -1, nil},
}

// Tdpvgsm28T3t13_2
var Tdpvgsm28T3t13_2 = map[int]defStructs.Champ{
	0: {"smid", -1, -1, nil},
}

// Tdpvgsm28T3t13_3
var Tdpvgsm28T3t13_3 = map[int]defStructs.Champ{
	0: {"svcPpMO", 1, -1, nil},
	1: {"classeMO", 1, -1, nil},
	2: {"factdeMO", 1, -1, nil},
	3: {"demacqMO", 1, -1, nil},
	4: {"recacqMO", 1, -1, nil},
	5: {"PrepaidMO", -1, -1, nil},
}

// Tdpvgsm28T3t13_4
var Tdpvgsm28T3t13_4 = map[int]defStructs.Champ{
	0: {"nbsuivi", 2, -1, nil},
	1: {"nbstat", 2, -1, nil},
	2: {"filler3", 1, -1, nil},
	3: {"filler4", -1, -1, nil},
}

// Tdpvgsm28T3t13_5
var Tdpvgsm28T3t13_5 = map[int]defStructs.Champ{
	0: {"subpart", 2, -1, nil},
	1: {"subdom", 2, -1, nil},
	2: {"chrgact", 1, -1, nil},
	3: {"nbfaxpag", -1, -1, nil},
}

// Tdpvgsm28T3t13
var Tdpvgsm28T3t13 = map[int]defStructs.Champ{
	0x07: {"T4t13.1", -1, -1, &Tdpvgsm28T3t13_1},
	0x08: {"T4t13.2", -1, -1, &Tdpvgsm28T3t13_2},
	0x09: {"T4t13.3", -1, -1, &Tdpvgsm28T3t13_3},
	0x0a: {"T4t13.3", -1, -1, &Tdpvgsm28T3t13_4},
	0x0b: {"T4t13.5", -1, -1, &Tdpvgsm28T3t13_5},
}

// Tdpvgsm28T3t14_1
var Tdpvgsm28T3t14_1 = map[int]defStructs.Champ{
	0: {"SSinvoc", -1, -1, nil},
}

// Tdpvgsm28T3t14_2
var Tdpvgsm28T3t14_2 = map[int]defStructs.Champ{
	0: {"svcSwitch", -1, -1, nil},
}

// Tdpvgsm28T3t14
var Tdpvgsm28T3t14 = map[int]defStructs.Champ{
	0x04: {"T4t14.1", -1, -1, &Tdpvgsm28T3t14_1},
	0x02: {"T4t14.2", -1, -1, &Tdpvgsm28T3t14_2},
}

// Tdpvgsm28T3t15_1
var Tdpvgsm28T3t15_1 = map[int]defStructs.Champ{
	0: {"callhold", -1, -1, nil},
}

// Tdpvgsm28T3t15_2
var Tdpvgsm28T3t15_2 = map[int]defStructs.Champ{
	0: {"NDS", -1, -1, nil},
}

// Tdpvgsm28T3t15
var Tdpvgsm28T3t15 = map[int]defStructs.Champ{
	0x10: {"T4t15.1", -1, -1, &Tdpvgsm28T3t15_1},
	0x11: {"T4t15.2", -1, -1, &Tdpvgsm28T3t15_2},
}

// Tdpvgsm28T3t16
var Tdpvgsm28T3t16 = map[int]defStructs.Champ{
	0: {"TrSortant", -1, -1, nil},
}

// Tdpvgsm28T3t17
var Tdpvgsm28T3t17 = map[int]defStructs.Champ{
	0: {"TrEntrant", -1, -1, nil},
}

// Tdpvgsm28T3t18_1
var Tdpvgsm28T3t18_1 = map[int]defStructs.Champ{
	0: {"C_CallRefNum", -1, -1, nil},
}

// Tdpvgsm28T3t18_2
var Tdpvgsm28T3t18_2 = map[int]defStructs.Champ{
	0: {"C_MSCAddress", -1, -1, nil},
}

// Tdpvgsm28T3t18_3
var Tdpvgsm28T3t18_3 = map[int]defStructs.Champ{
	0: {"C_SrvAddress1", -1, -1, nil},
}

// Tdpvgsm28T3t18_4
var Tdpvgsm28T3t18_4 = map[int]defStructs.Champ{
	0: {"C_DestNum1", -1, -1, nil},
}

// Tdpvgsm28T3t18_5
var Tdpvgsm28T3t18_5 = map[int]defStructs.Champ{
	0: {"C_Key1", 1, -1, nil},
	1: {"C_Level1", 1, -1, nil},
	2: {"C_InitCfInd1", 1, -1, nil},
	3: {"C_DefChInd1", 1, -1, nil},
	4: {"C_Modif1", -1, -1, nil},
}

// Tdpvgsm28T3t18_6
var Tdpvgsm28T3t18_6 = map[int]defStructs.Champ{
	0: {"C_CSEInfo1", -1, -1, nil},
}

// Tdpvgsm28T3t18_7
var Tdpvgsm28T3t18_7 = map[int]defStructs.Champ{
	0: {"C_SrvAddress2", -1, -1, nil},
}

// Tdpvgsm28T3t18_8
var Tdpvgsm28T3t18_8 = map[int]defStructs.Champ{
	0: {"C_DestNum2", -1, -1, nil},
}

// Tdpvgsm28T3t18_9
var Tdpvgsm28T3t18_9 = map[int]defStructs.Champ{
	0: {"C_Key2", 1, -1, nil},
	1: {"C_Level2", 1, -1, nil},
	2: {"C_InitCfInd2", 1, -1, nil},
	3: {"C_DefChInd2", 1, -1, nil},
	4: {"C_Modif2", -1, -1, nil},
}

// Tdpvgsm28T3t18_10
var Tdpvgsm28T3t18_10 = map[int]defStructs.Champ{
	0: {"C_CSEInfo2", -1, -1, nil},
}

// Tdpvgsm28T3t18
var Tdpvgsm28T3t18 = map[int]defStructs.Champ{
	0x01: {"T4t18.1", -1, -1, &Tdpvgsm28T3t18_1},
	0x02: {"T4t18.2", -1, -1, &Tdpvgsm28T3t18_2},
	0x03: {"T4t18.3", -1, -1, &Tdpvgsm28T3t18_3},
	0x04: {"T4t18.4", -1, -1, &Tdpvgsm28T3t18_4},
	0x05: {"T4t18.5", -1, -1, &Tdpvgsm28T3t18_5},
	0x06: {"T4t18.6", -1, -1, &Tdpvgsm28T3t18_6},
	0x07: {"T4t18.7", -1, -1, &Tdpvgsm28T3t18_7},
	0x08: {"T4t18.8", -1, -1, &Tdpvgsm28T3t18_8},
	0x09: {"T4t18.9", -1, -1, &Tdpvgsm28T3t18_9},
	0x0a: {"T4t18.10", -1, -1, &Tdpvgsm28T3t18_10},
}

// Tdpvgsm28T3t19
var Tdpvgsm28T3t19 = map[int]defStructs.Champ{
	0: {"PDPType", 1, -1, nil},
	1: {"PDPAddr", 4, -1, nil},
	2: {"accPointName3", 2, 1, nil},
}

// Tdpvgsm28T3t20
var Tdpvgsm28T3t20 = map[int]defStructs.Champ{
	0: {"remPDPAddr", -1, -1, nil},
}

// Tdpvgsm28T3t21_1
var Tdpvgsm28T3t21_1 = map[int]defStructs.Champ{
	0: {"qoSReqDelay", 1, -1, nil},
	1: {"qoSReqMean", 1, -1, nil},
	2: {"qoSReqPeak", 1, -1, nil},
	3: {"qoSReqPrec", 1, -1, nil},
	4: {"qoSReqReliab", -1, -1, nil},
}

// Tdpvgsm28T3t21_2
var Tdpvgsm28T3t21_2 = map[int]defStructs.Champ{
	0: {"qoS1Delay", 1, -1, nil},
	1: {"qoS1Mean", 1, -1, nil},
	2: {"qoS1Peak", 1, -1, nil},
	3: {"qoS1Prec", 1, -1, nil},
	4: {"qoS1Reliab", 1, -1, nil},
	5: {"qoS1VolOut", 6, -1, nil},
	6: {"qoS1VolIn", 6, -1, nil},
	7: {"qoS1TimeStamp", -1, -1, nil},
}

// Tdpvgsm28T3t21_3
var Tdpvgsm28T3t21_3 = map[int]defStructs.Champ{
	0: {"qoS2Delay", 1, -1, nil},
	1: {"qoS2Mean", 1, -1, nil},
	2: {"qoS2Peak", 1, -1, nil},
	3: {"qoS2Prec", 1, -1, nil},
	4: {"qoS2Reliab", 1, -1, nil},
	5: {"qoS2VolOut", 6, -1, nil},
	6: {"qoS2VolIn", 6, -1, nil},
	7: {"qoS2TimeStamp", -1, -1, nil},
}

// Tdpvgsm28T3t21_4
var Tdpvgsm28T3t21_4 = map[int]defStructs.Champ{
	0: {"qoS3Delay", 1, -1, nil},
	1: {"qoS3Mean", 1, -1, nil},
	2: {"qoS3Peak", 1, -1, nil},
	3: {"qoS3Prec", 1, -1, nil},
	4: {"qoS3Reliab", 1, -1, nil},
	5: {"qoS3VolOut", 6, -1, nil},
	6: {"qoS3VolIn", 6, -1, nil},
	7: {"qoS3TimeStamp", -1, -1, nil},
}

// Tdpvgsm28T3t21_5
var Tdpvgsm28T3t21_5 = map[int]defStructs.Champ{
	0: {"totVolOut", 4, 10, nil},
	1: {"totVolIn", -1, 10, nil},
}

// Tdpvgsm28T3t21_6
var Tdpvgsm28T3t21_6 = map[int]defStructs.Champ{
	0: {"chargingID", -1, -1, nil},
}

// Tdpvgsm28T3t21_7
var Tdpvgsm28T3t21_7 = map[int]defStructs.Champ{
	0: {"seqNum", -1, -1, nil},
}

// Tdpvgsm28T3t21_8
var Tdpvgsm28T3t21_8 = map[int]defStructs.Champ{
	0: {"PartialType", -1, -1, nil},
}

// Tdpvgsm28T3t21_9
var Tdpvgsm28T3t21_9 = map[int]defStructs.Champ{
	0: {"RecEntType1", 1, -1, nil},
	1: {"RecEntId1", 4, -1, nil},
	2: {"RecEntType2", 1, -1, nil},
	3: {"RecEntId2", 4, -1, nil},
	4: {"RecEntType3", 1, -1, nil},
	5: {"RecEntId3", -1, -1, nil},
}

// Tdpvgsm28T3t21
var Tdpvgsm28T3t21 = map[int]defStructs.Champ{
	0x01: {"T4t21.1", -1, -1, &Tdpvgsm28T3t21_1},
	0x02: {"T4t21.2", -1, -1, &Tdpvgsm28T3t21_2},
	0x03: {"T4t21.3", -1, -1, &Tdpvgsm28T3t21_3},
	0x04: {"T4t21.4", -1, -1, &Tdpvgsm28T3t21_4},
	0x05: {"T4t21.5", -1, -1, &Tdpvgsm28T3t21_5},
	0x06: {"T4t21.6", -1, -1, &Tdpvgsm28T3t21_6},
	0x07: {"T4t21.7", -1, -1, &Tdpvgsm28T3t21_7},
	0x08: {"T4t21.8", -1, -1, &Tdpvgsm28T3t21_8},
	0x09: {"T4t21.9", -1, -1, &Tdpvgsm28T3t21_9},
}

// Tdpvgsm28T3t30_1
var Tdpvgsm28T3t30_1 = map[int]defStructs.Champ{
	0: {"timeInCE", -1, -1, nil},
}

// Tdpvgsm28T3t30_2
var Tdpvgsm28T3t30_2 = map[int]defStructs.Champ{
	0: {"eec", -1, -1, nil},
}

// Tdpvgsm28T3t30
var Tdpvgsm28T3t30 = map[int]defStructs.Champ{
	0x20: {"T4t30.1", -1, -1, &Tdpvgsm28T3t30_1},
	0x21: {"T4t30.2", -1, -1, &Tdpvgsm28T3t30_2},
}

// Tdpvgsm28T3
var Tdpvgsm28T3 = map[int]defStructs.Champ{
	0x11: {"T3t1", -1, -1, &Tdpvgsm28T3t1},
	0x12: {"T3t2", -1, -1, &Tdpvgsm28T3t2},
	0x13: {"T3t3", -1, -1, &Tdpvgsm28T3t3},
	0x14: {"T3t4", -1, -1, &Tdpvgsm28T3t4},
	0x15: {"T3t5", -1, -1, &Tdpvgsm28T3t5},
	0x16: {"T3t6", -1, -1, &Tdpvgsm28T3t6},
	0x17: {"T3t7", -1, -1, &Tdpvgsm28T3t7},
	0x18: {"T3t8", -1, -1, &Tdpvgsm28T3t8},
	0x19: {"T3t9", -1, -1, &Tdpvgsm28T3t9},
	0x1a: {"T3t10", -1, -1, &Tdpvgsm28T3t10},
	0x1b: {"T3t11", -1, -1, &Tdpvgsm28T3t11},
	0x1c: {"T3t12", -1, -1, &Tdpvgsm28T3t12},
	0x1d: {"T3t13", -1, -1, &Tdpvgsm28T3t13},
	0x1e: {"T3t14", -1, -1, &Tdpvgsm28T3t14},
	0x1f: {"T3t15", -1, -1, &Tdpvgsm28T3t15},
	0x20: {"T3t16", -1, -1, &Tdpvgsm28T3t16},
	0x21: {"T3t17", -1, -1, &Tdpvgsm28T3t17},
	0x22: {"T3t18", -1, -1, &Tdpvgsm28T3t18},
	0x23: {"T3t19", -1, -1, &Tdpvgsm28T3t19},
	0x24: {"T3t20", -1, -1, &Tdpvgsm28T3t20},
	0x25: {"T3t21", -1, -1, &Tdpvgsm28T3t21},
	0x30: {"T3t30", -1, -1, &Tdpvgsm28T3t30},
}

// Tdpvgsm28T2
var Tdpvgsm28T2 = map[int]defStructs.Champ{
	0x02: {"T2", -1, -1, &Tdpvgsm28T3},
}

// Tdpvgsm28
var Tdpvgsm28 = map[int]defStructs.Champ{
	0x11: {"Tdpvgsm28", -1, -1, &Tdpvgsm28T2},
}

// Ddpvgsm28
var Ddpvgsm28 = map[int]defStructs.TabRechercheTypeDc{
	0: {"Tdpvgsm28", 0xffff0000, 0x80110000, 0, &Tdpvgsm28},
}
