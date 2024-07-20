package models

import "time"

type User struct {
	Id        *int       `json:"id" db:"id"`
	Login     *string    `json:"login" db:"login"`
	Password  *string    `json:"password" db:"password"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
}

type UserPublic struct {
	Id    *int    `json:"id" db:"id"`
	Login *string `json:"login" db:"login"`
}

type UserAuth struct {
	Login    *string `json:"login" db:"login" binding:"required"`
	Password *string `json:"password" db:"password" binding:"required"`
}
