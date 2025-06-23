package geoip

import (
	"fmt"
	"net"
	"sync"

	"github.com/oschwald/geoip2-golang"
)

// AI helped set up the global pointer for the GeoIP database and to initially get the files i needed to download
var (
	db   *geoip2.Reader
	once sync.Once
)

// for next 2 code blocks, most of this was direct from documentation but AI predictive text
// helped with edge case i wouldnt have otherwise had and made createing the functions way easier

// initiate GeoIP database
func InitDB(path string) error {
	var err error
	once.Do(func() {
		db, err = geoip2.Open(path)
	})
	if err != nil {
		return fmt.Errorf("failed to open GeoIP database: %w", err)
	}
	return nil
}

// lookup country by IP Address
func GetCountryByIP(ipStr string) (string, error) {
	if db == nil {
		return "", fmt.Errorf("GeoIP database is not initialized")
	}

	ip := net.ParseIP(ipStr)
	if ip == nil {
		return "", fmt.Errorf("invalid IP address: %s", ipStr)
	}

	record, err := db.Country(ip)
	if err != nil {
		return "", fmt.Errorf("failed to get country for IP %s: %w", ipStr, err)
	}

	if record.Country.IsoCode == "" {
		return "", fmt.Errorf("country not found for IP %s", ipStr)
	}

	return record.Country.IsoCode, nil
}

// this was recommended by ai as i had place a defer db.Close() in the wrong location which was in initDB instead of in main.go
// close database
func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
