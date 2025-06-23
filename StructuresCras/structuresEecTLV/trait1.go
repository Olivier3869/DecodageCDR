package structuresEecTLV

import "local/structures/defStructs"

// trait1
var Trait1 = map[int]defStructs.TabRechercheTypeDc{
	0:  {"Txx8he", 0xffff0000, 0x03150000, -1, &Txx8he},
	1:  {"Txx4he", 0xffff0000, 0x00140000, -1, &Txx4he},
	2:  {"Article", 0xff00ff00, 0x00000200, -1, &Article},
	3:  {"Article", 0xff00ff00, 0x00001800, -1, &Article},
	4:  {"Article", 0xff00ff00, 0x02000200, -1, &Article},
	5:  {"Article", 0xff00ff00, 0x02001800, -1, &Article},
	6:  {"Article", 0xff00ff00, 0x05000200, -1, &Article},
	7:  {"Article", 0xff00ff00, 0x05001800, -1, &Article},
	8:  {"Article", 0xff00ff00, 0x0e000200, -1, &Article},
	9:  {"Article", 0xff00ff00, 0x0e001800, -1, &Article},
	10: {"Article", 0xff00ff00, 0x0c000500, -1, &Article},
	11: {"Article", 0xff00ff00, 0x19000200, -1, &Article},
	12: {"Article", 0xff00ff00, 0x19001800, -1, &Article},
	13: {"Article", 0xff00ff00, 0x50000200, -1, &Article},
	14: {"Evenement", 0xff00ff00, 0x0c000100, -1, &Evenement},
	15: {"Evenement", 0xff00ff00, 0x0c000700, -1, &Evenement},
	16: {"Evenement", 0xff00ff00, 0x0c000800, -1, &Evenement},
	17: {"Evenement", 0xff000000, 0x11000000, -1, &Evenement},
	18: {"Evenement", 0xff000000, 0x12000000, -1, &Evenement},
	19: {"Evenement", 0xff000000, 0x13000000, -1, &Evenement},
	20: {"Evenement", 0xff000000, 0x14000000, -1, &Evenement},
	21: {"Evenement", 0xff000000, 0x15000000, -1, &Evenement},
	22: {"Evenement", 0xff000000, 0x16000000, -1, &Evenement},
	23: {"Evenement", 0xff000000, 0x17000000, -1, &Evenement},
}

var Txx8he = map[int]defStructs.Champ{
	0x03: {"Txx8he", -1, 0, &Txx8heChamps},
}

var Txx8heChamps = map[int]defStructs.Champ{
	0: {"BLOC", 3, 0, nil},
	1: {"DATE", 6, 0, nil},
	2: {"CAA", 6, 0, nil},
	3: {"EMI", 2, 0, nil},
	4: {"ART", 1, 0, nil},
	5: {"COD", 1, 0, nil},
	6: {"MINE", 1, 0, nil},
	7: {"APPLI", 1, 0, nil},
}

var Txx4he = map[int]defStructs.Champ{
	0x00: {"Txx4he", 0, 0, &Txx4heChamps},
}

var Txx4heChamps = map[int]defStructs.Champ{
	0: {"BLOC", 3, 0, nil},
	1: {"DATE", 6, 0, nil},
	2: {"CAA", 6, 0, nil},
	3: {"EMI", 2, 0, nil},
	4: {"ART", 1, 0, nil},
	5: {"NBM", 2, 0, nil},
}

var Evenement = map[int]defStructs.Champ{
	0x0c: {"evenement", -1, 0, &EvenementT2},
}

var EvenementT2 = map[int]defStructs.Champ{
	0x01: {"ttx", -1, 0, &EvTtxT3},
	0x05: {"cc", -1, 0, &EvCcT3},
	0x07: {"sy", -1, 0, &EvSyT3},
	0x08: {"Ev", -1, 0, &EvEvT3},
}

var EvTtxT3 = map[int]defStructs.Champ{
	0x3a: {"nart", -1, 0, &EvTtxNart},
	0x36: {"icar", -1, 0, &EvIcar},
}

