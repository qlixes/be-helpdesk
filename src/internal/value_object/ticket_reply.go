package valueobject

import (
	"github.com/qlixes/helpdesk/internal/domain"
	"gorm.io/gorm"
)

type Reply struct {
	gorm.Model
	ID       uint
	TicketID uint
	Tickets  Ticket
	UserID   uint
	Users    domain.User
	Body     string
}
