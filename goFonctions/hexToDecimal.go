package goFonctions

import "strconv"

// HexToDecimal convertit une chaîne hexadécimale en sa représentation décimale sous forme de chaîne.
//
// Paramètre :
// - hexStr : une chaîne de caractères représentant un nombre en hexadécimal.
//
// Retour :
// - une chaîne contenant la valeur décimale correspondante, ou une chaîne vide en cas d'erreur.
//
// Fonctionnement :
// - Analyse la chaîne hexadécimale en un entier de base 16.
// - Convertit cet entier en une chaîne décimale.
// - Retourne la chaîne décimale.
func HexToDecimal(hexStr string) string {
	// Convertir la chaîne hexadécimale en un entier (base 16)
	num, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		return ""
	}

	// Convertir l'entier en chaîne de caractères décimale
	decimalStr := strconv.FormatInt(num, 10)

	return decimalStr
}
