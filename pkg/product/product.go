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

// JSON returns the JSON form of Product
func (p Product) JSON() string {
	result, _ := json.Marshal(p)
	return string(result)
}

// CreateProduct creates a product type from a ProductFile
func CreateProduct(file Filer, logFilePath string) (product Product) {
	id, err := file.ID(logFilePath)
	if err != nil {
		id = err.Error()
	}
	message, err := file.ErrorMessage(logFilePath)
	if err != nil {
		message = err.Error()
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
