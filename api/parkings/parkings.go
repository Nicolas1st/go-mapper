package parkings

import (
	"fmt"
	"net/http"
)

func NewParkingsApi(database DatabaseInterface) http.HandlerFunc {
	parkingsResource := parkingsDependencies{
		datatbase: database,
	}

	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hit router")
		fmt.Println(r.URL.Path)
		fmt.Println(r.Method)
		if r.URL.Path == "/all" && r.Method == http.MethodGet {
			fmt.Println("hit all")
			parkingsResource.getAllParkings(w, r)
		} else if r.URL.Path == "/" && r.Method == http.MethodPost {
			fmt.Println("hit create")
			parkingsResource.createParkingPlace(w, r)
		} else if r.URL.Path == "/" && r.Method == http.MethodDelete {
			fmt.Println("hit delete")
			parkingsResource.removeParkingByID(w, r)
		} else {
			fmt.Println("dit not hit anything")
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
}
