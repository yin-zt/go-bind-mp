package model

import (
	"github.com/jinzhu/gorm"
)

type View struct {
	gorm.Model
	ViewName string    `gorm:"type:varchar(100);not null;comment:'视图名称'" json:"viewname"`
	Acl      string    `gorm:"type:varchar(250);not null;comment:'访问控制列表'" json:"acl"`
	Remark   string    `gorm:"type:varchar(250);comment:'备注'" json:"remark"`
	Zones    []*Zone   `gorm:"many2many:view_zone;" json:"zones"`
	Domains  []*Domain `gorm:"many2many:domain_view;" json:"domains"`
}
