package entity

import (
	"time"

	"github.com/google/uuid"
)

type Ship struct {
	ID           uuid.UUID
	Mode         string
	Status       string
	DelivaryDate time.Time
}
