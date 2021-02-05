package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID `json:"id"`
	Username         string    `json:"username"`
	Password         string    `json:"password"`
	Firstname        string    `json:"first_name`
	Lastname         string    `json:"last_name"`
}
