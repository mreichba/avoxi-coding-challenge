package handlers

import (
	"avoxi-geoip/geoip"
	"encoding/json"
	"log"
	"net/http"
)

// I had countries labeled wrong and AI helped me figure that out due to time limits
// incoming request struct
type VerifyRequest struct {
	IP               string   `json:"ip"`
	AllowedCountries []string `json:"countries"`
}

// outgoing response struct
type VerifyResponse struct {
	Allowed bool   `json:"allowed"`
	Message string `json:"message"`
}

// Ai predictive text was very helpful here and covered an edge case i would not have otherwise placed as well as the map function in the message: which i really liked, i would not have done it this way otherwise.
// Verifies IP address against countries
func VerifyIP(w http.ResponseWriter, r *http.Request) {
	var req VerifyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	country, err := geoip.GetCountryByIP(req.IP)
	if err != nil {
		http.Error(w, "IP lookup failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	allowed := false
	for _, allowedCountry := range req.AllowedCountries {
		if allowedCountry == country {
			allowed = true
			break
		}
	}
	response := VerifyResponse{
		Allowed: allowed,
		Message: "Country " + country + " is " + map[bool]string{true: "allowed", false: "not allowed"}[allowed],
	}

	// log request details
	log.Printf("IP: %v, Selected-Countries: %v, Valid-Country: %v, Allowed: %v", req.IP, req.AllowedCountries, country, allowed)

	//AI was needed for most of these headers as I was running into CORS issues and had some trouble figuring it out
	// sets CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
