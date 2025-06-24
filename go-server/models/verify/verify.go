package models

// VerifyRequest represents the incoming IP and allowed country list.
type VerifyRequest struct {
	IP               string   `json:"ip"`
	AllowedCountries []string `json:"countries"`
}

// VerifyResponse represents the API's response about IP allowance.
type VerifyResponse struct {
	Allowed bool   `json:"allowed"`
	Message string `json:"message"`
}
