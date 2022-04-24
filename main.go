package main

import (
	"api-jasa-pengiriman/api"
	"api-jasa-pengiriman/app/modules"
	"api-jasa-pengiriman/config"
	"api-jasa-pengiriman/utils"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	config := config.GetConfig()

	dbCon := utils.NewConnectionDatabase(config)

	defer dbCon.CloseConnection()

	controllers := modules.RegistrationModules(dbCon, config)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "success")
	})
	api.RegistrationPath(e, controllers)

	go func() {
		address := fmt.Sprintf(":%d", config.App.Port)
		if err := e.Start(address); err != nil {
			log.Fatal(err)
		}
	}()
	quit := make(chan os.Signal)
	<-quit
}
