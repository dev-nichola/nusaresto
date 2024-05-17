package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID
	Name       string
	Email      string
	Role_id    int
	Password   string
	Created_at time.Time
	Updated_at time.Time
}
