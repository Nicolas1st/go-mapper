package rate

import "gorm.io/gorm"

type RateDB struct {
	conn *gorm.DB
}

func NewRateDB(conn *gorm.DB) *RateDB {
	return &RateDB{
		conn: conn,
	}
}
