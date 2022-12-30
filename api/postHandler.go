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
	res := model.Response{}
	if err := ctx.ShouldBind(&p); err != nil {
		res.StatusCode = "01"
		res.Description = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	uid, err := util.ExtractTokenID(ctx)
	if err != nil {
		res.StatusCode = "01"
		res.Description = "token " + err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	p.UserId = uid
	p.CreatedAt = time.Now()
	p.Location = util.GetIPAddress()
	err = model.CreatePost(&p)
	if err != nil {
		res.StatusCode = "01"
		res.Description = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res.StatusCode = "00"
	res.Description = "berhasil membuat postingan baru"
	ctx.JSON(http.StatusOK, res)

}

func DeletePost(ctx *gin.Context) {
	res := model.Response{}
	pIdStr := ctx.Param("id")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		res.StatusCode = "01"
		res.Description = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	err = model.DeletePost(uint(pId))
	if err != nil {
		res.StatusCode = "01"
		res.Description = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res.StatusCode = "00"
	res.Description = "berhasil menghapus postingan"
	ctx.JSON(http.StatusOK, res)
}

func LikePost(ctx *gin.Context) {
	res := model.Response{}
	pl := model.PostLike{}
	pIdStr := ctx.Param("id")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		res.StatusCode = "01"
		res.Description = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	uid, err := util.ExtractTokenID(ctx)
	if err != nil {
		res.StatusCode = "01"
		res.Description = "token " + err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	pl.PostID = uint(pId)
	pl.UserID = uid
	pl.CreatedAt = time.Now()
	err = model.AddLike(&pl)
	if err != nil {
		res.StatusCode = "01"
		res.Description = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res.StatusCode = "00"
	res.Description = "berhasil menyukai postingan"
	ctx.JSON(http.StatusOK, res)
}

func UnlikePost(ctx *gin.Context) {
	pl := model.PostLike{}
	res := model.Response{}
	pIdStr := ctx.Param("id")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		res.StatusCode = "01"
		res.Description = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	uid, err := util.ExtractTokenID(ctx)
	if err != nil {
		res.StatusCode = "01"
		res.Description = "token " + err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	pl.PostID = uint(pId)
	pl.UserID = uid
	pl.CreatedAt = time.Now()
	err = model.RemoveLike(&pl)
	if err != nil {
		res.StatusCode = "01"
		res.Description = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res.StatusCode = "00"
	res.Description = "berhasil unlike postingan"
	ctx.JSON(http.StatusOK, res)
}

func GetLikesCount(ctx *gin.Context) {
	pl := model.PostLike{}
	res := model.Response{}
	pIdStr := ctx.Param("id")
	pId, err := strconv.Atoi(pIdStr)
	if err != nil {
		res.StatusCode = "01"
		res.Description = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	_, err = util.ExtractTokenID(ctx)
	if err != nil {
		res.StatusCode = "01"
		res.Description = "token " + err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	pl.PostID = uint(pId)
	count := model.GetLikesCount(&pl)
	likes := model.CountLike{
		Likes: count,
	}
	res.StatusCode = "00"
	res.Description = "berhasil get count likes"
	res.Data = likes
	ctx.JSON(http.StatusOK, res)
}

func GetAllPosts(ctx *gin.Context) {
	res := model.Response{}
	ps, err := model.GetAllPosts()
	if err != nil {
		res.StatusCode = "01"
		res.Description = err.Error()
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	res.StatusCode = "00"
	res.Description = "berhasil get all Posts"
	res.Data = ps
	ctx.JSON(http.StatusOK, res)
}
