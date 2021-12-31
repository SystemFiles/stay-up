package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/systemfiles/stay-up/api/models"
	"github.com/systemfiles/stay-up/api/types"
)

func GetStatus(c echo.Context) error {
	svc := models.Service{
		Name: "Bitwarden",
		Host: "bitwarden.sykesdev.ca",
		Port: 443,
		Protocol: types.PROTO_TCP,
		CurrentStatus: types.UP.String(),
		TimeoutMs: 5096,
		RefreshTimeMs: 200,
	}

		err := svc.CheckStatus()
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, svc)
}