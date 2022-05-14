package rate

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Rate categories dependent on the time of the day
type PeriodCategory uint

const (
	ActiveHours PeriodCategory = iota
	SlugishHours
)

// Discount Based on the time of the day
type Discount uint

// PeriodDisconts - dicounts should be applied
// based on the time of the day
type PeriodDiscount struct {
	gorm.Model
	PeriodCategory    PeriodCategory `gorm:"unique"`
	DisountInPercents Discount
}

func (db *RateDB) SetUpDiscount(activeHoursDiscount, sluggishHoursDiscount uint) error {
	if activeHoursDiscount > 100 || sluggishHoursDiscount > 100 {
		panic("the discount can not be bigger than 100%")
	}

	activeHours := PeriodDiscount{
		PeriodCategory:    ActiveHours,
		DisountInPercents: Discount(activeHoursDiscount),
	}
	if err := db.conn.Clauses(clause.OnConflict{DoNothing: true}).Create(&activeHours).Error; err != nil {
		return err
	}

	sluggishHours := PeriodDiscount{
		PeriodCategory:    SlugishHours,
		DisountInPercents: Discount(sluggishHoursDiscount),
	}
	if err := db.conn.Clauses(clause.OnConflict{DoNothing: true}).Create(&sluggishHours).Error; err != nil {
		return err
	}

	return nil
}
