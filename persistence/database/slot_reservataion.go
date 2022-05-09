package database

import (
	"fmt"
	"yaroslavl-parkings/persistence/model"
)

func (db Database) StoreSlotReservation(reservation *model.SlotReservation) (*model.SlotReservation, error) {
	result := db.db.Create(reservation)
	if result.Error != nil {
		return reservation, fmt.Errorf("could not store the parking place")
	}

	return reservation, nil
}
