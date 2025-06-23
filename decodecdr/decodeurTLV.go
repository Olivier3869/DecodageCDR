package main

import (
	"fmt"
	"local/structures/defStructs"
	"sort"
)

// decodeTLV décode les données TLV et affiche les CRAs (T1).
//
// Paramètres :
// - data ([]byte) : un tableau d'octets représentant les données à décoder.
func decodeTLV(data []byte) {
	index := 0
	cpt = 1
	indexColl = 0

	for index < len(data) {
		//fmt.Printf("-------------------------------%v :(%v): %v ------------------------------------------------------\n", index, indexColl, len(data))
		fin := 301
		// Si la longueur restante est inférieure à 301, ajuster fin
		if reste := len(data[index:]); reste < fin {
			fin = reste
		}

		index += gestionDebutFichier(data[index:index+fin], index)

		// Si on finit par un en pied
		if index >= len(data) {
			break
		}

		// 4 premiers octets
		firstFourBytes := data[index : index+4]
		//fmt.Printf("4BYTES %v\n", goFonctions.ByteArrayToHexString(firstFourBytes))
		tlvInfo, newIndex, nIndex := extractTLV(data, index)
		tlvType := tlvInfo[0].(int)
		tlvLength := tlvInfo[1].(int)
		//tlvValue := tlvInfo[2].([]byte)

		// Affichage CRA (quelle que soit l'option le CRA est toujours affiché)
		// - 2 pour TTmmm et TTnodes
		// TODO prévoir calcul dans extractTLV
		//nIndex := 4
		//if *typeTraitement == "TTmmm" || *typeTraitement == "TTnodes" || *typeTraitement == "trait1" {
		//	nIndex = 2
		//}
		cdr := new(Cdr)
		// CRA complet
		cdr.CdrByte = data[index : index+tlvLength+nIndex]

		StructureDC := rechercheTypeCRA(firstFourBytes, structureCRA)

		if StructureDC != nil {

			innerTag := *StructureDC.Tag
			if innerTag != nil {
				// Pour décoder l'intérieur du CRAs avec la fonction recursive et l'afficher
				if *afficherDataCra == true {
					decodeInnerTLV(cdr.CdrByte, innerTag, cdr, 0)
				}

				cpt += 1
			} else {
				fmt.Println(innerTag)
			}
		} else {
			fmt.Printf("ERR cas d'article 0x%02x non trouvé\n", tlvType)
		}

		AfficherDetailCRA(*cdr)
		// Décalage de l'index apres décodage du CRA
		index = newIndex

	}
}

// decodeInnerTLV décode récursivement les CRAs à partir des données fournies.
//
// Paramètres :
// - data ([]byte) : un tableau d'octets contenant les données à décoder.
// - structureDC (map[int]defStructs.Champ) : une carte contenant la structure des champs à décoder.
// - cdr (*Cdr) : un pointeur vers l'objet Cdr dans lequel stocker les résultats.
// - decalage (int) : le décalage à appliquer pour l'affichage des niveaux de profondeur.
func decodeInnerTLV(data []byte, structureDC map[int]defStructs.Champ, cdr *Cdr, decalage int) {
	index := 0
	//fmt.Println("data : ", goFonctions.ByteArrayToHexString(data))
	//fmt.Println(structureDC)

	//Décodage du CRA
	for index < len(data) {

		// Si la clé 0 est présente structure de type fixe
		if _, ok := structureDC[0]; ok && structureDC[0].Tag == nil && structureDC[0].Longueur != 0 {
			// Extraire les clés de la map et les trier
			tagKeys := make([]int, 0, len(structureDC))
			for tagKey := range structureDC {
				tagKeys = append(tagKeys, tagKey)
			}
			sort.Ints(tagKeys)

			for _, tagKey := range tagKeys {

				positionDebut := index
				positionFin := index

				// Cas ou la data plus petite que le tableau
				if positionDebut >= len(data) {
					break
				}

				tagValue := (structureDC)[tagKey]
				if tagValue.Longueur == -1 {
					positionFin = len(data)
					index = len(data)
				} else {
					positionFin = index + tagValue.Longueur
					index += tagValue.Longueur
				}

				if positionDebut < positionFin {
					val := data[positionDebut:positionFin]
					// Valeur du Tag arbitrairement placée à 9999 afin de définir un type d'affichage spécifique
					cdr.Champs = append(cdr.Champs, Values{Tag: 9999, Nom: tagValue.Nom, Codage: tagValue.Codage, Decalage: decalage + 1, Valeur: val})
				}
			}
		} else {
			// Sinon décodage tlv
			// Récursivité, car tlvValue est de type TLV
			tlvInfo, newIndex, _ := extractTLV(data, index)
			tlvType := tlvInfo[0].(int)
			//tlvLength := tlvInfo[1].(int)
			tlvValue := tlvInfo[2].([]byte)
			champsDC := structureDC[tlvType]

			if _, exists := structureDC[tlvType]; !exists {
				// Valeur du Tag arbitrairement placée à 9999 afin de définir un type d'affichage spécifique
				cdr.Champs = append(cdr.Champs, Values{Tag: 9999, Nom: fmt.Sprintf("UNKNOWN(%02x)", tlvType), Codage: 0, Decalage: decalage + 1, Valeur: tlvValue})
			} else if champsDC.Tag == nil {
				// Valeur du Tag arbitrairement placée à 9999 afin de définir un type d'affichage spécifique
				cdr.Champs = append(cdr.Champs, Values{Tag: 9999, Nom: champsDC.Nom, Codage: champsDC.Codage, Decalage: decalage + 1, Valeur: tlvValue})
			} else {
				cdr.Champs = append(cdr.Champs, Values{Tag: tlvType, Nom: champsDC.Nom, Codage: 0, Decalage: decalage})
				decodeInnerTLV(tlvValue, *champsDC.Tag, cdr, decalage+1)
			}
			index = newIndex
		}
	}

}
