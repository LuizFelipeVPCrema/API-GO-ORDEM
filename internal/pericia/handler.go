package pericia

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
	pericias, err := h.service.Listar(c.Request.Context())
	if err != nil {
		responderErro(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": NovasPericiasResponse(pericias)})
}

func (h *Handler) BuscarPorID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		responderErro(c, ErrIDInvalido)
		return
	}

	periciaEncontrada, err := h.service.BuscarPorID(c.Request.Context(), id)
	if err != nil {
		responderErro(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": NovaPericiaResponse(*periciaEncontrada)})
}

func (h *Handler) BuscarPorCodigo(c *gin.Context) {
	codigo := c.Param("codigo")

	periciaEncontrada, err := h.service.BuscarPorCodigo(c.Request.Context(), codigo)
	if err != nil {
		responderErro(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": NovaPericiaResponse(*periciaEncontrada)})
}

func responderErro(c *gin.Context, err error) {
	switch {
	case errors.Is(err, ErrIDInvalido):
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "INVALID_SKILL_ID",
					"message": err.Error(),
				},
			},
		)

	case errors.Is(err, ErrCodigoInvalido):
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "INVALID_SKILL_CODE",
					"message": err.Error(),
				},
			},
		)

	case errors.Is(err, ErrPericiaNaoEncontrada):
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": gin.H{
					"code":    "SKILL_NOT_FOUND",
					"message": err.Error(),
				},
			},
		)

	default:
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
