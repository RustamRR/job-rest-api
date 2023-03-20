package store

type Store interface {
	User() UserRepository
	Migrate() error
}
