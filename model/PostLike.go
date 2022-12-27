package model

import (
	"context"
	"go-pgx/database"
	"time"
)

type PostLike struct {
	Id        uint32    `json:"id"`
	PostID    uint      `json:"post_id"`
	UserID    uint      `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func AddLike(pl *PostLike) error {
	var err error
	_, err = database.DB.Exec(context.Background(), "INSERT INTO postlikes (post_id, user_id, created_at) values ($1, $2, $3)", pl.PostID, pl.UserID, pl.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func RemoveLike(pl *PostLike) error {
	var err error
	_, err = database.DB.Exec(context.Background(), "UPDATE postlikes SET deleted_at = $3 WHERE post_id = $1 AND user_id= $2", pl.PostID, pl.UserID, pl.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func GetLikesCount(pl *PostLike) uint {
	var count uint
	err := database.DB.QueryRow(context.Background(), "SELECT COUNT(user_id) FROM postlikes WHERE post_id = $1 ", pl.PostID).Scan(&count)
	if err != nil {
		return 0
	}
	return count
}
