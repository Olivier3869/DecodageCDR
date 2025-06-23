package goFonctions

import (
	"fmt"
)

// FormatDateTimeGMT formate une chaîne de date et heure de format "yyMMddHHmmss+xxxx" en "dd/MM/20yy HH:MM:SS +xxxx".
//
// Paramètre :
// - input : une chaîne de 18 caractères représentant la date, l'heure et le décalage GMT.
//
// Retour :
// - une chaîne formatée "dd/MM/20yy HH:MM:SS +xxxx" ou la chaîne d'entrée si sa longueur n'est pas 18 ou si une valeur n'est pas numérique.
//
// Fonctionnement :
// - Vérifie que la chaîne fait bien 18 caractères.
// - Extrait les composants année (en deux chiffres), mois, jour, heure, minute, seconde, et décalage GMT.
// - Vérifie que chaque composant est numérique.
// - Construit la chaîne formatée en ajoutant "20" devant l'année pour obtenir une année à quatre chiffres, et inclut le décalage GMT.
func FormatDateTimeGMT(input string) string {
	if len(input) != 18 {
		return input
	}
	// Extraire les différentes parties de la date
	year := input[0:2]
	month := input[2:4]
	day := input[4:6]
	hour := input[6:8]
	minute := input[8:10]
	second := input[10:12]
	offset := input[14:18]

	// Vérifier si les valeurs sont numériques
	if !IsNumeric(year) || !IsNumeric(month) || !IsNumeric(day) ||
		!IsNumeric(hour) || !IsNumeric(minute) || !IsNumeric(second) ||
		!IsNumeric(offset) {
		return ""
	}

	// Construire la date au format souhaité
	formattedDate := fmt.Sprintf("%s/%s/20%s %s:%s:%s +%s", day, month, year, hour, minute, second, offset)
	return formattedDate
}
