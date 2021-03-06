package user

import (
	"gorm.io/gorm"
)

type UserDB struct {
	conn *gorm.DB
}

func NewUserDB(conn *gorm.DB) *UserDB {
	return &UserDB{
		conn: conn,
	}
}

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

func (db UserDB) GetUserByID(id uint) (*User, error) {
	user := User{}
	return &user, db.conn.First(&user, id).Error
}

func (wrapper *UserDB) GetAllUsers() ([]User, error) {
	users := []User{}
	return users, wrapper.conn.Find(&users).Error
}
func (db *UserDB) CreateNewUser(username, email, password string, age uint) (*User, error) {
	passwordHash, err := HashPassword(password)
	if err != nil {
		return &User{}, err
	}

	user := &User{
		Username:     username,
		Email:        email,
		Age:          age,
		PasswordHash: string(passwordHash),
		Role:         Customer,
	}

	return user, db.conn.Create(user).Error
}

func (db *UserDB) RemoveUserByID(id uint) error {
	return db.conn.Delete(User{}, id).Error
}

func (db *UserDB) GetUserByName(username string) (*User, error) {
	user := User{}
	result := db.conn.Where("Username = ?", username).First(&user)

	return &user, result.Error
}
