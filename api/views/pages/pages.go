package pages

type PublicPages struct {
	SignIn *Page
	SignUp *Page
}

type PrivatePages struct {
	MakeOrder *Page
}

type AdminPages struct {
	MakeOrder     *Page
	AddParking    *Page
	RemoveParking *Page
}

type Pages struct {
	Public  *PublicPages
	Private *PrivatePages
	Admin   *AdminPages
}

func NewPages(pathToTemplates string) *Pages {
	public := &PublicPages{
		SignIn: BuildPage("SignIn", pathToTemplates, "layout", "layout.html", "sign-in-form.html", "not-signed-in-navbar.html", "footer.html"),
		SignUp: BuildPage("SignUp", pathToTemplates, "layout", "layout.html", "sign-up-form.html", "not-signed-in-navbar.html", "footer.html"),
	}

	private := &PrivatePages{
		MakeOrder: BuildPage("PrivateMakeOrder", pathToTemplates, "layout", "layout.html", "make-order-form.html", "signed-in-navbar.html", "footer.html"),
	}

	admin := &AdminPages{
		MakeOrder:     BuildPage("AdminMakeOrder", pathToTemplates, "layout", "layout.html", "make-order-form.html", "admin-navbar.html", "footer.html"),
		AddParking:    BuildPage("AdminAddParking", pathToTemplates, "layout", "layout.html", "admin/parkings/add-parking-place.html", "admin-navbar.html", "footer.html"),
		RemoveParking: BuildPage("AdminRemoveParking", pathToTemplates, "layout", "layout.html", "admin/parkings/remove-parking-place.html", "admin-navbar.html", "footer.html"),
	}

	return &Pages{
		Public:  public,
		Private: private,
		Admin:   admin,
	}
}
