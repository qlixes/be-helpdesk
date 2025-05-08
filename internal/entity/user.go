package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email"     gorm:""`
	Name     string `json:"name"      gorm:""`
	Password string
	RoleID   *Role `json:"role"      gorm:""`
	IsActive bool  `json:"is_active" gorm:""`
}

type UserRepository interface {
	Show() []User
	Filter(id int) User
	Search(key string) []User
	Update(id int, user User) (User, error)
	Delete(id int) error
}

var UserRepo = newUserRepository()

func newUserRepository() *User {
	return &User{}
}

func (e *User) Show() []User {

}

func (e *User) Filter(id int) User {

}

func (e *User) Search(key string) []User {

}

func (e *User) Update(id int, user User) (User, error) {

}

func (e *User) Delete(id int) error {

}
