package v1

import (
	"pp-bakcend/pkg/app"
	"pp-bakcend/pkg/logging"
	userservice "pp-bakcend/service/user_service"

	"github.com/gin-gonic/gin"
)

type SetUserSponsorForm struct {
	Mid     string `json:"mid" binding:"required"`
	Sponsor int64  `json:"sponsor" binding:"required"`
}

func SetUserSponsor(c *gin.Context) {
	App := app.Gin{C: c}

	var form SetUserSponsorForm
	var err error
	if err = c.ShouldBindJSON(&form); err != nil {
		App.Response(200, 400, "")
		logging.Info(err)
		return
	}
	err = userservice.SetUserSponsor(form.Mid, form.Sponsor)
	if err != nil {
		App.Response(200, 400, "")
		logging.Info(err)
		return
	}
	App.Response(200, 200, "")
}
