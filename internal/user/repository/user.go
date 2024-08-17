package repository

import (
	"github.com/evertonbzr/library_micro/internal/user/model"
	"github.com/evertonbzr/library_micro/pkg/infra/db"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) GetUserById(id int) (user *model.User, err error) {
	user = &model.User{}
	err = db.GetDB().First(user, id).Error
	return user, err
}

func (u *UserRepository) GetUserByEmail(email string) (user *model.User, err error) {
	user = &model.User{}
	err = db.GetDB().Where("email = ?", email).First(user).Error
	return user, err
}
func (u *UserRepository) CreateUser(user *model.User) (err error) {
	err = db.GetDB().Create(user).Error
	return err
}
