package domain

type Faq struct {
	ID     uint
	UserID uint
	Users  User
	Title  string
	Body   string
}
