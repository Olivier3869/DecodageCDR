package main

import (
	"fmt"
	"local/goFonctions"
	"os"
	"sort"
)

// var compteur string
var cpt int
var indexColl int

// decodeTASNP décode les données TASNP et affiche les CRAs.
//
// Paramètres :
// - data ([]byte) : un tableau d'octets représentant les données à décoder.
func decodeTASNP(data []byte) {
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

		//fmt.Printf("%v - %v\n", index, indexColl)

		// Décodage du HeaderBloc (on ne récupère que la taille)
		recordLengthBytes := data[index : index+1]
		recordLength := goFonctions.BytesToIntBigEndian(recordLengthBytes)
		headerBlock := data[index : index+*recordLength]

		// Affichage du headerBloc
		if *afficherEntete == true {
			fmt.Printf("--------------------HeaderBloc #%v------------------\n", cpt)
			fmt.Printf("%v%v\n", createIdAudit(0), goFonctions.ByteArrayToHexString(headerBlock))
			decodeHeaderBloc(headerBlock)
			fmt.Println("-------------------------------------------------------")
		}

		dataLengthInBlockBytes := data[index+6 : index+8]
		dataLengthInBlock := goFonctions.BytesToIntLittleEndian(dataLengthInBlockBytes)
		//fmt.Printf("dataLengthInBlock %v\n", dataLengthInBlock)
		var blocCDR []byte

		if dataLengthInBlock < len(data) {
			blocCDR = data[index+*recordLength : index+dataLengthInBlock]
			// Pour décoder l'intérieur du CRAs avec la fonction recursive et l'afficher
			decodeBlocCDR(blocCDR)

		} else {
			blocCDR = data[index+*recordLength:]
			if *afficherEntete == true {
				fmt.Printf("Trailer : %v%v\n", createIdAudit(0), goFonctions.ByteArrayToHexString(blocCDR))
			}
		}

		index += 65408 // Longueur réelle des blocs completés par des bit de bourrage ff
		//cpt += 1
	}
}

// decodeBlocCDR décode les blocs CDR à partir des données fournies.
//
// Paramètres :
// - data ([]byte) : un tableau d'octets contenant les données CDR à décoder.
func decodeBlocCDR(data []byte) {
	index := 0

	for index < len(data) {
		//fmt.Printf("-------------------------------%v : %v ------------------------------------------------------\n", index, len(data))

		cdr := new(Cdr)
		tailleCDRBytes := data[index : index+2]

		tailleCDR := goFonctions.BytesToIntLittleEndian(tailleCDRBytes)
		//fmt.Printf("tailleCDR %v\n", tailleCDR)
		if tailleCDR > 24 {
			// CRA complet
			cdr.CdrByte = data[index : index+tailleCDR]
			//fmt.Printf("%v%v\n", createIdAudit(cpt), goFonctions.ByteArrayToHexString(cra))
		}

		if tailleCDR == 24 {
			// Trailer
			trailerCDR := data[index : index+tailleCDR]
			if *afficherEntete == true {
				fmt.Printf("TrailerCDR : %v\n", goFonctions.ByteArrayToHexString(trailerCDR))
			}
			if index+tailleCDR == len(data) {
				break
			}
		} else if tailleCDR <= 2 {
			fmt.Println("Problème de décodage")
			os.Exit(0)
		} else if *afficherDataCra == true {

			// Pour décoder l'intérieur du CRAs avec la fonction recursive

			// CDR Header
			indexHeader, recordType := decodeHeaderCdr(data[index:index+tailleCDR], cdr)
			index += indexHeader

			// CDR Static
			indexStatic := decodeStaticCdr(data[index:index+tailleCDR], *recordType, cdr)
			index += indexStatic

			// CDR ASN1
			tailleCDR = tailleCDR - indexHeader - indexStatic
			asn1CDR := data[index+6 : index+tailleCDR]
			//fmt.Println(goFonctions.ByteArrayToHexString(asn1CDR))
			firstFourBytes := asn1CDR[0:5]

			// Selection de la structure
			StructureDC := rechercheTypeCRA(firstFourBytes, structureCRA)
			Tag := *StructureDC.Tag

			// Décodage du CDR ASN1

			decodeInnerASN1(asn1CDR, Tag, cdr, 0)

		}

		AfficherDetailCRA(*cdr)
		cpt += 1
		index += tailleCDR

	}
}

