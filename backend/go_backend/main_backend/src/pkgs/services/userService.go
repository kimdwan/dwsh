package services

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/kimdwan/dwsh/settings"
	"github.com/kimdwan/dwsh/src/dtos"
	"github.com/kimdwan/dwsh/src/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// user와 관련된 곳에서 사용하는 파싱 and 검증 함수
func ParseAndCheckBodyVerUser[T dtos.UserSignUpDto](ctx *gin.Context) (*T, error) {
	var (
		body T
		err  error
	)

	if err = ctx.ShouldBindBodyWithJSON(&body); err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return nil, errors.New("(json) 클라이언트에서 보낸 폼을 파싱하는데 오류가 발생했습니다")
	}

	validate := validator.New()

	if err = validate.Struct(body); err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return nil, errors.New("(validate) 클라이언트에서 보낸 폼을 파싱하는데 오류가 발생했습니다")
	}

	return &body, nil
}

func UserSignUpService(c context.Context, signup_dto *dtos.UserSignUpDto) (int, error) {
	var (
		db          *gorm.DB = settings.DB
		check_user  models.User
		errorStatus int
		err         error
	)

	// ds 유저의 경우 secret key의 중복성 여부를 파악해야 한다.
	if signup_dto.User_type == "DS" {
		if errorStatus, err = UserSignUpCheckSecretKeyFunc(c, db, signup_dto, check_user); err != nil {
			return errorStatus, err
		}
	}

	// 이메일 중복여부를 확인하는 함수
	if errorStatus, err = UserSignUpCheckEmailFunc(c, db, signup_dto, check_user); err != nil {
		return errorStatus, err
	}

	// 새로운 데이터를 등록하는 함수
	if err = UserSignUpCreateUserFunc(c, db, signup_dto); err != nil {
		return http.StatusInternalServerError, err
	}

	return 0, nil
}

// secret key의 중복성을 확인하는 함수
func UserSignUpCheckSecretKeyFunc(c context.Context, db *gorm.DB, signup_dto *dtos.UserSignUpDto, check_user models.User) (int, error) {

	result := db.WithContext(c).Where("user_secretkey = ?", *signup_dto.Secret_key).First(&check_user)

	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("시스템 오류: ", result.Error.Error())
			return http.StatusInternalServerError, errors.New("데이터 베이스에서 secret key와 동일한 값을 찾는데 오류가 발생했습니다")
		}
	} else {
		return http.StatusExpectationFailed, errors.New("이미 가입된 비밀키 입니다")
	}

	return 0, nil
}

// 이메일의 중복성을 확인하는 함수
func UserSignUpCheckEmailFunc(c context.Context, db *gorm.DB, signup_dto *dtos.UserSignUpDto, check_user models.User) (int, error) {
	result := db.WithContext(c).Where("user_email = ?", signup_dto.Email).First(&check_user)

	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			fmt.Println("시스템 오류: ", result.Error.Error())
			return http.StatusInternalServerError, errors.New("이메일에 해당하는 데이터를 찾는데 오류가 발생했습니다")
		}
	} else {
		return http.StatusNotExtended, errors.New("이미 존재하는 이메일 입니다")
	}

	return 0, nil
}

// 새로운 유저를 만드는 로직
func UserSignUpCreateUserFunc(c context.Context, db *gorm.DB, signup_dto *dtos.UserSignUpDto) (err error) {
	var (
		new_user models.User
	)

	// 기본정보 저장 (이메일, 동의서)
	new_user.User_email = signup_dto.Email
	new_user.User_name = signup_dto.User_name
	new_user.User_term_agree_1 = signup_dto.Term_agree_1
	new_user.User_term_agree_2 = signup_dto.Term_agree_2
	new_user.User_term_agree_3 = signup_dto.Term_agree_3

	// 비밀번호 해쉬화 후 저장
	var (
		salt_rounds int = 10
		hash        []byte
	)

	if hash, err = bcrypt.GenerateFromPassword([]byte(signup_dto.Password), salt_rounds); err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return errors.New("bcrypt로 해쉬화를 진행하는 중 오류가 발생했습니다")
	}

	new_user.User_hash = hash

	// DS 타입의 유저일때
	var (
		type_values []string = strings.Split(os.Getenv("DATABASE_USER_TYPE"), ",")
	)
	if signup_dto.User_type == "DS" {
		var (
			secret_keys     []string = strings.Split(os.Getenv("DS_USER_SECRET_KEY"), ",")
			isDsUserAllowed bool     = false
		)

		for idx, secret_key := range secret_keys {
			if secret_key == *signup_dto.Secret_key {
				new_user.User_secretkey = signup_dto.Secret_key
				new_user.User_type = type_values[idx]
				isDsUserAllowed = true
				break
			}
		}

		if !isDsUserAllowed {
			return errors.New("해당 비밀키에 해당하는 데이터를 찾을수가 없습니다")
		}
	} else {
		new_user.User_type = type_values[len(type_values)-1]
	}

	// 데이터 저장
	result := db.WithContext(c).Save(&new_user)
	if result.Error != nil {
		fmt.Println("시스템 오류: ", result.Error.Error())
		return errors.New("데이터 베이스에 새로운 유저를 저장하는데 오류가 발생했습니다")
	}

	return nil

}
