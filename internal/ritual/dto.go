package ritual

type FiltroConsulta struct {
	Elemento string
	Circulo  *int
}

type AprimoramentoResponse struct {
	ID int64 `json:"id"`

	Tipo string `json:"tipo"`

	CustoPEAdicional int `json:"custo_pe_adicional"`

	NEXMinimo     *int `json:"nex_minimo,omitempty"`
	CirculoMinimo *int `json:"circulo_minimo,omitempty"`

	DescricaoResumida *string `json:"descricao_resumida,omitempty"`

	OrdemExibicao int `json:"ordem_exibicao"`
}

type RitualResponse struct {
	ID     int64  `json:"id"`
	Codigo string `json:"codigo"`
	Nome   string `json:"nome"`

	Elemento string `json:"elemento"`
	Circulo  int    `json:"circulo"`
	Execucao string `json:"execucao"`

	Alcance string  `json:"alcance"`
	Alvo    *string `json:"alvo,omitempty"`
	Area    *string `json:"area,omitempty"`
	Duracao string  `json:"duracao"`

	Resistencia *string `json:"resistencia,omitempty"`

	CustoPEBase      int  `json:"custo_pe_base"`
	RequerComponente bool `json:"requer_componente"`

	DescricaoResumida *string `json:"descricao_resumida,omitempty"`

	FonteRegra       string  `json:"fonte_regra"`
	VersaoRegra      *string `json:"versao_regra,omitempty"`
	PaginaReferencia *int    `json:"pagina_referencia,omitempty"`
}

type RitualDetalhadoResponse struct {
	RitualResponse

	Aprimoramentos []AprimoramentoResponse `json:"aprimoramentos"`
}

func NovoRitualResponse(ritual Ritual) RitualResponse {
	return RitualResponse{
		ID:     ritual.ID,
		Codigo: string(ritual.Codigo),
		Nome:   ritual.Nome,

		Elemento: string(ritual.Elemento),
		Circulo:  int(ritual.Circulo),
		Execucao: string(ritual.Execucao),

		Alcance: ritual.Alcance,
		Alvo:    ritual.Alvo,
		Area:    ritual.Area,
		Duracao: ritual.Duracao,

		Resistencia: ritual.Resistencia,

		CustoPEBase:      ritual.CustoPEBase,
		RequerComponente: ritual.RequerComponente,

		DescricaoResumida: ritual.DescricaoResumida,

		FonteRegra:       ritual.FonteRegra,
		VersaoRegra:      ritual.VersaoRegra,
		PaginaReferencia: ritual.PaginaReferencia,
	}
}

func NovosRituaisResponse(rituais []Ritual) []RitualResponse {
	respostas := make([]RitualResponse, 0, len(rituais))

	for _, ritualAtual := range rituais {
		respostas = append(respostas, NovoRitualResponse(ritualAtual))
	}

	return respostas
}

func NovoAprimoramentoResponse(aprimoramento Aprimoramento) AprimoramentoResponse {
	var circuloMinimo *int

	if aprimoramento.CirculoMinimo != nil {
		valor := int(*aprimoramento.CirculoMinimo)
		circuloMinimo = &valor
	}

	return AprimoramentoResponse{
		ID: aprimoramento.ID,

		Tipo: string(aprimoramento.Tipo),

		CustoPEAdicional: aprimoramento.CustoPEAdicional,

		NEXMinimo:     aprimoramento.NEXMinimo,
		CirculoMinimo: circuloMinimo,

		DescricaoResumida: aprimoramento.DescricaoResumida,

		OrdemExibicao: aprimoramento.OrdemExibicao,
	}
}

func NovosAprimoramentosResponse(aprimoramentos []Aprimoramento) []AprimoramentoResponse {
	respostas := make(
		[]AprimoramentoResponse,
		0,
		len(aprimoramentos),
	)

	for _, aprimoramento := range aprimoramentos {
		respostas = append(
			respostas,
			NovoAprimoramentoResponse(
				aprimoramento,
			),
		)
	}

	return respostas
}

func NovoRitualDetalhadoResponse(ritual RitualDetalhado) RitualDetalhadoResponse {
	aprimoramentos := make([]AprimoramentoResponse, 0, len(ritual.Aprimoramentos))

	for _, aprimoramento := range ritual.Aprimoramentos {
		aprimoramentos = append(aprimoramentos, NovoAprimoramentoResponse(aprimoramento))
	}

	return RitualDetalhadoResponse{
		RitualResponse: NovoRitualResponse(ritual.Ritual),
		Aprimoramentos: aprimoramentos,
	}
}
