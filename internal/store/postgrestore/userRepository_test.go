package postgrestore

import (
	"github.com/RustamRR/job-rest-api/internal/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	store *Store
}

func (s *UserRepositoryTestSuite) SetupTest() {
	TestConfig(s.T())
	s.store = TestStore(s.T())

	if err := s.store.Migrate(); err != nil {
		s.T().Fatal(err)
	}
}

func (s *UserRepositoryTestSuite) TearDownTest() {
	if err := s.store.ClearDB(); err != nil {
		s.T().Fatal(err)
	}
}

func (s *UserRepositoryTestSuite) TestUserRepository_CreateSuccess() {
	user, repository := model.TestUser(s.T()), s.store.User()

	assert.NoError(s.T(), repository.Create(user))
}

func TestUserRepository_CreateSuccessTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}

func (s *UserRepositoryTestSuite) TestUserRepository_CreateFail() {
	testCases := []struct {
		name string
		u    func() *model.User
	}{
		{
			name: "empty password",
			u: func() *model.User {
				u := model.TestUser(s.T())
				u.Password = ""
				return u
			},
		},
		{
			name: "invalid min length password",
			u: func() *model.User {
				u := model.TestUser(s.T())
				u.Password = "123"
				return u
			},
		},
		{
			name: "invalid max length password",
			u: func() *model.User {
				u := model.TestUser(s.T())
				u.Password = "123ewrewrewqwerqwerqwrwereqwrqwerewqrweqrqwerwerweqr"
				return u
			},
		},
	}

	repository := s.store.User()

	for _, tc := range testCases {
		s.T().Run(tc.name, func(t *testing.T) {
			err := repository.Create(tc.u())
			assert.EqualError(s.T(), err, ErrCannotCreateUser.Error())
		})
	}
}

func TestUserRepository_CreateFailTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
