package main

import (
	"flag"
	"fmt"
	"local/goFonctions"
	"local/structures/defStructs"
	"path/filepath"
	"strings"
)

var typeTraitement *string

// var equipement *string
var afficherIndex *bool
var afficherEntete *bool
var afficherDataCra *bool
var afficherAide *bool
var afficherVersion *bool
var decoderColl bool
var structureCRA = make(map[int]defStructs.TabRechercheTypeDc)
var format string
var protocole string
var idAudit string

var versDecodeCDR string
var versStructures string
var versFonctions string

// Généré par script table.py dans A_projets
const listeTT = `
      METEOR/PALLADIUM                      COMETE              
┌─────────────┬─────────────┐ ┌───────────────┬────────────────┐
│ EEC         │ ECC         │ │           EEC │ ECC            │
├─────────────┼─────────────┤ ├───────────────┼────────────────┤
│ TT02F_HP    │ Dclientip   │ │         TTNBI │ D16xconv       │
│ TTaliasNb   │ Ddpvgsm     │ │         TTcgw │ Dafdb          │
│ TTasbc      │ Ddpvgsm32   │ │       TTnodes │ DallegroNew    │
│ TTbgw_gsn   │ DdpvgsmFree │ │        TTorca │ Dandromeri     │
│ TTchfcdr    │ Dpvgsm28    │ │       TTsonus │ Dcaps          │
│ TTcrae      │ Dsymp_home  │ │    TTsvaipTAG │ Dcasper        │
│ TTepdg      │ Dtapin      │ │     TTv2oipAS │ DcentaureBIV   │
│ TTims       │             │ │    TTv2oipSBC │ DcentaureDC    │
│ TTines      │             │ │    TTvoipvpn2 │ DcentaureHF    │
│ TTlmisf     │             │ │         TTwha │ DcentaureIP    │
│ TTmmm       │             │ │ TTwhaInternet │ DcentaureVG    │
│ TTmmscXura  │             │ │   TTwhaSelect │ DdeltaMinute   │
│ TTmmTelTas  │             │ │        trait1 │ Ddqd           │
│ TTmvf       │             │ │               │ Dentropi       │
│ TTnss9      │             │ │               │ Dentropi_sonus │
│ TTpgwcdr    │             │ │               │ Dfe            │
│ TTpgwe      │             │ │               │ Dferpv         │
│ TTsgwcdr    │             │ │               │ DkenobiWha     │
│ TTSms_cmg   │             │ │               │ Dmeteor        │
│ TTSms_mcohr │             │ │               │ Dmmdm          │
│ TTsmsoIP    │             │ │               │ DmmdmWha       │
│ TTsvaip     │             │ │               │ Dprims         │
│ TTtap310ZB  │             │ │               │ Drevod         │
│ TTtasnokia  │             │ │               │ Dspot          │
│ TTussd      │             │ │               │ Dsymphonie     │
│ TTwha       │             │ │               │ Dtrafic        │
└─────────────┴─────────────┘ └───────────────┴────────────────┘`

func main() {

	versDecodeCDR = "1.31"
	versStructures = "109"
	versFonctions = "23"

	// Déclaration des drapeaux pour les arguments de ligne de commande
	filePath := flag.String("f", "", "Chemin du fichier")
	typeTraitement = flag.String("t", "", "Type de décodeur")
	//equipement = flag.String("e", "", "Equipement")
	afficherIndex = flag.Bool("i", false, "Affiche les indices de CRAs")
	afficherEntete = flag.Bool("s", false, "Affiche les entêtes")
	afficherDataCra = flag.Bool("d", false, "Decode les CRAs")
	afficherAide = flag.Bool("h", false, "Affiche l'aide")
	afficherVersion = flag.Bool("v", false, "Affiche la version")

	// Analyse des arguments de ligne de commande
	flag.Parse()

	// Affichage de l'aide
	if *afficherAide == true {
		afficheAide()
		return
	}

	// Affichage de l'aide
	if *afficherVersion == true {
		fmt.Printf("Versions : decodeCDR  v%v\n", versDecodeCDR)
		fmt.Printf("           structures v%v\n", versStructures)
		fmt.Printf("           fonctions  v%v\n", versFonctions)
		return
	}

	// Vérifiez si un nom de fichier est fourni
	if *filePath == "" {
		fmt.Println("ERREUR: Un nom de fichier (-f) doit être spécifié.")
		fmt.Println("")
		afficheAide()
		return
	}

	// Vérifiez si le type de traitement est fourni
	// TODO qd -e effectif à corriger
	if *typeTraitement == "" {
		//fmt.Println("ERREUR : nom de l'équipement (-e) doit être spécifié ou ")
		fmt.Println("ERREUR : Le type de décodeur (-t) doit être spécifié :")
		fmt.Println(listeTT)
		return
	}

	// Si l'équipement est précisé ; recherche du type de décodage
	// TODO recherche du type de décodage dans la base mongo

	// Appeler la fonction pour lire le fichier binaire
	data, err := goFonctions.ReadBinaryFile(*filePath)
	if err != nil {
		fmt.Println("Erreur lors de la lecture du fichier :", err)
		return
	}

	filename := filepath.Base(*filePath)
	if strings.HasPrefix(filename, "coll_") {
		decoderColl = true
	} else {
		decoderColl = false
	}

	structureCRA, format, protocole = selectionStructure(*typeTraitement)
	//fmt.Println(structureCRA)
	if format == "TLV" {
		decodeTLV(data)
	} else if format == "FIXE" {
		decodeFIXE(data)
	} else if format == "CSV" {
		decodeCSV(data)
	} else if format == "ASN1" {
		decodeASN1(data)
	} else if format == "TASNP" {
		decodeTASNP(data)
	} else if format == "CHF" {
		decodeASN1(data)
	} else {
		fmt.Printf("Le type de traitement %v n'est pas pris en compte", *typeTraitement)
	}
}

func afficheAide() {
	fmt.Printf("decodeCDR v%v.%v.%v\n", versDecodeCDR, versStructures, versFonctions)
	fmt.Println("")
	fmt.Println("Syntaxe de la commande : decodeCDR -f <nomFichier> -t <typeDécodeur> [-d] [-s] [-i]")
	fmt.Println("    option -d : décodage des CDR")
	fmt.Println("    option -s : affichage des entêtes")
	fmt.Println("    option -i : affichage numéro du CDR")
	fmt.Println("    option -t : types de décodeurs disponibles :")
	fmt.Println(listeTT)
}
