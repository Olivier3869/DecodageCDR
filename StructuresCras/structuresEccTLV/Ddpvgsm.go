package structuresEccTLV

import "local/structures/defStructs"

// TdpvgsmT3t1
var TdpvgsmT3t1 = map[int]defStructs.Champ{
	0:  {"type", 1, -1, nil},
	1:  {"tentite", 1, -1, nil},
	2:  {"idemet", 9, -1, nil},
	3:  {"IMSI", 8, -1, nil},
	4:  {"DATE", 7, 8, nil},
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

// TdpvgsmT3t2
var TdpvgsmT3t2 = map[int]defStructs.Champ{
	0: {"idpart", -1, -1, nil},
}

// TdpvgsmT3t3
var TdpvgsmT3t3 = map[int]defStructs.Champ{
	0: {"local", -1, -1, nil},
}

// TdpvgsmT3t4
var TdpvgsmT3t4 = map[int]defStructs.Champ{
	0: {"MSISDN", -1, -1, nil},
}

// TdpvgsmT3t5
var TdpvgsmT3t5 = map[int]defStructs.Champ{
	0: {"IMEI", -1, -1, nil},
}

// TdpvgsmT3t6
var TdpvgsmT3t6 = map[int]defStructs.Champ{
	0: {"portOrigNet", -1, -1, nil},
}

// TdpvgsmT3t7
var TdpvgsmT3t7 = map[int]defStructs.Champ{
	0: {"estalloc", -1, 8, nil},
}

// TdpvgsmT3t8
var TdpvgsmT3t8 = map[int]defStructs.Champ{
	0: {"msgITX", -1, -1, nil},
}

// TdpvgsmT3t11_1
var TdpvgsmT3t11_1 = map[int]defStructs.Champ{
	0: {"nbsvcIN", -1, -1, nil},
}

// TdpvgsmT3t11_2
var TdpvgsmT3t11_2 = map[int]defStructs.Champ{
	0: {"lstSVCIN", -1, -1, nil},
}

// TdpvgsmT3t11_3
var TdpvgsmT3t11_3 = map[int]defStructs.Champ{
	0: {"SCPIN", -1, -1, nil},
}

// TdpvgsmT3t11
var TdpvgsmT3t11 = map[int]defStructs.Champ{
	0x01: {"T4t11.1", -1, -1, &TdpvgsmT3t11_1},
	0x02: {"T4t11.2", -1, -1, &TdpvgsmT3t11_2},
	0x03: {"T4t11.3", -1, -1, &TdpvgsmT3t11_3},
}

// TdpvgsmT3t12
var TdpvgsmT3t12 = map[int]defStructs.Champ{
	0: {"portpref", -1, -1, nil},
}

// TdpvgsmT3t13_1
var TdpvgsmT3t13_1 = map[int]defStructs.Champ{
	0: {"priorite", 1, -1, nil},
	1: {"classe", 1, -1, nil},
	2: {"factdest", 1, -1, nil},
	3: {"nbdemMT", 1, -1, nil},
	4: {"nbstatMT", -1, -1, nil},
	//5: {"nbrempMT", -1 , 1, "NULL", 5},
	//6: {"reserve1", -1 , -1, "NULL", 5},
}

// TdpvgsmT3t13_2
var TdpvgsmT3t13_2 = map[int]defStructs.Champ{
	0: {"smid", -1, -1, nil},
}

// TdpvgsmT3t13_3
var TdpvgsmT3t13_3 = map[int]defStructs.Champ{
	0: {"svcPpMO", 1, -1, nil},
	1: {"classeMO", 1, -1, nil},
	2: {"factdeMO", 1, -1, nil},
	3: {"demacqMO", 1, -1, nil},
	4: {"recacqMO", 1, -1, nil},
	5: {"PrepaidMO", -1, -1, nil},
}

// TdpvgsmT3t13_5
var TdpvgsmT3t13_5 = map[int]defStructs.Champ{
	0: {"subpart", 2, -1, nil},
	1: {"subdom", 2, -1, nil},
	2: {"chrgact", 1, -1, nil},
	3: {"nbfaxpag", -1, -1, nil},
}

// TdpvgsmT3t13
var TdpvgsmT3t13 = map[int]defStructs.Champ{
	0x07: {"T4t13.1", -1, -1, &TdpvgsmT3t13_1},
	0x08: {"T4t13.2", -1, -1, &TdpvgsmT3t13_2},
	0x09: {"T4t13.3", -1, -1, &TdpvgsmT3t13_3},
	0x0b: {"T4t13.5", -1, -1, &TdpvgsmT3t13_5},
}

// TdpvgsmT3t15_1
var TdpvgsmT3t15_1 = map[int]defStructs.Champ{
	0: {"callhold", -1, -1, nil},
}

// TdpvgsmT3t15_2
var TdpvgsmT3t15_2 = map[int]defStructs.Champ{
	0: {"NDS", -1, -1, nil},
}

// TdpvgsmT3t15
var TdpvgsmT3t15 = map[int]defStructs.Champ{
	0x10: {"T4t15.1", -1, -1, &TdpvgsmT3t15_1},
	0x11: {"T4t15.2", -1, -1, &TdpvgsmT3t15_2},
}

// TdpvgsmT3t16
var TdpvgsmT3t16 = map[int]defStructs.Champ{
	0: {"TrSortant", -1, 1, nil},
}

// TdpvgsmT3t17
var TdpvgsmT3t17 = map[int]defStructs.Champ{
	0: {"TrEntrant", -1, 1, nil},
}

// TdpvgsmT3t18_1
var TdpvgsmT3t18_1 = map[int]defStructs.Champ{
	0: {"C_CallRefNum", -1, -1, nil},
}

// TdpvgsmT3t18_2
var TdpvgsmT3t18_2 = map[int]defStructs.Champ{
	0: {"C_MSCAddress", -1, -1, nil},
}

// TdpvgsmT3t18_3
var TdpvgsmT3t18_3 = map[int]defStructs.Champ{
	0: {"C_SrvAddress1", -1, -1, nil},
}

// TdpvgsmT3t18_4
var TdpvgsmT3t18_4 = map[int]defStructs.Champ{
	0: {"C_DestNum1", -1, -1, nil},
}

// TdpvgsmT3t18_5
var TdpvgsmT3t18_5 = map[int]defStructs.Champ{
	0: {"C_Key1", 1, -1, nil},
	1: {"C_Level1", 1, -1, nil},
	2: {"C_InitCfInd1", 1, -1, nil},
	3: {"C_DefChInd1", 1, -1, nil},
	4: {"C_Modif1", -1, -1, nil},
}

// TdpvgsmT3t18_6
var TdpvgsmT3t18_6 = map[int]defStructs.Champ{
	0: {"C_CSEInfo1", -1, -1, nil},
}

// TdpvgsmT3t18_7
var TdpvgsmT3t18_7 = map[int]defStructs.Champ{
	0: {"C_AccPtNameNI", -1, -1, nil},
}

// TdpvgsmT3t18_8
var TdpvgsmT3t18_8 = map[int]defStructs.Champ{
	0: {"C_AccPtNameOI", -1, -1, nil},
}

// TdpvgsmT3t18_9
var TdpvgsmT3t18_9 = map[int]defStructs.Champ{
	0: {"C_SrvAddress2", -1, -1, nil},
}

// TdpvgsmT3t18_10
var TdpvgsmT3t18_10 = map[int]defStructs.Champ{
	0: {"C_DestNum2", -1, -1, nil},
}

// TdpvgsmT3t18_11
var TdpvgsmT3t18_11 = map[int]defStructs.Champ{
	0: {"C_Key2", 1, -1, nil},
	1: {"C_Level2", 1, -1, nil},
	2: {"C_InitCfInd2", 1, -1, nil},
	3: {"C_DefChInd2", 1, -1, nil},
	4: {"C_Modif2", -1, -1, nil},
}

// TdpvgsmT3t18_12
var TdpvgsmT3t18_12 = map[int]defStructs.Champ{
	0: {"C_CSEInfo2", -1, -1, nil},
}

// TdpvgsmT3t18_13
var TdpvgsmT3t18_13 = map[int]defStructs.Champ{
	0: {"accPointName1", -1, 1, nil},
}

// TdpvgsmT3t18_14
var TdpvgsmT3t18_14 = map[int]defStructs.Champ{
	0: {"accPointName2", -1, 1, nil},
}

// TdpvgsmT3t18
var TdpvgsmT3t18 = map[int]defStructs.Champ{
	0x01: {"T4t18.1", -1, -1, &TdpvgsmT3t18_1},
	0x02: {"T4t18.2", -1, -1, &TdpvgsmT3t18_2},
	0x03: {"T4t18.3", -1, -1, &TdpvgsmT3t18_3},
	0x04: {"T4t18.4", -1, -1, &TdpvgsmT3t18_4},
	0x05: {"T4t18.5", -1, -1, &TdpvgsmT3t18_5},
	0x06: {"T4t18.6", -1, -1, &TdpvgsmT3t18_6},
	0x07: {"T4t18.7", -1, -1, &TdpvgsmT3t18_7},
	0x08: {"T4t18.8", -1, -1, &TdpvgsmT3t18_8},
	0x09: {"T4t18.9", -1, -1, &TdpvgsmT3t18_9},
	0x0a: {"T4t18.10", -1, -1, &TdpvgsmT3t18_10},
	0x0b: {"T4t18.11", -1, -1, &TdpvgsmT3t18_11},
	0x0c: {"T4t18.12", -1, -1, &TdpvgsmT3t18_12},
	0x0d: {"T4t18.13", -1, -1, &TdpvgsmT3t18_13},
	0x0e: {"T4t18.14", -1, -1, &TdpvgsmT3t18_14},
}

// TdpvgsmT3t20_1
var TdpvgsmT3t20_1 = map[int]defStructs.Champ{
	0: {"PDPaddr", -1, -1, nil},
}

// TdpvgsmT3t20_2
var TdpvgsmT3t20_2 = map[int]defStructs.Champ{
	0: {"accPointName3", -1, 1, nil},
}

// TdpvgsmT3t20_3
var TdpvgsmT3t20_3 = map[int]defStructs.Champ{
	0: {"accPointName4", -1, 1, nil},
}

// TdpvgsmT3t20_4
var TdpvgsmT3t20_4 = map[int]defStructs.Champ{
	0: {"chargingID", -1, -1, nil},
}

// TdpvgsmT3t20_5
var TdpvgsmT3t20_5 = map[int]defStructs.Champ{
	0: {"SeqNum", -1, -1, nil},
}

// TdpvgsmT3t20_6
var TdpvgsmT3t20_6 = map[int]defStructs.Champ{
	0: {"PartialType", -1, -1, nil},
}

// TdpvgsmT3t20_7
var TdpvgsmT3t20_7 = map[int]defStructs.Champ{
	0: {"RecEntId1", -1, -1, nil},
}

// TdpvgsmT3t20_8
var TdpvgsmT3t20_8 = map[int]defStructs.Champ{
	0: {"RecEntId2", -1, -1, nil},
}

// TdpvgsmT3t20_9
var TdpvgsmT3t20_9 = map[int]defStructs.Champ{
	0: {"RecEntId3", -1, -1, nil},
}

// TdpvgsmT3t20_10
var TdpvgsmT3t20_10 = map[int]defStructs.Champ{
	0: {"totVolOut", 6, 10, nil},
	1: {"totVolIn", -1, 10, nil},
}

// TdpvgsmT3t20
var TdpvgsmT3t20 = map[int]defStructs.Champ{
	0x01: {"T4t20.1", -1, -1, &TdpvgsmT3t20_1},
	0x02: {"T4t20.2", -1, -1, &TdpvgsmT3t20_2},
	0x03: {"T4t20.3", -1, -1, &TdpvgsmT3t20_3},
	0x04: {"T4t20.4", -1, -1, &TdpvgsmT3t20_4},
	0x05: {"T4t20.5", -1, -1, &TdpvgsmT3t20_5},
	0x06: {"T4t20.6", -1, -1, &TdpvgsmT3t20_6},
	0x07: {"T4t20.7", -1, -1, &TdpvgsmT3t20_7},
	0x08: {"T4t20.8", -1, -1, &TdpvgsmT3t20_8},
	0x09: {"T4t20.9", -1, -1, &TdpvgsmT3t20_9},
	0x0a: {"T4t20.10", -1, -1, &TdpvgsmT3t20_10},
}

// TdpvgsmT3t21_1
var TdpvgsmT3t21_1 = map[int]defStructs.Champ{
	0: {"qoSReqDelay", 1, -1, nil},
	1: {"qoSReqMean", 1, -1, nil},
	2: {"qoSReqPeak", 1, -1, nil},
	3: {"qoSReqPrec", 1, -1, nil},
	4: {"qoSReqReliab", -1, -1, nil},
}

// TdpvgsmT3t21_2
var TdpvgsmT3t21_2 = map[int]defStructs.Champ{
	0: {"qoS1Delay", 1, -1, nil},
	1: {"qoS1Mean", 1, -1, nil},
	2: {"qoS1Peak", 1, -1, nil},
	3: {"qoS1Prec", 1, -1, nil},
	4: {"qoS1Reliab", 1, -1, nil},
	5: {"qoS1VolOut", 6, -1, nil},
	6: {"qoS1VolIn", 6, -1, nil},
	7: {"qoS1TimeStamp", -1, -1, nil},
}

// TdpvgsmT3t21_3
var TdpvgsmT3t21_3 = map[int]defStructs.Champ{
	0: {"qoS2Delay", 1, -1, nil},
	1: {"qoS2Mean", 1, -1, nil},
	2: {"qoS2Peak", 1, -1, nil},
	3: {"qoS2Prec", 1, -1, nil},
	4: {"qoS2Reliab", 1, -1, nil},
	5: {"qoS2VolOut", 6, -1, nil},
	6: {"qoS2VolIn", 6, -1, nil},
	7: {"qoS2TimeStamp", -1, -1, nil},
}

// TdpvgsmT3t21_4
var TdpvgsmT3t21_4 = map[int]defStructs.Champ{
	0: {"qoS3Delay", 1, -1, nil},
	1: {"qoS3Mean", 1, -1, nil},
	2: {"qoS3Peak", 1, -1, nil},
	3: {"qoS3Prec", 1, -1, nil},
	4: {"qoS3Reliab", 1, -1, nil},
	5: {"qoS3VolOut", 6, -1, nil},
	6: {"qoS3VolIn", 6, -1, nil},
	7: {"qoS3TimeStamp", -1, -1, nil},
}

// TdpvgsmT3t21
var TdpvgsmT3t21 = map[int]defStructs.Champ{
	0x01: {"T4t21.1", -1, -1, &TdpvgsmT3t21_1},
	0x02: {"T4t21.2", -1, -1, &TdpvgsmT3t21_2},
	0x03: {"T4t21.3", -1, -1, &TdpvgsmT3t21_3},
	0x04: {"T4t21.4", -1, -1, &TdpvgsmT3t21_4},
}

// TdpvgsmT3t22_1
var TdpvgsmT3t22_1 = map[int]defStructs.Champ{
	0:  {"qoSAllocRetPri", 1, -1, nil},
	1:  {"qoSTrafficClas", 1, -1, nil},
	2:  {"qoSDelivOrd", 1, -1, nil},
	3:  {"qoSDelivErr", 1, -1, nil},
	4:  {"qoSMaxSize", 1, -1, nil},
	5:  {"qoSMaxSRateUp", 1, -1, nil},
	6:  {"qoSMaxSRateDow", 1, -1, nil},
	7:  {"qoSResBER", 1, -1, nil},
	8:  {"qoSSDUErrRatio", 1, -1, nil},
	9:  {"qoSTransDelay", 1, -1, nil},
	10: {"qoSTrafficHand", -1, -1, nil},
	//11: {"qoSBitRateUp", -1, 1, "NULL", 5},
	//12: {"qoSBitRateDown", -1, -1, "NULL", 5},
}

// TdpvgsmT3t22_2
var TdpvgsmT3t22_2 = map[int]defStructs.Champ{
	0:  {"qoS1AllocRetPri", 1, -1, nil},
	1:  {"qoS1TrafficClas", 1, -1, nil},
	2:  {"qoS1DelivOrd", 1, -1, nil},
	3:  {"qoS1DelivErr", 1, -1, nil},
	4:  {"qoS1MaxSize", 1, -1, nil},
	5:  {"qoS1MaxSRateUp", 1, -1, nil},
	6:  {"qoS1MaxSRateDow", 1, -1, nil},
	7:  {"qoS1ResBER", 1, -1, nil},
	8:  {"qoS1SDUErrRati", 1, -1, nil},
	9:  {"qoS1TransDelay", 1, -1, nil},
	10: {"qoS1TrafficHan", 1, -1, nil},
	11: {"qoS1GuarRateUp", 1, -1, nil},
	12: {"qoS1GuarRateDo", 1, -1, nil},
	13: {"qoS1VOLOut", 6, -1, nil},
	14: {"qoS1VOLInqoS1VOLIn", 6, -1, nil},
	15: {"qoS1UMTSTime", -1, -1, nil},
}

// TdpvgsmT3t22_3
var TdpvgsmT3t22_3 = map[int]defStructs.Champ{
	0:  {"qoS2AllocRetPri", 1, -1, nil},
	1:  {"qoS2TrafficClas", 1, -1, nil},
	2:  {"qoS2DelivOrd", 1, -1, nil},
	3:  {"qoS2DelivErr", 1, -1, nil},
	4:  {"qoS2MaxSize", 1, -1, nil},
	5:  {"qoS21MaxSRateUp", 1, -1, nil},
	6:  {"qoS2MaxSRateDow", 1, -1, nil},
	7:  {"qoS2ResBER", 1, -1, nil},
	8:  {"qoS2SDUErrRati", 1, -1, nil},
	9:  {"qoS2TransDelay", 1, -1, nil},
	10: {"qoS2TrafficHan", 1, -1, nil},
	11: {"qoS2GuarRateUp", 1, -1, nil},
	12: {"qoS2GuarRateDo", 1, -1, nil},
	13: {"qoS2VOLOut", 6, -1, nil},
	14: {"qoS2VOLInqoS1VOLIn", 6, -1, nil},
	15: {"qoS2UMTSTime", -1, -1, nil},
}

// TdpvgsmT3t22_4
var TdpvgsmT3t22_4 = map[int]defStructs.Champ{
	0:  {"qoS3AllocRetPri", 1, -1, nil},
	1:  {"qoS3TrafficClas", 1, -1, nil},
	2:  {"qoS3DelivOrd", 1, -1, nil},
	3:  {"qoS3DelivErr", 1, -1, nil},
	4:  {"qoS3MaxSize", 1, -1, nil},
	5:  {"qoS31MaxSRateUp", 1, -1, nil},
	6:  {"qoS3MaxSRateDow", 1, -1, nil},
	7:  {"qoS3ResBER", 1, -1, nil},
	8:  {"qoS3SDUErrRati", 1, -1, nil},
	9:  {"qoS3TransDelay", 1, -1, nil},
	10: {"qoS3TrafficHan", 1, -1, nil},
	11: {"qoS3GuarRateUp", 1, -1, nil},
	12: {"qoS3GuarRateDo", 1, -1, nil},
	13: {"qoS3VOLOut", 6, -1, nil},
	14: {"qoS3VOLInqoS1VOLIn", 6, -1, nil},
	15: {"qoS3UMTSTime", -1, -1, nil},
}

// TdpvgsmT3t22
var TdpvgsmT3t22 = map[int]defStructs.Champ{
	0x01: {"T4t22.1", -1, -1, &TdpvgsmT3t22_1},
	0x02: {"T4t22.2", -1, -1, &TdpvgsmT3t22_2},
	0x03: {"T4t22.3", -1, -1, &TdpvgsmT3t22_3},
	0x04: {"T4t22.4", -1, -1, &TdpvgsmT3t22_4},
}

// TdpvgsmT3t23
var TdpvgsmT3t23 = map[int]defStructs.Champ{
	0: {"MailIdpart", -1, -1, nil},
}

// TdpvgsmT3t24
var TdpvgsmT3t24 = map[int]defStructs.Champ{
	0: {"MailMsOrEntity", -1, -1, nil},
}

// TdpvgsmT3t25
var TdpvgsmT3t25 = map[int]defStructs.Champ{
	0: {"MMSsubscribId2", -1, -1, nil},
}

// TdpvgsmT3t26
var TdpvgsmT3t26 = map[int]defStructs.Champ{
	0: {"MMStypeOfCont", -1, -1, nil},
}

// TdpvgsmT3t27
var TdpvgsmT3t27 = map[int]defStructs.Champ{
	0: {"MMSoutInterf", -1, 1, nil},
}

// TdpvgsmT3t28
var TdpvgsmT3t28 = map[int]defStructs.Champ{
	0: {"MMSinInterf", -1, 1, nil},
}

// TdpvgsmT3t29
var TdpvgsmT3t29 = map[int]defStructs.Champ{
	0: {"WlanHotSpotId", -1, -1, nil},
}

// TdpvgsmT3t30_1
var TdpvgsmT3t30_1 = map[int]defStructs.Champ{
	0: {"timeInCE", -1, 8, nil},
}

// TdpvgsmT3t30_2
var TdpvgsmT3t30_2 = map[int]defStructs.Champ{
	0: {"eec", -1, 10, nil},
}

// TdpvgsmT3t30
var TdpvgsmT3t30 = map[int]defStructs.Champ{
	0x20: {"T4t30.1", -1, -1, &TdpvgsmT3t30_1},
	0x21: {"T4t30.2", -1, -1, &TdpvgsmT3t30_2},
}

// TdpvgsmT3t31
var TdpvgsmT3t31 = map[int]defStructs.Champ{
	0: {"WlanSubscribId", -1, -1, nil},
}

// TdpvgsmT3t32
var TdpvgsmT3t32 = map[int]defStructs.Champ{
	0: {"WlanAccPointId", -1, -1, nil},
}

// TdpvgsmT3t33
var TdpvgsmT3t33 = map[int]defStructs.Champ{
	0: {"SMSrecepMSISDN", -1, -1, nil},
}

// TdpvgsmT3t34
var TdpvgsmT3t34 = map[int]defStructs.Champ{
	0: {"ZBprefix", -1, -1, nil},
}

// TdpvgsmT3t35_1
var TdpvgsmT3t35_1 = map[int]defStructs.Champ{
	0: {"duration", -1, -1, nil},
}

// TdpvgsmT3t35_2
var TdpvgsmT3t35_2 = map[int]defStructs.Champ{
	0: {"occurenceNb", -1, -1, nil},
}

// TdpvgsmT3t35_3
var TdpvgsmT3t35_3 = map[int]defStructs.Champ{
	0: {"participantsNb", -1, -1, nil},
}

// TdpvgsmT3t35_4
var TdpvgsmT3t35_4 = map[int]defStructs.Champ{
	0: {"presenceInd", -1, -1, nil},
}

// TdpvgsmT3t35
var TdpvgsmT3t35 = map[int]defStructs.Champ{
	0x01: {"T4t35.1", -1, -1, &TdpvgsmT3t35_1},
	0x02: {"T4t35.2", -1, -1, &TdpvgsmT3t35_2},
	0x03: {"T4t35.3", -1, -1, &TdpvgsmT3t35_3},
	0x04: {"T4t35.4", -1, -1, &TdpvgsmT3t35_4},
}

// TdpvgsmT3t36
var TdpvgsmT3t36 = map[int]defStructs.Champ{
	0: {"secondaryMSISDN", -1, -1, nil},
}

// TdpvgsmT3t37
var TdpvgsmT3t37 = map[int]defStructs.Champ{
	0: {"CSDaddInfo", -1, -1, nil},
}

// TdpvgsmT3t38_1
var TdpvgsmT3t38_1 = map[int]defStructs.Champ{
	0: {"origCallpartId", -1, 1, nil},
}

// TdpvgsmT3t38_2
var TdpvgsmT3t38_2 = map[int]defStructs.Champ{
	0: {"dialledDigits", -1, 1, nil},
}

// TdpvgsmT3t38_3
var TdpvgsmT3t38_3 = map[int]defStructs.Champ{
	0: {"equipmentType", -1, -1, nil},
}

// TdpvgsmT3t38_4
var TdpvgsmT3t38_4 = map[int]defStructs.Champ{
	0: {"originSystype", -1, -1, nil},
}

// TdpvgsmT3t38_5
var TdpvgsmT3t38_5 = map[int]defStructs.Champ{
	0: {"originBearSrv", -1, -1, nil},
}

// TdpvgsmT3t38
var TdpvgsmT3t38 = map[int]defStructs.Champ{
	0x01: {"T4t38.1", -1, -1, &TdpvgsmT3t38_1},
	0x02: {"T4t38.2", -1, -1, &TdpvgsmT3t38_2},
	0x03: {"T4t38.3", -1, -1, &TdpvgsmT3t38_3},
	0x04: {"T4t38.4", -1, -1, &TdpvgsmT3t38_4},
	0x05: {"T4t38.5", -1, -1, &TdpvgsmT3t38_5},
}

// TdpvgsmT3t39_1
var TdpvgsmT3t39_1 = map[int]defStructs.Champ{
	0: {"currencyCode", -1, -1, nil},
}

// TdpvgsmT3t39_2
var TdpvgsmT3t39_2 = map[int]defStructs.Champ{
	0: {"contVATCode", -1, -1, nil},
}

// TdpvgsmT3t39_3
var TdpvgsmT3t39_3 = map[int]defStructs.Champ{
	0: {"refundFlag", -1, -1, nil},
}

// TdpvgsmT3t39_4
var TdpvgsmT3t39_4 = map[int]defStructs.Champ{
	0: {"balanceAmount", -1, -1, nil},
}

// TdpvgsmT3t39_5
var TdpvgsmT3t39_5 = map[int]defStructs.Champ{
	0: {"balanceUnits", -1, -1, nil},
}

// TdpvgsmT3t39_6
var TdpvgsmT3t39_6 = map[int]defStructs.Champ{
	0: {"Amount", -1, -1, nil},
}

// TdpvgsmT3t39_7
var TdpvgsmT3t39_7 = map[int]defStructs.Champ{
	0: {"delivUnits", -1, -1, nil},
}

// TdpvgsmT3t39_8
var TdpvgsmT3t39_8 = map[int]defStructs.Champ{
	0: {"pricingUTtype", -1, -1, nil},
}

// TdpvgsmT3t39_9
var TdpvgsmT3t39_9 = map[int]defStructs.Champ{
	0: {"flagPitOut", -1, -1, nil},
}

// TdpvgsmT3t39
var TdpvgsmT3t39 = map[int]defStructs.Champ{
	0x01: {"T4t39.1", -1, -1, &TdpvgsmT3t39_1},
	0x02: {"T4t39.2", -1, -1, &TdpvgsmT3t39_2},
	0x03: {"T4t39.3", -1, -1, &TdpvgsmT3t39_3},
	0x04: {"T4t39.4", -1, -1, &TdpvgsmT3t39_4},
	0x05: {"T4t39.5", -1, -1, &TdpvgsmT3t39_5},
	0x06: {"T4t39.6", -1, -1, &TdpvgsmT3t39_6},
	0x07: {"T4t39.7", -1, -1, &TdpvgsmT3t39_7},
	0x08: {"T4t39.8", -1, -1, &TdpvgsmT3t39_8},
	0x09: {"T4t39.9", -1, -1, &TdpvgsmT3t39_9},
}

// TdpvgsmT3t40_1
var TdpvgsmT3t40_1 = map[int]defStructs.Champ{
	0: {"userIdNet", -1, -1, nil},
}

// TdpvgsmT3t40_2
var TdpvgsmT3t40_2 = map[int]defStructs.Champ{
	0: {"CBusageInd", -1, -1, nil},
}

// TdpvgsmT3t40
var TdpvgsmT3t40 = map[int]defStructs.Champ{
	0x01: {"T4t40.1", -1, -1, &TdpvgsmT3t40_1},
	0x02: {"T4t40.2", -1, -1, &TdpvgsmT3t40_2},
}

// TdpvgsmT3t41_1
var TdpvgsmT3t41_1 = map[int]defStructs.Champ{
	0: {"specifAttachVol", -1, -1, nil},
}

// TdpvgsmT3t41_2
var TdpvgsmT3t41_2 = map[int]defStructs.Champ{
	0: {"AttachNb", -1, -1, nil},
}

// TdpvgsmT3t41_3
var TdpvgsmT3t41_3 = map[int]defStructs.Champ{
	0: {"AttachFlag", -1, -1, nil},
}

// TdpvgsmT3t41_4
var TdpvgsmT3t41_4 = map[int]defStructs.Champ{
	0: {"CBserviceId", -1, -1, nil},
}

// TdpvgsmT3t41_5
var TdpvgsmT3t41_5 = map[int]defStructs.Champ{
	0: {"serviceID", -1, -1, nil},
}

// TdpvgsmT3t41_6
var TdpvgsmT3t41_6 = map[int]defStructs.Champ{
	0: {"specifServVol", -1, -1, nil},
}

// TdpvgsmT3t41
var TdpvgsmT3t41 = map[int]defStructs.Champ{
	0x01: {"T4t41.1", -1, -1, &TdpvgsmT3t41_1},
	0x02: {"T4t41.2", -1, -1, &TdpvgsmT3t41_2},
	0x03: {"T4t41.3", -1, -1, &TdpvgsmT3t41_3},
	0x04: {"T4t41.4", -1, -1, &TdpvgsmT3t41_4},
	0x05: {"T4t41.5", -1, -1, &TdpvgsmT3t41_5},
	0x06: {"T4t41.6", -1, -1, &TdpvgsmT3t41_6},
}

// TdpvgsmT3t42
var TdpvgsmT3t42 = map[int]defStructs.Champ{
	0: {"PDPType", -1, -1, nil},
}

// TdpvgsmT3t43_1
var TdpvgsmT3t43_1 = map[int]defStructs.Champ{
	0: {"salesUnit", -1, -1, nil},
}

// TdpvgsmT3t43_2
var TdpvgsmT3t43_2 = map[int]defStructs.Champ{
	0: {"refServID", -1, -1, nil},
}

// TdpvgsmT3t43_3
var TdpvgsmT3t43_3 = map[int]defStructs.Champ{
	0: {"transactionDesc", -1, 1, nil},
}

// TdpvgsmT3t43
var TdpvgsmT3t43 = map[int]defStructs.Champ{
	0x01: {"T4t43.1", -1, -1, &TdpvgsmT3t43_1},
	0x02: {"T4t43.2", -1, -1, &TdpvgsmT3t43_2},
	0x03: {"T4t43.3", -1, -1, &TdpvgsmT3t43_3},
}

// TdpvgsmT3t44_1
var TdpvgsmT3t44_1 = map[int]defStructs.Champ{
	0: {"transactionID", -1, -1, nil},
}

// TdpvgsmT3t44_2
var TdpvgsmT3t44_2 = map[int]defStructs.Champ{
	0: {"transactionType", -1, -1, nil},
}

// TdpvgsmT3t44
var TdpvgsmT3t44 = map[int]defStructs.Champ{
	0x01: {"T4t44.1", -1, -1, &TdpvgsmT3t44_1},
	0x02: {"T4t44.2", -1, -1, &TdpvgsmT3t44_2},
}

// TdpvgsmT3t45_1
var TdpvgsmT3t45_1 = map[int]defStructs.Champ{
	0: {"providerID", -1, -1, nil},
}

// TdpvgsmT3t45_2
var TdpvgsmT3t45_2 = map[int]defStructs.Champ{
	0: {"kitProviderID", -1, -1, nil},
}

// TdpvgsmT3t45_3
var TdpvgsmT3t45_3 = map[int]defStructs.Champ{
	0: {"shopName", -1, -1, nil},
}

// TdpvgsmT3t45_4
var TdpvgsmT3t45_4 = map[int]defStructs.Champ{
	0: {"ratingLevel", -1, -1, nil},
}

// TdpvgsmT3t45
var TdpvgsmT3t45 = map[int]defStructs.Champ{
	0x01: {"T4t45.1", -1, -1, &TdpvgsmT3t45_1},
	0x02: {"T4t45.2", -1, -1, &TdpvgsmT3t45_2},
	0x03: {"T4t45.3", -1, -1, &TdpvgsmT3t45_3},
	0x04: {"T4t45.4", -1, -1, &TdpvgsmT3t45_4},
}

// TdpvgsmT3t46_1
var TdpvgsmT3t46_1 = map[int]defStructs.Champ{
	0: {"TAEvolume", -1, -1, nil},
}

// TdpvgsmT3t46_2
var TdpvgsmT3t46_2 = map[int]defStructs.Champ{
	0: {"TAEduration", -1, -1, nil},
}

// TdpvgsmT3t46
var TdpvgsmT3t46 = map[int]defStructs.Champ{
	0x01: {"T4t46.1", -1, -1, &TdpvgsmT3t46_1},
	0x02: {"T4t46.2", -1, -1, &TdpvgsmT3t46_2},
}

// TdpvgsmT3t47_1
var TdpvgsmT3t47_1 = map[int]defStructs.Champ{
	0: {"onlinePurUsFlag", -1, -1, nil},
}

// TdpvgsmT3t47_2
var TdpvgsmT3t47_2 = map[int]defStructs.Champ{
	0: {"serviceClass", -1, 1, nil},
}

// TdpvgsmT3t47_3
var TdpvgsmT3t47_3 = map[int]defStructs.Champ{
	0: {"servPlatformID", -1, -1, nil},
}

// TdpvgsmT3t47
var TdpvgsmT3t47 = map[int]defStructs.Champ{
	0x01: {"T4t47.1", -1, -1, &TdpvgsmT3t47_1},
	0x02: {"T4t47.2", -1, -1, &TdpvgsmT3t47_2},
	0x03: {"T4t47.3", -1, -1, &TdpvgsmT3t47_3},
}

// TdpvgsmT3t48_1
var TdpvgsmT3t48_1 = map[int]defStructs.Champ{
	0: {"servIPaddress", -1, -1, nil},
}

// TdpvgsmT3t48_2
var TdpvgsmT3t48_2 = map[int]defStructs.Champ{
	0: {"servPortNb", -1, -1, nil},
}

// TdpvgsmT3t48_3
var TdpvgsmT3t48_3 = map[int]defStructs.Champ{
	0: {"servURL", -1, -1, nil},
}

// TdpvgsmT3t48_4
var TdpvgsmT3t48_4 = map[int]defStructs.Champ{
	0: {"carriedVolume", -1, -1, nil},
}

// TdpvgsmT3t48
var TdpvgsmT3t48 = map[int]defStructs.Champ{
	0x01: {"T4t48.1", -1, -1, &TdpvgsmT3t48_1},
	0x02: {"T4t48.2", -1, -1, &TdpvgsmT3t48_2},
	0x03: {"T4t48.3", -1, -1, &TdpvgsmT3t48_3},
	0x04: {"T4t48.4", -1, -1, &TdpvgsmT3t48_4},
}

// TdpvgsmT3t49_1
var TdpvgsmT3t49_1 = map[int]defStructs.Champ{
	0: {"usrAccType", -1, -1, nil},
}

// TdpvgsmT3t49_2
var TdpvgsmT3t49_2 = map[int]defStructs.Champ{
	0: {"usrChannelID", -1, -1, nil},
}

// TdpvgsmT3t49
var TdpvgsmT3t49 = map[int]defStructs.Champ{
	0x01: {"T4t49.1", -1, -1, &TdpvgsmT3t49_1},
	0x02: {"T4t49.2", -1, -1, &TdpvgsmT3t49_2},
}

// TdpvgsmT3t50_1
var TdpvgsmT3t50_1 = map[int]defStructs.Champ{
	0: {"partAccType", -1, -1, nil},
}

// TdpvgsmT3t50_2
var TdpvgsmT3t50_2 = map[int]defStructs.Champ{
	0: {"partChannelID", -1, -1, nil},
}

// TdpvgsmT3t50_3
var TdpvgsmT3t50_3 = map[int]defStructs.Champ{
	0: {"partTransID", -1, -1, nil},
}

// TdpvgsmT3t50
var TdpvgsmT3t50 = map[int]defStructs.Champ{
	0x01: {"T4t50.1", -1, -1, &TdpvgsmT3t50_1},
	0x02: {"T4t50.2", -1, -1, &TdpvgsmT3t50_2},
	0x03: {"T4t50.3", -1, -1, &TdpvgsmT3t50_3},
}

// TdpvgsmT3t51
var TdpvgsmT3t51 = map[int]defStructs.Champ{
	0: {"servPlatRoaInf", -1, -1, nil},
}

// TdpvgsmT3t52_1
var TdpvgsmT3t52_1 = map[int]defStructs.Champ{
	0: {"callingLocation", -1, -1, nil},
}

// TdpvgsmT3t52_2
var TdpvgsmT3t52_2 = map[int]defStructs.Champ{
	0: {"calledLocation", -1, -1, nil},
}

// TdpvgsmT3t52
var TdpvgsmT3t52 = map[int]defStructs.Champ{
	0x01: {"T4t52.1", -1, -1, &TdpvgsmT3t52_1},
	0x02: {"T4t52.2", -1, -1, &TdpvgsmT3t52_2},
}

// TdpvgsmT3t53_1
var TdpvgsmT3t53_1 = map[int]defStructs.Champ{
	0: {"callingGrpID", -1, -1, nil},
}

// TdpvgsmT3t53_2
var TdpvgsmT3t53_2 = map[int]defStructs.Champ{
	0: {"callingSiteID", -1, -1, nil},
}

// TdpvgsmT3t53_3
var TdpvgsmT3t53_3 = map[int]defStructs.Champ{
	0: {"calledGrpID", -1, -1, nil},
}

// TdpvgsmT3t53_4
var TdpvgsmT3t53_4 = map[int]defStructs.Champ{
	0: {"calledSiteID", -1, -1, nil},
}

// TdpvgsmT3t53
var TdpvgsmT3t53 = map[int]defStructs.Champ{
	0x01: {"T4t53.1", -1, -1, &TdpvgsmT3t53_1},
	0x02: {"T4t53.2", -1, -1, &TdpvgsmT3t53_2},
	0x03: {"T4t53.3", -1, -1, &TdpvgsmT3t53_3},
	0x04: {"T4t53.4", -1, -1, &TdpvgsmT3t53_4},
}

// TdpvgsmT3t54
var TdpvgsmT3t54 = map[int]defStructs.Champ{
	0: {"GFlag1", -1, -1, nil},
}

// TdpvgsmT3t55
var TdpvgsmT3t55 = map[int]defStructs.Champ{
	0: {"GFlag2", -1, -1, nil},
}

// TdpvgsmT3t56
var TdpvgsmT3t56 = map[int]defStructs.Champ{
	0: {"GValue", -1, -1, nil},
}

// TdpvgsmT3t57
var TdpvgsmT3t57 = map[int]defStructs.Champ{
	0: {"GDCB", -1, -1, nil},
}

// TdpvgsmT3t58
var TdpvgsmT3t58 = map[int]defStructs.Champ{
	0: {"GASCII", -1, -1, nil},
}

// TdpvgsmT3t59
var TdpvgsmT3t59 = map[int]defStructs.Champ{
	0: {"PrefixIMSIpart", 3, -1, nil},
	1: {"LocalTADIG", 5, -1, nil},
	2: {"PrefixIMSIclt", -1, -1, nil},
}

// TdpvgsmT3t60
var TdpvgsmT3t60 = map[int]defStructs.Champ{
	0: {"GMULTI2", -1, -1, nil},
}

// TdpvgsmT3t61_1
var TdpvgsmT3t61_1 = map[int]defStructs.Champ{
	0: {"calledPart", -1, -1, nil},
}

// TdpvgsmT3t61_2
var TdpvgsmT3t61_2 = map[int]defStructs.Champ{
	0: {"effectivePart", -1, -1, nil},
}

// TdpvgsmT3t61_3
var TdpvgsmT3t61_3 = map[int]defStructs.Champ{
	0: {"platRecCode", -1, -1, nil},
}

// TdpvgsmT3t61_4
var TdpvgsmT3t61_4 = map[int]defStructs.Champ{
	0: {"V4t61.4", -1, -1, nil},
}

// TdpvgsmT3t61_5
var TdpvgsmT3t61_5 = map[int]defStructs.Champ{
	0: {"V4t61.5", -1, -1, nil},
}

// TdpvgsmT3t61_6
var TdpvgsmT3t61_6 = map[int]defStructs.Champ{
	0: {"callID", -1, -1, nil},
}

// TdpvgsmT3t61_7
var TdpvgsmT3t61_7 = map[int]defStructs.Champ{
	0: {"GroupID", -1, -1, nil},
}

// TdpvgsmT3t61_8
var TdpvgsmT3t61_8 = map[int]defStructs.Champ{
	0: {"usrSpeechDur", -1, -1, nil},
}

// TdpvgsmT3t61_9
var TdpvgsmT3t61_9 = map[int]defStructs.Champ{
	0: {"speechBurstNb", -1, -1, nil},
}

// TdpvgsmT3t61
var TdpvgsmT3t61 = map[int]defStructs.Champ{
	0x01: {"T4t61.1", -1, -1, &TdpvgsmT3t61_1},
	0x02: {"T4t61.2", -1, -1, &TdpvgsmT3t61_2},
	0x03: {"T4t61.3", -1, -1, &TdpvgsmT3t61_3},
	0x04: {"T4t61.4", -1, -1, &TdpvgsmT3t61_4},
	0x05: {"T4t61.5", -1, -1, &TdpvgsmT3t61_5},
	0x06: {"T4t61.5", -1, -1, &TdpvgsmT3t61_6},
	0x07: {"T4t61.7", -1, -1, &TdpvgsmT3t61_7},
	0x08: {"T4t61.8", -1, -1, &TdpvgsmT3t61_8},
	0x09: {"T4t61.9", -1, -1, &TdpvgsmT3t61_9},
}

// TdpvgsmT3t62
var TdpvgsmT3t62 = map[int]defStructs.Champ{
	0: {"ASCIIString", -1, -1, nil},
}

// TdpvgsmT3t63
var TdpvgsmT3t63 = map[int]defStructs.Champ{
	0: {"LastCallingNb", -1, -1, nil},
}

// TdpvgsmT3t64
var TdpvgsmT3t64 = map[int]defStructs.Champ{
	0: {"RedirectingNb", -1, -1, nil},
}

// TdpvgsmT3t65
var TdpvgsmT3t65 = map[int]defStructs.Champ{
	0: {"RedirectReason", -1, -1, nil},
}

// TdpvgsmT3t66_1
var TdpvgsmT3t66_1 = map[int]defStructs.Champ{
	0: {"RecSeqNumber", 4, -1, nil},
	1: {"SigVolum", 4, -1, nil},
	2: {"ExtBearVolum", 4, -1, nil},
	3: {"TotalVolum", 4, -1, nil},
	4: {"CallEndTStamp", 7, 8, nil},
	5: {"CauseRecClose", 1, -1, nil},
	6: {"SessionTerm", 4, -1, nil},
	7: {"CORLStatus", -1, -1, nil},
}

// TdpvgsmT3t66_2
var TdpvgsmT3t66_2 = map[int]defStructs.Champ{
	0: {"SgsnAddr", -1, -1, nil},
}

// TdpvgsmT3t66_3
var TdpvgsmT3t66_3 = map[int]defStructs.Champ{
	0: {"GGSNAddr", -1, -1, nil},
}

// TdpvgsmT3t66_4
var TdpvgsmT3t66_4 = map[int]defStructs.Champ{
	0: {"SubRecordType", -1, -1, nil},
}

// TdpvgsmT3t66_5
var TdpvgsmT3t66_5 = map[int]defStructs.Champ{
	0: {"TransCalledNum", -1, -1, nil},
}

// TdpvgsmT3t66_6
var TdpvgsmT3t66_6 = map[int]defStructs.Champ{
	0: {"OriginatingIOI", -1, 1, nil},
}

// TdpvgsmT3t66_7
var TdpvgsmT3t66_7 = map[int]defStructs.Champ{
	0: {"TerminatingIOI", -1, 1, nil},
}

// TdpvgsmT3t66_8
var TdpvgsmT3t66_8 = map[int]defStructs.Champ{
	0: {"AccessNetInfo", -1, 1, nil},
}

// TdpvgsmT3t66_9
var TdpvgsmT3t66_9 = map[int]defStructs.Champ{
	0: {"SessionId", -1, 1, nil},
}

// TdpvgsmT3t66_10
var TdpvgsmT3t66_10 = map[int]defStructs.Champ{
	0: {"NodeIdentifier", -1, -1, nil},
}

// TdpvgsmT3t66_11
var TdpvgsmT3t66_11 = map[int]defStructs.Champ{
	0: {"AppliServerId", -1, 1, nil},
}

// TdpvgsmT3t66_12
var TdpvgsmT3t66_12 = map[int]defStructs.Champ{
	0: {"SIPMethod", -1, 1, nil},
}

// TdpvgsmT3t66_13
var TdpvgsmT3t66_13 = map[int]defStructs.Champ{
	0: {"PrivateUserId", -1, -1, nil},
}

// TdpvgsmT3t66
var TdpvgsmT3t66 = map[int]defStructs.Champ{
	0x01: {"T4t66.1", -1, -1, &TdpvgsmT3t66_1},
	0x02: {"T4t66.2", -1, -1, &TdpvgsmT3t66_2},
	0x03: {"T4t66.3", -1, -1, &TdpvgsmT3t66_3},
	0x04: {"T4t66.4", -1, -1, &TdpvgsmT3t66_4},
	0x05: {"T4t66.5", -1, -1, &TdpvgsmT3t66_5},
	0x06: {"T4t66.6", -1, -1, &TdpvgsmT3t66_6},
	0x07: {"T4t66.7", -1, -1, &TdpvgsmT3t66_7},
	0x08: {"T4t66.8", -1, -1, &TdpvgsmT3t66_8},
	0x09: {"T4t66.9", -1, -1, &TdpvgsmT3t66_9},
	0x10: {"T4t66.10", -1, -1, &TdpvgsmT3t66_10},
	0x11: {"T4t66.11", -1, -1, &TdpvgsmT3t66_11},
	0x12: {"T4t66.12", -1, -1, &TdpvgsmT3t66_12},
	0x13: {"T4t66.13", -1, -1, &TdpvgsmT3t66_13},
}

// TdpvgsmT3t67_1
var TdpvgsmT3t67_1 = map[int]defStructs.Champ{
	0: {"TermCause", 4, -1, nil},
	1: {"CauseCode", 1, -1, nil},
	2: {"AnswerIndicator", 1, -1, nil},
	3: {"RelCallIdReason", 4, -1, nil},
	4: {"OriCallIdReason", 4, -1, nil},
	5: {"ASRelParty", 6, -1, nil},
	6: {"UserTimeZone", -1, -1, nil},
}

// TdpvgsmT3t67_2
var TdpvgsmT3t67_2 = map[int]defStructs.Champ{
	0: {"AccSessionId", -1, -1, nil},
}

// TdpvgsmT3t67_3
var TdpvgsmT3t67_3 = map[int]defStructs.Champ{
	0: {"NetSessionId", -1, -1, nil},
}

// TdpvgsmT3t67_4
var TdpvgsmT3t67_4 = map[int]defStructs.Champ{
	0: {"RelatedCallId", -1, -1, nil},
}

// TdpvgsmT3t67_5
var TdpvgsmT3t67_5 = map[int]defStructs.Champ{
	0: {"LocalCallId", -1, -1, nil},
}

// TdpvgsmT3t67_6
var TdpvgsmT3t67_6 = map[int]defStructs.Champ{
	0: {"CodecUsed", -1, -1, nil},
}

// TdpvgsmT3t67_7
var TdpvgsmT3t67_7 = map[int]defStructs.Champ{
	0: {"CompanyId", -1, -1, nil},
}

// TdpvgsmT3t67_8
var TdpvgsmT3t67_8 = map[int]defStructs.Champ{
	0: {"RecCallingNum", -1, -1, nil},
}

// TdpvgsmT3t67_9
var TdpvgsmT3t67_9 = map[int]defStructs.Champ{
	0: {"NDI", -1, -1, nil},
}

// TdpvgsmT3t67
var TdpvgsmT3t67 = map[int]defStructs.Champ{
	0x01: {"T4t67.1", -1, -1, &TdpvgsmT3t67_1},
	0x02: {"T4t67.2", -1, -1, &TdpvgsmT3t67_2},
	0x03: {"T4t67.3", -1, -1, &TdpvgsmT3t67_3},
	0x04: {"T4t67.4", -1, -1, &TdpvgsmT3t67_4},
	0x05: {"T4t67.5", -1, -1, &TdpvgsmT3t67_5},
	0x06: {"T4t67.6", -1, -1, &TdpvgsmT3t67_6},
	0x07: {"T4t67.7", -1, -1, &TdpvgsmT3t67_7},
	0x08: {"T4t67.8", -1, -1, &TdpvgsmT3t67_8},
	0x09: {"T4t67.9", -1, -1, &TdpvgsmT3t67_9},
}

// TdpvgsmT3t68
var TdpvgsmT3t68 = map[int]defStructs.Champ{
	0: {"DataIllimitee", 1, -1, nil},
	1: {"CarteJumelle", 1, -1, nil},
	2: {"ReserveT3t68", -1, -1, nil},
}

// TdpvgsmT3t69
var TdpvgsmT3t69 = map[int]defStructs.Champ{
	0: {"PrefItinerance", -1, -1, nil},
}

// TdpvgsmT3t70
var TdpvgsmT3t70 = map[int]defStructs.Champ{
	0: {"PrefixeSVA", -1, -1, nil},
}

// TdpvgsmT3t71_1
var TdpvgsmT3t71_1 = map[int]defStructs.Champ{
	0: {"QoSClassId", -1, -1, nil},
}

// TdpvgsmT3t71_2
var TdpvgsmT3t71_2 = map[int]defStructs.Champ{
	0: {"maxReqBandUp", -1, -1, nil},
}

// TdpvgsmT3t71_3
var TdpvgsmT3t71_3 = map[int]defStructs.Champ{
	0: {"maxReqBandDown", -1, -1, nil},
}

// TdpvgsmT3t71_4
var TdpvgsmT3t71_4 = map[int]defStructs.Champ{
	0: {"garBitrateUp", -1, -1, nil},
}

// TdpvgsmT3t71_5
var TdpvgsmT3t71_5 = map[int]defStructs.Champ{
	0: {"garBitrateDown", -1, -1, nil},
}

// TdpvgsmT3t71_6
var TdpvgsmT3t71_6 = map[int]defStructs.Champ{
	0: {"RetentionPrior", -1, -1, nil},
}

// TdpvgsmT3t71
var TdpvgsmT3t71 = map[int]defStructs.Champ{
	0x01: {"T4t71.1", -1, -1, &TdpvgsmT3t71_1},
	0x02: {"T4t71.2", -1, -1, &TdpvgsmT3t71_2},
	0x03: {"T4t71.3", -1, -1, &TdpvgsmT3t71_3},
	0x04: {"T4t71.4", -1, -1, &TdpvgsmT3t71_4},
	0x05: {"T4t71.5", -1, -1, &TdpvgsmT3t71_5},
	0x06: {"T4t71.6", -1, -1, &TdpvgsmT3t71_6},
}

// TdpvgsmT3t72
var TdpvgsmT3t72 = map[int]defStructs.Champ{
	0: {"TypeEditeurSVA", -1, -1, nil},
}

// TdpvgsmT3t73
var TdpvgsmT3t73 = map[int]defStructs.Champ{
	0: {"CodeTarifSVA", -1, -1, nil},
}

// TdpvgsmT3t74_1
var TdpvgsmT3t74_1 = map[int]defStructs.Champ{
	0: {"ProviderName", -1, -1, nil},
}

// TdpvgsmT3t74_2
var TdpvgsmT3t74_2 = map[int]defStructs.Champ{
	0: {"CatContentCode", -1, -1, nil},
}

// TdpvgsmT3t74
var TdpvgsmT3t74 = map[int]defStructs.Champ{
	0x01: {"T4t74.1", -1, -1, &TdpvgsmT3t74_1},
	0x02: {"T4t74.2", -1, -1, &TdpvgsmT3t74_2},
}

// TdpvgsmT3t99_1
var TdpvgsmT3t99_1 = map[int]defStructs.Champ{
	0: {"GretaTag", -1, -1, nil},
}

// TdpvgsmT3t99_2
var TdpvgsmT3t99_2 = map[int]defStructs.Champ{
	0: {"GMSCForwTerm", -1, -1, nil},
}

// TdpvgsmT3t99_3
var TdpvgsmT3t99_3 = map[int]defStructs.Champ{
	0: {"FilterNumber", -1, -1, nil},
}

// TdpvgsmT3t99
var TdpvgsmT3t99 = map[int]defStructs.Champ{
	0x01: {"T4t99.1", -1, -1, &TdpvgsmT3t99_1},
	0x02: {"T4t99.2", -1, -1, &TdpvgsmT3t99_2},
	0x03: {"T4t99.3", -1, -1, &TdpvgsmT3t99_3},
}

// TdpvgsmT3
var TdpvgsmT3 = map[int]defStructs.Champ{
	0x11: {"T3t1", -1, -1, &TdpvgsmT3t1},
	0x12: {"T3t2", -1, -1, &TdpvgsmT3t2},
	0x13: {"T3t3", -1, -1, &TdpvgsmT3t3},
	0x14: {"T3t4", -1, -1, &TdpvgsmT3t4},
	0x15: {"T3t5", -1, -1, &TdpvgsmT3t5},
	0x16: {"T3t6", -1, -1, &TdpvgsmT3t6},
	0x17: {"T3t7", -1, -1, &TdpvgsmT3t7},
	0x18: {"T3t8", -1, -1, &TdpvgsmT3t8},
	0x1b: {"T3t11", -1, -1, &TdpvgsmT3t11},
	0x1c: {"T3t12", -1, -1, &TdpvgsmT3t12},
	0x1d: {"T3t13", -1, -1, &TdpvgsmT3t13},
	0x1f: {"T3t15", -1, -1, &TdpvgsmT3t15},
	0x20: {"T3t16", -1, -1, &TdpvgsmT3t16},
	0x21: {"T3t17", -1, -1, &TdpvgsmT3t17},
	0x22: {"T3t18", -1, -1, &TdpvgsmT3t18},
	0x24: {"T3t20", -1, -1, &TdpvgsmT3t20},
	0x25: {"T3t21", -1, -1, &TdpvgsmT3t21},
	0x26: {"T3t22", -1, -1, &TdpvgsmT3t22},
	0x27: {"T3t23", -1, -1, &TdpvgsmT3t23},
	0x28: {"T3t24", -1, -1, &TdpvgsmT3t24},
	0x29: {"T3t25", -1, -1, &TdpvgsmT3t25},
	0x2a: {"T3t26", -1, -1, &TdpvgsmT3t26},
	0x2b: {"T3t27", -1, -1, &TdpvgsmT3t27},
	0x2c: {"T3t28", -1, -1, &TdpvgsmT3t28},
	0x2e: {"T3t29", -1, -1, &TdpvgsmT3t29},
	0x30: {"T3t30", -1, -1, &TdpvgsmT3t30},
	0x2f: {"T3t31", -1, -1, &TdpvgsmT3t31},
	0x31: {"T3t32", -1, -1, &TdpvgsmT3t32},
	0x33: {"T3t33", -1, -1, &TdpvgsmT3t33},
	0x34: {"T3t34", -1, -1, &TdpvgsmT3t34},
	0x35: {"T3t35", -1, -1, &TdpvgsmT3t35},
	0x36: {"T3t36", -1, -1, &TdpvgsmT3t36},
	0x37: {"T3t37", -1, -1, &TdpvgsmT3t37},
	0x38: {"T3t38", -1, -1, &TdpvgsmT3t38},
	0x39: {"T3t39", -1, -1, &TdpvgsmT3t39},
	0x3a: {"T3t40", -1, -1, &TdpvgsmT3t40},
	0x3b: {"T3t41", -1, -1, &TdpvgsmT3t41},
	0x3c: {"T3t42", -1, -1, &TdpvgsmT3t42},
	0x3d: {"T3t43", -1, -1, &TdpvgsmT3t43},
	0x3e: {"T3t44", -1, -1, &TdpvgsmT3t44},
	0x3f: {"T3t45", -1, -1, &TdpvgsmT3t45},
	0x40: {"T3t46", -1, -1, &TdpvgsmT3t46},
	0x41: {"T3t47", -1, -1, &TdpvgsmT3t47},
	0x43: {"T3t48", -1, -1, &TdpvgsmT3t48},
	0x44: {"T3t49", -1, -1, &TdpvgsmT3t49},
	0x45: {"T3t50", -1, -1, &TdpvgsmT3t50},
	0x46: {"T3t51", -1, -1, &TdpvgsmT3t51},
	0x47: {"T3t52", -1, -1, &TdpvgsmT3t52},
	0x48: {"T3t53", -1, -1, &TdpvgsmT3t53},
	0x49: {"T3t54", -1, -1, &TdpvgsmT3t54},
	0x4a: {"T3t55", -1, -1, &TdpvgsmT3t55},
	0x4b: {"T3t56", -1, -1, &TdpvgsmT3t56},
	0x4c: {"T3t57", -1, -1, &TdpvgsmT3t57},
	0x4d: {"T3t58", -1, -1, &TdpvgsmT3t58},
	0x4e: {"T3t59", -1, -1, &TdpvgsmT3t59},
	0x4f: {"T3t60", -1, -1, &TdpvgsmT3t60},
	0x50: {"T3t61", -1, -1, &TdpvgsmT3t61},
	0x51: {"T3t62", -1, -1, &TdpvgsmT3t62},
	0x52: {"T3t63", -1, -1, &TdpvgsmT3t63},
	0x53: {"T3t64", -1, -1, &TdpvgsmT3t64},
	0x54: {"T3t65", -1, -1, &TdpvgsmT3t65},
	0x55: {"T3t66", -1, -1, &TdpvgsmT3t66},
	0x56: {"T3t67", -1, -1, &TdpvgsmT3t67},
	0x57: {"T3t68", -1, -1, &TdpvgsmT3t68},
	0x58: {"T3t69", -1, -1, &TdpvgsmT3t69},
	0x59: {"T3t70", -1, -1, &TdpvgsmT3t70},
	0x5a: {"T3t71", -1, -1, &TdpvgsmT3t71},
	0x5b: {"T3t72", -1, -1, &TdpvgsmT3t72},
	0x5c: {"T3t73", -1, -1, &TdpvgsmT3t73},
	0x5d: {"T3t74", -1, -1, &TdpvgsmT3t74},
	0x99: {"T3t99", -1, -1, &TdpvgsmT3t99},
}

// TdpvgsmT2
var TdpvgsmT2 = map[int]defStructs.Champ{
	0x02: {"T2", -1, -1, &TdpvgsmT3},
}

// Tdpvgsm
var Tdpvgsm = map[int]defStructs.Champ{
	0x4d: {"Tdpvgsm", -1, -1, &TdpvgsmT2},
}

// Tdpvgsm33
var Tdpvgsm33 = map[int]defStructs.Champ{
	0x31: {"Tdpvgsm33", -1, -1, &TdpvgsmT2},
}

// Ddpvgsm
var Ddpvgsm = map[int]defStructs.TabRechercheTypeDc{
	0: {"Tdpvgsm", 0xffff0000, 0x804d0000, 0, &Tdpvgsm},
	1: {"Tdpvgsm33", 0xffff0000, 0x80310000, 0, &Tdpvgsm33},
	// Rajout pour norego (SHM dans RJP)
	2: {"Tdpvgsm28", 0xffff0000, 0x80110000, 0, &Tdpvgsm28},
	// TEST
	3: {"Tdpvgsm32", 0xffff0000, 0x80140000, 0, &Tdpvgsm32},
	4: {"TdpvgsmFree", 0xffff0000, 0x804e0000, 0, &TdpvgsmFree},
}
