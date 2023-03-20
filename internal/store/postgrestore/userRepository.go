package postgrestore

import (
	"errors"
	"github.com/RustamRR/job-rest-api/internal/model"
)

var ErrCannotCreateUser error = errors.New("не получилось создать пользователя")

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(user *model.User) error {
	if err := r.store.db.Create(user).Error; err != nil {
		return ErrCannotCreateUser
	}

	return nil
}