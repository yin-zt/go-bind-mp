package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/go-bind-mp/pkg/model/request"
	"github.com/yin-zt/go-bind-mp/pkg/service/logic"
)

type ZoneController struct{}

func (z *ZoneController) List(c *gin.Context) {
	req := new(request.ZoneListReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Zone.List(c, req)
	})
}

func (z *ZoneController) Add(c *gin.Context) {
	req := new(request.ZoneDeleteReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Zone.Delete(c, req)
	})
}
