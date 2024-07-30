package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kimdwan/dwsh/src/middlewares"
	"github.com/kimdwan/dwsh/src/pkgs/controllers"
)

func UserRouter(router *gin.Engine) {
	// 테스트를 위한 라우터
	testrouter := router.Group("test")
	testrouter.GET("one", controllers.UserTestController)

	// 유저 개인을 위한 라우터
	userrouter := router.Group("user")

	// 기본 유저 라우터
	userrouter.POST("signup", middlewares.UserCheckUserTypeMiddleware(), controllers.UserSignUpController)
	userrouter.POST("login", controllers.UserLoginController)

	// 유저에게 집적 적인 서비스라고 불리기는 애매한 라우터
	etcrouter := router.Group("etc")
	etcrouter.GET("main/profile", controllers.UserEtcGetMainProfileController) // 메인 이미지를 전송하는 로직
	etcrouter.GET("main/logo", controllers.UserEtcGetMainLogoController)       // 메인 로고를 전송하는 로직
}
