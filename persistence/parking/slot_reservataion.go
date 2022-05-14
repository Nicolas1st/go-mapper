package parking

import "fmt"

func (db ParkingDB) StoreSlotReservation(reservation *SlotReservation) (*SlotReservation, error) {
	result := db.conn.Create(reservation)
	if result.Error != nil {
		return reservation, fmt.Errorf("could not store the parking place")
	}

	return reservation, nil
}
