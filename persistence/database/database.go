package database

import (
	"yaroslavl-parkings/persistence/model"

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
	err := db.db.AutoMigrate(&model.User{}, &model.ParkingPlace{}, &model.ParkingSlot{}, &model.SlotReservation{}, &model.Purchase{})
	if err != nil {
		panic("Could not initialize the tables in the database")
	}
}
