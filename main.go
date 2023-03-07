package main

import (
	"fmt"
	"net/http"
	"pp-bakcend/models"
	"pp-bakcend/pkg/gredis"
	"pp-bakcend/pkg/logging"
	"pp-bakcend/pkg/setting"
	"pp-bakcend/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	setting.Setup()
	logging.Setup()
	models.Setup()
	gredis.Setup()

}

func main() {
	// gin.SetMode()
	gin.SetMode(setting.ServerSetting.RunMode)
	router := routers.InitRouter()
	listen := fmt.Sprintf("%s", setting.ServerSetting.HttpPort)
	logging.Info("start listen on: ", listen)

	server := &http.Server{
		Addr:           listen,
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
