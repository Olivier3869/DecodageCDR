package structuresEccTLV

import "local/structures/defStructs"

// TrevodT4orp_7e
var TrevodT4orp_7e = map[int]defStructs.Champ{
	0x01: {"typeAcces", -1, 1, nil},
	0x02: {"typeClient", -1, 1, nil},
}

// TrevodT4orp_07
var TrevodT4orp_07 = map[int]defStructs.Champ{
	0x01: {"serviceType", -1, 1, nil},
}

// TrevodT4orp_09
var TrevodT4orp_09 = map[int]defStructs.Champ{
	0x01: {"terminalInfo", -1, 1, nil},
}

// TrevodT4orp_06
var TrevodT4orp_06 = map[int]defStructs.Champ{
	0x01: {"priceTTC", -1, 1, nil},
	0x02: {"priceHT", -1, 1, nil},
	0x03: {"currency", -1, 1, nil},
	0x04: {"catalogPrice", -1, 1, nil},
	0x05: {"paymentMethod", -1, 1, nil},
	0x06: {"specialOffer", -1, 1, nil},
	0x07: {"geographicZone", -1, 1, nil},
}

// TrevodT4orp_05
var TrevodT4orp_05 = map[int]defStructs.Champ{
	0x01: {"providerId", -1, 1, nil},
	0x02: {"providerName", -1, 1, nil},
}

// TrevodT4orp_04
var TrevodT4orp_04 = map[int]defStructs.Champ{
	0x02: {"pageName", -1, 1, nil},
	0x06: {"advisory", -1, 1, nil},
	0x07: {"pageId", -1, 1, nil},
	0x04: {"contentFormat", -1, 1, nil},
}

// TrevodT4orp_03
var TrevodT4orp_03 = map[int]defStructs.Champ{
	0x01: {"customerId", -1, 1, nil},
	0x40: {"phoneNumber", -1, 1, nil},
}

// TrevodT4orp_02
var TrevodT4orp_02 = map[int]defStructs.Champ{
	0x02: {"packageOrderId", -1, 1, nil},
	0x04: {"transactionDate", -1, 1, nil},
}

// TrevodT3orp
var TrevodT3orp = map[int]defStructs.Champ{
	0x02: {"transactionIdent", -1, 0, &TrevodT4orp_02},
	0x03: {"customerIdent", -1, 0, &TrevodT4orp_03},
	0x04: {"contentIdent", -1, 0, &TrevodT4orp_04},
	0x05: {"providerIdent", -1, 0, &TrevodT4orp_05},
	0x06: {"billingElement", -1, 0, &TrevodT4orp_06},
	0x07: {"contextIdent", -1, 0, &TrevodT4orp_07},
	0x01: {"CodePlateforme", -1, 1, nil},
	0x09: {"TerminalConnection", -1, 0, &TrevodT4orp_09},
	0x0c: {"codPromo", -1, 1, nil},
	0x0d: {"libPromo", -1, 1, nil},
	0x7e: {"ajout", -1, -1, &TrevodT4orp_7e},
}

// TrevodT4dlv_7e
var TrevodT4dlv_7e = map[int]defStructs.Champ{
	0x01: {"typeAcces", -1, 1, nil},
	0x02: {"typeClient", -1, 1, nil},
}

// TrevodT4dlv_08
var TrevodT4dlv_08 = map[int]defStructs.Champ{
	0x01: {"percentageWatched", -1, 1, nil},
	0x02: {"durationVisualization", -1, 1, nil},
	0x03: {"sessionStart", -1, 1, nil},
	0x04: {"sessionStop", -1, 1, nil},
	0x05: {"status", -1, 1, nil},
	0x06: {"errorCode", -1, 1, nil},
}

// TrevodT4dlv_07
var TrevodT4dlv_07 = map[int]defStructs.Champ{
	0x01: {"serviceType", -1, 1, nil},
	0x02: {"SVOSpackageId", -1, 1, nil},
}

// TrevodT4dlv_04
var TrevodT4dlv_04 = map[int]defStructs.Champ{
	0x01: {"mediaFileId", -1, 1, nil},
	0x02: {"contentTitle", -1, 1, nil},
	0x03: {"contentSubTitle", -1, 1, nil},
	0x04: {"contentFormat", -1, 1, nil},
	0x05: {"contentId", -1, 1, nil},
	0x06: {"asserId", -1, 1, nil},
	0x07: {"pageId", -1, 1, nil},
}

// TrevodT4dlv_03
var TrevodT4dlv_03 = map[int]defStructs.Champ{
	0x01: {"customerId", -1, 1, nil},
	0x40: {"phoneNumber", -1, 1, nil},
	0x03: {"terminalId", -1, 1, nil},
	0x04: {"videoServerId", -1, 1, nil},
}

