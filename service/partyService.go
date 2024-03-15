package service

import (
	"be-pemilu/models"

	"github.com/lib/pq"
)

type PartyService interface {
	PartyAll() ([]models.Party, error)
	PartyOne(id uint) (models.Party, error)
	CrePart(part models.Party) (models.Party, error)
	UpdPart(id uint, part models.Party) (models.Party, error)
	DeletePart(id uint) error
}

func (r *repository) PartyAll() ([]models.Party, error) {
	var parties []models.Party
	err := r.db.Model(&models.Party{}).Find(&parties).Error

	return parties, err
}

func (r *repository) PartyOne(id uint) (models.Party, error) {
	var party models.Party
	err := r.db.Model(&models.Party{}).First(&party, id).Error

	return party, err
}

func (r *repository) CrePart(party models.Party) (models.Party, error) {
	err := r.db.Create(&party).Error

	return party, err
}

func (r *repository) UpdPart(id uint, part models.Party) (models.Party, error) {
	err := r.db.Model(&models.Party{}).Where(id).Updates(&part).Error

	return part, err
}

func (r *repository) DeletePart(id uint) error {
	err := r.db.Delete(id).Error

	return err
}

type PartyAPI struct {
	ID       uint           `json:"id"`
	Name     string         `json:"name"`
	Leader   string         `json:"leader"`
	Image    string         `json:"image"`
	Visimisi pq.StringArray `json:"visimisi"`
	PaslonID uint           `json:"paslon_id"`
}
