package middlewares

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kimdwan/dwsh/settings"
	"github.com/kimdwan/dwsh/src/dtos"
	"github.com/kimdwan/dwsh/src/models"
	"gorm.io/gorm"
)

func AuthCheckJwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		// 토큰 읽기
		var (
			access_token string
			err          error
		)

		if access_token, err = ctx.Cookie("Authorization"); err != nil {
			fmt.Println("시스템 오류: ", err.Error())
			fmt.Println("클라이언트에서 보낸 jwt 토큰이 존재하지 않습니다")
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// access token을 파싱하는 로직
		var (
			payload         dtos.Payload
			jwt_secret_keys []string = []string{
				os.Getenv("JWT_ACCESS_SECRET"),
				os.Getenv("JWT_REFRESH_SECRET"),
			}
			counts int = 0
		)
		if err = AuthCheckJwtCheckJwtTokenFunc(access_token, jwt_secret_keys[0], &payload, &counts); err != nil {
			// refresh 인증 진행
			if counts > 0 {
				fmt.Println(err.Error())

				// refresh 토큰을 가져 오고 payload를 파싱하는 로직
				var (
					computer_number string = ctx.GetHeader("User-Computer-Number")
					refresh_token   string
					errorStatus     int
				)

				if computer_number == "" {
					fmt.Println("시스템 오류: header에 User-Computer-Number를 입력하지 않았습니다")
					ctx.AbortWithStatus(http.StatusBadRequest)
					return
				}

				c, cancel := context.WithTimeout(context.Background(), time.Second*100)
				defer cancel()

				// refresh 토큰을 가져오는 로직
				if errorStatus, err = AuthCheckJwtGetRefreshTokenFunc(c, computer_number, &refresh_token); err != nil {
					fmt.Println(err.Error())
					ctx.AbortWithStatus(errorStatus)
					return
				}

				// refresh token을 이용해서 다시 jwt 토큰을 검증하는 함수
				if err = AuthCheckJwtCheckJwtTokenFunc(refresh_token, jwt_secret_keys[1], &payload, &counts); err != nil {
					if counts > 1 {
						fmt.Println("refresh 토큰도 시간이 만료되었습니다 다시 로그인 해주세요")
					} else {
						fmt.Println(err.Error())
					}
					ctx.AbortWithStatus(http.StatusUnauthorized)
					return
				}

				// 새로운 access token을 발급하고 그에 대한 정리
				var (
					access_time_str string = os.Getenv("JWT_ACCESS_TIME")
				)
				if err = AuthCheckJwtGetNewAccessTokenFunc(ctx, c, jwt_secret_keys[0], access_time_str, &payload); err != nil {
					fmt.Println(err.Error())
					ctx.AbortWithStatus(http.StatusInternalServerError)
					return
				}

				// 인증 오류
			} else {
				fmt.Println(err.Error())
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}

		// 파싱한 payload를 다음 로직에게 전달하는 함수
		if err = AuthCheckJwtSendPayloadFunc(ctx, &payload); err != nil {
			fmt.Println(err.Error())
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		ctx.Next()

	}
}

// jwt 인증 로직
func AuthCheckJwtCheckJwtTokenFunc(jwt_token string, jwt_secret_key string, payload *dtos.Payload, counts *int) (err error) {

	// refresh 검증을 끝맞춘건지 확인하는 로직
	if *counts > 1 {
		return errors.New("refresh 검증도 끝맞쳤던 jwt 토큰입니다")
	}

	// jwt_token을 parsing 하는 로직
	parse_jwt_token, err := jwt.Parse(jwt_token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("클라이언트에서 보낸 jwt 토큰이 HMAC 방식인지 확인하시오")
		}
		return []byte(jwt_secret_key), nil
	})

	// jwt_token이 시간에 의한 오류인지 아니면 다른 이유인지 확인하는 로직
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			*counts += 1
			return errors.New("access token의 시간이 만료되었습니다 refresh 검증을 진행해야 합니다")
		} else {
			fmt.Println("시스템 오류: ", err.Error())
			return errors.New("jwt 토큰을 검증하는데 문제가 발생했습니단")
		}
	}

	// jwt 토큰에 저장된 payload를 끄집어 내는 중
	if claims, ok := parse_jwt_token.Claims.(jwt.MapClaims); ok {
		if payload_str, ok := claims["payload"]; ok {

			// 본격적으로 payload를 파싱하는 작업
			payload_byte, err := json.Marshal(payload_str)
			if err != nil {
				fmt.Println("시스템 오류: ", err.Error())
				return errors.New("payload를 byte화 하는데 오류가 발생했습니다")
			}

			if err = json.Unmarshal(payload_byte, payload); err != nil {
				fmt.Println("시스템 오류: ", err.Error())
				return errors.New("payload를 직렬화 하는데 오류가 발생했습니다")
			}

			return nil

		} else {
			return errors.New("claims에 저장된 payload를 파싱하는데 오류가 발생했습니다")
		}
	} else {
		return errors.New("jwt 토큰에 저장된 claims를 파싱하는데 오류가 발생했습니다")
	}

}

