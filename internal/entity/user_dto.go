package entity

type SigninDTO struct {
	Email string `validate:"required"`
	Password string `json:"required"`
}

type SignupDTO struct {
	Email string `validate:"require"`
	Name string `validate:"required"`
	Password string `validate:"required"`
	PasswordConfirmation string `validate:"required"`
	RoleID uint `validate:"required"`
	Role Role `validate:"required"`
}

type ForgetDTO struct {
	Email string `validate:"required"`
}