package logic

import (
	"fmt"
	"github.com/yin-zt/go-bind-mp/pkg/model"
	"github.com/yin-zt/go-bind-mp/pkg/service/isql"
	"github.com/yin-zt/go-bind-mp/pkg/util/tools"
)

var (
	ReqAssertErr = tools.NewRspError(tools.SystemErr, fmt.Errorf("请求异常"))

	Domain = &DomainLogic{}
	View   = &ViewLogic{}
	Zone   = &ZoneLogic{}
	Base   = &BaseLogic{}
)

// CommonAddDomain 标准创建域名
func CommonAddDomain(domain *model.Domain) error {

	// 先将域名记录添加到MySQL
	err := isql.Domain.Add(domain)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("向MySQL创建域名失败：" + err.Error()))
	}
	return nil
}

// CommonUpdateDomain 标准更新域名普通字段
func CommonUpdateDomain(newDomain *model.Domain) error {

	// 先将域名记录更新到MySQL
	err := isql.Domain.Update(newDomain)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("向MySQL更新域名失败：" + err.Error()))
	}
	return nil
}

// CommonAddView 标准创建视图
func CommonAddView(view *model.View) error {

	// 先将视图记录添加到MySQL
	err := isql.View.Add(view)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("向MySQL创建视图失败：" + err.Error()))
	}
	return nil
}

// CommonAddZone 标准创建zone
func CommonAddZone(zone *model.Zone) error {

	// 先将zone记录添加到MySQL
	err := isql.Zone.Add(zone)
	if err != nil {
		return tools.NewMySqlError(fmt.Errorf("向MySQL创建zone记录失败：" + err.Error()))
	}
	return nil
}
