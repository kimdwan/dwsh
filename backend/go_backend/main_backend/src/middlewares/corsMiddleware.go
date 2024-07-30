package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var (
			origin         = ctx.GetHeader("Origin")
			isAllowed      = false
			backend_hosts  = strings.Split(os.Getenv("GO_BACKEND_HOSTS"), ",")
			frontend_hosts = strings.Split(os.Getenv("GO_FRONTEND_HOSTS"), ",")
		)

		allowed_hosts := append(backend_hosts, frontend_hosts...)

		for _, allowed_host := range allowed_hosts {
			if allowed_host == origin {
				isAllowed = true
				break
			}
		}

		if isAllowed {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
			ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, X-Requested-With, Authorization, User-Computer-Number")
			ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		}

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		ctx.Next()
	}
}
