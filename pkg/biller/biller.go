package biller

type Biller struct{}

func NewBiller() *Biller {
	return &Biller{}
}

func (b *Biller) ComputePrice(basePrice uint, discount uint, forHowManyHours uint) uint {
	return basePrice * ((100 - discount) / 100) * forHowManyHours
}
