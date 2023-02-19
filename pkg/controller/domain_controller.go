package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/go-bind-mp/pkg/model/request"
	"github.com/yin-zt/go-bind-mp/pkg/service/logic"
)

type DomainController struct{}

func (d *DomainController) List(c *gin.Context) {
	req := new(request.DomainListReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Domain.List(c, req)
	})
}

// Add 增加域名记录
func (d *DomainController) Add(c *gin.Context) {
	req := new(request.DomainAddReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Domain.Add(c, req)
	})
}

// Delete 删除域名记录
func (d *DomainController) Delete(c *gin.Context) {
	req := new(request.DoaminDeleteReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Domain.Delete(c, req)
	})
}

// Update 更新域名记录
func (d *DomainController) Update(c *gin.Context) {
	req := new(request.DomainUpdateReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Domain.Update(c, req)
	})
}
