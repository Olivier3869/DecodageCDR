package main

import (
	"fmt"
	"local/goFonctions"
	"strings"
)

// AfficherDetailCRA affiche les détails d'un CRA (Call Detail Record).
// Cette fonction principale orchestre l'affichage des données CDR et des champs associés.
func AfficherDetailCRA(CRA Cdr) {
	// Étape 1 : Affichage de l'identifiant d'audit et des données CDR brutes
	afficherCdr(CRA)

	// Étape 2 : Affichage détaillé des champs du CRA si l'option est activée
	// *afficherDataCra est un flag global contrôlant ce niveau de détail
	if *afficherDataCra {
		// Parcours et affichage de tous les champs contenus dans le CRA
		for _, value := range CRA.Champs {
			afficherChamp(value)
		}
	}
}

// afficherCdr gère l'affichage des données CDR principales selon le format choisi
func afficherCdr(CRA Cdr) {
	// Format CSV : affichage des données brutes en tant que string
	if format == "CSV" {
		// createIdAudit(cpt) génère un identifiant unique pour l'audit
		// CRA.CdrByte contient les données binaires du CDR
		fmt.Printf("%v%s\n", createIdAudit(cpt), string(CRA.CdrByte))
	} else {
		// Autres formats : conversion des bytes en représentation hexadécimale
		// Plus lisible pour le débogage et l'analyse technique
		fmt.Printf("%v%v\n", createIdAudit(cpt), goFonctions.ByteArrayToHexString(CRA.CdrByte))
	}
}

// afficherChamp traite l'affichage d'un champ individuel selon le format spécifié
func afficherChamp(value Values) {
	// Création de l'indentation pour représenter la hiérarchie des données
	spaces := strings.Repeat(" ", value.Decalage)

	// Gestion du nom du champ (avec fallback si nom inconnu)
	nomChamp := value.Nom
	if nomChamp == "" {
		// Si le nom n'est pas défini, création d'un nom par défaut avec le tag
		nomChamp = fmt.Sprintf("INCONNU : TAG %v", value.Tag)
	}

	// Dispatch vers la fonction d'affichage appropriée selon le format
	switch format {
	case "TLV": // Tag-Length-Value format
		afficherChampTLV(value, spaces)
	case "ASN1": // Abstract Syntax Notation One format
		afficherChampASN1(value, spaces, nomChamp)
	case "TASNP": // Format personnalisé (TASNP)
		afficherChampTASNP(value, spaces, nomChamp)
	case "CSV": // Comma-Separated Values format
		fmt.Printf("%v%v : %s\n", createIdAudit(cpt), value.Nom, string(value.Valeur))
	case "FIXE": // Format à longueur fixe
		afficherChampFixe(value)
	}
}

// afficherChampTLV gère l'affichage spécifique au format TLV
func afficherChampTLV(value Values, spaces string) {
	// Tag spécial 9999 indique une valeur finale/leaf
	if value.Tag == 9999 {
		// Conversion de la valeur binaire en hexadécimal
		valeur := goFonctions.ByteArrayToHexString(value.Valeur)
		// Application d'une conversion selon le type d'encodage du champ
		valeurTraduite := conversionChamp(valeur, value.Codage)
		fmt.Printf("%v%v%v = %v\n", createIdAudit(cpt), spaces, value.Nom, valeurTraduite)
	} else {
		// Pour les autres tags, affichage du nom et tag en hexadécimal
		fmt.Printf("%v%v%v(%02x) \n", createIdAudit(cpt), spaces, value.Nom, value.Tag)
	}
}

// afficherChampASN1 gère l'affichage spécifique au format ASN.1
func afficherChampASN1(value Values, spaces string, nomChamp string) {
	// Vérification de la présence d'une valeur
	if value.Valeur != nil {
		// Traitement similaire au TLV mais avec syntaxe ASN.1
		valeur := goFonctions.ByteArrayToHexString(value.Valeur)
		valeurTraduite := conversionChamp(valeur, value.Codage)
		// Format: Nom(Tag) = Valeur
		fmt.Printf("%v%v%v(%v) = %v\n", createIdAudit(cpt), spaces, nomChamp, value.Tag, valeurTraduite)
	} else {
		// Si pas de valeur, affichage uniquement du nom et tag (nœud parent)
		fmt.Printf("%v%v%v(%v)\n", createIdAudit(cpt), spaces, nomChamp, value.Tag)
	}
}

// afficherChampTASNP combine les logiques TLV et ASN1 selon le contexte
func afficherChampTASNP(value Values, spaces string, nomChamp string) {
	// Même logique spéciale pour le tag 9999
	if value.Tag == 9999 {
		valeur := goFonctions.ByteArrayToHexString(value.Valeur)
		valeurTraduite := conversionChamp(valeur, value.Codage)
		// Format simplifié sans parenthèses pour tag 9999
		fmt.Printf("%v%v%v = %v\n", createIdAudit(cpt), spaces, nomChamp, valeurTraduite)
	} else {
		// Pour les autres tags, comportement identique à ASN1
		if value.Valeur != nil {
			valeur := goFonctions.ByteArrayToHexString(value.Valeur)
			valeurTraduite := conversionChamp(valeur, value.Codage)
			fmt.Printf("%v%v%v(%v) = %v\n", createIdAudit(cpt), spaces, nomChamp, value.Tag, valeurTraduite)
		} else {
			fmt.Printf("%v%v%v(%v)\n", createIdAudit(cpt), spaces, nomChamp, value.Tag)
		}
	}
}

// afficherChampFixe gère l'affichage pour le format à longueur fixe
func afficherChampFixe(value Values) {
	// Traitement direct sans indentation ni tag
	valeur := goFonctions.ByteArrayToHexString(value.Valeur)
	valeurTraduite := conversionChamp(valeur, value.Codage)
	// Format simple : Nom : Valeur
	fmt.Printf("%v%v : %v\n", createIdAudit(cpt), value.Nom, valeurTraduite)
}
