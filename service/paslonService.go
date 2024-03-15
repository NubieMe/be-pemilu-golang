package service

import (
	"be-pemilu/models"

	"github.com/lib/pq"
)

type PaslonService interface {
	PaslonAll() ([]models.Paslon, error)
	PaslonOne(id uint) (models.Paslon, error)
	CrePas(pas models.Paslon) (models.Paslon, error)
	UpdPas(id uint, pas PaslonAPI) (PaslonAPI, error)
	DeletePas(id uint) error
}

func (r *repository) PaslonAll() ([]models.Paslon, error) {
	var paslons []models.Paslon
	err := r.db.Model(&models.Paslon{}).Find(&paslons).Error

	return paslons, err
}

func (r *repository) PaslonOne(id uint) (models.Paslon, error) {
	var paslon models.Paslon
	err := r.db.Model(&models.Paslon{}).First(&paslon, id).Error

	return paslon, err
}

func (r *repository) CrePas(pas models.Paslon) (models.Paslon, error) {
	err := r.db.Create(&pas).Error

	return pas, err
}

func (r *repository) UpdPas(id uint, pas PaslonAPI) (PaslonAPI, error) {
	err := r.db.Model(&models.Paslon{}).Where(id).Updates(&pas).Error

	return pas, err
}

func (r *repository) DeletePas(id uint) error {
	err := r.db.Delete(id).Error

	return err
}

type PaslonAPI struct {
	ID        uint           `json:"id" gorm:"primarKey"`
	Name      string         `json:"name" gorm:"type: varchar(100)"`
	Image     string         `json:"image" gorm:"type: varchar(255)"`
	Visimisi  pq.StringArray `json:"visimisi"`
	Coalition pq.StringArray `json:"coalition"`
}
