package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID
	Name       string
	created_at time.Time
	updated_at time.Time
}