var EvCcT3 = map[int]defStructs.Champ{
	0x36: {"icar", -1, 0, &EvIcar},
	0x40: {"idus", -1, 0, &EvIdus},
	0x0e: {"caco", -1, 0, &EvCaco},
}

var EvSyT3 = map[int]defStructs.Champ{
	0x36: {"icar", -1, 0, &EvIcar},
	0x40: {"idus", -1, 0, &EvIdus},
	0x1f: {"g1", -1, 0, &EvGx},
	0x20: {"g2", -1, 0, &EvGx},
	0x21: {"g3", -1, 0, &EvGx},
	0x23: {"g5", -1, 0, &EvGx},
	0x24: {"g6", -1, 0, &EvGx},
	0x25: {"g7", -1, 0, &EvGx},
	0x26: {"g8", -1, 0, &EvGx},
}

var EvEvT3 = map[int]defStructs.Champ{
	0x36: {"icar", -1, 0, &EvIcar},
	0x11: {"ev1", -1, 0, &EvEv1},
	0x12: {"ev2", -1, 0, &EvEv2},
	0x13: {"ev3", -1, 0, &EvEv3},
	0x14: {"ev4", -1, 0, &EvEv4},
	0x15: {"ev5", -1, 0, &EvEv5},
	0x16: {"ev6", -1, 0, &EvEv6},
	0x17: {"ev7", -1, 0, &EvEv7},
}

var EvTtxNart = map[int]defStructs.Champ{
	0: {"Nart", 1, 0, nil},
}

var EvIcar = map[int]defStructs.Champ{
	0: {"ICAR", 5, 0, nil},
}

var EvIdus = map[int]defStructs.Champ{
	0: {"INF_IDUS", 1, 0, nil},
	1: {"NDR", -1, 0, nil},
}

var EvCaco = map[int]defStructs.Champ{
	0: {"typectrl", 1, 0, nil},
	1: {"CTX", 1, 0, nil},
	2: {"DUREE", 3, 0, nil},
	3: {"TAXE", 3, 0, nil},
}

var EvGx = map[int]defStructs.Champ{
	0: {"CAX", 1, 0, nil},
	1: {"CADET", 1, 0, nil},
	2: {"NDGP", 5, 0, nil},
	3: {"TAXE", 3, 0, nil},
}

var EvEv1 = map[int]defStructs.Champ{
	0: {"NUC", 1, 0, nil},
	1: {"NUT", 1, 0, nil},
	2: {"tarif", 1, 0, nil},
	3: {"jour", 1, 0, nil},
	4: {"CAA", 6, 0, nil},
}

var EvEv2 = map[int]defStructs.Champ{
	0: {"ICAR", 5, 0, nil},
	1: {"CAA", 6, 0, nil},
}

var EvEv3 = map[int]defStructs.Champ{
	0: {"etat", 1, 0, nil},
	1: {"CAA", 6, 0, nil},
}

var EvEv4 = map[int]defStructs.Champ{
	0: {"BLOC", 3, 0, nil},
	1: {"CAA", 6, 0, nil},
}

var EvEv5 = map[int]defStructs.Champ{
	0: {"CAA", 6, 0, nil},
}

var EvEv6 = map[int]defStructs.Champ{
	0: {"BLOC", 3, 0, nil},
	1: {"CAA", 6, 0, nil},
}

var EvEv7 = map[int]defStructs.Champ{
	0: {"BLOC", 3, 0, nil},
	1: {"CAA", 6, 0, nil},
}

var Article = map[int]defStructs.Champ{
	0x00: {"Analogique", 0, 0, &Tta6T2},
	0x02: {"Numeris", -1, 0, &Tns6T2},
	0x05: {"Radiomobile", -1, 0, &Trm6T2},
	0x0e: {"Faisceau", -1, 0, &Tfe6T2},
	0x19: {"Centrex", -1, 0, &CentrexT2},
	0x50: {"ZTAanalog", -1, 0, &ZTAta6T2},
	0x52: {"Zanumeris", -1, 0, &ZTAns6T2},
}

var Tta6T2 = map[int]defStructs.Champ{
	0x02: {"Dc", -1, 0, &Tta6dcT3},
	0x18: {"Dea", -1, 0, &Tta6deaT3},
}

