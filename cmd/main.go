package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"yaroslavl-parkings/api"
	"yaroslavl-parkings/api/auth"
	"yaroslavl-parkings/api/middlewares"
	"yaroslavl-parkings/api/orders"
	"yaroslavl-parkings/api/parkings"
	"yaroslavl-parkings/api/personal"
	"yaroslavl-parkings/api/rates"
	"yaroslavl-parkings/api/views"
	"yaroslavl-parkings/data"
	"yaroslavl-parkings/data/sessionstorer"
	"yaroslavl-parkings/pkg/biller"
	"yaroslavl-parkings/pkg/qiwi"
)

func main() {
	// get private key
	privateKey, exists := os.LookupEnv("qiwiKey")
	if !exists {
		panic("must have a qiwi privateKey to launch the application")
	}

	// init qiwi
	paymenter := qiwi.NewPaymenter(privateKey)

	// init database
	dsn := "host=localhost user=postgres password=password dbname=parkings port=5432 sslmode=disable"
	db := data.NewDatabase(dsn)
	db.InitTables()

	// create adming user if not already created
	db.User.SetUpAdminAccount("admin", "admin@admin.admin", "admin", 25)

	// set up tarrifs if not set up already
	db.Rate.SetUpAgeCategoryBasedHourlyRates(20, 15)
	db.Rate.SetUpDiscount(0, 20)

	// init session storage
	sessionStorage := sessionstorer.NewSessionStorer(time.Minute * 5)

	// serve static files
	files := http.FileServer(http.Dir("./web"))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// create views
	views := views.NewViews("api/views/html", sessionStorage, db.Rate, db.Order, paymenter)

	// init authentication functions
	auth := auth.NewAuthHandlers(sessionStorage, db.User)

	// set up routes for views
	http.HandleFunc(api.DefaultEndpoints.Root, views.SignIn)
	http.HandleFunc(api.DefaultEndpoints.LoginPage, views.SignIn)
	http.HandleFunc(api.DefaultEndpoints.SignUpPage, views.SignUp)
	http.HandleFunc(api.DefaultEndpoints.OrderPage, views.MakeOrder)
	http.HandleFunc(api.DefaultEndpoints.AddParkingPage, views.AddParkingPlace)
	http.HandleFunc(api.DefaultEndpoints.RemoveParkingPage, views.RemoveParkingPlace)
	http.HandleFunc(api.DefaultEndpoints.ProfilePage, views.Profile)

	http.HandleFunc(api.DefaultEndpoints.SeePricings, views.SeePricing)
	http.HandleFunc(api.DefaultEndpoints.ChangeActiveHoursDiscountPage, views.ChangeActiveHoursDiscount)
	http.HandleFunc(api.DefaultEndpoints.ChangeSluggishHoursDiscountPage, views.ChangeSluggishHoursDiscount)
	http.HandleFunc(api.DefaultEndpoints.ChangeAdultBaseRatePage, views.ChangeAdultBaseRate)
	http.HandleFunc(api.DefaultEndpoints.ChangeRetireeBaseRatePage, views.ChangeRetireeBaseRate)
	http.HandleFunc(api.DefaultEndpoints.PaymentPage, views.PaymentPage)
	http.HandleFunc(api.DefaultEndpoints.OrdersPage, views.OrdersPage)

	// account handlers
	accounts := personal.NewPersonalDataApi(db.User, sessionStorage)
	signupResultRedirector := middlewares.BuildRedirectOnApiCallResult(api.DefaultEndpoints.LoginPage, api.DefaultEndpoints.SignUpPage)
	http.HandleFunc(api.DefaultEndpoints.CreateAccountHandler, signupResultRedirector(accounts.CreateAccount))

	// setup authentication
	lougoutResultRedirector := middlewares.BuildRedirectOnApiCallResult(api.DefaultEndpoints.LoginPage, api.DefaultEndpoints.OrderPage)
	http.HandleFunc(api.DefaultEndpoints.LogoutHandler, lougoutResultRedirector(auth.LogoutHandler))

	authenticationResultRedirector := middlewares.BuildRedirectOnApiCallResult(api.DefaultEndpoints.OrderPage, api.DefaultEndpoints.LoginPage)
	http.HandleFunc(api.DefaultEndpoints.AuthenticateHandler, authenticationResultRedirector(auth.LoginHandler))

	// add parkings api
	parkingsApi := parkings.NewParkingsApi(db.Parking)
	http.Handle("/parkings/", http.StripPrefix("/parkings", parkingsApi))

	// add rates api
	ratesApiRedirector := middlewares.BuildRedirectOnApiCallResult(api.DefaultEndpoints.SeePricings, api.DefaultEndpoints.SeePricings)
	ratesApi := rates.NewRatesHandlers(db.Rate, sessionStorage)
	http.HandleFunc(api.DefaultEndpoints.ChangeActiveHoursDiscountHandler, ratesApiRedirector(ratesApi.UpdateActiveHoursDiscount))
	http.HandleFunc(api.DefaultEndpoints.ChangeSluggishHoursDiscountHandler, ratesApiRedirector(ratesApi.UpdateSluggishHoursDiscount))
	http.HandleFunc(api.DefaultEndpoints.ChangeAdultBaseRateHandler, ratesApiRedirector(ratesApi.UpdateAdultHourlyRate))
	http.HandleFunc(api.DefaultEndpoints.ChangeRetireeBaseRateHandler, ratesApiRedirector(ratesApi.UpdateRetireeHourlyRate))

	// add orders api
	biller := biller.NewBiller()
	orders := orders.NewOrdersApi(db.Order, db.Rate, sessionStorage, paymenter, biller, db.Parking)
	http.Handle(api.DefaultEndpoints.CreateOrderHandler, orders.MakeOrder)

	fmt.Println("Server started on http://localhost:8880")
	http.ListenAndServe(":8880", nil)
}
