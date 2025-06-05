package valueobject

import (
	"github.com/qlixes/be-helpdesk/internal/domain/entity"
	"gorm.io/gorm"
)

type Faq struct {
	gorm.Model
	ID     uint
	UserID uint
	Users  entity.User
	Title  string
	Body   string
}
