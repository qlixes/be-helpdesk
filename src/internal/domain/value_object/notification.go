package valueobject

import (
	"github.com/qlixes/be-helpdesk/internal/domain/entity"
	"gorm.io/gorm"
)

type Notification struct {
	gorm.Model
	ID       uint
	UserID   uint
	User     entity.User
	Path     string
	Data     string
	Response string
}
