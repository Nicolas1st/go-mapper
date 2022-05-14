package rate

import (
	"yaroslavl-parkings/data/user"

	"gorm.io/gorm"
)

// Base price
type PricePerHour uint

// Base - rates on to which
// discount can be applied
// to compute the final rate
type BaseRate struct {
	gorm.Model
	AgeCategory user.AgeCategory `gorm:"unique"`
	Price       PricePerHour
}

// SetUpAgeCategoryBasedHourlyRates - sets up prices per hour for age groups
func (db *RateDB) SetUpAgeCategoryBasedHourlyRates(forAdults, forRetirees uint) error {
	adultsRate := BaseRate{
		AgeCategory: user.Adult,
		Price:       PricePerHour(forAdults),
	}
	if err := db.conn.Create(&adultsRate).Error; err != nil {
		return err
	}

	retireesRate := BaseRate{
		AgeCategory: user.Retiree,
		Price:       PricePerHour(forRetirees),
	}
	if err := db.conn.Create(&retireesRate).Error; err != nil {
		return err
	}

	return nil
}

func (db *RateDB) GetAdultRatePerHour() (BaseRate, error) {
	var rate BaseRate
	// database ids start with one
	result := db.conn.First(&rate, user.Adult+1)

	return rate, result.Error
}

func (db *RateDB) GetRetireeRatePerHour() (BaseRate, error) {
	var rate BaseRate
	// database ids start with one
	result := db.conn.First(&rate, user.Retiree+1)

	return rate, result.Error
}

func (db *RateDB) UpdateAdultPricePerHour(pricePerHour uint) error {
	rate, err := db.GetAdultRatePerHour()
	if err != nil {
		return err
	}

	result := db.conn.Model(&rate).Updates(BaseRate{Price: PricePerHour(pricePerHour)})

	return result.Error
}

func (db *RateDB) UpdateRetireePricePerHour(pricePerHour uint) error {
	rate, err := db.GetRetireeRatePerHour()
	if err != nil {
		return err
	}

	result := db.conn.Model(&rate).Updates(BaseRate{Price: PricePerHour(pricePerHour)})

	return result.Error
}
