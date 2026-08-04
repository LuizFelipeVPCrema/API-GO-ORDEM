package classe

import "strings"

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
		ehLetraMaiuscula := caractere >= 'A' && caractere <= 'Z'

		ehSeparador := caractere == '_'

		if !ehLetraMaiuscula && !ehSeparador {
			return "", ErrCodigoInvalido
		}
	}

	return Codigo(codigoNormalizado), nil
}
