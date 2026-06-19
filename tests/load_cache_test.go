package tests

import (
	"testing"

	"github.com/go-openapi/testify/v2/require"
	"github.com/maneul0498-netizen/unicomer_tech_challenge/internal/appication/service"
)

func TestLoadCache(t *testing.T) {
	s := service.NewService()
	_, err := s.LoadCache()
	require.NoError(t, err)

}
