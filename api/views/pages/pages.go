package pages

type PublicPages struct {
	SignIn *Page
	SignUp *Page
}

type PrivatePages struct {
	MakeOrder *Page
	Profile   *Page
}

type AdminPages struct {
	MakeOrder                   *Page
	AddParking                  *Page
	RemoveParking               *Page
	Profile                     *Page
	ChangeAdultBaseRate         *Page
	ChangeRetireeBaseRate       *Page
	ChangeActiveHoursDiscount   *Page
	ChangeSluggishHoursDiscount *Page
	SeePricing                  *Page
}

type Pages struct {
	Public  *PublicPages
	Private *PrivatePages
	Admin   *AdminPages
}

func NewPages(pathToTemplates string) *Pages {
	mainTemplateName := "layout"
	commonFiles := []string{"layout.html", "footer.html"}

	commonPublicFiles := append(commonFiles, "not-signed-in-navbar.html")
	public := &PublicPages{
		SignIn: BuildPage("SignIn", pathToTemplates, mainTemplateName, append(commonPublicFiles, "sign-in-form.html")...),
		SignUp: BuildPage("SignUp", pathToTemplates, mainTemplateName, append(commonPublicFiles, "sign-up-form.html")...),
	}

	commonPrivateFiles := append(commonFiles, "signed-in-navbar.html")
	private := &PrivatePages{
		MakeOrder: BuildPage("PrivateMakeOrder", pathToTemplates, mainTemplateName, append(commonPrivateFiles, "make-order-form.html")...),
		Profile:   BuildPage("ProfilePage", pathToTemplates, mainTemplateName, append(commonPrivateFiles, "profile.html")...),
	}

	commonAdminFiles := append(commonFiles, "admin-navbar.html")
	admin := &AdminPages{
		MakeOrder:                   BuildPage("AdminMakeOrder", pathToTemplates, mainTemplateName, append(commonAdminFiles, "make-order-form.html")...),
		AddParking:                  BuildPage("AdminAddParking", pathToTemplates, mainTemplateName, append(commonAdminFiles, "admin/parkings/add-parking-place.html")...),
		RemoveParking:               BuildPage("AdminRemoveParking", pathToTemplates, mainTemplateName, append(commonAdminFiles, "admin/parkings/remove-parking-place.html")...),
		Profile:                     BuildPage("ProfilePage", pathToTemplates, mainTemplateName, append(commonAdminFiles, "profile.html")...),
		ChangeAdultBaseRate:         BuildPage("AdultBaseRate", pathToTemplates, mainTemplateName, append(commonAdminFiles, "change-adult-base-rate.html")...),
		ChangeRetireeBaseRate:       BuildPage("RetireeBaseRate", pathToTemplates, mainTemplateName, append(commonAdminFiles, "change-retiree-base-rate.html")...),
		ChangeActiveHoursDiscount:   BuildPage("ActiveHoursDiscount", pathToTemplates, mainTemplateName, append(commonAdminFiles, "change-active-hours-discount.html")...),
		ChangeSluggishHoursDiscount: BuildPage("SluggishHoursDiscount", pathToTemplates, mainTemplateName, append(commonAdminFiles, "change-sluggish-hours-discount.html")...),
		SeePricing:                  BuildPage("SeePricings", pathToTemplates, mainTemplateName, append(commonAdminFiles, "see-pricings.html")...),
	}

	return &Pages{
		Public:  public,
		Private: private,
		Admin:   admin,
	}
}
