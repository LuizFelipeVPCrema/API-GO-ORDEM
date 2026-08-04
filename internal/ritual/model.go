package ritual

type Codigo string

type Elemento string

const (
	ElementoSangue       Elemento = "SANGUE"
	ElementoMorte        Elemento = "MORTE"
	ElementoConhecimento Elemento = "CONHECIMENTO"
	ElementoEnergia      Elemento = "ENERGIA"
	ElementoMedo         Elemento = "MEDO"
)

type Circulo int

const (
	PrimeiroCirculo Circulo = iota + 1
	SegundoCirculo
	TerceiroCirculo
	QuartoCirculo
	QuintoCirculo
)

type TipoExecucao string

const (
	ExecucaoLivre     TipoExecucao = "LIVRE"
	ExecucaoMovimento TipoExecucao = "MOVIMENTO"
	ExecucaoPadrao    TipoExecucao = "PADRAO"
	ExecucaoCompleta  TipoExecucao = "COMPLETA"
	ExecucaoReacao    TipoExecucao = "REACAO"
	ExecucaoEspecial  TipoExecucao = "ESPECIAL"
)

type TipoAprimoramento string

const (
	AprimoramentoDiscente   TipoAprimoramento = "DISCENTE"
	AprimoramentoVerdadeiro TipoAprimoramento = "VERDADEIRO"
	AprimoramentoOutro      TipoAprimoramento = "OUTRO"
)

type Ritual struct {
	ID     int64
	Codigo Codigo
	Nome   string

	Elemento Elemento
	Circulo  Circulo
	Execucao TipoExecucao

	Alcance string
	Alvo    *string
	Area    *string
	Duracao string

	Resistencia *string

	CustoPEBase      int
	RequerComponente bool

	DescricaoResumida *string

	FonteRegra       string
	VersaoRegra      *string
	PaginaReferencia *int

	Ativo bool
}

type Aprimoramento struct {
	ID       int64
	RitualID int64

	Tipo TipoAprimoramento

	CustoPEAdicional int

	NEXMinimo         *int
	CirculoMinimo     *Circulo
	DescricaoResumida *string

	OrdemExibicao int
	Ativo         bool
}

type RitualDetalhado struct {
	Ritual
	Aprimoramentos []Aprimoramento
}

type Filtro struct {
	Elemento *Elemento
	Circulo  *Circulo
}
