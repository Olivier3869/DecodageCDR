package goFonctions

import "fmt"

// ByteArrayToHexString convertit un tableau d'octets en une chaîne hexadécimale.
//
// Paramètres :
// - byteArray : un tableau de bytes ([]byte) à convertir.
//
// Retour :
// - une chaîne de caractères (string) représentant la version hexadécimale du tableau d'octets.
//
// La fonction parcourt chaque octet du tableau, le formate en hexadécimal à deux chiffres,
// puis concatène le résultat dans une chaîne finale.
func ByteArrayToHexString(byteArray []byte) string {
	hexString := ""
	for _, b := range byteArray {
		hexString += fmt.Sprintf("%02x", b)
	}
	return hexString
}
