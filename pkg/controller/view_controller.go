package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/go-bind-mp/pkg/model/request"
	"github.com/yin-zt/go-bind-mp/pkg/service/logic"
)

type ViewController struct{}

func (v *ViewController) List(c *gin.Context) {
	req := new(request.ViewListReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.View.List(c, req)
	})
}

// Add 增加视图记录
func (v *ViewController) Add(c *gin.Context) {
	req := new(request.ViewAddReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.View.Add(c, req)
	})
}

// Delete 删除视图记录
func (v *ViewController) Delete(c *gin.Context) {
	req := new(request.ViewDeleteReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.View.Delete(c, req)
	})
}

// Update 更新视图记录
func (v *ViewController) Update(c *gin.Context) {
	req := new(request.ViewUpdateReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.View.Update(c, req)
	})
}
