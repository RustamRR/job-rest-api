package postgrestore

import (
	"github.com/RustamRR/job-rest-api/internal/model"
	"gorm.io/gorm"
)

type Store struct {
	db             *gorm.DB
	userRepository *UserRepository
}

func New(db *gorm.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() *UserRepository {
	if s.userRepository == nil {
		s.userRepository = &UserRepository{
			store: s,
		}
	}

	return s.userRepository
}

func (s *Store) Migrate() error {
	if err := s.db.AutoMigrate(&model.User{}); err != nil {
		return err
	}

	return nil
}

func (s *Store) ClearDB() error {
	if err := s.db.Migrator().DropTable(&model.User{}); err != nil {
		return err
	}

	return nil
}
