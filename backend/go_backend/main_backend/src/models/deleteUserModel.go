package models

import (
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DeleteUser struct {
	gorm.Model
	// 유저 아이디가 실제로 존재하는지 확인하는 로직 필요
	User_id uuid.UUID `gorm:"type:uuid;unique;not null"`
	// 유저의 타입이 지정한 환경변수 안에 있어야 함을 확인할 함수 만들어야 함
	User_type string `gorm:"type:varchar(255);not null;"`
	// 이메일 타입을 확인 하는 로직 필요
	User_email    string     `gorm:"not null;"`
	User_name     string     `gorm:"type:varchar(255);not null;"`
	User_nickname *string    `gorm:"type:varchar(255);"`
	User_birthday *time.Time `gorm:"type:date;"`
	// 조건의 1번과 2번이 제대로 저장 되었는지 확인 필요
	User_term_agree_1 bool   `gorm:"type:boolean;default:false;not null;"`
	User_term_agree_2 bool   `gorm:"type:boolean;default:false;not null;"`
	User_term_agree_3 bool   `gorm:"type:boolean;not null;"`
	User_feedback     string `gorm:"type:varchar(500);not null;"`
}

func (d *DeleteUser) BeforeCreate(tx *gorm.DB) error {

	// 유저가 존재하는지 확인하는 로직
	var (
		user User
	)
	c, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	result := tx.WithContext(c).Where("user_id = ?").First(&user)
	if result.Error != nil {
		fmt.Println("시스템 오류: ", result.Error.Error())
		return errors.New("데이터 베이스에서 유저 아이디와 동일한 유저를 찾는데 오류가 발생했습니다")
	}

	// 유저의 타입이 제대로 저장이 됬는지 확인하는 로직
	var (
		user_types        []string = strings.Split(os.Getenv("DATABASE_USER_TYPE"), ",")
		isUserTypeAllowed bool     = false
	)

	for _, user_type := range user_types {
		if user_type == d.User_type {
			isUserTypeAllowed = true
			break
		}
	}

	if !isUserTypeAllowed {
		var (
			userTypeErrorMsg string = "저장 가능한 유저 타입은 "
		)

		for idx, user_type := range user_types {
			userTypeErrorMsg += user_type

			if idx != 2 {
				userTypeErrorMsg += ", "
			}
		}

		fmt.Println("시스템 오류: 유저의 타입이 잘못되었음")
		return errors.New(userTypeErrorMsg + " 입니다")
	}

	// 이메일 형식을 확인하는 로직
	var (
		isEmailAllowed   bool   = false
		matchEmailString string = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	)

	re := regexp.MustCompile(matchEmailString)
	isEmailAllowed = re.MatchString(d.User_email)

	if !isEmailAllowed {
		fmt.Println("시스템 오류: 이메일 형식에 맞지 않음")
		return errors.New("이메일 형식에 맞지 않습니다")
	}

	// 이용약간의 동의여부를 파악하는 로직
	if !d.User_term_agree_1 || !d.User_term_agree_2 {
		fmt.Println("시스템 오류: 유저의 필수 동의 여부가 되어있지 않음")
		return errors.New("유저의 필수 동의여부가 되어있지 않습니다")
	}

	return nil
}

func (DeleteUser) TableName() string {
	return "DS_DELETE_USER_TB"
}
