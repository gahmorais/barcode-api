package models

type Product struct {
	Id         string
	Name       string
	Barcode    string
	Categories string
}

type ProductResponse struct {
	Id          string `bson:"id"`
	Description string `bson:"name"`
	Barcode     string `bson:"barcode"`
	Categories  string `bson:"categories"`
}
