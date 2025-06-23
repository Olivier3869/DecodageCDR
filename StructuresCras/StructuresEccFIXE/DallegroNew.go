package structuresEccFIXE

import "local/structures/defStructs"

// package structures
// TallegrowTrealer
var TallegrowTrealer = map[int]defStructs.Champ{
	0: {"CODART", 2, 1, nil},
	1: {"NBANNON", 6, 1, nil},
	2: {"FILLER", 174, -1, nil},
}

// TallegrowHeader
var TallegrowHeader = map[int]defStructs.Champ{
	0: {"CODART", 2, 1, nil},
	1: {"TYPDIV", 9, 1, nil},
	2: {"DATDIV", 7, 1, nil},
	3: {"BDEDIV", 5, 1, nil},
	4: {"FILLER", 1, 1, nil},
	5: {"EPRIN", 4, 1, nil},
	6: {"FILLER", 154, -1, nil},
}

// TallegrowPresta
var TallegrowPresta = map[int]defStructs.Champ{
	0:  {"CODART", 2, 1, nil},
	1:  {"AID", 14, 1, nil},
	2:  {"INDAID", 1, 1, nil},
	3:  {"CODPRES", 7, 1, nil},
	4:  {"DEBCRED", 1, 1, nil},
	5:  {"DATPREST", 8, 1, nil},
	6:  {"CODEMET", 4, 1, nil},
	7:  {"MONTANT", 12, 1, nil},
	8:  {"DEVISE", 3, 1, nil},
	9:  {"DUREE", 7, 1, nil},
	10: {"HEURE", 4, 1, nil},
	11: {"Z1", 8, 1, nil},
	12: {"Z2", 8, 1, nil},
	13: {"Z5", 30, 1, nil},
	14: {"Z6", 70, 1, nil},
	15: {"TYPEOP", 1, 1, nil},
	16: {"FILLER", 2, -1, nil},
}

// DallegroNew
var DallegroNew = map[int]defStructs.TabRechercheTypeDc{
	0: {"TallegrowHeader", 0xffff0000, 0x45300000, 182, &TallegrowHeader},
	1: {"TallegrowTrealer", 0xffff0000, 0x45390000, 182, &TallegrowTrealer},
	2: {"TallegrowPresta", 0xffff0000, 0x45310000, 182, &TallegrowPresta},
}
