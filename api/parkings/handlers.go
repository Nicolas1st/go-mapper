package parkings

import (
	"encoding/json"
	"fmt"
	"net/http"
	"yaroslavl-parkings/persistence/model"
)

type parkingsDependencies struct {
	datatbase DatabaseInterface
}

func (resource *parkingsDependencies) createParkingPlace(w http.ResponseWriter, r *http.Request) {
	// the parsed values are going to be stored here
	var parkingPlace model.ParkingPlace

	// parsing the body of json request
	err := json.NewDecoder(r.Body).Decode(&parkingPlace)
	fmt.Println(parkingPlace)
	fmt.Println(err)

	// return an error if parsing fails
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// storing the new parking place into the database
	resource.datatbase.StoreParkingPlace(&parkingPlace)

	json.NewEncoder(w).Encode(parkingPlace)
}

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

func (resource *parkingsDependencies) getAllParkings(w http.ResponseWriter, r *http.Request) {
	// querying the database
	parkingsPlaces := resource.datatbase.GetAllParkingPlaces()

	// sending to the user
	json.NewEncoder(w).Encode(&parkingsPlaces)
}
