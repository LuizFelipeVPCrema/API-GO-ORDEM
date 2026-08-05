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

	Modificacoes []ModificacaoResponse `json:"modificacoes"`
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
	return EquipamentoDetalhadoResponse{
		EquipamentoResponse: NovoEquipamentoResponse(
			equipamento.Equipamento,
		),
		Modificacoes: NovasModificacoesResponse(
			equipamento.Modificacoes,
		),
	}
}
