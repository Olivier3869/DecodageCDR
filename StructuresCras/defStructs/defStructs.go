package defStructs

type TabRechercheTypeDc struct {
	Nom      string
	Masque   int
	Valeur   int
	Longueur int
	Tag      *map[int]Champ
}

type Champ struct {
	Nom      string
	Longueur int
	Codage   int
	Tag      *map[int]Champ
}
