package valueobject

import (
	"github.com/qlixes/be-helpdesk/internal/domain/entity"
	"gorm.io/gorm"
)

type Reply struct {
	gorm.Model
	ID       uint
	TicketID uint
	Tickets  Ticket
	UserID   uint
	Users    entity.User
	Body     string
}
