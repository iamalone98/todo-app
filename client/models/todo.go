package models

import "time"

type Todo struct {
	Id          *int       `json:"todo_id" db:"todo_id"`
	Header      *string    `json:"header" db:"header"`
	Description *string    `json:"description" db:"description"`
	Completed   *bool      `json:"completed" db:"completed"`
	CreatedAt   *time.Time `json:"created_at" db:"todo_created_at"`
	TodoUser    `json:"user"`
}

type TodoUser struct {
	Id        *int       `json:"user_id" db:"user_id"`
	Login     *string    `json:"login" db:"login"`
	CreatedAt *time.Time `json:"created_at" db:"user_created_at"`
}

type TodoCreate struct {
	Header      *string `json:"header" db:"header" binding:"required"`
	Description *string `json:"description" db:"description"`
	Completed   *bool   `json:"completed" db:"completed"`
}

type TodoUpdate struct {
	Id          *int    `json:"todo_id" db:"todo_id" binding:"required"`
	Header      *string `json:"header" db:"header"`
	Description *string `json:"description" db:"description"`
	Completed   *bool   `json:"completed" db:"completed"`
}
