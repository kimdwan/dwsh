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

	// 인증 후 유저 기반의 서비스
	authuserrouter := authrouter.Group("user")
	authuserrouter.GET("get/profile", controllers.AuthGetProfileController)
	authuserrouter.GET("logout", controllers.AuthUserLogoutController)
}
