package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_CreateEnrichment(t *testing.T) {
	user := TestUser(t)
	assert.NoError(t, user.CreateEnrichment())
	assert.NotNil(t, user.ID)
	assert.NotNil(t, user.CreatedAt)
	assert.NotNil(t, user.UpdatedAt)
	assert.NotNil(t, user.Password)
}

func TestUser_ValidationCreate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *User {
				u := TestUser(t)
				_ = u.CreateEnrichment()
				return u
			},
			isValid: true,
		},
		{
			name: "not valid email",
			u: func() *User {
				u := TestUser(t)
				_ = u.CreateEnrichment()
				u.Email = "testcase"
				return u
			},
			isValid: false,
		},
		{
			name: "not valid password",
			u: func() *User {
				u := TestUser(t)
				_ = u.CreateEnrichment()
				u.Password = ""
				return u
			},
			isValid: false,
		},
		{
			name: "not valid sex",
			u: func() *User {
				u := TestUser(t)
				_ = u.CreateEnrichment()
				u.Sex = 0
				return u
			},
			isValid: false,
		},
		{
			name: "not valid first name",
			u: func() *User {
				u := TestUser(t)
				_ = u.CreateEnrichment()
				u.FirstName = ""
				return u
			},
			isValid: false,
		},
		{
			name: "not valid first name min length",
			u: func() *User {
				u := TestUser(t)
				_ = u.CreateEnrichment()
				u.FirstName = "w"
				return u
			},
			isValid: false,
		},
		{
			name: "not valid first name max length",
			u: func() *User {
				u := TestUser(t)
				_ = u.CreateEnrichment()
				u.FirstName = "weeeeeerwrwerwererwerewrwerwerwerwerwerewrewrewrw"
				return u
			},
			isValid: false,
		},
		{
			name: "not valid last name",
			u: func() *User {
				u := TestUser(t)
				_ = u.CreateEnrichment()
				u.LastName = ""
				return u
			},
			isValid: false,
		},
		{
			name: "not valid last name min length",
			u: func() *User {
				u := TestUser(t)
				_ = u.CreateEnrichment()
				u.LastName = "w"
				return u
			},
			isValid: false,
		},
		{
			name: "not valid last name max length",
			u: func() *User {
				u := TestUser(t)
				_ = u.CreateEnrichment()
				u.LastName = "weeeeeerwrwerwererwerewrwerwerwerwerwerewrewrewrw"
				return u
			},
			isValid: false,
		},
		{
			name: "valid birthday format",
			u: func() *User {
				u := TestUser(t)
				_ = u.CreateEnrichment()
				u.Birthday = "2008-07-30"
				return u
			},
			isValid: true,
		},
		{
			name: "not valid birthday format 1",
			u: func() *User {
				u := TestUser(t)
				_ = u.CreateEnrichment()
				u.Birthday = "2008-31-01"
				return u
			},
			isValid: false,
		},
		{
			name: "not valid birthday format 2",
			u: func() *User {
				u := TestUser(t)
				_ = u.CreateEnrichment()
				u.Birthday = "12.12.2008"
				return u
			},
			isValid: false,
		},
		{
			name: "not valid country",
			u: func() *User {
				u := TestUser(t)
				_ = u.CreateEnrichment()
				u.Country = ""
				return u
			},
			isValid: false,
		},
		{
			name: "not valid city",
			u: func() *User {
				u := TestUser(t)
				_ = u.CreateEnrichment()
				u.City = ""
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().ValidationCreate())
			} else {
				assert.Error(t, tc.u().ValidationCreate())
			}
		})
	}
}
