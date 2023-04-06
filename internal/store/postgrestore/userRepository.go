package postgrestore

import (
	"github.com/RustamRR/job-rest-api/internal/model"
	"github.com/google/uuid"
)

type UserRepository struct {
	store *Store
}

func (r *UserRepository) Create(user *model.User) error {
	if err := r.store.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) Update(user *model.User) error {
	if err := r.store.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) Find(id uuid.UUID) (model.User, error) {
	var user model.User

	if err := r.store.db.First(&user, "id = ?", id); err != nil {
		return user, err.Error
	}

	return user, nil
}

func (r *UserRepository) FindAll() ([]model.User, error) {
	var users []model.User

	result := r.store.db.Find(&users)
	if result.Error != nil {
		return users, result.Error
	}

	return users, nil
}

func (r *UserRepository) Delete(user *model.User) error {
	return nil
}
