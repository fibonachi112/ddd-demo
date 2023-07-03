package drivers

import (
	"context"
	"repo/config"
	"repo/internal/DI"
	"repo/internal/domain/entity"
)

type MockUserRepository struct {
	Config *config.Configuration
	Di     *DI.ServiceContainer
	Ctx    *context.Context
}

func (r *MockUserRepository) GetOne(id int) (entity.User, error) {
	//TODO implement me
	user, err := r.Di.InMemoryStorage.GetOneUser(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *MockUserRepository) SaveOne(user *entity.User) error {
	//TODO implement me
	newUser, err := r.Di.InMemoryStorage.AddUser(*user)
	if err != nil {
		return err
	}
	user.Id = newUser.Id
	user.CreatedAt = newUser.CreatedAt

	return nil
}

func (r *MockUserRepository) UpdateOne(user *entity.User) error {
	updatedUser, err := r.Di.InMemoryStorage.UpdateUser(*user)
	if err != nil {
		return err
	}
	user = &updatedUser
	return nil
}

func (r *MockUserRepository) DeleteOne(user *entity.User) error {
	err := r.Di.InMemoryStorage.RemoveUser(*user)
	if err != nil {
		return err
	}
	return nil
}

func (r *MockUserRepository) SetContext(ctx *context.Context)  {
	r.Ctx = ctx
}

func MakeMockDriver(config *config.Configuration, di *DI.ServiceContainer) *MockUserRepository {
	return &MockUserRepository{
		Config: config,
		Di:     di,
	}
}
