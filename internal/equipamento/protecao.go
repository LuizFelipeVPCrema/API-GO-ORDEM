package equipamento

type TipoProtecao string

const (
	ProtecaoLeve   TipoProtecao = "LEVE"
	ProtecaoPesada TipoProtecao = "PESADA"
	ProtecaoEscudo TipoProtecao = "ESCUDO"
	ProtecaoOutra  TipoProtecao = "OUTRA"
)

type Protecao struct {
	EquipamentoID int64

	TipoProtecao TipoProtecao

	BonusDefesa int

	PenalidadeTeste int
}
