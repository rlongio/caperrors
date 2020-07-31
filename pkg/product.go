package pkg

import (
	"encoding/json"
	"fmt"
	"log"
)

// Product represents an alert product
type Product struct {
	ProductID           string `json:"id"`
	ProductErrorMessage string `json:"message"`
}

// ToJSON returns the JSON form of Product
func (p Product) ToJSON() (j []byte, err error) {
	j, err = json.Marshal(p)
	return
}

// CreateProduct creates a product type from a ProductFile
func CreateProduct(file File, logFilePath string) (product Product) {
	log.Printf("processing %v", file.Base())
	productID, err := file.GetProductID(logFilePath)
	if err != nil {
		productID = fmt.Sprintf("ERROR in %v: %v", file.Base(), err)
		log.Printf("Error processing %v :%v", file.Base(), err)
	}
	productErrorInformation, err := file.GetProductErrorInformation(logFilePath)
	if err != nil {
		log.Panic(err)
	}
	product = Product{
		ProductID:           productID,
		ProductErrorMessage: productErrorInformation,
	}
	return
}
