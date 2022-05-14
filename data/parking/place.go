package parking

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type ParkingPlace struct {
	gorm.Model
	Address       string
	NumberOfSlots int
	Latitude      float64
	Longitude     float64
}

// StoreParkingPlace - stores parking place in the database
// returns the parking place and the error result
func (db ParkingDB) StoreParkingPlace(parkingPlace *ParkingPlace) (*ParkingPlace, error) {
	result := db.conn.Create(parkingPlace)
	if result.Error != nil {
		return parkingPlace, fmt.Errorf("could not store the parking place")
	}

	return parkingPlace, nil
}

// RemoveParkingPlaceByID - removes parking place from the databaes
func (db ParkingDB) RemoveParkingPlaceByID(id uint) {
	db.conn.Delete(&ParkingPlace{}, id)
}

// GetAllParkingPlaces - returns all parking places
func (db ParkingDB) GetAllParkingPlaces() []ParkingPlace {
	var parkingPlaces []ParkingPlace
	db.conn.Find(&parkingPlaces)

	return parkingPlaces
}

func (db ParkingDB) ReservePlace(parkingID uint, startTime time.Time, duration time.Duration) SlotReservation {
	return SlotReservation{}
}
