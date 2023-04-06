package postgrestore

import (
	"fmt"
	"github.com/RustamRR/job-rest-api/internal/model"
	"github.com/RustamRR/job-rest-api/internal/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	store *Store
}

var testUser *model.User

func (s *UserRepositoryTestSuite) SetupSuite() {
	TestConfig(s.T())
	db := utils.TestDB(s.T())
	s.store = New(db)

	if err := s.store.Migrate(); err != nil {
		s.T().Fatal(err)
	}
}

func (s *UserRepositoryTestSuite) SetupTest() {
	TestConfig(s.T())
	db := utils.TestDB(s.T())
	s.store = New(db)

	if err := s.store.Migrate(); err != nil {
		s.T().Fatal(err)
	}
}

func (s *UserRepositoryTestSuite) TearDownSuite() {
	if err := s.store.ClearDB(); err != nil {
		s.T().Fatal(err)
	}
}

func (s *UserRepositoryTestSuite) TestUserRepository_CreateSuccess() {
	user1, user2, repository := model.TestUser(s.T()), model.TestUser(s.T()), s.store.User()
	user2.Email = "duplicate@email.com"
	assert.NoError(s.T(), repository.Create(user1))
	assert.NoError(s.T(), repository.Create(user2))
	testUser = user1
}

func (s *UserRepositoryTestSuite) TestUserRepository_CreateFail() {
	testCases := []struct {
		name string
		u    func() *model.User
		err  string
	}{
		{
			name: "empty password",
			u: func() *model.User {
				u := model.TestUser(s.T())
				u.Password = ""
				return u
			},
			err: "cannot be blank",
		},
		{
			name: "invalid min length password",
			u: func() *model.User {
				u := model.TestUser(s.T())
				u.Password = "123"
				return u
			},
			err: "the length must be between 6 and 18",
		},
		{
			name: "invalid max length password",
			u: func() *model.User {
				u := model.TestUser(s.T())
				u.Password = "123ewrewrewqwerqwerqwrwereqwrqwerewqrweqrqwerwerweqr"
				return u
			},
			err: "the length must be between 6 and 18",
		},
	}

	repository := s.store.User()

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			err := repository.Create(tc.u())
			assert.EqualError(s.T(), err, tc.err)
		})
	}
}

func (s *UserRepositoryTestSuite) TestUserRepository_FindById() {
	testCases := []struct {
		name    string
		id      uuid.UUID
		isValid bool
	}{
		{
			name:    "valid",
			id:      testUser.ID,
			isValid: true,
		},
		{
			name:    "not valid",
			id:      uuid.New(),
			isValid: false,
		},
	}

	for _, tc := range testCases {
		if tc.isValid {
			s.T().Run(tc.name, func(t *testing.T) {
				u, err := s.store.User().Find(tc.id)
				assert.NoError(t, err)
				assert.Equal(t, testUser.FirstName, u.FirstName)
				assert.Equal(t, testUser.LastName, u.LastName)
				assert.Equal(t, testUser.City, u.City)
				assert.Equal(t, testUser.Country, u.Country)
				assert.NotNil(t, u.ID)
			})
		} else {
			s.T().Run(tc.name, func(t *testing.T) {
				_, err := s.store.User().Find(tc.id)
				assert.Error(t, err)
				assert.EqualError(t, err, gorm.ErrRecordNotFound.Error())
			})
		}
	}
}

func (s *UserRepositoryTestSuite) TestUserRepository_NotUniqueEmail() {
	userRepeat := model.TestUser(s.T())
	userRepeat.Email = "duplicate@email.com"
	err := s.store.User().Create(userRepeat)

	assert.Error(s.T(), err)
	assert.EqualError(s.T(), err, gorm.ErrDuplicatedKey.Error())
}

func (s *UserRepositoryTestSuite) TestUserRepository_FindAll() {
	users, err := s.store.User().FindAll()
	s.T().Run("No error when find all users", func(t *testing.T) {
		assert.NoError(s.T(), err)
	})

	s.T().Run("Count of users is 2", func(t *testing.T) {
		assert.Equal(s.T(), 2, len(users))
	})

	for idx, userModel := range users {
		s.T().Run(fmt.Sprintf("Assertions for %d user", idx+1), func(t *testing.T) {
			assert.IsType(s.T(), model.User{}, userModel)
		})
	}
}

func (s *UserRepositoryTestSuite) TestUserRepository_Update() {
	user, err := s.store.User().Find(testUser.ID)
	if err != nil {
		s.T().Errorf("error find test user in update test: %v", err)
	}

	user.Email = "mynot@real.mail"
	user.FirstName = "Orororororor"
	user.LastName = "Gogogo"
	user.Birthday = "1992-12-31"
	err = s.store.User().Update(&user)
	assert.NoError(s.T(), err)

	updatedUser, errFind := s.store.User().Find(testUser.ID)
	if errFind != nil {
		s.T().Errorf("error find user from db: %v", errFind)
	}
	assert.Equal(s.T(), user.FirstName, updatedUser.FirstName)
	assert.Equal(s.T(), user.LastName, updatedUser.LastName)
	assert.Equal(s.T(), user.Birthday, updatedUser.Birthday)
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
