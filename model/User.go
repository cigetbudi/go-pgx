package model

import (
	"context"
	"go-pgx/database"
	"go-pgx/util"
	"html"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       uint32    `json:"id"`
	Username string    `json:"username"`
	Fullname string    `json:"fullname"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	DOB      time.Time `json:"DOB"`
}

func AddUser(u *User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	_, err = database.DB.Exec(context.Background(), "INSERT INTO USERS (username,fullname,password,email,created_at,dob) values($1,$2,$3,$4,$5,$6)", u.Username, u.Fullname, hashedPassword, u.Email, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func VerifyPassword(pass, hashedPass string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
}

func LoginCheck(username, password string) (string, error) {
	var (
		un  string
		pa  string
		err error
		id  uint
	)
	err = database.DB.QueryRow(context.Background(), "SELECT id, username, password FROM users WHERE username = $1", username).Scan(&id, &un, &pa)
	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, pa)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := util.GenerateToken(id)
	if err != nil {
		return "", err
	}

	_, err = database.DB.Exec(context.Background(), "UPDATE users SET last_login = $1 WHERE id = $2 ", time.Now(), id)
	if err != nil {
		return "", err
	}

	return token, nil
}

func CheckLoginAttemp(username string) (uint, error) {
	var logCount uint
	err := database.DB.QueryRow(context.Background(), "SELECT login_attempt FROM users where username = $1 ", username).Scan(&logCount)
	if err != nil {
		return 0, err
	}
	return logCount, nil
}

func AddLoginAttemp(username string) error {
	_, err := database.DB.Exec(context.Background(), "UPDATE users SET login_attempt = login_attempt + 1 WHERE username = $1", username)
	if err != nil {
		return err
	}
	return nil
}
