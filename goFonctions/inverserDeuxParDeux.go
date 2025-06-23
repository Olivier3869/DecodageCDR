package goFonctions

// InverserDeuxParDeux échange chaque paire de caractères consécutifs dans une chaîne.
//
// Paramètre :
// - s : la chaîne de caractères à traiter.
//
// Retour :
// - la chaîne modifiée avec chaque paire de caractères inversée.
//
// Fonctionnement :
// - Convertit la chaîne en un tableau de runes.
// - Parcourt la liste par pas de deux, en échangeant chaque paire.
// - Reconstruit et retourne la chaîne modifiée.
func InverserDeuxParDeux(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i += 2 {
		runes[i], runes[i+1] = runes[i+1], runes[i]
	}
	return string(runes)
}
