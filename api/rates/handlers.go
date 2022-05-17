package rates

import (
	"errors"
	"net/http"
	"strconv"
	"yaroslavl-parkings/api"
)

// UpdateActiveHoursDiscount - update discount in the database
func (d *ratesDependencies) UpdateActiveHoursDiscount(w http.ResponseWriter, r *http.Request) error {
	if !api.IsAuthAndAdmin(d.sessions, r) {
		return errors.New("not an admin")
	}

	discount := r.PostFormValue("discount")

	discountInt, err := strconv.Atoi(discount)
	if err != nil {
		return err
	}

	return d.db.UpdateActiveHoursDiscount(uint(discountInt))
}

// UpdateSluggishHoursDiscount - update discount in the database
func (d *ratesDependencies) UpdateSluggishHoursDiscount(w http.ResponseWriter, r *http.Request) error {
	discount := r.PostFormValue("discount")

	discountInt, err := strconv.Atoi(discount)
	if err != nil {
		return err
	}

	return d.db.UpdateSluggishHoursDiscount(uint(discountInt))
}

// UpdateAdultHourlyRate - updates the base price for an hour of parking in the database
func (d *ratesDependencies) UpdateAdultHourlyRate(w http.ResponseWriter, r *http.Request) error {
	rate := r.PostFormValue("rate")

	rateInt, err := strconv.Atoi(rate)
	if err != nil {
		return err
	}

	return d.db.UpdateAdultPricePerHour(uint(rateInt))
}

// UpdateRetireeHourlyRate - updates the base price for an hour of parking in the database
func (d *ratesDependencies) UpdateRetireeHourlyRate(w http.ResponseWriter, r *http.Request) error {
	rate := r.PostFormValue("rate")

	rateInt, err := strconv.Atoi(rate)
	if err != nil {
		return err
	}

	return d.db.UpdateRetireePricePerHour(uint(rateInt))
}
