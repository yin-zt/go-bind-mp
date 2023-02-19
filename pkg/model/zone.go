package model

import (
	"github.com/jinzhu/gorm"
)

type Zone struct {
	gorm.Model
	ZoneName string    `gorm:"type:varchar(100);not null;comment:'zone名称'" json:"zonename"`
	AllowIps string    `gorm:"type:varchar(250);not null;comment:'允许同步IP'" json:"allowips"`
	SerialNu uint      `gorm:"type:uint;not null;default:1;comment:'序列号'" json:"serialnu"`
	Domains  []*Domain `gorm:"many2many:domain_zone"  json:"domains"`
	Views    []*View   `gorm:"many2many:view_zone;" json:"views"`
}
