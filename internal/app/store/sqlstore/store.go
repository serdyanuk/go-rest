package sqlstore

import (
	"database/sql"

	"github.com/serdyanuk/go-rest/config"
	"github.com/serdyanuk/go-rest/internal/app/store"
	"github.com/serdyanuk/go-rest/internal/pkg/logger"
)

type SqlStore struct {
	db           *sql.DB
	repositories *repositories
}

func New(cfg config.DBConfig) (*SqlStore, error) {
	db, err := createConnection(cfg)
	if err != nil {
		return nil, err
	}

	s := &SqlStore{db: db}
	s.repositories = newRepositories(s)

	return s, nil
}

func (s *SqlStore) Close() {
	s.db.Close()
}

func (s *SqlStore) User() store.UserRepostitory {
	return s.repositories.userRepository
}

func createConnection(cfg config.DBConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.DSN)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	l := logger.Get()
	l.Info("DB connect successful")

	return db, nil
}

type repositories struct {
	userRepository store.UserRepostitory
}

func newRepositories(s *SqlStore) *repositories {
	return &repositories{
		userRepository: NewUserRepostitory(s),
	}
}
