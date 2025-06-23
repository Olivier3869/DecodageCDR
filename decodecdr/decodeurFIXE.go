package main

import (
	"fmt"
	"local/structures/defStructs"
	"os"
	"sort"
)

// decodeFIXE décode les données FIXE et affiche les CRAs.
//
// Paramètres :
// - data ([]byte) : un tableau d'octets représentant les données à décoder.
func decodeFIXE(data []byte) {
	index := 0
	cpt = 1
	indexColl = 0

	for index < len(data) {
		//fmt.Printf("-------------------------------%v : %v\n", index, len(data))

		//  En remplacement du de la partie avant décodage
		//  Si presence entete coll ou, on avance l'index
		index += gestionDebutFichier(data[:301], index)

		cdr := new(Cdr)

		firstFourBytes := data[index : index+4]                       // Ajout apres commentaire
		StructureDC := rechercheTypeCRA(firstFourBytes, structureCRA) // Ajout apres commentaire
		// Si structure de CRA existante
		if StructureDC != nil {

			// Affichage CRA (quelle que soit l'option le CRA est toujours affiché)
			cdr.CdrByte = data[index : index+StructureDC.Longueur]
			//fmt.Printf("%v %v\n", createIdAudit(cpt), goFonctions.ByteArrayToHexString(cdr.CdrByte))

			innerTag := *StructureDC.Tag
			if innerTag != nil {
				// Pour décoder l'intérieur du CRAs avec la fonction recursive et l'afficher
				if *afficherDataCra == true {
					decodeCRA(cdr.CdrByte, innerTag, cdr)
				}

			} else {
				fmt.Println(innerTag)
			}

			AfficherDetailCRA(*cdr)
			cpt += 1
			// Décalage de l'index apres décodage du CRA
			index += StructureDC.Longueur

			// Sinon fin du script car impossible de determiner la taille du CRA suivant.
		} else {
			fmt.Printf("ERR cas d'article 0x%02x non trouvé\n", firstFourBytes)
			os.Exit(1)
		}

	}
}

// decodeCRA décode les données d'un CRA à partir des données fournies.
//
// Paramètres :
// - data ([]byte) : un tableau d'octets contenant les données du CRA à décoder.
// - structureDC (map[int]defStructs.Champ) : une carte contenant la structure des champs à décoder.
// - cdr (*Cdr) : un pointeur vers l'objet Cdr dans lequel stocker les résultats.
func decodeCRA(data []byte, structureDC map[int]defStructs.Champ, cdr *Cdr) {

	var index = 0

	tagKeys := make([]int, 0, len(structureDC))
	for tagKey := range structureDC {
		tagKeys = append(tagKeys, tagKey)
	}
	sort.Ints(tagKeys)

	for _, tagKey := range tagKeys {
		positionDebut := index
		positionFin := index
		tagValue := (structureDC)[tagKey]
		//spaces := strings.Repeat(" ", tagValue.Decalage)
		if tagValue.Longueur == -1 {
			positionFin = len(data)
			index = len(data)
		} else {
			positionFin = index + tagValue.Longueur
			index += tagValue.Longueur
		}
		val := data[positionDebut:positionFin]
		//valeur := goFonctions.ByteArrayToHexString(val)
		//valeurTraduite := conversionChamp(valeur, tagValue.Codage)
		//fmt.Printf("%v %v%v : %v\n", createIdAudit(cpt), spaces, tagValue.Nom, valeurTraduite)
		cdr.Champs = append(cdr.Champs, Values{Nom: tagValue.Nom, Codage: tagValue.Codage, Decalage: 0, Valeur: val})

	}

}
