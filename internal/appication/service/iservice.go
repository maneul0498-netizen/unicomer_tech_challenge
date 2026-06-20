package service

import (
	dto "github.com/maneul0498-netizen/unicomer_tech_challenge/internal/interfaces/http"
)

type IService interface {
	Get(filterFromDate, filterToDate, filterByType string) (dto.HoliDays, error)
	LoadCache() error
}
