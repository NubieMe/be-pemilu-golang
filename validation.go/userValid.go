package validation

type RegisterUser struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Address  string `json:"address" form:"address"`
	Gender   string `json:"gender" form:"gender"`
}

type UpdateUser struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Fullname string `json:"fullname" form:"fullname"`
	Address  string `json:"address" form:"address"`
	Gender   string `json:"gender" form:"gender"`
}

type Vote struct {
	PaslonID uint `json:"paslon_id" form:"paslon_id"`
}

type UserResponse struct {
	Fullname string `json:"fullname" form:"fullname" validate:"required"`
	Address  string `json:"address" form:"address" validate:"optional"`
	Gender   string `json:"gender" form:"gender" validate:"optional"`
}

type LoginUser struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
