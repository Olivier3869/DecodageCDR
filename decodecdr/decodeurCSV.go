package main

import (
	"bytes"
	"fmt"
	"local/structures/defStructs"
	"os"
)

// decodeCSV décode les données au format CSV et affiche les CRAs.
//
// Paramètres :
// - data ([]byte) : un tableau d'octets représentant les données CSV à décoder.
func decodeCSV(data []byte) {
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

		cdr := new(Cdr)

		index += gestionDebutFichier(data[index:index+fin], index)
		positionFin := 0

		// 4 premiers octets
		firstFourBytes := data[index : index+4] // ajout apres commentaire
		//fmt.Printf("-------- %v %v\n", index, goFonctions.ByteArrayToHexString(firstFourBytes))

		// Recherche de la structure de DC
		StructureDC := rechercheTypeCRA(firstFourBytes, structureCRA) // ajout apres commentaire

		// Si structure de CRA existante
		if StructureDC != nil {

			// Decoupage du CRA avec caractère \n
			for i := index; i < len(data); i++ {
				if data[i] == '\n' { // || i == indexColl
					// Extraire la ligne du début jusqu'au '\n'
					positionFin = i
					break
				}
				// Ajouter la dernière ligne si elle n'est pas suivie par un '\n'
				if positionFin < len(data) {
					positionFin = len(data)
				}
			}
			if positionFin > len(data) {
				positionFin = len(data)
			}
			//fmt.Printf(" %v - %v \n", index, positionFin)
			// Dans les structures, les entêtes sont définis pour ne par être affichés ni comptabilisés
			if StructureDC.Nom != "Entete" {
				// Affichage CRA (quelle que soit l'option le CRA est toujours affiché)
				cdr.CdrByte = data[index:positionFin]

				innerTag := *StructureDC.Tag
				separateur := byte(StructureDC.Longueur)

				if innerTag != nil {
					// Pour décoder l'intérieur du CRAs
					if *afficherDataCra == true {
						decodeCDR(cdr.CdrByte, innerTag, separateur, cdr)
					}

				} else {
					fmt.Println(innerTag)
				}

				AfficherDetailCRA(*cdr)
				cpt += 1
			}
			// Décalage de l'index apres décodage du CRA
			index = positionFin + 1

			// Sinon fin du script car impossible de determiner la taille du CRA suivant.
		} else {
			fmt.Printf("ERR cas d'article 0x%02x non trouvé\n", firstFourBytes)
			os.Exit(1)
		}
	}
}

// decodeCDR décode les données d'un CRA à partir des données fournies.
//
// Paramètres :
// - CRA ([]byte) : un tableau d'octets contenant les données du CRA à décoder.
// - structureDC (map[int]defStructs.Champ) : une carte contenant la structure des champs à décoder.
// - separateur (byte) : le caractère utilisé pour séparer les valeurs dans le CRA.
// - cdr (*Cdr) : un pointeur vers l'objet Cdr dans lequel stocker les résultats.
func decodeCDR(CRA []byte, structureDC map[int]defStructs.Champ, separateur byte, cdr *Cdr) {
	// Diviser la chaîne en utilisant le séparateur

	parts := bytes.Split(CRA, []byte{separateur})
	// Afficher les parties
	for i, part := range parts {
		cdr.Champs = append(cdr.Champs, Values{Nom: structureDC[i].Nom, Codage: structureDC[i].Codage, Decalage: 0, Valeur: part})
	}
}
