package modules

import (
	"api-jasa-pengiriman/api"
	ongkirApi "api-jasa-pengiriman/api/ongkir"
	ongkirBusiness "api-jasa-pengiriman/business/ongkir"
	"api-jasa-pengiriman/config"
	ongkirRepo "api-jasa-pengiriman/repository/ongkir"
	"api-jasa-pengiriman/utils"
)

func RegistrationModules(dbCon *utils.DatabaseConnection, config *config.AppConfig) api.Controller {
	usePermitRepository := ongkirRepo.RepositoryFactory(dbCon)
	usePermitService := ongkirBusiness.NewService(usePermitRepository)
	usePermitController := ongkirApi.NewController(usePermitService)

	controller := api.Controller{
		OngkirController: usePermitController,
	}
	return controller
}
