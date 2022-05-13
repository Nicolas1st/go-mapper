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
