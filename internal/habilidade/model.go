package habilidade

type Codigo string

type Categoria string

const (
	CategoriaHabilidadeClasse Categoria = "HABILIDADE_CLASSE"
	CategoriaPoderClasse      Categoria = "PODER_CLASSE"
	CategoriaHabilidadeOrigem Categoria = "HABILIDADE_ORIGEM"
	CategoriaHabilidadeTrilha Categoria = "HABILIDADE_TRILHA"
	CategoriaPoderParanormal  Categoria = "PODER_PARANORMAL"
	CategoriaPoderGeral       Categoria = "PODER_GERAL"
	CategoriaOutra            Categoria = "OUTRA"
)

type TipoAtivacao string

const (
	AtivacaoPassiva   TipoAtivacao = "PASSIVA"
	AtivacaoLivre     TipoAtivacao = "LIVRE"
	AtivacaoMovimento TipoAtivacao = "MOVIMENTO"
	AtivacaoPadrao    TipoAtivacao = "PADRAO"
	AtivacaoCompleta  TipoAtivacao = "COMPLETA"
	AtivacaoReacao    TipoAtivacao = "REACAO"
	AtivacaoEspecial  TipoAtivacao = "ESPECIAL"
)

type FormaAquisicao string

const (
	AquisicaoAutomatica FormaAquisicao = "AUTOMATICA"
	AquisicaoEscolha    FormaAquisicao = "ESCOLHA"
)

type Habilidade struct {
	ID     int64
	Codigo Codigo
	Nome   string

	Categoria    Categoria
	TipoAtivacao TipoAtivacao

	CustoPEBase     *int
	CustoPEVariavel bool

	DescricaoResumida *string

	FonteRegra  string
	VersaoRegra *string

	Ativa bool
}

type HabilidadeClasse struct {
	ID int64

	ClasseID      int64
	Habilidade    Habilidade
	NEXMinimo     int
	Aquisicao     FormaAquisicao
	OrdemExibicao int
}
