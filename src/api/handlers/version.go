package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func GetVersionHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"commit": os.Getenv("COMMIT"),
	})
}
