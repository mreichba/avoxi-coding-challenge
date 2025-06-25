package handlers

import (
	"avoxi-geoip/geoip"
	models "avoxi-geoip/models/verify"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIsCountryAllowed(t *testing.T) {
	tests := []struct {
		name             string
		country          string
		allowedCountries []string
		expected         bool
	}{
		{
			name:             "Allowed country present",
			country:          "US",
			allowedCountries: []string{"US", "CA", "GB"},
			expected:         true,
		},
		{
			name:             "Allowed country NOT present",
			country:          "FR",
			allowedCountries: []string{"US", "CA", "GB"},
			expected:         false,
		},
		{
			name:             "Empty allowed countries",
			country:          "US",
			allowedCountries: []string{},
			expected:         false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := isCountryAllowed((test.country), test.allowedCountries)
			if result != test.expected {
				t.Errorf("isCountryAllowed expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestBuildResponseMessage(t *testing.T) {
	tests := []struct {
		name     string
		country  string
		allowed  bool
		expected string
	}{
		{
			name:     "Allowed country",
			country:  "US",
			allowed:  true,
			expected: "Country US is allowed",
		},
		{
			name:     "Country NOT allowed",
			country:  "FR",
			allowed:  false,
			expected: "Country FR is not allowed",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := buildResponseMessage(test.country, test.allowed)
			if result != test.expected {
				t.Errorf("buildResponseMessage expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestVerifyIP(t *testing.T) {

	// Load the real GeoIP DB (use the actual path to the .mmdb file)
	err := geoip.InitDB("../data/GeoLite2-Country.mmdb") // Adjust the path as needed
	if err != nil {
		t.Fatalf("Failed to initialize GeoIP DB: %v", err)
	}
	defer geoip.CloseDB()

	tests := []struct {
		name            string
		requestBody     models.VerifyRequest
		expectedStatus  int
		expectedAllowed bool
		expectedMessage string
	}{
		{
			name: "Allowed Country (US)",
			requestBody: models.VerifyRequest{
				IP:               "8.8.8.8",
				AllowedCountries: []string{"US", "FR"},
			},
			expectedStatus:  http.StatusOK,
			expectedAllowed: true,
			expectedMessage: "Country US is allowed",
		},
		{
			name: "Country blocked (US IP not allowed)",
			requestBody: models.VerifyRequest{
				IP:               "8.8.8.8",
				AllowedCountries: []string{"GB", "FR"},
			},
			expectedStatus:  http.StatusOK,
			expectedAllowed: false,
			expectedMessage: "Country US is not allowed",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			// Marshal struct to JSON
			reqBodyBytes, err := json.Marshal(test.requestBody)
			if err != nil {
				t.Fatalf("Failed to marshal request body: %v", err)
			}

			// Create an HTTP request with the JSON body
			req := httptest.NewRequest(http.MethodPost, "/verify", bytes.NewReader(reqBodyBytes))
			req.Header.Set("Content-Type", "application/json")

			// Create a ResponseRecorder to capture the response
			rr := httptest.NewRecorder()

			// Call the handler directly
			VerifyIP(rr, req)
			// Assert the status code
			if rr.Code != http.StatusOK {
				t.Errorf("expected status code 200 OK, got %v", rr.Code)
			}

			// Decode the response into your VerifyResponse struct
			var resp models.VerifyResponse
			if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
				t.Fatalf("Failed to decode response: %v", err)
			}

			// Assert the returned values
			if resp.Allowed != test.expectedAllowed {
				t.Errorf("Expected Allowed to be %v, got %v", test.expectedAllowed, resp.Allowed)
			}
			if resp.Message != test.expectedMessage {
				t.Errorf("Expected message: %s, got %s", test.expectedMessage, resp.Message)
			}
		})
	}

	t.Run("Invalid method", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/verify", nil)
		rr := httptest.NewRecorder()

		VerifyIP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected 400 BadRequest for GET, got %d", rr.Code)
		}
	})

	t.Run("Invalid request body", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/verify", bytes.NewBuffer([]byte("bad json")))
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		VerifyIP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("Expected 400 BadRequest for invalid body, got %d", rr.Code)
		}
	})
}
