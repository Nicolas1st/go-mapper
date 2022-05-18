package parking

import (
	"fmt"
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
	fmt.Println("parking")

	for _, s := range slots {
		fmt.Println("slot")
		var reservations []SlotReservation
		db.conn.Where(&SlotReservation{ParkingSlotID: int(s.ID)}).Find(&reservations)

		numberOfOverlappingReservations := 0
		for _, r := range reservations {
			fmt.Println("reservation")
			if !((startTime.Before(r.OccupiedUntil)) || (r.OccupiedFrom.Before(endTime))) {
				numberOfOverlappingReservations++
			}
		}

		fmt.Println("reservations")
		fmt.Println(numberOfOverlappingReservations)
		if numberOfOverlappingReservations == 0 {
			return s, true
		}

		numberOfOverlappingReservations = 0
	}

	fmt.Println("outside")

	return ParkingSlot{}, false
}
