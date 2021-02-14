package handler

import (
	"github.com/midepeter/train-ticket/repository"
	"gorm.io/gorm"
)

type Handler struct {
	db   *gorm.DB
	Repo *repository.Repository
}

func NewHandler(db *gorm.DB) *Handler {
<<<<<<< HEAD
	return &Handler{db: db}
}
=======
	repo := repository.Repo(db)
	return &Handler{db: db, Repo: repo}
}
>>>>>>> c3d406c1586667a68bca168dd8d85b5bc4e01b07
