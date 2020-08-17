package internal

import (
	"testing"

	"github.com/SaltyCatFish/caperrors/pkg/product"
)

func Test(t *testing.T) {
	expected := "{\"name\":\"CapHandler Error\",\"status\":\"red\",\"short_summary\":\"Error found\",\"full_summary\":\"Error found in CapHandler log: 1234\",\"debug\":\"{\\\"id\\\":\\\"1234\\\",\\\"message\\\":\\\"Error Message Example\\\"}\",\"remediation\":\"Report to CapHandler team\"}"

	product := product.NewProduct("1234", "Error Message Example")
	output := NewDiagnosticOutput(product)
	if actual := output.JSON(); actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
