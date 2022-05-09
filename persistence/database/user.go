package database

import "yaroslavl-parkings/persistence/model"

func (db Database) GetUserByID(id uint) (*model.User, error) {
	user := model.User{}
	return &user, db.db.First(&user, id).Error
}

func (wrapper *Database) GetAllUsers() ([]model.User, error) {
	users := []model.User{}
	return users, wrapper.db.Find(&users).Error
}
func (db *Database) CreateNewUser(username, passwordHash string, role model.Role) (*model.User, error) {
	user := &model.User{
		Username:     username,
		PasswordHash: passwordHash,
		Role:         role,
	}

	return user, db.db.Create(user).Error
}

func (db *Database) RemoveUserByID(id uint) error {
	return db.db.Delete(model.User{}, id).Error
}

func (db *Database) GetUserByName(username string) (*model.User, error) {
	user := model.User{}
	result := db.db.Where("Username = ?", username).First(&user)

	return &user, result.Error
}