var Tns6T2 = map[int]defStructs.Champ{
	0x02: {"Dc", -1, 0, &Tns6dcT3},
	0x18: {"Dea", -1, 0, &Tns6deaT3},
}

var Trm6T2 = map[int]defStructs.Champ{
	0x02: {"Dc", -1, 0, &Trm6dcT3},
	0x18: {"Dea", -1, 0, &Trm6deaT3},
}

var Tfe6T2 = map[int]defStructs.Champ{
	0x02: {"Dc", -1, 0, &Tfe6dcT3},
	0x18: {"Dea", -1, 0, &Tch6deaT3},
}

var CentrexT2 = map[int]defStructs.Champ{
	0x02: {"Dc", -1, 0, &CentrexdcT3},
	0x18: {"Dea", -1, 0, &CentrexdeaT3},
}

var ZTAta6T2 = map[int]defStructs.Champ{
	0x02: {"Dc", -1, 0, &ZTAta6dcT3},
}

var ZTAns6T2 = map[int]defStructs.Champ{
	0x02: {"Dc", -1, 0, &ZTAns6dcT3},
}

var Tta6dcT3 = map[int]defStructs.Champ{
	0x03: {"G3", -1, 0, &T3g3},
	0x40: {"IDUS", -1, 0, &T3Idus},
	0x43: {"dede", -1, 0, &T3Dede},
	0x1a: {"link", -1, 0, &T3Link},
	0x1d: {"calc", -1, 0, &T3Calc},
	0x45: {"idav", -1, 0, &T3IDAV},
	0x44: {"coca", -1, 0, &T3Coca},
	0x46: {"cocav", -1, 0, &T3Cocav},
	0x42: {"nos", -1, 0, &T3Nos},
	0x38: {"indx", -1, 0, &T3Indx},
	0x08: {"cap1", -1, 0, &T3Cap1},
	0x06: {"cab", -1, 0, &T3Cab},
	0x3c: {"sidr", -1, 0, nil},
	0x3b: {"code", -1, 0, &T3Code},
	0x3f: {"inv", -1, 0, &T3Inv},
	0x3d: {"cle", -1, 0, nil},
	0x18: {"cler", -1, 0, &T3Cler},
	0x1b: {"sel", -1, 0, &T3Sel},
	0x1c: {"son", -1, 0, &T3Son},
	0x1f: {"nfsce", -1, 0, &T3Nfsce},
	0x20: {"nfscs", -1, 0, &T3Nfscs},
	0x21: {"ur", -1, 0, &T3Ur},
	0x7e: {"ce", -1, 0, nil},
}

var Tta6deaT3 = map[int]defStructs.Champ{
	0x04: {"G4", -1, 0, &T3g4},
	0x40: {"IDUS", -1, 0, &T3Idus},
	0x43: {"dede", -1, 0, &T3Dede},
	0x1a: {"link", -1, 0, &T3Link},
	0x1d: {"calc", -1, 0, &T3Calc},
	0x45: {"idav", -1, 0, &T3IDAV},
	0x44: {"coca", -1, 0, &T3Coca},
	0x46: {"cocav", -1, 0, &T3Cocav},
	0x42: {"nos", -1, 0, &T3Nos},
	0x3d: {"cle", -1, 0, nil},
	0x18: {"cler", -1, 0, &T3Cler},
	0x1b: {"sel", -1, 0, &T3Sel},
	0x1c: {"son", -1, 0, &T3Son},
	0x1f: {"nfsce", -1, 0, &T3Nfsce},
	0x20: {"nfscs", -1, 0, &T3Nfscs},
	0x21: {"ur", -1, 0, &T3Ur},
	0x7e: {"ce", -1, 0, nil},
	0x47: {"locus", -1, 0, &T3Locus},
	0x32: {"dur", -1, 0, &T3Dur},
	0x30: {"iprv", -1, 0, &T3Iprv},
	0x31: {"idrv", -1, 0, &T3Idrv},
	0x19: {"nprv", -1, 0, &T3Nbrv},
	0x07: {"ut", -1, 0, &T3Ut},
}

