package models

import (
	"errors"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	User_id            uuid.UUID  `gorm:"type:uuid;unique;"`
	User_type          string     `gorm:"type:varchar(255);not null;"`
	User_email         string     `gorm:"type:varchar(255);unique;not null;"`
	User_hash          string     `gorm:"type:varchar(255);not null;"`
	User_name          string     `gorm:"type:varchar(255);not null;"`
	User_pf_img        *string    `gorm:"type:varchar(255);"`
	User_nickname      *string    `gorm:"type:varchar(255);unique;"`
	User_birthday      *time.Time `gorm:"type:date;"`
	User_secretkey     *string    `gorm:"type:varchar(255);unique;"`
	User_term_agree_1  bool       `gorm:"type:boolean;default:false;not null;"`
	User_term_agree_2  bool       `gorm:"type:boolean;default:false;not null"`
	User_term_agree_3  bool       `gorm:"type:boolean;not null"`
	User_access_token  *string    `gorm:"type:varchar(255);unique"`
	User_refresh_token *string    `gorm:"type:varchar(255);unique"`
}

// 데이터 생성시 확인하는 로직
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {

	// user_id에 기본값을 부여하는 로직
	u.User_id = uuid.New()

	// 유저의 타입을 확인하는 로직
	var (
		isUserTypeAllowed bool     = false
		user_types        []string = strings.Split(os.Getenv("DATABASE_USER_TYPE"), ",")
	)

	for _, user_type := range user_types {
		if u.User_type == user_type {
			isUserTypeAllowed = true
			break
		}
	}

	if !isUserTypeAllowed {
		var (
			userTypeString string
		)

		for idx, user_type := range user_types {
			userTypeString += user_type
			if idx != (len(user_types) - 1) {
				userTypeString += ", "
			}
		}

		return errors.New("유저의 타입은 " + userTypeString + "중 한개만 가능합니다")
	}

	// 유저의 이메일을 확인하는 로직
	var (
		matchEmailString string = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		isEmailAllowed   bool   = false
	)

	re := regexp.MustCompile(matchEmailString)
	isEmailAllowed = re.MatchString(u.User_email)

	if !isEmailAllowed {
		return errors.New("이메일 형식에 맞지 않습니다")
	}

	// 프로필 이미지 형식을 확인하는 로직
	if u.User_pf_img != nil {
		var (
			imgstring_types     []string = strings.Split(os.Getenv("DATABASE_IMG_TYPE"), ",")
			userImgNames        []string = strings.Split(*u.User_pf_img, ".")
			userImgType         string   = userImgNames[len(userImgNames)-1]
			isProfileImgAllowed bool     = false
		)

		for _, imgstring_type := range imgstring_types {
			if userImgType == imgstring_type {
				isProfileImgAllowed = true
				break
			}
		}

		if !isProfileImgAllowed {
			var (
				imgTypeErrorMsg string
			)

			for idx, imgstring_type := range imgstring_types {
				imgTypeErrorMsg += imgstring_type

				if idx != len(imgstring_types)-1 {
					imgTypeErrorMsg += ", "
				}

			}

			return errors.New("저장할 수 있는 이미지 타입은 " + imgTypeErrorMsg + " 입니다.")
		}
	}

	// 년월일 확인하는 로직필요
	if u.User_birthday != nil {
		var (
			userBirthdayDate    string = u.User_birthday.Format("2006-01-02")
			matchBirthDayString string = `^\d{4}-\d{2}-\d{2}$`
			isBirthDayAllowed   bool   = false
		)

		re := regexp.MustCompile(matchBirthDayString)

		isBirthDayAllowed = re.MatchString(userBirthdayDate)

		if !isBirthDayAllowed {
			return errors.New("생년월일 형식에 맞지 않습니다 생년월일 형식은 YYYY-MM-DD여야 합니다")
		}

	}

	// 방문자가 아닐경우 환경변수에 있는 값인지 확인하는 로직
	if u.User_type != "방문자" {
		var (
			user_secretkeys        []string = strings.Split(os.Getenv("DS_USER_SECRET_KEY"), ",")
			isUserSecretKeyAllowed bool     = false
		)

		if u.User_secretkey != nil {

			for _, user_secretkey := range user_secretkeys {
				if *u.User_secretkey == user_secretkey {
					isUserSecretKeyAllowed = true
				}
			}

			if !isUserSecretKeyAllowed {
				var (
					secretKeyErrorMsg string
				)

				for idx, user_secretkey := range user_secretkeys {
					secretKeyErrorMsg += user_secretkey

					if idx != len(user_secretkeys)-1 {
						secretKeyErrorMsg += ", "
					}
				}

				return errors.New("user의 시크릿 키는 " + secretKeyErrorMsg + " 중 하나여야 합니다")
			}
		}
	}

	// 유저의 필수 동의 여부를 파악하는 로직

	if !u.User_term_agree_1 || !u.User_term_agree_2 {
		return errors.New("1번과 2번은 필수 동의 사항입니다")
	}

	return nil

}

