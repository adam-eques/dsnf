package models

import "github.com/google/uuid"

type Location struct {
	ID           uuid.UUID   `json:"id"`
	LocationName string      `json:"location_name"`
	PostalCode   string      `json:"postal_code"`
}
