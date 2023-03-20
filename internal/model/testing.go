package model

import (
	"testing"
)

func TestUser(t *testing.T) *User {
	return &User{
		FirstName: "User",
		LastName:  "LastUser",
		Email:     "test@test.com",
		Sex:       Male,
		Birthday:  "1990-02-22",
		Country:   "Россия",
		City:      "Санкт-Петербург",
		Password:  "password",
	}
}
