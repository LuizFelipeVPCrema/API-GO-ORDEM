package classe

type Codigo string

const (
	CodigoCombatente   Codigo = "COMBATENTE"
	CodigoEspecialista Codigo = "ESPECIALISTA"
	CodigoOcultista    Codigo = "OCULTISTA"
	CodigoSobrevivente Codigo = "SOBREVIVENTE"
)

type Classe struct {
	ID            int64
	Codigo        Codigo
	Nome          string
	OrdemExibicao int
	Ativa         bool
}