var Tns6dcT3 = map[int]defStructs.Champ{
	0x04: {"G4", -1, 0, &T3g4},
	0x40: {"IDUS", -1, 0, &T3Idus},
	0x43: {"dede", -1, 0, &T3Dede},
	0x1a: {"link", -1, 0, &T3Link},
	0x1d: {"calc", -1, 0, &T3Calc},
	0x45: {"idav", -1, 0, &T3IDAV},
	0x44: {"coca", -1, 0, &T3Coca},
	0x46: {"cocav", -1, 0, &T3Cocav},
	0x42: {"nos", -1, 0, &T3Nos},
	0x38: {"indx", -1, 0, &T3Indx},
	0x0a: {"cap2", -1, 0, &T3Cap2},
	0x06: {"cab", -1, 0, &T3Cab},
	0x3c: {"sidr", -1, 0, nil},
	0x3b: {"code", -1, 0, &T3Code},
	0x3f: {"inv", -1, 0, &T3Inv},
	0x3d: {"cle", -1, 0, nil},
	0x18: {"cler", -1, 0, &T3Cler},
	0x1b: {"sel", -1, 0, &T3Sel},
	0x1c: {"son", -1, 0, &T3Son},
	0x1f: {"nfsce", -1, 0, &T3Nfsce},
	0x20: {"nfscs", -1, 0, &T3Nfscs},
	0x21: {"ur", -1, 0, &T3Ur},
	0x7e: {"ce", -1, 0, nil},
	0x3e: {"ttx", -1, 0, nil},
	0x0c: {"cha", -1, 0, &T3Cha},
	0x0d: {"chap", -1, 0, &T3Chp},
}

var Tns6deaT3 = map[int]defStructs.Champ{
	0x04: {"G4", -1, 0, &T3g4},
	0x40: {"IDUS", -1, 0, &T3Idus},
	0x43: {"dede", -1, 0, &T3Dede},
	0x1a: {"link", -1, 0, &T3Link},
	0x1d: {"calc", -1, 0, &T3Calc},
	0x45: {"idav", -1, 0, &T3IDAV},
	0x44: {"coca", -1, 0, &T3Coca},
	0x46: {"cocav", -1, 0, &T3Cocav},
	0x42: {"nos", -1, 0, &T3Nos},
	0x3d: {"cle", -1, 0, nil},
	0x18: {"cler", -1, 0, &T3Cler},
	0x1b: {"sel", -1, 0, &T3Sel},
	0x1c: {"son", -1, 0, &T3Son},
	0x1f: {"nfsce", -1, 0, &T3Nfsce},
	0x20: {"nfscs", -1, 0, &T3Nfscs},
	0x21: {"ur", -1, 0, &T3Ur},
	0x47: {"locus", -1, 0, &T3Locus},
	0x32: {"dur", -1, 0, &T3Dur},
	0x30: {"iprv", -1, 0, &T3Iprv},
	0x31: {"idrv", -1, 0, &T3Idrv},
	0x19: {"nprv", -1, 0, &T3Nbrv},
	0x07: {"ut", -1, 0, &T3Ut},
}

var Trm6dcT3 = map[int]defStructs.Champ{
	0x04: {"G4", -1, 0, &T3g4},
	0x40: {"IDUS", -1, 0, &T3Idus},
	0x43: {"dede", -1, 0, &T3Dede},
	0x1a: {"link", -1, 0, &T3Link},
	0x1d: {"calc", -1, 0, &T3Calc},
	0x45: {"idav", -1, 0, &T3IDAV},
	0x44: {"coca", -1, 0, &T3Coca},
	0x46: {"cocav", -1, 0, &T3Cocav},
	0x0a: {"cap2", -1, 0, &T3Cap2},
	0x06: {"cab", -1, 0, &T3Cab},
	0x3d: {"cle", -1, 0, nil},
	0x18: {"cler", -1, 0, &T3Cler},
	0x1b: {"sel", -1, 0, &T3Sel},
	0x1c: {"son", -1, 0, &T3Son},
	0x1f: {"nfsce", -1, 0, &T3Nfsce},
	0x20: {"nfscs", -1, 0, &T3Nfscs},
	0x0c: {"cha", -1, 0, &T3Cha},
	0x47: {"locus", -1, 0, &T3Locus},
	0x41: {"icom", -1, 0, &T3Icom},
	0x30: {"iprv", -1, 0, &T3Iprv},
	0x31: {"idrv", -1, 0, &T3Idrv},
	0x19: {"nprv", -1, 0, &T3Nbrv},
	0x48: {"idsv", -1, 0, &T3Idsv},
}

