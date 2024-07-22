package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kimdwan/dwsh/src/dtos"
	"github.com/kimdwan/dwsh/src/pkgs/services"
)

func UserCheckUserTypeMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			body *dtos.UserSignUpDto
			err  error
		)

		// 클라이언트에서 보낸 폼 파싱하기
		body, err = services.ParseAndCheckBodyVerUser[dtos.UserSignUpDto](ctx)
		if err != nil {
			fmt.Println(err.Error())
			ctx.AbortWithStatus(http.StatusBadRequest)
		}

		// 이용약간 동의 확인하기
		if err = UserCheckUserTypeCheckTermAgreeOneAndTwoFunc(body); err != nil {
			fmt.Println(err.Error())
			ctx.AbortWithStatus(http.StatusBadGateway)
			return
		}

		// 유저 타입에 따른 만난날 또는 비밀키 확인하기
		if body.User_type == "DS" && body.Secret_key != nil && body.Our_first_day != nil {
			var (
				errorStatus int
			)

			errorStatus, err = UserCheckUserTypeCheckDsUserFunc(body)
			if err != nil {
				fmt.Println(err.Error())
				ctx.AbortWithStatus(errorStatus)
				return
			}
			fmt.Println("ds 계정 회원가입 시도")

		} else if body.User_type == "GUEST" && body.Secret_key == nil && body.Our_first_day == nil {
			fmt.Println("guest 계정 회원가입 시도")
		} else {
			fmt.Println("시스템 오류: 클라이언트에서 보낸 형식은 문제가 있습니다")
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		// 다음 미들웨어에서 사용할 수 있게 byte화 하는 로직
		var (
			signup_byte []byte
		)
		if signup_byte, err = json.Marshal(body); err != nil {
			fmt.Println("시스템 오류: ", err.Error())
			fmt.Println("body값을 다음 미들웨어가 사용할 수 있게 파싱하는데 오류가 발생했습니다")
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.Set("signup_byte", string(signup_byte))
		ctx.Next()
	}
}

func UserCheckUserTypeCheckDsUserFunc(body *dtos.UserSignUpDto) (int, error) {
	// 만난날 확인
	var (
		our_first_day_str string = os.Getenv("DS_USER_OUR_FIRST_DAY")
		our_first_day     time.Time
		system_first_day  time.Time
		err               error
	)

	if our_first_day, err = time.Parse("2006-01-02", our_first_day_str); err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return http.StatusInternalServerError, errors.New("첫번째 날을 파싱하는데 오류가 발생했습니다")
	}

	if system_first_day, err = time.Parse("2006-01-02", *body.Our_first_day); err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return http.StatusInternalServerError, errors.New("클라이언트가 보낸 첫번째 날을 파싱하는데 오류가 발생했습니다")
	}

	if our_first_day != system_first_day {
		return http.StatusNotAcceptable, errors.New("첫번째 날을 맞추지 못했습니다")
	}

	// 비밀키 확인
	var (
		secret_keys        []string = strings.Split(os.Getenv("DS_USER_SECRET_KEY"), ",")
		isSecretKeyAllowed bool     = false
	)

	for _, secret_key := range secret_keys {
		if secret_key == *body.Secret_key {
			isSecretKeyAllowed = true
			break
		}
	}

	if !isSecretKeyAllowed {
		return http.StatusUnauthorized, errors.New("secret key 입력이 잘못되었습니다")
	}

	return 0, nil
}

// 필수 사항을 동의 했는지 확인하는 로직
func UserCheckUserTypeCheckTermAgreeOneAndTwoFunc(body *dtos.UserSignUpDto) (err error) {
	if !body.Term_agree_1 || !body.Term_agree_2 {
		return errors.New("필수 사항 두개는 동의를 하셔야 합니다")
	}

	return nil
}
