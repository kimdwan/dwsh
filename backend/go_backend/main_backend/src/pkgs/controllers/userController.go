package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kimdwan/dwsh/src/dtos"
	"github.com/kimdwan/dwsh/src/pkgs/services"
)

// 회원가입을 담당하는 로직
func UserSignUpController(ctx *gin.Context) {

	var (
		signup_byte string = ctx.GetString("signup_byte")
		signup_dto  dtos.UserSignUpDto
		errorStatus int
		err         error
	)

	// 미들웨어에서 보낸 signup_byte를 직렬화 하는 로직
	if err = json.Unmarshal([]byte(signup_byte), &signup_dto); err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		fmt.Println("회원가입용 dto를 파싱하는데 오류가 발생했습니다")
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// 회원가입을 진행해 주는 함수
	c, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	if errorStatus, err = services.UserSignUpService(c, &signup_dto); err != nil {
		fmt.Println(err.Error())
		ctx.AbortWithStatus(errorStatus)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "회원가입 되었습니다.",
	})
}

// 로그인을 담당하는 로직
func UserLoginController(ctx *gin.Context) {
	var (
		login_dto       *dtos.UserLoginDto
		access_token    string
		computer_number uuid.UUID
		messages        string
		errorStatus     int
		err             error
	)

	// 로그인용 dto를 파싱하는 로직
	if login_dto, err = services.ParseAndCheckBodyVerUser[dtos.UserLoginDto](ctx); err != nil {
		fmt.Println(err.Error())
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// 로그인 로직
	c, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	if errorStatus, err = services.UserLoginService(c, login_dto, &access_token, &computer_number, &messages); err != nil {
		fmt.Println(err.Error())
		ctx.AbortWithStatus(errorStatus)
		return
	}

	// jwt access token을 쿠키로 보내는 로직 (나중에 더 보안을 강화하는 방법을 고려해야한다.)
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", access_token, 24*60*60, "", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{
		"message":         messages,
		"computer_number": computer_number,
	})
}
