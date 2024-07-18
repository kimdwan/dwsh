package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kimdwan/dwsh/settings"
	"github.com/kimdwan/dwsh/src/middlewares"
	"github.com/kimdwan/dwsh/src/pkgs/routes"
)

func init() {
	settings.LoadDotenv()
	settings.LoadDatabase()
}

func main() {
	var (
		port string = os.Getenv("GO_PORT")
	)

	if port == "" {
		fmt.Println("환경변수에 포트 번호를 입력하지 않았습니다.")
		panic("환경변수에 포트번호 입력 필요")
	}

	router := gin.Default()

	// 미들웨어 목록
	router.Use(gin.Logger())
	router.Use(middlewares.CorsMiddleware())

	// 라우터 목록
	routes.UserRouter(router)

	router.Run(port)
}
