package habilidade

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
	habilidades, err := h.service.Listar(c.Request.Context())
	if err != nil {
		responderErro(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": NovasHabilidadesResponse(habilidades)})
}

func (h *Handler) BuscarPorID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		responderErro(c, ErrIDInvalido)
		return
	}

	habilidadeEncontrada, err := h.service.BuscarPorID(c.Request.Context(), id)
	if err != nil {
		responderErro(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": NovaHabilidadeResponse(*habilidadeEncontrada)})
}

func (h *Handler) BuscarPorCodigo(c *gin.Context) {
	habilidadeEncontrada, err := h.service.BuscarPorCodigo(c.Request.Context(), c.Param("codigo"))
	if err != nil {
		responderErro(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": NovaHabilidadeResponse(*habilidadeEncontrada)})
}

func (h *Handler) ListarPorClasseID(c *gin.Context) {
	classeID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		responderErro(c, ErrClasseIDInvalido)
		return
	}

	habilidades, err := h.service.ListarPorClasseID(c.Request.Context(), classeID)
	if err != nil {
		responderErro(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": NovasHabilidadesClasseResponse(habilidades)})
}

func responderErro(c *gin.Context, err error) {
	switch {
	case errors.Is(err, ErrIDInvalido):
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "INVALID_ABILITY_ID",
					"message": err.Error(),
				},
			},
		)

	case errors.Is(err, ErrClasseIDInvalido):
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "INVALID_CLASS_ID",
					"message": err.Error(),
				},
			},
		)

	case errors.Is(err, ErrCodigoInvalido):
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "INVALID_ABILITY_CODE",
					"message": err.Error(),
				},
			},
		)

	case errors.Is(err, ErrHabilidadeNaoEncontrada):
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": gin.H{
					"code":    "ABILITY_NOT_FOUND",
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
					"message": "erro interno do servidor",
				},
			},
		)
	}
}
