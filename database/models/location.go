package models

import (
	"github.com/google/uuid"
)

type Location struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	LocationName string    `json:"location_name"`
	PostalCode   string    `json:"postal_code"`
}
