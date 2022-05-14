package user

import (
	"gorm.io/gorm/clause"
)

func (db *UserDB) SetUpAdminAccount(
	username,
	email,
	password string,
	age uint,
) error {
	passwordHash, err := hashPassword(password)
	if err != nil {
		return err
	}

	admin := User{
		Username:     username,
		Email:        email,
		Age:          age,
		PasswordHash: string(passwordHash),
		Role:         Admin,
	}

	return db.conn.Clauses(clause.OnConflict{DoNothing: true}).Create(&admin).Error
}
