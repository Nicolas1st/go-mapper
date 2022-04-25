package templates

type Templates struct {
	MakeOrderPage *makeOrderPage
	SignInPage    *signInPage
	SignUpPage    *signUpPage
}

func NewTemplates(pathToTemplates string) *Templates {
	return &Templates{
		MakeOrderPage: NewMakeOrderPage(pathToTemplates),
		SignInPage:    NewSignInPage(pathToTemplates),
		SignUpPage:    NewSignUpPage(pathToTemplates),
	}
}