// TrevodT4dlv_02
var TrevodT4dlv_02 = map[int]defStructs.Champ{
	0x02: {"transactionId", -1, 1, nil},
	0x03: {"ticketId", -1, 1, nil},
	0x04: {"transactionDate", -1, 1, nil},
	0x05: {"VODsessionId", -1, 1, nil},
}

// TrevodT3dlv
var TrevodT3dlv = map[int]defStructs.Champ{
	0x02: {"transactionIdent", -1, 0, &TrevodT4dlv_02},
	0x03: {"customerIdent", -1, 0, &TrevodT4dlv_03},
	0x04: {"contentIdent", -1, 0, &TrevodT4dlv_04},
	0x08: {"visualisation", -1, 0, &TrevodT4dlv_08},
	0x07: {"contextIdent", -1, 0, &TrevodT4dlv_07},
	0x01: {"CodePlateforme", -1, 1, nil},
	0x7e: {"ajout", -1, 1, &TrevodT4dlv_7e},
}

// TrevodT4ord_7e
var TrevodT4ord_7e = map[int]defStructs.Champ{
	0x01: {"typeAcces", -1, 1, nil},
	0x02: {"typeClient", -1, 1, nil},
}

// TrevodT4ord_09
var TrevodT4ord_09 = map[int]defStructs.Champ{
	0x01: {"terminalInfo", -1, 1, nil},
}

// TrevodT4ord_07
var TrevodT4ord_07 = map[int]defStructs.Champ{
	0x01: {"serviceType", -1, 1, nil},
	0x02: {"SVOSpackageId", -1, 1, nil},
}

// TrevodT4ord_06
var TrevodT4ord_06 = map[int]defStructs.Champ{
	0x01: {"priceTTC", -1, 1, nil},
	0x02: {"priceHT", -1, 1, nil},
	0x03: {"currency", -1, 1, nil},
	0x04: {"catalogPrice", -1, 1, nil},
	0x05: {"paymentMethod", -1, 1, nil},
	0x06: {"specialOffer", -1, 1, nil},
	0x07: {"geographicZone", -1, 1, nil},
}

// TrevodT4ord_05
var TrevodT4ord_05 = map[int]defStructs.Champ{
	0x01: {"providerId", -1, 1, nil},
	0x02: {"providerName", -1, 1, nil},
}

// TrevodT4ord_04
var TrevodT4ord_04 = map[int]defStructs.Champ{
	0x02: {"contentTitle", -1, 1, nil},
	0x03: {"contentSubTitle", -1, 1, nil},
	0x04: {"contentFormat", -1, 1, nil},
	0x05: {"contentId", -1, 1, nil},
	0x06: {"asserId", -1, 1, nil},
	0x07: {"pageId", -1, -1, nil},
}

// TrevodT4ord_03
var TrevodT4ord_03 = map[int]defStructs.Champ{
	0x01: {"customerId", -1, 1, nil},
	0x40: {"phoneNumber", -1, 1, nil},
}

// TrevodT4ord_02
var TrevodT4ord_02 = map[int]defStructs.Champ{
	0x02: {"transactionId", -1, 1, nil},
	0x04: {"transactionDate", -1, 1, nil},
}

// TrevodT3ord
var TrevodT3ord = map[int]defStructs.Champ{
	0x02: {"transactionIdent", -1, 0, &TrevodT4ord_02},
	0x03: {"customerIdent", -1, 0, &TrevodT4ord_03},
	0x04: {"contentIdent", -1, 0, &TrevodT4ord_04},
	0x05: {"providerIdent", -1, 0, &TrevodT4ord_05},
	0x06: {"billingElement", -1, 0, &TrevodT4ord_06},
	0x07: {"contextIdent", -1, 0, &TrevodT4ord_07},
	0x01: {"CodePlateforme", -1, 1, nil},
	0x09: {"TerminalConnection", -1, 0, &TrevodT4ord_09},
	0x0c: {"codPromo", -1, 1, nil},
	0x0d: {"libPromo", -1, 1, nil},
	0x7e: {"ajout", -1, -1, &TrevodT4ord_7e},
}

// TrevodT2
var TrevodT2 = map[int]defStructs.Champ{
	0x01: {"ORD", -1, 0, &TrevodT3ord},
	0x02: {"DLV", -1, 0, &TrevodT3dlv},
	0x03: {"ORP", -1, 0, &TrevodT3orp},
}

// Trevod
var Trevod = map[int]defStructs.Champ{
	0x4d: {"Trevod", -1, 0, &TrevodT2},
}

// Drevod
var Drevod = map[int]defStructs.TabRechercheTypeDc{
	0: {"Trevod", 0xffff0000, 0x804d0000, 0, &Trevod},
}
