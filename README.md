# Nearest Location Finder

Nearest Location Finder is a Go application that allows you to find the nearest location based on geographic coordinates using Google Maps integration and the Haversine formula for distance calculation.

## Getting Started

Follow these steps to get the application up and running:

### Prerequisites

- Go (Golang) installed on your machine
- Google Maps API Key (Instructions on how to obtain one can be found [here](https://developers.google.com/maps/gmp-get-started))

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/akhidnukhlis/nearby-location-finder-golang
   ```
   
2. Navigate to the project directory:

   ```
   cd nearby-location-finder-golang
   ```

3. Create a .env file in the project root and add your Google Maps API Key:

   ```
   GOOGLE_MAPS_API_KEY=YOUR_API_KEY
   ```

4. Build and run the application:

   ```
   go build
   ./nearby-location-finder-golang
   ```

### Installation

To find the nearest location, run the application with the target address as a command-line argument:

   ```
   ./nearby-location-finder-golang "Target Address"
   ```
** Replace "Target Address" with the address you want to find the nearest location to.

### Example
   ```
   ./nearby-location-finder-golang "1600 Amphitheatre Parkway, Mountain View, CA"
   ```

### Features
   * Geocoding of addresses using the Google Maps Geocoding API.
   * Calculation of distances between locations using the Haversine formula.
   * Finding the nearest location from a list of predefined locations.
   * Easy-to-use command-line interface.

### Acknowledgments
   * [Google Maps Platform](https://developers.google.com/maps/documentation)

### Author
   * Akhid Nukhlis
   * GitHub: [@akhidnukhlis](https://github.com/akhidnukhlis)
   * Email: nukhlis@gmail.com

