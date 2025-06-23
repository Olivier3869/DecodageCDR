package structuresEccTLV

import "local/structures/defStructs"

// TdpvgsmFreeT3t1
var TdpvgsmFreeT3t1 = map[int]defStructs.Champ{
	0: {"CallType", -1, -1, nil},
}

// TdpvgsmFreeT3t2
var TdpvgsmFreeT3t2 = map[int]defStructs.Champ{
	0: {"IMSI", -1, -1, nil},
}

// TdpvgsmFreeT3t3
var TdpvgsmFreeT3t3 = map[int]defStructs.Champ{
	0: {"BeginningDateTime", -1, -1, nil},
}

// TdpvgsmFreeT3t4
var TdpvgsmFreeT3t4 = map[int]defStructs.Champ{
	0: {"Duration", -1, -1, nil},
}

// TdpvgsmFreeT3t5
var TdpvgsmFreeT3t5 = map[int]defStructs.Champ{
	0: {"BearerService", -1, -1, nil},
}

// TdpvgsmFreeT3t6
var TdpvgsmFreeT3t6 = map[int]defStructs.Champ{
	0: {"Teleservice", -1, -1, nil},
}

// TdpvgsmFreeT3t7
var TdpvgsmFreeT3t7 = map[int]defStructs.Champ{
	0: {"SystemType", -1, -1, nil},
}

// TdpvgsmFreeT3t8
var TdpvgsmFreeT3t8 = map[int]defStructs.Champ{
	0: {"EfficiencyIndicator", -1, -1, nil},
}

// TdpvgsmFreeT3t9
var TdpvgsmFreeT3t9 = map[int]defStructs.Champ{
	0: {"CallTerminationType", -1, -1, nil},
}

// TdpvgsmFreeT3t10
var TdpvgsmFreeT3t10 = map[int]defStructs.Champ{
	0: {"CallPartyIdentity", -1, -1, nil},
}

// TdpvgsmFreeT3t11
var TdpvgsmFreeT3t11 = map[int]defStructs.Champ{
	0: {"LocationNumber", -1, -1, nil},
}

// TdpvgsmFreeT3t12
var TdpvgsmFreeT3t12 = map[int]defStructs.Champ{
	0: {"MSISDN", -1, -1, nil},
}

// TdpvgsmFreeT3t13
var TdpvgsmFreeT3t13 = map[int]defStructs.Champ{
	0: {"IMEI", -1, -1, nil},
}

// TdpvgsmFreeT3t14
var TdpvgsmFreeT3t14 = map[int]defStructs.Champ{
	0: {"CircuitAllocationTimestamp", -1, -1, nil},
}

// TdpvgsmFreeT3t15_1
var TdpvgsmFreeT3t15_1 = map[int]defStructs.Champ{
	0: {"PDPaddress", -1, 9, nil},
}

// TdpvgsmFreeT3t15_2
var TdpvgsmFreeT3t15_2 = map[int]defStructs.Champ{
	0: {"AccessPointNameNI", -1, 1, nil},
}

// TdpvgsmFreeT3t15_3
var TdpvgsmFreeT3t15_3 = map[int]defStructs.Champ{
	0: {"AccessPointNameOI", -1, 1, nil},
}

// TdpvgsmFreeT3t15_4
var TdpvgsmFreeT3t15_4 = map[int]defStructs.Champ{
	0: {"ChargingId", -1, -1, nil},
}

// TdpvgsmFreeT3t15_5
var TdpvgsmFreeT3t15_5 = map[int]defStructs.Champ{
	0: {"RecordSequenceNumber", -1, -1, nil},
}

// TdpvgsmFreeT3t15_6
var TdpvgsmFreeT3t15_6 = map[int]defStructs.Champ{
	0: {"PartialTypeIndicator", -1, -1, nil},
}

// TdpvgsmFreeT3t15_7
var TdpvgsmFreeT3t15_7 = map[int]defStructs.Champ{
	0: {"RecordingEntityIdentificationGGSN", -1, -1, nil},
}

// TdpvgsmFreeT3t15_8
var TdpvgsmFreeT3t15_8 = map[int]defStructs.Champ{
	0: {"RecordingEntityIdentificationSGSN1", -1, -1, nil},
}

// TdpvgsmFreeT3t15_9
var TdpvgsmFreeT3t15_9 = map[int]defStructs.Champ{
	0: {"RecordingEntityIdentificationSGSN2", -1, -1, nil},
}

// TdpvgsmFreeT3t15_10
var TdpvgsmFreeT3t15_10 = map[int]defStructs.Champ{
	0: {"TotalSumDataVolumeOutgoing", 6, -1, nil},
	1: {"TotalSumDataVolumeIncoming", -1, -1, nil},
}

