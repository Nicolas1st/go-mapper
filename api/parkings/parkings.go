package parkings

import (
	"net/http"
)

// NewParkingsApi - returns a router for parkings handlers
func NewParkingsApi(database DatabaseInterface) http.HandlerFunc {
	parkingsResource := parkingsDependencies{
		datatbase: database,
	}

	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/all" && r.Method == http.MethodGet {
			parkingsResource.getAllParkings(w, r)
		} else if r.URL.Path == "/" && r.Method == http.MethodPost {
			parkingsResource.createParkingPlace(w, r)
		} else if r.URL.Path == "/" && r.Method == http.MethodDelete {
			parkingsResource.removeParkingByID(w, r)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}
}
