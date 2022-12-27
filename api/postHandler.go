package api

import (
	"go-pgx/model"
	"go-pgx/util"
	"net/http"
	"strconv"
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

func LikePost(ctx *gin.Context) {
	pl := model.PostLike{}
	pIdStr := ctx.Param("id")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid, err := util.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "login tidak sah, harap login kembali"})
		return
	}
	pl.PostID = uint(pId)
	pl.UserID = uid
	pl.CreatedAt = time.Now()
	err = model.AddLike(&pl)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "berhasil like postingan ini"})
}

func UnlikePost(ctx *gin.Context) {
	pl := model.PostLike{}
	pIdStr := ctx.Param("id")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	uid, err := util.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "login tidak sah, harap login kembali"})
		return
	}
	pl.PostID = uint(pId)
	pl.UserID = uid
	pl.CreatedAt = time.Now()
	err = model.RemoveLike(&pl)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "berhasil unlike postingan ini"})

}

func GetLikesCount(ctx *gin.Context) {
	pl := model.PostLike{}
	pIdStr := ctx.Param("id")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = util.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "login tidak sah, harap login kembali"})
		return
	}
	pl.PostID = uint(pId)
	count := model.GetLikesCount(&pl)
	ctx.JSON(http.StatusOK, gin.H{"likes": count})
}
