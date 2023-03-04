package v1

import (
	"pp-bakcend/pkg/app"
	"pp-bakcend/pkg/enums"
	"pp-bakcend/pkg/logging"
	loginfoservice "pp-bakcend/service/loginfo_service"

	"github.com/gin-gonic/gin"
)

func GetAllLoginInfo(c *gin.Context) {
	App := app.Gin{C: c}

	result, err := loginfoservice.GetAllLoginInfo()
	if err != nil {
		logging.Error(err)
		App.Response(200, enums.ERROR, "")
		return
	}
	App.Response(200, enums.SUCCESS, result)
}
