package store

import "github.com/RustamRR/job-rest-api/internal/model"

type UserRepository interface {
	Create(*model.User) error
	Update(*model.User) error
	Find(string) error
	Delete(*model.User) error
}
