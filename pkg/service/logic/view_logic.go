package logic

import (
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/go-bind-mp/pkg/model"
	"github.com/yin-zt/go-bind-mp/pkg/model/request"
	"github.com/yin-zt/go-bind-mp/pkg/model/response"
	"github.com/yin-zt/go-bind-mp/pkg/service/isql"
	"github.com/yin-zt/go-bind-mp/pkg/util/tools"
)

type ViewLogic struct{}

// List 数据列表
func (v ViewLogic) List(c *gin.Context, req interface{}) (data interface{}, resError interface{}) {
	r, ok := req.(*request.ViewListReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	// 获取数据列表
	views, err := isql.View.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取视图列表失败: %s", err.Error()))
	}

	count, err := isql.View.Count()
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取视图总数失败"))
	}

	rets := make([]model.View, 0)
	for _, view := range views {
		rets = append(rets, *view)
	}

	return response.ViewListRsp{
		Total: count,
		Views: rets,
	}, nil
}

// Add 添加视图数据
func (v ViewLogic) Add(c *gin.Context, req interface{}) (data interface{}, resError interface{}) {
	r, ok := req.(*request.ViewAddReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	//if isql.Domain.Exist(tools.H{"doaminrecord": r.DomainRecord, "type": r.Type, "resolution": r.Resolution}) {
	//	return nil, tools.NewValidatorError(fmt.Errorf("相同的域名已存在，请勿重复添加"))
	//}

	zones, err := isql.Zone.GetZonesByIds(r.ZonesId)
	if err != nil {
		return nil, tools.NewValidatorError(fmt.Errorf("根据Zone ID查询zone信息失败"))
	}

	domains, err := isql.Domain.GetDomainsByIds(r.DomainsId)
	if err != nil {
		return nil, tools.NewValidatorError(fmt.Errorf("根据View ID查询域名信息失败"))
	}

	view := model.View{
		ViewName: r.ViewName,
		Acl:      r.Acl,
		Remark:   r.Remark,
		Zones:    zones,
		Domains:  domains,
	}
	err = CommonAddView(&view)
	if err != nil {
		log.Info("添加视图记录失败")
		return nil, tools.NewOperationError(fmt.Errorf("添加视图记录失败" + err.Error()))
	}
	return nil, nil
}

// Delete 删除视图记录
func (v ViewLogic) Delete(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.ViewDeleteReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	for _, id := range r.ViewIds {
		filter := tools.H{"id": int(id)}
		if !isql.View.Exist(filter) {
			return nil, tools.NewMySqlError(fmt.Errorf("视图资源不存在"))
		}
		if isql.View.CheckRelateResource(filter) {
			return nil, tools.NewMySqlError(fmt.Errorf("待删除的视图资源还关联其他domain或zone资源"))
		}
	}

	// 再将视图资源从MySQL中删除
	err := isql.View.Delete(r.ViewIds)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("在MySQL删除视图资源失败: " + err.Error()))
	}
	return nil, nil
}

// Update 更新域名数据
func (v ViewLogic) Update(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.ViewUpdateReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	if !isql.View.Exist(tools.H{"id": r.ID}) {
		return nil, tools.NewMySqlError(fmt.Errorf("该视图资源不存在"))
	}

	// 先获取域名信息
	oldView := new(model.View)
	err := isql.View.Find(tools.H{"id": r.ID}, oldView)
	if err != nil {
		return nil, tools.NewMySqlError(err)
	}

	// 拼装新的域名信息
	view := model.View{
		Model:    oldView.Model,
		ViewName: oldView.ViewName,
		Acl:      oldView.Acl,
		Remark:   r.Remark,
		Zones:    oldView.Zones,
		Domains:  oldView.Domains,
	}

	if err = CommonUpdateView(&view); err != nil {
		return nil, tools.NewOperationError(fmt.Errorf("更新视图资源失败" + err.Error()))
	}
	return nil, nil
}

// CommonUpdateView 标准更新视图资源普通字段
func CommonUpdateView(updateview *model.View) error {
	// 先将视图记录更新到MySQL
	err := isql.View.Update(updateview)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("向MySQL更新视图失败：" + err.Error()))
	}
	return nil
}
