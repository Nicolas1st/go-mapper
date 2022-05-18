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
}

// CreateParkingSlot - creates parking slot
func (db *ParkingDB) CreateParkingSlot(number int, parkingPlace ParkingPlace) error {
	slot := ParkingSlot{
		Number:         number,
		ParkingPlaceID: parkingPlace.ID,
		ParkingPlace:   parkingPlace,
	}

	return db.conn.Create(&slot).Error
}

// ReservePlace - reserves a place in specifed parking
func (db ParkingDB) FindSlot(parkingID uint, startTime, endTime time.Time) (ParkingSlot, bool) {
	var slots []ParkingSlot
	db.conn.Where(&ParkingSlot{ParkingPlaceID: parkingID}).Find(&slots)

	for _, s := range slots {
		var reservations []SlotReservation
		db.conn.Where(&SlotReservation{ParkingSlotID: int(s.ID)}).Find(&reservations)

		numberOfOverlappingReservations := 0
		for _, r := range reservations {
			if !((startTime.Before(r.OccupiedUntil)) || (r.OccupiedFrom.Before(endTime))) {
				numberOfOverlappingReservations++
			}
		}

		if numberOfOverlappingReservations == 0 {
			return s, true
		}

		numberOfOverlappingReservations = 0
	}

	return ParkingSlot{}, false
}
