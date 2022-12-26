package api

import (
	"go-pgx/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddBlacklisted(ctx *gin.Context) {
	b := model.Blacklisted{}
	if err := ctx.ShouldBindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "mohon untuk melengkapi semua isian"})
		return
	}

	err := model.AddBlacklisted(&b)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "berhasil membuat postingan baru"})
}
