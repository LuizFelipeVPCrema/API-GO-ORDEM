package patente

type LimiteItemResponse struct {
	Categoria        int `json:"categoria"`
	QuantidadeMaxima int `json:"quantidade_maxima"`
}

type PatenteResponse struct {
	ID                     int64                `json:"id"`
	Codigo                 string               `json:"codigo"`
	Nome                   string               `json:"nome"`
	PontosPrestigioMinimos int                  `json:"pontos_prestigio_minimos"`
	LimiteCredito          string               `json:"limite_credito"`
	NivelHierarquico       int                  `json:"nivel_hierarquico"`
	Limites                []LimiteItemResponse `json:"limites"`
}

func NovaPatenteResponse(p Patente) PatenteResponse {
	limites := make(
		[]LimiteItemResponse,
		0,
		len(p.Limites),
	)

	for _, limite := range p.Limites {
		limites = append(
			limites,
			LimiteItemResponse{
				Categoria:        int(limite.Categoria),
				QuantidadeMaxima: limite.QuantidadeMaxima,
			},
		)
	}

	return PatenteResponse{
		ID:                     int64(p.ID),
		Codigo:                 string(p.Codigo),
		Nome:                   p.Nome,
		PontosPrestigioMinimos: p.PontosPrestigioMinimos,
		LimiteCredito:          string(p.LimiteCredito),
		NivelHierarquico:       p.NivelHierarquico,
		Limites:                limites,
	}
}

func NovasPatentesResponse(
	patentes []Patente,
) []PatenteResponse {
	respostas := make(
		[]PatenteResponse,
		0,
		len(patentes),
	)

	for _, p := range patentes {
		respostas = append(
			respostas,
			NovaPatenteResponse(p),
		)
	}

	return respostas
}
