package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

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
// @Param fromDate query string true "fromDate (YYYY-MM-DD)"
// @Param toDate query string true "toDate (YYYY-MM-DD)"
// @Param type query string false "holiday type"
// @Param Accept header string false "application/json or application/xml"
// @Produce json,xml
// @Router /holidays [get]
func (h *Handler) Get(c *gin.Context) {

	filterFromDate := c.Query("fromDate")
	filterToDate := c.Query("toDate")
	filterByType := c.Query("type")

	d := h.service.Get()

	var data Dto
	err := json.Unmarshal(d, &data)

	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, ErrResponse{Message: err.Error(), Code: http.StatusInternalServerError})
		return
	}

	holiDays := HoliDays{}

	if filterFromDate != "" && filterToDate != "" {
		holiDays, err = FilterByDateRange(data.Data, filterFromDate, filterToDate)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, ErrResponse{Message: err.Error(), Code: http.StatusInternalServerError})
			return
		}

	} else {
		holiDays = data.Data
	}

	if filterByType == "Civil" || filterByType == "Religioso" {
		holiDays = FilterByType(holiDays, filterByType)
	}

	switch c.GetHeader("Accept") {
	case "application/xml":
		c.XML(http.StatusOK, &holiDays)
	default:
		c.JSON(http.StatusOK, &holiDays)
	}

}

func FilterByDateRange(holiDays HoliDays, fromDate, toDate string) (HoliDays, error) {

	const layout = "2006-01-02"

	from, err := time.Parse(layout, fromDate)
	if err != nil {
		return nil, err
	}

	to, err := time.Parse(layout, toDate)
	if err != nil {
		return nil, err
	}

	result := []HoliDay{}

	for _, holiday := range holiDays {

		holidayDate, err := time.Parse(layout, holiday.Date)
		if err != nil {
			continue
		}

		if (holidayDate.Equal(from) || holidayDate.After(from)) &&
			(holidayDate.Equal(to) || holidayDate.Before(to)) {

			result = append(result, holiday)
		}
	}

	return result, nil
}

func FilterByType(holiDays HoliDays, holidayType string) []HoliDay {

	result := []HoliDay{}

	for _, holiday := range holiDays {

		if strings.EqualFold(holiday.Type, holidayType) {
			result = append(result, holiday)
		}
	}

	return result
}
