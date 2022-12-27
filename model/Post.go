package model

import (
	"context"
	"go-pgx/database"
	"time"
)

type Post struct {
	Id        uint32    `json:"id"`
	Content   string    `json:"content"`
	Location  string    `json:"location"`
	UserId    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func CreatePost(p *Post) error {
	var err error

	_, err = database.DB.Exec(context.Background(), "INSERT INTO POSTS (content, location, user_id, created_at) values($1,$2,$3,$4)", p.Content, p.Location, p.UserId, p.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
