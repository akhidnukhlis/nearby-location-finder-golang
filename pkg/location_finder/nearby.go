package pkg

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/url"
)

// Location represents a geographic location.
type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

// Geocoder is responsible for converting addresses into geographic coordinates.
type Geocoder interface {
	Geocode(address string) (Location, error)
}

// GeocodeResponse is an implementation of the Geocoder interface using the Response Google Maps API.
type GeocodeResponse struct {
	Results []struct {
		Geometry struct {
			Location Location `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
}

// GoogleMapsGeocoder is an implementation of the Geocoder interface using the Google Maps Geocoding API.
type GoogleMapsGeocoder struct {
	APIKey string
}

func (g *GoogleMapsGeocoder) Geocode(address string) (Location, error) {
	// Construct the API request URL
	apiURL := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s", url.QueryEscape(address), g.APIKey)

	// Make the HTTP GET request
	resp, err := http.Get(apiURL)
	if err != nil {
		return Location{}, err
	}
	defer resp.Body.Close()

	// Check for HTTP status code errors
	if resp.StatusCode != http.StatusOK {
		return Location{}, fmt.Errorf("Google Maps API request failed with status code: %d", resp.StatusCode)
	}

	// Parse the JSON response
	var geocodeResponse GeocodeResponse
	if err := json.NewDecoder(resp.Body).Decode(&geocodeResponse); err != nil {
		return Location{}, err
	}

	// Check if any results were found
	if len(geocodeResponse.Results) == 0 {
		return Location{}, fmt.Errorf("No results found for address: %s", address)
	}

	// Extract the latitude and longitude
	location := geocodeResponse.Results[0].Geometry.Location
	return location, nil
}

// DistanceCalculator calculates the distance between two locations.
type DistanceCalculator interface {
	CalculateDistance(loc1, loc2 Location) float64
}

// HaversineCalculator is an implementation of the DistanceCalculator interface using the Haversine formula.
type HaversineCalculator struct{}

func (h *HaversineCalculator) CalculateDistance(loc1, loc2 Location) float64 {
	// Convert latitude and longitude from degrees to radians
	lat1 := loc1.Latitude * (math.Pi / 180.0)
	lon1 := loc1.Longitude * (math.Pi / 180.0)
	lat2 := loc2.Latitude * (math.Pi / 180.0)
	lon2 := loc2.Longitude * (math.Pi / 180.0)

	// Haversine formula
	dLat := lat2 - lat1
	dLon := lon2 - lon1
	a := math.Pow(math.Sin(dLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	R := 6371.0 // Earth's radius in kilometers
	distance := R * c

	return distance
}

// LocationFinder finds the nearest location using a Geocoder and DistanceCalculator.
type LocationFinder struct {
	Geocoder           Geocoder
	DistanceCalculator DistanceCalculator
	Locations          map[string]Location
}

func (lf *LocationFinder) FindNearestLocation(address string) (string, float64, error) {
	targetLocation, err := lf.Geocoder.Geocode(address)
	if err != nil {
		return "", 0, err
	}

	var nearestName string
	var nearestDistance float64

	for name, loc := range lf.Locations {
		distance := lf.DistanceCalculator.CalculateDistance(targetLocation, loc)
		if nearestName == "" || distance < nearestDistance {
			nearestName = name
			nearestDistance = distance
		}
	}

	return nearestName, nearestDistance, nil
}
