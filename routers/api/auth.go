package api

import (
	"pp-bakcend/models"
	"pp-bakcend/pkg/app"
	"pp-bakcend/pkg/enums"
	"pp-bakcend/pkg/logging"
	"pp-bakcend/pkg/setting"
	authservice "pp-bakcend/service/auth_service"
	userservice "pp-bakcend/service/user_service"
	"pp-bakcend/utils"

	"github.com/gin-gonic/gin"
)

type AppLoginForm struct {
	Name    string `json:"name" binding:"required" `
	Avatar  string `json:"avatar" binding:"required"`
	Jct     string `json:"jct" binding:"required"`
	Version string `json:"version" binding:"required"`
}

func AppLogin(c *gin.Context) {
	App := app.Gin{C: c}
	var form AppLoginForm
	var err error
	if err = c.ShouldBindJSON(&form); err != nil {
		App.Response(200, 400, "")
		logging.Info(err)
		return
	}
	user := models.Users{}
	user.Mid, err = c.Cookie("mid")
	cryptStr, err := c.Cookie("crypto")
	hostname, err := c.Cookie("hostname")
	if err != nil || user.Mid == "" {
		logging.Error(err)
		App.Response(200, enums.ERROR_AUTH_FAILED, "")
		return
	}
	user.Avatar = form.Avatar
	user.Name = form.Name

	// if no user ,create
	userservice.CreateIfNotExist(&user)

	if authservice.CheckAuth(user.Mid, hostname, form.Jct, cryptStr, form.Version) != true {
		App.Response(200, enums.ERROR_AUTH_FAILED, "")
		return
	}

	err = userservice.Update(&user)
	if err != nil {
		App.Response(200, enums.ERROR, "")
		logging.Error(err)
		return
	}
	App.Response(200, enums.SUCCESS, "")

}

type BrowserLoginForm struct {
	Password string `json:"password" binding:"required"`
}

func BrowserLogin(c *gin.Context) {
	App := app.Gin{C: c}
	var form BrowserLoginForm
	var err error
	if err = c.ShouldBindJSON(&form); err != nil {
		App.Response(200, 400, "")
		logging.Info(err)
		return
	}
	if !utils.IsEqual(form.Password, setting.AppSetting.AdminPass) {
		App.Response(200, enums.ERROR_AUTH_FAILED, "")
		return
	}
	// add session
	session, err := authservice.AddSession("admin", form.Password)
	if err != nil {
		App.Response(200, enums.ERROR_AUTH_FAILED, "")
		logging.Error(err)
		return
	}
	c.SetCookie("session-id", session, 3600*24*7, "/", "", false, false)
	App.Response(200, enums.SUCCESS, "")
}
