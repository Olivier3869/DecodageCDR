package goFonctions

import (
	"encoding/binary"
)

// BytesToIntLittleEndian convertit un tableau de bytes en un entier (Little Endian).
//
// Paramètres :
// - data : un tableau de bytes ([]byte) à convertir.
//
// Retour :
// - un entier (int) représentant la valeur convertie. Retourne 0 si la conversion n'est pas possible.
//
// Fonctionnement :
// - Vérifie si le tableau est vide ou trop grand (plus de 8 octets).
// - Selon la taille du tableau, convertit en utilisant la représentation Little Endian pour 1, 2 ou 4 octets.
// - Si la taille n'est pas supportée, retourne 0.
func BytesToIntLittleEndian(data []byte) int { // , error supprimé
	// Vérification de la longueur du tableau
	if len(data) == 0 {
		return 0 //, errors.New("le tableau de bytes est vide")
	}

	// Si le tableau est plus grand que 8 octets, nous ne pouvons pas le convertir en un int64
	if len(data) > 8 {
		return 0 //, errors.New("taille du tableau de bytes trop grande, doit être entre 1 et 8 octets")
	}

	// Conversion du tableau de bytes en entier selon la taille du tableau
	var value int
	switch len(data) {
	case 1:
		value = int(data[0])
	case 2:
		value = int(binary.LittleEndian.Uint16(data))
	case 4:
		value = int(binary.LittleEndian.Uint32(data))
	//case 8: value = int64(binary.BigEndian.Uint64(data))
	default:
		return 0 // , fmt.Errorf("taille du tableau non supportée: %d octets", len(data))
	}

	// Retourner la valeur décodée
	return value // , nil
}
