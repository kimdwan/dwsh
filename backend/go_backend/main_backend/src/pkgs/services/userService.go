package services

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/kimdwan/dwsh/settings"
	"github.com/kimdwan/dwsh/src/dtos"
	"github.com/kimdwan/dwsh/src/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// user와 관련된 곳에서 사용하는 파싱 and 검증 함수
func ParseAndCheckBodyVerUser[T dtos.UserSignUpDto | dtos.UserLoginDto](ctx *gin.Context) (*T, error) {
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

// 로그인을 담당하는 함수
func UserLoginService(c context.Context, login_dto *dtos.UserLoginDto, access_token *string, computer_number *uuid.UUID, messages *string, user_type *string) (int, error) {
	var (
		db          *gorm.DB = settings.DB
		check_user  models.User
		errorStatus int
		err         error
	)

	// 이메일 확인
	if err = UserLoginCheckEmailFunc(c, db, login_dto, &check_user, &errorStatus); err != nil {
		return errorStatus, err
	}

	// 비밀번호 확인
	if err = UserLoginCheckPasswordFunc(db, login_dto, &check_user, &errorStatus); err != nil {
		return errorStatus, err
	}

	// 기존에 로그인 되어있는지 확인하는 함수
	UserLoginCheckComputerNumberFunc(&check_user, messages)

	// jwt 토큰을 만드는 함수
	if err = UserLoginMakeJwtTokenFunc(&check_user, access_token, &errorStatus); err != nil {
		return errorStatus, err
	}

	// computer_number를 저장하고 유저 타입까지 불러오기 데이터 베이스에 마무리까지 하는 함수
	if err = UserLoginMakeComputerNumberAndSaveDatabaseFunc(c, db, &check_user, computer_number, user_type, &errorStatus); err != nil {
		return errorStatus, err
	}

	return 0, nil

}

// 이메일을 확인하는 함수
func UserLoginCheckEmailFunc(c context.Context, db *gorm.DB, login_dto *dtos.UserLoginDto, check_user *models.User, errorStatus *int) error {
	result := db.WithContext(c).Where("user_email = ?", login_dto.Email).First(check_user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			*errorStatus = http.StatusNotFound
			return errors.New("해당 이메일은 존재하지 않습니다")
		} else {
			fmt.Println("시스템 오류: ", result.Error.Error())
			*errorStatus = http.StatusInternalServerError
			return errors.New("이메일에 맞는 데이터 베이스를 찾는데 오류가 발생했습니다")
		}
	}
	return nil
}

// 비밀번호를 확인하는 함수
func UserLoginCheckPasswordFunc(db *gorm.DB, login_dto *dtos.UserLoginDto, check_user *models.User, errorStatus *int) (err error) {
	if err = bcrypt.CompareHashAndPassword(check_user.User_hash, []byte(login_dto.Password)); err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		*errorStatus = http.StatusUnauthorized
		return errors.New("비밀번호가 틀렸습니다")
	}
	return nil
}

// 컴퓨터 함수의 존재 유무를 파악하는 함수
func UserLoginCheckComputerNumberFunc(check_user *models.User, messages *string) {
	if check_user.User_computer_number != nil {
		*messages = "기존에 로그인 되어있는 계정입니다. 기존의 계정을 로그아웃 합니다."
	} else {
		*messages = "로그인 되었습니다."
	}
}

// jwt 토큰을 만들고 저장하는 함수
func UserLoginMakeJwtTokenFunc(check_user *models.User, access_token *string, errorStatus *int) (err error) {
	var (
		sub     dtos.Sub
		payload dtos.Payload
	)

	// payload 제작
	sub.Email = check_user.User_email
	sub.Name = check_user.User_name
	sub.User_type = check_user.User_type

	payload.User_id = check_user.User_id
	payload.Sub = sub

	var (
		jwt_secret_keys []string = []string{
			os.Getenv("JWT_ACCESS_SECRET"),
			os.Getenv("JWT_REFRESH_SECRET"),
		}
		jwt_time_strs []string = []string{
			os.Getenv("JWT_ACCESS_TIME"),
			os.Getenv("JWT_REFRESH_TIME"),
		}
		jwt_tokens []string
	)

	for idx, jwt_secret_key := range jwt_secret_keys {

		// jwt의 시간을 파싱하는 로직
		var (
			jwt_time int
		)

		if jwt_time, err = strconv.Atoi(jwt_time_strs[idx]); err != nil {
			fmt.Println("시스템 오류: ", err.Error())
			*errorStatus = http.StatusInternalServerError
			return errors.New("문자열인 시간을 숫자로 파싱하는데 오류가 발생했습니다")
		}

		// jwt 토큰을 만드는 함수
		var (
			jwt_token string
		)

		jwt_token_str := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"exp":     jwt_time,
			"payload": payload,
		})

		if jwt_token, err = jwt_token_str.SignedString([]byte(jwt_secret_key)); err != nil {
			fmt.Println("시스템 오류: ", err.Error())
			*errorStatus = http.StatusInternalServerError
			return errors.New("jwt 토큰을 생성하는데 오류가 발생했습니다")
		}

		jwt_tokens = append(jwt_tokens, jwt_token)

	}

	// 데이터베이스에 킵
	check_user.User_access_token = &jwt_tokens[0]
	check_user.User_refresh_token = &jwt_tokens[1]

	*access_token = jwt_tokens[0]

	return nil
}

// 컴퓨터 넘버를 만들고 데이터 베이스에 저장까지 하는 함수
func UserLoginMakeComputerNumberAndSaveDatabaseFunc(c context.Context, db *gorm.DB, check_user *models.User, computer_number *uuid.UUID, user_type *string, errorStatus *int) error {

	// 컴퓨터 넘버 만들기
	*computer_number = uuid.New()

	// 유저 타입 지정
	*user_type = check_user.User_type

	// 데이터 베이스에 저장하는 로직
	check_user.User_computer_number = computer_number

	result := db.WithContext(c).Save(&check_user)
	if result.Error != nil {
		fmt.Println("시스템 오류: ", result.Error.Error())
		*errorStatus = http.StatusInternalServerError
		return errors.New("데이터 베이스에 인증 데이터 관련을 저장하는데 오류가 발생했습니다")
	}

	return nil
}

// 메인 이미지의 파일 위치를 찾고 보내는 로직
func UserEtcGetMainProfileService(ImageType *dtos.ImageType) error {
	var (
		data_server_path      string = os.Getenv("DATA_FILE_SERVER")
		baseimg_file_path     string = os.Getenv("DATA_FILE_BASE_SERVER")
		main_profile_img_path string = os.Getenv("DATABASE_BASE_MAIN_IMG")
	)

	// 메인 이미지의 위치가 나오는 장소
	img_path := filepath.Join(data_server_path, baseimg_file_path, main_profile_img_path)

	// 이미지를 바이트 형식으로 읽고 base64로 인코딩 하는 로직
	imgData, err := ioutil.ReadFile(img_path)
	if err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return errors.New("메인 이미지를 읽는데 오류가 발생했습니다")
	}

	ImageType.Base64Img = base64.StdEncoding.EncodeToString(imgData)

	// 메인 이미지에 타입을 정하는 로직
	var (
		main_profile_img_path_list []string = strings.Split(main_profile_img_path, ".")
	)
	ImageType.ImgType = main_profile_img_path_list[len(main_profile_img_path_list)-1]

	return nil
}
