package internal

import (
	"encoding/json"
	"fmt"

	"github.com/SaltyCatFish/caperrors/pkg/product"
)

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

// NewDiagnosticOutput returns a new instance of DiagnosticOuput
func NewDiagnosticOutput(p product.Product) DiagnosticOutput {
	return DiagnosticOutput{
		Name:         "CapHandler Error",
		Status:       "red",
		ShortSummary: "Error found",
		FullSummary:  fmt.Sprintf("Error found in CapHandler log: %v", p.ID),
		Debug:        fmt.Sprintf("%v", p.JSON()),
		Remediation:  "Report to CapHandler team",
	}
}

// JSON converts DiagnosticOuput to JSON
func (d DiagnosticOutput) JSON() string {
	result, _ := json.Marshal(d)
	return string(result)
}
