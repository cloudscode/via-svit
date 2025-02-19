package router

import (
	"github.com/cloudscode/via-svit/backend/api"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

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
