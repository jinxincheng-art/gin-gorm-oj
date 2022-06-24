package router

import (
	_ "gin-gorm-oj/docs"
	"gin-gorm-oj/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	// Swagger 配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// 路由规则
	r.GET("/ping", service.Ping)
	r.GET("/problem-list", service.GetProblemList)
	r.GET("/problem-detail",service.GetProblemDetail)
	return r
}
