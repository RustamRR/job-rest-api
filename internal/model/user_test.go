package model

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type UserModelTestSuite struct {
	suite.Suite
}

func (s *UserModelTestSuite) TestUser_CreateEnrichment() {
	user := TestUser(s.T())
	assert.NoError(s.T(), user.CreateEnrichment())
	assert.NotNil(s.T(), user.ID)
	assert.NotNil(s.T(), user.CreatedAt)
	assert.NotNil(s.T(), user.UpdatedAt)
	assert.NotNil(s.T(), user.Password)
}

func (s *UserModelTestSuite) TestUser_ValidationCreate() {
	testCases := []struct {
		name    string
		u       func() *User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *User {
				u := TestUser(s.T())
				_ = u.CreateEnrichment()
				return u
			},
			isValid: true,
		},
		{
			name: "not valid email",
			u: func() *User {
				u := TestUser(s.T())
				_ = u.CreateEnrichment()
				u.Email = "testcase"
				return u
			},
			isValid: false,
		},
		{
			name: "not valid password",
			u: func() *User {
				u := TestUser(s.T())
				_ = u.CreateEnrichment()
				u.HashedPassword = ""
				return u
			},
			isValid: false,
		},
		{
			name: "not valid sex",
			u: func() *User {
				u := TestUser(s.T())
				_ = u.CreateEnrichment()
				u.Sex = 0
				return u
			},
			isValid: false,
		},
		{
			name: "not valid first name",
			u: func() *User {
				u := TestUser(s.T())
				_ = u.CreateEnrichment()
				u.FirstName = ""
				return u
			},
			isValid: false,
		},
		{
			name: "not valid first name min length",
			u: func() *User {
				u := TestUser(s.T())
				_ = u.CreateEnrichment()
				u.FirstName = "w"
				return u
			},
			isValid: false,
		},
		{
			name: "not valid first name max length",
			u: func() *User {
				u := TestUser(s.T())
				_ = u.CreateEnrichment()
				u.FirstName = "weeeeeerwrwerwererwerewrwerwerwerwerwerewrewrewrw"
				return u
			},
			isValid: false,
		},
		{
			name: "not valid last name",
			u: func() *User {
				u := TestUser(s.T())
				_ = u.CreateEnrichment()
				u.LastName = ""
				return u
			},
			isValid: false,
		},
		{
			name: "not valid last name min length",
			u: func() *User {
				u := TestUser(s.T())
				_ = u.CreateEnrichment()
				u.LastName = "w"
				return u
			},
			isValid: false,
		},
		{
			name: "not valid last name max length",
			u: func() *User {
				u := TestUser(s.T())
				_ = u.CreateEnrichment()
				u.LastName = "weeeeeerwrwerwererwerewrwerwerwerwerwerewrewrewrw"
				return u
			},
			isValid: false,
		},
		{
			name: "valid birthday format",
			u: func() *User {
				u := TestUser(s.T())
				_ = u.CreateEnrichment()
				u.Birthday = "2008-07-30"
				return u
			},
			isValid: true,
		},
		{
			name: "not valid birthday format 1",
			u: func() *User {
				u := TestUser(s.T())
				_ = u.CreateEnrichment()
				u.Birthday = "2008-31-01"
				return u
			},
			isValid: false,
		},
		{
			name: "not valid birthday format 2",
			u: func() *User {
				u := TestUser(s.T())
				_ = u.CreateEnrichment()
				u.Birthday = "12.12.2008"
				return u
			},
			isValid: false,
		},
		{
			name: "not valid country",
			u: func() *User {
				u := TestUser(s.T())
				_ = u.CreateEnrichment()
				u.Country = ""
				return u
			},
			isValid: false,
		},
		{
			name: "not valid city",
			u: func() *User {
				u := TestUser(s.T())
				_ = u.CreateEnrichment()
				u.City = ""
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().ValidationCreate())
			} else {
				assert.Error(t, tc.u().ValidationCreate())
			}
		})
	}
}

func (s *UserModelTestSuite) TestUser_ValidationUpdate() {
	testCases := []struct {
		name    string
		user    func() *User
		isValid bool
	}{
		{
			name: "valid validation update",
			user: func() *User {
				u := TestUser(s.T())
				err := u.CreateEnrichment()
				u.UpdatedAt = time.Now()
				if err != nil {
					s.T().Errorf("enrichment user error: %v", err)
				}
				return u
			},
			isValid: true,
		},
		{
			name: "not valid first name",
			user: func() *User {
				u := TestUser(s.T())
				err := u.CreateEnrichment()
				u.UpdatedAt = time.Now()
				u.FirstName = ""
				if err != nil {
					s.T().Errorf("enrichment user error: %v", err)
				}
				return u
			},
			isValid: false,
		},
		{
			name: "not valid last name",
			user: func() *User {
				u := TestUser(s.T())
				err := u.CreateEnrichment()
				u.UpdatedAt = time.Now()
				u.LastName = ""
				if err != nil {
					s.T().Errorf("enrichment user error: %v", err)
				}
				return u
			},
			isValid: false,
		},
		{
			name: "not valid birth date",
			user: func() *User {
				u := TestUser(s.T())
				err := u.CreateEnrichment()
				u.UpdatedAt = time.Now()
				u.Birthday = ""
				if err != nil {
					s.T().Errorf("enrichment user error: %v", err)
				}
				return u
			},
			isValid: false,
		},
		{
			name: "not valid city",
			user: func() *User {
				u := TestUser(s.T())
				err := u.CreateEnrichment()
				u.UpdatedAt = time.Now()
				u.City = ""
				if err != nil {
					s.T().Errorf("enrichment user error: %v", err)
				}
				return u
			},
			isValid: false,
		},
		{
			name: "not valid country",
			user: func() *User {
				u := TestUser(s.T())
				err := u.CreateEnrichment()
				u.UpdatedAt = time.Now()
				u.Country = ""
				if err != nil {
					s.T().Errorf("enrichment user error: %v", err)
				}
				return u
			},
			isValid: false,
		},
		{
			name: "not valid sex",
			user: func() *User {
				u := TestUser(s.T())
				err := u.CreateEnrichment()
				u.UpdatedAt = time.Now()
				u.Sex = 9
				if err != nil {
					s.T().Errorf("enrichment user error: %v", err)
				}
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			u := tc.user()
			assert.True(t, u.CreatedAt.Before(u.UpdatedAt))
			if tc.isValid {
				assert.NoError(t, u.ValidationUpdate())
			} else {
				assert.Error(t, u.ValidationUpdate())
			}
		})
	}

	user := TestUser(s.T())
	err := user.CreateEnrichment()
	if err != nil {
		s.T().Errorf("enrichment user error: %v", err)
	}

	s.T().Run("not valid updated at", func(t *testing.T) {
		assert.True(t, user.CreatedAt.Equal(user.UpdatedAt))
	})
}

func TestUserModelTestSuite(t *testing.T) {
	suite.Run(t, new(UserModelTestSuite))
}
