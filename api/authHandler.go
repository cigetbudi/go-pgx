package api

import (
	"go-pgx/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	u := model.User{}
	res := model.Response{}

	if err := ctx.ShouldBindJSON(&u); err != nil {
		res.StatusCode = "01"
		res.Description = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	err := model.AddUser(&u)
	if err != nil {
		res.StatusCode = "01"
		res.Description = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res.StatusCode = "00"
	res.Description = "berhasil mendaftar akun"
	ctx.JSON(http.StatusOK, res)
}

func Login(ctx *gin.Context) {
	u := model.User{}
	res := model.Response{}
	if err := ctx.ShouldBindJSON(&u); err != nil {
		res.StatusCode = "01"
		res.Description = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	logCount, err := model.CheckLoginAttemp(u.Username)
	if err != nil {
		res.StatusCode = "01"
		res.Description = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	if logCount > 3 {
		res.StatusCode = "01"
		res.Description = "akun anda terkunci, mohon menghubungi admin"
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	token, err := model.LoginCheck(u.Username, u.Password)
	if err != nil {
		err = model.AddLoginAttemp(u.Username)
		if err != nil {
			log.Fatal()
		}
		res.StatusCode = "01"
		res.Description = "username atau password tidak sesuai"
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	t := model.TokenResp{
		Token: token,
	}

	res.StatusCode = "00"
	res.Description = "berhasil login"
	res.Data = t
	ctx.JSON(http.StatusOK, res)
}
