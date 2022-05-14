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

func (d *ordersDependencies) makeOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hit order")
	if !api.IsAuth(d.sessions, r) {
		http.Redirect(w, r, api.DefaultEndpoints.OrderPage, http.StatusSeeOther)
		fmt.Println("is auth")
		return
	}

	session, valid := api.GetSessionIfValid(d.sessions, r)
	if !valid {
		http.Redirect(w, r, api.DefaultEndpoints.OrderPage, http.StatusSeeOther)
		fmt.Println("is valid")
		return
	}

	startHour, _ := strconv.Atoi(r.PostFormValue("start-hour"))
	startMinute, _ := strconv.Atoi(r.PostFormValue("start-minute"))

	endHour, _ := strconv.Atoi(r.PostFormValue("end-hour"))
	endMinute, _ := strconv.Atoi(r.PostFormValue("end-minute"))

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

	// after this time it won't possible to complete the order
	timeOut := time.Now().Add(5 * time.Minute)
	qiwiID := uuid.NewString()
	qiwiPaymentFormURL, _ := d.paymenter.CreateNewBill(qiwiID, sum, qiwi.RUB, timeOut, "Parking App")
	orderID, _ := d.orderDB.CreateOrder(*session.User, uint(sum), qiwiPaymentFormURL, qiwiID)

	http.Redirect(w, r, api.DefaultEndpoints.PaymentPage+fmt.Sprintf("?orderID=%v", orderID), http.StatusSeeOther)
}
