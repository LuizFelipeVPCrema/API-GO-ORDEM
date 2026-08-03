package patente

import "context"

type Service struct {
	repository Repository
}

func NovoService(
	repository Repository,
) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Listar(
	ctx context.Context,
) ([]Patente, error) {
	return s.repository.Listar(ctx)
}

func (s *Service) BuscarPorID(
	ctx context.Context,
	id int64,
) (*Patente, error) {
	if err := validarID(id); err != nil {
		return nil, err
	}

	return s.repository.BuscarPorID(ctx, id)
}

func (s *Service) BuscarPorPrestigio(
	ctx context.Context,
	pontos int,
) (*Patente, error) {
	if err := validarPontosPrestigio(pontos); err != nil {
		return nil, err
	}

	return s.repository.BuscarPorPrestigio(
		ctx,
		pontos,
	)

}
