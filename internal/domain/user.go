package domain

import (
	"context"

	"github.com/qlixes/helpdesk/internal/domain"
	"github.com/qlixes/helpdesk/internal/infrastructure/dto"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint
	Email    string
	Password string
	Name     string
	RoleID   uint
	Role     *Role
}

type UserRepository interface {
	Create(ctx context.Context, data *dto.UserDto) (*domain.User, error)
	Read(ctx context.Context, data *dto.UserDto) (*domain.User, error)
	Update(ctx context.Context, data *dto.UserDto) (*domain.User, error)
	Delete(ctx context.Context, data *dto.UserDto) error
}
