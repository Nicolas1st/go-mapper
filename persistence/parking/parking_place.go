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

type ParkingSlot struct {
	gorm.Model
	Number         int
	ParkingPlaceID int
	ParkingPlace   ParkingPlace
}

type SlotReservation struct {
	gorm.Model

	ParkingPlaceID int
	ParkingPlace   ParkingPlace

	ParkingSlotID int
	ParkingSlot   ParkingSlot

	OccupiedFrom  *time.Time
	OccupiedUntil *time.Time
}

func (db ParkingDB) StoreParkingPlace(parkingPlace *ParkingPlace) (*ParkingPlace, error) {
	result := db.conn.Create(parkingPlace)
	if result.Error != nil {
		return parkingPlace, fmt.Errorf("could not store the parking place")
	}

	return parkingPlace, nil
}

func (db ParkingDB) RemoveParkingPlaceByID(id uint) {
	db.conn.Delete(&ParkingPlace{}, id)
}

func (db ParkingDB) GetAllParkingPlaces() []ParkingPlace {
	var parkingPlaces []ParkingPlace
	db.conn.Find(&parkingPlaces)

	return parkingPlaces
}
