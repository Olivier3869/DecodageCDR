package structuresEecASN1

import "local/structures/defStructs"

// TTsssas
var TTsssas = map[int]defStructs.TabRechercheTypeDc{

	0: {"Tsssas", 0x00ff0000, 0x00820000, 0, &Tsssas},
}

var Tsssas = map[int]defStructs.Champ{
	16: {"StopSecretRecord", -1, 0, &StopSecret},
}

var StopSecret = map[int]defStructs.Champ{
	0:  {"callingPartyAddress", -1, 1, nil},
	1:  {"initialCalledPartyAddress", -1, 1, nil},
	2:  {"stopSecretAdress", -1, 1, nil},
	3:  {"calledPartyAddress", -1, 1, nil},
	4:  {"iCID", -1, 1, nil},
	5:  {"STOPSecretAction", -1, 0, nil},
	6:  {"sessionStartTime", -1, 1, nil},
	7:  {"sessionEndTime", -1, 1, nil},
	8:  {"sessionDuration", -1, 1, nil},
	10: {"SessionWithCalledStartTime", -1, 0, nil},
	11: {"SessionWithCalledDuration", -1, 0, nil},
	15: {"SessionEstablishedStartTime", -1, 0, nil},
	16: {"sessionEtablishedDuration", -1, 0, nil},
	21: {"MUBCallerUserSSservice", -1, 0, &media},
	22: {"MUBDestUserStop", -1, 0, &media},
	23: {"MUBCallerDest", -1, 0, &sdpmedia},
	24: {"ReasingParty", -1, 0, nil},
	25: {"TerminationCause", -1, 0, nil},
}

var ReleasingParty = map[int]defStructs.Champ{
	0: {"caller", -1, 0, nil},
	1: {"called", -1, 0, nil},
	2: {"stopSecret", -1, 0, nil},
	3: {"network", -1, 0, nil},
	4: {"error", -1, 0, nil},
	5: {"OMS", -1, 0, nil},
	6: {"VM", -1, 0, nil},
	7: {"UNKNOWN", -1, 0, nil},
}

var stopSecretAction = map[int]defStructs.Champ{}

//var sdpmedia = map[int]defStructs.Champ{
//	16: {"SDPmedia", -1, 0, &sdpmedia},
//}

var sdpmedia = map[int]defStructs.Champ{
	1: {"listOfMedia", -1, 0, &sequenceMedia},
	2: {"timeOfMediaChange", -1, 0, nil},
	3: {"duration", -1, 0, nil},
}

var sequenceMedia = map[int]defStructs.Champ{
	16: {"SequenceMedia1", -1, 0, &media},
}

//
//var sdpmedia2 = map[int]defStructs.Champ{
//	1: {"listOfMedia2", -1, 0, &sequenceMedia2},
//	2: {"timeOfMediaChange", -1, 0, nil},
//	3: {"duration", -1, 0, nil},
//}
//
//var sequenceMedia2 = map[int]defStructs.Champ{
//	16: {"SequenceMedia2", -1, 0, &mediaa},
//}
//
//var sdpmedia3 = map[int]defStructs.Champ{
//	1: {"listOfMedia3", -1, 0, &sequenceMedia3},
//	2: {"timeOfMediaChange", -1, 0, nil},
//	3: {"duration", -1, 0, nil},
//}

//var sequenceMedia3 = map[int]defStructs.Champ{
//	16: {"SequenceMedia3", -1, 0, &mediaa},
//}
//
//var sdpmedia4 = map[int]defStructs.Champ{
//	1: {"listOfMedia4", -1, 0, &sequenceMedia4},
//	2: {"timeOfMediaChange", -1, 0, nil},
//	3: {"duration", -1, 0, nil},
//}
//
//var sequenceMedia4 = map[int]defStructs.Champ{
//	16: {"SequenceMedia4", -1, 0, &mediaa},
//}
//
//var sdpmedia5 = map[int]defStructs.Champ{
//	1: {"listOfMedia5", -1, 0, &sequenceMedia5},
//	2: {"timeOfMediaChange", -1, 0, nil},
//	3: {"duration", -1, 0, nil},
//}
//
//var sequenceMedia5 = map[int]defStructs.Champ{
//	16: {"SequenceMedia5", -1, 0, &mediaa},
//}
//
//var mediaa = map[int]defStructs.Champ{
//	16: {"Mediaa", -1, 0, &mediaa1},
//}

var media = map[int]defStructs.Champ{
	0: {"MediaName", -1, 0, nil},
	3: {"ListeDescriptionA", -1, 0, &listeDescription},
}

//var listeDescription = map[int]defStructs.Champ{
//	16: {"ListeDescrip", -1, 0, &listeDescripA},
//}

var listeDescription = map[int]defStructs.Champ{
	22: {"ListeStringA", -1, 0, nil},
}
