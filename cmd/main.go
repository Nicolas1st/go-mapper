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
	"yaroslavl-parkings/pkg/qiwi"
)

func getEnvVar(key string) string {
	envvar, exists := os.LookupEnv(key)
	if !exists {
		panic("provide the " + key + " env var")
	}

	return envvar
}

func main() {
	// read in the env vars
	APP_QIWIKEY := getEnvVar("APP_QIWIKEY")
	APP_DATABASE_HOST := getEnvVar("APP_DATABASE_HOST")
	APP_DATABASE_PORT := getEnvVar("APP_DATABASE_PORT")
	APP_DATABASE_PASSWORD := getEnvVar("APP_DATABASE_PASSWORD")
	APP_DATABASE_NAME := getEnvVar("APP_DATABASE_NAME")
	APP_DATABASE_USER := getEnvVar("APP_DATABASE_USER")

	// init qiwi
	paymenter := qiwi.NewPaymenter(APP_QIWIKEY)

	// init database
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", APP_DATABASE_HOST, APP_DATABASE_USER, APP_DATABASE_PASSWORD, APP_DATABASE_NAME, APP_DATABASE_PORT)
	db := data.NewDatabase(dsn)
	db.InitTables()

	// set up initial app state if not already set up
	db.User.SetUpAdminAccount("admin", "admin@admin.admin", "admin", 25)
	db.Rate.SetUpAgeCategoryBasedHourlyRates(20, 15)
	db.Rate.SetUpDiscount(0, 20)

	// init session storage
	sessionStorage := sessionstorer.NewSessionStorer(time.Minute * 5)

	// serve static files
	files := http.FileServer(http.Dir("./web"))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// set up views
	views := views.NewViews("./api/views/html", sessionStorage, db.Rate, db.Order, paymenter)
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

	// set up account handlers
	accounts := personal.NewPersonalDataApi(db.User, sessionStorage)
	signupResultRedirector := middlewares.BuildRedirectOnApiCallResult(api.DefaultEndpoints.LoginPage, api.DefaultEndpoints.SignUpPage)
	http.HandleFunc(api.DefaultEndpoints.CreateAccountHandler, signupResultRedirector(accounts.CreateAccount))

	// set up auth functions
	auth := auth.NewAuthHandlers(sessionStorage, db.User)

	lougoutResultRedirector := middlewares.BuildRedirectOnApiCallResult(
		api.DefaultEndpoints.LoginPage,
		api.DefaultEndpoints.OrderPage,
	)
	http.HandleFunc(api.DefaultEndpoints.LogoutHandler, lougoutResultRedirector(auth.LogoutHandler))

	authenticationResultRedirector := middlewares.BuildRedirectOnApiCallResult(
		api.DefaultEndpoints.OrderPage,
		api.DefaultEndpoints.LoginPage,
	)
	http.HandleFunc(api.DefaultEndpoints.AuthenticateHandler, authenticationResultRedirector(auth.LoginHandler))

	// set up parkings api
	parkingsApi := parkings.NewParkingsApi(db.Parking)
	http.Handle("/parkings/", http.StripPrefix("/parkings", parkingsApi))

	// set up rates api
	ratesApi := rates.NewRatesHandlers(db.Rate, sessionStorage)
	ratesApiRedirector := middlewares.BuildRedirectOnApiCallResult(api.DefaultEndpoints.SeePricings, api.DefaultEndpoints.SeePricings)
	http.HandleFunc(api.DefaultEndpoints.ChangeActiveHoursDiscountHandler, ratesApiRedirector(ratesApi.UpdateActiveHoursDiscount))
	http.HandleFunc(api.DefaultEndpoints.ChangeSluggishHoursDiscountHandler, ratesApiRedirector(ratesApi.UpdateSluggishHoursDiscount))
	http.HandleFunc(api.DefaultEndpoints.ChangeAdultBaseRateHandler, ratesApiRedirector(ratesApi.UpdateAdultHourlyRate))
	http.HandleFunc(api.DefaultEndpoints.ChangeRetireeBaseRateHandler, ratesApiRedirector(ratesApi.UpdateRetireeHourlyRate))

	// set up orders api
	orders := orders.NewOrdersApi(db.Order, db.Rate, sessionStorage, paymenter, db.Parking)
	http.Handle(api.DefaultEndpoints.CreateOrderHandler, orders.MakeOrder)

	fmt.Println("Server started on http://localhost:8880")
	fmt.Println(http.ListenAndServe(":8880", nil))
}
