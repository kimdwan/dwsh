package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserSignUpController(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "회원가입 되었습니다.",
	})
}