var Trm6deaT3 = map[int]defStructs.Champ{
	0x04: {"G4", -1, 0, &T3g4},
	0x40: {"IDUS", -1, 0, &T3Idus},
	0x43: {"dede", -1, 0, &T3Dede},
	0x1a: {"link", -1, 0, &T3Link},
	0x1d: {"calc", -1, 0, &T3Calc},
	0x45: {"idav", -1, 0, &T3IDAV},
	0x44: {"coca", -1, 0, &T3Coca},
	0x46: {"cocav", -1, 0, &T3Cocav},
	0x42: {"nos", -1, 0, &T3Nos},
	0x3d: {"cle", -1, 0, nil},
	0x18: {"cler", -1, 0, &T3Cler},
	0x1b: {"sel", -1, 0, &T3Sel},
	0x1c: {"son", -1, 0, &T3Son},
	0x1f: {"nfsce", -1, 0, &T3Nfsce},
	0x20: {"nfscs", -1, 0, &T3Nfscs},
	0x47: {"locus", -1, 0, &T3Locus},
	0x32: {"dur", -1, 0, &T3Dur},
	0x30: {"iprv", -1, 0, &T3Iprv},
	0x31: {"idrv", -1, 0, &T3Idrv},
	0x19: {"nprv", -1, 0, &T3Nbrv},
	0x07: {"ut", -1, 0, &T3Ut},
	0x41: {"icom", -1, 0, &T3Icom},
	0x48: {"idsv", -1, 0, &T3Idsv},
}

var Tfe6dcT3 = map[int]defStructs.Champ{
	0x04: {"G4", -1, 0, &T3g4},
	0x40: {"IDUS", -1, 0, &T3Idus},
	0x43: {"dede", -1, 0, &T3Dede},
	0x1a: {"link", -1, 0, &T3Link},
	0x1d: {"calc", -1, 0, &T3Calc},
	0x45: {"idav", -1, 0, &T3IDAV},
	0x44: {"coca", -1, 0, &T3Coca},
	0x46: {"cocav", -1, 0, &T3Cocav},
	0x0a: {"cap2", -1, 0, &T3Cap2},
	0x06: {"cab", -1, 0, &T3Cab},
	0x3d: {"cle", -1, 0, nil},
	0x18: {"cler", -1, 0, &T3Cler},
	0x1b: {"sel", -1, 0, &T3Sel},
	0x1c: {"son", -1, 0, &T3Son},
	0x1f: {"nfsce", -1, 0, &T3Nfsce},
	0x20: {"nfscs", -1, 0, &T3Nfscs},
	0x0c: {"cha", -1, 0, &T3Cha},
	0x47: {"locus", -1, 0, &T3Locus},
	0x41: {"icom", -1, 0, &T3Icom},
	0x30: {"iprv", -1, 0, &T3Iprv},
	0x31: {"idrv", -1, 0, &T3Idrv},
	0x19: {"nprv", -1, 0, &T3Nbrv},
	0x48: {"idsv", -1, 0, &T3Idsv},
	0x3c: {"sidr", -1, 0, &T3Idsv},
}

