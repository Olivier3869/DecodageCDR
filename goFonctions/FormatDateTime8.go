package goFonctions

import "fmt"

// FormatDateTime8 formate une chaîne de date et heure de format "yyMMddHHmmss" en "dd/MM/20yy HH:MM:SS".
//
// Paramètre :
// - dateTime : une chaîne de 12 caractères représentant la date et l'heure.
//
// Retour :
// - une chaîne formatée "dd/MM/20yy HH:MM:SS" ou une chaîne vide si l'entrée est invalide.
//
// Fonctionnement :
// - Vérifie que la chaîne fait bien 12 caractères.
// - Extrait les composants année (en deux chiffres), mois, jour, heure, minute, seconde.
// - Vérifie que chaque composant est numérique.
// - Construit la chaîne formatée en ajoutant "20" devant l'année pour obtenir une année à quatre chiffres.
func FormatDateTime8(dateTime string) string {
	if len(dateTime) != 12 {
		return ""
	}

	// Extraire les différentes parties de la date
	year := dateTime[0:2]
	month := dateTime[2:4]
	day := dateTime[4:6]
	hour := dateTime[6:8]
	minute := dateTime[8:10]
	second := dateTime[10:12]

	// Vérifier si les valeurs sont numériques
	if !IsNumeric(year) || !IsNumeric(month) || !IsNumeric(day) ||
		!IsNumeric(hour) || !IsNumeric(minute) || !IsNumeric(second) {
		return ""
	}

	// Construire la date au format souhaité
	formattedDateTime := fmt.Sprintf("%s/%s/20%s %s:%s:%s", day, month, year, hour, minute, second)
	return formattedDateTime
}
