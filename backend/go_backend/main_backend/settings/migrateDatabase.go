package settings

import (
	"fmt"

	"github.com/kimdwan/dwsh/src/models"
)

func MigrateDatabase() {

	if err := DB.AutoMigrate(&models.User{}, &models.DeleteUser{}); err != nil {
		fmt.Println("시스템 오류: ", err.Error())
		panic("데이터 베이스를 마이그레이션 하는데 오류가 발생했습니다.")
	}

}
