package rates

import "yaroslavl-parkings/data/sessionstorer"

type DatabaseInterface interface {
	UpdateAdultPricePerHour(pricePerHour uint) error
	UpdateRetireePricePerHour(pricePerHour uint) error

	UpdateActiveHoursDiscount(discountInPercents uint) error
	UpdateSluggishHoursDiscount(discountInPercents uint) error
}

type SessionsInterface interface {
	IsSessionValid(sessionToken string) (*sessionstorer.Session, bool)
}
