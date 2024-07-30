package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kimdwan/dwsh/src/middlewares"
	"github.com/kimdwan/dwsh/src/pkgs/controllers"
)

func AuthRouter(router *gin.Engine) {
	authrouter := router.Group("auth")

	// jwt 토큰 인증 미들웨어
	authrouter.Use(middlewares.AuthCheckJwtMiddleware())

	// test 라우터
	authrouter.GET("test", controllers.AuthTestController)
}
