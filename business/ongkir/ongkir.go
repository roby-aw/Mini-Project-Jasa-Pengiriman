package ongkir

type Kota struct {
	ID          uint   `gorm:"primaryKey"`
	Kota_Nama   string `json:"kota_nama"`
	Kode_Pos    string `json:"postal_code"`
	Tipe        string `json:"tipe"`
	Provinsi_ID string `json:"provinsi_id"`
}

type Province struct {
	ID            uint   `gorm:"primaryKey"`
	Provinsi_Nama string `json:"provinsi_nama"`
}

type Ongkir struct {
	Asal   string `json:"asal"`
	Tujuan string `json:"tujuan"`
	Berat  string `json:"berat"`
}

type HasilOngkir struct {
	Biaya_Pengiriman string `json:"biaya_pengiriman"`
}

type CostResponse struct {
	Rajaongkir struct {
		Results []Results `json:"results"`
	} `json:"rajaongkir"`
}
type Results struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	Costs []Cost `json:"costs"`
}

type Cost struct {
	Service string `json:"service"`
	Cost    []struct {
		Value int    `json:"value"`
		ETD   string `json:"etd"`
		Note  string `json:"note"`
	} `json:"cost"`
}
