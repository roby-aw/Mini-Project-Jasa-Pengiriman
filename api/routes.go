package api

import (
	"api-jasa-pengiriman/api/ongkir"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	OngkirController *ongkir.Controller
}

func RegistrationPath(e *echo.Echo, controller Controller) {
	e.GET("/dataongkir", controller.OngkirController.GetCity)
	e.POST("/cost", controller.OngkirController.GetCost)
}
