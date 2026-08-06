package personagem

import "errors"

var (
	ErrPersonagemNaoEncontrado    = errors.New("personagem nao encontrado")
	ErrClasseNaoEncontrada        = errors.New("classe nao encontrada")
	ErrIDInvalido                 = errors.New("id do personagem invalido")
	ErrClasseIDInvalido           = errors.New("id da classe invalido")
	ErrNomeObrigatorio            = errors.New("nome do personagem e obrigatorio")
	ErrNEXInvalido                = errors.New("NEX do personagem invalido")
	ErrPrestigioInvalido          = errors.New("pontos de prestigio invalido")
	ErrIdadeInvalida              = errors.New("idade do personagem invalida")
	ErrAtributoInvalido           = errors.New("valor de atributo invalido")
	ErrRecursoInvalido            = errors.New("valor de recurso invalido")
	ErrRecursoAtualMaiorQueMaximo = errors.New("recurso atual nao pode ser maior que o maximo")
	ErrRequisicaoInvalida         = errors.New("requisicao invalida")
)
