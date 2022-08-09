package model

import (
	"time"

	"github.com/google/uuid"
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
