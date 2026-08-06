package personagem

import "strings"

const (
	limitePadrao = 20
	limiteMaximo = 100
)

func validarID(id int64) error {
	if id <= 0 {
		return ErrIDInvalido
	}

	return nil
}

func validarClasseID(id int64) error {
	if id <= 0 {
		return ErrClasseIDInvalido
	}

	return nil
}

func validarPersonagem(personagem Personagem) error {
	personagem.Nome = strings.TrimSpace(personagem.Nome)

	if personagem.Nome == "" {
		return ErrNomeObrigatorio
	}

	if err := validarClasseID(personagem.ClasseID); err != nil {
		return err
	}

	if personagem.NEX < 0 || personagem.NEX > 100 {
		return ErrNEXInvalido
	}

	if personagem.PontosPrestigio < 0 {
		return ErrPrestigioInvalido
	}

	if personagem.Idade != nil && *personagem.Idade <= 0 {
		return ErrIdadeInvalida
	}

	if err := validarAtributos(personagem.Atributos); err != nil {
		return err
	}

	return validarRecursos(personagem.Recursos)
}

func validarAtributos(atributos Atributos) error {
	if atributos.Agilidade < 0 ||
		atributos.Forca < 0 ||
		atributos.Intelecto < 0 ||
		atributos.Presenca < 0 ||
		atributos.Vigor < 0 {
		return ErrAtributoInvalido
	}

	return nil
}

func validarRecursos(recursos Recursos) error {
	if recursos.PEAtual < 0 ||
		recursos.PVMaximo < 0 ||
		recursos.PEAtual < 0 ||
		recursos.PEMaximo < 0 ||
		recursos.SanidadeAtual < 0 ||
		recursos.SanidadeMaxima < 0 {
		return ErrRecursoInvalido
	}

	if recursos.PVAtual > recursos.PVMaximo ||
		recursos.PEAtual > recursos.PEMaximo ||
		recursos.SanidadeAtual > recursos.SanidadeMaxima {
		return ErrRecursoAtualMaiorQueMaximo
	}

	return nil
}

func normalizarTextoOpcional(texto *string) *string {
	if texto == nil {
		return nil
	}

	valor := strings.TrimSpace(*texto)

	if valor == "" {
		return nil
	}

	return &valor
}

func normalizarFiltro(filtro Filtro) Filtro {
	filtro.Nome = strings.TrimSpace(
		filtro.Nome,
	)

	if filtro.Limit <= 0 {
		filtro.Limit = limitePadrao
	}

	if filtro.Limit > limiteMaximo {
		filtro.Limit = limiteMaximo
	}

	if filtro.Offset < 0 {
		filtro.Offset = 0
	}

	return filtro
}

func validarAtualizacao(request AtualizarPersonagemRequest) error {
	temAlteracao :=
		request.Nome != nil ||
			request.Jogador != nil ||
			request.ClasseID != nil ||
			request.NEX != nil ||
			request.PontosPrestigio != nil ||
			request.Idade != nil ||
			request.Aparencia != nil ||
			request.Personalidade != nil ||
			request.Historia != nil ||
			request.Objetivo != nil ||
			temAtualizacaoAtributos(
				request.Atributos,
			) ||
			temAtualizacaoRecursos(
				request.Recursos,
			)

	if !temAlteracao {
		return ErrRequisicaoInvalida
	}

	return nil
}

func temAtualizacaoAtributos(request *AtualizarAtributosRequest) bool {
	if request == nil {
		return false
	}

	return request.Agilidade != nil ||
		request.Forca != nil ||
		request.Intelecto != nil ||
		request.Presenca != nil ||
		request.Vigor != nil
}

func temAtualizacaoRecursos(request *AtualizarRecursosRequest) bool {
	if request == nil {
		return false
	}

	return request.PVAtual != nil ||
		request.PVMaximo != nil ||
		request.PEAtual != nil ||
		request.PEMaximo != nil ||
		request.SanidadeAtual != nil ||
		request.SanidadeMaxima != nil
}
