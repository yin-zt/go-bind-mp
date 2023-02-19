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

type DomainService struct {
}

// 当前域名信息缓存，避免频繁获取数据库
var domainInfoCache = cache.New(24*time.Hour, 48*time.Hour)

// List 获取数据列表
func (d DomainService) List(req *request.DomainListReq) ([]*model.Domain, error) {
	var list []*model.Domain
	db := common.DB.Model(&model.Domain{}).Order("created_at DESC")

	record := strings.TrimSpace(req.DomainRecord)
	if record != "" {
		db = db.Where("domainrecord LIKE ?", fmt.Sprintf("%%%s%%", record))
	}

	pageReq := tools.NewPageOption(req.PageNum, req.PageSize)
	err := db.Offset(pageReq.PageNum).Limit(pageReq.PageSize).Preload("BelongView").Preload("BelongZone").Find(&list).Error
	return list, err
}

// Count 获取资源总数
func (d DomainService) Count() (int64, error) {
	var count int64
	err := common.DB.Model(&model.Domain{}).Count(&count).Error
	return count, err
}

// Exist 判断资源是否存在
func (d DomainService) Exist(filter map[string]interface{}) bool {
	var dataObj model.Domain
	err := common.DB.Where(filter).First(&dataObj).Error
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

// Add 添加域名记录
func (d DomainService) Add(domain *model.Domain) error {
	//result := common.DB.Create(user)
	//return user.ID, result.Error
	return common.DB.Create(domain).Error
}

// Find 获取单个资源
func (d DomainService) Find(filter map[string]interface{}, data *model.Domain) error {
	return common.DB.Where(filter).Preload("BelongView").Preload("BelongZone").First(&data).Error
}

// Delete 域名批量删除
func (d DomainService) Delete(ids []uint) error {
	var domains []model.Domain
	for _, id := range ids {
		// 根据ID获取域名信息
		filter := tools.H{"id": id}

		domain := new(model.Domain)
		err := d.Find(filter, domain)
		if err != nil {
			return fmt.Errorf("获取域名信息失败，err: %v", err)
		}
		domains = append(domains, *domain)
	}

	err := common.DB.Debug().Select("BelongView", "BelongZone").Unscoped().Delete(&domains).Error
	if err != nil {
		return err
	} else {
		// 删除域名成功，则删除域名信息缓存
		for _, domain := range domains {
			domainInfoCache.Delete(domain.DomainRecord)
		}
	}

	// 删除域名与view的关联
	err = common.DB.Debug().Exec("DELETE FROM domain_view WHERE domain_id IN (?)", ids).Error
	if err != nil {
		return err
	}

	// 删除域名与zone的关联
	err = common.DB.Debug().Exec("DELETE FROM domain_zone WHERE domain_id IN (?)", ids).Error
	if err != nil {
		return err
	}

	return err
}

// Update 更新域名资源
func (d DomainService) Update(domain *model.Domain) error {
	err := common.DB.Model(domain).Updates(domain).Error
	if err != nil {
		return err
	}
	if domain.BelongZone != nil {
		err = common.DB.Model(Domain).Association("BelongView").Replace(domain.BelongView)
	}
	if domain.BelongZone != nil {
		err = common.DB.Model(Domain).Association("BelongZone").Replace(domain.BelongZone)
	}

	// 如果更新成功就更新域名信息缓存
	if err == nil {
		domainDb := &model.Domain{}
		common.DB.Where("domain_record = ?", domain.DomainRecord).Preload("BelongView").First(&domainDb)
		domainInfoCache.Set(domain.DomainRecord, *domainDb, cache.DefaultExpiration)
	}
	return err
}

// 根据域名ID获取域名
func (d DomainService) GetDomainsByIds(domainIds []uint) ([]*model.Domain, error) {
	var list []*model.Domain
	err := common.DB.Where("id IN (?)", domainIds).Find(&list).Error
	return list, err
}