// 업데이트에 확인해야 하는 로직

func (u *User) BeforeSave(tx *gorm.DB) (err error) {

	// 유저의 타입 확인
	var (
		user_types        []string = strings.Split(os.Getenv("DATABASE_USER_TYPE"), ",")
		isUserTypeAllowed bool     = false
	)

	for _, user_type := range user_types {
		if u.User_type == user_type {
			isUserTypeAllowed = true
			break
		}
	}

	if !isUserTypeAllowed {
		var (
			userTypeErrorMsg string
		)

		for idx, user_type := range user_types {
			userTypeErrorMsg += user_type

			if idx != len(user_types)-1 {
				userTypeErrorMsg += ", "
			}
		}

		return errors.New("유저의 타입은 " + userTypeErrorMsg + " 중 하나여야 합니다")
	}

	// 유저의 이메일 로직 확인
	var (
		emailMatchString string = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		isEmailAllowed   bool   = false
	)

	re := regexp.MustCompile(emailMatchString)
	isEmailAllowed = re.MatchString(u.User_email)
	if !isEmailAllowed {
		return errors.New("이메일 타입을 유지해 주세요")
	}

	// 프로필 이미지 타입 확인
	if u.User_pf_img != nil {
		var (
			img_types           []string = strings.Split(os.Getenv("DATABASE_IMG_TYPE"), ",")
			isProfileImgAllowed bool     = false
		)

		for _, img_type := range img_types {
			if *u.User_pf_img == img_type {
				isProfileImgAllowed = true
				break
			}
		}

		if !isProfileImgAllowed {
			var (
				profileImgErrorMsg string
			)

			for idx, img_type := range img_types {
				profileImgErrorMsg += img_type

				if idx != len(img_types)-1 {
					profileImgErrorMsg += ", "
				}
			}

			return errors.New("유저가 프로필 이미지를 저장할수 있는 타입은 " + profileImgErrorMsg + " 중 하나입니다")

		}
	}

	// 생년월일 확인 로직
	if u.User_birthday != nil {
		var (
			user_birthday_string string = u.User_birthday.Format("2006-01-02")
			birthdayMatchString  string = `^\d{4}-\d{2}-\d{2}$`
			isBirthDayAllowed    bool   = false
		)

		re := regexp.MustCompile(birthdayMatchString)

		isBirthDayAllowed = re.MatchString(user_birthday_string)

		if !isBirthDayAllowed {
			return errors.New("생년월일의 타입은 YYYY-MM-DD 입니다")
		}
	}

	// 유저의 시크릿 키를 확인하는 로직
	if u.User_type != "방문자" {
		var (
			isSecretKeyAllowed bool     = false
			user_secretkeys    []string = strings.Split(os.Getenv("DS_USER_SECRET_KEY"), ",")
		)

		if u.User_secretkey != nil {
			for _, user_secretkey := range user_secretkeys {
				if *u.User_secretkey == user_secretkey {
					isSecretKeyAllowed = true
					break
				}
			}
		}

		if !isSecretKeyAllowed {
			var (
				secretkeyErrorMsg string
			)

			for idx, user_secretkey := range user_secretkeys {
				secretkeyErrorMsg += user_secretkey

				if idx != len(user_secretkeys)-1 {
					secretkeyErrorMsg += ", "
				}

			}

			return errors.New("유저의 시크릿 키는 " + secretkeyErrorMsg + " 중 하나입니다")
		}
	}

	// 필수 동의사항 체크 여부 확인 로직
	if !u.User_term_agree_1 || !u.User_term_agree_2 {
		return errors.New("유저는 1번과 2번 동의사항은 필수로 작성해야 합니다")
	}

	return nil

}

func (User) TableName() string {
	return "DS_USER_TB"
}
