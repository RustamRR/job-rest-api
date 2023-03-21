package model

import (
	"github.com/RustamRR/job-rest-api/internal/utils"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Sex int

const (
	Male Sex = iota + 1
	Female
)

func (s Sex) String() string {
	return [...]string{"Мужской", "Женский"}[s-1]
}

type User struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;" json:"id"`
	Email     string    `gorm:"not null" json:"email"`
	FirstName string    `gorm:"not null" json:"first_name"`
	LastName  string    `gorm:"not null" json:"last_name"`
	Birthday  string    `json:"birthday"`
	Country   string    `json:"country"`
	City      string    `json:"city"`
	Sex       Sex       `json:"sex"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (u *User) Sanitize() {
	u.Password = ""
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if err := u.CreateEnrichment(); err != nil {
		return err
	}

	if err := u.ValidationCreate(); err != nil {
		return err
	}

	return nil
}

func (u *User) CreateEnrichment() error {
	u.ID = uuid.New()
	if err := validation.Validate(
		u.Password,
		validation.Required,
		validation.Length(6, 18),
	); err != nil {
		return err
	}

	password, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	u.Password = password

	currentTime := time.Now()
	u.CreatedAt, u.UpdatedAt = currentTime, currentTime

	return nil
}

func (u *User) ValidationCreate() error {
	return validation.ValidateStruct(u,
		validation.Field(&u.ID, validation.Required, is.UUID),
		validation.Field(&u.FirstName, validation.Required, validation.Length(2, 15)),
		validation.Field(&u.LastName, validation.Required, validation.Length(2, 25)),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Birthday, validation.Required, validation.Date("2006-01-02")),
		validation.Field(&u.Country, validation.Required),
		validation.Field(&u.City, validation.Required),
		validation.Field(&u.Password, validation.Required),
		validation.Field(&u.Sex, validation.Required, validation.In(Male, Female)),
		validation.Field(&u.CreatedAt, validation.Required),
		validation.Field(&u.UpdatedAt, validation.Required),
	)
}
