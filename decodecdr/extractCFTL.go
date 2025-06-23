package main

import (
	"local/goFonctions"
)

// extractCFTL extrait la classe, le format, le tag et la longueur d'un format ASN1 à partir des données fournies.
//
// Paramètres :
// - data ([]byte) : un tableau d'octets contenant les données à analyser.
//
// Retourne :
// - byte : la classe du tag (2 premiers bits).
// - byte : le format du tag (3ᵉ bit).
// - byte : le tag (5 derniers bits).
// - *int : un pointeur vers la longueur décimale, ou nil si non applicable.
// - int : l'index mis à jour après l'extraction.
func extractCFTL(data []byte) (byte, byte, byte, *int, int) {
	index := 0
	octet := data[index]
	classe := (octet >> 6) & 0b11 // Les 2 premiers bits (bit 6 et bit 7)
	form := (octet >> 5) & 1      // Le 3ᵉ bit (bit 3)
	tag := octet & 0b00011111     // Les 5 derniers bits
	index += 1
	//fmt.Printf("class: %b - form: %b - tag: %b / %d\n", classe, form, tag, tag)

	// TAG étendu
	if tag > 30 {
		tag = 0

		// Boucle pour concaténer les 7 bits de poids faible
		for index < len(data) {
			tag = (tag << 7) | data[index]&0x7F // Décalage et ajout des 7 bits

			if data[index]&0x80 == 0 { // Vérifier si le bit de poids fort est à 0
				index++
				break // Sortir de la boucle si c'est le dernier octet
			}
			index++

		}

	}

	// Décodage de la longueur
	var longueurDecimale *int
	octet = data[index]
	if octet&0x80 != 0 {
		// Forme longue de la longueur
		lLongueur := int(octet & 0x7F) // Les 7 derniers bits
		index++

		// Récupération des octets de la longueur
		newLongueur := data[index : index+lLongueur]
		// Conversion de la longueur en entier (big-endian)
		longueurDecimale = goFonctions.BytesToIntBigEndian(newLongueur)
		index += lLongueur
	} else {
		// Forme courte de la longueur
		longueurDecimale = ExtractLast7Bits(octet) // Les 7 derniers bits
		index++

	}

	return classe, form, tag, longueurDecimale, index
}

// ExtractLast7Bits extrait les 7 derniers bits d'un octet et retourne un pointeur vers la valeur.
//
// Paramètres :
// - octet (byte) : l'octet dont les 7 derniers bits doivent être extraits.
//
// Retourne :
// - *int : un pointeur vers la valeur extraite.
func ExtractLast7Bits(octet byte) *int {
	// Extraire les 7 derniers bits
	value := int(octet & 0x7F)

	// Retourner un pointeur vers la valeur
	return &value
}
