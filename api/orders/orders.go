package orders

import (
	"net/http"
	"time"
	"yaroslavl-parkings/data/parking"
	"yaroslavl-parkings/data/rate"
	"yaroslavl-parkings/data/sessionstorer"
	"yaroslavl-parkings/data/user"
	"yaroslavl-parkings/pkg/qiwi"
)

type OrderDatabase interface {
	CreateOrder(user user.User, sum uint, paymentURL, stringID string, paymentTimeout time.Time) (uint, error)
}

type RateDatabase interface {
	GetActiveHoursDiscount() (rate.PeriodDiscount, error)
	GetSluggishHoursDiscount() (rate.PeriodDiscount, error)
	GetAdultRatePerHour() (rate.BaseRate, error)
	GetRetireeRatePerHour() (rate.BaseRate, error)
}

type SessionsInterface interface {
	IsSessionValid(sessionToken string) (*sessionstorer.Session, bool)
}

type ParkerInterface interface {
	FindSlot(parkingID uint, startTime, endTime time.Time) (parking.ParkingSlot, bool)
	CreateSlotReservation(slot parking.ParkingSlot, from, until time.Time) error
}

type ordersDependencies struct {
	orderDB   OrderDatabase
	rateDB    RateDatabase
	sessions  SessionsInterface
	paymenter *qiwi.Paymenter
	parker    ParkerInterface
}

type ordersApi struct {
	MakeOrder http.HandlerFunc
}

// NewOrdersApi - construct a make order function
func NewOrdersApi(
	orderDB OrderDatabase,
	rateDB RateDatabase,
	sessions SessionsInterface,
	paymenter *qiwi.Paymenter,
	parker ParkerInterface,
) *ordersApi {
	dependencies := &ordersDependencies{
		orderDB:   orderDB,
		rateDB:    rateDB,
		sessions:  sessions,
		paymenter: paymenter,
		parker:    parker,
	}

	return &ordersApi{
		MakeOrder: dependencies.makeOrder,
	}
}
