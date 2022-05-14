package rates

import (
	"net/http"
	"yaroslavl-parkings/data/sessionstorer"
)

type DatabaseInterface interface {
	UpdateAdultPricePerHour(pricePerHour uint) error
	UpdateRetireePricePerHour(pricePerHour uint) error

	UpdateActiveHoursDiscount(discountInPercents uint) error
	UpdateSluggishHoursDiscount(discountInPercents uint) error
}

type SessionsInterface interface {
	IsSessionValid(sessionToken string) (*sessionstorer.Session, bool)
}

type ratesDependencies struct {
	db       DatabaseInterface
	sessions SessionsInterface
}

type RatesHandlers struct {
	UpdateActiveHoursDiscount   func(w http.ResponseWriter, r *http.Request) error
	UpdateSluggishHoursDiscount func(w http.ResponseWriter, r *http.Request) error
	UpdateAdultHourlyRate       func(w http.ResponseWriter, r *http.Request) error
	UpdateRetireeHourlyRate     func(w http.ResponseWriter, r *http.Request) error
}

func NewRatesHandlers(db DatabaseInterface, sessions SessionsInterface) *RatesHandlers {
	dependencies := ratesDependencies{
		db:       db,
		sessions: sessions,
	}

	return &RatesHandlers{
		UpdateActiveHoursDiscount:   dependencies.UpdateActiveHoursDiscount,
		UpdateSluggishHoursDiscount: dependencies.UpdateSluggishHoursDiscount,
		UpdateAdultHourlyRate:       dependencies.UpdateAdultHourlyRate,
		UpdateRetireeHourlyRate:     dependencies.UpdateRetireeHourlyRate,
	}
}
