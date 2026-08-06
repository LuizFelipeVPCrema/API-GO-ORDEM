package personagem

type CriarPersonagemRequest struct {
	Nome    string  `json:"nome"`
	Jogador *string `json:"jogador"`

	ClasseID int64 `json:"classe_id"`

	NEX             int `json:"nex"`
	PontosPrestigio int `json:"pontos_prestigio"`

	Idade *int `json:"idade"`

	Aparencia     *string `json:"aparencia"`
	Personalidade *string `json:"personalidade"`
	Historia      *string `json:"historia"`
	Objetivo      *string `json:"objetivo"`

	Atributos AtributosRequest `json:"atributos"`
	Recursos  RecursosRequest  `json:"recursos"`
}

type AtributosRequest struct {
	Agilidade int `json:"agilidade"`
	Forca     int `json:"forca"`
	Intelecto int `json:"intelecto"`
	Presenca  int `json:"presenca"`
	Vigor     int `json:"vigor"`
}

type RecursosRequest struct {
	PVAtual  int `json:"pv_atual"`
	PVMaximo int `json:"pv_maximo"`

	PEAtual  int `json:"pe_atual"`
	PEMaximo int `json:"pe_maximo"`

	SanidadeAtual  int `json:"sanidade_atual"`
	SanidadeMaxima int `json:"sanidade_maxima"`
}

type AtualizarPersonagemRequest struct {
	Nome    *string `json:"nome"`
	Jogador *string `json:"jogador"`

	ClasseID *int64 `json:"classe_id"`

	NEX             *int `json:"nex"`
	PontosPrestigio *int `json:"pontos_prestigio"`

	Idade *int `json:"idade"`

	Aparencia     *string `json:"aparencia"`
	Personalidade *string `json:"personalidade"`
	Historia      *string `json:"historia"`
	Objetivo      *string `json:"objetivo"`

	Atributos *AtualizarAtributosRequest `json:"atributos"`
	Recursos  *AtualizarRecursosRequest  `json:"recursos"`
}

type AtualizarAtributosRequest struct {
	Agilidade *int `json:"agilidade"`
	Forca     *int `json:"forca"`
	Intelecto *int `json:"intelecto"`
	Presenca  *int `json:"presenca"`
	Vigor     *int `json:"vigor"`
}

type AtualizarRecursosRequest struct {
	PVAtual  *int `json:"pv_atual"`
	PVMaximo *int `json:"pv_maximo"`

	PEAtual  *int `json:"pe_atual"`
	PEMaximo *int `json:"pe_maximo"`

	SanidadeAtual  *int `json:"sanidade_atual"`
	SanidadeMaxima *int `json:"sanidade_maxima"`
}

type AtributosResponse struct {
	Agilidade int `json:"agilidade"`
	Forca     int `json:"forca"`
	Intelecto int `json:"intelecto"`
	Presenca  int `json:"presenca"`
	Vigor     int `json:"vigor"`
}

type RecursosResponse struct {
	PVAtual  int `json:"pv_atual"`
	PVMaximo int `json:"pv_maximo"`

	PEAtual  int `json:"pe_atual"`
	PEMaximo int `json:"pe_maximo"`

	SanidadeAtual  int `json:"sanidade_atual"`
	SanidadeMaxima int `json:"sanidade_maxima"`
}

type PersonagemResponse struct {
	ID int64 `json:"id"`

	Nome    string  `json:"nome"`
	Jogador *string `json:"jogador,omitempty"`

	ClasseID int64 `json:"classe_id"`

	NEX             int `json:"nex"`
	PontosPrestigio int `json:"pontos_prestigio"`

	Idade *int `json:"idade,omitempty"`

	Aparencia     *string `json:"aparencia,omitempty"`
	Personalidade *string `json:"personalidade,omitempty"`
	Historia      *string `json:"historia,omitempty"`
	Objetivo      *string `json:"objetivo,omitempty"`

	Atributos AtributosResponse `json:"atributos"`
	Recursos  RecursosResponse  `json:"recursos"`
}

func NovoPersonagemResponse(
	personagem Personagem,
) PersonagemResponse {
	return PersonagemResponse{
		ID: personagem.ID,

		Nome:    personagem.Nome,
		Jogador: personagem.Jogador,

		ClasseID: personagem.ClasseID,

		NEX:             personagem.NEX,
		PontosPrestigio: personagem.PontosPrestigio,

		Idade: personagem.Idade,

		Aparencia:     personagem.Aparencia,
		Personalidade: personagem.Personalidade,
		Historia:      personagem.Historia,
		Objetivo:      personagem.Objetivo,

		Atributos: AtributosResponse{
			Agilidade: personagem.Atributos.Agilidade,
			Forca:     personagem.Atributos.Forca,
			Intelecto: personagem.Atributos.Intelecto,
			Presenca:  personagem.Atributos.Presenca,
			Vigor:     personagem.Atributos.Vigor,
		},

		Recursos: RecursosResponse{
			PVAtual:  personagem.Recursos.PVAtual,
			PVMaximo: personagem.Recursos.PVMaximo,

			PEAtual:  personagem.Recursos.PEAtual,
			PEMaximo: personagem.Recursos.PEMaximo,

			SanidadeAtual:  personagem.Recursos.SanidadeAtual,
			SanidadeMaxima: personagem.Recursos.SanidadeMaxima,
		},
	}
}

func NovosPersonagensResponse(
	personagens []Personagem,
) []PersonagemResponse {
	respostas := make(
		[]PersonagemResponse,
		0,
		len(personagens),
	)

	for _, personagemAtual := range personagens {
		respostas = append(
			respostas,
			NovoPersonagemResponse(personagemAtual),
		)
	}

	return respostas
}
