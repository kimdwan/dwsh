package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kimdwan/dwsh/src/dtos"
	"github.com/kimdwan/dwsh/src/pkgs/services"
)

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
