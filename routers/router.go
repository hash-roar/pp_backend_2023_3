package routers

import (
	"pp-bakcend/middlewares"
	"pp-bakcend/routers/api"
	v1 "pp-bakcend/routers/api/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.DefaultLogger())
	conf := cors.DefaultConfig()
  conf.AllowAllOrigins = true
	conf.AllowCredentials = true
	r.Use(cors.New(conf))

	r.POST("/alogin", api.AppLogin)
	r.POST("/blogin", api.BrowserLogin)

	netToServerRouter := r.Group("net").Use(middlewares.SessionAuth())
	{
		netToServerRouter.POST("loginfo", v1.GetAllLoginInfo)
		netToServerRouter.POST("shield", v1.GetAllBlockWords)
		netToServerRouter.POST("set-visible", v1.SetWordVisibility)
	}

	appToServerRouter := r.Group("app").Use(middlewares.JctAuth())
	{
		appToServerRouter.POST("add-words", v1.HandleBlockWords)
		appToServerRouter.POST("get-words", v1.UserGetBlockWord)
	}

	return r
}
