package repository

import (
	"github.com/iamalone98/todo-app/internal/db"
	"github.com/iamalone98/todo-app/models"
)

type UserRepository interface {
	Create(user models.UserAuth) error
	Get(login string) (*models.User, error)
}

type userRepository struct {
	db *db.Storage
}

func NewUserRepository(db *db.Storage) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u userRepository) Create(user models.UserAuth) error {
	if _, err := u.db.DB.Exec("INSERT INTO users (login, password, created_at) VALUES ($1, $2, CURRENT_TIMESTAMP)", user.Login, user.Password); err != nil {
		return err
	}

	return nil
}

func (u userRepository) Get(login string) (*models.User, error) {
	user := models.User{}

	if err := u.db.DB.QueryRow("SELECT * FROM users WHERE login = $1", login).Scan(&user.Id, &user.Login, &user.Password, &user.CreatedAt); err != nil {
		return nil, err
	}

	return &user, nil
}
