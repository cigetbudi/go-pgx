package model

import (
	"context"
	"go-pgx/database"
)

type Blacklisted struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func AddBlacklisted(b *Blacklisted) error {
	_, err := database.DB.Exec(context.Background(), "insert into blacklisted(username, email) values ($1, $2)", b.Username, b.Email)
	return err
}
