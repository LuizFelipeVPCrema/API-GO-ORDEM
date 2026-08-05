package equipamento

func CalcularResultadoModificacoes(
	equipamentoBase Equipamento,
	modificacoes []ModificacaoAplicada,
) (ResultadoModificacoes, error) {
	categoriaFinal := int(
		equipamentoBase.CategoriaBase,
	)

	espacosFinais :=
		equipamentoBase.EspacosBase

	for _, aplicacao := range modificacoes {
		if aplicacao.Quantidade <= 0 {
			return ResultadoModificacoes{},
				ErrQuantidadeModificacaoInvalida
		}

		modificacao := aplicacao.Modificacao

		if aplicacao.Quantidade >
			modificacao.LimitePorItem {
			return ResultadoModificacoes{},
				ErrLimiteModificacaoExcedido
		}

		if !modificacao.Cumulativa &&
			aplicacao.Quantidade > 1 {
			return ResultadoModificacoes{},
				ErrModificacaoNaoCumulativa
		}

		categoriaFinal +=
			modificacao.IncrementoCategoria *
				aplicacao.Quantidade

		espacosFinais +=
			modificacao.IncrementoEspacos *
				aplicacao.Quantidade
	}

	if categoriaFinal > int(CategoriaIV) {
		return ResultadoModificacoes{},
			ErrCategoriaFinalInvalida
	}

	if categoriaFinal < int(CategoriaZero) {
		categoriaFinal = int(CategoriaZero)
	}

	if espacosFinais < 0 {
		espacosFinais = 0
	}

	return ResultadoModificacoes{
		CategoriaFinal: Categoria(
			categoriaFinal,
		),
		EspacosFinais: espacosFinais,
	}, nil
}
