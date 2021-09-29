package dto

type UserUpdateDTO struct {
	Id       uint64 `json:"id" form:"id" binding:"required"`

	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min:6"`
	Status   uint8  `json:"status,string,omitempty" form:"status,omitempty"`
}

type UserCreateDTO struct {
	
	Username string `json:"username" form:"username" binding:"required"`
	Password string `json:"password,omitempty" form:"password,omitempty" binding:"required" validate:"min:6"`
	Status   uint8  `json:"status,string,omitempty" form:"status,omitempty"`
}
