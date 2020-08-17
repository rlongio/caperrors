package product

import (
	"fmt"
	"testing"
)

func TestJSON(t *testing.T) {
	expected := "{\"id\":\"0000\",\"message\":\"Error Message\"}"
	product := NewProduct("0000", "Error Message")
	if result := product.JSON(); result != expected {
		t.Errorf("%v does not equal %v", result, expected)
	}
}

type MockFiler struct {
	id      string
	message string
}

func NewMockFiler(id string, message string) MockFiler {
	return MockFiler{
		id:      id,
		message: message,
	}
}

func (m MockFiler) ID(logFilePath string) (id string, err error) {
	return m.id, nil
}

func (m MockFiler) ErrorMessage(logFilePath string) (id string, err error) {
	return m.message, nil
}

func TestCreateProduct(t *testing.T) {
	mockFile := NewMockFiler("0000", "message")
	product := CreateProduct(mockFile, "")
	if product.ID != "0000" {
		t.Errorf("%v does not equal %v", product.ID, "0000")
	}
	if product.Message != "message" {
		t.Errorf("%v does not equal %v", product.Message, "message")
	}
}

type MockErrorFiler struct{}

func (m MockErrorFiler) ID(logFilePath string) (id string, err error) {
	return "", fmt.Errorf("error id")
}

func (m MockErrorFiler) ErrorMessage(logFilePath string) (id string, err error) {
	return "", fmt.Errorf("error message")
}

func TestCreateProductError(t *testing.T) {
	mockFile := MockErrorFiler{}
	product := CreateProduct(mockFile, "")
	if product.ID != "error id" {
		t.Errorf("%v does not equal %v", product.ID, "error id")
	}
	if product.Message != "error message" {
		t.Errorf("%v does not equal %v", product.Message, "error message")
	}
}
