package patente

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
	patentes, err := h.service.Listar(
		c.Request.Context(),
	)
	if err != nil {
		responderErro(c, err)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": NovasPatentesResponse(patentes),
		},
	)
}

func (h *Handler) BuscarPorID(c *gin.Context) {
	id, err := strconv.ParseInt(
		c.Param("id"),
		10,
		64,
	)
	if err != nil {
		responderErro(c, ErrIDInvalido)
		return
	}

	patenteEncontrada, err := h.service.BuscarPorID(
		c.Request.Context(),
		id,
	)
	if err != nil {
		responderErro(c, err)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": NovaPatenteResponse(
				*patenteEncontrada,
			),
		},
	)
}

func (h *Handler) BuscarPorPrestigio(c *gin.Context) {
	pontos, err := strconv.Atoi(
		c.Param("pontos"),
	)
	if err != nil {
		responderErro(c, ErrPrestigioInvalido)
		return
	}

	patenteEncontrada, err :=
		h.service.BuscarPorPrestigio(
			c.Request.Context(),
			pontos,
		)
	if err != nil {
		responderErro(c, err)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"data": NovaPatenteResponse(
				*patenteEncontrada,
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
					"code":    "INVALID_PATENT_ID",
					"message": err.Error(),
				},
			},
		)
	case errors.Is(err, ErrPrestigioInvalido):
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": gin.H{
					"code":    "INVALID_PRESTIGE_POINTS",
					"message": err.Error(),
				},
			},
		)
	case errors.Is(err, ErrPatenteNaoEncontrada):
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": gin.H{
					"code":    "PATENT_NOT_FOUND",
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
