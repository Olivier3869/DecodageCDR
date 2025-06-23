package main

import "local/structures/defStructs"

// FindInnerTag recherche et retourne le tag interne d'une structure donnée.
//
// Paramètres :
// - structureDC (defStructs.TabRechercheTypeDc) : la structure contenant le tag à rechercher.
//
// Retourne :
// - map[int]defStructs.Champ : une carte contenant le tag interne si présent, sinon nil.
func FindInnerTag(structureDC defStructs.TabRechercheTypeDc) map[int]defStructs.Champ {
	innerTag = *structureDC.Tag
	if innerTag != nil {
		return innerTag
	} else {
		return nil
	}

}
