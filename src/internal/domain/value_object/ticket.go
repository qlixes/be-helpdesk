package valueobject

import (
	"github.com/qlixes/be-helpdesk/internal/entity"
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
	Companies entity.Company
	UserID    uint
	Users     entity.User
	Group     TicketGroup
	Title     string
	Status    TicketStatus
}
