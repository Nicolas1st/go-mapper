package parkings

import (
	"encoding/json"
	"net/http"
	"yaroslavl-parkings/data/parking"
)

type parkingsDependencies struct {
	datatbase DatabaseInterface
}

func (resource *parkingsDependencies) createParkingPlace(w http.ResponseWriter, r *http.Request) {
	// the parsed values are going to be stored here
	var parkingPlace parking.ParkingPlace

	// parsing the body of json request
	err := json.NewDecoder(r.Body).Decode(&parkingPlace)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// storing the new parking place into the database
	resource.datatbase.StoreParkingPlace(&parkingPlace)

	json.NewEncoder(w).Encode(parkingPlace)
}

// removeParkingByID - removes a parking in the database,
// if no parking id has been provided,
//returns bad request status code
func (resource *parkingsDependencies) removeParkingByID(w http.ResponseWriter, r *http.Request) {
	jsonBody := struct {
		ID uint `json:"ID"`
	}{}

	// decoding
	err := json.NewDecoder(r.Body).Decode(&jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resource.datatbase.RemoveParkingPlaceByID(jsonBody.ID)
	json.NewEncoder(w).Encode(jsonBody)
}

// getAllParkings - returns all parkings places from the database
func (resource *parkingsDependencies) getAllParkings(w http.ResponseWriter, r *http.Request) {
	// querying the database
	parkingsPlaces := resource.datatbase.GetAllParkingPlaces()

	// sending to the user
	json.NewEncoder(w).Encode(&parkingsPlaces)
}
