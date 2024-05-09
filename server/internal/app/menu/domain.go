package menu

import (
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	UUID       uuid.UUID
	Name       string
	Update_at  time.Time
	Created_at time.Time
}
