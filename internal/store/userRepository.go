package store

import (
	"github.com/RustamRR/job-rest-api/internal/model"
	"github.com/google/uuid"
)

type UserRepository interface {
	Create(*model.User) error
	Update(*model.User) error
	Find(uuid.UUID) (result model.User, err error)
	FindAll() ([]model.User, error)
	Delete(*model.User) error
}
