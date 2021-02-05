package handler

import (
	"gorm.io/gorm"
	"github.com/midepeter/train-ticket/repository"
)

type Handler struct {
	db *gorm.DB
	Repo *repository.Repository
}

