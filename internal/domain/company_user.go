package domain

type CompanyUser struct {
	ID        uint
	CompanyID uint
	Companies Company
	UserID    uint
	Users     User
	IsActive  bool
}
