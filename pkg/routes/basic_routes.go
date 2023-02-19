package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/go-bind-mp/pkg/controller"
)

// 注册基础路由
func InitBaseRoutes(r *gin.RouterGroup) gin.IRoutes {
	base := r.Group("/base")
	{
		base.GET("ping", controller.Demo)
		//base.GET("getpasswd", controller.Base.GetPasswd) // 将明文字符串转为MySQL识别的密码
		//// 登录登出刷新token无需鉴权 登录前需要先将密码转换为mysql识别的密码
		//base.POST("/login", authMiddleware.LoginHandler)
		//base.POST("/logout", authMiddleware.LogoutHandler)
		//base.POST("/refreshToken", authMiddleware.RefreshHandler)
		//base.GET("/dashboard", controller.Base.Dashboard) // 系统首页展示数据
	}
	return r
}
