package habilidade

import "strings"

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

func normalizarCodigo(codigo string) (Codigo, error) {
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
