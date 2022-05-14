package order

import (
	"fmt"
	"yaroslavl-parkings/persistence/parking"
	"yaroslavl-parkings/persistence/user"

	"gorm.io/gorm"
)

type Purchase struct {
	gorm.Model
	Amount            int
	UserID            int
	User              user.User
	SlotReservationID int
	SlotReservation   parking.SlotReservation
}

func (db OrderDB) StorePurchase(purchase *Purchase) (*Purchase, error) {
	result := db.conn.Create(purchase)
	if result.Error != nil {
		return purchase, fmt.Errorf("could not store the parking place")
	}

	return purchase, nil
}

func (db OrderDB) GetAllPurchases() []Purchase {
	var purchases []Purchase
	db.conn.Find(purchases)

	return purchases
}
