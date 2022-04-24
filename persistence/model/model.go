package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string
	PasswordHash string
}

type ParkingPlace struct {
	gorm.Model
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

type Purchase struct {
	gorm.Model
	Amount            int
	UserID            int
	User              User
	SlotReservationID int
	SlotReservation   SlotReservation
}
