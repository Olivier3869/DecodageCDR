package main

import (
	"encoding/binary"
	"local/structures/defStructs"
)

// rechercheTypeCRA recherche le type de CRA en fonction des premiers octets et d'une structure donnée.
//
// Paramètres :
// - firstFourBytes ([]byte) : un tableau d'octets représentant les premiers quatre octets à analyser.
// - structureCRA (map[int]defStructs.TabRechercheTypeDc) : une carte contenant des structures de recherche de type CRA.
//
// Retourne :
// - *defStructs.TabRechercheTypeDc : un pointeur vers la structure correspondante si un type est trouvé, sinon nil.

func rechercheTypeCRA(firstFourBytes []byte, structureCRA map[int]defStructs.TabRechercheTypeDc) *defStructs.TabRechercheTypeDc {
	// Parcours de la structure pour définir type de CRA
	for _, value := range structureCRA {
		masque := value.Masque
		cle := value.Valeur

		// Appliquer le masque aux premiers octets
		maskedValue := binary.BigEndian.Uint32(firstFourBytes) & uint32(masque)

		if maskedValue == uint32(cle) {
			return &value

		}
	}

	return nil
}
