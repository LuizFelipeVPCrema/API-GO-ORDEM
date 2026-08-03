package pericia

import (
	"context"
	"strings"
)

type Service struct {
	repository Repository
}

func NovoService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Listar(ctx context.Context) ([]Pericia, error) {
	return s.repository.Listar(ctx)
}

func (s *Service) BuscarPorID(ctx context.Context, id int64) (*Pericia, error) {
	if err := validarID(id); err != nil {
		return nil, err
	}

	return s.repository.BuscarPorID(ctx, id)
}

func (s *Service) BuscarPorCodigo(ctx context.Context, codigo string) (*Pericia, error) {
	codigoNormalizado, err := normalizarCodigo(codigo)
	if err != nil {
		return nil, err
	}

	return s.repository.BuscarPorCodigo(ctx, codigoNormalizado)
}

func validarID(id int64) error {
	if id <= 0 {
		return ErrIDInvalido
	}

	return nil
}

func normalizarCodigo(codigo string) (Codigo, error) {
	codigoNormalizado := strings.ToUpper(
		strings.TrimSpace(codigo),
	)

	if codigoNormalizado == "" {
		return "", ErrCodigoInvalido
	}

	for _, caractere := range codigoNormalizado {
		ehLetraMaiuscula :=
			caractere >= 'A' && caractere <= 'Z'

		ehSeparador := caractere == '_'

		if !ehLetraMaiuscula && !ehSeparador {
			return "", ErrCodigoInvalido
		}
	}

	return Codigo(codigoNormalizado), nil
}
