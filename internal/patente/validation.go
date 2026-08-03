package patente

func validarID(id int64) error {
	if id <= 0 {
		return ErrIDInvalido
	}

	return nil
}

func validarPontosPrestigio(pontos int) error {
	if pontos < 0 {
		return ErrPrestigioInvalido
	}

	return nil
}
