package domain

import "gorm.io/gorm"

type Acl struct {
	gorm.Model
	ID        uint
	RoleID    uint
	Roles     Role
	ServiceID uint
	Services  Service
}
