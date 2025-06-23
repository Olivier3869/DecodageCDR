package goFonctions

// IsNumeric vérifie si une chaîne de caractères ne contient que des chiffres.
//
// Paramètre :
// - s : la chaîne à vérifier.
//
// Retour :
// - true si la chaîne ne contient que des chiffres (0-9), sinon false.
//
// Fonctionnement :
// - Parcourt chaque caractère de la chaîne.
// - Vérifie si le caractère est compris entre '0' et '9'.
// - Retourne false dès qu'un caractère non numérique est rencontré, sinon true.
func IsNumeric(s string) bool {
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}
