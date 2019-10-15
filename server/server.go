package server

import (
	"github.com/gin-gonic/gin"
	"mygo/service/v1"
	"os"
)


/**
路由设置并开启gin服务
 */
func init(){

	port := os.Getenv("SERVER_PORT")
	server := gin.Default()

	// 心跳测试
	server.GET("ping",v1.Ping)

	// 发送短信
	server.GET("sendphone",v1.PhoneCaptcha)

	// 注册用户
	server.POST("registeruser",v1.UserRegister)

	err := server.Run(port)
	if err != nil {
		panic(err.Error())
	}
}