// refresh 토큰을 불러오는 로직
func AuthCheckJwtGetRefreshTokenFunc(c context.Context, computer_number string, refresh_token *string) (int, error) {
	var (
		db   *gorm.DB = settings.DB
		user models.User
	)

	// 데이터베이스에서 computer_number에 해당하는 refresh token을 가져옴
	result := db.WithContext(c).Where("user_computer_number = ?", computer_number).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("시스템 오류: computer_number가 잘못 되었음")
			return http.StatusUnauthorized, errors.New("computer_number를 다시 확인해 주세요")
		} else {
			fmt.Println("시스템 오류: ", result.Error.Error())
			return http.StatusInternalServerError, errors.New("computer_number에 맞는 유저 테이블을 찾는데 오류가 발생했습니다")
		}
	}

	*refresh_token = *user.User_refresh_token

	return 0, nil
}

func AuthCheckJwtGetNewAccessTokenFunc(ctx *gin.Context, c context.Context, access_secret_key string, access_time_str string, payload *dtos.Payload) error {

	// 시간을 파싱 하는 로직
	access_time, err := strconv.Atoi(access_time_str)
	if err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return errors.New("jwt 토큰의 access token의 시간을 정수화 하는데 오류가 발생했습니다")
	}

	// 새로운 access token을 만드는 함수
	jwt_token_str := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"payload": payload,
		"exp":     time.Now().Add(time.Duration(access_time) * time.Second).Unix(),
	})

	new_access_token, err := jwt_token_str.SignedString([]byte(access_secret_key))
	if err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return errors.New("새로운 access token을 만드는데 오류가 발생했습니다")
	}

	// 새로운 access 토큰을 저장하고 클라이언트에게 보냄
	var (
		db   *gorm.DB = settings.DB
		user models.User
	)

	result := db.WithContext(c).Where("user_id = ?", payload.User_id).First(&user)
	if result.Error != nil {
		fmt.Println("시스템 오류: ", result.Error.Error())
		return errors.New("user id와 동일한 데이터 베이스를 찾는데 오류가 발생했습니다")
	}

	user.User_access_token = &new_access_token
	if result = db.WithContext(c).Save(&user); result.Error != nil {
		fmt.Println("시스템 오류: ", result.Error.Error())
		return errors.New("데이터 베이스에 새로운 access token을 저장하는데 오류가 발생했습니다")
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", new_access_token, 24*60*60, "", "", false, true)
	return nil
}

func AuthCheckJwtSendPayloadFunc(ctx *gin.Context, paylaod *dtos.Payload) error {

	payload_byte, err := json.Marshal(paylaod)
	if err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return errors.New("payload를 직렬화하는데 오류가 발생했습니다")
	}

	ctx.Set("payload_byte", string(payload_byte))

	return nil
}
