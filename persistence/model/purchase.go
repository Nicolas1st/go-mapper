package model

import "gorm.io/gorm"

type Purchase struct {
	gorm.Model
	Amount            int
	UserID            int
	User              User
	SlotReservationID int
	SlotReservation   SlotReservation
}
