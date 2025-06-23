package goFonctions

import "fmt"

// ConvertToStringSlice convertit un slice de interfaces ([]interface{}) en un slice de chaînes ([]string).
//
// Paramètre :
// - input : un slice de interfaces ([]interface{}).
//
// Retour :
// - un slice de chaînes ([]string) contenant les éléments convertis.
// - une erreur si un élément n'est pas une chaîne.
//
// Fonctionnement :
// - Parcourt chaque élément du slice d'entrée.
// - Vérifie si l'élément est une chaîne.
// - Si oui, l'ajoute au slice de sortie.
// - Sinon, retourne une erreur indiquant l'index de l'élément non conforme.
func ConvertToStringSlice(input []interface{}) ([]string, error) {
	groupingFields := make([]string, len(input))
	for i, item := range input {
		str, ok := item.(string)
		if !ok {
			return nil, fmt.Errorf("élément à l'index %d n'est pas une chaîne", i)
		}
		groupingFields[i] = str
	}
	return groupingFields, nil
}
