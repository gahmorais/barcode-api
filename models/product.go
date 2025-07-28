package models

type Product struct {
	Id         string
	Name       string
	Barcode    string
	Categories string
}

type ProductResponse struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Barcode    string `json:"barcode"`
	Categories string `json:"categories"`
}
