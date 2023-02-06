package main

import (
	"go_deliver/controller"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("/api/deliver", controller.MakeList)
	return r
}
