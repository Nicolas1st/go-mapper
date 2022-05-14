package model

import (
	"gorm.io/gorm"
)

type OrderStatus string

const (
	PAID   OrderStatus = "PAID"
	UNPAID OrderStatus = "UNPAID"
)

type Order struct {
	gorm.Model
	UserID            uint
	User              User
	Sum               uint
	Status            OrderStatus
	SlotReservationID uint
	SlotReservation   SlotReservation
}

// CreateOrder - creates order in the database
// return id of the order
func (db *Database) CreateOrder(
	user User,
	sum uint,
	slotReservation SlotReservation,
) (uint, error) {
	order := &Order{
		UserID:            user.ID,
		User:              user,
		Sum:               sum,
		SlotReservationID: slotReservation.ID,
		Status:            UNPAID,
	}

	result := db.db.Create(order)
	return order.ID, result.Error
}

// MarkOrderAsPaid - changes order status to paid
func (db *Database) MarkOrderAsPaid(id uint) error {
	var order Order
	result := db.db.First(&order, id)
	if result.Error != nil {
		return result.Error
	}

	order.Status = PAID
	result = db.db.Save(&order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
