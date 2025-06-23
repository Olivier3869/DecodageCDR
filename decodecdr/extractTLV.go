package main

// extractTLV extrait un élément TLV (Type-Length-Value) à partir des données fournies.
//
// Paramètres :
// - data ([]byte) : un tableau d'octets contenant les données à décoder.
// - index (int) : l'index à partir duquel commencer l'extraction.
//
// Retourne :
// - []interface{} : un tableau contenant le type TLV, la longueur, la valeur et la longueur totale de l'élément TLV.
// - int : l'index mis à jour après l'extraction de l'élément TLV.
func extractTLV(data []byte, index int) ([]interface{}, int, int) {
	var tlvType int   // int
	var tlvLength int // int
	typeLen := 0
	lengthLen := 0

	// Extraire le type
	for (data[index] & 0x80) == 0x80 {
		tlvType = (tlvType << 7) | int(data[index]&0x7F) // int
		index++
		typeLen++
	}
	tlvType = (tlvType << 7) | int(data[index]&0x7F) // int
	index++
	typeLen++

	// Extraire la longueur
	for (data[index] & 0x80) == 0x80 {
		tlvLength = (tlvLength << 7) | int(data[index]&0x7F) // int
		index++
		lengthLen++
	}
	tlvLength = (tlvLength << 7) | int(data[index]&0x7F) // int
	index++
	lengthLen++

	// Extraire la valeur
	tlvValue := data[index : index+tlvLength]
	index += tlvLength
	indexTL := typeLen + lengthLen
	totalLength := typeLen + lengthLen + tlvLength

	return []interface{}{tlvType, tlvLength, tlvValue, totalLength}, index, indexTL
}
