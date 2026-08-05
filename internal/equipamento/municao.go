package equipamento

type UnidadeDuracao string

const (
	DuracaoCena     UnidadeDuracao = "CENA"
	DuracaoDisparo  UnidadeDuracao = "DISPARO"
	DuracaoUso      UnidadeDuracao = "USO"
	DuracaoMissao   UnidadeDuracao = "MISSAO"
	DuracaoEspecial UnidadeDuracao = "ESPECIAL"
)

type Municao struct {
	EquipamentoID int64

	DuracaoQuantidade *int
	DuracaoUnidade    UnidadeDuracao

	Consumivel bool
}

type MunicaoCompativel struct {
	Equipamento Equipamento
	Municao     Municao
}
