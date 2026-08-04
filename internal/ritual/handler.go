package ritual

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
	consulta := FiltroConsulta{
		Elemento: c.Query("elemento"),
	}

	circuloTexto := c.Query("circulo")
	if circuloTexto != "" {
		circulo, err := strconv.Atoi(circuloTexto)
		if err != nil {
			responderErro(c, ErrCirculoInvalido)
			return
		}

		consulta.Circulo = &circulo
	}

	rituais, err := h.service.Listar(c.Request.Context(), consulta)
	if err != nil {
		responderErro(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": NovosRituaisResponse(rituais)})
}

func (h *Handler) BuscarPorID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		responderErro(c, ErrIDInvalido)
		return
	}

	ritualEncontrado, err := h.service.BuscarPorID(c.Request.Context(), id)
	if err != nil {
		responderErro(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": NovoRitualDetalhadoResponse(*ritualEncontrado)})
}

func (h *Handler) BuscarPorCodigo(c *gin.Context) {
	ritualEncontrado, err := h.service.BuscarPorCodigo(c.Request.Context(), c.Param("codigo"))
	if err != nil {
		responderErro(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": NovoRitualDetalhadoResponse(*ritualEncontrado)})
}

func (h *Handler) ListarAprimoramentos(c *gin.Context) {
	ritualID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		responderErro(c, ErrIDInvalido)
		return
	}

	aprimoramentos, err := h.service.ListarAprimoramentos(c.Request.Context(), ritualID)
	if err != nil {
		responderErro(c, err)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": NovosAprimoramentosResponse(
				aprimoramentos,
			),
		},
	)
}

func responderErro(c *gin.Context, err error) {
	switch {
	case errors.Is(err, ErrIDInvalido):
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "INVALID_RITUAL_ID",
					"message": err.Error(),
				},
			},
		)

	case errors.Is(err, ErrCodigoInvalido):
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "INVALID_RITUAL_CODE",
					"message": err.Error(),
				},
			},
		)

	case errors.Is(err, ErrElementoInvalido):
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "INVALID_RITUAL_ELEMENT",
					"message": err.Error(),
				},
			},
		)

	case errors.Is(err, ErrCirculoInvalido):
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "INVALID_RITUAL_CIRCLE",
					"message": err.Error(),
				},
			},
		)

	case errors.Is(err, ErrRitualNaoEncontrado):
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": gin.H{
					"code":    "RITUAL_NOT_FOUND",
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
