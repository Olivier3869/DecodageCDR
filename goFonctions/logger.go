package goFonctions

import (
	"fmt"
	"strings"
	"time"
)

var timerReference string

// Couleurs ANSI
const (
	Reset         = "\033[0m"
	Red           = "\033[1;31m"
	Yellow        = "\033[33m"
	DarkTurquoise = "\033[38;5;30m" // Couleur approximative pour bleu turquoise sombre
	BoldGreen     = "\033[1;32m"
	Turquoise     = "\033[36m" // Turquoise
)

// formatTime retourne l'heure actuelle en format HH:MM:SS.
// Si l'heure est identique à la dernière appelée, retourne une chaîne vide.
// Sinon, met à jour la référence et retourne l'heure entre crochets.
func formatTime() string {
	loc, _ := time.LoadLocation("Europe/Paris")
	now := time.Now().In(loc)
	temps := now.Format("15:04:05")

	if temps == timerReference {
		return "          "
	} else {
		timerReference = temps
		return fmt.Sprintf("[%s]", temps)
	}

}

// Log affiche un message avec l'heure en turquoise sombre, sans couleur pour le message.
//
// Exemple : [HH:MM:SS] message
func Log(message string) {
	fmt.Printf("%s%s%s %v", DarkTurquoise, formatTime(), Reset, message)
}

// LogAlerte affiche un message avec l'heure en turquoise sombre et le texte en rouge.
//
// Exemple : [HH:MM:SS] message
func LogAlerte(message string) {
	fmt.Printf("%s%s %s%v%s", DarkTurquoise, formatTime(), Red, message, Reset)
}

// LogInfo affiche un message avec l'heure en turquoise sombre et le texte en jaune.
//
// Exemple : [HH:MM:SS] message
func LogInfo(message string) {
	fmt.Printf("%s%s %s%v%s", DarkTurquoise, formatTime(), Yellow, message, Reset)
}

// LogSuccess affiche un message avec l'heure en turquoise sombre et le texte en vert gras.
//
// Exemple : [HH:MM:SS] message
func LogSuccess(message string) {
	fmt.Printf("%s%s %s%v%s", DarkTurquoise, formatTime(), BoldGreen, message, Reset)
}

// LogRule affiche un message encadré par des traits horizontaux de chaque côté, en turquoise.
//
// Exemple :
// ─────────────────────────────────── message ─────────────────────────────────────────
func LogRule(message string) {
	totalWidth := 150
	messageWidth := len(message)
	// Calcul du nombre de traits de chaque côté
	sideWidth := (totalWidth - messageWidth - 2) / 2 // -2 pour les espaces autour du message

	// Crée les traits de chaque côté
	side := strings.Repeat("─", sideWidth)

	// Affiche le message encadré
	fmt.Printf("%s%s %s %s%s\n", Turquoise, side, message, side, Reset)
}
