package models

type Fetch struct {
	UUID			string	`json:"UUID"`
	AreaProvinsi	string	`json:"area_provinsi"`
	Komoditas		string	`json:"komoditas"`
	AreaKota		string 	`json:"area_kota"`
	Size			string 	`json:"size"`
	Price			string 	`json:"price"`
	TglParsed		string 	`json:"tgl_parsed"`
	Timestamp		string 	`json:"timestamp"`
}

type USDValue struct {
	Value			float32	`json:"IDR_USD"`
}

type FetchCollection struct {
	Fetchs []Fetch `json:"items"`
}

