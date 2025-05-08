package entity

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name     string `json:"role"      gorm:""`
	IsActive bool   `json:"is_active" gorm:""`
}
