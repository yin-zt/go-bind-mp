package logic

import (
	"github.com/yin-zt/go-bind-mp/pkg/model/request"
	"github.com/yin-zt/go-bind-mp/pkg/util/tools"

	"github.com/gin-gonic/gin"
)

type BaseLogic struct{}

// GetPasswd
func (l BaseLogic) GetPasswd(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.GetPasswdReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	return tools.NewGenPasswd(r.Passwd), nil
}
