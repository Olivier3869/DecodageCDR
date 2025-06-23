package main

import (
	"encoding/binary"
	"fmt"
	"local/goFonctions"
)

// decodeColl décode les données d'un en-tête COLL et retourne la taille de l'en-tête et la taille des données.
//
// Paramètres :
// - data ([]byte) : un tableau d'octets représentant les données à décoder.
//
// Retourne :
// - int : la taille de l'en-tête décodé.
// - int : la taille des données décodées.
func decodeColl(data []byte) (int, int) {
	// Interpréter la version
	version := binary.BigEndian.Uint32(data[0:4])

	// Déclarer toutes les variables à l'extérieur du bloc conditionnel
	var tailleEntete uint16
	var numEEC uint16
	var horodate string
	var tailleDonnees uint32
	var fichier string

	var idAuditVersion, idAuditNumDuplicata, idAuditNumTransit, idAuditDecade string
	var idAuditDate uint32
	// var idAuditSeqSource uint16
	var idAuditIdfSource uint16
	var idAuditModule string

	if version == 2 {
		tailleEntete = binary.BigEndian.Uint16(data[4:6])
		numEEC = binary.BigEndian.Uint16(data[12:14])
		//numSequence := binary.BigEndian.Uint16(data[14:18])
		dateHeure := goFonctions.ByteArrayToHexString(data[18:22])
		horodate = goFonctions.HexToDecimal(dateHeure)
		tailleDonnees = binary.BigEndian.Uint32(data[25:29])

		idAuditVersion = goFonctions.HexToDecimal(goFonctions.ByteArrayToHexString(data[29:30]))
		idAuditNumDuplicata = goFonctions.HexToDecimal(goFonctions.ByteArrayToHexString(data[30:31]))
		idAuditNumTransit = goFonctions.HexToDecimal(goFonctions.ByteArrayToHexString(data[31:32]))
		idAuditDecade = goFonctions.HexToDecimal(goFonctions.ByteArrayToHexString(data[32:33]))
		idAuditDate = binary.LittleEndian.Uint32(data[33:37])
		//idAuditSeqSource = binary.LittleEndian.Uint16(data[37:41])
		idAuditIdfSource = binary.LittleEndian.Uint16(data[41:43])
		idAuditModule = goFonctions.HexToDecimal(goFonctions.ByteArrayToHexString(data[43:45]))

		nomFichier := goFonctions.ByteArrayToHexString(data[46:tailleEntete])
		fichier = goFonctions.HexToString(nomFichier)

	} else {
		tailleEntete = 38
		numEEC = binary.BigEndian.Uint16(data[4:6])
		dateHeure := goFonctions.ByteArrayToHexString(data[10:14])
		horodate = goFonctions.HexToDecimal(dateHeure)
		tailleDonnees = binary.BigEndian.Uint32(data[17:21])

		idAuditVersion = goFonctions.HexToDecimal(goFonctions.ByteArrayToHexString(data[21:22]))
		idAuditNumDuplicata = goFonctions.HexToDecimal(goFonctions.ByteArrayToHexString(data[22:23]))
		idAuditNumTransit = goFonctions.HexToDecimal(goFonctions.ByteArrayToHexString(data[23:24]))
		idAuditDecade = goFonctions.HexToDecimal(goFonctions.ByteArrayToHexString(data[24:25]))
		idAuditDate = binary.LittleEndian.Uint32(data[25:29])
		//idAuditSeqSource = binary.LittleEndian.Uint16(data[29:33])
		idAuditIdfSource = binary.LittleEndian.Uint16(data[33:35])
		idAuditModule = goFonctions.HexToDecimal(goFonctions.ByteArrayToHexString(data[35:37]))

	}
	audit := fmt.Sprintf("%v.%v.%v_%v_0_%v_%v_%v", idAuditVersion, idAuditDecade, idAuditDate, idAuditIdfSource, idAuditNumDuplicata, idAuditNumTransit, idAuditModule)
	//audit := fmt.Sprintf("%v.%v.%v_%v_%v_%v_%v_%v", idAuditVersion, idAuditDecade, idAuditDate, idAuditIdfSource, idAuditSeqSource, idAuditNumDuplicata, idAuditNumTransit, idAuditModule)
	idAudit = fmt.Sprintf("%v.%v.%v_%v_", idAuditVersion, idAuditDecade, idAuditDate, idAuditIdfSource)

	if *afficherEntete == true {
		fmt.Println("-------------------bloc: #0001-------------------")
		fmt.Println("version     : ", version)
		fmt.Println("eec         : ", numEEC)
		fmt.Println("horodate    : ", horodate)
		fmt.Println("id audit    : ", audit)
		fmt.Println("taille      : ", tailleDonnees)
		if version == 2 {
			fmt.Println("fichier     : ", fichier)
		}
		fmt.Println("-------------------------------------------------")
	}
	return int(tailleEntete), int(tailleDonnees)
}
