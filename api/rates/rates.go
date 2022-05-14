package rates

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"yaroslavl-parkings/api"
)

func (d *ratesDependencies) UpdateActiveHoursDiscount(w http.ResponseWriter, r *http.Request) error {
	if !api.IsAuthAndAdmin(d.sessions, r) {
		return errors.New("not an admin")
	}

	discount := r.PostFormValue("discount")

	discountInt, err := strconv.Atoi(discount)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return d.db.UpdateActiveHoursDiscount(uint(discountInt))
}

func (d *ratesDependencies) UpdateSluggishHoursDiscount(w http.ResponseWriter, r *http.Request) error {
	discount := r.PostFormValue("discount")

	discountInt, err := strconv.Atoi(discount)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return d.db.UpdateSluggishHoursDiscount(uint(discountInt))
}

func (d *ratesDependencies) UpdateAdultHourlyRate(w http.ResponseWriter, r *http.Request) error {
	rate := r.PostFormValue("rate")

	rateInt, err := strconv.Atoi(rate)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return d.db.UpdateAdultPricePerHour(uint(rateInt))
}

func (d *ratesDependencies) UpdateRetireeHourlyRate(w http.ResponseWriter, r *http.Request) error {
	rate := r.PostFormValue("rate")

	rateInt, err := strconv.Atoi(rate)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return d.db.UpdateRetireePricePerHour(uint(rateInt))
}
