package repository

import "github.com/rht6226/event-management-app/model"

type UserRepository interface {
	Save(*model.User) (*model.User, error)
	FindByID(id uint) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindAll() ([]*model.User, error)
	Update(user *model.User) (*model.User, error)
	Delete(id uint) error
}
