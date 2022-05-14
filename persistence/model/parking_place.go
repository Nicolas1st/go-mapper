package model

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

func (db Database) StoreParkingPlace(parkingPlace *ParkingPlace) (*ParkingPlace, error) {
	result := db.db.Create(parkingPlace)
	if result.Error != nil {
		return parkingPlace, fmt.Errorf("could not store the parking place")
	}

	return parkingPlace, nil
}

func (db Database) RemoveParkingPlaceByID(id uint) {
	db.db.Delete(&ParkingPlace{}, id)
}

func (db Database) GetAllParkingPlaces() []ParkingPlace {
	var parkingPlaces []ParkingPlace
	db.db.Find(&parkingPlaces)

	return parkingPlaces
}
