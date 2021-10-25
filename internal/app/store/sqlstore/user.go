package sqlstore

import (
	"github.com/serdyanuk/go-rest/internal/app/core"
	"github.com/serdyanuk/go-rest/internal/app/model"
	"github.com/serdyanuk/go-rest/internal/app/store"
)

type UserRepostitory struct {
	*SqlStore
}

func NewUserRepostitory(s *SqlStore) store.UserRepostitory {
	return &UserRepostitory{s}
}

// Create creates a new user
func (s *UserRepostitory) Create(login, password string) (*model.User, error) {
	u := &model.User{
		Login:        login,
		PasswordHash: password,
	}
	if err := u.BeforeCreate(); err != nil {
		return nil, err
	}
	err := s.db.QueryRow("INSERT INTO users (login, password_hash) VALUES($1, $2) RETURNING id", u.Login, u.PasswordHash).Scan(&u.ID)
	if err != nil {
		return nil, core.NewSystemError(err, "user create failed")
	}
	return u, nil
}
