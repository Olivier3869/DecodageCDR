package main

import (
	"fmt"
	"local/goFonctions"
)

// conversionChamp convertit un champ donné en fonction d'un code de conversion spécifié.
//
// Paramètres :
// - champ (string) : la chaîne de caractères à convertir.
// - codage (int) : un entier représentant le type de conversion à appliquer.
//
// Retourne :
// - string : la chaîne d'origine suivie de la conversion correspondante, ou la chaîne d'origine si le codage n'est pas reconnu.
//
// Codes de conversion pris en charge :
// - 1 : Conversion hexadécimale en chaîne (HexToString) : 323038303137393030363234303337 => 208017900624037
// - 2 : Conversion EBCDIC en chaîne (EbcdicToString)
// - 3 : Conversion hexadécimale en adresse IPv4 (HexToIPv4)
// - 5 : Inversion de la chaîne par paires (InverserDeuxParDeux) : 803386094000f0 => 0833689004000f
// - 6 : Formatage de la date et l'heure en GMT (FormatDateTimeGMT) : 2506050801522b0200 => 05/06/2025 08:01:52 +0200
// - 8 : Formatage de la date et l'heure (FormatDateTime)
// - 9 : Formatage de la date et l'heure avec un format spécifique (FormatDateTime8)
// - 10 : Conversion hexadécimale en décimal (HexToDecimal) : 3f => 63
// - 11 : Conversion hexadécimale en entier (HexToIntLittleEndian)
// - 12 : Formatage de la date et l'heure inversé (FormatDateTimeInv)
// - 13 : Formatage de la date et l'heure avec un format spécifique (FormatDateTime14)
// - 14 : Conversion de la date (ConvertDate)
// - 15 : Conversion de l'heure (ConvertTime)
//
// Si le codage n'est pas reconnu, la fonction retourne simplement le champ d'origine.

func conversionChamp(champ string, codage int) string {
	switch codage {
	case 1:
		return champ + " => " + goFonctions.HexToString(champ)
	case 2:
		return champ + " => " + goFonctions.EbcdicToString(champ)
	case 3:
		return champ + " => " + goFonctions.HexToIPv4(champ)
	case 5:
		return champ + " => " + goFonctions.InverserDeuxParDeux(champ)
	case 6:
		return champ + " => " + goFonctions.FormatDateTimeGMT(champ)
	case 7:
		return champ + " => " + goFonctions.HexTimestampToDate(champ)
	case 8:
		return champ + " => " + goFonctions.FormatDateTime(champ)
	case 9:
		return champ + " => " + goFonctions.FormatDateTime8(champ)
	case 10:
		return champ + " => " + goFonctions.HexToDecimal(champ)
	case 11:
		return champ + " => " + fmt.Sprintf("%d", goFonctions.HexToIntLittleEndian(champ))
	case 12:
		return champ + " => " + goFonctions.FormatDateTimeInv(champ)
	case 13:
		return champ + " => " + goFonctions.FormatDateTime14(champ)
	case 14:
		return champ + " => " + goFonctions.ConvertDate(champ)
	case 15:
		return champ + " => " + goFonctions.ConvertTime(champ)
	default:
		return champ
	}
}
