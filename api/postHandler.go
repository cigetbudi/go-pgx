package api

import (
	"go-pgx/model"
	"go-pgx/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreatePost(ctx *gin.Context) {
	p := model.Post{}
	if err := ctx.ShouldBind(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "mohon untuk melengkapi semua isian"})
		return
	}
	uid, err := util.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "login tidak sah, harap login kembali"})
		return
	}
	p.UserId = uid
	p.CreatedAt = time.Now()
	p.Location = util.GetIPAddress()
	err = model.CreatePost(&p)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "berhasil membuat postingan baru"})
}
