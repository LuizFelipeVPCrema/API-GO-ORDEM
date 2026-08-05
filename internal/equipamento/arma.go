package equipamento

type TipoArma string

const (
	ArmaCorpoACorpo TipoArma = "CORPO_A_CORPO"
	ArmaDistancia   TipoArma = "DISTANCIA"
)

type Empunhadura string

const (
	EmpunhaduraLeve     Empunhadura = "LEVE"
	EmpunhaduraUmaMao   Empunhadura = "UMA_MAO"
	EmpunhaduraDuasMaos Empunhadura = "DUAS_MAOS"
)

type Alcance string

const (
	AlcanceCorpoACorpo Alcance = "CORPO_A_CORPO"
	AlcanceCurto       Alcance = "CURTO"
	AlcanceMedio       Alcance = "MEDIO"
	AlcanceLongo       Alcance = "LONGO"
	AlcanceExtremo     Alcance = "EXTREMO"
)

type Arma struct {
	EquipamentoID int64

	TipoArma TipoArma

	DanoBase string
	TipoDano string

	CriticoMargem        int
	CriticoMultiplicador int

	Alcance     Alcance
	Empunhadura Empunhadura

	Recarga *string
}

type ExpressaoDano struct {
	QuantidadeDados int
	Faces           int
	Bonus           int
}
