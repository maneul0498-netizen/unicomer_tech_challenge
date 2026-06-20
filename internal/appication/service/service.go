package service

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	dto "github.com/maneul0498-netizen/unicomer_tech_challenge/internal/interfaces/http"
)

type Service struct {
	data       dto.Dto
	boostr_url string
}

func NewService(boostr_ur string) IService {
	s := &Service{
		boostr_url: boostr_ur,
	}
	err := s.LoadCache()
	if err != nil {
		log.Println("Error trying to get data")
		panic(err)
	}
	return s
}

func (s *Service) LoadCache() error {
	log.Println("getting cache data")

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, s.boostr_url, nil)

	if err != nil {
		log.Println(err)
		return err
	}

	res, err := client.Do(req)

	if err != nil {
		log.Println(err)
		return err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Println(err)
		return err
	}

	err = json.Unmarshal(body, &s.data)

	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Get(filterFromDate, filterToDate, filterByType string) (dto.HoliDays, error) {
	holiDays := s.data.Data
	var err error

	if filterFromDate != "" && filterToDate != "" {
		holiDays, err = FilterByDateRange(holiDays, filterFromDate, filterToDate)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}

	if filterByType != "" {
		holiDays = FilterByType(holiDays, filterByType)
	}
	return holiDays, nil
}

func FilterByDateRange(holiDays dto.HoliDays, fromDate, toDate string) (dto.HoliDays, error) {

	const layout = "2006-01-02"

	from, err := time.Parse(layout, fromDate)
	if err != nil {
		return nil, err
	}

	to, err := time.Parse(layout, toDate)
	if err != nil {
		return nil, err
	}

	result := dto.HoliDays{}

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

func FilterByType(holiDays dto.HoliDays, holidayType string) dto.HoliDays {

	result := dto.HoliDays{}

	for _, holiday := range holiDays {

		if strings.EqualFold(holiday.Type, holidayType) {
			result = append(result, holiday)
		}
	}

	return result
}
