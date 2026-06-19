package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maneul0498-netizen/unicomer_tech_challenge/internal/appication/service"
)

type Handler struct {
	service service.IService
}

func NewHandler(s service.IService) *Handler {
	return &Handler{
		service: s,
	}
}

// GetUser godoc
// @Summary Get
// @Tags HolyDays
// @Param filter path string true "filter"
// @Param Accept header string false "application/json or application/xml"
// Produce json,xml
// @Router / [get]
func (h *Handler) Get(c *gin.Context) {

	h.service.Get()

	switch c.GetHeader("Accept") {
	case "application/xml":
		c.XML(http.StatusOK, nil)
	default:
		c.JSON(http.StatusOK, nil)
	}

}
