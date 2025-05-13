package entity

type RoleRequest struct {
	Name string `validate:"required"`
}