package ongkir

import (
	"api-jasa-pengiriman/business/ongkir"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

type MysqlRepository struct {
	db *gorm.DB
}

func NewMysqlRepository(db *gorm.DB) *MysqlRepository {
	return &MysqlRepository{
		db: db,
	}
}

func (repo *MysqlRepository) City(DataOngkir *ongkir.Ongkir) (*ongkir.Kota, error) {
	var kota ongkir.Kota
	err := repo.db.Where("kota_nama = ?", DataOngkir.Asal).First(&kota).Error
	if err != nil {
		return nil, err
	}
	return &kota, err
}

func (repo *MysqlRepository) Cost(Dataongkir *ongkir.Ongkir) (*ongkir.CostResponse, error) {
	var kotaasal ongkir.Kota
	var kotatujuan ongkir.Kota
	err := repo.db.Where("kota_nama = ?", Dataongkir.Asal).First(&kotaasal).Error
	if err != nil {
		return nil, err
	}
	err = repo.db.Where("kota_nama = ?", Dataongkir.Tujuan).First(&kotatujuan).Error
	if err != nil {
		return nil, err
	}
	url := "https://api.rajaongkir.com/starter/cost"
	var queryString string = fmt.Sprintf("origin=%d&destination=%d&weight=%s&courier=jne", kotaasal.ID, kotatujuan.ID, Dataongkir.Berat)
	payload := strings.NewReader(queryString)
	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("key", "75ac21477faeddb1b34f8dd63fc61f45")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	hasil := string(body)
	fmt.Println(hasil)
	var cost ongkir.CostResponse
	err = json.Unmarshal([]byte(hasil), &cost)
	if err != nil {
		return nil, err
	}
	return &cost, nil
}
