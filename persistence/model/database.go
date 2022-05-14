package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase(dsn string) *Database {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &Database{db: db}
}

func (db *Database) InitTables() {
	err := db.db.AutoMigrate(&User{}, &ParkingPlace{}, &ParkingSlot{}, &SlotReservation{}, &Purchase{})
	if err != nil {
		panic("Could not initialize the tables in the database")
	}
}
