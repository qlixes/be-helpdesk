package entity

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID   uint
	Name string
}
