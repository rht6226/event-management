package repository

import (
	"github.com/rht6226/event-management-app/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

// Delete implements UserRepository.
func (r *userRepository) Delete(id uint) error {
	if err := r.db.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

// FindAll implements UserRepository.
func (r *userRepository) FindAll() ([]*model.User, error) {
	var users []*model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// FindByEmail implements UserRepository.
func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var user *model.User = &model.User{}
	if err := r.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// FindByID implements UserRepository.
func (r *userRepository) FindByID(id uint) (*model.User, error) {
	var user *model.User = &model.User{}
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Save implements UserRepository.
func (r *userRepository) Save(user *model.User) (*model.User, error) {
	if err := r.db.FirstOrCreate(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Update implements UserRepository.
func (r *userRepository) Update(user *model.User) (*model.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

var _ UserRepository = (*userRepository)(nil)
