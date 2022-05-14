package order

import (
	"yaroslavl-parkings/persistence/parking"
	"yaroslavl-parkings/persistence/user"

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
	User              user.User
	Sum               uint
	Status            OrderStatus
	SlotReservationID uint
	SlotReservation   parking.SlotReservation
}

// CreateOrder - creates order in the database
// return id of the order
func (db *OrderDB) CreateOrder(
	user user.User,
	sum uint,
	slotReservation parking.SlotReservation,
) (uint, error) {
	order := &Order{
		UserID:            user.ID,
		User:              user,
		Sum:               sum,
		SlotReservationID: slotReservation.ID,
		Status:            UNPAID,
	}

	result := db.conn.Create(order)
	return order.ID, result.Error
}

// MarkOrderAsPaid - changes order status to paid
func (db *OrderDB) MarkOrderAsPaid(id uint) error {
	var order Order
	result := db.conn.First(&order, id)
	if result.Error != nil {
		return result.Error
	}

	order.Status = PAID
	result = db.conn.Save(&order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
