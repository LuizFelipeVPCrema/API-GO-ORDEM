package pericia

import "errors"

var (
	ErrPericiaNaoEncontrada = errors.New("pericia nao encontrada")
	ErrIDInvalido           = errors.New("id da pericia invalido")
	ErrCodigoInvalido       = errors.New("codigo da pericia invalido")
)
