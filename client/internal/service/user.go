package service

import (
	"github.com/iamalone98/todo-app/internal/repository"
	"github.com/iamalone98/todo-app/internal/utils"
	"github.com/iamalone98/todo-app/models"
)

type UserService interface {
	Create(user models.UserAuth) error
	Get(login string) (*models.User, error)
}

type userService struct {
	r repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return userService{
		r: r,
	}
}

func (u userService) Create(user models.UserAuth) error {
	hashPass, err := utils.HashPassword([]byte(*user.Password), 14)

	if err != nil {
		return err
	}

	user.Password = &hashPass

	if err := u.r.Create(user); err != nil {
		return err
	}

	return nil
}

func (u userService) Get(login string) (*models.User, error) {
	return u.r.Get(login)
}
