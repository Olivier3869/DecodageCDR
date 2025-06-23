# StructuresCras

Version : 110

Ce dépot héberge les structures de CRAs communes aux outils de décodage de CRAs pour :  
 - **norego**, l'outil de NONREG
 - **decodeCDR** pour le décodage des CRAs  

## Formats des données


## Utilisation

Pour utiliser les fonctions dans un autre projet `Go` il faut configurer vote projet avec l'utilisation d'un fichier `go.mod`.  
Dans ce fichier, ajouter les lignes :

```mod
replace local/structures => ../StructuresCras
require local/structures v0.0.0
```

Pour chaque fichier Go ou il est nécessaire d'utiliser une structure, il faut ajouter l'`import` :

```go
import (
"local/structures/defStructs"
structuresEccCSV "local/structures/structuresEccCSV"
structuresEccFIXE "local/structures/structuresEccFIXE"
structuresEccTLV "local/structures/structuresEccTLV"
)
```

Et pour utiliser une des structures : 

```go
func selectionStructure(tag string) (map[int]defStructs.TabRechercheTypeDc, string) {
	switch {
	// ECC format TLV
	case tag == "Dafdb":
		return structuresEccTLV.Dafdb, "TLV"
```

## Structures 

Les structures sont disponibles pour les flux suivants :
```txt
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
└─────────────┴─────────────┘ └───────────────┴────────────────┘ 
```