var Tch6deaT3 = map[int]defStructs.Champ{
	0x04: {"G4", -1, 0, &T3g4},
	0x40: {"IDUS", -1, 0, &T3Idus},
	0x43: {"dede", -1, 0, &T3Dede},
	0x1a: {"link", -1, 0, &T3Link},
	0x1d: {"calc", -1, 0, &T3Calc},
	0x45: {"idav", -1, 0, &T3IDAV},
	0x44: {"coca", -1, 0, &T3Coca},
	0x46: {"cocav", -1, 0, &T3Cocav},
	0x42: {"nos", -1, 0, &T3Nos},
	0x3d: {"cle", -1, 0, nil},
	0x18: {"cler", -1, 0, &T3Cler},
	0x1b: {"sel", -1, 0, &T3Sel},
	0x1c: {"son", -1, 0, &T3Son},
	0x1f: {"nfsce", -1, 0, &T3Nfsce},
	0x20: {"nfscs", -1, 0, &T3Nfscs},
	0x47: {"locus", -1, 0, &T3Locus},
	0x32: {"dur", -1, 0, &T3Dur},
	0x30: {"iprv", -1, 0, &T3Iprv},
	0x31: {"idrv", -1, 0, &T3Idrv},
	0x19: {"nprv", -1, 0, &T3Nbrv},
	0x07: {"ut", -1, 0, &T3Ut},
	0x41: {"icom", -1, 0, &T3Icom},
	0x48: {"idsv", -1, 0, &T3Idsv},
}

var CentrexdcT3 = map[int]defStructs.Champ{
	0x03: {"G3", -1, 0, &T3g3},
	0x40: {"IDUS", -1, 0, &T3Idus},
	0x43: {"dede", -1, 0, &T3Dede},
	0x1a: {"link", -1, 0, &T3Link},
	0x1d: {"calc", -1, 0, &T3Calc},
	0x45: {"idav", -1, 0, &T3IDAV},
	0x44: {"coca", -1, 0, &T3Coca},
	0x46: {"cocav", -1, 0, &T3Cocav},
	0x42: {"nos", -1, 0, &T3Nos},
	0x38: {"indx", -1, 0, &T3Indx},
	0x08: {"cap1", -1, 0, &T3Cap1},
	0x06: {"cab", -1, 0, &T3Cab},
	0x3c: {"sidr", -1, 0, nil},
	0x3b: {"code", -1, 0, &T3Code},
	0x3f: {"inv", -1, 0, &T3Inv},
	0x3d: {"cle", -1, 0, nil},
	0x18: {"cler", -1, 0, &T3Cler},
	0x1b: {"sel", -1, 0, &T3Sel},
	0x1c: {"son", -1, 0, &T3Son},
	0x1f: {"nfsce", -1, 0, &T3Nfsce},
	0x20: {"nfscs", -1, 0, &T3Nfscs},
	0x21: {"ur", -1, 0, &T3Ur},
	0x7e: {"ce", -1, 0, nil},
}

var CentrexdeaT3 = map[int]defStructs.Champ{
	0x04: {"G4", -1, 0, &T3g4},
	0x40: {"IDUS", -1, 0, &T3Idus},
	0x43: {"dede", -1, 0, &T3Dede},
	0x1a: {"link", -1, 0, &T3Link},
	0x1d: {"calc", -1, 0, &T3Calc},
	0x45: {"idav", -1, 0, &T3IDAV},
	0x44: {"coca", -1, 0, &T3Coca},
	0x46: {"cocav", -1, 0, &T3Cocav},
	0x42: {"nos", -1, 0, &T3Nos},
	0x3d: {"cle", -1, 0, nil},
	0x18: {"cler", -1, 0, &T3Cler},
	0x1b: {"sel", -1, 0, &T3Sel},
	0x1c: {"son", -1, 0, &T3Son},
	0x1f: {"nfsce", -1, 0, &T3Nfsce},
	0x20: {"nfscs", -1, 0, &T3Nfscs},
	0x21: {"ur", -1, 0, &T3Ur},
	0x47: {"locus", -1, 0, &T3Locus},
	0x32: {"dur", -1, 0, &T3Dur},
	0x30: {"iprv", -1, 0, &T3Iprv},
	0x31: {"idrv", -1, 0, &T3Idrv},
	0x19: {"nprv", -1, 0, &T3Nbrv},
	0x07: {"ut", -1, 0, &T3Ut},
}

