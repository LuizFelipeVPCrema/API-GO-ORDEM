package ritual

import "context"

type Service struct {
	repository Repository
}

func NovoService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) Listar(ctx context.Context, consulta FiltroConsulta) ([]Ritual, error) {
	filtro := Filtro{}

	if consulta.Elemento != "" {
		elemento, err := normalizarElemento(consulta.Elemento)
		if err != nil {
			return nil, err
		}

		filtro.Elemento = &elemento
	}

	if consulta.Circulo != nil {
		circulo := Circulo(*consulta.Circulo)

		if err := validarCirculo(circulo); err != nil {
			return nil, err
		}

		filtro.Circulo = &circulo
	}

	return s.repository.Listar(ctx, filtro)
}

func (s *Service) BuscarPorID(ctx context.Context, id int64) (*RitualDetalhado, error) {
	if err := validarID(id); err != nil {
		return nil, err
	}

	ritualEncontrado, err := s.repository.BuscarPorID(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.carregarDetalhes(ctx, ritualEncontrado)
}

func (s *Service) BuscarPorCodigo(ctx context.Context, codigo string) (*RitualDetalhado, error) {
	codigoNormalizado, err := normalizarCodigo(codigo)
	if err != nil {
		return nil, err
	}

	ritualEncontrado, err := s.repository.BuscarPorCodigo(ctx, codigoNormalizado)
	if err != nil {
		return nil, err
	}

	return s.carregarDetalhes(ctx, ritualEncontrado)
}

func (s *Service) ListarAprimoramentos(ctx context.Context, ritualID int64) ([]Aprimoramento, error) {
	if err := validarID(ritualID); err != nil {
		return nil, err
	}

	_, err := s.repository.BuscarPorID(ctx, ritualID)
	if err != nil {
		return nil, err
	}

	return s.repository.ListarAprimoramentos(ctx, ritualID)

}

func (s *Service) carregarDetalhes(ctx context.Context, ritualEncontrado *Ritual) (*RitualDetalhado, error) {
	aprimoramentos, err :=
		s.repository.ListarAprimoramentos(
			ctx,
			ritualEncontrado.ID,
		)

	if err != nil {
		return nil, err
	}

	return &RitualDetalhado{
		Ritual:         *ritualEncontrado,
		Aprimoramentos: aprimoramentos,
	}, nil
}
