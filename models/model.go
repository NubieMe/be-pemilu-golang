package models

import (
	"time"

	"github.com/lib/pq"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"type: varchar(100)"`
	Password string `json:"password" gorm:"type: varchar(255)"`
	Fullname string `json:"fullname" gorm:"type: varchar(100)"`
	Address  string `json:"address" gorm:"type: varchar(100)"`
	Gender   string `json:"gender" gorm:"type: varchar(10)"`
	Article  []Article
	PaslonID uint `json:"paslon_id" gorm:"default:NULL"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Article struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"type: varchar(100)"`
	Description string `json:"description"`
	Image       string `json:"image" gorm:"type: varchar(255)"`
	UserID      uint   `json:"user_id" gorm:"default:NULL"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Party struct {
	ID       uint           `json:"id" gorm:"primaryKey"`
	Name     string         `json:"name" gorm:"type: varchar(100)"`
	Leader   string         `json:"leader" gorm:"type: varchar(100)"`
	Image    string         `json:"image" gorm:"type: varchar(255)"`
	Visimisi pq.StringArray `json:"visimisi" gorm:"type: text[]"`
	PaslonID uint           `json:"paslon_id" gorm:"default:NULL"`
}

type Paslon struct {
	ID        uint           `json:"id" gorm:"primarKey"`
	Name      string         `json:"name" gorm:"type: varchar(100)"`
	Image     string         `json:"image" gorm:"type: varchar(255)"`
	Visimisi  pq.StringArray `json:"visimisi" gorm:"type: text[]"`
	Coalition []Party
	Voter     []User
}
