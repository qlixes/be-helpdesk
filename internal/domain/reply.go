package domain

import "gorm.io/gorm"

type Reply struct {
	gorm.Model
	ID       uint
	TicketID uint
	Tickets  Ticket
	UserID   uint
	Users    User
	Body     string
}
