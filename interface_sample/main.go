package main

import (
	"fmt"
)

// Account account interface
type Account interface {
	GetEmail()
}

// User user struct
type User struct {
	Email string
}

// GetEmail email func
type GetEmail func(u *User) string

// Handler handler
type Handler interface {
	SendMail()
}

// SendMail send mail func
func (f GetEmail) SendMail() {
	fmt.Printf("send mail to: %v", f)
}

func main() {
	user := new(User)
	f := GetEmail(func(u *User) string {
		u.Email = "yyy@test.com"
		return u.Email
	})

	// pp.Println(user)
	// pp.Println(f)
	f.SendMail()

}
