package service

import (
	"github.com/iamalone98/todo-app/internal/repository"
	"github.com/iamalone98/todo-app/models"
)

type TodoService interface {
	Get(id int) (*models.Todo, error)
	Create(todo models.Todo)
	Update(todo models.Todo)
	Delete(id int)
}

type todoService struct {
	r repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return todoService{
		r: todoRepo,
	}
}

func (t todoService) Get(id int) (*models.Todo, error) {
	todo, err := t.r.Get(id)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (t todoService) Create(todo models.Todo) {}
func (t todoService) Update(todo models.Todo) {}
func (t todoService) Delete(id int)           {}
