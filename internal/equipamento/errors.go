package equipamento

import "errors"

var (
	ErrEquipamentoNaoEncontrado      = errors.New("equipamento nao encontrado")
	ErrIDInvalido                    = errors.New("id do equipamento invalido")
	ErrCodigoInvalido                = errors.New("codigo do equipamento invalido")
	ErrTipoInvalido                  = errors.New("tipo de equipamento invalido")
	ErrCategoriaInvalida             = errors.New("categoria do equipamento invalida")
	ErrQuantidadeModificacaoInvalida = errors.New("quantidade da modificacao invalida")
	ErrLimiteModificacaoExcedido     = errors.New("limite de aplicacoes da modificacoes excedido")
	ErrModificacaoNaoCumulativa      = errors.New("esta modificacao nao pode ser aplicada mais de uma vez")
	ErrCategoriaFinalInvalida        = errors.New("a categoria final do equipamento excede a categoria IV")
)
