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

type ZoneLogic struct{}

// List 数据列表
func (z ZoneLogic) List(c *gin.Context, req interface{}) (data interface{}, resError interface{}) {
	r, ok := req.(*request.ZoneListReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	// 获取数据列表
	zones, err := isql.Zone.List(r)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取zone列表失败: %s", err.Error()))
	}

	count, err := isql.Zone.Count()
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("获取zone总数失败"))
	}

	rets := make([]model.Zone, 0)
	for _, zone := range zones {
		rets = append(rets, *zone)
	}

	return response.ZoneListRsp{
		Total: count,
		Zones: rets,
	}, nil
}

// Add 添加Zone数据
func (z ZoneLogic) Add(c *gin.Context, req interface{}) (data interface{}, resError interface{}) {
	r, ok := req.(*request.ZoneAddReq)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	//if isql.Domain.Exist(tools.H{"doaminrecord": r.DomainRecord, "type": r.Type, "resolution": r.Resolution}) {
	//	return nil, tools.NewValidatorError(fmt.Errorf("相同的域名已存在，请勿重复添加"))
	//}

	domains, err := isql.Domain.GetDomainsByIds(r.DomainsId)
	if err != nil {
		return nil, tools.NewValidatorError(fmt.Errorf("根据Domain ID查询域名信息失败"))
	}

	views, err := isql.View.GetViewsByIds(r.ViewsId)
	if err != nil {
		return nil, tools.NewValidatorError(fmt.Errorf("根据View ID查询view信息失败"))
	}

	zone := model.Zone{
		ZoneName: r.ZoneName,
		AllowIps: r.AllowIps,
		SerialNu: r.SerialNu,
		Views:    views,
		Domains:  domains,
	}
	err = CommonAddZone(&zone)
	if err != nil {
		log.Info("添加zone记录失败")
		return nil, tools.NewOperationError(fmt.Errorf("添加zone记录失败" + err.Error()))
	}
	log.Info("添加zone记录成功")
	return nil, nil
}
