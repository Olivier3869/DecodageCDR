package main

import (
	"local/goFonctions"
	"strings"
)

// gestionDebutFichier traite les premiers octets d'un fichier pour déterminer s'il s'agit d'un en-tête ou d'un fichier COLL.
//
// Paramètres :
// - data ([]byte) : un tableau d'octets représentant le contenu du fichier à analyser.
// - index (int) : l'index à partir duquel commencer le traitement.
//
// Retourne :
// - int : l'index mis à jour après le traitement des en-têtes et des données.
func gestionDebutFichier(data []byte, index int) int {
	innerIndex := 0
	// Pour décodage des entêtes de fichiers coll_
	if decoderColl == true && index >= indexColl {
		//fmt.Printf("%v\n", goFonctions.ByteArrayToHexString(data[0:]))
		// En version 2 ; la taille de l'entête est de 301 octets maximum
		tailleEntete, tailleDonnees := decodeColl(data)
		//fmt.Printf("entete : %v - Donnees : %v - audit : %v\n", tailleEntete, tailleDonnees, idAudit)
		//fmt.Printf("%v\n", goFonctions.ByteArrayToHexString(data[index:index+tailleEntete]))
		innerIndex += tailleEntete

		indexColl = index + tailleDonnees + tailleEntete
		cpt = 1
		//os.Exit(0)
	}

	// 6 octets suivants pour verifier si c'est un entete MQSeries RabbitMQ Etc
	firstSixBytes := data[0:6]
	//fmt.Printf("Entete first bytes : %v\n", goFonctions.ByteArrayToHexString(data[0:100]))

	if ok, ind := goFonctions.EstEntete(firstSixBytes); ok {
		d := 0
		//fmt.Println(goFonctions.ByteArrayToHexString(data[0:ind]))
		// ATTENTION : les entêtes en entrée sont plus longs que ceux en sortie...
		if strings.HasPrefix(*typeTraitement, "T") {
			if ind == 65 {
				d = 1
			}
			ind += 32
			idAudit = goFonctions.ByteArrayToHexString(data[d:ind])
			cpt = 1
		} else if strings.HasPrefix(*typeTraitement, "D") {
			protocole = "Kafka"
		} else {
			idAudit = goFonctions.ByteArrayToHexString(data[31:93])
		}

		//fmt.Printf("Longueur Entete : %v\n", ind)
		indexFin := ind // + index
		if indexFin > len(data) {
			indexFin = len(data)
		}
		afficheEntete(data[0:indexFin])
		innerIndex += ind
	}

	// cas des Fichier TTmmm 3fc2 en debut de fichier
	if *typeTraitement == "TTmmm" && goFonctions.ByteArrayToHexString(data[innerIndex:innerIndex+1]) == "3f" {
		innerIndex += 2
	}

	return innerIndex
}
