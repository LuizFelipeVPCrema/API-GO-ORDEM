package patente

import "time"

type Codigo string

const (
	CodigoRecruta          Codigo = "RECRUTA"
	CodigoOperador         Codigo = "OPERADOR"
	CodigoAgenteEspecial   Codigo = "AGENTE_ESPECIAL"
	CodigoOficialOperacoes Codigo = "OFICIAL_OPERACOES"
	CodigoAgenteElite      Codigo = "AGENTE_ELITE"
)

type LimiteCredito string

const (
	CreditoBaixo     LimiteCredito = "BAIXO"
	CreditoMedio     LimiteCredito = "MEDIO"
	CreditoAlto      LimiteCredito = "ALTO"
	CreditoIlimitado LimiteCredito = "ILIMITADO"
)

type CategoriaItem int

const (
	CartegoriaI   CategoriaItem = 1
	CartegoriaII  CategoriaItem = 2
	CartegoriaIII CategoriaItem = 3
	CartegoriaIV  CategoriaItem = 4
)

type LimiteItem struct {
	Categoria        CategoriaItem
	QuantidadeMaxima int
}

type Patente struct {
	ID                     uint
	Codigo                 Codigo
	Nome                   string
	PontosPrestigioMinimos int
	LimiteCredito          LimiteCredito
	NivelHierarquico       int
	Ativa                  bool
	Limites                []LimiteItem
	CriadoEm               time.Time
	AtualizadoEm           time.Time
	ExcluidoEm             time.Time
}

type PatenteComLimites struct {
	Patente
	Limites []LimiteItem
}
