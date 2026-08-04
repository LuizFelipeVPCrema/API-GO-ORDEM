package classe

type ClasseResponse struct {
	ID            int64  `json:"id"`
	Codigo        string `json:"codigo"`
	Nome          string `json:"nome"`
	OrdemExibicao int    `json:"ordem_exibicao"`
}

func NovaClasseResponse(classeEncontrada Classe) ClasseResponse {
	return ClasseResponse{
		ID:            classeEncontrada.ID,
		Codigo:        string(classeEncontrada.Codigo),
		Nome:          classeEncontrada.Nome,
		OrdemExibicao: classeEncontrada.OrdemExibicao,
	}
}

func NovasClassesResponse(classes []Classe) []ClasseResponse {
	respostas := make([]ClasseResponse, 0, len(classes))

	for _, classeAtual := range classes {
		respostas = append(respostas, NovaClasseResponse(classeAtual))
	}

	return respostas
}
