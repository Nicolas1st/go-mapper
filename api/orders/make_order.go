package orders

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"yaroslavl-parkings/api"
	"yaroslavl-parkings/data/rate"
	"yaroslavl-parkings/data/user"
	"yaroslavl-parkings/pkg/qiwi"

	"github.com/google/uuid"
)

// makeOrders - creates an order in the system,
// creates a payment in the qiwi system,
// returns the link to the payment page
func (d *ordersDependencies) makeOrder(w http.ResponseWriter, r *http.Request) {
	// redirect if not auth
	if !api.IsAuth(d.sessions, r) {
		http.Redirect(w, r, api.DefaultEndpoints.OrderPage, http.StatusSeeOther)
		return
	}

	session, valid := api.GetSessionIfValid(d.sessions, r)
	if !valid {
		http.Redirect(w, r, api.DefaultEndpoints.OrderPage, http.StatusSeeOther)
		return
	}

	// extract values from the form
	startHour, err := strconv.Atoi(r.PostFormValue("start-hour"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	startMinute, err := strconv.Atoi(r.PostFormValue("start-minute"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	endHour, err := strconv.Atoi(r.PostFormValue("end-hour"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	endMinute, err := strconv.Atoi(r.PostFormValue("end-minute"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// compute for how long the user is going to use the parkingr
	var duration float64 = (float64(endHour*60+endMinute) - float64(startHour*60+startMinute)) / 60

	// get time of the day discount
	var discount rate.Discount
	if time.Now().Hour() < 12 {
		temp, _ := d.rateDB.GetSluggishHoursDiscount()
		discount = temp.DiscountInPercents
	} else {
		temp, _ := d.rateDB.GetActiveHoursDiscount()
		discount = temp.DiscountInPercents
	}

	// get the price per hour
	ageCategory := session.User.GetAgeCategory()
	var pricePerHour rate.PricePerHour
	if ageCategory == user.Adult {
		temp, _ := d.rateDB.GetAdultRatePerHour()
		pricePerHour = temp.Price
	} else {
		temp, _ := d.rateDB.GetRetireeRatePerHour()
		pricePerHour = temp.Price
	}

	// compute the check
	sum := int(duration * float64(pricePerHour) * float64(100-discount) / 100)

	// get the qiwi payment url
	timeOut := time.Now().Add(5 * time.Minute)
	qiwiID := uuid.NewString()
	qiwiPaymentFormURL, _ := d.paymenter.CreateNewBill(qiwiID, sum, qiwi.RUB, timeOut, "Parking App")

	// record the information about th order
	orderID, _ := d.orderDB.CreateOrder(*session.User, uint(sum), qiwiPaymentFormURL, qiwiID)

	http.Redirect(w, r, api.DefaultEndpoints.PaymentPage+fmt.Sprintf("?orderID=%v", orderID), http.StatusSeeOther)
}
