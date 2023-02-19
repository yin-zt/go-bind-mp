package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/go-bind-mp/pkg/controller"
)

func InitDomainRoutes(r *gin.RouterGroup) gin.IRoutes {
	domain := r.Group("/domain")
	{
		domain.GET("/list", controller.Domain.List)
		domain.POST("/add", controller.Domain.Add)
		domain.POST("/delete", controller.Domain.Delete)
		domain.POST("/update", controller.Domain.Update)
	}
	return r
}
