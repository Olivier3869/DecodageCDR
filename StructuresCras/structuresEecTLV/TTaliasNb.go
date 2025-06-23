package structuresEecTLV

import "local/structures/defStructs"

// T3UseCaseT3
var T3UseCaseT3 = map[int]defStructs.Champ{
	0: {"UseCase", -1, 0, nil},
}

// T3TypeT3
var T3TypeT3 = map[int]defStructs.Champ{
	0: {"Type", -1, 0, nil},
}

// T3MobileConcerneT3
var T3MobileConcerneT3 = map[int]defStructs.Champ{
	0: {"MobileConcerne", -1, 0, nil},
}

// T3CorrepondantT3
var T3CorrepondantT3 = map[int]defStructs.Champ{
	0: {"Correpondant", -1, 0, nil},
}

// T3HorodatageT3
var T3HorodatageT3 = map[int]defStructs.Champ{
	0: {"Horodatage", -1, 12, nil},
}

// T3DureeAppelT3
var T3DureeAppelT3 = map[int]defStructs.Champ{
	0: {"DureeAppel", -1, 0, nil},
}

// T3LigneRenvoiT3
var T3LigneRenvoiT3 = map[int]defStructs.Champ{
	0: {"LigneRenvoi", -1, 0, nil},
}

// T3MSISDNSecondT3
var T3MSISDNSecondT3 = map[int]defStructs.Champ{
	0: {"MSISDNSecond", -1, 0, nil},
}

// T3MSISDNTermiT3
var T3MSISDNTermiT3 = map[int]defStructs.Champ{
	0: {"MSISDNTermi", -1, 0, nil},
}

// T3PassertedIT3
var T3PassertedIT3 = map[int]defStructs.Champ{
	0: {"PassertedI", -1, 0, nil},
}

// T3GenericNumberT3
var T3GenericNumberT3 = map[int]defStructs.Champ{
	0: {"GenericNumber", -1, 0, nil},
}

// T3RedirectingNumT3
var T3RedirectingNumT3 = map[int]defStructs.Champ{
	0: {"RedirectingNum", -1, 0, nil},
}

// T3EquipmentIdentT3
var T3EquipmentIdentT3 = map[int]defStructs.Champ{
	0: {"EquipmentIdent", -1, 0, nil},
}

// T3EfficiencyIndicT3
var T3EfficiencyIndicT3 = map[int]defStructs.Champ{
	0: {"EfficiencyIndic", -1, 0, nil},
}

// T3TerminationTypeT3
var T3TerminationTypeT3 = map[int]defStructs.Champ{
	0: {"TerminationType", -1, 0, nil},
}

// T3PaniT3
var T3PaniT3 = map[int]defStructs.Champ{
	0: {"Pani", -1, 0, nil},
}

var TaliasNbT3 = map[int]defStructs.Champ{
	0x01: {"T3UseCase", -1, 3, &T3UseCaseT3},
	0x02: {"T3Type", -1, 3, &T3TypeT3},
	0x03: {"T3MobileConcerne", -1, 3, &T3MobileConcerneT3},
	0x04: {"T3Correpondant", -1, 3, &T3CorrepondantT3},
	0x05: {"T3Horodatage", -1, 3, &T3HorodatageT3},
	0x06: {"T3DureeAppel", -1, 3, &T3DureeAppelT3},
	0x07: {"T3LigneRenvoi", -1, 3, &T3LigneRenvoiT3},
	0x08: {"T3MSISDNSecond", -1, 3, &T3MSISDNSecondT3},
	0x09: {"T3MSISDNTermi", -1, 3, &T3MSISDNTermiT3},
	0x0a: {"T3PassertedI", -1, 3, &T3PassertedIT3},
	0x0b: {"T3GenericNumber", -1, 3, &T3GenericNumberT3},
	0x0c: {"T3RedirectingNum", -1, 3, &T3RedirectingNumT3},
	0x0d: {"T3EquipmentIdent", -1, 3, &T3EquipmentIdentT3},
	0x0e: {"T3EfficiencyIndic", -1, 3, &T3EfficiencyIndicT3},
	0x0f: {"T3TerminationType", -1, 3, &T3TerminationTypeT3},
	0x10: {"T3Pani", -1, 3, &T3PaniT3},
}

// TaliasNbT2
var TaliasNbT2 = map[int]defStructs.Champ{
	0x02: {"Tsvaip", -1, -1, &TaliasNbT3},
}

// TaliasNb
var TaliasNb = map[int]defStructs.Champ{
	0x64: {"Tsvaip", -1, -1, &TaliasNbT2},
}

// TTaliasNb
var TTaliasNb = map[int]defStructs.TabRechercheTypeDc{
	0: {"TTsvaip", 0xffff0000, 0x80640000, 0, &TaliasNb},
}
