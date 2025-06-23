package goFonctions

import "encoding/hex"

// HexToIntLittleEndian convertit une chaîne hexadécimale en un entier en respectant l'ordre Little-endian.
//
// Paramètre :
// - hexString : une chaîne hexadécimale représentant le nombre.
//
// Retour :
// - l'entier correspondant à la valeur hexadécimale en Little-endian, ou 0 en cas d'erreur.
//
// Fonctionnement :
// - Décode la chaîne hexadécimale en un slice de bytes.
// - Inverse l'ordre des octets pour respecter le format Little-endian.
// - Convertit les bytes en un entier en utilisant la décomposition par décalage.
func HexToIntLittleEndian(hexString string) int {
	// Convertir la chaîne hexadécimale en un slice de bytes
	bytesData, err := hex.DecodeString(hexString)
	if err != nil {
		return 0
	}

	// Inverser les octets pour respecter l'ordre Little-endian
	for i := 0; i < len(bytesData)/2; i++ {
		bytesData[i], bytesData[len(bytesData)-1-i] = bytesData[len(bytesData)-1-i], bytesData[i]
	}

	// Convertir les bytes en entier (Little-endian)
	var result int
	for i := 0; i < len(bytesData); i++ {
		result |= int(bytesData[i]) << (8 * i)
	}

	return result
}
