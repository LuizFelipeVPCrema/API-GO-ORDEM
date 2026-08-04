package habilidade

import "errors"

var (
	ErrHabilidadeNaoEncontrada = errors.New("habilidade nao encontrada")
	ErrClasseNaoEncontrada     = errors.New("classe nao encontrada")
	ErrIDInvalido              = errors.New("id da habilidade invalido")
	ErrClasseIDInvalido        = errors.New("id da classe invalido")
	ErrCodigoInvalido          = errors.New("codigo da habilidade invalido")
)
