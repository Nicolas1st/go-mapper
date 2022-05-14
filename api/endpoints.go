package api

type Endpoints struct {
	Root              string
	LoginPage         string
	SignUpPage        string
	OrderPage         string
	ProfilePage       string
	AddParkingPage    string
	RemoveParkingPage string
	OrdersPage        string

	AuthenticateHandler  string
	LogoutHandler        string
	CreateAccountHandler string

	SeePricings                     string
	ChangeActiveHoursDiscountPage   string
	ChangeSluggishHoursDiscountPage string
	ChangeAdultBaseRatePage         string
	ChangeRetireeBaseRatePage       string

	ChangeActiveHoursDiscountHandler   string
	ChangeAdultBaseRateHandler         string
	ChangeRetireeBaseRateHandler       string
	ChangeSluggishHoursDiscountHandler string

	CreateOrderHandler string
	PaymentPage        string
}

var DefaultEndpoints = Endpoints{
	// html pages
	Root:              "/",
	LoginPage:         "/signin",
	SignUpPage:        "/signup",
	OrderPage:         "/order",
	ProfilePage:       "/profile",
	AddParkingPage:    "/parkings/creation",
	RemoveParkingPage: "/parkings/removal",
	OrdersPage:        "/orders",

	// authentiation handlers
	AuthenticateHandler:  "/auth/authenticate",
	LogoutHandler:        "/auth/logout",
	CreateAccountHandler: "/account/creation",

	// pricings pages
	SeePricings:                     "/admin/pricings",
	ChangeActiveHoursDiscountPage:   "/admin/discount/active-hours",
	ChangeSluggishHoursDiscountPage: "/admin/discount/sluggish-hours",
	ChangeAdultBaseRatePage:         "/admin/rates/adult",
	ChangeRetireeBaseRatePage:       "/admin/rates/retiree",

	// pricing handlers
	ChangeActiveHoursDiscountHandler:   "/admin/discount/active-hours/change",
	ChangeSluggishHoursDiscountHandler: "/admin/discount/slugish-hours/change",
	ChangeAdultBaseRateHandler:         "/admin/rates/adult/change",
	ChangeRetireeBaseRateHandler:       "/admin/rates/retiree/change",

	// apis
	CreateOrderHandler: "/orders/creation",
	PaymentPage:        "/payments",
}
