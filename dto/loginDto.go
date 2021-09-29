package dto

type LoginDTO struct {
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:5"`
}
