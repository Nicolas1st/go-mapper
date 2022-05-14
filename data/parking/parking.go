package parking

import "gorm.io/gorm"

type ParkingDB struct {
	conn *gorm.DB
}

func NewParkingDB(conn *gorm.DB) *ParkingDB {
	return &ParkingDB{
		conn: conn,
	}
}
