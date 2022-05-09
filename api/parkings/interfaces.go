package parkings

import "yaroslavl-parkings/persistence/model"

type DatabaseInterface interface {
	StoreParkingPlace(parkingPlace *model.ParkingPlace) (*model.ParkingPlace, error)
	RemoveParkingPlaceByID(id uint)
	GetAllParkingPlaces() []model.ParkingPlace
}
