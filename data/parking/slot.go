package parking

import (
	"time"

	"gorm.io/gorm"
)

type ParkingSlot struct {
	gorm.Model
	Number         int
	ParkingPlaceID uint
	ParkingPlace   ParkingPlace

	OccupiedFrom  *time.Time
	OccupiedUntil *time.Time
}

func (db *ParkingDB) CreateParkingSlot(number int, parkingPlace ParkingPlace) error {
	slot := ParkingSlot{
		Number:         number,
		ParkingPlaceID: parkingPlace.ID,
		ParkingPlace:   parkingPlace,

		OccupiedFrom:  nil,
		OccupiedUntil: nil,
	}

	return db.conn.Create(&slot).Error
}

func (db *ParkingDB) OccupySlot(id uint, from, until *time.Time) error {
	var slot ParkingSlot

	if err := db.conn.First(&slot, id).Error; err != nil {
		return err
	}

	slot.OccupiedFrom = from
	slot.OccupiedUntil = until

	return nil
}

func (db *ParkingDB) CheckIfSlotIsFree(from, until *time.Time) bool {
	return true
}

func (db *ParkingDB) FindFreeSlots(parkingID uint) []ParkingSlot {
	return []ParkingSlot{}
}
