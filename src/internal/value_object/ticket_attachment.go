package valueobject

import (
	"gorm.io/gorm"
)

type TicketAttachment struct {
	gorm.Model
	ID       uint
	ReplyID  uint
	Replies  Reply
	Filename string
}
