package repository

import (
	s "github.com/iamalone98/todo-app/internal/db"
	"github.com/iamalone98/todo-app/models"
)

type TodoRepository interface {
	Get(id int) (*models.Todo, error)
	Create(todo models.Todo)
	Update(todo models.Todo)
	Delete(id int)
}

type todoRepository struct {
	db *s.Storage
}

func NewTodoRepository(db *s.Storage) TodoRepository {
	return todoRepository{
		db: db,
	}
}

func (t todoRepository) Get(id int) (*models.Todo, error) {
	todo := &models.Todo{}
	err := t.db.DB.Get(todo, "SELECT * FROM todos WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (t todoRepository) Create(todo models.Todo) {}

func (t todoRepository) Update(todo models.Todo) {}

func (t todoRepository) Delete(id int) {}
