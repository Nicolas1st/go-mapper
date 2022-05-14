package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model
	Amount            int
	UserID            int
	User              User
	SlotReservationID int
	SlotReservation   SlotReservation
}

func (db Database) StorePurchase(purchase *Purchase) (*Purchase, error) {
	result := db.db.Create(purchase)
	if result.Error != nil {
		return purchase, fmt.Errorf("could not store the parking place")
	}

	return purchase, nil
}

func (db Database) GetAllPurchases() []Purchase {
	var purchases []Purchase
	db.db.Find(purchases)

	return purchases
}
