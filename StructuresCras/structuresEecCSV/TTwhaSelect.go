package structuresEecCSV

import "local/structures/defStructs"

// TTwhaSelect
var TTwhaSelect = map[int]defStructs.TabRechercheTypeDc{
	0: {"TwhaSelect", 0x0000ffff, 0x00004632, 0x3b, &TwhaSelect},
	1: {"Entete", 0xffff0000, 0x49640000, 0x3b, nil},
}

var TwhaSelect = map[int]defStructs.Champ{
	0:  {"Identifiant_source", -1, 0, nil},
	1:  {"AID_NBH", -1, 0, nil},
	2:  {"AID_BSS", -1, 0, nil},
	3:  {"Code_prestation", -1, 0, nil},
	4:  {"Type_de_prestation", -1, 0, nil},
	5:  {"Code_debit_credit", -1, 0, nil},
	6:  {"Date_effet_de_la_prestation", -1, 0, nil},
	7:  {"Montant_HT", -1, 0, nil},
	8:  {"Code_devise", -1, 0, nil},
	9:  {"Heure_de_la_prestation", -1, 0, nil},
	10: {"Date_de_debut_abonnement", -1, 0, nil},
	11: {"Date_fin_abonnement", -1, 0, nil},
	12: {"timezone", -1, 0, nil},
	13: {"Indicateur_operateur_du_service", -1, 0, nil},
	14: {"Fournisseur", -1, 0, nil},
	15: {"Libelle", -1, 0, nil},
	16: {"Duree_appel_au_service", -1, 0, nil},
	17: {"Categorie_du_service", -1, 0, nil},
	18: {"Identifiant_du_ticket", -1, 0, nil},
	19: {"Login", -1, 0, nil},
	20: {"Montant_unitaire", -1, 0, nil},
	21: {"Quantite", -1, 0, nil},
}
