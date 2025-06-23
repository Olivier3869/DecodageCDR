package symp_home

import "local/structures/defStructs"

// T2sympTwha
var T2sympTwha = map[int]defStructs.Champ{
	0x01: {"TwhaContact", -1, -1, &TwhaContact},
	0x02: {"TwhaInternet", -1, -1, &TwhaInternet},
	0x03: {"TwhaSelect", -1, -1, &TwhaSelect},
}

// TwhaSelectT3ace
var TwhaSelectT3ace = map[int]defStructs.Champ{
	0x10: {"T4date", -1, 1, nil},
	0x50: {"T4sgtqs", -1, 1, nil},
	0x51: {"T4eccfact", -1, 1, nil},
}

// TwhaSelect
var TwhaSelect = map[int]defStructs.Champ{
	0x01: {"AID_NBH", -1, 1, nil},
	0x02: {"AID_BSS", -1, 1, nil},
	0x03: {"CodPres", -1, 1, nil},
	0x04: {"TypPres", -1, 1, nil},
	0x05: {"DebCre", -1, 1, nil},
	0x06: {"Date", -1, 1, nil},
	0x07: {"MontantHT", -1, 1, nil},
	0x08: {"Devise", -1, 1, nil},
	0x09: {"HeurePres", -1, 1, nil},
	0x0a: {"CodeCSA", -1, 1, nil},
	0x0b: {"TypeOpServ", -1, 1, nil},
	0x0c: {"EdiProdServ", -1, 1, nil},
	0x0e: {"startPeriod", -1, 1, nil},
	0x0f: {"endPeriod", -1, 1, nil},
	0x12: {"LibProdServ", -1, 1, nil},
	0x14: {"TimeZone", -1, 1, nil},
	0x15: {"CodEmet", -1, 1, nil},
	0x18: {"DurAppServ", -1, 1, nil},
	0x20: {"IdTicket", -1, 1, nil},
	0x21: {"Login", -1, 1, nil},
	0x22: {"PxUnitHT", -1, 1, nil},
	0x23: {"Quantite", -1, 1, nil},
	0x7e: {"ace", -1, -1, &TwhaSelectT3ace},
}

// TwhaInternetT3ace
var TwhaInternetT3ace = map[int]defStructs.Champ{
	0x10: {"T4date", -1, 1, nil},
	0x50: {"T4sgtqs", -1, 1, nil},
	0x51: {"T4eccfact", -1, 1, nil},
}

// TwhaInternet
var TwhaInternet = map[int]defStructs.Champ{
	0x01: {"AID_NBH", -1, 1, nil},
	0x02: {"AID_BSS", -1, 1, nil},
	0x03: {"CodPres", -1, 1, nil},
	0x04: {"TypPres", -1, 1, nil},
	0x05: {"DebCre", -1, 1, nil},
	0x06: {"Date", -1, 1, nil},
	0x07: {"MontantHT", -1, 1, nil},
	0x08: {"Devise", -1, 1, nil},
	0x09: {"HeurePres", -1, 1, nil},
	0x0a: {"CodeCSA", -1, 1, nil},
	0x0b: {"TypeOpServ", -1, 1, nil},
	0x0c: {"EdiProdServ", -1, 1, nil},
	0x0e: {"startPeriod", -1, 1, nil},
	0x0f: {"endPeriod", -1, 1, nil},
	0x12: {"LibProdServ", -1, 1, nil},
	0x14: {"TimeZone", -1, 1, nil},
	0x15: {"CodEmet", -1, 1, nil},
	0x18: {"DurAppServ", -1, 1, nil},
	0x20: {"IdTicket", -1, 1, nil},
	0x21: {"Login", -1, 1, nil},
	0x22: {"PxUnitHT", -1, 1, nil},
	0x23: {"Quantite", -1, 1, nil},
	0x7e: {"ace", -1, -1, &TwhaInternetT3ace},
}

// TwhaContact
var TwhaContact = map[int]defStructs.Champ{
	0x01: {"AID_NBH", -1, 1, nil},
	0x02: {"AID_BSS", -1, 1, nil},
	0x03: {"CodPres", -1, 1, nil},
	0x04: {"TypPres", -1, 1, nil},
	0x05: {"DebCre", -1, 1, nil},
	0x06: {"Date", -1, 1, nil},
	0x07: {"MontantHT", -1, 1, nil},
	0x08: {"Devise", -1, 1, nil},
	0x09: {"HeurePres", -1, 1, nil},
	0x0a: {"CodeCSA", -1, 1, nil},
	0x0b: {"TypeOpServ", -1, 1, nil},
	0x0c: {"EdiProdServ", -1, 1, nil},
	0x0e: {"startPeriod", -1, 1, nil},
	0x0f: {"endPeriod", -1, 1, nil},
	0x12: {"LibProdServ", -1, 1, nil},
	0x14: {"TimeZone", -1, 1, nil},
	0x15: {"CodEmet", -1, 1, nil},
	0x18: {"DurAppServ", -1, 1, nil},
	0x20: {"IdTicket", -1, 1, nil},
	0x21: {"Login", -1, 1, nil},
	0x22: {"PxUnitHT", -1, 1, nil},
	0x23: {"Quantite", -1, 1, nil},
	0x7e: {"ajoutPlatine", -1, -1, &T4sympAjoutPlatine},
}
