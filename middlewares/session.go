package middlewares

import (
	"net/http"
	"pp-bakcend/pkg/enums"
	authservice "pp-bakcend/service/auth_service"

	"github.com/gin-gonic/gin"
)

const mid = "admin"

func SessionAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		session, err := ctx.Cookie("session-id")
		if err != nil || !authservice.CheckSession(mid, session) {
			ctx.JSON(http.StatusOK, gin.H{
				"code": enums.ERROR_AUTH_FAILED,
				"msg":  "",
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
