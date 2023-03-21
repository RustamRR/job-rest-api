package user

import (
	"bytes"
	"encoding/json"
	"github.com/RustamRR/job-rest-api/internal/app/apiserver/controller"
	"github.com/RustamRR/job-rest-api/internal/model"
	"github.com/RustamRR/job-rest-api/internal/store"
	"github.com/RustamRR/job-rest-api/internal/store/postgrestore"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type UserControllerTestSuite struct {
	suite.Suite
	router *echo.Echo
	store  store.Store
}

func (s *UserControllerTestSuite) SetupTest() {
	controller.TestConfig(s.T())
	s.store = postgrestore.TestStore(s.T())

	if err := s.store.Migrate(); err != nil {
		s.T().Fatal(err)
	}

	s.router = echo.New()
	InitRoutes(s.router, s.store)
}

func (s *UserControllerTestSuite) TearDownTest() {
	if err := s.store.ClearDB(); err != nil {
		s.T().Fatal(err)
	}
}

func (s *UserControllerTestSuite) TestUserController_CreateOk() {
	body := []byte(`{
		"email":"httptest@user.com", 
		"first_name":"John", 
		"last_name":"Doe",
		"birthday":"1990-01-22",
		"country":"Россия",
		"city":"Санкт-Петербург",
		"sex":1,
		"password":"userpassword"
	}`)
	req := httptest.NewRequest(http.MethodPost, "/users/create", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	context := s.router.NewContext(req, rec)

	err := create(context)
	assert.NoError(s.T(), err)

	var user model.User
	err = json.Unmarshal([]byte(rec.Body.String()), &user)
	assert.NoError(s.T(), err)

	assert.Equal(s.T(), http.StatusCreated, rec.Code)
	assert.NotEmpty(s.T(), user.ID)
	assert.NotEmpty(s.T(), user.CreatedAt)
	assert.NotEmpty(s.T(), user.UpdatedAt)
	assert.Equal(s.T(), user.Birthday, "1990-01-22")
	assert.Equal(s.T(), user.Email, "httptest@user.com")
	assert.Equal(s.T(), user.City, "Санкт-Петербург")
	assert.Equal(s.T(), user.Country, "Россия")
	assert.Equal(s.T(), user.Sex, model.Male)
	assert.Equal(s.T(), user.FirstName, "John")
	assert.Equal(s.T(), user.LastName, "Doe")
	assert.Empty(s.T(), user.Password)
}

func TestUserController_CreateOkTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
