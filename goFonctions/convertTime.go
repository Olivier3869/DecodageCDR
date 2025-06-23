package goFonctions

import (
	"fmt"
	"strconv"
)

// ConvertTime transforme une heure au format "hhmmss" (en hexadécimal) en "hh:mm:ss".
//
// Paramètre :
// - timeHex : une chaîne de caractères représentant l'heure en hexadécimal.
//
// Retour :
// - une chaîne formatée "hh:mm:ss". Retourne une chaîne vide en cas d'erreur.
//
// Fonctionnement :
// - Convertit la chaîne hexadécimale en entier.
// - Extrait les heures, minutes et secondes à partir de l'entier.
// - Formate la sortie en "hh:mm:ss".
func ConvertTime(timeHex string) string {
	// Convertir la chaîne hexadécimale en entier
	timeInt, err := strconv.ParseInt(timeHex, 16, 32)
	if err != nil {
		return ""
	}

	// Extraire les heures, minutes et secondes
	hours := int((timeInt >> 16) & 0xFF)  // Les 2 premiers octets
	minutes := int((timeInt >> 8) & 0xFF) // Les 2 suivants
	seconds := int(timeInt & 0xFF)        // Les 2 derniers

	// Formater l'heure
	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
