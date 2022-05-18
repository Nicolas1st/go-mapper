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

	OccupiedFrom  time.Time
	OccupiedUntil time.Time
}

func (db ParkingDB) CreateSlotReservation(slot ParkingSlot, from, until time.Time) error {
	r := SlotReservation{
		ParkingPlace:   slot.ParkingPlace,
		ParkingPlaceID: int(slot.ParkingPlaceID),
		ParkingSlot:    slot,
		ParkingSlotID:  int(slot.ID),
		OccupiedFrom:   from,
		OccupiedUntil:  until,
	}

	result := db.conn.Create(&r)
	if result.Error != nil {
		return fmt.Errorf("could not store the parking place")
	}

	return nil
}
