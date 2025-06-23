package structuresEccCSV

import "local/structures/defStructs"

// TmmdmWha
var TmmdmWhaDc = map[int]defStructs.Champ{
	0:  {"OrderId", -1, -1, nil},
	1:  {"TransactionId", -1, -1, nil},
	2:  {"date", -1, -1, nil},
	3:  {"AID", -1, -1, nil},
	4:  {"ND", -1, -1, nil},
	5:  {"ProductDescription", -1, -1, nil},
	6:  {"ProductFormat", -1, -1, nil},
	7:  {"ProductId", -1, -1, nil},
	8:  {"PartnerId", -1, -1, nil},
	9:  {"PartnerName", -1, -1, nil},
	10: {"TrxGrossAmount", -1, -1, nil},
	11: {"TrxNetAmount", -1, -1, nil},
	12: {"Currency", -1, -1, nil},
	13: {"CatalogAmount", -1, -1, nil},
	14: {"PaymentMethod", -1, -1, nil},
	15: {"Timezone", -1, -1, nil},
	16: {"ServiceLabel", -1, -1, nil},
	17: {"TypeAccess", -1, -1, nil},
	18: {"SGTQS", -1, -1, nil},
}

// DmmdmWha
var DmmdmWha = map[int]defStructs.TabRechercheTypeDc{
	0: {"TmmdmWhaDc", 0x00000000, 0x00000000, -1, &TmmdmWhaDc},
}
