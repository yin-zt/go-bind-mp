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

type DomainLogic struct{}

// List 数据列表
func (d DomainLogic) List(c *gin.Context, req interface{}) (data interface{}, resError interface{}) {
	r, ok := req.(*request.DomainListReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	// 获取数据列表
	domains, err := isql.Domain.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取域名列表失败: %s", err.Error()))
	}

	count, err := isql.Domain.Count()
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取域名总数失败"))
	}

	rets := make([]model.Domain, 0)
	for _, domain := range domains {
		rets = append(rets, *domain)
	}

	return response.DomainListRsp{
		Total:   count,
		Domains: rets,
	}, nil
}

// Add 添加域名数据
func (d DomainLogic) Add(c *gin.Context, req interface{}) (data interface{}, resError interface{}) {
	r, ok := req.(*request.DomainAddReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	//if isql.Domain.Exist(tools.H{"doaminrecord": r.DomainRecord, "type": r.Type, "resolution": r.Resolution}) {
	//	return nil, tools.NewValidatorError(fmt.Errorf("相同的域名已存在，请勿重复添加"))
	//}

	zones, err := isql.Zone.GetZonesByIds(r.BelongZoneId)
	if err != nil {
		return nil, tools.NewValidatorError(fmt.Errorf("根据Zone ID查询zone信息失败"))
	}

	views, err := isql.View.GetViewsByIds(r.BelongViewId)
	if err != nil {
		return nil, tools.NewValidatorError(fmt.Errorf("根据View ID查询view信息失败"))
	}

	domain := model.Domain{
		DomainRecord: r.DomainRecord,
		Type:         r.Type,
		Resolution:   r.Resolution,
		Status:       "2",
		Monitor:      "1",
		Protocol:     "TCP",
		Port:         80,
		Notify:       "mailong",
		BelongSystem: "李白系统",
		BelongZone:   zones,
		BelongView:   views,
	}
	err = CommonAddDomain(&domain)
	if err != nil {
		log.Info("添加域名记录失败")
		return nil, tools.NewOperationError(fmt.Errorf("添加域名记录失败" + err.Error()))
	}
	return nil, nil
}

// Delete 删除域名记录
func (d DomainLogic) Delete(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.DoaminDeleteReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	for _, id := range r.DomainIds {
		filter := tools.H{"id": int(id)}
		if !isql.Domain.Exist(filter) {
			return nil, tools.NewMySqlError(fmt.Errorf("域名不存在"))
		}
	}

	// 再将域名从MySQL中删除
	err := isql.Domain.Delete(r.DomainIds)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("在MySQL删除域名失败: " + err.Error()))
	}
	return nil, nil
}

// Update 更新域名数据
func (d DomainLogic) Update(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.DomainUpdateReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	if !isql.Domain.Exist(tools.H{"id": r.ID}) {
		return nil, tools.NewMySqlError(fmt.Errorf("该记录不存在"))
	}

	// 先获取域名信息
	oldDomain := new(model.Domain)
	err := isql.Domain.Find(tools.H{"id": r.ID}, oldDomain)
	if err != nil {
		return nil, tools.NewMySqlError(err)
	}

	// 拼装新的域名信息
	domain := model.Domain{
		Model:        oldDomain.Model,
		DomainRecord: oldDomain.DomainRecord,
		Type:         oldDomain.Type,
		Resolution:   oldDomain.Resolution,
		Status:       oldDomain.Status,
		Monitor:      r.Monitor,
		Protocol:     r.Protocol,
		Port:         r.Port,
		Notify:       r.Notify,
		BelongSystem: r.BelongSystem,
	}

	if err = CommonUpdateDomain(&domain); err != nil {
		return nil, tools.NewOperationError(fmt.Errorf("更新域名失败" + err.Error()))
	}

	return nil, nil
}
