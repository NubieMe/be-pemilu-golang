package service

import (
	"be-pemilu/models"
	"time"
)

type ArticleService interface {
	ArtAll() ([]ArticleAPI, error)
	ArtOne(id uint) (ArticleOneAPI, error)
	CreArt(artic models.Article) (models.Article, error)
	UpdArt(id uint, artic UpdateArticle) (UpdateArticle, error)
	DeleteArt(id uint) error
}

func (r *repository) ArtAll() ([]ArticleAPI, error) {
	var artics []ArticleAPI
	err := r.db.Model(&models.Article{}).Find(&artics).Error

	return artics, err
}

func (r *repository) ArtOne(id uint) (ArticleOneAPI, error) {
	var artic ArticleOneAPI
	err := r.db.Model(&models.Article{}).First(&artic).Error

	return artic, err
}

func (r *repository) CreArt(artic models.Article) (models.Article, error) {
	err := r.db.Create(&artic).Error

	return artic, err
}

func (r *repository) UpdArt(id uint, artic UpdateArticle) (UpdateArticle, error) {
	err := r.db.Model(&models.Article{}).Where(id).Updates(&artic).Error

	return artic, err
}

func (r *repository) DeleteArt(id uint) error {
	err := r.db.Delete(id).Error

	return err
}

type ArticleAPI struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Image     string `json:"image"`
	CreatedAt string `json:"created_at"`
}

type ArticleOneAPI struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
	UserID      uint   `json:"user_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type UpdateArticle struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	UserID      uint      `json:"user_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}
