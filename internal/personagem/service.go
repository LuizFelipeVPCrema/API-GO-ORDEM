package personagem

import (
	"context"
	"errors"
	"strings"

	"github.com/LuizFelipeVPCrema/api-go-ordem/internal/classe"
)

type ClasseProvider interface {
	BuscarPorID(ctx context.Context, id int64) (*classe.Classe, error)
}

type Service struct {
	repository     Repository
	classeProvider ClasseProvider
}

func NovoService(repository Repository, classeProvider ClasseProvider) *Service {
	return &Service{
		repository:     repository,
		classeProvider: classeProvider,
	}
}

func (s *Service) Listar(ctx context.Context, filtro Filtro) ([]Personagem, error) {
	filtro = normalizarFiltro(filtro)

	if filtro.ClasseID != nil {
		if err := validarClasseID(*filtro.ClasseID); err != nil {
			return nil, err
		}
	}

	return s.repository.Listar(ctx, filtro)
}

func (s *Service) BuscarPorID(ctx context.Context, id int64) (*Personagem, error) {
	if err := validarID(id); err != nil {
		return nil, err
	}

	return s.repository.BuscarPorID(ctx, id)
}

func (s *Service) Criar(ctx context.Context, request CriarPersonagemRequest) (*Personagem, error) {
	personagemNovo := Personagem{
		Nome: strings.TrimSpace(
			request.Nome,
		),

		Jogador: normalizarTextoOpcional(
			request.Jogador,
		),

		ClasseID: request.ClasseID,

		NEX:             request.NEX,
		PontosPrestigio: request.PontosPrestigio,

		Idade: request.Idade,

		Aparencia: normalizarTextoOpcional(
			request.Aparencia,
		),

		Personalidade: normalizarTextoOpcional(
			request.Personalidade,
		),

		Historia: normalizarTextoOpcional(
			request.Historia,
		),

		Objetivo: normalizarTextoOpcional(
			request.Objetivo,
		),

		Atributos: Atributos{
			Agilidade: request.Atributos.Agilidade,
			Forca:     request.Atributos.Forca,
			Intelecto: request.Atributos.Intelecto,
			Presenca:  request.Atributos.Presenca,
			Vigor:     request.Atributos.Vigor,
		},

		Recursos: Recursos{
			PVAtual:  request.Recursos.PVAtual,
			PVMaximo: request.Recursos.PVMaximo,

			PEAtual:  request.Recursos.PEAtual,
			PEMaximo: request.Recursos.PEMaximo,

			SanidadeAtual: request.Recursos.SanidadeAtual,

			SanidadeMaxima: request.Recursos.SanidadeMaxima,
		},

		Ativa: true,
	}

	if err := validarPersonagem(personagemNovo); err != nil {
		return nil, err
	}

	if err := s.validarClasse(ctx, personagemNovo.ClasseID); err != nil {
		return nil, err
	}

	return s.repository.Criar(ctx, personagemNovo)
}

func (s *Service) Atualizar(ctx context.Context, id int64, request AtualizarPersonagemRequest) (*Personagem, error) {
	if err := validarID(id); err != nil {
		return nil, err
	}

	if err := validarAtualizacao(request); err != nil {
		return nil, err
	}

	personagemAtual, err := s.repository.BuscarPorID(ctx, id)
	if err != nil {
		return nil, err
	}

	aplicarAtualizacaoGeral(personagemAtual, request)

	aplicarAtualizacaoAtributos(&personagemAtual.Atributos, request.Atributos)

	aplicarAtualizacaoRecursos(&personagemAtual.Recursos, request.Recursos)

	if err := validarPersonagem(*personagemAtual); err != nil {
		return nil, err
	}

	if request.ClasseID != nil {
		if err := s.validarClasse(ctx, personagemAtual.ClasseID); err != nil {
			return nil, err
		}
	}

	if err := s.repository.Atualizar(ctx, *personagemAtual); err != nil {
		return nil, err
	}

	return personagemAtual, nil
}

func (s *Service) Remover(ctx context.Context, id int64) error {
	if err := validarID(id); err != nil {
		return err
	}

	return s.repository.Desativar(ctx, id)
}

func (s *Service) validarClasse(ctx context.Context, classeID int64) error {
	_, err := s.classeProvider.BuscarPorID(
		ctx,
		classeID,
	)

	if errors.Is(
		err,
		classe.ErrClasseNaoEncontrada,
	) {
		return ErrClasseNaoEncontrada
	}

	return err
}

func aplicarAtualizacaoGeral(
	personagemAtual *Personagem,
	request AtualizarPersonagemRequest,
) {
	if request.Nome != nil {
		personagemAtual.Nome =
			strings.TrimSpace(*request.Nome)
	}

	if request.Jogador != nil {
		personagemAtual.Jogador =
			normalizarTextoOpcional(
				request.Jogador,
			)
	}

	if request.ClasseID != nil {
		personagemAtual.ClasseID =
			*request.ClasseID
	}

	if request.NEX != nil {
		personagemAtual.NEX =
			*request.NEX
	}

	if request.PontosPrestigio != nil {
		personagemAtual.PontosPrestigio =
			*request.PontosPrestigio
	}

	if request.Idade != nil {
		personagemAtual.Idade =
			request.Idade
	}

	if request.Aparencia != nil {
		personagemAtual.Aparencia =
			normalizarTextoOpcional(
				request.Aparencia,
			)
	}

	if request.Personalidade != nil {
		personagemAtual.Personalidade =
			normalizarTextoOpcional(
				request.Personalidade,
			)
	}

	if request.Historia != nil {
		personagemAtual.Historia =
			normalizarTextoOpcional(
				request.Historia,
			)
	}

	if request.Objetivo != nil {
		personagemAtual.Objetivo =
			normalizarTextoOpcional(
				request.Objetivo,
			)
	}
}

func aplicarAtualizacaoAtributos(
	atributos *Atributos,
	request *AtualizarAtributosRequest,
) {
	if request == nil {
		return
	}

	if request.Agilidade != nil {
		atributos.Agilidade =
			*request.Agilidade
	}

	if request.Forca != nil {
		atributos.Forca =
			*request.Forca
	}

	if request.Intelecto != nil {
		atributos.Intelecto =
			*request.Intelecto
	}

	if request.Presenca != nil {
		atributos.Presenca =
			*request.Presenca
	}

	if request.Vigor != nil {
		atributos.Vigor =
			*request.Vigor
	}
}

func aplicarAtualizacaoRecursos(
	recursos *Recursos,
	request *AtualizarRecursosRequest,
) {
	if request == nil {
		return
	}

	if request.PVAtual != nil {
		recursos.PVAtual =
			*request.PVAtual
	}

	if request.PVMaximo != nil {
		recursos.PVMaximo =
			*request.PVMaximo
	}

	if request.PEAtual != nil {
		recursos.PEAtual =
			*request.PEAtual
	}

	if request.PEMaximo != nil {
		recursos.PEMaximo =
			*request.PEMaximo
	}

	if request.SanidadeAtual != nil {
		recursos.SanidadeAtual =
			*request.SanidadeAtual
	}

	if request.SanidadeMaxima != nil {
		recursos.SanidadeMaxima =
			*request.SanidadeMaxima
	}
}
