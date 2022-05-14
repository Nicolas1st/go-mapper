package order

import "gorm.io/gorm"

type OrderDB struct {
	conn *gorm.DB
}

func NewOrderDB(conn *gorm.DB) *OrderDB {
	return &OrderDB{
		conn: conn,
	}
}
