package validation

import "time"

type CreateArticle struct {
	Title       string `form:"title" validate:"required"`
	Description string `form:"description" validate:"required"`
	Image       string `form:"image"`
	UserID      uint   `form:"user_id"`
}

type UpdateArticle struct {
	Title       string `form:"title"`
	Description string `form:"description"`
	Image       string `form:"image"`
	UserID      uint   `form:"user_id"`
	UpdatedAt   time.Time
}
