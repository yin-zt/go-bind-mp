package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/go-bind-mp/pkg/controller"
)

func InitZoneRoutes(r *gin.RouterGroup) gin.IRoutes {
	domain := r.Group("/zone")
	{
		domain.GET("/list", controller.Zone.List)
		domain.POST("/add", controller.Zone.Add)
		domain.POST("/delete", controller.Zone.Delete)
	}
	return r
}
