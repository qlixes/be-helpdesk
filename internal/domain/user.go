package domain

type User struct {
	ID       uint
	Email    string
	Password string
	Name     string
	RoleID   uint
	Role     *Role
}

type UserRepository interface {
}
