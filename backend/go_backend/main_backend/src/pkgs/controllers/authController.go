package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kimdwan/dwsh/src/pkgs/services"
)

func AuthTestController(ctx *gin.Context) {

	// payload가 제대로 파싱이 됬는지 확인하는 로직
	var (
		errorStatus int
		err         error
	)

	// 테스트 진행
	if errorStatus, err = services.AuthTestService(ctx); err != nil {
		fmt.Println(err.Error())
		ctx.AbortWithStatus(errorStatus)
		return
	}

	// 확인
	ctx.JSON(http.StatusOK, gin.H{
		"message": "테스트 완료",
	})
}