// decodeHeaderBloc décode le bloc d'en-tête (implémentation à compléter).
//
// Paramètres :
// - data ([]byte) : un tableau d'octets représentant le bloc d'en-tête à décoder.
func decodeHeaderBloc(data []byte) {
	// Implémentation à compléter
}

// decodeHeaderCdr décode l'en-tête du CDR et retourne l'index mis à jour et le type d'enregistrement.
//
// Paramètres :
// - data ([]byte) : un tableau d'octets représentant le CDR à décoder.
// - cdr (*Cdr) : un pointeur vers l'objet Cdr dans lequel stocker les résultats.
//
// Retourne :
// - int : l'index mis à jour après le décodage.
// - *int : un pointeur vers le type d'enregistrement.
func decodeHeaderCdr(data []byte, cdr *Cdr) (index int, recordType *int) {
	index = 0

	structureHeader := structureCRA[2]
	champsHeader := *structureHeader.Tag

	tagKeys := make([]int, 0, len(champsHeader))
	for tagKey := range champsHeader {
		tagKeys = append(tagKeys, tagKey)
	}
	sort.Ints(tagKeys)

	for _, tagKey := range tagKeys {
		positionDebut := index
		positionFin := index
		tagValue := (champsHeader)[tagKey]

		if tagValue.Longueur == -1 {
			positionFin = len(data)
			index = len(data)
		} else {
			positionFin = index + tagValue.Longueur
			index += tagValue.Longueur
		}
		val := data[positionDebut:positionFin]

		if tagValue.Nom == "record_type" {
			recordType = goFonctions.BytesToIntBigEndian(val)
		}

		cdr.Champs = append(cdr.Champs, Values{Tag: 9999, Nom: tagValue.Nom, Codage: tagValue.Codage, Decalage: 0, Valeur: val})
	}
	return index, recordType
}

// decodeStaticCdr décode les données statiques du CDR et retourne l'index mis à jour.
//
// Paramètres :
// - data ([]byte) : un tableau d'octets représentant le CDR à décoder.
// - recordType int : le type d'enregistrement du CDR.
// - cdr (*Cdr) : un pointeur vers l'objet Cdr dans lequel stocker les résultats.
//
// Retourne :
// - int : l'index mis à jour après le décodage.
func decodeStaticCdr(data []byte, recordType int, cdr *Cdr) (index int) {
	index = 0

	structureStatic := structureCRA[3]
	typeStatic := *structureStatic.Tag
	typeCdr := typeStatic[recordType]

	cdr.Champs = append(cdr.Champs, Values{Tag: recordType, Nom: typeCdr.Nom, Codage: 0, Decalage: 0})

	champsStatic := *typeCdr.Tag

	tagKeys := make([]int, 0, len(champsStatic))
	for tagKey := range champsStatic {
		tagKeys = append(tagKeys, tagKey)
	}
	sort.Ints(tagKeys)

	for _, tagKey := range tagKeys {
		positionDebut := index
		positionFin := index
		tagValue := (champsStatic)[tagKey]

		if tagValue.Longueur == -1 {
			positionFin = len(data)
			index = len(data)
		} else {
			positionFin = index + tagValue.Longueur
			index += tagValue.Longueur
		}
		val := data[positionDebut:positionFin]

		// Valeur du Tag arbitrairement placée à 9999 afin de définir un type d'affichage spécifique
		cdr.Champs = append(cdr.Champs, Values{Tag: 9999, Nom: tagValue.Nom, Codage: tagValue.Codage, Decalage: 1, Valeur: val})
	}
	return index
}
