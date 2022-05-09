package database

import (
	"fmt"
	"yaroslavl-parkings/persistence/model"
)

func (db Database) StoreParkingPlace(parkingPlace *model.ParkingPlace) (*model.ParkingPlace, error) {
	result := db.db.Create(parkingPlace)
	if result.Error != nil {
		return parkingPlace, fmt.Errorf("could not store the parking place")
	}

	return parkingPlace, nil
}

func (db Database) RemoveParkingPlaceByID(id uint) {
	db.db.Delete(&model.ParkingPlace{}, id)
}

func (db Database) GetAllParkingPlaces() []model.ParkingPlace {
	var parkingPlaces []model.ParkingPlace
	db.db.Find(&parkingPlaces)

	return parkingPlaces
}
