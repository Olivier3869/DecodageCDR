package structuresEccCSV

import "local/structures/defStructs"

// TkenobiWha
var TkenobiWhaDc = map[int]defStructs.Champ{
	0:  {"operation", -1, -1, nil},
	1:  {"TransactionId", -1, -1, nil},
	2:  {"AID", -1, -1, nil},
	3:  {"ND", -1, -1, nil},
	4:  {"date", -1, -1, nil},
	5:  {"godet", -1, -1, nil},
	6:  {"godetTransactionId", -1, -1, nil},
	7:  {"productId", -1, -1, nil},
	8:  {"ProductDescription", -1, -1, nil},
	9:  {"montantTTC", -1, -1, nil},
	10: {"montantHT", -1, -1, nil},
	11: {"montantTVA", -1, -1, nil},
	12: {"timeZone", -1, -1, nil},
	13: {"SGTQS", -1, -1, nil},
	14: {"typeClient", -1, -1, nil},
	15: {"typeAcces", -1, -1, nil},
	16: {"codePrest", -1, -1, nil},
	17: {"dateLocale", -1, -1, nil},
}

// DkenobiWha
var DkenobiWha = map[int]defStructs.TabRechercheTypeDc{
	0: {"TkenobiWhaDc", 0x00000000, 0x00000000, 0x3b, &TkenobiWhaDc},
}
