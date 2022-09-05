package user

import (
	"github.com/google/uuid"
)

type UserID string

func NewUserID() UserID {
	return UserID(uuid.NewString())
}

type User struct {
	ID       UserID
	Username string
	Password string
}

func NewUser(username, password string) *User {
	return &User{ID: NewUserID(), Username: username, Password: password}
}