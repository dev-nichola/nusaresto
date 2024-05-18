package menu

import (
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Update_at   time.Time `json:"update_at,omitempty"`
	Created_at  time.Time `json:"created_at,omitempty"`
}
