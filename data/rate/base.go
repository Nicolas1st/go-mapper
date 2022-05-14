package rate

import (
	"yaroslavl-parkings/data/user"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	if err := db.conn.Clauses(clause.OnConflict{DoNothing: true}).Create(&adultsRate).Error; err != nil {
		return err
	}

	retireesRate := BaseRate{
		AgeCategory: user.Retiree,
		Price:       PricePerHour(forRetirees),
	}
	if err := db.conn.Clauses(clause.OnConflict{DoNothing: true}).Create(&retireesRate).Error; err != nil {
		return err
	}

	return nil
}
