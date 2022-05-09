package templates

type Templates struct {
	MakeOrderPage     *makeOrderPage
	SignInPage        *signInPage
	SignUpPage        *signUpPage
	AddParkingPage    *addParkingPage
	RemoveParkingPage *removeParkingPage
}

func NewTemplates(pathToTemplates string) *Templates {
	return &Templates{
		MakeOrderPage:     NewMakeOrderPage(pathToTemplates),
		SignInPage:        NewSignInPage(pathToTemplates),
		SignUpPage:        NewSignUpPage(pathToTemplates),
		AddParkingPage:    NewAddParkingPage(pathToTemplates),
		RemoveParkingPage: NewRemoveParkingPage(pathToTemplates),
	}
}
