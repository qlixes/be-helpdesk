package entity

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `json:"email"`
	Name      string `json:"name"`
	Password  string
	RoleID    uint `json:"role_id"`
	Role Role
	IsActive  bool `json:"is_active"`
	CreatedAt time.Time
}