package user

import (
	"bytes"
	"encoding/json"
	"github.com/RustamRR/job-rest-api/internal/app/apiserver/controller"
	"github.com/RustamRR/job-rest-api/internal/model"
	"github.com/RustamRR/job-rest-api/internal/store"
	"github.com/RustamRR/job-rest-api/internal/store/postgrestore"
	"github.com/RustamRR/job-rest-api/internal/utils"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

var testUser model.User

type UserControllerTestSuite struct {
	suite.Suite
	router *echo.Echo
	store  store.Store
}

func (s *UserControllerTestSuite) SetupSuite() {
	controller.TestConfig(s.T())
	db := utils.TestDB(s.T())
	s.store = postgrestore.New(db)
	if err := s.store.Migrate(); err != nil {
		s.T().Fatal(err)
	}

	s.router = echo.New()
	InitRoutes(s.router, s.store)
}

func (s *UserControllerTestSuite) TearDownSuite() {
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
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	context := s.router.NewContext(req, rec)

	err := create(context)
	assert.NoError(s.T(), err)

	err = json.Unmarshal([]byte(rec.Body.String()), &testUser)
	assert.NoError(s.T(), err)

	assert.Equal(s.T(), http.StatusCreated, rec.Code)
	assert.NotEmpty(s.T(), testUser.ID)
	assert.NotEmpty(s.T(), testUser.CreatedAt)
	assert.NotEmpty(s.T(), testUser.UpdatedAt)
	assert.Equal(s.T(), testUser.Birthday, "1990-01-22")
	assert.Equal(s.T(), testUser.Email, "httptest@user.com")
	assert.Equal(s.T(), testUser.City, "Санкт-Петербург")
	assert.Equal(s.T(), testUser.Country, "Россия")
	assert.Equal(s.T(), testUser.Sex, model.Male)
	assert.Equal(s.T(), testUser.FirstName, "John")
	assert.Equal(s.T(), testUser.LastName, "Doe")
	assert.Empty(s.T(), testUser.Password)
}

func (s *UserControllerTestSuite) TestUserController_GetOk() {
	uuidValue := testUser.ID.String()
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	context := s.router.NewContext(req, rec)
	context.SetParamNames("id")
	context.SetParamValues(uuidValue)
	err := get(context)

	assert.NoError(s.T(), err)

	var user model.User
	err = json.Unmarshal([]byte(rec.Body.String()), &user)
	assert.NoError(s.T(), err)

	assert.Equal(s.T(), http.StatusOK, rec.Code)
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

func (s *UserControllerTestSuite) TestUserController_GetAllOk() {
	testUser2 := model.TestUser(s.T())
	testUser2.Email = "second@user.test"

	err2 := s.store.User().Create(testUser2)
	if err2 != nil {
		s.T().Errorf("Error when create user 2: %v", err2)
	}

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	context := s.router.NewContext(req, rec)
	err := getAll(context)

	s.T().Run("No error when get all users", func(t *testing.T) {
		assert.NoError(t, err)
	})

	var users []model.User
	err = json.Unmarshal([]byte(rec.Body.String()), &users)
	s.T().Run("No error when unmarshal response body", func(t *testing.T) {
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, 2, len(users))
	})
}

func (s *UserControllerTestSuite) TestUserController_UpdateOk() {
	id := testUser.ID.String()
	body := []byte(`{
		"first_name":"Peter",
		"last_name":"Meter",
		"sex":2,
		"birthday":"1992-04-18",
		"country":"Belarus"
	}`)

	req := httptest.NewRequest(http.MethodPatch, "/users", bytes.NewBuffer(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	context := s.router.NewContext(req, rec)
	context.SetPath("/users/:id")
	context.SetParamNames("id")
	context.SetParamValues(id)

	if err := update(context); err != nil {
		s.T().Errorf("update user error: %v", err)
	}

	s.T().Log(rec.Body.String())

	var user model.User
	if err := json.Unmarshal([]byte(rec.Body.String()), &user); err != nil {
		s.T().Errorf("unmarshal error:%v", err)
	}

	assert.Equal(s.T(), "Peter", user.FirstName)
	assert.Equal(s.T(), "Meter", user.LastName)
	assert.Equal(s.T(), model.Female, user.Sex)
	assert.Equal(s.T(), "1992-04-18", user.Birthday)
	assert.Equal(s.T(), "Belarus", user.Country)
	assert.Equal(s.T(), testUser.City, user.City)
}

func TestUserControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UserControllerTestSuite))
}
