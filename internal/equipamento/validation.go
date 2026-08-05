package equipamento

import "strings"

func validarID(id int64) error {
	if id <= 0 {
		return ErrIDInvalido
	}

	return nil
}

func validarCategoria(categoria Categoria) error {
	if categoria < CategoriaZero || categoria > CategoriaIV {
		return ErrCategoriaInvalida
	}

	return nil
}

func normalizaCodigo(codigo string) (Codigo, error) {
	codigoNormalizado := strings.ToUpper(strings.TrimSpace(codigo))
	if codigoNormalizado == "" {
		return "", ErrCodigoInvalido
	}

	for _, caractere := range codigoNormalizado {
		ehLetra := caractere >= 'A' && caractere <= 'Z'

		ehNumero := caractere >= '0' && caractere <= '9'

		ehSeparador := caractere == '_'

		if !ehLetra && !ehNumero && !ehSeparador {
			return "", ErrCodigoInvalido
		}
	}

	return Codigo(codigoNormalizado), nil
}

func normalizarTipo(tipo string) (Tipo, error) {
	tipoNormalizado := Tipo(strings.ToUpper(strings.TrimSpace(tipo)))

	switch tipoNormalizado {
	case TipoEquipamentoArma,
		TipoEquipamentoProtecao,
		TipoEquipamentoAcessorio,
		TipoEquipamentoMunicao,
		TipoEquipamentoItemOperacional,
		TipoEquipamentoItemParanormal,
		TipoEquipamentoOutro:
		return tipoNormalizado, nil
	default:
		return "", ErrTipoInvalido
	}
}
