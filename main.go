package main

import (
	"fmt"
	pkg "playground/nearby-location-finder-golang/pkg/location_finder"
)

func main() {
	apiKey := "YOUR_API_KEY" // Replace with your Google Maps API key

	geocoder := &pkg.GoogleMapsGeocoder{APIKey: apiKey}

	address := "1600 Amphitheatre Parkway, Mountain View, CA" // Replace with the address you want to geocode

	location, err := geocoder.Geocode(address)
	if err != nil {
		fmt.Println("Error geocoding address:", err)
		return
	}

	fmt.Printf("Geocoded location for address '%s':\n", address)
	fmt.Printf("Latitude: %.6f\n", location.Latitude)
	fmt.Printf("Longitude: %.6f\n", location.Longitude)
}
