package repository

import (
	"github.com/iamalone98/todo-app/internal/db"
	"github.com/iamalone98/todo-app/models"
)

type UserRepository interface {
	Create(user models.UserAuth) (*models.UserPublic, error)
	Get(login string) (*models.User, error)
	GetById(id int) (*models.User, error)
}

type userRepository struct {
	db *db.Storage
}

func NewUserRepository(db *db.Storage) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u userRepository) Create(user models.UserAuth) (*models.UserPublic, error) {
	userRet := models.UserPublic{}

	if err := u.db.DB.QueryRow("INSERT INTO users (login, password, created_at) VALUES ($1, $2, current_timestamp) RETURNING user_id, login, created_at", user.Login, user.Password).Scan(&userRet.Id, &userRet.Login, &userRet.CreatedAt); err != nil {
		return nil, err
	}

	return &userRet, nil
}

func (u userRepository) Get(login string) (*models.User, error) {
	user := models.User{}

	if err := u.db.DB.QueryRow("SELECT * FROM users WHERE login = $1", login).Scan(&user.Id, &user.Login, &user.Password, &user.CreatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u userRepository) GetById(id int) (*models.User, error) {
	user := models.User{}

	if err := u.db.DB.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.Id, &user.Login, &user.Password, &user.CreatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}
