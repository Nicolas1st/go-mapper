package parkings

import "yaroslavl-parkings/data/parking"

type DatabaseInterface interface {
	StoreParkingPlace(parkingPlace *parking.ParkingPlace) (*parking.ParkingPlace, error)
	RemoveParkingPlaceByID(id uint)
	GetAllParkingPlaces() []parking.ParkingPlace
}
