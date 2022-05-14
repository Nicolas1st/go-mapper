package parking

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type SlotReservation struct {
	gorm.Model

	ParkingPlaceID int
	ParkingPlace   ParkingPlace

	ParkingSlotID int
	ParkingSlot   ParkingSlot

	OccupiedFrom  *time.Time
	OccupiedUntil *time.Time
}

func (db ParkingDB) StoreSlotReservation(reservation *SlotReservation) (*SlotReservation, error) {
	result := db.conn.Create(reservation)
	if result.Error != nil {
		return reservation, fmt.Errorf("could not store the parking place")
	}

	return reservation, nil
}
