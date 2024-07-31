package controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kimdwan/dwsh/src/dtos"
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

func AuthGetProfileController(ctx *gin.Context) {
	var (
		payload     *dtos.Payload
		image_types dtos.ImageType
		err         error
	)

	// payload를 읽는 함수
	if payload, err = services.AuthParsePayloadService(ctx); err != nil {
		fmt.Println(err.Error())
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// 이미지를 읽고 보냄
	c, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	if errorStatus, err := services.AuthGetProfileService(c, payload, &image_types); err != nil {
		fmt.Println(err.Error())
		ctx.AbortWithStatus(errorStatus)
		return
	}

	ctx.JSON(http.StatusOK, image_types)

}

func AuthUserLogoutController(ctx *gin.Context) {
	var (
		payload *dtos.Payload
		err     error
	)

	// payload를 읽어오는 함수
	if payload, err = services.AuthParsePayloadService(ctx); err != nil {
		fmt.Println(err.Error())
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// 유저의 데이터 베이스를 비워주는 함수
	if errorStatus, err := services.AuthUserLogoutService(payload); err != nil {
		fmt.Println(err.Error())
		ctx.AbortWithStatus(errorStatus)
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", "", 24*60*60, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{
		"message": "로그아웃 되었습니다.",
	})
}
