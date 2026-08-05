package equipamento

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
		Tipo: c.Query("tipo"),
	}

	categoriaTexto := c.Query("categoria")
	if categoriaTexto != "" {
		categoria, err := strconv.Atoi(categoriaTexto)
		if err != nil {
			responderErro(c, ErrCategoriaInvalida)
			return
		}

		consulta.Categoria = &categoria
	}

	equipamento, err := h.service.Listar(c.Request.Context(), consulta)
	if err != nil {
		responderErro(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": NovosEquipamentosResponse(equipamento)})
}

func (h *Handler) BuscarPorID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		responderErro(c, ErrIDInvalido)
		return
	}

	equipamentoEncontrado, err := h.service.BuscarPorID(c.Request.Context(), id)
	if err != nil {
		responderErro(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": NovoEquipamentoDetalhadoResponse(*equipamentoEncontrado)})
}

func (h *Handler) BuscarPorCodigo(c *gin.Context) {
	equipamentoEncontrado, err := h.service.BuscarPorCodigo(c.Request.Context(), c.Param("codigo"))
	if err != nil {
		responderErro(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": NovoEquipamentoDetalhadoResponse(*equipamentoEncontrado)})
}

func (h *Handler) ListarModificacoes(c *gin.Context) {
	equipamentoID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		responderErro(c, ErrIDInvalido)
		return
	}

	modificacoes, err := h.service.ListarModificacoes(c.Request.Context(), equipamentoID)
	if err != nil {
		responderErro(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": NovasModificacoesResponse(modificacoes)})
}

func responderErro(c *gin.Context, err error) {
	switch {
	case errors.Is(err, ErrIDInvalido):
		c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "INVALID_EQUIPMENT_ID", "message": err.Error()}})
	case errors.Is(err, ErrCodigoInvalido):
		c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "INVALID_EQUIPMENT_CODE", "message": err.Error()}})
	case errors.Is(err, ErrTipoInvalido):
		c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "INVALID_EQUIPMENT_TYPE", "message": err.Error()}})
	case errors.Is(err, ErrCategoriaInvalida):
		c.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"code": "INVALID_EQUIPMENT_CATEGORY", "message": err.Error()}})
	case errors.Is(err, ErrEquipamentoNaoEncontrado):
		c.JSON(http.StatusNotFound, gin.H{"error": gin.H{"code": "EQUIPMENT_NOT_FOUND", "message": err.Error()}})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"code": "INTERNAL_SERVER_ERROR", "message": "erro interno do servidor"}})
	}
}
