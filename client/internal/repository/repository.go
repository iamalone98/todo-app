package repository

import "github.com/iamalone98/todo-app/internal/db"

type Repository struct {
	TodoRepository TodoRepository
	UserRepository UserRepository
}

func NewRepository(db *db.Storage) Repository {
	return Repository{
		TodoRepository: NewTodoRepository(db),
		UserRepository: NewUserRepository(db),
	}
}
