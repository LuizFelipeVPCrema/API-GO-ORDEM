package habilidade

type HabilidadeResponse struct {
	ID     int64  `json:"id"`
	Codigo string `json:"codigo"`
	Nome   string `json:"nome"`

	Categoria    string `json:"categoria"`
	TipoAtivacao string `json:"tipo_ativacao"`

	CustoPEBase     *int `json:"custo_pe_base,omitempty"`
	CustoPEVariavel bool `json:"custo_pe_variavel"`

	DescricaoResumida *string `json:"descricao_resumida,omitempty"`

	FonteRegra  string  `json:"fonte_regra"`
	VersaoRegra *string `json:"versao_regra,omitempty"`
}

type HabilidadeClasseResponse struct {
	ID       int64 `json:"id"`
	ClasseID int64 `json:"classe_id"`

	NEXMinimo int    `json:"nex_minimo"`
	Aquisicao string `json:"aquisicao"`

	OrdemExibicao int `json:"ordem_exibicao"`

	Habilidade HabilidadeResponse `json:"habilidade"`
}

func NovaHabilidadeResponse(habilidade Habilidade) HabilidadeResponse {
	return HabilidadeResponse{
		ID:     habilidade.ID,
		Codigo: string(habilidade.Codigo),
		Nome:   habilidade.Nome,

		Categoria:    string(habilidade.Categoria),
		TipoAtivacao: string(habilidade.TipoAtivacao),

		CustoPEBase:     habilidade.CustoPEBase,
		CustoPEVariavel: habilidade.CustoPEVariavel,

		DescricaoResumida: habilidade.DescricaoResumida,

		FonteRegra:  habilidade.FonteRegra,
		VersaoRegra: habilidade.VersaoRegra,
	}
}

func NovasHabilidadesResponse(habilidades []Habilidade) []HabilidadeResponse {
	respostas := make([]HabilidadeResponse, 0, len(habilidades))

	for _, habilidadeAtual := range habilidades {
		respostas = append(respostas, NovaHabilidadeResponse(habilidadeAtual))
	}

	return respostas
}

func NovaHabilidadeClasseResponse(vinculo HabilidadeClasse) HabilidadeClasseResponse {
	return HabilidadeClasseResponse{
		ID:       vinculo.ID,
		ClasseID: vinculo.ClasseID,

		NEXMinimo: vinculo.NEXMinimo,
		Aquisicao: string(vinculo.Aquisicao),

		OrdemExibicao: vinculo.OrdemExibicao,

		Habilidade: NovaHabilidadeResponse(vinculo.Habilidade),
	}
}

func NovasHabilidadesClasseResponse(vinculos []HabilidadeClasse) []HabilidadeClasseResponse {
	respostas := make([]HabilidadeClasseResponse, 0, len(vinculos))

	for _, vinculo := range vinculos {
		respostas = append(respostas, NovaHabilidadeClasseResponse(vinculo))
	}

	return respostas
}
