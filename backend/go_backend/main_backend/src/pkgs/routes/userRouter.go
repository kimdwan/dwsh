package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kimdwan/dwsh/src/middlewares"
	"github.com/kimdwan/dwsh/src/pkgs/controllers"
)

func UserRouter(router *gin.Engine) {
	userrouter := router.Group("user")

	// 기본 유저 라우터
	userrouter.POST("signup", middlewares.UserCheckUserTypeMiddleware(), controllers.UserSignUpController)
	userrouter.POST("login", controllers.UserLoginController)
}
