package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostPingHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}
