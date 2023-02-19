package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/go-bind-mp/pkg/model/request"
	"github.com/yin-zt/go-bind-mp/pkg/service/logic"
)

type BaseController struct{}

// GetPasswd 生成加密密码
func (m *BaseController) GetPasswd(c *gin.Context) {
	req := new(request.GetPasswdReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Base.GetPasswd(c, req)
	})
}
