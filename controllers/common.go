package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Common struct {
}

func (_self Common) Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"ID":       1,
		"Email":    "1",
		"Password": "1",
	})
}
