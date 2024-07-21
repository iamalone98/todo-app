package repository

import (
	s "github.com/iamalone98/todo-app/internal/db"
	"github.com/iamalone98/todo-app/models"
)

type TodoRepository interface {
	Get(id int, userId int) (*models.Todo, error)
	GetAll(userId int) ([]models.Todo, error)
	Create(userId int, todo models.TodoCreate) (int, error)
	Update(userId int, todo models.TodoUpdate) error
	Delete(id int, userId int) error
}

type todoRepository struct {
	db *s.Storage
}

func NewTodoRepository(db *s.Storage) TodoRepository {
	return todoRepository{
		db: db,
	}
}

func (t todoRepository) Get(id int, userId int) (*models.Todo, error) {
	todo := models.Todo{}
	err := t.db.DB.Get(&todo, `SELECT 
			todos.todo_id, 
			todos.header, 
			todos.description, 
			todos.completed, 
			todos.created_at AS todo_created_at, 
			users.user_id, 
			users.login, 
			users.created_at AS user_created_at 
			FROM todos 
			LEFT JOIN users ON todos.user_id = users.user_id 
			WHERE todos.todo_id = $1 AND todos.user_id = $2
		`, id, userId)

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func (t todoRepository) GetAll(userId int) ([]models.Todo, error) {
	todos := []models.Todo{}
	err := t.db.DB.Select(&todos, `SELECT 
			todos.todo_id, 
			todos.header, 
			todos.description, 
			todos.completed, 
			todos.created_at AS todo_created_at, 
			users.user_id, 
			users.login, 
			users.created_at AS user_created_at 
			FROM todos 
			LEFT JOIN users ON todos.user_id = users.user_id 
			WHERE todos.user_id = $1
		`, userId)

	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (t todoRepository) Create(userId int, todo models.TodoCreate) (int, error) {
	var todoId int

	if todo.Completed == nil {
		defaultCompleted := false
		todo.Completed = &defaultCompleted
	}

	if err := t.db.DB.QueryRow("INSERT INTO todos (user_id, header, description, completed, created_at) VALUES ($1, $2, $3, $4, current_timestamp) RETURNING todo_id", userId, todo.Header, todo.Description, todo.Completed).Scan(&todoId); err != nil {
		return todoId, err
	}

	return todoId, nil
}

func (t todoRepository) Update(userId int, todo models.TodoUpdate) error {
	_, err := t.db.DB.Exec("UPDATE todos SET header = COALESCE($1, header), description = COALESCE($2, description), completed = COALESCE($3, completed) WHERE todo_id = $4 AND user_id = $5", todo.Header, todo.Description, todo.Completed, todo.Id, userId)
	if err != nil {
		return err
	}

	return nil
}

func (t todoRepository) Delete(id int, userId int) error {
	_, err := t.db.DB.Exec("DELETE FROM todos WHERE todo_id = $1 AND user_id = $2", id, userId)
	if err != nil {
		return err
	}

	return nil
}
