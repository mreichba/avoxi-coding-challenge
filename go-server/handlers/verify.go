package handlers

import (
	"avoxi-geoip/geoip"
	models "avoxi-geoip/models/verify"
	"encoding/json"
	"log"
	"net/http"
)

func VerifyIP(w http.ResponseWriter, r *http.Request) {

	//Decode the request body into VerifyRequest struck
	var req models.VerifyRequest
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
	resp := models.VerifyResponse{
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
