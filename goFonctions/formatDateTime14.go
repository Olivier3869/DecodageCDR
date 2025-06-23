package goFonctions

import "time"

// FormatDateTime14 reformate une chaîne de date et heure de format "yyyyMMddHHmmss" en "dd/MM/yyyy HH:MM:SS".
//
// Paramètre :
// - dateTime : une chaîne de 14 caractères représentant la date et l'heure dans le format "yyyyMMddHHmmss".
//
// Retour :
// - une chaîne formatée "dd/MM/yyyy HH:MM:SS" ou une chaîne vide si l'entrée est invalide.
//
// Fonctionnement :
// - Définit le layout correspondant au format attendu.
// - Tente de parser la chaîne d'entrée selon ce layout.
// - Si le parsing échoue, retourne une chaîne vide.
// - Sinon, formate la date en "02/01/2006 15:04:05".
// - Retourne la chaîne formatée.
func FormatDateTime14(dateTime string) string {
	layout := "20060102150405"
	t, err := time.Parse(layout, dateTime)
	if err != nil {
		return ""
	}
	return t.Format("02/01/2006 15:04:05")
}
