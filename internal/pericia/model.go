package pericia

type Codigo string

type AtributoBase string

const (
	AtributoAgilidade AtributoBase = "AGI"
	AtributoForca     AtributoBase = "FOR"
	AtributoIntelecto AtributoBase = "INT"
	AtributoPresenca  AtributoBase = "PRE"
	AtributoVigor     AtributoBase = "VIG"
)

const (
	CodigoProfissao Codigo = "PROFISSAO"
)

type Pericia struct {
	ID     int64
	Codigo Codigo
	Nome   string

	AtributoBase AtributoBase

	SomenteTreinada       bool
	PenalidadeCarga       bool
	PermiteEspecializacao bool

	OrdemExibicao int
	Ativa         bool
}
