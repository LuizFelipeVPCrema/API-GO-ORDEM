package equipamento

type Codigo string

type Tipo string

const (
	TipoArma            Tipo = "ARMA"
	TipoProtecao        Tipo = "PROTECAO"
	TipoAcessorio       Tipo = "ACESSORIO"
	TipoMunicao         Tipo = "MUNICAO"
	TipoItemOperacional Tipo = "ITEM_OPERACIONAL"
	TipoItemParanormal  Tipo = "ITEM_PARANORMAL"
	TipoOutro           Tipo = "OUTRO"
)

type TipoAplicabilidade string

const (
	AplicabilidadeArma            TipoAplicabilidade = "ARMA"
	AplicabilidadeProtecao        TipoAplicabilidade = "PROTECAO"
	AplicabilidadeAcessorio       TipoAplicabilidade = "ACESSORIO"
	AplicabilidadeMunicao         TipoAplicabilidade = "MUNICAO"
	AplicabilidadeItemOperacional TipoAplicabilidade = "ITEM_OPERACIONAL"
	AplicabilidadeItemParanormal  TipoAplicabilidade = "ITEM_PARANORMAL"
	AplicabilidadeOutro           TipoAplicabilidade = "OUTRO"
	AplicabilidadeTodos           TipoAplicabilidade = "TODOS"
)

type Categoria int

const (
	CategoriaZero Categoria = iota
	CategoriaI
	CategoriaII
	CategoriaIII
	CategoriaIV
)

type Equipamento struct {
	ID     int64
	Codigo Codigo
	Nome   string
	Tipo   Tipo

	CategoriaBase Categoria
	EspacosBase   int

	DescricaoResumida *string

	FonteRegra       string
	VersaoRegra      *string
	PaginaReferencia *int

	Ativa bool
}

type Modificacao struct {
	ID     int64
	Codigo Codigo
	Nome   string

	IncrementoCategoria int
	IncrementoEspacos   int

	LimitePorItem int
	Cumulativa    bool

	DescricaoResumida *string

	FonteRegra       string
	VersaoRegra      *string
	PaginaReferencia *int

	Ativa bool
}

type EquipamentoDetalhado struct {
	Equipamento
	Modificacoes []Modificacao
}

type Filtro struct {
	Tipo      *Tipo
	Categoria *Categoria
}

type ModificacaoAplicada struct {
	Modificacao Modificacao
	Quantidade  int
}

type ResultadoModificacoes struct {
	CategoriaFinal Categoria
	EspacosFinais  int
}
