package util

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(userid uint) (string, error) {
	tok_lifespan, err := strconv.Atoi(os.Getenv("TOKEN_LIFESPAN"))
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userid
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(tok_lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET")))
}

func TokenValid(ctx *gin.Context) error {
	tokenString := ExtractToken(ctx)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil
}
func ExtractToken(ctx *gin.Context) string {
	token := ctx.Query("token")
	if token != "" {
		return token
	}
	bearerTok := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerTok, " ")) == 2 {
		return strings.Split(bearerTok, " ")[1]
	}
	return ""
}

func ExtractTokenID(ctx *gin.Context) (uint, error) {
	tokString := ExtractToken(ctx)
	token, err := jwt.Parse(tokString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})
	if err != nil {
		return 0, nil
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, nil
		}
		return uint(uid), nil
	}
	return 0, nil

}
