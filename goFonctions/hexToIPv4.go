package goFonctions

import (
	"fmt"
)

// HexToIPv4 convertit une chaîne hexadécimale de 8 caractères en adresse IPv4.
//
// Paramètre :
// - hex : une chaîne hexadécimale de 8 caractères représentant une adresse IPv4.
//
// Retour :
// - une chaîne formatée "x.x.x.x" correspondant à l'adresse IPv4, ou une chaîne vide si l'entrée est invalide.
//
// Fonctionnement :
// - Vérifie que la chaîne fait exactement 8 caractères.
// - Extrait chaque octet en utilisant fmt.Sscanf.
// - Vérifie que chaque octet est compris entre 0 et 255.
// - Construit et retourne l'adresse IP sous forme de chaîne.
func HexToIPv4(hex string) string {
	if len(hex) != 8 {
		return ""
	}

	// Création d'un tableau pour les octets
	ip := make([]byte, 4)
	for i := 0; i < 4; i++ {
		var byteValue int
		_, err := fmt.Sscanf(hex[i*2:i*2+2], "%x", &byteValue)
		if err != nil || byteValue < 0 || byteValue > 255 {
			return ""
		}
		ip[i] = byte(byteValue)
	}

	// Retourne l'adresse IP sous forme de chaîne
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}
