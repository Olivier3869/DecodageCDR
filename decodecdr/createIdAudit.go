package main

import (
	"fmt"
	"local/goFonctions"
)

// createIdAudit génère un identifiant d'audit basé sur le compteur et les paramètres globaux.
//
// Paramètres :
// - cpt (int) : le compteur utilisé pour générer l'identifiant d'audit.
//
// Retourne :
// - string : l'identifiant d'audit formaté.
func createIdAudit(cpt int) string {
	audit := ""
	//fmt.Println(idAudit)
	if *afficherIndex == true {
		if decoderColl == true {
			audit = fmt.Sprintf("%v%d_0_0_0 : ", idAudit, cpt)
			//} else if strings.HasPrefix(*typeTraitement, "T") && protocole == "" {
			//	audit = fmt.Sprintf("0.0.0000000000_0000_%d_0_0_0 : ", cpt)
		} else if *typeTraitement == "TTines" && len(idAudit) > 0 {
			audit = fmt.Sprintf("%v_%d : ", goFonctions.HexToString(idAudit[78:182]), cpt)
		} else if *typeTraitement == "TTchfcdr" && len(idAudit) > 0 {
			audit = fmt.Sprintf("%v_%d : ", goFonctions.HexToString(idAudit[78:182]), cpt)
		} else {
			audit = fmt.Sprintf("CDR_%09d : ", cpt)
		}
	}
	return audit
}
