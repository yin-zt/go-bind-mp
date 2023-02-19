package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/go-bind-mp/pkg/controller"
)

func InitViewRoutes(r *gin.RouterGroup) gin.IRoutes {
	view := r.Group("/view")
	{
		view.GET("/list", controller.View.List)
		view.POST("/add", controller.View.Add)
		view.POST("/delete", controller.View.Delete)
		view.POST("/update", controller.View.Update)
	}
	return r
}
