package isql

import (
	"fmt"
	"github.com/yin-zt/go-bind-mp/pkg/model"
	"github.com/yin-zt/go-bind-mp/pkg/model/request"
	"github.com/yin-zt/go-bind-mp/pkg/util/common"
	"github.com/yin-zt/go-bind-mp/pkg/util/tools"
	"strings"
)

type ZoneService struct {
}

// List 获取数据列表
func (z ZoneService) List(req *request.ZoneListReq) ([]*model.Zone, error) {
	var list []*model.Zone
	db := common.DB.Model(&model.Zone{}).Order("created_at DESC")

	zonename := strings.TrimSpace(req.ZoneName)
	if zonename != "" {
		db = db.Where("zonename LIKE ?", fmt.Sprintf("%%%s%%", zonename))
	}

	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Preload("Domains").Preload("Views").Find(&list).Error
	return list, err
}

// Count 获取资源总数
func (z ZoneService) Count() (int64, error) {
	var count int64
	err := common.DB.Model(&model.Zone{}).Count(&count).Error
	return count, err
}

//  根据角色ID获取角色
func (z ZoneService) GetZonesByIds(zoneIds []uint) ([]*model.Zone, error) {
	var list []*model.Zone
	err := common.DB.Where("id IN (?)", zoneIds).Find(&list).Error
	return list, err
}

// Add 添加zone记录
func (z ZoneService) Add(zone *model.Zone) error {
	//result := common.DB.Create(user)
	//return user.ID, result.Error
	return common.DB.Create(zone).Error
}
