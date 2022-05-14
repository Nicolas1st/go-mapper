package rate

import (
	"gorm.io/gorm"
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
	PeriodCategory     PeriodCategory `gorm:"unique"`
	DiscountInPercents Discount
}

func (db *RateDB) SetUpDiscount(activeHoursDiscount, sluggishHoursDiscount uint) error {
	if activeHoursDiscount > 100 || sluggishHoursDiscount > 100 {
		panic("the discount can not be bigger than 100%")
	}

	activeHours := PeriodDiscount{
		PeriodCategory:     ActiveHours,
		DiscountInPercents: Discount(activeHoursDiscount),
	}
	if err := db.conn.Create(&activeHours).Error; err != nil {
		return err
	}

	sluggishHours := PeriodDiscount{
		PeriodCategory:     SlugishHours,
		DiscountInPercents: Discount(sluggishHoursDiscount),
	}
	if err := db.conn.Create(&sluggishHours).Error; err != nil {
		return err
	}

	return nil
}

func (db *RateDB) GetActiveHoursDiscount() (PeriodDiscount, error) {
	var discount PeriodDiscount
	// database ids start with one
	result := db.conn.First(&discount, ActiveHours+1)

	return discount, result.Error
}

func (db *RateDB) GetSluggishHoursDiscount() (PeriodDiscount, error) {
	var discount PeriodDiscount
	// database ids start with one
	result := db.conn.First(&discount, SlugishHours+1)

	return discount, result.Error
}

func (db *RateDB) UpdateActiveHoursDiscount(discountInPercents uint) error {
	discount, err := db.GetActiveHoursDiscount()
	if err != nil {
		return err
	}

	result := db.conn.Model(&discount).Updates(PeriodDiscount{DiscountInPercents: Discount(discountInPercents)})

	return result.Error
}

func (db *RateDB) UpdateSluggishHoursDiscount(discountInPercents uint) error {
	discount, err := db.GetSluggishHoursDiscount()
	if err != nil {
		return err
	}

	result := db.conn.Model(&discount).Updates(PeriodDiscount{DiscountInPercents: Discount(discountInPercents)})

	return result.Error
}
