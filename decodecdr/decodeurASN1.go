package main

import (
	"fmt"
	"local/structures/defStructs"
)

var innerTag map[int]defStructs.Champ

// decodeASN1 décode les données au format ASN.1 et affiche les CRAs.
//
// Paramètres :
// - data ([]byte) : un tableau d'octets représentant les données à décoder.
func decodeASN1(data []byte) {
	index := 0
	cpt = 1
	indexColl = 0

	for index < len(data) {
		fin := 301
		// Si la longueur restante est inférieure à 301, ajuster fin
		if reste := len(data[index:]); reste < fin {
			fin = reste
		}

		index += gestionDebutFichier(data[index:index+fin], index)

		// Décodage des CRAs
		//fmt.Printf("---------------------------------------------------%v :(%v): %v ------------------------------------------------------\n", index, indexColl, len(data))
		//fmt.Printf("%v %v\n", compteur, goFonctions.ByteArrayToHexString(data[index:index+32]))

		// init du cdr et debut du cdr pour forme indéfinie de la longueur
		indexDebut := index
		cdr := new(Cdr)

		// Décodage class form tag longueur
		//fmt.Printf("8 bytes : %v\n", goFonctions.ByteArrayToHexString(data[index:index+8]))
		_, _, _, longueurDecimale, decIndex := extractCFTL(data[index:])
		//if longueurDecimale != nil {
		//	fmt.Printf("class: %b - form: %b - tag: %b / %d - longueur: %d\n", classe, form, tag, tag, *longueurDecimale)
		//} else {
		//	fmt.Printf("class: %b - form: %b - tag: %b / %d - newIndex: %d\n", classe, form, tag, tag, newIndex)
		//}

		// Definition de la structure et du Tag
		firstFourBytes := data[index : index+4]
		StructureDC := rechercheTypeCRA(firstFourBytes, structureCRA)

		innerTag = FindInnerTag(*StructureDC)

		if longueurDecimale != nil {
			// CRA complet
			cdr.CdrByte = data[index : index+*longueurDecimale+decIndex]

			if innerTag != nil {
				// Pour décoder l'intérieur du CRAs avec la fonction recursive et l'afficher
				if *afficherDataCra == true {
					// on connait la longueur du CRA donc pas besoin de l'index
					decodeInnerASN1(cdr.CdrByte, innerTag, cdr, 0)
				}

			} else {
				fmt.Println("ERR : innerTag nil")
			}
			index += *longueurDecimale + decIndex
		} else {
			// Forme indéfinie de la longueur donc besoin de l'index
			//fmt.Println("ERR : longueurDecimale nil")
			newIndex := decodeInnerASN1(data[index:], innerTag, cdr, 0)
			cdr.CdrByte = data[indexDebut : indexDebut+newIndex]

			index += newIndex
		}

		AfficherDetailCRA(*cdr)

		cpt += 1
		//os.Exit(0)
	}
}

// decodeInnerASN1 décode récursivement les données d'un CRA au format ASN.1.
//
// Paramètres :
// - data ([]byte) : un tableau d'octets contenant les données à décoder.
// - structureDC (map[int]defStructs.Champ) : une carte contenant la structure des champs à décoder.
// - cdr (*Cdr) : un pointeur vers l'objet Cdr dans lequel stocker les résultats.
// - decalage (int) : le décalage à appliquer pour l'affichage des niveaux de profondeur.
//
// Retourne :
// - int : l'index mis à jour après le décodage.
func decodeInnerASN1(data []byte, structureDC map[int]defStructs.Champ, cdr *Cdr, decalage int) int {
	index := 0

	for index < len(data) {

		// Cas de la forme indéfinie de la longueur: la fin du cra est 0000
		octet1 := data[index]
		octet2 := data[index+1]
		// Si fin du CRA ; break de la boucle pour passer au suivant
		if octet1 == 0x00 && octet2 == 0x00 {
			index += 2
			break
		}

		// Décodage class form tag longueur
		// 4 premiers octets
		//firstTwoBytes := data[index : index+4]
		//fmt.Printf("-------- %v %v --------\n", compteur, goFonctions.ByteArrayToHexString(firstTwoBytes))

		_, form, tag, longueurDecimale, newIndex := extractCFTL(data[index:])
		//if longueurDecimale != nil {
		//	fmt.Printf("class: %b - form: %b - tag: %b / %d - longueur: %d\n", classe, form, tag, tag, *longueurDecimale)
		//} else {
		//	fmt.Printf("class: %b - form: %b - tag: %b / %d - newIndex: %d\n", classe, form, tag, tag, newIndex)
		//}

		champsDC := structureDC[int(tag)]

		if longueurDecimale != nil {

			index += newIndex

			value := data[index : index+*longueurDecimale]

			if form == 1 && champsDC.Tag != nil {
				cdr.Champs = append(cdr.Champs, Values{Tag: int(tag), Nom: champsDC.Nom, Codage: champsDC.Codage, Decalage: decalage})
				decodeInnerASN1(value, *champsDC.Tag, cdr, decalage+1)
				index += *longueurDecimale
			} else {
				cdr.Champs = append(cdr.Champs, Values{Tag: int(tag), Nom: champsDC.Nom, Codage: champsDC.Codage, Decalage: decalage, Valeur: value})
				index += *longueurDecimale
			}
		} else {
			cdr.Champs = append(cdr.Champs, Values{Tag: int(tag), Nom: champsDC.Nom, Codage: champsDC.Codage, Decalage: decalage})
			structureDC = *champsDC.Tag
			index += 2
		}
	}

	return index
}
