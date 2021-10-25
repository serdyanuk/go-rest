package model

type User struct {
	ID           int    `json:"id"`
	Login        string `json:"login"`
	PasswordHash string `json:"-"`
}

func (u *User) BeforeCreate() error {
	return nil
}
