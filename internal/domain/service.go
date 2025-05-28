package domain

import "gorm.io/gorm"

type Service struct {
	gorm.Model
	ID   uint
	Name string
}
