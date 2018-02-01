package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// User user struct
type User struct {
	ID        uuid.UUID
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
