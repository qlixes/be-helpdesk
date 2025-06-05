package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `gorm:"primary_key;auto_increment" json:"id"`
	Email    string `gorm:"not null;" json:"email"`
	Password string `gorm:"" json:"password"`
	Name     string `gorm:"" json:"name"`
	RoleID   uint   `gorm:"" json:"role_id"`
	Role     *Role
}

type UserRepository interface {
	Create() error
	Select() error
	Update() error
	Delete() error
}
