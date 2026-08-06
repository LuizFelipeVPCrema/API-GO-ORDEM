package personagem

type Personagem struct {
	ID int64

	Nome    string
	Jogador *string

	ClasseID int64

	NEX             int
	PontosPrestigio int

	Idade *int

	Aparencia     *string
	Personalidade *string
	Historia      *string
	Objetivo      *string

	Atributos Atributos
	Recursos  Recursos

	Ativa bool
}

type Atributos struct {
	Agilidade int
	Forca     int
	Intelecto int
	Presenca  int
	Vigor     int
}

type Recursos struct {
	PVAtual  int
	PVMaximo int

	PEAtual  int
	PEMaximo int

	SanidadeAtual  int
	SanidadeMaxima int
}

type Filtro struct {
	Nome     string
	ClasseID *int64

	Limit  int
	Offset int
}
