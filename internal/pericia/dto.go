package pericia

type PericiaResponse struct {
	ID                    int64  `json:"id"`
	Codigo                string `json:"codigo"`
	Nome                  string `json:"nome"`
	AtributoBase          string `json:"atributo_base"`
	SomenteTreinada       bool   `json:"somente_treinada"`
	PenalidadeCarga       bool   `json:"penalidade_carga"`
	PermiteEspecializacao bool   `json:"permite_especializacao"`
	OrdemExibicao         int    `json:"ordem_exibicao"`
}

func NovaPericiaResponse(pericia Pericia) PericiaResponse {
	return PericiaResponse{
		ID:                    pericia.ID,
		Codigo:                string(pericia.Codigo),
		Nome:                  pericia.Nome,
		AtributoBase:          string(pericia.AtributoBase),
		SomenteTreinada:       pericia.SomenteTreinada,
		PenalidadeCarga:       pericia.PenalidadeCarga,
		PermiteEspecializacao: pericia.PermiteEspecializacao,
		OrdemExibicao:         pericia.OrdemExibicao,
	}
}

func NovasPericiasResponse(pericias []Pericia) []PericiaResponse {
	respostas := make([]PericiaResponse, 0, len(pericias))

	for _, periciaAtual := range pericias {
		respostas = append(respostas, NovaPericiaResponse(periciaAtual))
	}

	return respostas
}
