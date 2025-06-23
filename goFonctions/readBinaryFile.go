package goFonctions

import (
	"encoding/binary"
	"os"
)

// ReadBinaryFile lit le contenu binaire d'un fichier.
//
// Paramètre :
// - filename : le chemin vers le fichier à lire.
//
// Retour :
// - un tableau d'octets contenant le contenu du fichier,
// - une erreur si une opération échoue.
//
// Fonctionnement :
// - Ouvre le fichier en mode lecture.
// - Récupère la taille du fichier pour allouer un tableau d'octets.
// - Lit le contenu du fichier dans le tableau en utilisant binary.Read avec l'endianness LittleEndian.
// - Ferme le fichier avant de retourner.
func ReadBinaryFile(filename string) ([]byte, error) {
	// Ouvrir le fichier en mode lecture
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Obtenir la taille du fichier pour allouer un tableau d'octets
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}
	fileSize := fileInfo.Size()

	// Lire le contenu du fichier dans un tableau d'octets
	data := make([]byte, fileSize)
	if err := binary.Read(file, binary.LittleEndian, &data); err != nil {
		return nil, err
	}

	return data, nil
}
