package settings

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func LoadDatabase() {

	var (
		dsn string = os.Getenv("POSTGRES_DATABASE_DSN")
		err error
	)

	if dsn == "" {
		fmt.Println("시스템 오류: 환경변수에 데이터 베이스 환경변수를 지정하지 않았습니다")
		panic("환경변수에 데이터베이스 dsn 소실")
	}

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		panic("데이터 베이스를 연결하는데 오류가 발생했습니다.")
	}

}