var ZTAta6dcT3 = map[int]defStructs.Champ{
	0x03: {"G3", -1, 0, &T3g3},
	0x40: {"IDUS", -1, 0, &T3Idus},
	0x43: {"dede", -1, 0, &T3Dede},
	0x1a: {"link", -1, 0, &T3Link},
	0x1d: {"calc", -1, 0, &T3Calc},
	0x45: {"idav", -1, 0, &T3IDAV},
	0x44: {"coca", -1, 0, &T3Coca},
	0x46: {"cocav", -1, 0, &T3Cocav},
	0x42: {"nos", -1, 0, &T3Nos},
	0x38: {"indx", -1, 0, &T3Indx},
	0x08: {"cap1", -1, 0, &T3Cap1},
	0x06: {"cab", -1, 0, &T3Cab},
	0x3c: {"sidr", -1, 0, nil},
	0x3b: {"code", -1, 0, &T3Code},
	0x3f: {"inv", -1, 0, &T3Inv},
	0x3d: {"cle", -1, 0, nil},
	0x18: {"cler", -1, 0, &T3Cler},
	0x1b: {"sel", -1, 0, &T3Sel},
	0x1c: {"son", -1, 0, &T3Son},
	0x1f: {"nfsce", -1, 0, &T3Nfsce},
	0x20: {"nfscs", -1, 0, &T3Nfscs},
	0x21: {"ur", -1, 0, &T3Ur},
	0x7e: {"ce", -1, 0, nil},
}

var ZTAns6dcT3 = map[int]defStructs.Champ{
	0x04: {"G4", -1, 0, &T3g4},
	0x40: {"IDUS", -1, 0, &T3Idus},
	0x43: {"dede", -1, 0, &T3Dede},
	0x1a: {"link", -1, 0, &T3Link},
	0x1d: {"calc", -1, 0, &T3Calc},
	0x45: {"idav", -1, 0, &T3IDAV},
	0x44: {"coca", -1, 0, &T3Coca},
	0x46: {"cocav", -1, 0, &T3Cocav},
	0x42: {"nos", -1, 0, &T3Nos},
	0x38: {"indx", -1, 0, &T3Indx},
	0x0a: {"cap2", -1, 0, &T3Cap2},
	0x06: {"cab", -1, 0, &T3Cab},
	0x3c: {"sidr", -1, 0, &T3Idsv},
	0x3b: {"code", -1, 0, &T3Code},
	0x3f: {"inv", -1, 0, &T3Inv},
	0x3d: {"cle", -1, 0, nil},
	0x18: {"cler", -1, 0, &T3Cler},
	0x1b: {"sel", -1, 0, &T3Sel},
	0x1c: {"son", -1, 0, &T3Son},
	0x1f: {"nfsce", -1, 0, &T3Nfsce},
	0x20: {"nfscs", -1, 0, &T3Nfscs},
	0x21: {"ur", -1, 0, &T3Ur},
	0x7e: {"ce", -1, 0, nil},
	0x3e: {"ttx", -1, 0, nil},
	0x0c: {"cha", -1, 0, &T3Cha},
	0x0d: {"chap", -1, 0, &T3Chp},
}

var T3g3 = map[int]defStructs.Champ{
	0: {"STY", 1, 0, nil},
	1: {"SCT", 1, 0, nil},
	2: {"ICAR", 5, 0, nil},
	3: {"INEF", 1, 0, nil},
}

var T3g4 = map[int]defStructs.Champ{
	0: {"STY", 1, 0, nil},
	1: {"SCT", 1, 0, nil},
	2: {"ICAR", 5, 0, nil},
	3: {"INEF", 1, 0, nil},
	4: {"REF", -1, 0, nil},
}

var T3Idus = map[int]defStructs.Champ{
	0: {"INF_IDUS", 1, 0, nil},
	1: {"NDR", -1, 0, nil},
}

var T3Dede = map[int]defStructs.Champ{
	0: {"INF_dede", 1, 0, nil},
	1: {"NDE", -1, 0, nil},
}

var T3Link = map[int]defStructs.Champ{
	0: {"APPEL", 3, 0, nil},
	1: {"NŒUD", -1, 0, nil},
}

var T3Calc = map[int]defStructs.Champ{
	0: {"APPEL", 1, 0, nil},
}

var T3IDAV = map[int]defStructs.Champ{
	0: {"INF_IDAV", 1, 0, nil},
	1: {"NUM_IDAV", -1, 0, nil},
}

var T3Coca = map[int]defStructs.Champ{
	0: {"BETA", -1, 0, nil},
}

