package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `json:"email"     gorm:""`
	Name      string `json:"name"      gorm:""`
	Password  string
	RoleID    Role `json:"role"      gorm:""`
	IsActive  bool `json:"is_active" gorm:""`
	CreatedAt time.Time
}

type UserRepository interface {
	Show() []User
	Filter(id int) User
	Search(key string) []User
	Update(id int, user User) (User, error)
	Delete(id int) error
}
