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
			origin         string   = ctx.GetHeader("Origin")
			isAllowed      bool     = false
			backend_hosts  []string = strings.Split(os.Getenv("GO_BACKEND_HOSTS"), ",")
			frontend_hosts []string = strings.Split(os.Getenv("GO_FRONTEND_HOSTS"), ",")
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
			ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin")
			ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")
		}

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}

		ctx.Next()
	}
}
