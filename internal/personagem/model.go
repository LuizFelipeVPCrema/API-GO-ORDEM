package personagem

import "time"

type Personagem struct {
	ID        uint
	UsuarioID uint

	OrigemID  uint
	ClasseID  uint
	TrilhaID  uint
	PatenteID uint
	NexID     uint

	Nome        string
	NomeJogador string
	Idade       *int

	Aparencia     string
	Personalidade string
	Historia      string

	Versao int

	CriadoEm     time.Time
	AtualizadoEm time.Time
	ExcluidoEm   time.Time
}

type FichaCompleta struct {
	Personagem

	// Atributos  Atributos
	// Recursos   Recursos
	// Pericias   []Pericia
	// Rituais    []Ritual
	// Habilidades []Habilidade
	// Itens      []Item
}
