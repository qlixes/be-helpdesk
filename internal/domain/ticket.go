package domain

import "gorm.io/gorm"

type TicketGroup int

const (
	General TicketGroup = iota
	Technical
	Finance
)

type TicketStatus int

const (
	Pending TicketStatus = iota
	Solved
)

type Ticket struct {
	gorm.Model
	ID        uint
	CompanyID uint
	Companies Company
	UserID    uint
	Users     User
	Group     TicketGroup
	Title     string
	Status    TicketStatus
}
