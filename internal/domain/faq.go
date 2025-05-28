package domain

import "gorm.io/gorm"

type Faq struct {
	gorm.Model
	ID     uint
	UserID uint
	Users  User
	Title  string
	Body   string
}
