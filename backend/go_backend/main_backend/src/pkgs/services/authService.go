package services

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kimdwan/dwsh/settings"
	"github.com/kimdwan/dwsh/src/dtos"
	"github.com/kimdwan/dwsh/src/models"
	"gorm.io/gorm"
)

// auth 라우터를 통해 생성된 payload를 뿌려주는 서비스
func AuthParsePayloadService(ctx *gin.Context) (*dtos.Payload, error) {
	var (
		payload_byte string = ctx.GetString("payload_byte")
		payload      dtos.Payload
		err          error
	)

	if payload_byte == "" {
		return nil, errors.New("미들웨어에서 보낸 payload_byte가 존재하지 않습니다")
	}

	if err = json.Unmarshal([]byte(payload_byte), &payload); err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return nil, errors.New("payload byte를 역직렬화 하는데 오류가 발생했습니다")
	}

	return &payload, nil
}

// auth 라우터가 보내는 값이 제대로 작동을 하는지 확인하는 서비스
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

// 이미지 저장 위치를 파악하고 프로필 이미지를 제공하는 함수
func AuthGetProfileService(c context.Context, payload *dtos.Payload, image_types *dtos.ImageType) (int, error) {
	var (
		profile_name string
		err          error
	)

	// 데이터 베이스에서 유저의 프로필 위치 정보 찾기
	if err = AuthGetProfileGetUserProfileImgNameFunc(c, payload, &profile_name); err != nil {
		return http.StatusInternalServerError, err
	}

	// 프로필 존재 유무를 확인
	if profile_name == "" {
		return 0, nil
	}

	// 프로필 사진을 읽고 보냄
	if err = AuthGetProfileFindProfileImgFunc(payload, &profile_name, image_types); err != nil {
		return http.StatusNotFound, err
	}

	return 0, nil

}

// 데이터 베이스에서 유저의 프로필 이름정보를 찾고 저장
func AuthGetProfileGetUserProfileImgNameFunc(c context.Context, payload *dtos.Payload, proflie_name *string) error {
	var (
		db   *gorm.DB = settings.DB
		user models.User
	)

	result := db.WithContext(c).Where("user_id = ?", payload.User_id).First(&user)
	if result.Error != nil {
		fmt.Println("시스템 오류: ", result.Error.Error())
		return errors.New("데이터 베이스에서 user_id에 해당하는 user 테이블을 찾는데 오류가 발생했습니다")
	}

	if user.User_pf_img != nil {
		*proflie_name = *user.User_pf_img
	}

	return nil
}

// 서버에서 유저가 저장한 파일을 찾고 프로필 이미지 찾기
func AuthGetProfileFindProfileImgFunc(payload *dtos.Payload, profile_name *string, image_types *dtos.ImageType) error {
	var (
		server_name         string = os.Getenv("DATA_FILE_SERVER")
		profile_server_name string = os.Getenv("DATA_FILE_PROFILE_IMG_SERVER")
	)

	// 프로필이 저장된 위치를 찾아냄
	user_proflie_server_name := filepath.Join(server_name, profile_server_name, payload.User_id.String(), *profile_name)

	// 프로필 이미지를 읽음
	img_byte, err := os.ReadFile(user_proflie_server_name)
	if err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		return errors.New("유저의 프로필 이름에 해당하는 프로필 이미지를 찾는데 오류가 발생했습니다")
	}

	image_types.Base64Img = base64.StdEncoding.EncodeToString(img_byte)

	// 타입을 읽고 보내줌
	var (
		profile_img_list []string = strings.Split(*profile_name, ".")
	)
	image_types.ImgType = profile_img_list[len(profile_img_list)-1]

	return nil
}

// 유저를 로그아웃 하는 서비스
func AuthUserLogoutService(payload *dtos.Payload) (int, error) {

	// 유저의 인증과 관련된 데이터를 수정함
	if err := AuthUserLogoutRemoveJwtTokenFunc(payload); err != nil {
		return http.StatusInternalServerError, err
	}

	return 0, nil
}

// 유저 데이터베이스에서 유저의 jwt 토큰이나 인증 토큰을 없앰
func AuthUserLogoutRemoveJwtTokenFunc(payload *dtos.Payload) error {
	var (
		db   *gorm.DB = settings.DB
		user models.User
	)
	c, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	// 데이터베이스에서 user_id에 해당하는 데이터를 찾고 jwt 토큰과 computer number를 없앰
	result := db.WithContext(c).Where("user_id = ?", payload.User_id).First(&user)
	if result.Error != nil {
		fmt.Println("시스템 오류: ", result.Error.Error())
		return errors.New("데이터 베이스에서 user_id에 해당하는 유저 테이블을 찾는데 오류가 발생했습니다")
	}

	*user.User_access_token = ""
	*user.User_refresh_token = ""
	user.User_computer_number = nil

	// 유저 데이터에서 테이블을 업로드 하는데 오류가 발생했습니다.
	if result = db.WithContext(c).Save(&user); result.Error != nil {
		fmt.Println("시스템 오류: ", result.Error.Error())
		return errors.New("데이터 베이스에 유저 테이블을 업로드 하는데 오류가 발생했습니다")
	}

	return nil
}
