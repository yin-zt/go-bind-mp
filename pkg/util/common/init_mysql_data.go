package common

import (
	"fmt"
	"github.com/yin-zt/go-bind-mp/pkg/config"
	"github.com/yin-zt/go-bind-mp/pkg/model"

	"github.com/jinzhu/gorm"
)

// 初始化mysql数据
func InitData() {
	// 是否初始化数据
	if !config.Conf.System.InitData {
		return
	}

	// 1.写入View数据
	newViews := make([]*model.View, 0)
	views := []*model.View{
		{
			Model:    gorm.Model{ID: 1},
			ViewName: "shanghai",
			Acl:      "10.1.0.0/16",
			Remark:   "上海虹桥",
		},
		{
			Model:    gorm.Model{ID: 2},
			ViewName: "beijing",
			Acl:      "10.2.0.0/16",
			Remark:   "北京故宫",
		},
	}

	for _, view := range views {
		err := DB.First(&view, view.ID).Error
		//if errors.Is(err, gorm.ErrRecordNotFound) {
		//if err == gorm.ErrRecordNotFound {
		if err != nil {
			newViews = append(newViews, view)
		}
	}

	if len(newViews) > 0 {
		err := DB.Create(&newViews).Error
		if err != nil {
			//Log.Errorf("写入系统角色数据失败：%v", err)
			fmt.Println("写入views数据失败")
		}
	}

	// 2.写入Zone数据
	newZones := make([]*model.Zone, 0)
	zones := []*model.Zone{
		{
			Model:    gorm.Model{ID: 1},
			ZoneName: "baidu.com",
			AllowIps: "10.1.48.70,10.2.28.70",
			Views:    views[:2],
		},
		{
			Model:    gorm.Model{ID: 2},
			ZoneName: "douyin.com",
			AllowIps: "10.1.48.70",
			Views:    views[:1],
		},
	}

	for _, zone := range zones {
		err := DB.First(&zone, zone.ID).Error
		//if errors.Is(err, gorm.ErrRecordNotFound) {
		//if err == gorm.ErrRecordNotFound {
		if err != nil {
			newZones = append(newZones, zone)
		}
	}

	if len(newZones) > 0 {
		err := DB.Create(&newZones).Error
		if err != nil {
			fmt.Println("写入zones数据失败")
		}
	}

	// 3.写入Domain数据
	newDomains := make([]*model.Domain, 0)
	domains := []*model.Domain{
		{
			Model:        gorm.Model{ID: 1},
			DomainRecord: "a1.baidu.com",
			Type:         "1",
			Resolution:   "10.1.1.1",
			Status:       "1",
			Monitor:      "1",
			Protocol:     "TCP",
			Port:         80,
			Notify:       "user1",
			BelongSystem: "日炎系统",
			BelongView:   views[:1],
			BelongZone:   zones[:1],
		},
		{
			Model:        gorm.Model{ID: 2},
			DomainRecord: "a2.baidu.com",
			Type:         "1",
			Resolution:   "10.1.1.2",
			Status:       "2",
			Monitor:      "0",
			Protocol:     "HTTP",
			Port:         80,
			Notify:       "user2",
			BelongSystem: "无尽系统",
			BelongView:   views[:1],
			BelongZone:   zones[:1],
		},
		{
			Model:        gorm.Model{ID: 3},
			DomainRecord: "d1.douyin.com",
			Type:         "1",
			Resolution:   "10.2.1.1",
			Status:       "1",
			Monitor:      "1",
			Protocol:     "TCP",
			Port:         80,
			Notify:       "user3",
			BelongSystem: "抖音系统",
			BelongView:   views[:1],
			BelongZone:   zones[:1],
		},
	}

	for _, domain := range domains {
		err := DB.First(&domain, domain.ID).Error
		//if errors.Is(err, gorm.ErrRecordNotFound) {
		//if err == gorm.ErrRecordNotFound {
		if err != nil {
			newDomains = append(newDomains, domain)
		}
	}

	if len(newDomains) > 0 {
		err := DB.Create(&newDomains).Error
		if err != nil {
			fmt.Println("写入domains数据失败")
		}
	}
}
