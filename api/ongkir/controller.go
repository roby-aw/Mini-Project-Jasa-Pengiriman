package ongkir

import (
	"api-jasa-pengiriman/business/ongkir"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	service ongkir.Service
}

func NewController(service ongkir.Service) *Controller {
	return &Controller{
		service: service,
	}
}

func (Controller *Controller) GetCity(c echo.Context) error {
	ongkir := ongkir.Ongkir{}
	c.Bind(&ongkir)
	ongkirs, err := Controller.service.GetCity(&ongkir)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, ongkirs.ID)
}

func (Controller *Controller) GetCost(c echo.Context) error {
	ongkir := ongkir.Ongkir{}
	c.Bind(&ongkir)
	ongkirs, err := Controller.service.GetCost(&ongkir)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, ongkirs.Rajaongkir.Results)
}
