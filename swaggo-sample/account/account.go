package account

import uuid "github.com/satori/go.uuid"

// Account account struct
type Account struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

// Payload account payload struct
type Payload struct {
	ID       uuid.UUID `json:"id"`
	Email    string    `json:"email"`
	Passowrd string    `json:"password"`
}
