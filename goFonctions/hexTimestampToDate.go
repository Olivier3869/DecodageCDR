package goFonctions

import (
	"strconv"
	"time"
)

// HexTimestampToDate convertit une chaîne hexadécimale représentant un timestamp Unix en une date formatée "dd/MM/yyyy HH:MM:SS".
//
// Paramètre :
// - hexStr : une chaîne de caractères en hexadécimal représentant un timestamp Unix.
//
// Retour :
// - une chaîne formatée "dd/MM/yyyy HH:MM:SS" ou une chaîne vide si la conversion échoue.
//
// Fonctionnement :
// - Tente de convertir la chaîne hexadécimale en un entier 64 bits.
// - Si la conversion échoue, retourne une chaîne vide.
// - Sinon, crée un objet time.Time à partir du timestamp.
// - Formate la date et l'heure selon le modèle "02/01/2006 15:04:05".
// - Retourne la chaîne formatée.
func HexTimestampToDate(hexStr string) string {
	// Convertir la chaîne hexadécimale en entier
	timestampInt, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		return ""
	}

	// Convertir le timestamp en temps
	t := time.Unix(timestampInt, 0)

	// Formater la date en "dd/mm/yyyy hh:mm:ss"
	return t.Format("02/01/2006 15:04:05")
}
