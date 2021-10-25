package store

import (
	_ "github.com/lib/pq"
	"github.com/serdyanuk/go-rest/internal/app/model"
)

type Store interface {
	User() UserRepostitory
}

type UserRepostitory interface {
	Create(login, password string) (*model.User, error)
}
