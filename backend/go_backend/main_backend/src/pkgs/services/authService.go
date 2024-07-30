package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kimdwan/dwsh/src/dtos"
)

func AuthTestService(ctx *gin.Context) (int, error) {

	var (
		err error
	)

	// payload byte를 체크하는 함수
	if err = AuthTestCheckPayloadFunc(ctx); err != nil {
		return http.StatusInternalServerError, err
	}

	return 0, nil
}

func AuthTestCheckPayloadFunc(ctx *gin.Context) (err error) {
	var (
		payload_byte string = ctx.GetString("payload_byte")
		payload      dtos.Payload
	)

	if err = json.Unmarshal([]byte(payload_byte), &payload); err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return errors.New("payload byte를 역직렬화 하는데 오류가 발생했습니다")
	}

	fmt.Println("payload 이상 무")
	fmt.Println(payload)
	return nil
}
