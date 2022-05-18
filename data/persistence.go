package data

import (
	"yaroslavl-parkings/data/order"
	"yaroslavl-parkings/data/parking"
	"yaroslavl-parkings/data/rate"
	"yaroslavl-parkings/data/user"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Database - struct containing all
// database tables
type Database struct {
	Conn *gorm.DB

	User    *user.UserDB
	Parking *parking.ParkingDB
	Order   *order.OrderDB
	Rate    *rate.RateDB
}

// NewDatabase - sets up a new database connection,
// terminates the program if it's not possible to set up
// the connection
func NewDatabase(dsn string) *Database {
	conn, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &Database{
		Conn: conn,

		User:    user.NewUserDB(conn),
		Parking: parking.NewParkingDB(conn),
		Order:   order.NewOrderDB(conn),
		Rate:    rate.NewRateDB(conn),
	}
}

// InitTables - initializes tables
// if they are not already created
// on the database to which
// connection has be established
func (db *Database) InitTables() {
	err := db.Conn.AutoMigrate(
		&user.User{},
		&parking.ParkingPlace{},
		&parking.ParkingSlot{},
		&parking.SlotReservation{},
		&order.Order{},
		&rate.BaseRate{},
		&rate.PeriodDiscount{},
	)

	if err != nil {
		panic("Could not initialize the tables in the database")
	}
}
