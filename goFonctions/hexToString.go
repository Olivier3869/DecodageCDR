package goFonctions

import (
	"fmt"
	"strconv"
	"strings"
)

// HexToString convertit une chaîne hexadécimale en une chaîne de caractères imprimables.
//
// Paramètre :
// - hexString : une chaîne hexadécimale où chaque paire de caractères représente un octet.
//
// Retour :
// - la chaîne de caractères correspondante, ou une chaîne vide en cas d'erreur.
//
// Fonctionnement :
// - Parcourt la chaîne hexadécimale par paires de caractères.
// - Convertit chaque paire en un octet.
// - Construit la chaîne de caractères à partir des octets, en filtrant les caractères non imprimables.
func HexToString(hexString string) string {
	var output strings.Builder // Initialise un builder pour construire la chaîne de sortie

	// Parcours de la chaîne hexadécimale par pas de 2 caractères
	for i := 0; i < len(hexString); i += 2 {
		// Extraction de chaque paire hexadécimale
		hex := hexString[i : i+2]

		// Conversion de la paire hex en un entier (base 16)
		decimal, err := strconv.ParseInt(hex, 16, 64)
		if err != nil {
			// En cas d'erreur de conversion, affiche un message et quitte la fonction
			fmt.Printf("Erreur lors de la conversion hexadécimale: %v\n", err)
			return ""
		}

		// Conversion de l'entier en byte
		b := byte(decimal)

		// Vérification si le byte est un caractère ASCII imprimable (espaces, lettres, chiffres, ponctuation)
		if b >= 32 && b <= 126 {
			// Ajout du caractère à la chaîne de sortie si il est imprimable
			output.WriteByte(b)
		}
	}

	// Retourne la chaîne construite
	return output.String()
}
