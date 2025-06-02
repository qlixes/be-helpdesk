package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint
	Email    string
	Password string
	Name     string
	RoleID   uint
	Role     *Role
}

type UserRepository interface {
}
