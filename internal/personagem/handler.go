package personagem

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NovoHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Listar(c *gin.Context) {
	filtro := Filtro{
		Nome: c.Query("nome"),
	}

	if classeTexto := c.Query("classe_id"); classeTexto != "" {
		classeID, err := strconv.ParseInt(classeTexto, 10, 64)
		if err != nil {
			responderErro(c, ErrClasseIDInvalido)
			return
		}

		filtro.ClasseID = &classeID
	}

	limit, err := lerInteiroConsulta(c.Query("limit"), 20)
	if err != nil {
		responderErro(c, ErrRequisicaoInvalida)
		return
	}

	offset, err := lerInteiroConsulta(c.Query("offset"), 0)
	if err != nil {
		responderErro(c, ErrRequisicaoInvalida)
		return
	}

	filtro.Limit = limit
	filtro.Offset = offset

	personagens, err := h.service.Listar(c.Request.Context(), filtro)
	if err != nil {
		responderErro(c, err)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": NovosPersonagensResponse(personagens),
			"pagination": gin.H{
				"limit":  filtro.Limit,
				"offset": filtro.Offset,
				"count":  len(personagens),
			},
		},
	)
}

func (h *Handler) BuscarPorID(c *gin.Context) {
	id, err := lerID(c.Param("id"))
	if err != nil {
		responderErro(c, err)
		return
	}

	personagemEncontrada, err := h.service.BuscarPorID(c.Request.Context(), id)
	if err != nil {
		responderErro(c, err)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": NovoPersonagemResponse(*personagemEncontrada),
		})
}

func (h *Handler) Criar(c *gin.Context) {
	var request CriarPersonagemRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		responderErro(c, ErrRequisicaoInvalida)
		return
	}

	personagemCriado, err := h.service.Criar(c.Request.Context(), request)
	if err != nil {
		responderErro(c, err)
		return
	}

	c.JSON(
		http.StatusCreated,
		gin.H{
			"data": NovoPersonagemResponse(*personagemCriado),
		},
	)
}

func (h *Handler) Atualizar(c *gin.Context) {
	id, err := lerID(c.Param("id"))
	if err != nil {
		responderErro(c, err)
		return
	}

	var request AtualizarPersonagemRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		responderErro(c, ErrRequisicaoInvalida)
		return
	}

	personagemAtualizada, err := h.service.Atualizar(c.Request.Context(), id, request)
	if err != nil {
		responderErro(c, err)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": NovoPersonagemResponse(*personagemAtualizada),
		},
	)
}

func (h *Handler) Remover(c *gin.Context) {
	id, err := lerID(c.Param("id"))
	if err != nil {
		responderErro(c, err)
		return
	}

	if err := h.service.Remover(c.Request.Context(), id); err != nil {
		responderErro(c, err)
		return
	}
	c.Status(http.StatusNoContent)
}

func lerID(valor string) (int64, error) {
	id, err := strconv.ParseInt(valor, 10, 64)
	if err != nil || id <= 0 {
		return 0, ErrIDInvalido
	}

	return id, nil
}

func lerInteiroConsulta(valor string, padrao int) (int, error) {
	if valor == "" {
		return padrao, nil
	}

	numero, err := strconv.Atoi(valor)
	if err != nil || numero < 0 {
		return 0, ErrRequisicaoInvalida
	}

	return numero, nil
}

func responderErro(c *gin.Context, err error) {
	switch {
	case errors.Is(err, ErrIDInvalido),
		errors.Is(err, ErrClasseIDInvalido),
		errors.Is(err, ErrNomeObrigatorio),
		errors.Is(err, ErrNEXInvalido),
		errors.Is(err, ErrPrestigioInvalido),
		errors.Is(err, ErrIdadeInvalida),
		errors.Is(err, ErrAtributoInvalido),
		errors.Is(err, ErrRecursoInvalido),
		errors.Is(err, ErrRecursoAtualMaiorQueMaximo),
		errors.Is(err, ErrRequisicaoInvalida):
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "INVALID_CHARACTER_REQUEST",
					"message": err.Error(),
				},
			},
		)
	case errors.Is(err, ErrPersonagemNaoEncontrado):
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": gin.H{
					"code":    "CARACTER_NOT_FOUND",
					"message": err.Error(),
				},
			},
		)

	case errors.Is(err, ErrClasseNaoEncontrada):
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": gin.H{
					"code":    "CLASS_NOT_FOUND",
					"message": err.Error(),
				},
			},
		)
	default:
		_ = c.Error(err)

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": gin.H{
					"code":    "INTERNAL_ERROR",
					"message": "erro interna do servidor",
				},
			},
		)
	}
}
