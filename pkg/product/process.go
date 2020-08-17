package product

import (
	"encoding/json"
	"fmt"
)

// Process is a function that operates on a product
type Process func(Product)

// Print prints the product json to stdout
func (p Product) Print() {
	fmt.Println(string(p.JSON()))
}

// DiagnosticOutput contains the information needed to send
// a diagnostic update to Big Brother
type DiagnosticOutput struct {
	Name         string `json:"name"`
	Status       string `json:"status"`
	ShortSummary string `json:"short_summary"`
	FullSummary  string `json:"full_summary"`
	Debug        string `json:"debug"`
	Remediation  string `json:"remediation"`
}

// PrintDiagnosticOutput prints product for sending to big brother
func (p Product) PrintDiagnosticOutput() {
	d := DiagnosticOutput{
		Name:         "CapHandler Error",
		Status:       "red",
		ShortSummary: "Error found",
		FullSummary:  fmt.Sprintf("Error found in CapHandler log: %v", p.ID),
		Debug:        fmt.Sprintf("%v", p.JSON()),
		Remediation:  "Report to CapHandler team",
	}
	result, _ := json.Marshal(d)
	fmt.Println(string(result))
}
