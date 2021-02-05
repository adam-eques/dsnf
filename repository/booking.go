package repository

import (
	"github.com/midepeter/train-ticket/database/models"
)

func (r *Repository) Get(id uint, data map[string]interface{}) (*models.Ticket, error) {
	var ticket models.Ticket

	r.db.Model(&ticket).Where("id = ?", id).First(data)
	return &ticket, nil
}

func (r *Repository) Update(id uint, data map[string]interface{}) (*models.Ticket, error) {
	var ticket models.Ticket

	r.db.Model(&ticket).Where("id = ?", id).Update(data)
	return &ticket, nil
}

func (r *Repository) Delete(id uint) error {
	var ticket models.Ticket

	r.db.Model(&ticket).Where("id= ?", id).Delete()
	return nil
}
