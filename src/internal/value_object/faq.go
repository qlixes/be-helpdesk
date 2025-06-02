package valueobject

import (
	"github.com/qlixes/helpdesk/internal/domain"
	"gorm.io/gorm"
)

type Faq struct {
	gorm.Model
	ID     uint
	UserID uint
	Users  domain.User
	Title  string
	Body   string
}
