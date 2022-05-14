package model

import "fmt"

func (db Database) StoreSlotReservation(reservation *SlotReservation) (*SlotReservation, error) {
	result := db.db.Create(reservation)
	if result.Error != nil {
		return reservation, fmt.Errorf("could not store the parking place")
	}

	return reservation, nil
}
