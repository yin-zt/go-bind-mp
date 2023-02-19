package model

import (
	"github.com/jinzhu/gorm"
)

type Domain struct {
	gorm.Model
	DomainRecord string  `gorm:"type:varchar(100);not null;comment:'域名记录'" json:"doaminrecord"`
	Type         string  `gorm:"type:tinyint(1);default:1;comment:'记录类型:1A, 2CNAME, 3MX'" json:"type"`
	Resolution   string  `gorm:"type:varchar(100);not null;comment:'解析目标'" json:"resolution"`
	Status       string  `gorm:"type:tinyint(1);default:1;comment:'域名状态: 1上线, 2待上线, 3待下线, 4下线'" json:"status"`
	Monitor      string  `gorm:"type:tinyint(1);default:1;comment:'是否监控: 1是, 2否'" json:"monitor_or_not"`
	Protocol     string  `gorm:"type:varchar(50);comment:'服务协议'" json:"protocol"`
	Port         uint    `gorm:"type:uint;comment:'服务端口'" json:"port"`
	Notify       string  `gorm:"type:varchar(100);comment:'监控告警接收人'" json:"notify"`
	BelongSystem string  `gorm:"type:varchar(100);comment:'所属系统'" json:"belongsystem"`
	BelongView   []*View `gorm:"many2many:domain_view;" json:"belongview"`
	BelongZone   []*Zone `gorm:"many2many:domain_zone"  json:"belongzone"`
}
