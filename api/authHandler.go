package api

import (
	"go-pgx/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	u := model.User{}
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "mohon untuk melengkapi semua isian"})
		return
	}

	err := model.AddUser(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "berhasil mendaftar akun"})
}

func Login(ctx *gin.Context) {
	u := model.User{}
	if err := ctx.ShouldBindJSON(&u); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "mohon untuk melengkapi semua isian"})
		return
	}

	logCount, err := model.CheckLoginAttemp(u.Username)
	if err != nil {
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	}
	if logCount > 3 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "akun anda terkunci, silahkan hubungi admin untuk proses pembukaan"})
		return
	}

	token, err := model.LoginCheck(u.Username, u.Password)
	if err != nil {
		err = model.AddLoginAttemp(u.Username)
		if err != nil {
			log.Fatal()
		}
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "username atau password tidak sesuai"})
		return
	}

	ctx.JSON(http.StatusOK, token)

}
