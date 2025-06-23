package goFonctions

import (
	"fmt"
	"strconv"
)

// ConvertDate transforme une date au format "ddMMyy" (en hexadécimal) en "dd/MM/yyyy".
//
// Paramètre :
// - dateHex : une chaîne de caractères représentant la date en hexadécimal.
//
// Retour :
// - une chaîne formatée "dd/MM/yyyy". Retourne une chaîne vide en cas d'erreur.
//
// Fonctionnement :
// - Convertit la chaîne hexadécimale en entier.
// - Extrait le jour, le mois et l'année à partir de l'entier.
// - Formate la date en "dd/MM/yyyy" avec l'année corrigée (ajout de 2000).
func ConvertDate(dateHex string) string {
	// Convertir la chaîne hexadécimale en entier
	dateInt, err := strconv.ParseInt(dateHex, 16, 32)
	if err != nil {
		return ""
	}

	// Extraire le jour, le mois et l'année
	day := int(dateInt & 0xFF)             // Les 2 derniers octets
	month := int((dateInt >> 8) & 0xFF)    // Les 2 suivants
	year := int((dateInt>>16)&0xFF) + 2000 // Les 2 premiers octets + 2000

	// Formater la date
	return fmt.Sprintf("%02d/%02d/%d", day, month, year)
}
