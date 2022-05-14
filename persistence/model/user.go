package model

import "gorm.io/gorm"

type Role int

const (
	Admin Role = iota
	Customer
)

type User struct {
	gorm.Model
	Username     string `gorm:"unique"`
	Email        string `gorm:"unique"`
	Age          uint
	Role         Role
	PasswordHash string
}

// IsAdmin - returns true if the user admin, false otherwise
func (user *User) IsAdmin() bool {
	return user.Role == Admin
}

func (db Database) GetUserByID(id uint) (*User, error) {
	user := User{}
	return &user, db.db.First(&user, id).Error
}

func (wrapper *Database) GetAllUsers() ([]User, error) {
	users := []User{}
	return users, wrapper.db.Find(&users).Error
}
func (db *Database) CreateNewUser(username, email, passwordHash string, age uint, role Role) (*User, error) {
	user := &User{
		Username:     username,
		Email:        email,
		Age:          age,
		PasswordHash: passwordHash,
		Role:         role,
	}

	return user, db.db.Create(user).Error
}

func (db *Database) UpdateUserAge(id, newAge uint) error {
	user, err := db.GetUserByID(id)
	if err != nil {
		return err
	}

	return db.db.Model(&user).Updates(User{Age: newAge}).Error
}

func (db *Database) UpdateUserEmail(id uint, newEmail string) error {
	user, err := db.GetUserByID(id)
	if err != nil {
		return err
	}

	return db.db.Model(&user).Updates(User{Email: newEmail}).Error
}

func (db *Database) RemoveUserByID(id uint) error {
	return db.db.Delete(User{}, id).Error
}

func (db *Database) GetUserByName(username string) (*User, error) {
	user := User{}
	result := db.db.Where("Username = ?", username).First(&user)

	return &user, result.Error
}
