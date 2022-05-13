package api

type Endpoints struct {
	Root              string
	LoginPage         string
	SignUpPage        string
	OrderPage         string
	ProfilePage       string
	AddParkingPage    string
	RemoveParkingPage string

	AuthenticateHandler  string
	LogoutHandler        string
	CreateAccountHandler string
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

	// authentiation handlers
	AuthenticateHandler:  "/auth/authenticate",
	LogoutHandler:        "/auth/logout",
	CreateAccountHandler: "/account/creation",
}
