package domain

type Acl struct {
	ID        uint
	RoleID    uint
	Roles     Role
	ServiceID uint
	Services  Service
}
