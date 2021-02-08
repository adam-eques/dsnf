package repository

import (
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func Repo(db *gorm.DB) *Repository {
	return &Repository{db: db}
}
