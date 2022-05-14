package persistence

import (
	"yaroslavl-parkings/persistence/order"
	"yaroslavl-parkings/persistence/parking"
	"yaroslavl-parkings/persistence/user"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Database - struct containing all
// database tables
type Database struct {
	Conn *gorm.DB

	User    *user.UserDB
	Parking *parking.ParkingDB
	Order   *order.OrderDB
}

// NewDatabase - sets up a new database connection,
// terminates the program if it's not possible to set up
// the connection
func NewDatabase(dsn string) *Database {
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &Database{
		Conn: conn,

		User:    user.NewUserDB(conn),
		Parking: parking.NewParkingDB(conn),
		Order:   order.NewOrderDB(conn),
	}
}

// InitTables - initializes tables
// if they are not already created
// on the database to which
// connection has be established
func (db *Database) InitTables() {
	err := db.Conn.AutoMigrate(&user.User{}, &parking.ParkingPlace{}, &parking.ParkingSlot{}, &parking.SlotReservation{}, &order.Purchase{})
	if err != nil {
		panic("Could not initialize the tables in the database")
	}
}
