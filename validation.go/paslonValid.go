package validation

type CreatePaslon struct {
	Name     string   `form:"name" validate:"required"`
	Image    string   `form:"image"`
	Visimisi []string `form:"visimisi"`
}

type UpdatePaslon struct {
	Name     string   `form:"name"`
	Image    string   `form:"image"`
	Visimisi []string `form:"visimisi"`
}
