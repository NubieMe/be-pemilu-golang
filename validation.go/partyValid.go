package validation

type CreateParty struct {
	Name     string   `form:"name" validate:"required"`
	Leader   string   `form:"leader" validate:"required"`
	Image    string   `form:"image"`
	Visimisi []string `form:"visimisi"`
	PaslonID uint     `form:"paslon_id"`
}

type UpdateParty struct {
	Name     string   `form:"name"`
	Leader   string   `form:"leader"`
	Image    string   `form:"image"`
	Visimisi []string `form:"visimisi"`
	PaslonID uint     `form:"paslon_id"`
}
