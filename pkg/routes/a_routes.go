package routes

import (
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/go-bind-mp/pkg/config"
)

func InitRoutes() *gin.Engine {
	//设置模式
	gin.SetMode(config.Conf.System.Mode)

	// 创建带有默认中间件的路由:
	// 日志与恢复中间件
	r := gin.Default()

	// 路由分组
	apiGroup := r.Group("/" + config.Conf.System.UrlPathPrefix)

	// 注册路由
	InitBaseRoutes(apiGroup)   // 注册基础路由, 不需要jwt认证中间件,不需要casbin中间件
	InitDomainRoutes(apiGroup) // 注册域名路由
	InitViewRoutes(apiGroup)   // 注册视图路由
	InitZoneRoutes(apiGroup)   // 注册Zone路由

	log.Info("初始化路由完成！")
	return r
}
