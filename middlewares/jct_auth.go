package middlewares

import (
	"net/http"
	"pp-bakcend/pkg/enums"
	authservice "pp-bakcend/service/auth_service"

	"github.com/gin-gonic/gin"
)

func JctAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		mid, err := c.Cookie("mid")
		crypto, err := c.Cookie("crypto")
		hostname, err := c.Cookie("hostname")

		if err != nil || !authservice.CheckAuthWithHostname(mid, crypto, hostname) {
			c.JSON(http.StatusOK, gin.H{
				"code": enums.ERROR_AUTH_FAILED,
				"msg":  "",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}
