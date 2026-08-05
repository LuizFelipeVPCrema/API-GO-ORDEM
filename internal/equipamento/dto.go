package equipamento

type FiltroConsulta struct {
	Tipo      string
	Categoria *int
}

type EquipamentoResponse struct {
	ID     int64  `json:"id"`
	Codigo string `json:"codigo"`
	Nome   string `json:"nome"`
	Tipo   string `json:"tipo"`

	CategoriaBase       int    `json:"categoria_base"`
	CategoriaBaseRotulo string `json:"categoria_base_rotulo"`
	EspacosBase         int    `json:"espacos_base"`

	DescricaoResumida *string `json:"descricao_resumida,omitempty"`

	FonteRegra       string  `json:"fonte_regra"`
	VersaoRegra      *string `json:"versao_regra,omitempty"`
	PaginaReferencia *int    `json:"pagina_referencia,omitempty"`
}

type ModificacaoResponse struct {
	ID     int64  `json:"id"`
	Codigo string `json:"codigo"`
	Nome   string `json:"nome"`

	IncrementoCategoria int `json:"incremento_categoria"`
	IncrementoEspacos   int `json:"incremento_espacos"`

	LimitePorItem int  `json:"limite_por_item"`
	Cumulativa    bool `json:"cumulativa"`

	DescricaoResumida *string `json:"descricao_resumida,omitempty"`

	FonteRegra       string  `json:"font_regra"`
	VersaoRegra      *string `json:"versao_regra,omitempty"`
	PaginaReferencia *int    `json:"pagina_referencia,omitempty"`
}

type EquipamentoDetalhadoResponse struct {
	EquipamentoResponse

	Arma     *ArmaResponse     `json:"arma,omitempty"`
	Protecao *ProtecaoResponse `json:"protecao,omitempty"`
	Municao  *MunicaoResponse  `json:"municao,omitempty"`

	Modificacoes []ModificacaoResponse `json:"modificacoes"`

	MunicoesCompativeis []MunicaoCompativelResponse `json:"municoes_compativeis,omitempty"`
}

type ArmaResponse struct {
	TipoArma string `json:"tipo_arma"`

	DanoBase string `json:"dano_base"`
	TipoDano string `json:"tipo_dano"`

	CriticoMargem        int `json:"critico_margem"`
	CriticoMultiplicador int `json:"critico_multiplicador"`

	Alcance     string `json:"alcance"`
	Empunhadura string `json:"empunhadura"`

	Recarga *string `json:"recarga,omitempty"`
}

type ProtecaoResponse struct {
	TipoProtecao string `json:"tipo_protecao"`

	BonusDefesa      int `json:"bonus_defesa"`
	PenalidadeTestes int `json:"penalidade_testes"`
}

type MunicaoResponse struct {
	DuracaoQuantidade *int   `json:"duracao_quantidade,omitempty"`
	DuracaoUnidade    string `json:"duracao_unidade"`

	Consumivel bool `json:"consumivel"`
}

type MunicaoCompativelResponse struct {
	Equipamento EquipamentoResponse `json:"equipamento"`
	Municao     MunicaoResponse     `json:"municao"`
}

func rotuloCategoria(categoria Categoria) string {
	switch categoria {
	case CategoriaZero:
		return "0"
	case CategoriaI:
		return "I"
	case CategoriaII:
		return "II"
	case CategoriaIII:
		return "III"
	case CategoriaIV:
		return "IV"
	default:
		return ""
	}
}

func NovoEquipamentoResponse(equipamento Equipamento) EquipamentoResponse {
	return EquipamentoResponse{
		ID:     equipamento.ID,
		Codigo: string(equipamento.Codigo),
		Nome:   equipamento.Nome,
		Tipo:   string(equipamento.Tipo),

		CategoriaBase:       int(equipamento.CategoriaBase),
		CategoriaBaseRotulo: rotuloCategoria(equipamento.CategoriaBase),
		EspacosBase:         equipamento.EspacosBase,

		DescricaoResumida: equipamento.DescricaoResumida,

		FonteRegra:       equipamento.FonteRegra,
		VersaoRegra:      equipamento.VersaoRegra,
		PaginaReferencia: equipamento.PaginaReferencia,
	}
}

