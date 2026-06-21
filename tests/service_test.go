package tests

import (
	"testing"

	"github.com/go-openapi/testify/assert"
	"github.com/go-openapi/testify/v2/require"
	"github.com/maneul0498-netizen/unicomer_tech_challenge/internal/appication/service"
	dto "github.com/maneul0498-netizen/unicomer_tech_challenge/internal/interfaces/http"
)

func TestLoadCache(t *testing.T) {
	s := service.NewService("https://api.victorsanmartin.com/holidays/en.json")
	err := s.LoadCache()
	require.NoError(t, err)
}

func TestFilterByDateRange(t *testing.T) {
	holiDays := dto.HoliDays{
		dto.HoliDay{
			Date:        "2026-01-01",
			Title:       "Año Nuevo",
			Type:        "Civil",
			Inalienable: true,
			Extra:       "Civil e Irrenunciable",
		},
		dto.HoliDay{
			Date:        "2026-04-03",
			Title:       "Viernes Santo",
			Type:        "Civil",
			Inalienable: false,
			Extra:       "Civil",
		},
		dto.HoliDay{
			Date:        "2026-04-04",
			Title:       "Sábado Santo",
			Type:        "Civil",
			Inalienable: false,
			Extra:       "Civil",
		},
	}

	filteredHoliDays, err := service.FilterByDateRange(holiDays, "2026-01-01", "2026-04-04")
	require.NoError(t, err)
	assert.Equal(t, 3, len(filteredHoliDays))

	filteredHoliDays, err = service.FilterByDateRange(holiDays, "2026-01-01", "2026-04-03")
	require.NoError(t, err)
	assert.Equal(t, 2, len(filteredHoliDays))

	filteredHoliDays, err = service.FilterByDateRange(holiDays, "2026-01-01", "2026-01-01")
	require.NoError(t, err)
	assert.Equal(t, 1, len(filteredHoliDays))

}

func TestFilterByDateRangeWrongDates(t *testing.T) {
	holiDays := dto.HoliDays{
		dto.HoliDay{
			Date:        "2026-01-01",
			Title:       "Año Nuevo",
			Type:        "Civil",
			Inalienable: true,
			Extra:       "Civil e Irrenunciable",
		},
		dto.HoliDay{
			Date:        "2026-04-03",
			Title:       "Viernes Santo",
			Type:        "Civil",
			Inalienable: false,
			Extra:       "Civil",
		},
		dto.HoliDay{
			Date:        "2026-04-04",
			Title:       "Sábado Santo",
			Type:        "Civil",
			Inalienable: false,
			Extra:       "Civil",
		},
	}

	//wrong date
	_, err := service.FilterByDateRange(holiDays, "2026-01-011", "2026-04-04")
	require.Error(t, err)

	//wrong date
	_, err = service.FilterByDateRange(holiDays, "2026-01-01", "2026-04-044")
	require.Error(t, err)

}

func TestFilterByType(t *testing.T) {
	holiDays := dto.HoliDays{
		dto.HoliDay{
			Date:        "2026-01-01",
			Title:       "Año Nuevo",
			Type:        "Civil",
			Inalienable: true,
			Extra:       "Civil e Irrenunciable",
		},
		dto.HoliDay{
			Date:        "2026-04-03",
			Title:       "Viernes Santo",
			Type:        "Civil",
			Inalienable: false,
			Extra:       "Civil",
		},
		dto.HoliDay{
			Date:        "2026-04-04",
			Title:       "Sábado Santo",
			Type:        "Civil",
			Inalienable: false,
			Extra:       "Civil",
		},
		//it is not a real holiday
		dto.HoliDay{
			Date:        "2026-05-05",
			Title:       "Fake holiday",
			Type:        "Religioso",
			Inalienable: false,
			Extra:       "Religioso",
		},
		//it is not a real holiday
		dto.HoliDay{
			Date:        "2026-06-06",
			Title:       "Fake holiday",
			Type:        "Religioso",
			Inalienable: false,
			Extra:       "Religioso",
		},
	}

	filteredHoliDays := service.FilterByType(holiDays, "Civil")
	assert.Equal(t, 3, len(filteredHoliDays))

	filteredHoliDays = service.FilterByType(holiDays, "Religioso")
	assert.Equal(t, 2, len(filteredHoliDays))

}
