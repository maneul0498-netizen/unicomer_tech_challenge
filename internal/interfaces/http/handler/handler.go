package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/maneul0498-netizen/unicomer_tech_challenge/internal/appication/service"
	dto "github.com/maneul0498-netizen/unicomer_tech_challenge/internal/interfaces/http"
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
// @Param fromDate query string false "fromDate (YYYY-MM-DD)"
// @Param toDate query string false "toDate (YYYY-MM-DD)"
// @Param type query string false "holiday type"
// @Param Accept header string false "application/json or application/xml"
// @Produce json
// @Produce xml
// @Failure 500 {object} http.ErrorResponse
// @Router /holidays [get]
func (h *Handler) Get(c *gin.Context) {

	filterFromDate := c.Query("fromDate")
	filterToDate := c.Query("toDate")
	filterByType := c.Query("type")

	holiDays, err := h.service.Get(filterFromDate, filterToDate, filterByType)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, dto.ErrorResponse{Message: err.Error(), Code: http.StatusInternalServerError})
		return
	}

	switch c.GetHeader("Accept") {
	case "application/xml":
		c.XML(http.StatusOK, &holiDays)
	default:
		c.JSON(http.StatusOK, &holiDays)
	}

}
