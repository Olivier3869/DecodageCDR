package structuresEecTLV

import "local/structures/defStructs"

// TTmmm
var TTmmm = map[int]defStructs.TabRechercheTypeDc{
	0: {"EnteteBloc", 0xffff0000, 0x00180000, 0, &EnteteBloc},
	1: {"Tmmm", 0xff00ff00, 0x0d000200, 0, &Tmmm},
	2: {"TmmmOL", 0xff00ff00, 0x0d000300, 0, &TmmmOL},
}

// EnteteBloc
var EnteteBloc = map[int]defStructs.Champ{
	0x00: {"Entete", -1, 1, &m_EnteteChamps},
}

// m_EnteteChamps
var m_EnteteChamps = map[int]defStructs.Champ{
	0: {"RangBloc", 4, 0, nil},
	1: {"DateHeure", 6, 8, nil},
	2: {"IdentiteEntiteEmettrice", 9, 0, nil},
	3: {"AdresseSortieEmetteur", 2, 0, nil},
	4: {"NumeroVersionPresentation", 1, 0, nil},
	5: {"NumeroBandeMagnetique", 2, 0, nil},
}

// Tmmm
var Tmmm = map[int]defStructs.Champ{
	0x0d: {"ReseauMobileGsmF1", -1, 1, &TmmmT2},
}

// TmmmT2
var TmmmT2 = map[int]defStructs.Champ{
	0x02: {"ArticleFD", -1, 0, &TmmmT3},
}

// TmmmT3
var TmmmT3 = map[int]defStructs.Champ{
	0x01: {"ParametreCompose", -1, 3, &m_T3G3},
	0x02: {"IdentiteEntiteEmettrice", -1, 3, nil},
	0x03: {"MSISDNAbonne", -1, 0, nil},
	0x04: {"IdentiteDemande", -1, 0, nil},
	0x05: {"NombrePagesFax", -1, 0, nil},
}

// m_T3G3
var m_T3G3 = map[int]defStructs.Champ{
	0: {"TypeAppel", 1, 0, nil},
	1: {"TypeEntiteEmettrice", 1, 0, nil},
	2: {"Date", 7, 8, nil},
	3: {"Duree", 3, 0, nil},
	4: {"IMSIAbonne", 8, 0, nil},
	5: {"Partition", 2, 0, nil},
	6: {"Domaine", 2, 0, nil},
}

// TmmmOL
var TmmmOL = map[int]defStructs.Champ{
	0x0d: {"ReseauMobileGSMOL", -1, 1, &TmmmOLT2},
}

// TmmmOLT2
var TmmmOLT2 = map[int]defStructs.Champ{
	0x03: {"ArtcleFD", -1, 0, &TmmmOLT3},
}

// TmmmOLT3
var TmmmOLT3 = map[int]defStructs.Champ{
	0x01: {"ParametreCompose", -1, 0, &T3G3_OL},
	0x02: {"identiteEntiteEmettrice", -1, 0, nil},
	0x03: {"MSISDNAbonne", -1, 0, nil},
	0x04: {"identiteDemande", -1, 0, nil},
	0x05: {"NombrePagesFax", -1, 0, nil},
	0x06: {"adresseMail", -1, 0, nil},
	0x07: {"portailAcces", -1, 0, nil},
	0x08: {"Volume", -1, 0, nil},
}

// T3G3_OL
var T3G3_OL = map[int]defStructs.Champ{
	0: {"TypeAppel", 1, 0, nil},
	1: {"TypeEntiteEmettrice", 1, 0, nil},
	2: {"Date", 7, 8, nil},
	3: {"Duree", 3, 0, nil},
	4: {"IMSIabonne", 8, 0, nil},
	5: {"Partition", 2, 0, nil},
	6: {"Domaine", 2, 0, nil},
}
