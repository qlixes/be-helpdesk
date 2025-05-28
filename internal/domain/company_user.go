package domain

import "gorm.io/gorm"

type CompanyUser struct {
	gorm.Model
	ID        uint
	CompanyID uint
	Companies Company
	UserID    uint
	Users     User
	IsActive  bool
}
