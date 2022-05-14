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
	CreateOrder(user user.User, sum uint, paymentURL, stringID string) (uint, error)
	MarkOrderAsPaid(id uint) error
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

type BillerInterface interface {
	ComputePrice(basePrice uint, discount uint, forHowManyMinutes uint) uint
}

type ParkerInterface interface {
	ReservePlace(parkingID uint, startTime time.Time, forHowManyMinute uint) parking.SlotReservation
}

type ordersDependencies struct {
	orderDB   OrderDatabase
	rateDB    RateDatabase
	sessions  SessionsInterface
	paymenter *qiwi.Paymenter
	biller    BillerInterface
	parker    ParkerInterface
}

type ordersApi struct {
	MakeOrder http.HandlerFunc
}

func NewOrdersApi(
	orderDB OrderDatabase,
	rateDB RateDatabase,
	sessions SessionsInterface,
	paymenter *qiwi.Paymenter,
	biller BillerInterface,
	parker ParkerInterface,
) *ordersApi {
	dependencies := &ordersDependencies{
		orderDB:   orderDB,
		rateDB:    rateDB,
		sessions:  sessions,
		paymenter: paymenter,
		biller:    biller,
		parker:    parker,
	}

	return &ordersApi{
		MakeOrder: dependencies.makeOrder,
	}
}
