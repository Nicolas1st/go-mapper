package rates

import "net/http"

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

// NewRatesHandlers - contstructs all functions needed to update the tarriffs data
// in the database
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
