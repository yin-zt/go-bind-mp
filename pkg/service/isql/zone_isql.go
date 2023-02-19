package isql

import (
	"errors"
	"fmt"
	"github.com/yin-zt/go-bind-mp/pkg/model"
	"github.com/yin-zt/go-bind-mp/pkg/model/request"
	"github.com/yin-zt/go-bind-mp/pkg/util/common"
	"github.com/yin-zt/go-bind-mp/pkg/util/tools"
	"gorm.io/gorm"
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

// Exist 判断zone资源是否存在
func (z ZoneService) Exist(filter map[string]interface{}) bool {
	var dataObj model.Zone
	err := common.DB.Where(filter).First(&dataObj).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// CheckRelateDomains 判断zone资源是否关联其他domains实例
func (z ZoneService) CheckRelateDomains(filter map[string]interface{}) bool {
	var dataObj model.Zone
	_ = common.DB.Where(filter).Preload("Domains").First(&dataObj).Error
	if len(dataObj.Domains) != 0 {
		return true
	}
	return false
}

// Find 获取单个资源
func (z ZoneService) Find(filter map[string]interface{}, data *model.Zone) error {
	return common.DB.Where(filter).Preload("Views").Preload("Domains").First(&data).Error
}

// Delete zone批量删除
func (z ZoneService) Delete(ids []uint) error {
	var zones []model.Zone
	for _, id := range ids {
		// 根据ID获取域名信息
		filter := tools.H{"id": id}

		zone := new(model.Zone)
		err := z.Find(filter, zone)
		if err != nil {
			return fmt.Errorf("获取zone资源失败，err: %v", err)
		}
		zones = append(zones, *zone)
	}

	err := common.DB.Debug().Select("Views", "Domains").Unscoped().Delete(&zones).Error
	if err != nil {
		return err
	}
	return err
}
