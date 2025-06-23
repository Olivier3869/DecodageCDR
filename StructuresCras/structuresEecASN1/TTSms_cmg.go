package structuresEecASN1

import "local/structures/defStructs"

// TTSms_cmg
var TTSms_cmg = map[int]defStructs.TabRechercheTypeDc{
	0: {"TTsms_cmg", 0xff000000, 0x30000000, 0, &TSms_cmg},
}

// TSms_cmg
var TSms_cmg = map[int]defStructs.Champ{
	16: {"TTsms_cmg", -1, 0, &CallDetailRecord},
}

var CallDetailRecord = map[int]defStructs.Champ{
	0:  {"origAddress", -1, 5, &notifAddressExt},
	1:  {"origAddressGSM", -1, 5, nil},
	2:  {"recipAddress", -1, 5, &recipAddressExt},
	3:  {"recipAddressGSM", -1, 5, nil},
	4:  {"submitDate", -1, 14, nil},
	5:  {"submitTime", -1, 15, nil},
	6:  {"status", -1, 0, nil},
	7:  {"terminDate", -1, 14, nil},
	8:  {"terminTime", -1, 15, nil},
	9:  {"lengthOfMessage", -1, 0, nil},
	10: {"prioIndicator", -1, 0, nil},
	11: {"validityPeriod", -1, 0, &validityPeriodExt},
	12: {"deferIndicator", -1, 0, nil},
	13: {"deferPeriod", -1, 0, &validityPeriodExt},
	14: {"notifIndicator", -1, 0, nil},
	15: {"notifAddress", -1, 0, &notifAddressExt},
	16: {"notifAddressGSM", -1, 5, nil},
	17: {"vsmscid", -1, 0, nil},
	18: {"vsmsctype", -1, 0, nil},
	19: {"dgtiAddress", -1, 0, &notifAddressExt},
	20: {"dgtiAddressGSM", -1, 0, nil},
	21: {"destPointCode", -1, 0, nil},
	22: {"ogtiAddress", -1, 0, &notifAddressExt},
	23: {"ogtiAddressGSM", -1, 5, nil},
	24: {"origPointCode", -1, 0, nil},
	25: {"orglSubmitDate", -1, 14, nil},
	26: {"orglSubmitTime", -1, 15, nil},
	27: {"transparentPid", -1, 0, nil},
	28: {"mesgReplyPath", -1, 0, nil},
	29: {"intlMobileSubId", -1, 5, nil},
	30: {"callingLineId", -1, 5, &notifAddressExt},
	31: {"callingLineIdGSM", -1, 5, nil},
	32: {"consolidation", -1, 0, nil},
	33: {"portNumber", -1, 0, nil},
	38: {"origIntlMobileSubId", -1, 0, nil},
	39: {"billid", -1, 0, nil},
	40: {"lang", -1, 0, nil},
	41: {"cbat", -1, 0, nil},
	42: {"ppPser", -1, 0, nil},
	43: {"ppAser", -1, 0, nil},
	44: {"ppAserDuringJam", -1, 0, nil},
	45: {"ppAserFree", -1, 0, nil},
	46: {"ppAserRecip", -1, 0, nil},
	47: {"smeReference", -1, 0, nil},
	48: {"smsContents", -1, 0, nil},
	49: {"smsContentDcs", -1, 0, nil},
	50: {"cmReferenceNr", -1, 0, nil},
	51: {"currentSegment", -1, 0, nil},
	52: {"segmentsTotal", -1, 0, nil},
	53: {"textFormatting", -1, 0, nil},
	54: {"bytesCompressedData", -1, 0, nil},
	55: {"predefinedAnimations", -1, 0, nil},
	56: {"userDefinedAnimations", -1, 0, nil},
	57: {"predefinedSounds", -1, 0, nil},
	58: {"userDefinedSounds", -1, 0, nil},
	59: {"blackWhitePictures", -1, 0, nil},
	60: {"standardWvg", -1, 0, nil},
	61: {"characterSizeWvg", -1, 0, nil},
	62: {"greyscalePictures", -1, 0, nil},
	63: {"colourPictures", -1, 0, nil},
	64: {"businessCards", -1, 0, nil},
	65: {"calendarEntries", -1, 0, nil},
	66: {"polyphonicMelodies", -1, 0, nil},
	67: {"bit8PortNumberDest", -1, 0, nil},
	68: {"bit16PortNumberDest", -1, 0, nil},
	69: {"lmsgNrSeg", -1, 0, nil},
	70: {"messageReference", -1, 0, nil},
	71: {"boolSer", -1, 0, nil},
	72: {"origLASN", -1, 0, nil},
	73: {"recipLASN", -1, 0, nil},
	74: {"origMsgID", -1, 0, nil},
	75: {"recipMsgID", -1, 0, nil},
	76: {"receiptDate", -1, 14, nil},
	77: {"receiptTime", -1, 15, nil},
	78: {"isr", -1, 0, nil},
	79: {"recipAltAddress", -1, 0, &notifAddressExt},
	80: {"generatedSegments", -1, 0, nil},
	81: {"serviceType", -1, 0, nil},
}

var recipAddressExt = map[int]defStructs.Champ{
	0: {"TypeOfNumber", -1, 0, nil},
	1: {"NumberingPlanIndicator", -1, 0, nil},
	2: {"ProtocolIdentifier", -1, 0, nil},
	3: {"NumericString", -1, 0, nil},
}

var validityPeriodExt = map[int]defStructs.Champ{
	0: {"hours", -1, 0, nil},
	1: {"minutes", -1, 0, nil},
}

var notifAddressExt = map[int]defStructs.Champ{
	0: {"TypeOfNumber", -1, 0, nil},
	1: {"NumberingPlanIndicator", -1, 0, nil},
	2: {"ProtocolIdentifier", -1, 0, nil},
	3: {"NumericString", -1, 0, nil},
}
