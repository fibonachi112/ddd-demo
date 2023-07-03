package UserRepository

import (
	"context"
	"repo/internal/domain/entity"
)

type IUserRepository interface {
	GetOne(id int) (entity.User, error)
	SaveOne(user *entity.User) error
	UpdateOne(user *entity.User) error
	DeleteOne(user *entity.User) error
	SetContext(ctx *context.Context)
}
