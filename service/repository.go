package service

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

func Repo2DB(db *gorm.DB) *repository {
	return &repository{db}
}
