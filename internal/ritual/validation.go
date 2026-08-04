package ritual

import "strings"

func validarID(id int64) error {
	if id <= 0 {
		return ErrIDInvalido
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

func normalizarElemento(elemento string) (Elemento, error) {
	elementoNormalizado := Elemento(strings.ToUpper(strings.TrimSpace(elemento)))

	switch elementoNormalizado {
	case ElementoSangue,
		ElementoMorte,
		ElementoConhecimento,
		ElementoEnergia,
		ElementoMedo:
		return elementoNormalizado, nil
	default:
		return "", ErrElementoInvalido
	}
}

func validarCirculo(circulo Circulo) error {
	switch circulo {
	case PrimeiroCirculo,
		SegundoCirculo,
		TerceiroCirculo,
		QuartoCirculo,
		QuintoCirculo:
		return nil

	default:
		return ErrCirculoInvalido
	}
}
