package service

import "be-pemilu/models"

type UserService interface {
	Result() ([]UserAPI, error)
	FindOne(id uint) (UserOneAPI, error)
	Register(user models.User) (models.User, error)
	UserUpdate(id uint, user models.User) (models.User, error)
	Vote(id uint, user models.User) (models.User, error)
	DeleteUser(id uint) error
}

func (r *repository) Vote(id uint, user models.User) (models.User, error) {
	err := r.db.Model(&user).Where("id = ?", id).Updates(user).Error

	return user, err
}

func (r *repository) Result() ([]UserAPI, error) {
	var users []UserAPI
	err := r.db.Model(&models.User{}).Find(&users).Error

	return users, err
}

func (r *repository) FindOne(id uint) (UserOneAPI, error) {
	var user UserOneAPI
	err := r.db.Model(&models.User{}).First(&user, id).Error

	return user, err
}

func (r *repository) Register(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) UserUpdate(id uint, user models.User) (models.User, error) {
	err := r.db.Model(&user).Where("id = ?", id).Updates(user).Error

	return user, err
}

func (r *repository) DeleteUser(id uint) error {
	err := r.db.Delete(id).Error

	return err
}

type UserAPI struct {
	ID       uint   `json:"id"`
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
	Gender   string `json:"gender"`
	PaslonID uint   `json:"paslon_id"`
}

type UserOneAPI struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
	Gender   string `json:"gender"`
	PaslonID uint   `json:"paslon_id"`
}
