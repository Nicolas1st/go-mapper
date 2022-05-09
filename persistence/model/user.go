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
	Role         Role
	PasswordHash string
}
