// Package product contains tools for creating and working on NOAA products.
package product

import (
	"encoding/json"
)

// Product represents an alert product
type Product struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

// ToJSON returns the JSON form of Product
func (p Product) ToJSON() (j []byte, err error) {
	j, err = json.Marshal(p)
	return
}

// CreateProduct creates a product type from a ProductFile
func CreateProduct(file File, logFilePath string) (product Product) {
	id, err := file.ID(logFilePath)
	if err != nil {
		return
	}
	message, err := file.ErrorMessage(logFilePath)
	if err != nil {
		return
	}
	return NewProduct(id, message)
}

// NewProduct returns a new instance of Product
func NewProduct(ID string, message string) Product {
	return Product{
		ID:      ID,
		Message: message,
	}
}
