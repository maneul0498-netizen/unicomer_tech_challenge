package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func loadInfo() {

}

// GetUser godoc
// @Summary Get
// @Tags Get
// @Param filter path string true "filter"
// @Router / [get]
func (h *Handler) Get(c *gin.Context) {
	log.Println("AAAAAAAAAa")
	switch c.GetHeader("Accept") {
	case "application/xml":
		c.XML(http.StatusOK, nil)
	default:
		c.JSON(http.StatusOK, nil)
	}

}
