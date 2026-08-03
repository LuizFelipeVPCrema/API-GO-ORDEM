package patente

import "errors"

var (
	ErrPatenteNaoEncontrada = errors.New("patente não encontrada")
	ErrPatenteInativa       = errors.New("patente inativa")
	ErrIDInvalido           = errors.New("id da patente inválido")
	ErrPrestigioInvalido    = errors.New("pontos de prestígio inválidos")
)
