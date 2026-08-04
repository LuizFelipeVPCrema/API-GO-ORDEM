package classe

import "errors"

var (
	ErrClasseNaoEncontrada = errors.New("classe nao encontrada")
	ErrIDInvalido          = errors.New("id da classe invalido")
	ErrCodigoInvalido      = errors.New("codigo da classe invalido")
)
