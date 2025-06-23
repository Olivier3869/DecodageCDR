package goFonctions

import "encoding/binary"

// EstEntete vérifie si un tableau de bytes correspond à une entête spécifique.
//
// Paramètre :
// - data : un tableau de bytes ([]byte).
//
// Retour :
// - bool : true si c'est une entête reconnue, false sinon.
// - int : la longueur de l'entête si reconnu, 0 sinon.
//
// Fonctionnement :
// - Convertit la donnée en un entier 64 bits en utilisant l'endianness BigEndian.
// - Applique des masques pour comparer avec des valeurs spécifiques.
// - Retourne true avec la longueur associée si une correspondance est trouvée.
// - Sinon, retourne false et 0.
func EstEntete(data []byte) (bool, int) {
	value := binary.BigEndian.Uint64(append(make([]byte, 2), data...))
	// ATTENTION : les entetes en entrée sont plus longs que ceux en sortie ...
	if value&0xFFFFFFFFFFFF == 0x0a2d2d2d2075 {
		// Cas des séparateurs RabbitMQ PALLADIUM ECC
		return true, 65
	} else if value&0xFFFFFFFFFFFF == 0x2d2d2d207577 {
		// Cas des entêtes PALLADIUM ECC
		return true, 64
	} else if value&0x000000FFFFFF == 0x000000017200 {
		// Cas des entêtes MQSeries COMETE Symphonie
		return true, 18
	} else if value&0x000000FFFFFF == 0x00000001eb00 {
		// Cas des entêtes MQSeries COMETE Casper
		return true, 18
	} else if value&0x000000FFFFFF == 0x00000001f900 {
		// Cas des entêtes MQSeries METEOR Symphonie HOME
		return true, 18
	} else if value&0x000000FFFFFF == 0x00000001b100 {
		// Cas des entêtes MQSeries METEOR B Symphonie GSM
		return true, 18
	} else if value&0x000000FFFFFF == 0x00000001f200 {
		// Cas des entêtes MQSeries METEOR Symphonie GSM
		return true, 18
	} else if value&0x000000FFFFFF == 0x00000001d200 {
		// Cas des entêtes MQSeries METEOR A Symphonie GSM
		return true, 18
	} else {
		return false, 0
	}

}
