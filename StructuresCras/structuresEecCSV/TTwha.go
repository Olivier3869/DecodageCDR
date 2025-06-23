package structuresEecCSV

import "local/structures/defStructs"

// TTwha
var TTwha = map[int]defStructs.TabRechercheTypeDc{
	0: {"Twha", 0xff000000, 0x37000000, 0x3b, &Twha},
	1: {"TwhaTVPrepaid", 0xffff0000, 0x43460000, 0x3b, &TwhaTVPrepaid},
	2: {"TwhaGamesOrder", 0xffff0000, 0x43450000, 0x3b, &TwhaGamesOrder},
	3: {"Entete", 0xffff0000, 0x4f700000, 0x3b, nil},
}

// Twha
var Twha = map[int]defStructs.Champ{
	0:  {"Canal", -1, -1, nil},
	1:  {"Numero_demandeur", -1, -1, nil},
	2:  {"Transcodification_du_marchand", -1, -1, nil},
	3:  {"Transaction_id", -1, -1, nil},
	4:  {"Date_de_debut", -1, -1, nil},
	5:  {"Montant_HT", -1, -1, nil},
	6:  {"Code_TVA", -1, -1, nil},
	7:  {"Identifiant_editeur", -1, -1, nil},
	8:  {"Shop_name", -1, -1, nil},
	9:  {"Type_operateur", -1, -1, nil},
	10: {"Categorie", -1, -1, nil},
	11: {"Type_achat", -1, -1, nil},
	12: {"Type_transaction", -1, -1, nil},
	13: {"Description_produit", -1, -1, nil},
	14: {"Numero_appele", -1, -1, nil},
	15: {"Evenement_facturant", -1, -1, nil},
}

var TwhaTVPrepaid = map[int]defStructs.Champ{
	0:  {"Operation", -1, 0, nil},
	1:  {"TransactionId", -1, 0, nil},
	2:  {"AID", -1, 0, nil},
	3:  {"ND", -1, 0, nil},
	4:  {"Date", -1, 0, nil},
	5:  {"Godet", -1, 0, nil},
	6:  {"GodetTrxId", -1, 0, nil},
	7:  {"productId", -1, 0, nil},
	8:  {"productDescription", -1, 0, nil},
	9:  {"trxGrossAmount", -1, 0, nil},
	10: {"trxNetAmount", -1, 0, nil},
	11: {"trxVATAmount", -1, 0, nil},
	12: {"Timezone", -1, 0, nil},
	13: {"CodeOperator", -1, 0, nil},
	14: {"productEditor", -1, 0, nil},
	15: {"productCSACategory", -1, 0, nil},
	16: {"trxType", -1, 0, nil},
	17: {"startPeriodDate", -1, 0, nil},
	18: {"endPeriodDate", -1, 0, nil},
	19: {"unitPriceNetAmount", -1, 0, nil},
	20: {"quantity", -1, 0, nil},
}

var TwhaGamesOrder = map[int]defStructs.Champ{
	0:  {"OrderId", -1, 0, nil},
	1:  {"TransactionId", -1, 0, nil},
	2:  {"Date", -1, 0, nil},
	3:  {"AID", -1, 0, nil},
	4:  {"ND", -1, 0, nil},
	5:  {"ProductDescription", -1, 0, nil},
	6:  {"ProductFormat", -1, 0, nil},
	7:  {"ProductId", -1, 0, nil},
	8:  {"PartnerId", -1, 0, nil},
	9:  {"PartnerName", -1, 0, nil},
	10: {"TrxGrossAmount", -1, 0, nil},
	11: {"TrxNetAmount", -1, 0, nil},
	12: {"Currency", -1, 0, nil},
	13: {"CatalogAmount", -1, 0, nil},
	14: {"Payment method", -1, 0, nil},
	15: {"Timezone", -1, 0, nil},
	16: {"ServiceLabel", -1, 0, nil},
}
