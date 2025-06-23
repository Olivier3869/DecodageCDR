package goFonctions

import "fmt"

// FormatDateTimeInv reformate une chaîne de date et heure de format "ssMMhhddMMyyyy" en "dd/MM/yyyy HH:MM:SS".
//
// Paramètre :
// - dateTime : une chaîne de 14 caractères représentant la date et l'heure dans un ordre inversé.
//
// Retour :
// - une chaîne formatée "dd/MM/yyyy HH:MM:SS" ou une chaîne vide si l'entrée est invalide.
//
// Fonctionnement :
// - Vérifie que la chaîne fait bien 14 caractères.
// - Extrait les composants année, mois, jour, heure, minute, seconde en inversant l'ordre.
// - Vérifie que chaque composant est numérique.
// - Construit la chaîne formatée dans le format souhaité.
func FormatDateTimeInv(dateTime string) string {
	if len(dateTime) != 14 {
		return ""
	}

	// Extraire les différentes parties de la date
	year := dateTime[12:14] + dateTime[10:12]
	month := dateTime[8:10]
	day := dateTime[6:8]
	hour := dateTime[4:6]
	minute := dateTime[2:4]
	second := dateTime[0:2]

	// Vérifier si les valeurs sont numériques
	if !IsNumeric(year) || !IsNumeric(month) || !IsNumeric(day) ||
		!IsNumeric(hour) || !IsNumeric(minute) || !IsNumeric(second) {
		return ""
	}

	// Construire la date au format souhaité
	formattedDateTime := fmt.Sprintf("%s/%s/%s %s:%s:%s", day, month, year, hour, minute, second)
	return formattedDateTime
}
