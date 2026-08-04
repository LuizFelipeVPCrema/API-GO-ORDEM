package habilidade

import "context"

type Service struct {
	repository Repository
}

func NovoService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Listar(ctx context.Context) ([]Habilidade, error) {
	return s.repository.Listar(ctx)
}

func (s *Service) BuscarPorID(ctx context.Context, id int64) (*Habilidade, error) {
	if err := validarID(id); err != nil {
		return nil, err
	}

	return s.repository.BuscarPorID(ctx, id)
}

func (s *Service) BuscarPorCodigo(ctx context.Context, codigo string) (*Habilidade, error) {
	codigoNormalizado, err := normalizarCodigo(codigo)

	if err != nil {
		return nil, err
	}

	return s.repository.BuscarPorCodigo(ctx, codigoNormalizado)
}

func (s *Service) ListarPorClasseID(ctx context.Context, classeID int64) ([]HabilidadeClasse, error) {
	if err := validarClasseID(classeID); err != nil {
		return nil, err
	}

	return s.repository.ListarPorClasseID(ctx, classeID)
}
