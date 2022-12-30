package model

import (
	"context"
	"go-pgx/database"
	"log"
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

func DeletePost(postID uint) error {
	var err error

	_, err = database.DB.Exec(context.Background(), "UPDATE posts SET deleted_at = $1 WHERE id = $2", time.Now(), postID)
	if err != nil {
		return err
	}
	return nil
}

func GetAllPosts() (interface{}, error) {
	p := Post{}
	rows, err := database.DB.Query(context.Background(), "SELECT id, content, location, user_id, created_at FROM posts WHERE deleted_at IS NULL")
	if err != nil {
		log.Fatalf("Gagal eksekusi query: %v\n", err)
	}
	defer rows.Close()

	ps := []Post{}
	for rows.Next() {
		err = rows.Scan(&p.Id, &p.Content, &p.Location, &p.UserId, &p.CreatedAt)
		if err != nil {
			log.Fatalf("Gagal Scan query: %v\n", err)
		}
		ps = append(ps, p)
	}
	return ps, nil
}
