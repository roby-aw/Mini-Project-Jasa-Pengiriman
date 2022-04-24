package ongkir

import (
	"api-jasa-pengiriman/business/ongkir"
	"api-jasa-pengiriman/utils"
)

func RepositoryFactory(dbCon *utils.DatabaseConnection) ongkir.Repository {
	ongkirRepo := NewMysqlRepository(dbCon.Mysql)
	return ongkirRepo
}