func NovosEquipamentosResponse(equipamentos []Equipamento) []EquipamentoResponse {
	respostas := make([]EquipamentoResponse, 0, len(equipamentos))

	for _, equipamentoAtual := range equipamentos {
		respostas = append(respostas, NovoEquipamentoResponse(equipamentoAtual))
	}

	return respostas
}

func NovaModificacaoResponse(modificacao Modificacao) ModificacaoResponse {
	return ModificacaoResponse{
		ID:     modificacao.ID,
		Codigo: string(modificacao.Codigo),
		Nome:   modificacao.Nome,

		IncrementoCategoria: modificacao.IncrementoCategoria,
		IncrementoEspacos:   modificacao.IncrementoEspacos,

		LimitePorItem: modificacao.LimitePorItem,
		Cumulativa:    modificacao.Cumulativa,

		DescricaoResumida: modificacao.DescricaoResumida,

		FonteRegra:       modificacao.FonteRegra,
		VersaoRegra:      modificacao.VersaoRegra,
		PaginaReferencia: modificacao.PaginaReferencia,
	}
}

func NovasModificacoesResponse(modificacoes []Modificacao) []ModificacaoResponse {
	respostas := make(
		[]ModificacaoResponse,
		0,
		len(modificacoes),
	)

	for _, modificacaoAtual := range modificacoes {
		respostas = append(
			respostas,
			NovaModificacaoResponse(
				modificacaoAtual,
			),
		)
	}

	return respostas
}

func NovoEquipamentoDetalhadoResponse(equipamento EquipamentoDetalhado) EquipamentoDetalhadoResponse {
	resposta := EquipamentoDetalhadoResponse{
		EquipamentoResponse: NovoEquipamentoResponse(
			equipamento.Equipamento,
		),

		Modificacoes: NovasModificacoesResponse(
			equipamento.Modificacoes,
		),

		MunicoesCompativeis: NovasMunicoesCompativeisResponse(
			equipamento.MunicaoCompativeis,
		),
	}

	if equipamento.Arma != nil {
		arma := NovaArmaResponse(
			*equipamento.Arma,
		)

		resposta.Arma = &arma
	}

	if equipamento.Protecao != nil {
		protecao := NovaProtecaoResponse(
			*equipamento.Protecao,
		)

		resposta.Protecao = &protecao
	}

	if equipamento.Municao != nil {
		municao := NovaMunicaoResponse(
			*equipamento.Municao,
		)

		resposta.Municao = &municao
	}

	return resposta
}

func NovaArmaResponse(
	arma Arma,
) ArmaResponse {
	return ArmaResponse{
		TipoArma: string(arma.TipoArma),

		DanoBase: arma.DanoBase,
		TipoDano: arma.TipoDano,

		CriticoMargem:        arma.CriticoMargem,
		CriticoMultiplicador: arma.CriticoMultiplicador,

		Alcance:     string(arma.Alcance),
		Empunhadura: string(arma.Empunhadura),

		Recarga: arma.Recarga,
	}
}

func NovaProtecaoResponse(
	protecao Protecao,
) ProtecaoResponse {
	return ProtecaoResponse{
		TipoProtecao: string(
			protecao.TipoProtecao,
		),

		BonusDefesa: protecao.BonusDefesa,

		PenalidadeTestes: protecao.PenalidadeTeste,
	}
}

func NovaMunicaoResponse(
	municao Municao,
) MunicaoResponse {
	return MunicaoResponse{
		DuracaoQuantidade: municao.DuracaoQuantidade,

		DuracaoUnidade: string(
			municao.DuracaoUnidade,
		),

		Consumivel: municao.Consumivel,
	}
}

func NovasMunicoesCompativeisResponse(
	municoes []MunicaoCompativel,
) []MunicaoCompativelResponse {
	respostas := make(
		[]MunicaoCompativelResponse,
		0,
		len(municoes),
	)

	for _, municao := range municoes {
		respostas = append(
			respostas,
			MunicaoCompativelResponse{
				Equipamento: NovoEquipamentoResponse(
					municao.Equipamento,
				),
				Municao: NovaMunicaoResponse(
					municao.Municao,
				),
			},
		)
	}

	return respostas
}
