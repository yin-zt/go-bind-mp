package isql

import (
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	"github.com/yin-zt/go-bind-mp/pkg/model"
	"github.com/yin-zt/go-bind-mp/pkg/model/request"
	"github.com/yin-zt/go-bind-mp/pkg/util/common"
	"github.com/yin-zt/go-bind-mp/pkg/util/tools"
	"gorm.io/gorm"
	"strings"
	"time"
)

type ViewService struct {
}

// 当前视图信息缓存，避免频繁获取数据库 【just test】
var viewInfoCache = cache.New(24*time.Hour, 48*time.Hour)

// List 获取数据列表
func (v ViewService) List(req *request.ViewListReq) ([]*model.View, error) {
	var list []*model.View
	db := common.DB.Model(&model.View{}).Order("created_at DESC")

	viewname := strings.TrimSpace(req.ViewName)
	if viewname != "" {
		db = db.Where("viewname LIKE ?", fmt.Sprintf("%%%s%%", viewname))
	}

	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Preload("Zones").Preload("Domains").Find(&list).Error
	return list, err
}

// Count 获取资源总数
func (v ViewService) Count() (int64, error) {
	var count int64
	err := common.DB.Model(&model.View{}).Count(&count).Error
	return count, err
}

//  根据角色ID获取角色
func (v ViewService) GetViewsByIds(viewIds []uint) ([]*model.View, error) {
	var list []*model.View
	err := common.DB.Where("id IN (?)", viewIds).Find(&list).Error
	return list, err
}

// Add 添加视图记录
func (v ViewService) Add(view *model.View) error {
	//result := common.DB.Create(user)
	//return user.ID, result.Error
	return common.DB.Create(view).Error
}

// Find 获取单个资源
func (v ViewService) Find(filter map[string]interface{}, data *model.View) error {
	return common.DB.Where(filter).Preload("Zones").Preload("Domains").First(&data).Error
}

// Exist 判断视图资源是否存在
func (v ViewService) Exist(filter map[string]interface{}) bool {
	var dataObj model.View
	err := common.DB.Where(filter).First(&dataObj).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// CheckRelateResource 判断视图资源是否关联其他资源存在
func (v ViewService) CheckRelateResource(filter map[string]interface{}) bool {
	var dataObj model.View
	_ = common.DB.Where(filter).Preload("Zones").Preload("Domains").First(&dataObj).Error
	if len(dataObj.Zones) != 0 {
		return true
	}
	if len(dataObj.Domains) != 0 {
		return true
	}
	return false
}

// Delete 视图批量删除
func (v ViewService) Delete(ids []uint) error {
	var views []model.View
	for _, id := range ids {
		// 根据ID获取域名信息
		filter := tools.H{"id": id}

		view := new(model.View)
		err := v.Find(filter, view)
		if err != nil {
			return fmt.Errorf("获取视图资源失败，err: %v", err)
		}
		views = append(views, *view)
	}

	err := common.DB.Debug().Select("Zones", "Domains").Unscoped().Delete(&views).Error
	if err != nil {
		return err
	} else {
		// 删除视图资源成功，则删除视图资源信息缓存
		for _, view := range views {
			domainInfoCache.Delete(view.ViewName)
		}
	}
	return err
}

// Update 更新视图资源
func (v ViewService) Update(view *model.View) error {
	err := common.DB.Model(view).Updates(view).Error
	if err != nil {
		return err
	}
	if view.Zones != nil {
		err = common.DB.Model(view).Association("Zones").Replace(view.Zones)
	}
	if view.Domains != nil {
		err = common.DB.Model(view).Association("Domains").Replace(view.Domains)
	}

	// 如果更新成功就更新域名信息缓存
	if err == nil {
		viewDb := &model.View{}
		common.DB.Where("view_name = ?", view.ViewName).Preload("Zones").First(&viewDb)
		viewInfoCache.Set(view.ViewName, *viewDb, cache.DefaultExpiration)
	}
	return err
}
