package goFonctions

import "encoding/hex"

// EbcdicToString convertit une chaîne hexadécimale EBCDIC en une chaîne ASCII.
//
// Paramètre :
// - ebcdicHex : une chaîne hexadécimale représentant des caractères EBCDIC.
//
// Retour :
// - une chaîne ASCII correspondant à la conversion. Les caractères inconnus sont remplacés par '?'.
//
// Fonctionnement :
// - Décode la chaîne hexadécimale en un tableau de bytes.
// - Pour chaque byte, convertit en ASCII en utilisant une table de correspondance.
// - Si la correspondance n'existe pas, utilise '?'.
// - Retourne la chaîne ASCII résultante.
func EbcdicToString(ebcdicHex string) string {
	// Conversion de la chaîne hexadécimale en tableau de bytes
	ebcdicBytes, _ := hex.DecodeString(ebcdicHex)

	var asciiChars []byte
	for _, ebcdicByte := range ebcdicBytes {
		asciiChar, exists := ebcdicToAscii[ebcdicByte]
		if !exists {
			asciiChar = '?'
		}
		asciiChars = append(asciiChars, asciiChar)
	}
	return string(asciiChars)
}
