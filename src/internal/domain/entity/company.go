package entity

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	ID   uint
	Name string
}
