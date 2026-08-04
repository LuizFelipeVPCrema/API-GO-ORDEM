package ritual

import "errors"

var (
	ErrRitualNaoEncontrado = errors.New("ritual nao encontrado")
	ErrIDInvalido          = errors.New("id do ritual invalido")
	ErrCodigoInvalido      = errors.New("codigo do ritual invalido")
	ErrElementoInvalido    = errors.New("elemento do ritual invalido")
	ErrCirculoInvalido     = errors.New("circulo do ritual invalido")
)
