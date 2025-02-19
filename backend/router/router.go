package router

import (
	"github.com/cloudscode/via-svit/backend/api"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 使用 Gin 官方 CORS 中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://wails.localhost:34115"}, // 允许的前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 订单模块路由
	orderGroup := r.Group("/orders")
	{
		orderGroup.GET("", api.GetOrders)
		orderGroup.POST("", api.CreateOrder)
		orderGroup.PUT("", api.UpdateOrder)
		orderGroup.DELETE("/:id", api.DeleteOrder)
	}

	// 如有其他模块，可继续添加…

	return r
}
