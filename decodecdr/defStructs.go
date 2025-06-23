package main

// Cdr représente un enregistrement de données CDR (Call Detail Record).
type Cdr struct {
	CdrByte []byte   // CdrByte contient les données brutes du CDR sous forme de tableau d'octets.
	Champs  []Values // Champs est un tableau de valeurs associées au CDR, contenant des informations supplémentaires.
}

// Values représentent les valeurs associées à un tag dans un CDR.
type Values struct {
	Tag      int    // Tag identifie le champ dans le CDR.
	Nom      string // Nom est le nom descriptif du champ.
	Valeur   []byte // Valeur contient la valeur brute associée au tag sous forme de tableau d'octets.
	Codage   int    // Codage spécifie le type de codage utilisé pour la valeur.
	Decalage int    // Decalage indique un éventuel décalage d'affichage à appliquer.
}