// TdpvgsmFreeT3t15
var TdpvgsmFreeT3t15 = map[int]defStructs.Champ{
	0x01: {"T4t15_1", -1, -1, &TdpvgsmFreeT3t15_1},
	0x02: {"T4t15.2", -1, -1, &TdpvgsmFreeT3t15_2},
	0x03: {"T4t15.3", -1, -1, &TdpvgsmFreeT3t15_3},
	0x04: {"T4t15.4", -1, -1, &TdpvgsmFreeT3t15_4},
	0x05: {"T4t15.5", -1, -1, &TdpvgsmFreeT3t15_5},
	0x06: {"T4t15.6", -1, -1, &TdpvgsmFreeT3t15_6},
	0x07: {"T4t15.7", -1, -1, &TdpvgsmFreeT3t15_7},
	0x08: {"T4t15.8", -1, -1, &TdpvgsmFreeT3t15_8},
	0x09: {"T4t15.9", -1, -1, &TdpvgsmFreeT3t15_9},
	0x0a: {"T4t15.10", -1, -1, &TdpvgsmFreeT3t15_10},
}

// TdpvgsmFreeT3t16_1
var TdpvgsmFreeT3t16_1 = map[int]defStructs.Champ{
	0: {"CDRreceptionDate", -1, -1, nil},
}

// TdpvgsmFreeT3t16_2
var TdpvgsmFreeT3t16_2 = map[int]defStructs.Champ{
	0: {"EECode", -1, -1, nil},
}

// TdpvgsmFreeT3t16
var TdpvgsmFreeT3t16 = map[int]defStructs.Champ{
	0x20: {"T3t16_1", -1, -1, &TdpvgsmFreeT3t16_1},
	0x21: {"T3t16_2", -1, -1, &TdpvgsmFreeT3t16_2},
}

// TdpvgsmFreeT3t17
var TdpvgsmFreeT3t17 = map[int]defStructs.Champ{
	0: {"IncomingZonesBlanchesPrefix", -1, -1, nil},
}

// TdpvgsmFreeT3t18
var TdpvgsmFreeT3t18 = map[int]defStructs.Champ{
	0: {"OutgoingZonesBlanchesPrefix", -1, -1, nil},
}

// TdpvgsmFreeT3t19
var TdpvgsmFreeT3t19 = map[int]defStructs.Champ{
	0: {"LastCallingNumber", -1, -1, nil},
}

// TdpvgsmFreeT3t20
var TdpvgsmFreeT3t20 = map[int]defStructs.Champ{
	0: {"RedirectingNumber", -1, -1, nil},
}

// TdpvgsmFreeT3t21
var TdpvgsmFreeT3t21 = map[int]defStructs.Champ{
	0: {"PrefixeItinerance", -1, -1, nil},
}

// TdpvgsmFreeT3
var TdpvgsmFreeT3 = map[int]defStructs.Champ{
	0x01: {"T3t1", -1, -1, &TdpvgsmFreeT3t1},
	0x02: {"T3t2", -1, -1, &TdpvgsmFreeT3t2},
	0x03: {"T3t3", -1, -1, &TdpvgsmFreeT3t3},
	0x04: {"T3t4", -1, -1, &TdpvgsmFreeT3t4},
	0x05: {"T3t5", -1, -1, &TdpvgsmFreeT3t5},
	0x06: {"T3t6", -1, -1, &TdpvgsmFreeT3t6},
	0x07: {"T3t7", -1, -1, &TdpvgsmFreeT3t7},
	0x08: {"T3t8", -1, -1, &TdpvgsmFreeT3t8},
	0x09: {"T3t9", -1, -1, &TdpvgsmFreeT3t9},
	0x0a: {"T3t10", -1, -1, &TdpvgsmFreeT3t10},
	0x0b: {"T3t11", -1, -1, &TdpvgsmFreeT3t11},
	0x0c: {"T3t12", -1, -1, &TdpvgsmFreeT3t12},
	0x0d: {"T3t13", -1, -1, &TdpvgsmFreeT3t13},
	0x0e: {"T3t14", -1, -1, &TdpvgsmFreeT3t14},
	0x0f: {"SpecificInformationForPacket", -1, -1, &TdpvgsmFreeT3t15},
	0x10: {"T3t16", -1, -1, &TdpvgsmFreeT3t16},
	0x11: {"T3t17", -1, -1, &TdpvgsmFreeT3t17},
	0x12: {"T3t18", -1, -1, &TdpvgsmFreeT3t18},
	0x13: {"T3t19", -1, -1, &TdpvgsmFreeT3t19},
	0x14: {"T3t20", -1, -1, &TdpvgsmFreeT3t20},
	0x15: {"T3t21", -1, -1, &TdpvgsmFreeT3t21},
}

// TdpvgsmFreeT2
var TdpvgsmFreeT2 = map[int]defStructs.Champ{
	0x02: {"T2", -1, -1, &TdpvgsmFreeT3},
}

// TdpvgsmFree
var TdpvgsmFree = map[int]defStructs.Champ{
	0x4e: {"TdpvgsmFree", -1, -1, &TdpvgsmFreeT2},
}

// DdpvgsmFree
var DdpvgsmFree = map[int]defStructs.TabRechercheTypeDc{
	0: {"TdpvgsmFree", 0xffff0000, 0x804e0000, 0, &TdpvgsmFree},
}
