package service

import (
	"github.com/iamalone98/todo-app/internal/repository"
	"github.com/iamalone98/todo-app/internal/utils"
	"github.com/iamalone98/todo-app/models"
)

type UserService interface {
	Create(user models.UserAuth) (*models.UserPublic, error)
	Get(login string) (*models.User, error)
	GetById(id int) (*models.User, error)
}

type userService struct {
	r repository.UserRepository
}

func NewUserService(r repository.UserRepository) userService {
	return userService{
		r: r,
	}
}

func (u userService) Create(user models.UserAuth) (*models.UserPublic, error) {
	hashPass, err := utils.HashPassword([]byte(*user.Password), 14)

	if err != nil {
		return nil, err
	}

	user.Password = &hashPass

	userRet, err := u.r.Create(user)

	if err != nil {
		return nil, err
	}

	return userRet, nil
}

func (u userService) Get(login string) (*models.User, error) {
	return u.r.Get(login)
}

func (u userService) GetById(id int) (*models.User, error) {
	return u.r.GetById(id)
}
