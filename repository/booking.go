package repository

import (
	"fmt"

	"github.com/midepeter/train-ticket/database/models"
)

func (r *Repository) GetBooking(id uint, t *models.Ticket) (err error) {
	if err = r.db.Where("id = ?", id).First(t).Error; err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateBooking(id uint, t *models.Ticket) (err error) {
	fmt.Println(t)
	r.db.Save(t)
	return nil
}

func (r *Repository) DeleteBooking(id uint, t *models.Ticket) error {
	r.db.Where("id= ?", id).Delete(t)
	return nil
}
