package order

import (
	"errors"
	"yaroslavl-parkings/data/user"

	"gorm.io/gorm"
)

type OrderStatus string

const (
	PAID   OrderStatus = "PAID"
	UNPAID OrderStatus = "UNPAID"
)

type Order struct {
	gorm.Model
	StringID   string
	UserID     uint
	User       user.User
	Sum        uint
	Status     OrderStatus
	PaymentURL string
}

// CreateOrder - creates order in the database
// return id of the order
func (db *OrderDB) CreateOrder(
	user user.User,
	sum uint,
	paymentURL,
	stringID string,
) (uint, error) {
	order := &Order{
		UserID:     user.ID,
		User:       user,
		Sum:        sum,
		Status:     UNPAID,
		PaymentURL: paymentURL,
		StringID:   stringID,
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

// GetOrderPaymentURLByID - returns the url
// needed to complete the payment
func (db *OrderDB) GetOrderPaymentURLByID(id uint) string {
	var order Order
	result := db.conn.First(&order, id)
	if result.Error != nil {
		return ""
	}

	return order.PaymentURL
}

// GetOrderByID - returns an order based on id
func (db *OrderDB) GetOrderByID(id uint) (*Order, error) {
	var order Order
	result := db.conn.First(&order, id)
	if result.Error != nil {
		return &Order{}, errors.New("not found")
	}

	return &order, nil
}

// GetAllOrders - returns all orders
func (db *OrderDB) GetAllOrders() []Order {
	var orders []Order
	db.conn.Find(&orders)

	return orders
}

// GetAllOrdersByUserID - gets all orders for a specific user
func (db *OrderDB) GetAllOrdersByUserID(uid uint) []Order {
	var orders []Order
	db.conn.Where(&Order{UserID: uid}).Find(&orders)

	return orders
}
