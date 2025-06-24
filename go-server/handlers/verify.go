package handlers

import (
	"avoxi-geoip/geoip"
	"encoding/json"
	"log"
	"net/http"
)

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

func VerifyIP(w http.ResponseWriter, r *http.Request) {

	//Decode the request body into VerifyRequest struck
	var req VerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Failed to decode request: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Get the country from provided IP
	country, err := geoip.GetCountryByIP(req.IP)
	if err != nil {
		log.Printf("IP lookup failed: %v", err)
		http.Error(w, "IP lookup failed", http.StatusInternalServerError)
		return
	}

	// Check if the country is allowed
	allowed := isCountryAllowed(country, req.AllowedCountries)

	// Prepare response
	resp := VerifyResponse{
		Allowed: allowed,
		Message: buildResponseMessage(country, allowed),
	}

	// log the request details
	log.Printf("IP: %s, Selected-Countries: %v, Valid-Country: %s, Allowed: %v", req.IP, req.AllowedCountries, country, allowed)

	// Set CORS headers
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Failed to encode response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func isCountryAllowed(country string, allowedCountries []string) bool {
	for _, allowedCountry := range allowedCountries {
		if allowedCountry == country {
			return true
		}
	}
	return false
}

func buildResponseMessage(country string, allowed bool) string {
	if allowed {
		return "Country " + country + " is allowed"
	} else {
		return "Country " + country + " is not allowed"
	}
}
