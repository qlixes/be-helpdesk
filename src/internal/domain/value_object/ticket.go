package valueobject

import (
	"github.com/qlixes/be-helpdesk/internal/domain"
	"gorm.io/gorm"
)

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
	Companies domain.Company
	UserID    uint
	Users     domain.User
	Group     TicketGroup
	Title     string
	Status    TicketStatus
}
