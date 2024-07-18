package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kimdwan/dwsh/src/pkgs/controllers"
)

func UserRouter(router *gin.Engine) {
	userrouter := router.Group("user")

	// 기본 유저 회원가입 경로
	userrouter.POST("signup", controllers.UserSignUpController)
}
