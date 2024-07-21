package service

import (
	"github.com/iamalone98/todo-app/internal/repository"
	"github.com/iamalone98/todo-app/models"
)

type TodoService interface {
	Get(id int, userId int) (*models.Todo, error)
	GetAll(userId int) ([]models.Todo, error)
	Create(userId int, todo models.TodoCreate) (*models.Todo, error)
	Update(userId int, todo models.TodoUpdate) (*models.Todo, error)
	Delete(id int, userId int) (*models.Todo, error)
}

type todoService struct {
	r repository.TodoRepository
}

func NewTodoService(todoRepo repository.TodoRepository) TodoService {
	return todoService{
		r: todoRepo,
	}
}

func (t todoService) Get(id int, userId int) (*models.Todo, error) {
	todo, err := t.r.Get(id, userId)

	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (t todoService) GetAll(userId int) ([]models.Todo, error) {
	todos, err := t.r.GetAll(userId)

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (t todoService) Create(userId int, todo models.TodoCreate) (*models.Todo, error) {
	todoId, err := t.r.Create(userId, todo)
	if err != nil {
		return nil, err
	}

	todoRet, err := t.r.Get(todoId, userId)
	if err != nil {
		return nil, err
	}

	return todoRet, nil
}

func (t todoService) Update(userId int, todo models.TodoUpdate) (*models.Todo, error) {
	err := t.r.Update(userId, todo)
	if err != nil {
		return nil, err
	}

	todoRet, err := t.r.Get(*todo.Id, userId)
	if err != nil {
		return nil, err
	}

	return todoRet, err
}

func (t todoService) Delete(id int, userId int) (*models.Todo, error) {
	todo, err := t.Get(id, userId)
	if err != nil {
		return nil, err
	}

	err = t.r.Delete(id, userId)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