var T3Indx = map[int]defStructs.Champ{
	0: {"BETAV", -1, 0, nil},
}

var T3Nos = map[int]defStructs.Champ{
	0: {"INF_NOS", 1, 0, nil},
	1: {"NUMNOS", -1, 0, nil},
}

var T3Cocav = map[int]defStructs.Champ{
	0: {"INDEX", -1, 0, nil},
}

var T3Cap1 = map[int]defStructs.Champ{
	0: {"CAX", 1, 0, nil},
	1: {"TAP", 1, 0, nil},
	2: {"DUREE", 2, 0, nil},
	3: {"TAXE", 2, 0, nil},
}

var T3Cap2 = map[int]defStructs.Champ{
	0: {"CAX", 1, 0, nil},
	1: {"TAP", 1, 0, nil},
	2: {"DUREE", 2, 0, nil},
	3: {"TAXE", 2, 0, nil},
	4: {"SUP", 1, 0, nil},
}

var T3Cab = map[int]defStructs.Champ{
	0: {"CAX", 1, 0, nil},
	1: {"TAP", 1, 0, nil},
	2: {"DUREE", 3, 0, nil},
	3: {"TAXE", 3, 0, nil},
	4: {"SUP", 1, 0, nil},
}

var T3Code = map[int]defStructs.Champ{
	0: {"CODE", -1, 0, nil},
}

var T3Inv = map[int]defStructs.Champ{
	0: {"INVOC", -1, 0, nil},
}

var T3Cle = map[int]defStructs.Champ{
	0: {"CLE1", 1, 0, nil},
	1: {"CLE2", 1, 0, nil},
	2: {"CLE3", 1, 0, nil},
	3: {"CLE4", 1, 0, nil},
	4: {"CLE5", 1, 0, nil},
	5: {"CLE6", 1, 0, nil},
	6: {"CLE7", 1, 0, nil},
	7: {"CLE8", -1, 0, nil},
}

var T3Cler = map[int]defStructs.Champ{
	0: {"REP", -1, 0, nil},
}

var T3Sel = map[int]defStructs.Champ{
	0: {"DUSEL", -1, 0, nil},
}

var T3Son = map[int]defStructs.Champ{
	0: {"DUSON", -1, 0, nil},
}

var T3Nfsce = map[int]defStructs.Champ{
	0: {"nfsce", 5, 0, nil},
	1: {"afcte", -1, 0, nil},
}

var T3Nfscs = map[int]defStructs.Champ{
	0: {"nfscs", 5, 0, nil},
	1: {"afcts", -1, 0, nil},
}

var T3Ur = map[int]defStructs.Champ{
	0: {"typur", 1, 0, nil},
	1: {"ur", -1, 0, nil},
}

var T3Locus = map[int]defStructs.Champ{
	0: {"INFLOCUS", 1, 0, nil},
	1: {"NUMLOCUS", -1, 0, nil},
}

var T3Dur = map[int]defStructs.Champ{
	0: {"DUREE", -1, 0, nil},
}

var T3Iprv = map[int]defStructs.Champ{
	0: {"INFIPRVS", 1, 0, nil},
	1: {"NUMIPRV", -1, 0, nil},
}

var T3Idrv = map[int]defStructs.Champ{
	0: {"INFIDRVS", 1, 0, nil},
	1: {"NUMIDRV", -1, 0, nil},
}

var T3Nbrv = map[int]defStructs.Champ{
	0: {"NRENV", -1, 0, nil},
}

var T3Ut = map[int]defStructs.Champ{
	0: {"TAXE", -1, 0, nil},
}

var T3Cha = map[int]defStructs.Champ{
	0: {"MES1", 2, 0, nil},
}

var T3Chp = map[int]defStructs.Champ{
	0: {"MENV", 1, 0, nil},
	1: {"MERE", 1, 0, nil},
	2: {"ECHO", 3, 0, nil},
}

var T3Icom = map[int]defStructs.Champ{
	0: {"INF_ICOM", 1, 0, nil},
	1: {"NUM_ICOM", -1, 0, nil},
}

var T3Idsv = map[int]defStructs.Champ{
	0: {"IDSV", -1, 0, nil},
}
