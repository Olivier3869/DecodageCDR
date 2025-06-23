package main

import (
	"fmt"
	"local/goFonctions"
)

// afficheEntete affiche l'en-tête des données si l'option d'affichage est activée.
//
// Paramètres :
// - data ([]byte) : un tableau d'octets représentant les données dont l'en-tête doit être affiché.
func afficheEntete(data []byte) {
	if *afficherEntete == true {
		ind := len(data)

		//enteteRabbit := data[index:indexFin]
		if ind == 18 {
			fmt.Printf("Entete %v : %v\n", protocole, goFonctions.ByteArrayToHexString(data[:18]))
		} else if ind == 65 || ind == 97 {
			fin := ind - 1
			entete := goFonctions.ByteArrayToHexString(data[1:fin])
			fmt.Printf("Entete %v : %v\n", protocole, goFonctions.HexToString(entete))
		} else if ind == 64 || ind == 96 {
			fin := ind - 1
			entete := goFonctions.ByteArrayToHexString(data[:fin])
			fmt.Printf("Entete %v : %v\n", protocole, goFonctions.HexToString(entete))
		}
	}
}
