package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID             uuid.UUID
	FirstName      string
	MiddleName     string
	LastName       string
	GenderID       int
	GenderName     string
	HashedPassword string
	Email          string
	Phone          string
	Bio            string
	PhotoURL       string
	ActiveFrom     time.Time
	ActiveTo       time.Time
}

func (u *User) ComparePassword(password string) error {
	if bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(password)) != nil {
		return errors.New("wrong password")
	}

	return nil
}
