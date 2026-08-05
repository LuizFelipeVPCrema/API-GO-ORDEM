package equipamento

import "context"

type Service struct {
	repository Repository
}

func NovoService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Listar(ctx context.Context, consulta FiltroConsulta) ([]Equipamento, error) {
	filtro := Filtro{}

	if consulta.Tipo != "" {
		tipo, err := normalizarTipo(consulta.Tipo)
		if err != nil {
			return nil, err
		}

		filtro.Tipo = &tipo
	}

	if consulta.Categoria != nil {
		categoria := Categoria(*consulta.Categoria)
		if err := validarCategoria(categoria); err != nil {
			return nil, err
		}

		filtro.Categoria = &categoria
	}

	return s.repository.Listar(ctx, filtro)
}

func (s *Service) BuscarPorID(ctx context.Context, id int64) (*EquipamentoDetalhado, error) {
	if err := validarID(id); err != nil {
		return nil, err
	}

	equipamentoEncontrado, err := s.repository.BuscarPorID(ctx, id)
	if err != nil {
		return nil, err
	}

	return s.carregarDetalhes(ctx, equipamentoEncontrado)
}

func (s *Service) BuscarPorCodigo(ctx context.Context, codigo string) (*EquipamentoDetalhado, error) {
	codigoNormalizado, err := normalizaCodigo(codigo)
	if err != nil {
		return nil, err
	}

	equipamentoEncontrado, err := s.repository.BuscarPorCodigo(ctx, codigoNormalizado)
	if err != nil {
		return nil, err
	}

	return s.carregarDetalhes(ctx, equipamentoEncontrado)
}

func (s *Service) ListarModificacoes(ctx context.Context, equipamentoID int64) ([]Modificacao, error) {
	if err := validarID(equipamentoID); err != nil {
		return nil, err
	}

	_, err := s.repository.BuscarPorID(ctx, equipamentoID)
	if err != nil {
		return nil, err
	}

	return s.repository.ListarModificacoes(ctx, equipamentoID)
}

func (s *Service) carregarDetalhes(ctx context.Context, equipamentoEncontrado *Equipamento) (*EquipamentoDetalhado, error) {
	modificacoes, err :=
		s.repository.ListarModificacoes(
			ctx,
			equipamentoEncontrado.ID,
		)
	if err != nil {
		return nil, err
	}

	detalhado := &EquipamentoDetalhado{
		Equipamento:  *equipamentoEncontrado,
		Modificacoes: modificacoes,

		MunicaoCompativeis: make(
			[]MunicaoCompativel,
			0,
		),
	}

	switch equipamentoEncontrado.Tipo {
	case TipoEquipamentoArma:
		arma, err :=
			s.repository.BuscarArmaPorEquipamentoID(
				ctx,
				equipamentoEncontrado.ID,
			)
		if err != nil {
			return nil, err
		}

		municoes, err :=
			s.repository.ListarMunicoesCompativeis(
				ctx,
				equipamentoEncontrado.ID,
			)
		if err != nil {
			return nil, err
		}

		detalhado.Arma = arma
		detalhado.MunicaoCompativeis = municoes

	case TipoEquipamentoProtecao:
		protecao, err :=
			s.repository.BuscarProtecaoPorEquipamentoID(
				ctx,
				equipamentoEncontrado.ID,
			)
		if err != nil {
			return nil, err
		}

		detalhado.Protecao = protecao

	case TipoEquipamentoMunicao:
		municao, err :=
			s.repository.BuscarMunicaoPorEquipamentoID(
				ctx,
				equipamentoEncontrado.ID,
			)
		if err != nil {
			return nil, err
		}

		detalhado.Municao = municao
	}

	return detalhado, nil
}
