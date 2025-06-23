package goFonctions

import (
	"fmt"
	"io"
	"os"
)

// CopyFile copie le contenu d'un fichier source vers un fichier de destination.
//
// Paramètres :
// - src : chemin du fichier source.
// - dst : chemin du fichier de destination.
//
// Retour :
// - une erreur si une opération échoue, sinon nil.
//
// Fonctionnement :
// - Ouvre le fichier source en lecture.
// - Crée ou écrase le fichier de destination.
// - Copie le contenu du fichier source dans le fichier destination.
// - Ferme les fichiers après opération.
func CopyFile(src, dst string) error {
	// Ouvrir le fichier source
	source, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("erreur lors de l'ouverture du fichier source : %v", err)
	}
	defer source.Close()

	// Créer le fichier de destination
	destination, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("erreur lors de la création du fichier de destination : %v", err)
	}
	defer destination.Close()

	// Copier le contenu du fichier source vers le fichier de destination
	_, err = io.Copy(destination, source)
	if err != nil {
		return fmt.Errorf("erreur lors de la copie du fichier : %v", err)
	}

	return nil
}
