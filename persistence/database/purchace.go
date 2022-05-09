package database

import (
	"fmt"
	"yaroslavl-parkings/persistence/model"
)

func (db Database) StorePurchase(purchase *model.Purchase) (*model.Purchase, error) {
	result := db.db.Create(purchase)
	if result.Error != nil {
		return purchase, fmt.Errorf("could not store the parking place")
	}

	return purchase, nil
}

func (db Database) GetAllPurchases() []model.Purchase {
	var purchases []model.Purchase
	db.db.Find(purchases)

	return purchases
}
